package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
)

// A Step Functions Task to run AWS Batch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var instanceType instanceType
//   var payload interface{}
//
//   runBatchJob := awscdk.Aws_stepfunctions_tasks.NewRunBatchJob(&runBatchJobProps{
//   	jobDefinitionArn: jsii.String("jobDefinitionArn"),
//   	jobName: jsii.String("jobName"),
//   	jobQueueArn: jsii.String("jobQueueArn"),
//
//   	// the properties below are optional
//   	arraySize: jsii.Number(123),
//   	attempts: jsii.Number(123),
//   	containerOverrides: &containerOverrides{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		gpuCount: jsii.Number(123),
//   		instanceType: instanceType,
//   		memory: jsii.Number(123),
//   		vcpus: jsii.Number(123),
//   	},
//   	dependsOn: []jobDependency{
//   		&jobDependency{
//   			jobId: jsii.String("jobId"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	payload: map[string]interface{}{
//   		"payloadKey": payload,
//   	},
//   	timeout: duration,
//   })
//
// Deprecated: use `BatchSubmitJob`.
type RunBatchJob interface {
	awsstepfunctions.IStepFunctionsTask
	// Called when the task object is used in a workflow.
	// Deprecated: use `BatchSubmitJob`.
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for RunBatchJob
type jsiiProxy_RunBatchJob struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: use `BatchSubmitJob`.
func NewRunBatchJob(props *RunBatchJobProps) RunBatchJob {
	_init_.Initialize()

	j := jsiiProxy_RunBatchJob{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunBatchJob",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: use `BatchSubmitJob`.
func NewRunBatchJob_Override(r RunBatchJob, props *RunBatchJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunBatchJob",
		[]interface{}{props},
		r,
	)
}

func (r *jsiiProxy_RunBatchJob) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}

