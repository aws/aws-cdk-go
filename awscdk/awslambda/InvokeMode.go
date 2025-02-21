package awslambda


// The invoke modes for a Lambda function.
//
// Example:
//   var fn function
//
//
//   fn.AddFunctionUrl(&FunctionUrlOptions{
//   	AuthType: lambda.FunctionUrlAuthType_NONE,
//   	InvokeMode: lambda.InvokeMode_RESPONSE_STREAM,
//   })
//
type InvokeMode string

const (
	// Default option.
	//
	// Lambda invokes your function using the Invoke API operation.
	// Invocation results are available when the payload is complete.
	// The maximum payload size is 6 MB.
	InvokeMode_BUFFERED InvokeMode = "BUFFERED"
	// Your function streams payload results as they become available.
	//
	// Lambda invokes your function using the InvokeWithResponseStream API operation.
	// The maximum response payload size is 20 MB, however, you can request a quota increase.
	InvokeMode_RESPONSE_STREAM InvokeMode = "RESPONSE_STREAM"
)

