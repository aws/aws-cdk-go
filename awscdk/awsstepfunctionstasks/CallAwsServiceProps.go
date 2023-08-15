package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for calling an AWS service's API action from your state machine.
//
// Example:
//   detectLabels := tasks.NewCallAwsService(this, jsii.String("DetectLabels"), &CallAwsServiceProps{
//   	Service: jsii.String("rekognition"),
//   	Action: jsii.String("detectLabels"),
//   	IamResources: []*string{
//   		jsii.String("*"),
//   	},
//   	AdditionalIamStatements: []policyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   			Actions: []*string{
//   				jsii.String("s3:getObject"),
//   			},
//   			Resources: []*string{
//   				jsii.String("arn:aws:s3:::my-bucket/*"),
//   			},
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/supported-services-awssdk.html
//
type CallAwsServiceProps struct {
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
	// The API action to call.
	//
	// Use camelCase.
	Action *string `field:"required" json:"action" yaml:"action"`
	// The resources for the IAM statement that will be added to the state machine role's policy to allow the state machine to make the API call.
	//
	// By default the action for this IAM statement will be `service:action`.
	IamResources *[]*string `field:"required" json:"iamResources" yaml:"iamResources"`
	// The AWS service to call.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/supported-services-awssdk.html
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
	// The action for the IAM statement that will be added to the state machine role's policy to allow the state machine to make the API call.
	//
	// Use in the case where the IAM action name does not match with the
	// API service/action name, e.g. `s3:ListBuckets` requires `s3:ListAllMyBuckets`.
	// Default: - service:action.
	//
	IamAction *string `field:"optional" json:"iamAction" yaml:"iamAction"`
	// Parameters for the API action call.
	//
	// Use PascalCase for the parameter names.
	// Default: - no parameters.
	//
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

