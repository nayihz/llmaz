/*
Copyright 2024.

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
	lws "sigs.k8s.io/lws/api/leaderworkerset/v1"

	api "inftyai.com/llmaz/api/v1alpha1"
)

// ServiceSpec defines the desired state of Service.
// Service controller will maintain multi-flavor of workloads with
// different accelerators for cost or performance considerations.
type ServiceSpec struct {
	// MultiModelsClaims represents multiple modelClaim, which is useful when different
	// sub-workload has different accelerator requirements, like the state-of-the-art
	// technology called splitwise, the workload template is shared by both.
	// Most of the time, one modelClaim is enough.
	// Note: properties (nodeSelectors, resources, e.g.) of the model flavors
	// will be applied to the workload if not exist.
	// +kubebuilder:validation:MinItems=1
	MultiModelsClaims []api.MultiModelsClaim `json:"multiModelsClaims,omitempty"`
	// WorkloadTemplate defines the underlying workload layout and configuration.
	// Note: the LWS spec might be twisted to support different technologies
	// like splitwise and accelerator fungibility and several LWSs will be created.
	WorkloadTemplate lws.LeaderWorkerSetSpec `json:"workloadTemplate"`
	// ElasticConfig defines the configuration for elastic usage,
	// e.g. the max/min replicas. Default to 0 ~ Inf+.
	// +optional
	ElasticConfig *ElasticConfig `json:"elasticConfig,omitempty"`
}

// ServiceStatus defines the observed state of Service
type ServiceStatus struct {
	// Conditions represents the Inference condition.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Service is the Schema for the services API
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceSpec   `json:"spec,omitempty"`
	Status ServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ServiceList contains a list of Service
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
