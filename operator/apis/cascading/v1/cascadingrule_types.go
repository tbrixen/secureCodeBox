/*
Copyright 2020 iteratec GmbH.

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
	executionv1 "github.com/secureCodeBox/secureCodeBox-v2-alpha/operator/apis/execution/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CascadingRuleSpec defines the desired state of CascadingRule
type CascadingRuleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CascadingRule. Edit CascadingRule_types.go to remove/update
	Matches  []MatchesRule        `json:"matches"`
	ScanSpec executionv1.ScanSpec `json:"scanSpec"`
}

// MatchesRule is a generic map which is used to model the structure of a finding for which the CascadingRule should take effect
type MatchesRule struct {
	Name        string                        `json:"name,omitempty"`
	Category    string                        `json:"category,omitempty"`
	Description string                        `json:"description,omitempty"`
	Location    string                        `json:"location,omitempty"`
	Severity    string                        `json:"severity,omitempty"`
	OsiLayer    string                        `json:"osi_layer,omitempty"`
	Attributes  map[string]intstr.IntOrString `json:"attributes,omitempty"`
}

// CascadingRuleStatus defines the observed state of CascadingRule
type CascadingRuleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// CascadingRule is the Schema for the cascadingrules API
type CascadingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CascadingRuleSpec   `json:"spec,omitempty"`
	Status CascadingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CascadingRuleList contains a list of CascadingRule
type CascadingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CascadingRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CascadingRule{}, &CascadingRuleList{})
}
