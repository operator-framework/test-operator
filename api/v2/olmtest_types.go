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

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OLMTestSpec defines the desired state of OLMTest.
type OLMTestSpec struct {
	// NewField1 is an example field that is used to validate interop with webhooks.
	NewField1 string `json:"newField1,omitempty"`
	// NewField2 is an example field that is used to validate interop with webhooks.
	NewField2 string `json:"newField2,omitempty"`
	// NewField3 is an example field that is used to validate interop with webhooks.
	NewField3 string `json:"newField3,omitempty"`
	// NewField4 is an example field that is used to validate interop with webhooks.
	NewField4 string `json:"newField4,omitempty"`
}

// OLMTestStatus defines the observed state of OLMTest.
type OLMTestStatus struct {
	// None
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:conversion:hub
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

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
