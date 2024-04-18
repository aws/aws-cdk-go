package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for calling an external HTTP endpoint with HttpInvoke.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   connection := events.NewConnection(this, jsii.String("Connection"), &ConnectionProps{
//   	Authorization: events.Authorization_Basic(jsii.String("username"), awscdk.SecretValue_UnsafePlainText(jsii.String("password"))),
//   })
//
//   tasks.NewHttpInvoke(this, jsii.String("Invoke HTTP API"), &HttpInvokeProps{
//   	ApiRoot: jsii.String("https://api.example.com"),
//   	ApiEndpoint: sfn.TaskInput_FromText(jsii.String("https://api.example.com/path/to/resource")),
//   	Body: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"foo": jsii.String("bar"),
//   	}),
//   	Connection: Connection,
//   	Headers: sfn.TaskInput_*FromObject(map[string]interface{}{
//   		"Content-Type": jsii.String("application/json"),
//   	}),
//   	Method: sfn.TaskInput_*FromText(jsii.String("POST")),
//   	QueryStringParameters: sfn.TaskInput_*FromObject(map[string]interface{}{
//   		"id": jsii.String("123"),
//   	}),
//   	UrlEncodingFormat: tasks.URLEncodingFormat_BRACKETS,
//   })
//
type HttpInvokeProps struct {
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
	// The API endpoint to call, relative to `apiRoot`.
	//
	// Example:
	//   sfn.TaskInput_FromText(jsii.String("path/to/resource"))
	//
	ApiEndpoint awsstepfunctions.TaskInput `field:"required" json:"apiEndpoint" yaml:"apiEndpoint"`
	// Permissions are granted to call all resources under this path.
	//
	// Example:
	//   "https://api.example.com"
	//
	ApiRoot *string `field:"required" json:"apiRoot" yaml:"apiRoot"`
	// The EventBridge Connection to use for authentication.
	Connection awsevents.IConnection `field:"required" json:"connection" yaml:"connection"`
	// The HTTP method to use.
	//
	// Example:
	//   sfn.TaskInput_FromText(jsii.String("GET"))
	//
	Method awsstepfunctions.TaskInput `field:"required" json:"method" yaml:"method"`
	// The body to send to the HTTP endpoint.
	// Default: - No body is sent with the request.
	//
	Body awsstepfunctions.TaskInput `field:"optional" json:"body" yaml:"body"`
	// The headers to send to the HTTP endpoint.
	//
	// Example:
	//   sfn.TaskInput_FromObject(map[string]interface{}{
	//   	"Content-Type": jsii.String("application/json"),
	//   })
	//
	// Default: - No additional headers are added to the request.
	//
	Headers awsstepfunctions.TaskInput `field:"optional" json:"headers" yaml:"headers"`
	// The query string parameters to send to the HTTP endpoint.
	// Default: - No query string parameters are sent in the request.
	//
	QueryStringParameters awsstepfunctions.TaskInput `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
	// Determines whether to apply URL encoding to the request body, and which array encoding format to use.
	//
	// `URLEncodingFormat.NONE` passes the JSON-serialized `RequestBody` field as the HTTP request body.
	// Otherwise, the HTTP request body is the URL-encoded form data of the `RequestBody` field using the
	// specified array encoding format, and the `Content-Type` header is set to `application/x-www-form-urlencoded`.
	// Default: - URLEncodingFormat.NONE
	//
	UrlEncodingFormat URLEncodingFormat `field:"optional" json:"urlEncodingFormat" yaml:"urlEncodingFormat"`
}

