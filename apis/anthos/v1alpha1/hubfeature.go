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

// HubFeatureParameters is the cluster configs.
type HubFeatureParameters struct {
	// Location: The name of the Google Compute
	Location    string `json:"location"`
	FeatureName string `json:"featureName,omitempty"`
}

// HubFeatureObservation is the cluster output.
type HubFeatureObservation struct {
	// State of the resource.
	State string `json:"state,omitempty"`
}

// A HubFeatureSpec defines the desired state of a Cluster.
type HubFeatureSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HubFeatureParameters `json:"forProvider"`
}

// A HubFeatureStatus represents the observed state of a Cluster.
type HubFeatureStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HubFeatureObservation `json:"atProvider,omitempty"`
}

// HubFeature is a managed resource that represents a Google Kubernetes Engine
// cluster.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.forProvider.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type HubFeature struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HubFeatureSpec   `json:"spec"`
	Status HubFeatureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HubFeatureList contains a list of Cluster items
type HubFeatureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HubFeature `json:"items"`
}
