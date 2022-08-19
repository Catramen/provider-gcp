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

// AttachedClusterParameters is the cluster configs.
type AttachedClusterParameters struct {
	// Location: The name of the Google Compute
	Location string `json:"location"`
	// PlatformVersion is the platform version.
	PlatformVersion string    `json:"platformVersion,omitempty"`
	Authority       Authority `json:"authority,omitempty"`
	Fleet           Fleet     `json:"fleet,omitempty"`
}

// Authority is the OIDC authority.
type Authority struct {
	IssuerURL string `json:"issuerUrl,omitempty"`
}

// Fleet is the attached cluster fleet configuration.
type Fleet struct {
	Project string `json:"project,omitempty"`
}

// AttachedClusterObservation is the cluster output.
type AttachedClusterObservation struct {
	// CreationTimestamp: Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Id: The unique identifier for the resource. This
	// identifier is defined by the server.
	ID uint64 `json:"id,omitempty"`

	// SelfLink: Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// State of the resource.
	State string `json:"state,omitempty"`

	// MembershipID is the membership id.
	MembershipID string `json:"membershipId,omitempty"`

	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
}

// A AttachedClusterSpec defines the desired state of a Cluster.
type AttachedClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AttachedClusterParameters `json:"forProvider"`
}

// A AttachedClusterStatus represents the observed state of a Cluster.
type AttachedClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AttachedClusterObservation `json:"atProvider,omitempty"`
}

// AttachedCluster is a managed resource that represents a Google Kubernetes Engine
// cluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.status"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.forProvider.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=AttachedCluster,categories={crossplane,managed,gcp}
type AttachedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AttachedClusterSpec   `json:"spec"`
	Status AttachedClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttachedClusterList contains a list of Cluster items
type AttachedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AttachedCluster `json:"items"`
}
