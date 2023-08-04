package awslambda


// Options to add a url to a Lambda function.
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
type FunctionUrlOptions struct {
	// The type of authentication that your function URL uses.
	// Default: FunctionUrlAuthType.AWS_IAM
	//
	AuthType FunctionUrlAuthType `field:"optional" json:"authType" yaml:"authType"`
	// The cross-origin resource sharing (CORS) settings for your function URL.
	// Default: - No CORS configuration.
	//
	Cors *FunctionUrlCorsOptions `field:"optional" json:"cors" yaml:"cors"`
	// The type of invocation mode that your Lambda function uses.
	// Default: InvokeMode.BUFFERED
	//
	InvokeMode InvokeMode `field:"optional" json:"invokeMode" yaml:"invokeMode"`
}

