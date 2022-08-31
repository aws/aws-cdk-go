package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
)

// Invoke a Lambda function as a Task.
//
// OUTPUT: the output of this task is either the return value of Lambda's
// Invoke call, or whatever the Lambda Function posted back using
// `SendTaskSuccess/SendTaskFailure` in `waitForTaskToken` mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var taskInput taskInput
//
//   runLambdaTask := awscdk.Aws_stepfunctions_tasks.NewRunLambdaTask(function_, &runLambdaTaskProps{
//   	clientContext: jsii.String("clientContext"),
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	invocationType: awscdk.*Aws_stepfunctions_tasks.invocationType_REQUEST_RESPONSE,
//   	payload: taskInput,
//   	qualifier: jsii.String("qualifier"),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-lambda.html
//
// Deprecated: Use `LambdaInvoke`.
type RunLambdaTask interface {
	awsstepfunctions.IStepFunctionsTask
	// Called when the task object is used in a workflow.
	// Deprecated: Use `LambdaInvoke`.
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for RunLambdaTask
type jsiiProxy_RunLambdaTask struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: Use `LambdaInvoke`.
func NewRunLambdaTask(lambdaFunction awslambda.IFunction, props *RunLambdaTaskProps) RunLambdaTask {
	_init_.Initialize()

	if err := validateNewRunLambdaTaskParameters(lambdaFunction, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RunLambdaTask{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunLambdaTask",
		[]interface{}{lambdaFunction, props},
		&j,
	)

	return &j
}

// Deprecated: Use `LambdaInvoke`.
func NewRunLambdaTask_Override(r RunLambdaTask, lambdaFunction awslambda.IFunction, props *RunLambdaTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunLambdaTask",
		[]interface{}{lambdaFunction, props},
		r,
	)
}

func (r *jsiiProxy_RunLambdaTask) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	if err := r.validateBindParameters(_task); err != nil {
		panic(err)
	}
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}

