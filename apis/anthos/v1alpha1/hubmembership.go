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

// HubMembershipParameters is the cluster configs.
type HubMembershipParameters struct {
	MembershipID string `json:"membershipId"`
	GKEClusterID string `json:"GkeClusterId,omitempty"`
	Issuer       string `json:"Issuer,omitempty"`
}

// HubMembershipObservation is the cluster output.
type HubMembershipObservation struct {
	// State of the resource.
	State string `json:"state,omitempty"`
}

// A HubMembershipSpec defines the desired state of a Cluster.
type HubMembershipSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HubMembershipParameters `json:"forProvider"`
}

// A HubMembershipStatus represents the observed state of a Cluster.
type HubMembershipStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HubMembershipObservation `json:"atProvider,omitempty"`
}

// HubMembership is a managed resource that represents a Google Kubernetes Engine
// cluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type HubMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HubMembershipSpec   `json:"spec"`
	Status HubMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HubMembershipList contains a list of Cluster items
type HubMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HubMembership `json:"items"`
}
