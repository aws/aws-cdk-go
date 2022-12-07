package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
)

// Invoke a Glue job as a Task.
//
// OUTPUT: the output of this task is a JobRun structure, for details consult
// https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-runs.html#aws-glue-api-jobs-runs-JobRun
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   runGlueJobTask := awscdk.Aws_stepfunctions_tasks.NewRunGlueJobTask(jsii.String("glueJobName"), &runGlueJobTaskProps{
//   	arguments: map[string]*string{
//   		"argumentsKey": jsii.String("arguments"),
//   	},
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	notifyDelayAfter: duration,
//   	securityConfiguration: jsii.String("securityConfiguration"),
//   	timeout: duration,
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-glue.html
//
// Deprecated: use `GlueStartJobRun`.
type RunGlueJobTask interface {
	awsstepfunctions.IStepFunctionsTask
	// Called when the task object is used in a workflow.
	// Deprecated: use `GlueStartJobRun`.
	Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for RunGlueJobTask
type jsiiProxy_RunGlueJobTask struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: use `GlueStartJobRun`.
func NewRunGlueJobTask(glueJobName *string, props *RunGlueJobTaskProps) RunGlueJobTask {
	_init_.Initialize()

	if err := validateNewRunGlueJobTaskParameters(glueJobName, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RunGlueJobTask{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunGlueJobTask",
		[]interface{}{glueJobName, props},
		&j,
	)

	return &j
}

// Deprecated: use `GlueStartJobRun`.
func NewRunGlueJobTask_Override(r RunGlueJobTask, glueJobName *string, props *RunGlueJobTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunGlueJobTask",
		[]interface{}{glueJobName, props},
		r,
	)
}

func (r *jsiiProxy_RunGlueJobTask) Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	if err := r.validateBindParameters(task); err != nil {
		panic(err)
	}
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

