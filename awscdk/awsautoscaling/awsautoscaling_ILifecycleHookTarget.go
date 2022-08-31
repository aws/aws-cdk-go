package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
)

// Interface for autoscaling lifecycle hook targets.
// Experimental.
type ILifecycleHookTarget interface {
	// Called when this object is used as the target of a lifecycle hook.
	// Experimental.
	Bind(scope constructs.Construct, options *BindHookTargetOptions) *LifecycleHookTargetConfig
}

// The jsii proxy for ILifecycleHookTarget
type jsiiProxy_ILifecycleHookTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_ILifecycleHookTarget) Bind(scope constructs.Construct, options *BindHookTargetOptions) *LifecycleHookTargetConfig {
	var returns *LifecycleHookTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

