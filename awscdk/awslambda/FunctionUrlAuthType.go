package awslambda


// The auth types for a function url.
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
type FunctionUrlAuthType string

const (
	// Restrict access to authenticated IAM users only.
	FunctionUrlAuthType_AWS_IAM FunctionUrlAuthType = "AWS_IAM"
	// Bypass IAM authentication to create a public endpoint.
	FunctionUrlAuthType_NONE FunctionUrlAuthType = "NONE"
)

