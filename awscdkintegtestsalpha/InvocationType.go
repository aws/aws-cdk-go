// CDK Integration Testing Constructs
package awscdkintegtestsalpha


// The type of invocation.
//
// Default is REQUEST_RESPONE.
//
// Example:
//   var app app
//   var stack stack
//   var queue queue
//   var fn iFunction
//
//
//   integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &IntegTestProps{
//   	TestCases: []*stack{
//   		stack,
//   	},
//   })
//
//   integ.Assertions.InvokeFunction(&LambdaInvokeFunctionProps{
//   	FunctionName: fn.FunctionName,
//   	InvocationType: awscdkintegtestsalpha.InvocationType_EVENT,
//   	Payload: jSON.stringify(map[string]*string{
//   		"status": jsii.String("OK"),
//   	}),
//   })
//
//   message := integ.Assertions.AwsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"), map[string]interface{}{
//   	"QueueUrl": queue.queueUrl,
//   	"WaitTimeSeconds": jsii.Number(20),
//   })
//
//   message.AssertAtPath(jsii.String("Messages.0.Body"), awscdkintegtestsalpha.ExpectedResult_ObjectLike(map[string]interface{}{
//   	"requestContext": map[string]*string{
//   		"condition": jsii.String("Success"),
//   	},
//   	"requestPayload": map[string]*string{
//   		"status": jsii.String("OK"),
//   	},
//   	"responseContext": map[string]*f64{
//   		"statusCode": jsii.Number(200),
//   	},
//   	"responsePayload": jsii.String("success"),
//   }))
//
// Experimental.
type InvocationType string

const (
	// Invoke the function asynchronously.
	//
	// Send events that fail multiple times to the function's
	// dead-letter queue (if it's configured).
	// The API response only includes a status code.
	// Experimental.
	InvocationType_EVENT InvocationType = "EVENT"
	// Invoke the function synchronously.
	//
	// Keep the connection open until the function returns a response or times out.
	// The API response includes the function response and additional data.
	// Experimental.
	InvocationType_REQUEST_RESPONE InvocationType = "REQUEST_RESPONE"
	// Validate parameter values and verify that the user or role has permission to invoke the function.
	// Experimental.
	InvocationType_DRY_RUN InvocationType = "DRY_RUN"
)

