/*


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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Guestbook1Spec defines the desired state of Guestbook1
type Guestbook1Spec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Guestbook1. Edit Guestbook1_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// Guestbook1Status defines the observed state of Guestbook1
type Guestbook1Status struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Guestbook1 is the Schema for the guestbook1s API
type Guestbook1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Guestbook1Spec   `json:"spec,omitempty"`
	Status Guestbook1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Guestbook1List contains a list of Guestbook1
type Guestbook1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Guestbook1 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Guestbook1{}, &Guestbook1List{})
}
