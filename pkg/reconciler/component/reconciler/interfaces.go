package reconciler

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	apicommon "github.com/triggermesh/scoby/pkg/apis/scoby.triggermesh.io/common"
	"github.com/triggermesh/scoby/pkg/reconciler/resources"
)

type Base interface {
	reconcile.Reconciler
}

type Object interface {
	// Kubernetes object
	client.Object
	AsKubeObject() client.Object

	// Status management
	GetStatusManager() StatusManager

	// Rendering from registration
	ObjectRender
}

// type ObjectStatus interface {
// 	StatusGetObservedGeneration() int64
// 	StatusSetObservedGeneration(generation int64)
// 	StatusGetCondition(conditionType string) *apicommon.Condition
// 	StatusSetCondition(condition *apicommon.Condition)
// 	StatusSetAddressURL(string)
// 	// StatusIsEqual(client.Object) bool
// 	StatusSetValue(value interface{}, path ...string) error
// }

// ObjectRender contains rendering methods, either to be used by the renderer
// to convert the registration information into workload assets, or by the
// reconciler to retrieve those assets.
type ObjectRender interface {

	// AddEnvVar is used by a renderer to add a new a new environment
	// variable to the rendered object informing tracking information
	// about the JSON path of the object element that originates
	// the variable.
	AddEnvVar(path string, value *corev1.EnvVar)

	// GetEnvVarAtPath is used to retrieve environment variables set
	// after an object's path element.
	//
	// The main (and probably only) case is a status rendering option
	// that references a value set after an object's path
	GetEnvVarAtPath(path string) *corev1.EnvVar

	// Once rendered an object can be queried about the container options
	// that they resulting worload must include.
	AsContainerOptions() []resources.ContainerOption
}

type StatusManager interface {
	GetObservedGeneration() int64
	SetObservedGeneration(int64)
	GetCondition(conditionType string) *apicommon.Condition
	SetCondition(condition *apicommon.Condition)
	GetAddressURL() string
	SetAddressURL(string)
	SetValue(value interface{}, path ...string) error
}

type StatusManagerFactory interface {
	// Configures the internal set of conditions for the
	// generated status managers.
	UpdateConditionSet(happyCond string, conditions ...string)

	// Create a new status manager for the object.
	ForObject(object *unstructured.Unstructured) StatusManager
}

type Resolver interface {
	Resolve(ctx context.Context, ref *corev1.ObjectReference) (string, error)
}

type ObjectRenderer interface {
	Render(context.Context, Object) error
}

type ObjectManager interface {
	NewObject() Object
	GetRenderer() ObjectRenderer
}

type FormFactor interface {
	SetupController(name string, c controller.Controller, owner runtime.Object) error
	Reconcile(context.Context, Object) (ctrl.Result, error)
}
