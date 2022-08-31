package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
)

// A Step Functions Task to invoke a Lambda function.
//
// The Lambda function Arn is defined as Resource in the state machine definition.
//
// OUTPUT: the output of this task is the return value of the Lambda Function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var payload interface{}
//
//   invokeFunction := awscdk.Aws_stepfunctions_tasks.NewInvokeFunction(function_, &invokeFunctionProps{
//   	payload: map[string]interface{}{
//   		"payloadKey": payload,
//   	},
//   })
//
// Deprecated: Use `LambdaInvoke`.
type InvokeFunction interface {
	awsstepfunctions.IStepFunctionsTask
	// Called when the task object is used in a workflow.
	// Deprecated: Use `LambdaInvoke`.
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for InvokeFunction
type jsiiProxy_InvokeFunction struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: Use `LambdaInvoke`.
func NewInvokeFunction(lambdaFunction awslambda.IFunction, props *InvokeFunctionProps) InvokeFunction {
	_init_.Initialize()

	j := jsiiProxy_InvokeFunction{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.InvokeFunction",
		[]interface{}{lambdaFunction, props},
		&j,
	)

	return &j
}

// Deprecated: Use `LambdaInvoke`.
func NewInvokeFunction_Override(i InvokeFunction, lambdaFunction awslambda.IFunction, props *InvokeFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.InvokeFunction",
		[]interface{}{lambdaFunction, props},
		i,
	)
}

func (i *jsiiProxy_InvokeFunction) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}

