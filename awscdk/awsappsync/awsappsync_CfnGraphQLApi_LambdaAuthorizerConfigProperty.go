package awsappsync


// Configuration for AWS Lambda function authorization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaAuthorizerConfigProperty := &lambdaAuthorizerConfigProperty{
//   	authorizerResultTtlInSeconds: jsii.Number(123),
//   	authorizerUri: jsii.String("authorizerUri"),
//   	identityValidationExpression: jsii.String("identityValidationExpression"),
//   }
//
type CfnGraphQLApi_LambdaAuthorizerConfigProperty struct {
	// The number of seconds a response should be cached for.
	//
	// The default is 5 minutes (300 seconds). The Lambda function can override this by returning a `ttlOverride` key in its response. A value of 0 disables caching of responses.
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

