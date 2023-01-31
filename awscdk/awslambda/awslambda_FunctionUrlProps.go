package awslambda


// Properties for a FunctionUrl.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var function_ function
//
//   functionUrlProps := &functionUrlProps{
//   	function: function_,
//
//   	// the properties below are optional
//   	authType: awscdk.Aws_lambda.functionUrlAuthType_AWS_IAM,
//   	cors: &functionUrlCorsOptions{
//   		allowCredentials: jsii.Boolean(false),
//   		allowedHeaders: []*string{
//   			jsii.String("allowedHeaders"),
//   		},
//   		allowedMethods: []httpMethod{
//   			awscdk.*Aws_lambda.*httpMethod_GET,
//   		},
//   		allowedOrigins: []*string{
//   			jsii.String("allowedOrigins"),
//   		},
//   		exposedHeaders: []*string{
//   			jsii.String("exposedHeaders"),
//   		},
//   		maxAge: duration,
//   	},
//   }
//
// Experimental.
type FunctionUrlProps struct {
	// The type of authentication that your function URL uses.
	// Experimental.
	AuthType FunctionUrlAuthType `field:"optional" json:"authType" yaml:"authType"`
	// The cross-origin resource sharing (CORS) settings for your function URL.
	// Experimental.
	Cors *FunctionUrlCorsOptions `field:"optional" json:"cors" yaml:"cors"`
	// The function to which this url refers.
	//
	// It can also be an `Alias` but not a `Version`.
	// Experimental.
	Function IFunction `field:"required" json:"function" yaml:"function"`
}

