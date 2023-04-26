package awsappsync


// Configuration for AWS Lambda function authorization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaAuthorizerConfigProperty := &LambdaAuthorizerConfigProperty{
//   	AuthorizerResultTtlInSeconds: jsii.Number(123),
//   	AuthorizerUri: jsii.String("authorizerUri"),
//   	IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   }
//
type CfnGraphQLApi_LambdaAuthorizerConfigProperty struct {
	// The number of seconds a response should be cached for.
	//
	// The default is 0 seconds, which disables caching. If you don't specify a value for `authorizerResultTtlInSeconds` , the default value is used. The maximum value is one hour (3600 seconds). The Lambda function can override this by returning a `ttlOverride` key in its response.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// The ARN of the Lambda function to be called for authorization.
	//
	// This may be a standard Lambda ARN, a version ARN ( `.../v3` ) or alias ARN.
	//
	// *Note* : This Lambda function must have the following resource-based policy assigned to it. When configuring Lambda authorizers in the console, this is done for you. To do so with the AWS CLI , run the following:
	//
	// `aws lambda add-permission --function-name "arn:aws:lambda:us-east-2:111122223333:function:my-function" --statement-id "appsync" --principal appsync.amazonaws.com --action lambda:InvokeFunction`
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// A regular expression for validation of tokens before the Lambda function is called.
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
}

