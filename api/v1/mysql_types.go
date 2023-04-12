/*
Copyright 2023.

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

// MysqlSpec defines the desired state of Mysql
type MysqlSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Name    string   `json:"name,omitempty"`
	Image   string   `json:"image,omitempty"`
	Command []string `json:"command,omitempty" protobuf:"bytes,3,rep,name=command"`

	// +optional
	//+kubebuilder:default:=1
	//+kubebuilder:validation:Minimum:=1
	Replicas *int32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
}

type MySqlCode string

const (
	SuccessCode MySqlCode = "success"
	FailedCode  MySqlCode = "failed"
)

// MysqlStatus defines the observed state of Mysql
type MysqlStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Code          MySqlCode `json:"code,omitempty"`
	Replicas      int32     `json:"replicas"`
	ReadyReplicas int32     `json:"readyReplicas"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas,selectorpath=.status.labelSelector
//+kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.code",description="The phase of game."
//+kubebuilder:printcolumn:name="DESIRED",type="integer",JSONPath=".spec.replicas",description="The desired number of pods."
//+kubebuilder:printcolumn:name="CURRENT",type="integer",JSONPath=".status.replicas",description="The number of currently all pods."
//+kubebuilder:printcolumn:name="READY",type="integer",JSONPath=".status.readyReplicas",description="The number of pods ready."

// Mysql is the Schema for the mysqls API
type Mysql struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MysqlSpec   `json:"spec,omitempty"`
	Status MysqlStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MysqlList contains a list of Mysql
type MysqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mysql `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Mysql{}, &MysqlList{})
}
