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
//   runBatchJob := awscdk.Aws_stepfunctions_tasks.NewRunBatchJob(&RunBatchJobProps{
//   	JobDefinitionArn: jsii.String("jobDefinitionArn"),
//   	JobName: jsii.String("jobName"),
//   	JobQueueArn: jsii.String("jobQueueArn"),
//
//   	// the properties below are optional
//   	ArraySize: jsii.Number(123),
//   	Attempts: jsii.Number(123),
//   	ContainerOverrides: &ContainerOverrides{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		GpuCount: jsii.Number(123),
//   		InstanceType: instanceType,
//   		Memory: jsii.Number(123),
//   		Vcpus: jsii.Number(123),
//   	},
//   	DependsOn: []jobDependency{
//   		&jobDependency{
//   			JobId: jsii.String("jobId"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	IntegrationPattern: awscdk.Aws_stepfunctions.ServiceIntegrationPattern_FIRE_AND_FORGET,
//   	Payload: map[string]interface{}{
//   		"payloadKey": payload,
//   	},
//   	Timeout: duration,
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

	if err := validateNewRunBatchJobParameters(props); err != nil {
		panic(err)
	}
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

