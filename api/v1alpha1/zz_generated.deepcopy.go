//go:build !ignore_autogenerated

// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/nutanix-cloud-native/cluster-api-runtime-extensions-nutanix/api/external/sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
	"k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMILookup) DeepCopyInto(out *AMILookup) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMILookup.
func (in *AMILookup) DeepCopy() *AMILookup {
	if in == nil {
		return nil
	}
	out := new(AMILookup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AMISpec) DeepCopyInto(out *AMISpec) {
	*out = *in
	if in.Lookup != nil {
		in, out := &in.Lookup, &out.Lookup
		*out = new(AMILookup)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AMISpec.
func (in *AMISpec) DeepCopy() *AMISpec {
	if in == nil {
		return nil
	}
	out := new(AMISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterConfig) DeepCopyInto(out *AWSClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterConfig.
func (in *AWSClusterConfig) DeepCopy() *AWSClusterConfig {
	if in == nil {
		return nil
	}
	out := new(AWSClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterConfigSpec) DeepCopyInto(out *AWSClusterConfigSpec) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSSpec)
		(*in).DeepCopyInto(*out)
	}
	in.GenericClusterConfigSpec.DeepCopyInto(&out.GenericClusterConfigSpec)
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(AWSControlPlaneNodeConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraAPIServerCertSANs != nil {
		in, out := &in.ExtraAPIServerCertSANs, &out.ExtraAPIServerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterConfigSpec.
func (in *AWSClusterConfigSpec) DeepCopy() *AWSClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AWSClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSControlPlaneNodeConfigSpec) DeepCopyInto(out *AWSControlPlaneNodeConfigSpec) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSControlPlaneNodeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSControlPlaneNodeConfigSpec.
func (in *AWSControlPlaneNodeConfigSpec) DeepCopy() *AWSControlPlaneNodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AWSControlPlaneNodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSControlPlaneNodeSpec) DeepCopyInto(out *AWSControlPlaneNodeSpec) {
	*out = *in
	in.AWSGenericNodeSpec.DeepCopyInto(&out.AWSGenericNodeSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSControlPlaneNodeSpec.
func (in *AWSControlPlaneNodeSpec) DeepCopy() *AWSControlPlaneNodeSpec {
	if in == nil {
		return nil
	}
	out := new(AWSControlPlaneNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSGenericNodeSpec) DeepCopyInto(out *AWSGenericNodeSpec) {
	*out = *in
	if in.AMISpec != nil {
		in, out := &in.AMISpec, &out.AMISpec
		*out = new(AMISpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalSecurityGroups != nil {
		in, out := &in.AdditionalSecurityGroups, &out.AdditionalSecurityGroups
		*out = make(AdditionalSecurityGroup, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSGenericNodeSpec.
func (in *AWSGenericNodeSpec) DeepCopy() *AWSGenericNodeSpec {
	if in == nil {
		return nil
	}
	out := new(AWSGenericNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSLoadBalancerSpec) DeepCopyInto(out *AWSLoadBalancerSpec) {
	*out = *in
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(v1beta2.ELBScheme)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSLoadBalancerSpec.
func (in *AWSLoadBalancerSpec) DeepCopy() *AWSLoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(AWSLoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSNetwork) DeepCopyInto(out *AWSNetwork) {
	*out = *in
	if in.VPC != nil {
		in, out := &in.VPC, &out.VPC
		*out = new(VPC)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(Subnets, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSNetwork.
func (in *AWSNetwork) DeepCopy() *AWSNetwork {
	if in == nil {
		return nil
	}
	out := new(AWSNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSpec) DeepCopyInto(out *AWSSpec) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(Region)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(AWSNetwork)
		(*in).DeepCopyInto(*out)
	}
	if in.ControlPlaneLoadBalancer != nil {
		in, out := &in.ControlPlaneLoadBalancer, &out.ControlPlaneLoadBalancer
		*out = new(AWSLoadBalancerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSpec.
func (in *AWSSpec) DeepCopy() *AWSSpec {
	if in == nil {
		return nil
	}
	out := new(AWSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSWorkerNodeConfig) DeepCopyInto(out *AWSWorkerNodeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSWorkerNodeConfig.
func (in *AWSWorkerNodeConfig) DeepCopy() *AWSWorkerNodeConfig {
	if in == nil {
		return nil
	}
	out := new(AWSWorkerNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSWorkerNodeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSWorkerNodeConfigSpec) DeepCopyInto(out *AWSWorkerNodeConfigSpec) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSWorkerNodeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSWorkerNodeConfigSpec.
func (in *AWSWorkerNodeConfigSpec) DeepCopy() *AWSWorkerNodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AWSWorkerNodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSWorkerNodeSpec) DeepCopyInto(out *AWSWorkerNodeSpec) {
	*out = *in
	in.AWSGenericNodeSpec.DeepCopyInto(&out.AWSGenericNodeSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSWorkerNodeSpec.
func (in *AWSWorkerNodeSpec) DeepCopy() *AWSWorkerNodeSpec {
	if in == nil {
		return nil
	}
	out := new(AWSWorkerNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AdditionalSecurityGroup) DeepCopyInto(out *AdditionalSecurityGroup) {
	{
		in := &in
		*out = make(AdditionalSecurityGroup, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalSecurityGroup.
func (in AdditionalSecurityGroup) DeepCopy() AdditionalSecurityGroup {
	if in == nil {
		return nil
	}
	out := new(AdditionalSecurityGroup)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addons) DeepCopyInto(out *Addons) {
	*out = *in
	if in.CNI != nil {
		in, out := &in.CNI, &out.CNI
		*out = new(CNI)
		**out = **in
	}
	if in.NFD != nil {
		in, out := &in.NFD, &out.NFD
		*out = new(NFD)
		**out = **in
	}
	if in.ClusterAutoscaler != nil {
		in, out := &in.ClusterAutoscaler, &out.ClusterAutoscaler
		*out = new(ClusterAutoscaler)
		**out = **in
	}
	if in.CCM != nil {
		in, out := &in.CCM, &out.CCM
		*out = new(CCM)
		(*in).DeepCopyInto(*out)
	}
	if in.CSIProviders != nil {
		in, out := &in.CSIProviders, &out.CSIProviders
		*out = new(CSI)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceLoadBalancer != nil {
		in, out := &in.ServiceLoadBalancer, &out.ServiceLoadBalancer
		*out = new(ServiceLoadBalancer)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addons.
func (in *Addons) DeepCopy() *Addons {
	if in == nil {
		return nil
	}
	out := new(Addons)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CCM) DeepCopyInto(out *CCM) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CCM.
func (in *CCM) DeepCopy() *CCM {
	if in == nil {
		return nil
	}
	out := new(CCM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNI) DeepCopyInto(out *CNI) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNI.
func (in *CNI) DeepCopy() *CNI {
	if in == nil {
		return nil
	}
	out := new(CNI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSI) DeepCopyInto(out *CSI) {
	*out = *in
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]CSIProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.DefaultStorage = in.DefaultStorage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSI.
func (in *CSI) DeepCopy() *CSI {
	if in == nil {
		return nil
	}
	out := new(CSI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIProvider) DeepCopyInto(out *CSIProvider) {
	*out = *in
	if in.StorageClassConfig != nil {
		in, out := &in.StorageClassConfig, &out.StorageClassConfig
		*out = make([]StorageClassConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIProvider.
func (in *CSIProvider) DeepCopy() *CSIProvider {
	if in == nil {
		return nil
	}
	out := new(CSIProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAutoscaler) DeepCopyInto(out *ClusterAutoscaler) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAutoscaler.
func (in *ClusterAutoscaler) DeepCopy() *ClusterAutoscaler {
	if in == nil {
		return nil
	}
	out := new(ClusterAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneEndpointSpec) DeepCopyInto(out *ControlPlaneEndpointSpec) {
	*out = *in
	if in.VirtualIPSpec != nil {
		in, out := &in.VirtualIPSpec, &out.VirtualIPSpec
		*out = new(ControlPlaneVirtualIPSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneEndpointSpec.
func (in *ControlPlaneEndpointSpec) DeepCopy() *ControlPlaneEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneVirtualIPSpec) DeepCopyInto(out *ControlPlaneVirtualIPSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneVirtualIPSpec.
func (in *ControlPlaneVirtualIPSpec) DeepCopy() *ControlPlaneVirtualIPSpec {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneVirtualIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultStorage) DeepCopyInto(out *DefaultStorage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultStorage.
func (in *DefaultStorage) DeepCopy() *DefaultStorage {
	if in == nil {
		return nil
	}
	out := new(DefaultStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerClusterConfig) DeepCopyInto(out *DockerClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerClusterConfig.
func (in *DockerClusterConfig) DeepCopy() *DockerClusterConfig {
	if in == nil {
		return nil
	}
	out := new(DockerClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerClusterConfigSpec) DeepCopyInto(out *DockerClusterConfigSpec) {
	*out = *in
	if in.Docker != nil {
		in, out := &in.Docker, &out.Docker
		*out = new(DockerSpec)
		**out = **in
	}
	in.GenericClusterConfigSpec.DeepCopyInto(&out.GenericClusterConfigSpec)
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(DockerNodeConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraAPIServerCertSANs != nil {
		in, out := &in.ExtraAPIServerCertSANs, &out.ExtraAPIServerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerClusterConfigSpec.
func (in *DockerClusterConfigSpec) DeepCopy() *DockerClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DockerClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerNodeConfig) DeepCopyInto(out *DockerNodeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerNodeConfig.
func (in *DockerNodeConfig) DeepCopy() *DockerNodeConfig {
	if in == nil {
		return nil
	}
	out := new(DockerNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerNodeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerNodeConfigSpec) DeepCopyInto(out *DockerNodeConfigSpec) {
	*out = *in
	if in.Docker != nil {
		in, out := &in.Docker, &out.Docker
		*out = new(DockerNodeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerNodeConfigSpec.
func (in *DockerNodeConfigSpec) DeepCopy() *DockerNodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DockerNodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerNodeSpec) DeepCopyInto(out *DockerNodeSpec) {
	*out = *in
	if in.CustomImage != nil {
		in, out := &in.CustomImage, &out.CustomImage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerNodeSpec.
func (in *DockerNodeSpec) DeepCopy() *DockerNodeSpec {
	if in == nil {
		return nil
	}
	out := new(DockerNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerSpec) DeepCopyInto(out *DockerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerSpec.
func (in *DockerSpec) DeepCopy() *DockerSpec {
	if in == nil {
		return nil
	}
	out := new(DockerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Etcd) DeepCopyInto(out *Etcd) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(Image)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Etcd.
func (in *Etcd) DeepCopy() *Etcd {
	if in == nil {
		return nil
	}
	out := new(Etcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericClusterConfigSpec) DeepCopyInto(out *GenericClusterConfigSpec) {
	*out = *in
	if in.KubernetesImageRepository != nil {
		in, out := &in.KubernetesImageRepository, &out.KubernetesImageRepository
		*out = new(string)
		**out = **in
	}
	if in.Etcd != nil {
		in, out := &in.Etcd, &out.Etcd
		*out = new(Etcd)
		(*in).DeepCopyInto(*out)
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(HTTPProxy)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageRegistries != nil {
		in, out := &in.ImageRegistries, &out.ImageRegistries
		*out = make([]ImageRegistry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GlobalImageRegistryMirror != nil {
		in, out := &in.GlobalImageRegistryMirror, &out.GlobalImageRegistryMirror
		*out = new(GlobalImageRegistryMirror)
		(*in).DeepCopyInto(*out)
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = new(Addons)
		(*in).DeepCopyInto(*out)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericClusterConfigSpec.
func (in *GenericClusterConfigSpec) DeepCopy() *GenericClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(GenericClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalImageRegistryMirror) DeepCopyInto(out *GlobalImageRegistryMirror) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(RegistryCredentials)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalImageRegistryMirror.
func (in *GlobalImageRegistryMirror) DeepCopy() *GlobalImageRegistryMirror {
	if in == nil {
		return nil
	}
	out := new(GlobalImageRegistryMirror)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProxy) DeepCopyInto(out *HTTPProxy) {
	*out = *in
	if in.AdditionalNo != nil {
		in, out := &in.AdditionalNo, &out.AdditionalNo
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProxy.
func (in *HTTPProxy) DeepCopy() *HTTPProxy {
	if in == nil {
		return nil
	}
	out := new(HTTPProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistry) DeepCopyInto(out *ImageRegistry) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(RegistryCredentials)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistry.
func (in *ImageRegistry) DeepCopy() *ImageRegistry {
	if in == nil {
		return nil
	}
	out := new(ImageRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFD) DeepCopyInto(out *NFD) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFD.
func (in *NFD) DeepCopy() *NFD {
	if in == nil {
		return nil
	}
	out := new(NFD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixCategoryIdentifier) DeepCopyInto(out *NutanixCategoryIdentifier) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixCategoryIdentifier.
func (in *NutanixCategoryIdentifier) DeepCopy() *NutanixCategoryIdentifier {
	if in == nil {
		return nil
	}
	out := new(NutanixCategoryIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixClusterConfig) DeepCopyInto(out *NutanixClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixClusterConfig.
func (in *NutanixClusterConfig) DeepCopy() *NutanixClusterConfig {
	if in == nil {
		return nil
	}
	out := new(NutanixClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NutanixClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixClusterConfigSpec) DeepCopyInto(out *NutanixClusterConfigSpec) {
	*out = *in
	if in.Nutanix != nil {
		in, out := &in.Nutanix, &out.Nutanix
		*out = new(NutanixSpec)
		(*in).DeepCopyInto(*out)
	}
	in.GenericClusterConfigSpec.DeepCopyInto(&out.GenericClusterConfigSpec)
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(NutanixNodeConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraAPIServerCertSANs != nil {
		in, out := &in.ExtraAPIServerCertSANs, &out.ExtraAPIServerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixClusterConfigSpec.
func (in *NutanixClusterConfigSpec) DeepCopy() *NutanixClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NutanixClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixMachineDetails) DeepCopyInto(out *NutanixMachineDetails) {
	*out = *in
	out.MemorySize = in.MemorySize.DeepCopy()
	in.Image.DeepCopyInto(&out.Image)
	in.Cluster.DeepCopyInto(&out.Cluster)
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]NutanixResourceIdentifier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalCategories != nil {
		in, out := &in.AdditionalCategories, &out.AdditionalCategories
		*out = make([]NutanixCategoryIdentifier, len(*in))
		copy(*out, *in)
	}
	out.SystemDiskSize = in.SystemDiskSize.DeepCopy()
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(NutanixResourceIdentifier)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixMachineDetails.
func (in *NutanixMachineDetails) DeepCopy() *NutanixMachineDetails {
	if in == nil {
		return nil
	}
	out := new(NutanixMachineDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixNodeConfig) DeepCopyInto(out *NutanixNodeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixNodeConfig.
func (in *NutanixNodeConfig) DeepCopy() *NutanixNodeConfig {
	if in == nil {
		return nil
	}
	out := new(NutanixNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NutanixNodeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixNodeConfigSpec) DeepCopyInto(out *NutanixNodeConfigSpec) {
	*out = *in
	if in.Nutanix != nil {
		in, out := &in.Nutanix, &out.Nutanix
		*out = new(NutanixNodeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixNodeConfigSpec.
func (in *NutanixNodeConfigSpec) DeepCopy() *NutanixNodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NutanixNodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixNodeSpec) DeepCopyInto(out *NutanixNodeSpec) {
	*out = *in
	in.MachineDetails.DeepCopyInto(&out.MachineDetails)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixNodeSpec.
func (in *NutanixNodeSpec) DeepCopy() *NutanixNodeSpec {
	if in == nil {
		return nil
	}
	out := new(NutanixNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixPrismCentralEndpointSpec) DeepCopyInto(out *NutanixPrismCentralEndpointSpec) {
	*out = *in
	out.Credentials = in.Credentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixPrismCentralEndpointSpec.
func (in *NutanixPrismCentralEndpointSpec) DeepCopy() *NutanixPrismCentralEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(NutanixPrismCentralEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixResourceIdentifier) DeepCopyInto(out *NutanixResourceIdentifier) {
	*out = *in
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixResourceIdentifier.
func (in *NutanixResourceIdentifier) DeepCopy() *NutanixResourceIdentifier {
	if in == nil {
		return nil
	}
	out := new(NutanixResourceIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixSpec) DeepCopyInto(out *NutanixSpec) {
	*out = *in
	in.ControlPlaneEndpoint.DeepCopyInto(&out.ControlPlaneEndpoint)
	out.PrismCentralEndpoint = in.PrismCentralEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixSpec.
func (in *NutanixSpec) DeepCopy() *NutanixSpec {
	if in == nil {
		return nil
	}
	out := new(NutanixSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryCredentials) DeepCopyInto(out *RegistryCredentials) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryCredentials.
func (in *RegistryCredentials) DeepCopy() *RegistryCredentials {
	if in == nil {
		return nil
	}
	out := new(RegistryCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLoadBalancer) DeepCopyInto(out *ServiceLoadBalancer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLoadBalancer.
func (in *ServiceLoadBalancer) DeepCopy() *ServiceLoadBalancer {
	if in == nil {
		return nil
	}
	out := new(ServiceLoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassConfig) DeepCopyInto(out *StorageClassConfig) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReclaimPolicy != nil {
		in, out := &in.ReclaimPolicy, &out.ReclaimPolicy
		*out = new(v1.PersistentVolumeReclaimPolicy)
		**out = **in
	}
	if in.VolumeBindingMode != nil {
		in, out := &in.VolumeBindingMode, &out.VolumeBindingMode
		*out = new(storagev1.VolumeBindingMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassConfig.
func (in *StorageClassConfig) DeepCopy() *StorageClassConfig {
	if in == nil {
		return nil
	}
	out := new(StorageClassConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Subnets) DeepCopyInto(out *Subnets) {
	{
		in := &in
		*out = make(Subnets, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnets.
func (in Subnets) DeepCopy() Subnets {
	if in == nil {
		return nil
	}
	out := new(Subnets)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	if in.SSHAuthorizedKeys != nil {
		in, out := &in.SSHAuthorizedKeys, &out.SSHAuthorizedKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPC) DeepCopyInto(out *VPC) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPC.
func (in *VPC) DeepCopy() *VPC {
	if in == nil {
		return nil
	}
	out := new(VPC)
	in.DeepCopyInto(out)
	return out
}
