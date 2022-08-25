package awslambda


// The auth types for a function url.
//
// Example:
//   // Can be a Function or an Alias
//   var fn function
//
//
//   fnUrl := fn.addFunctionUrl(&functionUrlOptions{
//   	authType: lambda.functionUrlAuthType_NONE,
//   })
//
//   awscdk.NewCfnOutput(this, jsii.String("TheUrl"), &cfnOutputProps{
//   	value: fnUrl.url,
//   })
//
type FunctionUrlAuthType string

const (
	// Restrict access to authenticated IAM users only.
	FunctionUrlAuthType_AWS_IAM FunctionUrlAuthType = "AWS_IAM"
	// Bypass IAM authentication to create a public endpoint.
	FunctionUrlAuthType_NONE FunctionUrlAuthType = "NONE"
)

