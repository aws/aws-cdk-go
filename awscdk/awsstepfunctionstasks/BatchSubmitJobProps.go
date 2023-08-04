package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for RunBatchJob.
//
// Example:
//   import batch "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//   var batchJobDefinition ecsJobDefinition
//   var batchQueue jobQueue
//
//
//   task := tasks.NewBatchSubmitJob(this, jsii.String("Submit Job"), &BatchSubmitJobProps{
//   	JobDefinitionArn: batchJobDefinition.JobDefinitionArn,
//   	JobName: jsii.String("MyJob"),
//   	JobQueueArn: batchQueue.JobQueueArn,
//   })
//
type BatchSubmitJobProps struct {
	// An optional description for this state.
	// Default: - No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	// Default: - None (Task is executed using the State Machine's execution role).
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	// Default: - None.
	//
	// Deprecated: use `heartbeatTimeout`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// Timeout for the heartbeat.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	HeartbeatTimeout awsstepfunctions.Timeout `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: - The entire task input (JSON path '$').
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Default: - `IntegrationPattern.REQUEST_RESPONSE` for most tasks.
	// `IntegrationPattern.RUN_JOB` for the following exceptions:
	// `BatchSubmitJob`, `EmrAddStep`, `EmrCreateCluster`, `EmrTerminationCluster`, and `EmrContainersStartJobRun`.
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: - The entire JSON node determined by the state input, the task result,
	// and resultPath is passed to the next state (JSON path '$').
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Default: - Replaces the entire input with the result (JSON path '$').
	//
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Default: - None.
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the task.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	TaskTimeout awsstepfunctions.Timeout `field:"optional" json:"taskTimeout" yaml:"taskTimeout"`
	// Timeout for the task.
	// Default: - None.
	//
	// Deprecated: use `taskTimeout`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The arn of the job definition used by this job.
	JobDefinitionArn *string `field:"required" json:"jobDefinitionArn" yaml:"jobDefinitionArn"`
	// The name of the job.
	//
	// The first character must be alphanumeric, and up to 128 letters (uppercase and lowercase),
	// numbers, hyphens, and underscores are allowed.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// The arn of the job queue into which the job is submitted.
	JobQueueArn *string `field:"required" json:"jobQueueArn" yaml:"jobQueueArn"`
	// The array size can be between 2 and 10,000.
	//
	// If you specify array properties for a job, it becomes an array job.
	// For more information, see Array Jobs in the AWS Batch User Guide.
	// Default: - No array size.
	//
	ArraySize *float64 `field:"optional" json:"arraySize" yaml:"arraySize"`
	// The number of times to move a job to the RUNNABLE status.
	//
	// You may specify between 1 and 10 attempts.
	// If the value of attempts is greater than one,
	// the job is retried on failure the same number of attempts as the value.
	// Default: 1.
	//
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
	// A list of container overrides in JSON format that specify the name of a container in the specified job definition and the overrides it should receive.
	// See: https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html#Batch-SubmitJob-request-containerOverrides
	//
	// Default: - No container overrides.
	//
	ContainerOverrides *BatchContainerOverrides `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// A list of dependencies for the job.
	//
	// A job can depend upon a maximum of 20 jobs.
	// See: https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html#Batch-SubmitJob-request-dependsOn
	//
	// Default: - No dependencies.
	//
	DependsOn *[]*BatchJobDependency `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// The payload to be passed as parameters to the batch job.
	// Default: - No parameters are passed.
	//
	Payload awsstepfunctions.TaskInput `field:"optional" json:"payload" yaml:"payload"`
	// The tags applied to the job request.
	// Default: {} - no tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

