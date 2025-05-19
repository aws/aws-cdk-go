package awsappsync


// A `LambdaAuthorizerConfig` specifies how to authorize AWS AppSync API access when using the `AWS_LAMBDA` authorizer mode.
//
// Be aware that an AWS AppSync API can have only one AWS Lambda authorizer configured at a time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaAuthorizerConfigProperty := &LambdaAuthorizerConfigProperty{
//   	AuthorizerUri: jsii.String("authorizerUri"),
//
//   	// the properties below are optional
//   	AuthorizerResultTtlInSeconds: jsii.Number(123),
//   	IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-lambdaauthorizerconfig.html
//
type CfnApi_LambdaAuthorizerConfigProperty struct {
	// The Amazon Resource Name (ARN) of the Lambda function to be called for authorization.
	//
	// This can be a standard Lambda ARN, a version ARN ( `.../v3` ), or an alias ARN.
	//
	// *Note* : This Lambda function must have the following resource-based policy assigned to it. When configuring Lambda authorizers in the console, this is done for you. To use the AWS Command Line Interface ( AWS CLI ), run the following:
	//
	// `aws lambda add-permission --function-name "arn:aws:lambda:us-east-2:111122223333:function:my-function" --statement-id "appsync" --principal appsync.amazonaws.com --action lambda:InvokeFunction`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-lambdaauthorizerconfig.html#cfn-appsync-api-lambdaauthorizerconfig-authorizeruri
	//
	AuthorizerUri *string `field:"required" json:"authorizerUri" yaml:"authorizerUri"`
	// The number of seconds a response should be cached for.
	//
	// The default is 0 seconds, which disables caching. If you don't specify a value for `authorizerResultTtlInSeconds` , the default value is used. The maximum value is one hour (3600 seconds). The Lambda function can override this by returning a `ttlOverride` key in its response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-lambdaauthorizerconfig.html#cfn-appsync-api-lambdaauthorizerconfig-authorizerresultttlinseconds
	//
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// A regular expression for validation of tokens before the Lambda function is called.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-lambdaauthorizerconfig.html#cfn-appsync-api-lambdaauthorizerconfig-identityvalidationexpression
	//
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
}

