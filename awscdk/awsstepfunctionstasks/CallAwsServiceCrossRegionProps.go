package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for calling an AWS service's API action from your state machine across regions.
//
// Example:
//   var myBucket Bucket
//
//   getObject := tasks.NewCallAwsServiceCrossRegion(this, jsii.String("GetObject"), &CallAwsServiceCrossRegionProps{
//   	Region: jsii.String("ap-northeast-1"),
//   	Service: jsii.String("s3"),
//   	Action: jsii.String("getObject"),
//   	Parameters: map[string]interface{}{
//   		"Bucket": myBucket.bucketName,
//   		"Key": sfn.JsonPath_stringAt(jsii.String("$.key")),
//   	},
//   	IamResources: []*string{
//   		myBucket.ArnForObjects(jsii.String("*")),
//   	},
//   })
//
type CallAwsServiceCrossRegionProps struct {
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
	// The API action to call.
	//
	// Use camelCase.
	Action *string `field:"required" json:"action" yaml:"action"`
	// The resources for the IAM statement that will be added to the Lambda function role's policy to allow the state machine to make the API call.
	IamResources *[]*string `field:"required" json:"iamResources" yaml:"iamResources"`
	// The AWS region to call this AWS API for.
	//
	// Example:
	//   "us-east-1"
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// The AWS service to call in AWS SDK for JavaScript v3 format.
	//
	// Example:
	//   "s3"
	//
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/v3/latest/
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// Additional IAM statements that will be added to the state machine role's policy.
	//
	// Use in the case where the call requires more than a single statement to
	// be executed, e.g. `rekognition:detectLabels` requires also S3 permissions
	// to read the object on which it must act.
	// Default: - no additional statements are added.
	//
	AdditionalIamStatements *[]awsiam.PolicyStatement `field:"optional" json:"additionalIamStatements" yaml:"additionalIamStatements"`
	// The AWS API endpoint.
	// Default: Do not override API endpoint.
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The action for the IAM statement that will be added to the Lambda function role's policy to allow the state machine to make the API call.
	//
	// By default the action for this IAM statement will be `service:action`.
	//
	// Use in the case where the IAM action name does not match with the
	// API service/action name, e.g. `s3:ListBuckets` requires `s3:ListAllMyBuckets`.
	// Default: - service:action.
	//
	IamAction *string `field:"optional" json:"iamAction" yaml:"iamAction"`
	// Parameters for the API action call in AWS SDK for JavaScript v3 format.
	// Default: - no parameters.
	//
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Whether to retry on the backend Lambda service exceptions.
	//
	// This handles `Lambda.ServiceException`, `Lambda.AWSLambdaException`,
	// `Lambda.SdkClientException`, and `Lambda.ClientExecutionTimeoutException`
	// with an interval of 2 seconds, a back-off rate
	// of 2 and 6 maximum attempts.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/bp-lambda-serviceexception.html
	//
	// Default: true.
	//
	RetryOnServiceExceptions *bool `field:"optional" json:"retryOnServiceExceptions" yaml:"retryOnServiceExceptions"`
}

