package awslambda


// The auth types for a function url.
//
// Example:
//   // Can be a Function or an Alias
//   var fn function
//
//
//   fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
//   	AuthType: lambda.FunctionUrlAuthType_NONE,
//   })
//
//   awscdk.NewCfnOutput(this, jsii.String("TheUrl"), &CfnOutputProps{
//   	Value: fnUrl.Url,
//   })
//
type FunctionUrlAuthType string

const (
	// Restrict access to authenticated IAM users only.
	FunctionUrlAuthType_AWS_IAM FunctionUrlAuthType = "AWS_IAM"
	// Bypass IAM authentication to create a public endpoint.
	FunctionUrlAuthType_NONE FunctionUrlAuthType = "NONE"
)

