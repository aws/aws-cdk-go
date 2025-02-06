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
//   	ApiEndpoint: sfn.TaskInput_FromText(jsii.String("path/to/resource")),
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

