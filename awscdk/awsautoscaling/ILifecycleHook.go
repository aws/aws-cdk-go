package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsautoscaling"
	"github.com/aws/constructs-go/constructs/v10"
)

// A basic lifecycle hook object.
type ILifecycleHook interface {
	interfacesawsautoscaling.ILifecycleHookRef
	awscdk.IResource
	// The role for the lifecycle hook to execute.
	// Default: - A default role is created if 'notificationTarget' is specified.
	// Otherwise, no role is created.
	//
	Role() awsiam.IRole
}

// The jsii proxy for ILifecycleHook
type jsiiProxy_ILifecycleHook struct {
	internal.Type__interfacesawsautoscalingILifecycleHookRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ILifecycleHook) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ILifecycleHook) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ILifecycleHook) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILifecycleHook) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILifecycleHook) LifecycleHookRef() *interfacesawsautoscaling.LifecycleHookReference {
	var returns *interfacesawsautoscaling.LifecycleHookReference
	_jsii_.Get(
		j,
		"lifecycleHookRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILifecycleHook) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILifecycleHook) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

