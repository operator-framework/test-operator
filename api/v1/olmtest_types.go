/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OLMTestSpec defines the desired state of OLMTest.
type OLMTestSpec struct {
	// Field1 is an example field that is used to validate interop with webhooks.
	Field1 string `json:"field1,omitempty"`
	// Field2 is an example field that is used to validate interop with webhooks.
	Field2 string `json:"field2,omitempty"`
	// Field3 is an example field that is used to validate interop with webhooks.
	Field3 string `json:"field3,omitempty"`
	// Field4 is an example field that is used to validate interop with webhooks.
	Field4 string `json:"field4,omitempty"`
}

// OLMTestStatus defines the observed state of OLMTest.
type OLMTestStatus struct {
	// NO Status
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// OLMTest is the Schema for the olmtests API.
type OLMTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OLMTestSpec   `json:"spec,omitempty"`
	Status OLMTestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OLMTestList contains a list of OLMTest.
type OLMTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OLMTest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OLMTest{}, &OLMTestList{})
}
