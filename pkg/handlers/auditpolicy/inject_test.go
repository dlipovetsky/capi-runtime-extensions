// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package auditpolicy

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/d2iq-labs/capi-runtime-extensions/common/pkg/testutils/capitest"
)

func TestGeneratePatches(t *testing.T) {
	capitest.ValidateGeneratePatches(
		t,
		NewPatch,
		capitest.PatchTestDef{
			Name: "unset variable",
		},
		capitest.PatchTestDef{
			Name:        "http proxy set for KubeadmControlPlaneTemplate",
			RequestItem: capitest.NewKubeadmControlPlaneTemplateRequestItem(),
			ExpectedPatchMatchers: []capitest.JSONPatchMatcher{{
				Operation:    "add",
				Path:         "/spec/template/spec/kubeadmConfigSpec/files",
				ValueMatcher: HaveLen(1),
			}, {
				Operation: "add",
				Path:      "/spec/template/spec/kubeadmConfigSpec/clusterConfiguration",
				ValueMatcher: HaveKeyWithValue(
					"apiServer",
					SatisfyAll(
						HaveKeyWithValue(
							"extraArgs",
							map[string]interface{}{
								"audit-log-maxbackup": "10",
								"audit-log-maxsize":   "100",
								"audit-log-path":      "/var/log/audit/kube-apiserver-audit.log",
								"audit-policy-file":   "/etc/kubernetes/audit-policy/apiserver-audit-policy.yaml",
								"audit-log-maxage":    "30",
							},
						),
						HaveKeyWithValue(
							"extraVolumes",
							[]interface{}{
								map[string]interface{}{
									"hostPath":  "/etc/kubernetes/audit-policy/",
									"mountPath": "/etc/kubernetes/audit-policy/",
									"name":      "audit-policy",
									"readOnly":  true,
								},
								map[string]interface{}{
									"name":      "audit-logs",
									"hostPath":  "/var/log/kubernetes/audit",
									"mountPath": "/var/log/audit/",
								},
							},
						),
					),
				),
			}},
		},
	)
}