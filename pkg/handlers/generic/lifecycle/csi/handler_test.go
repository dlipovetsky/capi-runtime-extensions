// Copyright 2024 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package csi

import (
	"context"
	"testing"

	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	runtimehooksv1 "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/nutanix-cloud-native/cluster-api-runtime-extensions-nutanix/api/v1alpha1"
	"github.com/nutanix-cloud-native/cluster-api-runtime-extensions-nutanix/api/variables"
)

type fakeCSIProvider struct{}

func (n *fakeCSIProvider) Apply(
	ctx context.Context,
	provider v1alpha1.CSIProvider,
	defaultStorageConfig *v1alpha1.DefaultStorage,
	req *runtimehooksv1.AfterControlPlaneInitializedRequest,
	log logr.Logger,
) error {
	return nil
}

var testProviderHandlers = map[string]CSIProvider{
	"test1": &fakeCSIProvider{},
	"test2": &fakeCSIProvider{},
}

func testReq(csi *v1alpha1.CSI) (*runtimehooksv1.AfterControlPlaneInitializedRequest, error) {
	cv, err := variables.MarshalToClusterVariable(
		"clusterConfig",
		&v1alpha1.GenericClusterConfigSpec{
			Addons: &v1alpha1.Addons{
				CSIProviders: csi,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &runtimehooksv1.AfterControlPlaneInitializedRequest{
		Cluster: clusterv1.Cluster{
			Spec: clusterv1.ClusterSpec{
				Topology: &clusterv1.Topology{
					Variables: []clusterv1.ClusterVariable{
						*cv,
					},
				},
			},
		},
	}, nil
}

func Test_AfterControlPlaneInitialized(t *testing.T) {
	tests := []struct {
		name       string
		csi        *v1alpha1.CSI
		wantStatus runtimehooksv1.ResponseStatus
	}{
		{
			name:       "csi variable is optional",
			csi:        nil,
			wantStatus: runtimehooksv1.ResponseStatusSuccess,
		},
		{
			name: "if csi variable set, must specify one or more providers",
			csi: &v1alpha1.CSI{
				Providers: []v1alpha1.CSIProvider{},
			},
			wantStatus: runtimehooksv1.ResponseStatusFailure,
		},
		{
			name: "if csi variable set, a provider must have a storage class config",
			csi: &v1alpha1.CSI{
				Providers: []v1alpha1.CSIProvider{
					{
						Name:               "test1",
						StorageClassConfig: []v1alpha1.StorageClassConfig{},
					},
				},
			},
			wantStatus: runtimehooksv1.ResponseStatusFailure,
		},
		{
			name: "if csi variable set with one provider, we derive the default storage",
			csi: &v1alpha1.CSI{
				Providers: []v1alpha1.CSIProvider{
					{
						Name: "test1",
						StorageClassConfig: []v1alpha1.StorageClassConfig{
							{
								Name: "example",
							},
						},
					},
				},
			},
			wantStatus: runtimehooksv1.ResponseStatusSuccess,
		},
		{
			name: "if csi variable set with two or more providers, user must set default storage",
			csi: &v1alpha1.CSI{
				Providers: []v1alpha1.CSIProvider{
					{
						Name: "test1",
						StorageClassConfig: []v1alpha1.StorageClassConfig{
							{
								Name: "example",
							},
						},
					},
					{
						Name: "test2",
						StorageClassConfig: []v1alpha1.StorageClassConfig{
							{
								Name: "example",
							},
						},
					},
				},
			},
			wantStatus: runtimehooksv1.ResponseStatusFailure,
		},
		{
			name: "if csi variable set with two providers and default storage",
			csi: &v1alpha1.CSI{
				Providers: []v1alpha1.CSIProvider{
					{
						Name: "test1",
						StorageClassConfig: []v1alpha1.StorageClassConfig{
							{
								Name: "example",
							},
						},
					},
					{
						Name: "test2",
						StorageClassConfig: []v1alpha1.StorageClassConfig{
							{
								Name: "example",
							},
						},
					},
				},
				DefaultStorage: &v1alpha1.DefaultStorage{
					ProviderName:           "test1",
					StorageClassConfigName: "example",
				},
			},
			wantStatus: runtimehooksv1.ResponseStatusSuccess,
		},
		{
			name: "csi provider is unknown",
			csi: &v1alpha1.CSI{
				Providers: []v1alpha1.CSIProvider{
					{
						Name: "not-test1-or-test2",
						StorageClassConfig: []v1alpha1.StorageClassConfig{
							{
								Name: "example",
							},
						},
					},
				},
			},
			wantStatus: runtimehooksv1.ResponseStatusFailure,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			client := fake.NewClientBuilder().Build()
			handler := New(client, testProviderHandlers)
			resp := &runtimehooksv1.AfterControlPlaneInitializedResponse{}

			req, err := testReq(tt.csi)
			if err != nil {
				t.Fatalf("failed to create test request: %s", err)
			}

			handler.AfterControlPlaneInitialized(ctx, req, resp)
			if diff := cmp.Diff(tt.wantStatus, resp.Status); diff != "" {
				t.Errorf("response Status mismatch (-want +got):\n%s. Message: %s", diff, resp.Message)
			}
		})
	}
}