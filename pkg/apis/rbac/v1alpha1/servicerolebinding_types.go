package v1alpha1

import (
	"bufio"
	"bytes"
	"log"

	"github.com/gogo/protobuf/jsonpb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	istiov1alpha1 "istio.io/api/rbac/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServiceRoleBindingSpec defines the desired state of ServiceRoleBinding
type ServiceRoleBindingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	istiov1alpha1.ServiceRoleBinding
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRoleBinding is the Schema for the servicerolebindings API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ServiceRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceRoleBindingSpec   `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceRoleBindingList contains a list of ServiceRoleBinding
type ServiceRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceRoleBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceRoleBinding{}, &ServiceRoleBindingList{})
}

func (p *ServiceRoleBindingSpec) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	marshaler := jsonpb.Marshaler{}
	err := marshaler.Marshal(writer, &p.ServiceRoleBinding)
	if err != nil {
		log.Printf("Could not marshal ServiceRoleBindingSpec. Error: %v", err)
		return nil, err
	}

	writer.Flush()
	return buffer.Bytes(), nil
}

func (p *ServiceRoleBindingSpec) UnmarshalJSON(b []byte) error {
	reader := bytes.NewReader(b)
	unmarshaler := jsonpb.Unmarshaler{}
	err := unmarshaler.Unmarshal(reader, &p.ServiceRoleBinding)
	if err != nil {
		log.Printf("Could not unmarshal ServiceRoleBindingSpec. Error: %v", err)
		return err
	}
	return nil
}

// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
// Based of https://github.com/istio/istio/blob/release-0.8/pilot/pkg/config/kube/crd/types.go#L450
func (in *ServiceRoleBindingSpec) DeepCopyInto(out *ServiceRoleBindingSpec) {
	*out = *in
}