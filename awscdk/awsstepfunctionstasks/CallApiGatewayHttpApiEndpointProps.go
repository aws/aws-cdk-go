package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for calling an HTTP API Endpoint.
//
// Example:
//   import apigatewayv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   httpApi := apigatewayv2.NewHttpApi(this, jsii.String("MyHttpApi"))
//
//   invokeTask := tasks.NewCallApiGatewayHttpApiEndpoint(this, jsii.String("Call HTTP API"), &CallApiGatewayHttpApiEndpointProps{
//   	ApiId: httpApi.ApiId,
//   	ApiStack: awscdk.*stack_Of(httpApi),
//   	Method: tasks.HttpMethod_GET,
//   })
//
type CallApiGatewayHttpApiEndpointProps struct {
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

