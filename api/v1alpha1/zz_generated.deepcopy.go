//go:build !ignore_autogenerated

// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/api/external/sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
	"k8s.io/api/core/v1"
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
func (in *AWSNodeSpec) DeepCopyInto(out *AWSNodeSpec) {
	*out = *in
	if in.IAMInstanceProfile != nil {
		in, out := &in.IAMInstanceProfile, &out.IAMInstanceProfile
		*out = new(IAMInstanceProfile)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(InstanceType)
		**out = **in
	}
	if in.AMISpec != nil {
		in, out := &in.AMISpec, &out.AMISpec
		*out = new(AMISpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalSecurityGroups != nil {
		in, out := &in.AdditionalSecurityGroups, &out.AdditionalSecurityGroups
		*out = make(AdditionalSecurityGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSNodeSpec.
func (in *AWSNodeSpec) DeepCopy() *AWSNodeSpec {
	if in == nil {
		return nil
	}
	out := new(AWSNodeSpec)
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
func (in AdditionalSecurityGroup) DeepCopyInto(out *AdditionalSecurityGroup) {
	{
		in := &in
		*out = make(AdditionalSecurityGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
		*out = new(v1.LocalObjectReference)
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
	if in.DefaultStorage != nil {
		in, out := &in.DefaultStorage, &out.DefaultStorage
		*out = new(DefaultStorage)
		**out = **in
	}
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
		*out = new(v1.LocalObjectReference)
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
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigSpec) DeepCopyInto(out *ClusterConfigSpec) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Docker != nil {
		in, out := &in.Docker, &out.Docker
		*out = new(DockerSpec)
		**out = **in
	}
	if in.Nutanix != nil {
		in, out := &in.Nutanix, &out.Nutanix
		*out = new(NutanixSpec)
		(*in).DeepCopyInto(*out)
	}
	in.GenericClusterConfig.DeepCopyInto(&out.GenericClusterConfig)
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(NodeConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigSpec.
func (in *ClusterConfigSpec) DeepCopy() *ClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneEndpointSpec) DeepCopyInto(out *ControlPlaneEndpointSpec) {
	*out = *in
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
func (in *DockerNodeSpec) DeepCopyInto(out *DockerNodeSpec) {
	*out = *in
	if in.CustomImage != nil {
		in, out := &in.CustomImage, &out.CustomImage
		*out = new(OCIImage)
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
func (in ExtraAPIServerCertSANs) DeepCopyInto(out *ExtraAPIServerCertSANs) {
	{
		in := &in
		*out = make(ExtraAPIServerCertSANs, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraAPIServerCertSANs.
func (in ExtraAPIServerCertSANs) DeepCopy() ExtraAPIServerCertSANs {
	if in == nil {
		return nil
	}
	out := new(ExtraAPIServerCertSANs)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericClusterConfig) DeepCopyInto(out *GenericClusterConfig) {
	*out = *in
	if in.KubernetesImageRepository != nil {
		in, out := &in.KubernetesImageRepository, &out.KubernetesImageRepository
		*out = new(KubernetesImageRepository)
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
	if in.ExtraAPIServerCertSANs != nil {
		in, out := &in.ExtraAPIServerCertSANs, &out.ExtraAPIServerCertSANs
		*out = make(ExtraAPIServerCertSANs, len(*in))
		copy(*out, *in)
	}
	if in.ImageRegistries != nil {
		in, out := &in.ImageRegistries, &out.ImageRegistries
		*out = make(ImageRegistries, len(*in))
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
		*out = make(Users, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericClusterConfig.
func (in *GenericClusterConfig) DeepCopy() *GenericClusterConfig {
	if in == nil {
		return nil
	}
	out := new(GenericClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericNodeConfig) DeepCopyInto(out *GenericNodeConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericNodeConfig.
func (in *GenericNodeConfig) DeepCopy() *GenericNodeConfig {
	if in == nil {
		return nil
	}
	out := new(GenericNodeConfig)
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
func (in ImageRegistries) DeepCopyInto(out *ImageRegistries) {
	{
		in := &in
		*out = make(ImageRegistries, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistries.
func (in ImageRegistries) DeepCopy() ImageRegistries {
	if in == nil {
		return nil
	}
	out := new(ImageRegistries)
	in.DeepCopyInto(out)
	return *out
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
func (in *NodeConfig) DeepCopyInto(out *NodeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfig.
func (in *NodeConfig) DeepCopy() *NodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigSpec) DeepCopyInto(out *NodeConfigSpec) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSNodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Docker != nil {
		in, out := &in.Docker, &out.Docker
		*out = new(DockerNodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Nutanix != nil {
		in, out := &in.Nutanix, &out.Nutanix
		*out = new(NutanixNodeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigSpec.
func (in *NodeConfigSpec) DeepCopy() *NodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NodeConfigSpec)
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
		*out = make(NutanixResourceIdentifiers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.SystemDiskSize = in.SystemDiskSize.DeepCopy()
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
	if in.AdditionalTrustBundle != nil {
		in, out := &in.AdditionalTrustBundle, &out.AdditionalTrustBundle
		*out = new(string)
		**out = **in
	}
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
func (in NutanixResourceIdentifiers) DeepCopyInto(out *NutanixResourceIdentifiers) {
	{
		in := &in
		*out = make(NutanixResourceIdentifiers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NutanixResourceIdentifiers.
func (in NutanixResourceIdentifiers) DeepCopy() NutanixResourceIdentifiers {
	if in == nil {
		return nil
	}
	out := new(NutanixResourceIdentifiers)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NutanixSpec) DeepCopyInto(out *NutanixSpec) {
	*out = *in
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	in.PrismCentralEndpoint.DeepCopyInto(&out.PrismCentralEndpoint)
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
		*out = new(v1.LocalObjectReference)
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
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
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
func (in *StorageClassConfig) DeepCopyInto(out *StorageClassConfig) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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
func (in Users) DeepCopyInto(out *Users) {
	{
		in := &in
		*out = make(Users, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Users.
func (in Users) DeepCopy() Users {
	if in == nil {
		return nil
	}
	out := new(Users)
	in.DeepCopyInto(out)
	return *out
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
