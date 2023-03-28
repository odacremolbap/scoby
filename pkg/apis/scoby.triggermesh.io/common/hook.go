package common

import "knative.dev/pkg/apis"

type Hook struct {
}

type Address struct {
	// Ref points to an addressable object.
	// +optional
	Ref *Reference `json:"ref,omitempty"`

	// URI can be an absolute URL(non-empty scheme and non-empty host) pointing to the target or a relative URI. Relative URIs will be resolved using the base URI retrieved from Ref.
	// +optional
	URI *apis.URL `json:"uri,omitempty"`

	// Timeout for hook calls.
	// +optional
	Timeout *string
}

// Reference contains enough information to refer to another object.
// It's a trimmed down version of corev1.ObjectReference.
type Reference struct {
	// Kind of the referent.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind"`

	// Namespace of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	// This is optional field, it gets defaulted to the object holding it if left out.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Name of the hook address.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name"`

	// API version of the hook address.
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`

	// Initialization configuration for the hook
	// +optional
	Initialization *HookServiceConfiguration

	// Finalization configuration for the hook
	// +optional
	Finalization *HookServiceConfiguration
}

// Initialization configuration for the hook
type HookServiceConfiguration struct {
	// Whether the hook service is supported.
	// +optional
	Enabled *bool

	// API version implemented by the hook service.
	// +optional
	APIVersion *string
}
