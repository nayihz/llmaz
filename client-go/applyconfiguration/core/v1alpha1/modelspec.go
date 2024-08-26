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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/inftyai/llmaz/api/core/v1alpha1"
)

// ModelSpecApplyConfiguration represents an declarative configuration of the ModelSpec type for use
// with apply.
type ModelSpecApplyConfiguration struct {
	FamilyName       *v1alpha1.ModelName            `json:"familyName,omitempty"`
	Source           *ModelSourceApplyConfiguration `json:"source,omitempty"`
	InferenceFlavors []FlavorApplyConfiguration     `json:"inferenceFlavors,omitempty"`
}

// ModelSpecApplyConfiguration constructs an declarative configuration of the ModelSpec type for use with
// apply.
func ModelSpec() *ModelSpecApplyConfiguration {
	return &ModelSpecApplyConfiguration{}
}

// WithFamilyName sets the FamilyName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FamilyName field is set to the value of the last call.
func (b *ModelSpecApplyConfiguration) WithFamilyName(value v1alpha1.ModelName) *ModelSpecApplyConfiguration {
	b.FamilyName = &value
	return b
}

// WithSource sets the Source field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Source field is set to the value of the last call.
func (b *ModelSpecApplyConfiguration) WithSource(value *ModelSourceApplyConfiguration) *ModelSpecApplyConfiguration {
	b.Source = value
	return b
}

// WithInferenceFlavors adds the given value to the InferenceFlavors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InferenceFlavors field.
func (b *ModelSpecApplyConfiguration) WithInferenceFlavors(values ...*FlavorApplyConfiguration) *ModelSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithInferenceFlavors")
		}
		b.InferenceFlavors = append(b.InferenceFlavors, *values[i])
	}
	return b
}
