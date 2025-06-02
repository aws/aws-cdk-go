package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for invoking a Bedrock Model.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//
//   var outputBucket iBucket
//   var trainingBucket iBucket
//   var validationBucket iBucket
//   var kmsKey iKey
//   var vpc iVpc
//
//
//   model := bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_EXPRESS_V1())
//
//   task := tasks.NewBedrockCreateModelCustomizationJob(this, jsii.String("CreateModelCustomizationJob"), &BedrockCreateModelCustomizationJobProps{
//   	BaseModel: model,
//   	ClientRequestToken: jsii.String("MyToken"),
//   	CustomizationType: tasks.CustomizationType_FINE_TUNING,
//   	CustomModelKmsKey: kmsKey,
//   	CustomModelName: jsii.String("MyCustomModel"),
//   	 // required
//   	CustomModelTags: []customModelTag{
//   		&customModelTag{
//   			Key: jsii.String("key1"),
//   			Value: jsii.String("value1"),
//   		},
//   	},
//   	HyperParameters: map[string]*string{
//   		"batchSize": jsii.String("10"),
//   	},
//   	JobName: jsii.String("MyCustomizationJob"),
//   	 // required
//   	JobTags: []*customModelTag{
//   		&customModelTag{
//   			Key: jsii.String("key2"),
//   			Value: jsii.String("value2"),
//   		},
//   	},
//   	OutputData: &OutputBucketConfiguration{
//   		Bucket: outputBucket,
//   		 // required
//   		Path: jsii.String("output-data/"),
//   	},
//   	TrainingData: &TrainingBucketConfiguration{
//   		Bucket: trainingBucket,
//   		Path: jsii.String("training-data/data.json"),
//   	},
//   	 // required
//   	// If you don't provide validation data, you have to specify `Evaluation percentage` hyperparameter.
//   	ValidationData: []validationBucketConfiguration{
//   		&validationBucketConfiguration{
//   			Bucket: validationBucket,
//   			Path: jsii.String("validation-data/data.json"),
//   		},
//   	},
//   	VpcConfig: map[string][]iSecurityGroup{
//   		"securityGroups": []*iSecurityGroup{
//   			ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   				"vpc": vpc,
//   			}),
//   		},
//   		"subnets": vpc.privateSubnets,
//   	},
//   })
//
type BedrockCreateModelCustomizationJobProps struct {
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
	// The base model.
	BaseModel awsbedrock.IModel `field:"required" json:"baseModel" yaml:"baseModel"`
	// A name for the resulting custom model.
	//
	// The maximum length is 63 characters.
	CustomModelName *string `field:"required" json:"customModelName" yaml:"customModelName"`
	// A name for the fine-tuning job.
	//
	// The maximum length is 63 characters.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// The S3 bucket configuration where the output data is stored.
	// See: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_OutputDataConfig.html
	//
	OutputData *OutputBucketConfiguration `field:"required" json:"outputData" yaml:"outputData"`
	// The S3 bucket configuration where the training data is stored.
	// See: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_TrainingDataConfig.html
	//
	TrainingData *TrainingBucketConfiguration `field:"required" json:"trainingData" yaml:"trainingData"`
	// A unique, case-sensitive identifier to ensure that the API request completes no more than one time.
	//
	// If this token matches a previous request, Amazon Bedrock ignores the request, but does not return an error.
	//
	// The maximum length is 256 characters.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	//
	// Default: - no client request token.
	//
	ClientRequestToken *string `field:"optional" json:"clientRequestToken" yaml:"clientRequestToken"`
	// The customization type.
	// Default: FINE_TUNING.
	//
	CustomizationType CustomizationType `field:"optional" json:"customizationType" yaml:"customizationType"`
	// The custom model is encrypted at rest using this key.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/encryption-custom-job.html
	//
	// Default: - encrypted with the AWS owned key.
	//
	CustomModelKmsKey awskms.IKey `field:"optional" json:"customModelKmsKey" yaml:"customModelKmsKey"`
	// Tags to attach to the resulting custom model.
	//
	// The maximum number of tags is 200.
	// Default: - no tags.
	//
	CustomModelTags *[]*CustomModelTag `field:"optional" json:"customModelTags" yaml:"customModelTags"`
	// Parameters related to tuning the model.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models-hp.html
	//
	// Default: - use default hyperparameters.
	//
	HyperParameters *map[string]*string `field:"optional" json:"hyperParameters" yaml:"hyperParameters"`
	// Tags to attach to the job.
	//
	// The maximum number of tags is 200.
	// Default: - no tags.
	//
	JobTags *[]*CustomModelTag `field:"optional" json:"jobTags" yaml:"jobTags"`
	// The IAM role that Amazon Bedrock can assume to perform tasks on your behalf.
	//
	// For example, during model training, Amazon Bedrock needs your permission to read input data from an S3 bucket,
	// write model artifacts to an S3 bucket.
	// To pass this role to Amazon Bedrock, the caller of this API must have the iam:PassRole permission.
	// Default: - use auto generated role.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The S3 bucket configuration where the validation data is stored.
	//
	// If you don't provide a validation dataset, specify the evaluation percentage by the `Evaluation percentage` hyperparameter.
	//
	// The maximum number is 10.
	// See: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_Validator.html
	//
	// Default: undefined - validate using a subset of the training data.
	//
	ValidationData *[]*ValidationBucketConfiguration `field:"optional" json:"validationData" yaml:"validationData"`
	// The VPC configuration.
	// Default: - no VPC configuration.
	//
	VpcConfig IBedrockCreateModelCustomizationJobVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

