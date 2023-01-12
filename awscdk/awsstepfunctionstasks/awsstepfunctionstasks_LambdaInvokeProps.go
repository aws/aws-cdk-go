package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for invoking a Lambda function with LambdaInvoke.
//
// Example:
//   var fn function
//
//   tasks.NewLambdaInvoke(this, jsii.String("Invoke with empty object as payload"), &lambdaInvokeProps{
//   	lambdaFunction: fn,
//   	payload: sfn.taskInput.fromObject(map[string]interface{}{
//   	}),
//   })
//
//   // use the output of fn as input
//   // use the output of fn as input
//   tasks.NewLambdaInvoke(this, jsii.String("Invoke with payload field in the state input"), &lambdaInvokeProps{
//   	lambdaFunction: fn,
//   	payload: sfn.*taskInput.fromJsonPathAt(jsii.String("$.Payload")),
//   })
//
type LambdaInvokeProps struct {
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Lambda function to invoke.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Up to 3583 bytes of base64-encoded data about the invoking client to pass to the function.
	ClientContext *string `field:"optional" json:"clientContext" yaml:"clientContext"`
	// Invocation type of the Lambda function.
	InvocationType LambdaInvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The JSON that will be supplied as input to the Lambda function.
	Payload awsstepfunctions.TaskInput `field:"optional" json:"payload" yaml:"payload"`
	// Invoke the Lambda in a way that only returns the payload response without additional metadata.
	//
	// The `payloadResponseOnly` property cannot be used if `integrationPattern`, `invocationType`,
	// `clientContext`, or `qualifier` are specified.
	// It always uses the REQUEST_RESPONSE behavior.
	PayloadResponseOnly *bool `field:"optional" json:"payloadResponseOnly" yaml:"payloadResponseOnly"`
	// Version or alias to invoke a published version of the function.
	//
	// You only need to supply this if you want the version of the Lambda Function to depend
	// on data in the state machine state. If not, you can pass the appropriate Alias or Version object
	// directly as the `lambdaFunction` argument.
	// Deprecated: pass a Version or Alias object as lambdaFunction instead.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
	// Whether to retry on Lambda service exceptions.
	//
	// This handles `Lambda.ServiceException`, `Lambda.AWSLambdaException` and
	// `Lambda.SdkClientException` with an interval of 2 seconds, a back-off rate
	// of 2 and 6 maximum attempts.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/bp-lambda-serviceexception.html
	//
	RetryOnServiceExceptions *bool `field:"optional" json:"retryOnServiceExceptions" yaml:"retryOnServiceExceptions"`
}

