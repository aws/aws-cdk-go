package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for calling an HTTP API Endpoint using JSONata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//   var outputs interface{}
//   var stack stack
//   var taskInput taskInput
//   var taskRole taskRole
//   var timeout timeout
//
//   callApiGatewayHttpApiEndpointJsonataProps := &CallApiGatewayHttpApiEndpointJsonataProps{
//   	ApiId: jsii.String("apiId"),
//   	ApiStack: stack,
//   	Method: awscdk.Aws_stepfunctions_tasks.HttpMethod_GET,
//
//   	// the properties below are optional
//   	ApiPath: jsii.String("apiPath"),
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	AuthType: awscdk.*Aws_stepfunctions_tasks.AuthType_NO_AUTH,
//   	Comment: jsii.String("comment"),
//   	Credentials: &Credentials{
//   		Role: taskRole,
//   	},
//   	Headers: taskInput,
//   	Heartbeat: cdk.Duration_Minutes(jsii.Number(30)),
//   	HeartbeatTimeout: timeout,
//   	IntegrationPattern: awscdk.Aws_stepfunctions.IntegrationPattern_REQUEST_RESPONSE,
//   	Outputs: outputs,
//   	QueryLanguage: awscdk.*Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	QueryParameters: taskInput,
//   	RequestBody: taskInput,
//   	StageName: jsii.String("stageName"),
//   	StateName: jsii.String("stateName"),
//   	TaskTimeout: timeout,
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type CallApiGatewayHttpApiEndpointJsonataProps struct {
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
	// Http method for the API.
	Method HttpMethod `field:"required" json:"method" yaml:"method"`
	// Path parameters appended after API endpoint.
	// Default: - No path.
	//
	ApiPath *string `field:"optional" json:"apiPath" yaml:"apiPath"`
	// Authentication methods.
	// Default: AuthType.NO_AUTH
	//
	AuthType AuthType `field:"optional" json:"authType" yaml:"authType"`
	// HTTP request information that does not relate to contents of the request.
	// Default: - No headers.
	//
	Headers awsstepfunctions.TaskInput `field:"optional" json:"headers" yaml:"headers"`
	// Query strings attatched to end of request.
	// Default: - No query parameters.
	//
	QueryParameters awsstepfunctions.TaskInput `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// HTTP Request body.
	// Default: - No request body.
	//
	RequestBody awsstepfunctions.TaskInput `field:"optional" json:"requestBody" yaml:"requestBody"`
	// The Id of the API to call.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The Stack in which the API is defined.
	ApiStack awscdk.Stack `field:"required" json:"apiStack" yaml:"apiStack"`
	// Name of the stage where the API is deployed to in API Gateway.
	// Default: '$default'.
	//
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

