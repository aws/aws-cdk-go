package awslambda


// Properties for defining a `CfnUrl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUrlProps := &CfnUrlProps{
//   	AuthType: jsii.String("authType"),
//   	TargetFunctionArn: jsii.String("targetFunctionArn"),
//
//   	// the properties below are optional
//   	Cors: &CorsProperty{
//   		AllowCredentials: jsii.Boolean(false),
//   		AllowHeaders: []*string{
//   			jsii.String("allowHeaders"),
//   		},
//   		AllowMethods: []*string{
//   			jsii.String("allowMethods"),
//   		},
//   		AllowOrigins: []*string{
//   			jsii.String("allowOrigins"),
//   		},
//   		ExposeHeaders: []*string{
//   			jsii.String("exposeHeaders"),
//   		},
//   		MaxAge: jsii.Number(123),
//   	},
//   	InvokeMode: jsii.String("invokeMode"),
//   	Qualifier: jsii.String("qualifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html
//
type CfnUrlProps struct {
	// The type of authentication that your function URL uses.
	//
	// Set to `AWS_IAM` if you want to restrict access to authenticated users only. Set to `NONE` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Security and auth model for Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html#cfn-lambda-url-authtype
	//
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// The name of the Lambda function.
	//
	// **Name formats** - *Function name* - `my-function` .
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:my-function` .
	// - *Partial ARN* - `123456789012:function:my-function` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html#cfn-lambda-url-targetfunctionarn
	//
	TargetFunctionArn *string `field:"required" json:"targetFunctionArn" yaml:"targetFunctionArn"`
	// The [Cross-Origin Resource Sharing (CORS)](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for your function URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html#cfn-lambda-url-cors
	//
	Cors interface{} `field:"optional" json:"cors" yaml:"cors"`
	// Use one of the following options:.
	//
	// - `BUFFERED` – This is the default option. Lambda invokes your function using the `Invoke` API operation. Invocation results are available when the payload is complete. The maximum payload size is 6 MB.
	// - `RESPONSE_STREAM` – Your function streams payload results as they become available. Lambda invokes your function using the `InvokeWithResponseStream` API operation. The maximum response payload size is 20 MB, however, you can [request a quota increase](https://docs.aws.amazon.com/servicequotas/latest/userguide/request-quota-increase.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html#cfn-lambda-url-invokemode
	//
	InvokeMode *string `field:"optional" json:"invokeMode" yaml:"invokeMode"`
	// The alias name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-url.html#cfn-lambda-url-qualifier
	//
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

