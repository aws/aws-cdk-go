package awslambda


// Properties for defining a `CfnUrl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUrlProps := &cfnUrlProps{
//   	authType: jsii.String("authType"),
//   	targetFunctionArn: jsii.String("targetFunctionArn"),
//
//   	// the properties below are optional
//   	cors: &corsProperty{
//   		allowCredentials: jsii.Boolean(false),
//   		allowHeaders: []*string{
//   			jsii.String("allowHeaders"),
//   		},
//   		allowMethods: []*string{
//   			jsii.String("allowMethods"),
//   		},
//   		allowOrigins: []*string{
//   			jsii.String("allowOrigins"),
//   		},
//   		exposeHeaders: []*string{
//   			jsii.String("exposeHeaders"),
//   		},
//   		maxAge: jsii.Number(123),
//   	},
//   	invokeMode: jsii.String("invokeMode"),
//   	qualifier: jsii.String("qualifier"),
//   }
//
type CfnUrlProps struct {
	// The type of authentication that your function URL uses.
	//
	// Set to `AWS_IAM` if you want to restrict access to authenticated `IAM` users only. Set to `NONE` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Security and auth model for Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html) .
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// The name of the Lambda function.
	//
	// **Name formats** - *Function name* - `my-function` .
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:my-function` .
	// - *Partial ARN* - `123456789012:function:my-function` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	TargetFunctionArn *string `field:"required" json:"targetFunctionArn" yaml:"targetFunctionArn"`
	// The [Cross-Origin Resource Sharing (CORS)](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for your function URL.
	Cors interface{} `field:"optional" json:"cors" yaml:"cors"`
	// `AWS::Lambda::Url.InvokeMode`.
	InvokeMode *string `field:"optional" json:"invokeMode" yaml:"invokeMode"`
	// The alias name.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

