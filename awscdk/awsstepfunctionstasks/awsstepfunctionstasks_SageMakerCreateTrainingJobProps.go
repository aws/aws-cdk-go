package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for creating an Amazon SageMaker training job.
//
// Example:
//   tasks.NewSageMakerCreateTrainingJob(this, jsii.String("TrainSagemaker"), &sageMakerCreateTrainingJobProps{
//   	trainingJobName: sfn.jsonPath.stringAt(jsii.String("$.JobName")),
//   	algorithmSpecification: &algorithmSpecification{
//   		algorithmName: jsii.String("BlazingText"),
//   		trainingInputMode: tasks.inputMode_FILE,
//   	},
//   	inputDataConfig: []channel{
//   		&channel{
//   			channelName: jsii.String("train"),
//   			dataSource: &dataSource{
//   				s3DataSource: &s3DataSource{
//   					s3DataType: tasks.s3DataType_S3_PREFIX,
//   					s3Location: tasks.s3Location.fromJsonExpression(jsii.String("$.S3Bucket")),
//   				},
//   			},
//   		},
//   	},
//   	outputDataConfig: &outputDataConfig{
//   		s3OutputLocation: tasks.*s3Location.fromBucket(s3.bucket.fromBucketName(this, jsii.String("Bucket"), jsii.String("mybucket")), jsii.String("myoutputpath")),
//   	},
//   	resourceConfig: &resourceConfig{
//   		instanceCount: jsii.Number(1),
//   		instanceType: ec2.NewInstanceType(sfn.*jsonPath.stringAt(jsii.String("$.InstanceType"))),
//   		volumeSize: awscdk.Size.gibibytes(jsii.Number(50)),
//   	},
//   	 // optional: default is 1 instance of EC2 `M4.XLarge` with `10GB` volume
//   	stoppingCondition: &stoppingCondition{
//   		maxRuntime: awscdk.Duration.hours(jsii.Number(2)),
//   	},
//   })
//
type SageMakerCreateTrainingJobProps struct {
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	// Deprecated: use `heartbeatTimeout`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// Timeout for the heartbeat.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	HeartbeatTimeout awsstepfunctions.Timeout `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
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
	// Timeout for the task.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	TaskTimeout awsstepfunctions.Timeout `field:"optional" json:"taskTimeout" yaml:"taskTimeout"`
	// Timeout for the task.
	// Deprecated: use `taskTimeout`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Identifies the training algorithm to use.
	AlgorithmSpecification *AlgorithmSpecification `field:"required" json:"algorithmSpecification" yaml:"algorithmSpecification"`
	// Describes the various datasets (e.g. train, validation, test) and the Amazon S3 location where stored.
	InputDataConfig *[]*Channel `field:"required" json:"inputDataConfig" yaml:"inputDataConfig"`
	// Identifies the Amazon S3 location where you want Amazon SageMaker to save the results of model training.
	OutputDataConfig *OutputDataConfig `field:"required" json:"outputDataConfig" yaml:"outputDataConfig"`
	// Training Job Name.
	TrainingJobName *string `field:"required" json:"trainingJobName" yaml:"trainingJobName"`
	// Isolates the training container.
	//
	// No inbound or outbound network calls can be made to or from the training container.
	EnableNetworkIsolation *bool `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Environment variables to set in the Docker container.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Algorithm-specific parameters that influence the quality of the model.
	//
	// Set hyperparameters before you start the learning process.
	// For a list of hyperparameters provided by Amazon SageMaker.
	// See: https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html
	//
	Hyperparameters *map[string]interface{} `field:"optional" json:"hyperparameters" yaml:"hyperparameters"`
	// Specifies the resources, ML compute instances, and ML storage volumes to deploy for model training.
	ResourceConfig *ResourceConfig `field:"optional" json:"resourceConfig" yaml:"resourceConfig"`
	// Role for the Training Job.
	//
	// The role must be granted all necessary permissions for the SageMaker training job to
	// be able to operate.
	//
	// See https://docs.aws.amazon.com/fr_fr/sagemaker/latest/dg/sagemaker-roles.html#sagemaker-roles-createtrainingjob-perms
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Sets a time limit for training.
	StoppingCondition *StoppingCondition `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// Tags to be applied to the train job.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the VPC that you want your training job to connect to.
	VpcConfig *VpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

