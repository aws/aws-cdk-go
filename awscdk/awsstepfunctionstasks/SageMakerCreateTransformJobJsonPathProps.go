package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for creating an Amazon SageMaker transform job task using JSONPath.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//   var instanceType instanceType
//   var key key
//   var resultSelector interface{}
//   var role role
//   var size size
//   var taskRole taskRole
//   var timeout timeout
//
//   sageMakerCreateTransformJobJsonPathProps := &SageMakerCreateTransformJobJsonPathProps{
//   	ModelName: jsii.String("modelName"),
//   	TransformInput: &TransformInput{
//   		TransformDataSource: &TransformDataSource{
//   			S3DataSource: &TransformS3DataSource{
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				S3DataType: awscdk.Aws_stepfunctions_tasks.S3DataType_MANIFEST_FILE,
//   			},
//   		},
//
//   		// the properties below are optional
//   		CompressionType: awscdk.*Aws_stepfunctions_tasks.CompressionType_NONE,
//   		ContentType: jsii.String("contentType"),
//   		SplitType: awscdk.*Aws_stepfunctions_tasks.SplitType_NONE,
//   	},
//   	TransformJobName: jsii.String("transformJobName"),
//   	TransformOutput: &TransformOutput{
//   		S3OutputPath: jsii.String("s3OutputPath"),
//
//   		// the properties below are optional
//   		Accept: jsii.String("accept"),
//   		AssembleWith: awscdk.*Aws_stepfunctions_tasks.AssembleWith_NONE,
//   		EncryptionKey: key,
//   	},
//
//   	// the properties below are optional
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	BatchStrategy: awscdk.*Aws_stepfunctions_tasks.BatchStrategy_MULTI_RECORD,
//   	Comment: jsii.String("comment"),
//   	Credentials: &Credentials{
//   		Role: taskRole,
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	Heartbeat: cdk.Duration_Minutes(jsii.Number(30)),
//   	HeartbeatTimeout: timeout,
//   	InputPath: jsii.String("inputPath"),
//   	IntegrationPattern: awscdk.Aws_stepfunctions.IntegrationPattern_REQUEST_RESPONSE,
//   	MaxConcurrentTransforms: jsii.Number(123),
//   	MaxPayload: size,
//   	ModelClientOptions: &ModelClientOptions{
//   		InvocationsMaxRetries: jsii.Number(123),
//   		InvocationsTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	},
//   	OutputPath: jsii.String("outputPath"),
//   	QueryLanguage: awscdk.*Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	ResultPath: jsii.String("resultPath"),
//   	ResultSelector: map[string]interface{}{
//   		"resultSelectorKey": resultSelector,
//   	},
//   	Role: role,
//   	StateName: jsii.String("stateName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TaskTimeout: timeout,
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	TransformResources: &TransformResources{
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: instanceType,
//
//   		// the properties below are optional
//   		VolumeEncryptionKey: key,
//   	},
//   }
//
type SageMakerCreateTransformJobJsonPathProps struct {
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
	// See: https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: $.
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: $.
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Default: $.
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

