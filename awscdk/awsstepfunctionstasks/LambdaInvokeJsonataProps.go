package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for invoking a Lambda function with LambdaInvoke using Jsonata.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var callApiFunc function
//   var useVariableFunc function
//
//   step1 := tasks.LambdaInvoke_Jsonata(this, jsii.String("Step 1"), &LambdaInvokeJsonataProps{
//   	LambdaFunction: callApiFunc,
//   	Assign: map[string]interface{}{
//   		"x": jsii.String("{% $states.result.Payload.x %}"),
//   	},
//   })
//   step2 := sfn.Pass_Jsonata(this, jsii.String("Step 2"))
//   step3 := sfn.Pass_Jsonata(this, jsii.String("Step 3"))
//   step4 := sfn.Pass_Jsonata(this, jsii.String("Step 4"))
//   step5 := tasks.LambdaInvoke_Jsonata(this, jsii.String("Step 5"), &LambdaInvokeJsonataProps{
//   	LambdaFunction: useVariableFunc,
//   	Payload: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"x": jsii.String("{% $x %}"),
//   	}),
//   })
//
type LambdaInvokeJsonataProps struct {
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
	// Lambda function to invoke.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Up to 3583 bytes of base64-encoded data about the invoking client to pass to the function.
	// Default: - No context.
	//
	ClientContext *string `field:"optional" json:"clientContext" yaml:"clientContext"`
	// Invocation type of the Lambda function.
	// Default: InvocationType.REQUEST_RESPONSE
	//
	InvocationType LambdaInvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The JSON that will be supplied as input to the Lambda function.
	// Default: - The state input (JSONata: '{% $states.input %}', JSONPath: '$')
	//
	Payload awsstepfunctions.TaskInput `field:"optional" json:"payload" yaml:"payload"`
	// Invoke the Lambda in a way that only returns the payload response without additional metadata.
	//
	// The `payloadResponseOnly` property cannot be used if `integrationPattern`, `invocationType`,
	// `clientContext`, or `qualifier` are specified.
	// It always uses the REQUEST_RESPONSE behavior.
	// Default: false.
	//
	PayloadResponseOnly *bool `field:"optional" json:"payloadResponseOnly" yaml:"payloadResponseOnly"`
	// Version or alias to invoke a published version of the function.
	//
	// You only need to supply this if you want the version of the Lambda Function to depend
	// on data in the state machine state. If not, you can pass the appropriate Alias or Version object
	// directly as the `lambdaFunction` argument.
	// Default: - Version or alias inherent to the `lambdaFunction` object.
	//
	// Deprecated: pass a Version or Alias object as lambdaFunction instead.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
	// Whether to retry on Lambda service exceptions.
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

