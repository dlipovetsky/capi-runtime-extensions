// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package mirrors

import (
	"context"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	bootstrapv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1beta1"
	controlplanev1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1beta1"
	runtimehooksv1 "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/api/v1alpha1"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/common/pkg/capi/clustertopology/patches"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/common/pkg/capi/clustertopology/patches/selectors"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/common/pkg/capi/clustertopology/variables"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/clusterconfig"
)

type globalMirrorPatchHandler struct {
	client ctrlclient.Client

	variableName      string
	variableFieldPath []string
}

func NewPatch(
	cl ctrlclient.Client,
) *globalMirrorPatchHandler {
	return newGlobalMirrorPatchHandler(
		cl,
		clusterconfig.MetaVariableName,
		GlobalMirrorVariableName,
	)
}

func newGlobalMirrorPatchHandler(
	cl ctrlclient.Client,
	variableName string,
	variableFieldPath ...string,
) *globalMirrorPatchHandler {
	scheme := runtime.NewScheme()
	_ = bootstrapv1.AddToScheme(scheme)
	_ = controlplanev1.AddToScheme(scheme)
	return &globalMirrorPatchHandler{
		client:            cl,
		variableName:      variableName,
		variableFieldPath: variableFieldPath,
	}
}

func (h *globalMirrorPatchHandler) Mutate(
	ctx context.Context,
	obj *unstructured.Unstructured,
	vars map[string]apiextensionsv1.JSON,
	holderRef runtimehooksv1.HolderReference,
	clusterKey ctrlclient.ObjectKey,
) error {
	log := ctrl.LoggerFrom(ctx).WithValues(
		"holderRef", holderRef,
	)

	globalMirror, found, err := variables.Get[v1alpha1.GlobalImageRegistryMirror](
		vars,
		h.variableName,
		h.variableFieldPath...,
	)
	if err != nil {
		return err
	}
	if !found {
		log.V(5).Info("Global registry mirror variable not defined")
		return nil
	}

	log = log.WithValues(
		"variableName",
		h.variableName,
		"variableFieldPath",
		h.variableFieldPath,
		"variableValue",
		globalMirror,
	)

	if err := patches.MutateIfApplicable(
		obj, vars, &holderRef, selectors.ControlPlane(), log,
		func(obj *controlplanev1.KubeadmControlPlaneTemplate) error {
			mirrorConfig, err := mirrorConfigForGlobalMirror(
				ctx,
				h.client,
				globalMirror,
				obj,
			)
			if err != nil {
				return err
			}
			files, commands, generateErr := generateFilesAndCommands(
				mirrorConfig,
				globalMirror)
			if generateErr != nil {
				return generateErr
			}

			log.WithValues(
				"patchedObjectKind", obj.GetObjectKind().GroupVersionKind().String(),
				"patchedObjectName", ctrlclient.ObjectKeyFromObject(obj),
			).Info("adding global registry mirror files to control plane kubeadm config spec")
			obj.Spec.Template.Spec.KubeadmConfigSpec.Files = append(
				obj.Spec.Template.Spec.KubeadmConfigSpec.Files,
				files...,
			)

			log.WithValues(
				"patchedObjectKind", obj.GetObjectKind().GroupVersionKind().String(),
				"patchedObjectName", ctrlclient.ObjectKeyFromObject(obj),
			).Info("adding PreKubeadmCommands to control plane kubeadm config spec")
			obj.Spec.Template.Spec.KubeadmConfigSpec.PreKubeadmCommands = append(
				obj.Spec.Template.Spec.KubeadmConfigSpec.PreKubeadmCommands,
				commands...,
			)

			return nil
		}); err != nil {
		return err
	}

	if err := patches.MutateIfApplicable(
		obj, vars, &holderRef, selectors.WorkersKubeadmConfigTemplateSelector(), log,
		func(obj *bootstrapv1.KubeadmConfigTemplate) error {
			mirrorConfig, err := mirrorConfigForGlobalMirror(
				ctx,
				h.client,
				globalMirror,
				obj,
			)
			if err != nil {
				return err
			}
			files, commands, generateErr := generateFilesAndCommands(
				mirrorConfig,
				globalMirror)
			if generateErr != nil {
				return generateErr
			}

			log.WithValues(
				"patchedObjectKind", obj.GetObjectKind().GroupVersionKind().String(),
				"patchedObjectName", ctrlclient.ObjectKeyFromObject(obj),
			).Info("adding global registry mirror files to worker node kubeadm config template")
			obj.Spec.Template.Spec.Files = append(obj.Spec.Template.Spec.Files, files...)

			log.WithValues(
				"patchedObjectKind", obj.GetObjectKind().GroupVersionKind().String(),
				"patchedObjectName", ctrlclient.ObjectKeyFromObject(obj),
			).Info("adding PreKubeadmCommands to worker node kubeadm config template")
			obj.Spec.Template.Spec.PreKubeadmCommands = append(obj.Spec.Template.Spec.PreKubeadmCommands, commands...)

			return nil
		}); err != nil {
		return err
	}

	return nil
}

func generateFilesAndCommands(
	mirrorConfig *mirrorConfig,
	globalMirror v1alpha1.GlobalImageRegistryMirror,
) ([]bootstrapv1.File, []string, error) {
	// generate default registry mirror file
	files, err := generateGlobalRegistryMirrorFile(mirrorConfig)
	if err != nil {
		return nil, nil, err
	}
	// generate CA certificate file for registry mirror
	mirrorCAFile := generateMirrorCACertFile(mirrorConfig, globalMirror)
	files = append(files, mirrorCAFile...)
	// generate Containerd registry config drop-in file
	registryConfigDropIn := generateContainerdRegistryConfigDropInFile()
	files = append(files, registryConfigDropIn...)
	// generate Containerd apply patches script and command
	var commands []string
	applyPatchesFile, applyPatchesCommand, err := generateContainerdApplyPatchesScript()
	if err != nil {
		return nil, nil, err
	}
	files = append(files, applyPatchesFile...)
	commands = append(commands, applyPatchesCommand)

	return files, commands, err
}
