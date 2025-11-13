package interfacesawsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LifecycleHook.
// Experimental.
type ILifecycleHookRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LifecycleHook resource.
	// Experimental.
	LifecycleHookRef() *LifecycleHookReference
}

// The jsii proxy for ILifecycleHookRef
type jsiiProxy_ILifecycleHookRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILifecycleHookRef) LifecycleHookRef() *LifecycleHookReference {
	var returns *LifecycleHookReference
	_jsii_.Get(
		j,
		"lifecycleHookRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILifecycleHookRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILifecycleHookRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

