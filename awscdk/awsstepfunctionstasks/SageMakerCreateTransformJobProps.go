package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for creating an Amazon SageMaker transform job task.
//
// Example:
//   tasks.NewSageMakerCreateTransformJob(this, jsii.String("Batch Inference"), &SageMakerCreateTransformJobProps{
//   	TransformJobName: jsii.String("MyTransformJob"),
//   	ModelName: jsii.String("MyModelName"),
//   	ModelClientOptions: &ModelClientOptions{
//   		InvocationsMaxRetries: jsii.Number(3),
//   		 // default is 0
//   		InvocationsTimeout: awscdk.Duration_Minutes(jsii.Number(5)),
//   	},
//   	TransformInput: &TransformInput{
//   		TransformDataSource: &TransformDataSource{
//   			S3DataSource: &TransformS3DataSource{
//   				S3Uri: jsii.String("s3://inputbucket/train"),
//   				S3DataType: tasks.S3DataType_S3_PREFIX,
//   			},
//   		},
//   	},
//   	TransformOutput: &TransformOutput{
//   		S3OutputPath: jsii.String("s3://outputbucket/TransformJobOutputPath"),
//   	},
//   	TransformResources: &TransformResources{
//   		InstanceCount: jsii.Number(1),
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_M4, ec2.InstanceSize_XLARGE),
//   	},
//   })
//
type SageMakerCreateTransformJobProps struct {
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
	//
	// Depending on the AWS Service, the Service Integration Pattern availability will vary.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-supported-services.html
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
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
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
	// Name of the model that you want to use for the transform job.
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// Dataset to be transformed and the Amazon S3 location where it is stored.
	TransformInput *TransformInput `field:"required" json:"transformInput" yaml:"transformInput"`
	// Transform Job Name.
	TransformJobName *string `field:"required" json:"transformJobName" yaml:"transformJobName"`
	// S3 location where you want Amazon SageMaker to save the results from the transform job.
	TransformOutput *TransformOutput `field:"required" json:"transformOutput" yaml:"transformOutput"`
	// Number of records to include in a mini-batch for an HTTP inference request.
	// Default: - No batch strategy.
	//
	BatchStrategy BatchStrategy `field:"optional" json:"batchStrategy" yaml:"batchStrategy"`
	// Environment variables to set in the Docker container.
	// Default: - No environment variables.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Maximum number of parallel requests that can be sent to each instance in a transform job.
	// Default: - Amazon SageMaker checks the optional execution-parameters to determine the settings for your chosen algorithm.
	// If the execution-parameters endpoint is not enabled, the default value is 1.
	//
	MaxConcurrentTransforms *float64 `field:"optional" json:"maxConcurrentTransforms" yaml:"maxConcurrentTransforms"`
	// Maximum allowed size of the payload, in MB.
	// Default: 6.
	//
	MaxPayload awscdk.Size `field:"optional" json:"maxPayload" yaml:"maxPayload"`
	// Configures the timeout and maximum number of retries for processing a transform job invocation.
	// Default: - 0 retries and 60 seconds of timeout.
	//
	ModelClientOptions *ModelClientOptions `field:"optional" json:"modelClientOptions" yaml:"modelClientOptions"`
	// Role for the Transform Job.
	// Default: - A role is created with `AmazonSageMakerFullAccess` managed policy.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Tags to be applied to the train job.
	// Default: - No tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// ML compute instances for the transform job.
	// Default: - 1 instance of type M4.XLarge
	//
	TransformResources *TransformResources `field:"optional" json:"transformResources" yaml:"transformResources"`
}

