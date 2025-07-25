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
	"log"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	testv2 "github.com/operator-framework/test-operator/api/v2"
)

// ConvertTo converts this OLMTest (v1) to the Hub version (v2).
func (src *OLMTest) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*testv2.OLMTest)
	log.Printf("ConvertTo: Converting OLMTest from Spoke version v1 to Hub version v2;"+
		"source: %s/%s, target: %s/%s", src.Namespace, src.Name, dst.Namespace, dst.Name)

	dst.Spec.NewField1 = src.Spec.Field1
	dst.Spec.NewField2 = src.Spec.Field2
	dst.Spec.NewField3 = src.Spec.Field3
	dst.Spec.NewField4 = src.Spec.Field4
	return nil
}

// ConvertFrom converts the Hub version (v2) to this OLMTest (v1).
func (dst *OLMTest) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*testv2.OLMTest)
	log.Printf("ConvertFrom: Converting OLMTest from Hub version v2 to Spoke version v1;"+
		"source: %s/%s, target: %s/%s", src.Namespace, src.Name, dst.Namespace, dst.Name)

	dst.Spec.Field1 = src.Spec.NewField1
	dst.Spec.Field2 = src.Spec.NewField2
	dst.Spec.Field3 = src.Spec.NewField3
	dst.Spec.Field4 = src.Spec.NewField4
	return nil
}
