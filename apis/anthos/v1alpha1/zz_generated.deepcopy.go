//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachedCluster) DeepCopyInto(out *AttachedCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedCluster.
func (in *AttachedCluster) DeepCopy() *AttachedCluster {
	if in == nil {
		return nil
	}
	out := new(AttachedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttachedCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachedClusterList) DeepCopyInto(out *AttachedClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AttachedCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedClusterList.
func (in *AttachedClusterList) DeepCopy() *AttachedClusterList {
	if in == nil {
		return nil
	}
	out := new(AttachedClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AttachedClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachedClusterObservation) DeepCopyInto(out *AttachedClusterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedClusterObservation.
func (in *AttachedClusterObservation) DeepCopy() *AttachedClusterObservation {
	if in == nil {
		return nil
	}
	out := new(AttachedClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachedClusterParameters) DeepCopyInto(out *AttachedClusterParameters) {
	*out = *in
	out.Authority = in.Authority
	out.Fleet = in.Fleet
	in.ClusterCredentials.DeepCopyInto(&out.ClusterCredentials)
	if in.AdminUsers != nil {
		in, out := &in.AdminUsers, &out.AdminUsers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedClusterParameters.
func (in *AttachedClusterParameters) DeepCopy() *AttachedClusterParameters {
	if in == nil {
		return nil
	}
	out := new(AttachedClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachedClusterSpec) DeepCopyInto(out *AttachedClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedClusterSpec.
func (in *AttachedClusterSpec) DeepCopy() *AttachedClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AttachedClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachedClusterStatus) DeepCopyInto(out *AttachedClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedClusterStatus.
func (in *AttachedClusterStatus) DeepCopy() *AttachedClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AttachedClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authority) DeepCopyInto(out *Authority) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authority.
func (in *Authority) DeepCopy() *Authority {
	if in == nil {
		return nil
	}
	out := new(Authority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCredentials) DeepCopyInto(out *ClusterCredentials) {
	*out = *in
	in.CommonCredentialSelectors.DeepCopyInto(&out.CommonCredentialSelectors)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCredentials.
func (in *ClusterCredentials) DeepCopy() *ClusterCredentials {
	if in == nil {
		return nil
	}
	out := new(ClusterCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSync) DeepCopyInto(out *ConfigSync) {
	*out = *in
	out.Git = in.Git
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSync.
func (in *ConfigSync) DeepCopy() *ConfigSync {
	if in == nil {
		return nil
	}
	out := new(ConfigSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSyncGitConfig) DeepCopyInto(out *ConfigSyncGitConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSyncGitConfig.
func (in *ConfigSyncGitConfig) DeepCopy() *ConfigSyncGitConfig {
	if in == nil {
		return nil
	}
	out := new(ConfigSyncGitConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fleet) DeepCopyInto(out *Fleet) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleet.
func (in *Fleet) DeepCopy() *Fleet {
	if in == nil {
		return nil
	}
	out := new(Fleet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeature) DeepCopyInto(out *HubFeature) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeature.
func (in *HubFeature) DeepCopy() *HubFeature {
	if in == nil {
		return nil
	}
	out := new(HubFeature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HubFeature) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeatureList) DeepCopyInto(out *HubFeatureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HubFeature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeatureList.
func (in *HubFeatureList) DeepCopy() *HubFeatureList {
	if in == nil {
		return nil
	}
	out := new(HubFeatureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HubFeatureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeatureMembership) DeepCopyInto(out *HubFeatureMembership) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeatureMembership.
func (in *HubFeatureMembership) DeepCopy() *HubFeatureMembership {
	if in == nil {
		return nil
	}
	out := new(HubFeatureMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HubFeatureMembership) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeatureMembershipList) DeepCopyInto(out *HubFeatureMembershipList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HubFeatureMembership, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeatureMembershipList.
func (in *HubFeatureMembershipList) DeepCopy() *HubFeatureMembershipList {
	if in == nil {
		return nil
	}
	out := new(HubFeatureMembershipList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HubFeatureMembershipList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeatureMembershipObservation) DeepCopyInto(out *HubFeatureMembershipObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeatureMembershipObservation.
func (in *HubFeatureMembershipObservation) DeepCopy() *HubFeatureMembershipObservation {
	if in == nil {
		return nil
	}
	out := new(HubFeatureMembershipObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeatureMembershipParameters) DeepCopyInto(out *HubFeatureMembershipParameters) {
	*out = *in
	out.ConfigManagement = in.ConfigManagement
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeatureMembershipParameters.
func (in *HubFeatureMembershipParameters) DeepCopy() *HubFeatureMembershipParameters {
	if in == nil {
		return nil
	}
	out := new(HubFeatureMembershipParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeatureMembershipSpec) DeepCopyInto(out *HubFeatureMembershipSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeatureMembershipSpec.
func (in *HubFeatureMembershipSpec) DeepCopy() *HubFeatureMembershipSpec {
	if in == nil {
		return nil
	}
	out := new(HubFeatureMembershipSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeatureMembershipStatus) DeepCopyInto(out *HubFeatureMembershipStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeatureMembershipStatus.
func (in *HubFeatureMembershipStatus) DeepCopy() *HubFeatureMembershipStatus {
	if in == nil {
		return nil
	}
	out := new(HubFeatureMembershipStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeatureObservation) DeepCopyInto(out *HubFeatureObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeatureObservation.
func (in *HubFeatureObservation) DeepCopy() *HubFeatureObservation {
	if in == nil {
		return nil
	}
	out := new(HubFeatureObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeatureParameters) DeepCopyInto(out *HubFeatureParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeatureParameters.
func (in *HubFeatureParameters) DeepCopy() *HubFeatureParameters {
	if in == nil {
		return nil
	}
	out := new(HubFeatureParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeatureSpec) DeepCopyInto(out *HubFeatureSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeatureSpec.
func (in *HubFeatureSpec) DeepCopy() *HubFeatureSpec {
	if in == nil {
		return nil
	}
	out := new(HubFeatureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubFeatureStatus) DeepCopyInto(out *HubFeatureStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubFeatureStatus.
func (in *HubFeatureStatus) DeepCopy() *HubFeatureStatus {
	if in == nil {
		return nil
	}
	out := new(HubFeatureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubMembership) DeepCopyInto(out *HubMembership) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubMembership.
func (in *HubMembership) DeepCopy() *HubMembership {
	if in == nil {
		return nil
	}
	out := new(HubMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HubMembership) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubMembershipList) DeepCopyInto(out *HubMembershipList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HubMembership, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubMembershipList.
func (in *HubMembershipList) DeepCopy() *HubMembershipList {
	if in == nil {
		return nil
	}
	out := new(HubMembershipList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HubMembershipList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubMembershipObservation) DeepCopyInto(out *HubMembershipObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubMembershipObservation.
func (in *HubMembershipObservation) DeepCopy() *HubMembershipObservation {
	if in == nil {
		return nil
	}
	out := new(HubMembershipObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubMembershipParameters) DeepCopyInto(out *HubMembershipParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubMembershipParameters.
func (in *HubMembershipParameters) DeepCopy() *HubMembershipParameters {
	if in == nil {
		return nil
	}
	out := new(HubMembershipParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubMembershipSpec) DeepCopyInto(out *HubMembershipSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubMembershipSpec.
func (in *HubMembershipSpec) DeepCopy() *HubMembershipSpec {
	if in == nil {
		return nil
	}
	out := new(HubMembershipSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubMembershipStatus) DeepCopyInto(out *HubMembershipStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubMembershipStatus.
func (in *HubMembershipStatus) DeepCopy() *HubMembershipStatus {
	if in == nil {
		return nil
	}
	out := new(HubMembershipStatus)
	in.DeepCopyInto(out)
	return out
}