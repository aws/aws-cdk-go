package awsautoscalinghooktargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscalinghooktargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Lambda Function as a hook target.
//
// Internally creates a Topic to make the connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ Function
//   var key Key
//
//   functionHook := awscdk.Aws_autoscaling_hooktargets.NewFunctionHook(function_, key)
//
type FunctionHook interface {
	awsautoscaling.ILifecycleHookTarget
	// If the `IRole` does not exist in `options`, will create an `IRole` and an SNS Topic and attach both to the lifecycle hook.
	//
	// If the `IRole` does exist in `options`, will only create an SNS Topic and attach it to the lifecycle hook.
	Bind(scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) *awsautoscaling.LifecycleHookTargetConfig
}

// The jsii proxy struct for FunctionHook
type jsiiProxy_FunctionHook struct {
	internal.Type__awsautoscalingILifecycleHookTarget
}

func NewFunctionHook(fn awslambda.IFunction, encryptionKey awskms.IKey) FunctionHook {
	_init_.Initialize()

	if err := validateNewFunctionHookParameters(fn); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionHook{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling_hooktargets.FunctionHook",
		[]interface{}{fn, encryptionKey},
		&j,
	)

	return &j
}

func NewFunctionHook_Override(f FunctionHook, fn awslambda.IFunction, encryptionKey awskms.IKey) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling_hooktargets.FunctionHook",
		[]interface{}{fn, encryptionKey},
		f,
	)
}

func (f *jsiiProxy_FunctionHook) Bind(scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) *awsautoscaling.LifecycleHookTargetConfig {
	if err := f.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *awsautoscaling.LifecycleHookTargetConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

