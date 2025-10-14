package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LifecycleHook.
// Experimental.
type ILifecycleHookRef interface {
	constructs.IConstruct
	// A reference to a LifecycleHook resource.
	// Experimental.
	LifecycleHookRef() *LifecycleHookReference
}

// The jsii proxy for ILifecycleHookRef
type jsiiProxy_ILifecycleHookRef struct {
	internal.Type__constructsIConstruct
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

