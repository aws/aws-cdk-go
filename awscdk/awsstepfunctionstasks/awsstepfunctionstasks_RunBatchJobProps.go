package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for RunBatchJob.
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
//   runBatchJobProps := &RunBatchJobProps{
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
//   }
//
// Deprecated: use `BatchSubmitJob`.
type RunBatchJobProps struct {
	// The arn of the job definition used by this job.
	// Deprecated: use `BatchSubmitJob`.
	JobDefinitionArn *string `field:"required" json:"jobDefinitionArn" yaml:"jobDefinitionArn"`
	// The name of the job.
	//
	// The first character must be alphanumeric, and up to 128 letters (uppercase and lowercase),
	// numbers, hyphens, and underscores are allowed.
	// Deprecated: use `BatchSubmitJob`.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// The arn of the job queue into which the job is submitted.
	// Deprecated: use `BatchSubmitJob`.
	JobQueueArn *string `field:"required" json:"jobQueueArn" yaml:"jobQueueArn"`
	// The array size can be between 2 and 10,000.
	//
	// If you specify array properties for a job, it becomes an array job.
	// For more information, see Array Jobs in the AWS Batch User Guide.
	// Deprecated: use `BatchSubmitJob`.
	ArraySize *float64 `field:"optional" json:"arraySize" yaml:"arraySize"`
	// The number of times to move a job to the RUNNABLE status.
	//
	// You may specify between 1 and 10 attempts.
	// If the value of attempts is greater than one,
	// the job is retried on failure the same number of attempts as the value.
	// Deprecated: use `BatchSubmitJob`.
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
	// A list of container overrides in JSON format that specify the name of a container in the specified job definition and the overrides it should receive.
	// See: https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html#Batch-SubmitJob-request-containerOverrides
	//
	// Deprecated: use `BatchSubmitJob`.
	ContainerOverrides *ContainerOverrides `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// A list of dependencies for the job.
	//
	// A job can depend upon a maximum of 20 jobs.
	// See: https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html#Batch-SubmitJob-request-dependsOn
	//
	// Deprecated: use `BatchSubmitJob`.
	DependsOn *[]*JobDependency `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// The service integration pattern indicates different ways to call TerminateCluster.
	//
	// The valid value is either FIRE_AND_FORGET or SYNC.
	// Deprecated: use `BatchSubmitJob`.
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// The payload to be passed as parametrs to the batch job.
	// Deprecated: use `BatchSubmitJob`.
	Payload *map[string]interface{} `field:"optional" json:"payload" yaml:"payload"`
	// The timeout configuration for this SubmitJob operation.
	//
	// The minimum value for the timeout is 60 seconds.
	// See: https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html#Batch-SubmitJob-request-timeout
	//
	// Deprecated: use `BatchSubmitJob`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

