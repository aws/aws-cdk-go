package triggers


// The invocation type to apply to a trigger.
//
// This determines whether the trigger function should await the result of the to be triggered function or not.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   import triggers "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//
//
//   func := lambda.NewFunction(stack, jsii.String("MyFunction"), &functionProps{
//   	handler: jsii.String("index.handler"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	code: lambda.code.fromInline(jsii.String("foo")),
//   })
//
//   triggers.NewTrigger(stack, jsii.String("MyTrigger"), &triggerProps{
//   	handler: func,
//   	timeout: awscdk.Duration.minutes(jsii.Number(10)),
//   	invocationType: triggers.invocationType_EVENT,
//   })
//
type InvocationType string

const (
	// Invoke the function synchronously.
	//
	// Keep the connection open until the function returns a response or times out.
	// The API response includes the function response and additional data.
	InvocationType_EVENT InvocationType = "EVENT"
	// Invoke the function asynchronously.
	//
	// Send events that fail multiple times to the function's dead-letter queue (if one is configured).
	// The API response only includes a status code.
	InvocationType_REQUEST_RESPONSE InvocationType = "REQUEST_RESPONSE"
	// Validate parameter values and verify that the user or role has permission to invoke the function.
	InvocationType_DRY_RUN InvocationType = "DRY_RUN"
)

