package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
)

// A Step Functions Task to invoke an Activity worker.
//
// An Activity can be used directly as a Resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var activity activity
//   var duration duration
//
//   invokeActivity := awscdk.Aws_stepfunctions_tasks.NewInvokeActivity(activity, &invokeActivityProps{
//   	heartbeat: duration,
//   })
//
// Deprecated: use `StepFunctionsInvokeActivity`.
type InvokeActivity interface {
	awsstepfunctions.IStepFunctionsTask
	// Called when the task object is used in a workflow.
	// Deprecated: use `StepFunctionsInvokeActivity`.
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for InvokeActivity
type jsiiProxy_InvokeActivity struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: use `StepFunctionsInvokeActivity`.
func NewInvokeActivity(activity awsstepfunctions.IActivity, props *InvokeActivityProps) InvokeActivity {
	_init_.Initialize()

	if err := validateNewInvokeActivityParameters(activity, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InvokeActivity{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.InvokeActivity",
		[]interface{}{activity, props},
		&j,
	)

	return &j
}

// Deprecated: use `StepFunctionsInvokeActivity`.
func NewInvokeActivity_Override(i InvokeActivity, activity awsstepfunctions.IActivity, props *InvokeActivityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.InvokeActivity",
		[]interface{}{activity, props},
		i,
	)
}

func (i *jsiiProxy_InvokeActivity) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	if err := i.validateBindParameters(_task); err != nil {
		panic(err)
	}
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}

