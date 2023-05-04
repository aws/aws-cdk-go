package awslambda


// Options to add a url to a Lambda function.
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
type FunctionUrlOptions struct {
	// The type of authentication that your function URL uses.
	AuthType FunctionUrlAuthType `field:"optional" json:"authType" yaml:"authType"`
	// The cross-origin resource sharing (CORS) settings for your function URL.
	Cors *FunctionUrlCorsOptions `field:"optional" json:"cors" yaml:"cors"`
}

