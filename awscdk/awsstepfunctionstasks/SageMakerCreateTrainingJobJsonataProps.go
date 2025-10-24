package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for creating an Amazon SageMaker training job using JSONata.
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
//   var dockerImage DockerImage
//   var hyperparameters interface{}
//   var instanceType InstanceType
//   var key Key
//   var outputs interface{}
//   var role Role
//   var s3Location S3Location
//   var size Size
//   var subnet Subnet
//   var subnetFilter SubnetFilter
//   var taskRole TaskRole
//   var timeout Timeout
//   var vpc Vpc
//
//   sageMakerCreateTrainingJobJsonataProps := &SageMakerCreateTrainingJobJsonataProps{
//   	AlgorithmSpecification: &AlgorithmSpecification{
//   		AlgorithmName: jsii.String("algorithmName"),
//   		MetricDefinitions: []MetricDefinition{
//   			&MetricDefinition{
//   				Name: jsii.String("name"),
//   				Regex: jsii.String("regex"),
//   			},
//   		},
//   		TrainingImage: dockerImage,
//   		TrainingInputMode: awscdk.Aws_stepfunctions_tasks.InputMode_PIPE,
//   	},
//   	OutputDataConfig: &OutputDataConfig{
//   		S3OutputLocation: s3Location,
//
//   		// the properties below are optional
//   		EncryptionKey: key,
//   	},
//   	TrainingJobName: jsii.String("trainingJobName"),
//
//   	// the properties below are optional
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	Comment: jsii.String("comment"),
//   	Credentials: &Credentials{
//   		Role: taskRole,
//   	},
//   	EnableNetworkIsolation: jsii.Boolean(false),
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	Heartbeat: cdk.Duration_Minutes(jsii.Number(30)),
//   	HeartbeatTimeout: timeout,
//   	Hyperparameters: map[string]interface{}{
//   		"hyperparametersKey": hyperparameters,
//   	},
//   	InputDataConfig: []Channel{
//   		&Channel{
//   			ChannelName: jsii.String("channelName"),
//   			DataSource: &DataSource{
//   				S3DataSource: &S3DataSource{
//   					S3Location: s3Location,
//
//   					// the properties below are optional
//   					AttributeNames: []*string{
//   						jsii.String("attributeNames"),
//   					},
//   					S3DataDistributionType: awscdk.*Aws_stepfunctions_tasks.S3DataDistributionType_FULLY_REPLICATED,
//   					S3DataType: awscdk.*Aws_stepfunctions_tasks.S3DataType_MANIFEST_FILE,
//   				},
//   			},
//
//   			// the properties below are optional
//   			CompressionType: awscdk.*Aws_stepfunctions_tasks.CompressionType_NONE,
//   			ContentType: jsii.String("contentType"),
//   			InputMode: awscdk.*Aws_stepfunctions_tasks.InputMode_PIPE,
//   			RecordWrapperType: awscdk.*Aws_stepfunctions_tasks.RecordWrapperType_NONE,
//   			ShuffleConfig: &ShuffleConfig{
//   				Seed: jsii.Number(123),
//   			},
//   		},
//   	},
//   	IntegrationPattern: awscdk.Aws_stepfunctions.IntegrationPattern_REQUEST_RESPONSE,
//   	Outputs: outputs,
//   	QueryLanguage: awscdk.*Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	ResourceConfig: &ResourceConfig{
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: instanceType,
//   		VolumeSize: size,
//
//   		// the properties below are optional
//   		VolumeEncryptionKey: key,
//   	},
//   	Role: role,
//   	StateName: jsii.String("stateName"),
//   	StoppingCondition: &StoppingCondition{
//   		MaxRuntime: cdk.Duration_*Minutes(jsii.Number(30)),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TaskTimeout: timeout,
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	VpcConfig: &VpcConfig{
//   		Vpc: vpc,
//
//   		// the properties below are optional
//   		Subnets: &SubnetSelection{
//   			AvailabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			OnePerAz: jsii.Boolean(false),
//   			SubnetFilters: []SubnetFilter{
//   				subnetFilter,
//   			},
//   			SubnetGroupName: jsii.String("subnetGroupName"),
//   			Subnets: []ISubnet{
//   				subnet,
//   			},
//   			SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   		},
//   	},
//   }
//
type SageMakerCreateTrainingJobJsonataProps struct {
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
	// Identifies the training algorithm to use.
	AlgorithmSpecification *AlgorithmSpecification `field:"required" json:"algorithmSpecification" yaml:"algorithmSpecification"`
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
	// Describes the various datasets (e.g. train, validation, test) and the Amazon S3 location where stored.
	// Default: - No inputDataConfig.
	//
	InputDataConfig *[]*Channel `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
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

