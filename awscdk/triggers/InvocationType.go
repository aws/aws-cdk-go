package triggers


// The invocation type to apply to a trigger.
//
// This determines whether the trigger function should await the result of the to be triggered function or not.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   func := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Code: lambda.Code_FromInline(jsii.String("foo")),
//   })
//
//   triggers.NewTrigger(this, jsii.String("MyTrigger"), &TriggerProps{
//   	Handler: func,
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(10)),
//   	InvocationType: triggers.InvocationType_EVENT,
//   })
//
type InvocationType string

const (
	// Invoke the function asynchronously.
	//
	// Send events that fail multiple times to the function's dead-letter queue (if one is configured).
	// The API response only includes a status code.
	InvocationType_EVENT InvocationType = "EVENT"
	// Invoke the function synchronously.
	//
	// Keep the connection open until the function returns a response or times out.
	// The API response includes the function response and additional data.
	InvocationType_REQUEST_RESPONSE InvocationType = "REQUEST_RESPONSE"
	// Validate parameter values and verify that the user or role has permission to invoke the function.
	InvocationType_DRY_RUN InvocationType = "DRY_RUN"
)

