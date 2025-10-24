package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for BatchSubmitJob using JSONata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//   var instanceType InstanceType
//   var outputs interface{}
//   var size Size
//   var taskInput TaskInput
//   var taskRole TaskRole
//   var timeout Timeout
//
//   batchSubmitJobJsonataProps := &BatchSubmitJobJsonataProps{
//   	JobDefinitionArn: jsii.String("jobDefinitionArn"),
//   	JobName: jsii.String("jobName"),
//   	JobQueueArn: jsii.String("jobQueueArn"),
//
//   	// the properties below are optional
//   	ArraySize: jsii.Number(123),
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	Attempts: jsii.Number(123),
//   	Comment: jsii.String("comment"),
//   	ContainerOverrides: &BatchContainerOverrides{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		GpuCount: jsii.Number(123),
//   		InstanceType: instanceType,
//   		Memory: size,
//   		Vcpus: jsii.Number(123),
//   	},
//   	Credentials: &Credentials{
//   		Role: taskRole,
//   	},
//   	DependsOn: []BatchJobDependency{
//   		&BatchJobDependency{
//   			JobId: jsii.String("jobId"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Heartbeat: cdk.Duration_Minutes(jsii.Number(30)),
//   	HeartbeatTimeout: timeout,
//   	IntegrationPattern: awscdk.Aws_stepfunctions.IntegrationPattern_REQUEST_RESPONSE,
//   	Outputs: outputs,
//   	Payload: taskInput,
//   	QueryLanguage: awscdk.*Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	StateName: jsii.String("stateName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TaskTimeout: timeout,
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type BatchSubmitJobJsonataProps struct {
	// A comment describing this state.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The name of the query language used by the state.
	//
	// If the state does not contain a `queryLanguage` field,
	// then it will use the query language specified in the top-level `queryLanguage` field.
	// Default: - JSONPath.
	//
	QueryLanguage awsstepfunctions.QueryLanguage `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
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
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	//
	// Depending on the AWS Service, the Service Integration Pattern availability will vary.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-supported-services.html
	//
	// Default: - `IntegrationPattern.REQUEST_RESPONSE` for most tasks.
	// `IntegrationPattern.RUN_JOB` for the following exceptions:
	// `BatchSubmitJob`, `EmrAddStep`, `EmrCreateCluster`, `EmrTerminationCluster`, and `EmrContainersStartJobRun`.
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
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
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
	// Used to specify and transform output from the state.
	//
	// When specified, the value overrides the state output default.
	// The output field accepts any JSON value (object, array, string, number, boolean, null).
	// Any string value, including those inside objects or arrays,
	// will be evaluated as JSONata if surrounded by {% %} characters.
	// Output also accepts a JSONata expression directly.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-input-output-filtering.html
	//
	// Default: - $states.result or $states.errorOutput
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
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

