package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for creating an Amazon SageMaker transform job task.
//
// Example:
//   tasks.NewSageMakerCreateTransformJob(this, jsii.String("Batch Inference"), &sageMakerCreateTransformJobProps{
//   	transformJobName: jsii.String("MyTransformJob"),
//   	modelName: jsii.String("MyModelName"),
//   	modelClientOptions: &modelClientOptions{
//   		invocationsMaxRetries: jsii.Number(3),
//   		 // default is 0
//   		invocationsTimeout: awscdk.Duration.minutes(jsii.Number(5)),
//   	},
//   	transformInput: &transformInput{
//   		transformDataSource: &transformDataSource{
//   			s3DataSource: &transformS3DataSource{
//   				s3Uri: jsii.String("s3://inputbucket/train"),
//   				s3DataType: tasks.s3DataType_S3_PREFIX,
//   			},
//   		},
//   	},
//   	transformOutput: &transformOutput{
//   		s3OutputPath: jsii.String("s3://outputbucket/TransformJobOutputPath"),
//   	},
//   	transformResources: &transformResources{
//   		instanceCount: jsii.Number(1),
//   		instanceType: ec2.instanceType.of(ec2.instanceClass_M4, ec2.instanceSize_XLARGE),
//   	},
//   })
//
type SageMakerCreateTransformJobProps struct {
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
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
	BatchStrategy BatchStrategy `field:"optional" json:"batchStrategy" yaml:"batchStrategy"`
	// Environment variables to set in the Docker container.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Maximum number of parallel requests that can be sent to each instance in a transform job.
	MaxConcurrentTransforms *float64 `field:"optional" json:"maxConcurrentTransforms" yaml:"maxConcurrentTransforms"`
	// Maximum allowed size of the payload, in MB.
	MaxPayload awscdk.Size `field:"optional" json:"maxPayload" yaml:"maxPayload"`
	// Configures the timeout and maximum number of retries for processing a transform job invocation.
	ModelClientOptions *ModelClientOptions `field:"optional" json:"modelClientOptions" yaml:"modelClientOptions"`
	// Role for the Transform Job.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Tags to be applied to the train job.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// ML compute instances for the transform job.
	TransformResources *TransformResources `field:"optional" json:"transformResources" yaml:"transformResources"`
}

