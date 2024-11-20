package awslambda


// Options to add a url to a Lambda function.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
//   	AuthType: lambda.FunctionUrlAuthType_NONE,
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewFunctionUrlOrigin(fnUrl),
//   	},
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

