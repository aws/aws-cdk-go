package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for invoking a Bedrock Model.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   model := bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_EXPRESS_V1())
//
//   task := tasks.NewBedrockInvokeModel(this, jsii.String("Prompt Model"), &BedrockInvokeModelProps{
//   	Model: Model,
//   	Body: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"inputText": jsii.String("Generate a list of five first names."),
//   		"textGenerationConfig": map[string]*f64{
//   			"maxTokenCount": jsii.Number(100),
//   			"temperature": jsii.Number(1),
//   		},
//   	}),
//   	ResultSelector: map[string]interface{}{
//   		"names": sfn.JsonPath_stringAt(jsii.String("$.Body.results[0].outputText")),
//   	},
//   })
//
type BedrockInvokeModelProps struct {
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
	// The Bedrock model that the task will invoke.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/api-methods-run.html
	//
	Model awsbedrock.IModel `field:"required" json:"model" yaml:"model"`
	// The desired MIME type of the inference body in the response.
	// See: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_InvokeModel.html
	//
	// Default: 'application/json'.
	//
	Accept *string `field:"optional" json:"accept" yaml:"accept"`
	// The input data for the Bedrock model invocation.
	//
	// The inference parameters contained in the body depend on the Bedrock model being used.
	//
	// The body must be in the format specified in the `contentType` field.
	// For example, if the content type is `application/json`, the body must be
	// JSON formatted.
	//
	// The body must be up to 256 KB in size. For input data that exceeds 256 KB,
	// use `input` instead to retrieve the input data from S3.
	//
	// You must specify either the `body` or the `input` field, but not both.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html
	//
	// Default: - Input data is retrieved from the location specified in the `input` field.
	//
	Body awsstepfunctions.TaskInput `field:"optional" json:"body" yaml:"body"`
	// The MIME type of the input data in the request.
	// See: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_InvokeModel.html
	//
	// Default: 'application/json'.
	//
	// Deprecated: This property does not require configuration because the only acceptable value is 'application/json'.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The guardrail is applied to the invocation.
	// Default: - No guardrail is applied to the invocation.
	//
	Guardrail Guardrail `field:"optional" json:"guardrail" yaml:"guardrail"`
	// The source location to retrieve the input data from.
	// Default: - Input data is retrieved from the `body` field.
	//
	Input *BedrockInvokeModelInputProps `field:"optional" json:"input" yaml:"input"`
	// The destination location where the API response is written.
	//
	// If you specify this field, the API response body is replaced with a reference to the
	// output location.
	// Default: - The API response body is returned in the result.
	//
	Output *BedrockInvokeModelOutputProps `field:"optional" json:"output" yaml:"output"`
	// Specifies whether to enable or disable the Bedrock trace.
	// Default: - Trace is not enabled for the invocation.
	//
	TraceEnabled *bool `field:"optional" json:"traceEnabled" yaml:"traceEnabled"`
}

