/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// HubFeatureMembershipParameters is the cluster configs.
type HubFeatureMembershipParameters struct {
	// Location: The name of the Google Compute
	Location         string     `json:"location"`
	Feature          string     `json:"feature"`
	Membership       string     `json:"membership,omitempty"`
	ConfigManagement ConfigSync `json:"configManagement,omitempty"`
}

// ConfigSync is the config sync configs.
type ConfigSync struct {
	Version string              `json:"version"`
	Git     ConfigSyncGitConfig `json:"git"`
}

// ConfigSyncGitConfig is the git config of config sync
type ConfigSyncGitConfig struct {
	Repo       string `json:"repo"`
	Branch     string `json:"branch"`
	SecretType string `json:"secretType"`
	PolicyDir  string `json:"policyDir"`
}

// HubFeatureMembershipObservation is the cluster output.
type HubFeatureMembershipObservation struct {
	// State of the resource.
	State string `json:"state,omitempty"`
}

// A HubFeatureMembershipSpec defines the desired state of a Cluster.
type HubFeatureMembershipSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HubFeatureMembershipParameters `json:"forProvider"`
}

// A HubFeatureMembershipStatus represents the observed state of a Cluster.
type HubFeatureMembershipStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HubFeatureMembershipObservation `json:"atProvider,omitempty"`
}

// HubFeatureMembership is a managed resource that represents a Google Kubernetes Engine
// cluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.status"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.forProvider.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type HubFeatureMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HubFeatureMembershipSpec   `json:"spec"`
	Status HubFeatureMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HubFeatureMembershipList contains a list of Cluster items
type HubFeatureMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HubFeatureMembership `json:"items"`
}
