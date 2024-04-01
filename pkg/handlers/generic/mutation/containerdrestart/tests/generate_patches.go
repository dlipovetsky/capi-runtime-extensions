// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package tests

import (
	"testing"

	"github.com/onsi/gomega"
	runtimehooksv1 "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1"

	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/common/pkg/capi/clustertopology/handlers/mutation"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/common/pkg/testutils/capitest"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/common/pkg/testutils/capitest/request"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/mutation/containerdrestart"
)

func TestGeneratePatches(
	t *testing.T,
	generatorFunc func() mutation.GeneratePatches,
) {
	t.Helper()

	capitest.ValidateGeneratePatches(
		t,
		generatorFunc,
		capitest.PatchTestDef{
			Name:        "restart script and command added to control plane kubeadm config spec",
			RequestItem: request.NewKubeadmControlPlaneTemplateRequestItem(""),
			ExpectedPatchMatchers: []capitest.JSONPatchMatcher{
				{
					Operation: "add",
					Path:      "/spec/template/spec/kubeadmConfigSpec/files",
					ValueMatcher: gomega.ContainElements(
						gomega.HaveKeyWithValue(
							"path", containerdrestart.ContainerdRestartScriptOnRemote,
						),
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/kubeadmConfigSpec/preKubeadmCommands",
					ValueMatcher: gomega.ContainElements(
						containerdrestart.ContainerdRestartScriptOnRemoteCommand,
					),
				},
			},
		},
		capitest.PatchTestDef{
			Name: "restart script and command added to worker node kubeadm config template",
			Vars: []runtimehooksv1.Variable{
				capitest.VariableWithValue(
					"builtin",
					map[string]any{
						"machineDeployment": map[string]any{
							"class": "*",
						},
					},
				),
			},
			RequestItem: request.NewKubeadmConfigTemplateRequestItem(""),
			ExpectedPatchMatchers: []capitest.JSONPatchMatcher{
				{
					Operation: "add",
					Path:      "/spec/template/spec/files",
					ValueMatcher: gomega.ContainElements(
						gomega.HaveKeyWithValue(
							"path", containerdrestart.ContainerdRestartScriptOnRemote,
						),
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/preKubeadmCommands",
					ValueMatcher: gomega.ContainElements(
						containerdrestart.ContainerdRestartScriptOnRemoteCommand,
					),
				},
			},
		},
	)
}
