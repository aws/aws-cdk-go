package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for creating an Amazon SageMaker training job.
//
// Example:
//   tasks.NewSageMakerCreateTrainingJob(this, jsii.String("TrainSagemaker"), &SageMakerCreateTrainingJobProps{
//   	TrainingJobName: sfn.JsonPath_StringAt(jsii.String("$.JobName")),
//   	AlgorithmSpecification: &AlgorithmSpecification{
//   		AlgorithmName: jsii.String("BlazingText"),
//   		TrainingInputMode: tasks.InputMode_FILE,
//   	},
//   	InputDataConfig: []channel{
//   		&channel{
//   			ChannelName: jsii.String("train"),
//   			DataSource: &DataSource{
//   				S3DataSource: &S3DataSource{
//   					S3DataType: tasks.S3DataType_S3_PREFIX,
//   					S3Location: tasks.S3Location_FromJsonExpression(jsii.String("$.S3Bucket")),
//   				},
//   			},
//   		},
//   	},
//   	OutputDataConfig: &OutputDataConfig{
//   		S3OutputLocation: tasks.S3Location_FromBucket(s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("mybucket")), jsii.String("myoutputpath")),
//   	},
//   	ResourceConfig: &ResourceConfig{
//   		InstanceCount: jsii.Number(1),
//   		InstanceType: ec2.NewInstanceType(sfn.JsonPath_*StringAt(jsii.String("$.InstanceType"))),
//   		VolumeSize: awscdk.Size_Gibibytes(jsii.Number(50)),
//   	},
//   	 // optional: default is 1 instance of EC2 `M4.XLarge` with `10GB` volume
//   	StoppingCondition: &StoppingCondition{
//   		MaxRuntime: awscdk.Duration_Hours(jsii.Number(2)),
//   	},
//   })
//
type SageMakerCreateTrainingJobProps struct {
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
	// Default: false.
	//
	EnableNetworkIsolation *bool `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Environment variables to set in the Docker container.
	// Default: - No environment variables.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Algorithm-specific parameters that influence the quality of the model.
	//
	// Set hyperparameters before you start the learning process.
	// For a list of hyperparameters provided by Amazon SageMaker.
	// See: https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html
	//
	// Default: - No hyperparameters.
	//
	Hyperparameters *map[string]interface{} `field:"optional" json:"hyperparameters" yaml:"hyperparameters"`
	// Specifies the resources, ML compute instances, and ML storage volumes to deploy for model training.
	// Default: - 1 instance of EC2 `M4.XLarge` with `10GB` volume
	//
	ResourceConfig *ResourceConfig `field:"optional" json:"resourceConfig" yaml:"resourceConfig"`
	// Role for the Training Job.
	//
	// The role must be granted all necessary permissions for the SageMaker training job to
	// be able to operate.
	//
	// See https://docs.aws.amazon.com/fr_fr/sagemaker/latest/dg/sagemaker-roles.html#sagemaker-roles-createtrainingjob-perms
	// Default: - a role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Sets a time limit for training.
	// Default: - max runtime of 1 hour.
	//
	StoppingCondition *StoppingCondition `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// Tags to be applied to the train job.
	// Default: - No tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the VPC that you want your training job to connect to.
	// Default: - No VPC.
	//
	VpcConfig *VpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

