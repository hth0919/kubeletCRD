// NOTE: Boilerplate only.  Ignore this file.

// Package v1 contains API Schema definitions for the keti v1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=keti.migration
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "keti.migration", Version: "v1"}

	// SchemeBuilder is the scheme builder with scheme init functions to run for this API package
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// AddToScheme is a common registration function for mapping packaged scoped group & version keys to a scheme
	AddToScheme = SchemeBuilder.AddToScheme
)

// addKnownTypes adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Pod{},
		&PodList{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}