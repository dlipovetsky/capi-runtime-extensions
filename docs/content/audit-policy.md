---
title: "Audit policy"
---

Kubernetes auditing provides a security-relevant, chronological set of records documenting the sequence of actions in a
cluster. The cluster audits the activities generated by users, by applications that use the Kubernetes API, and by the
control plane itself. The `auditpolicypatch` external patch will generate appropriate configuration for the Kubernetes
control plane.

To enable the audit policy enable the `auditpolicypatch` external patch on `ClusterClass`.

```yaml
apiVersion: cluster.x-k8s.io/v1beta1
kind: ClusterClass
metadata:
  name: <NAME>
spec:
  patches:
    - name: audit-policy
      external:
        generateExtension: "auditpolicypatch.<external-config-name>"
```

Applying this configuration will result in new bootstrap files on the `KubeadmControlPlaneTemplate`.