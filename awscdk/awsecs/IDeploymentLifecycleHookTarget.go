package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for deployment lifecycle hook targets.
type IDeploymentLifecycleHookTarget interface {
	// Bind this target to a deployment lifecycle hook.
	Bind(scope constructs.IConstruct) *DeploymentLifecycleHookTargetConfig
}

// The jsii proxy for IDeploymentLifecycleHookTarget
type jsiiProxy_IDeploymentLifecycleHookTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IDeploymentLifecycleHookTarget) Bind(scope constructs.IConstruct) *DeploymentLifecycleHookTargetConfig {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *DeploymentLifecycleHookTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

