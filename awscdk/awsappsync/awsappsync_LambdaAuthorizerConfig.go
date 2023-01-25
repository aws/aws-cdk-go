package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Configuration for Lambda authorization in AppSync.
//
// Note that you can only have a single AWS Lambda function configured to authorize your API.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var authFunction function
//
//
//   appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
//   	name: jsii.String("api"),
//   	schema: appsync.schema.fromAsset(path.join(__dirname, jsii.String("appsync.test.graphql"))),
//   	authorizationConfig: &authorizationConfig{
//   		defaultAuthorization: &authorizationMode{
//   			authorizationType: appsync.authorizationType_LAMBDA,
//   			lambdaAuthorizerConfig: &lambdaAuthorizerConfig{
//   				handler: authFunction,
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type LambdaAuthorizerConfig struct {
	// The authorizer lambda function.
	//
	// Note: This Lambda function must have the following resource-based policy assigned to it.
	// When configuring Lambda authorizers in the console, this is done for you.
	// To do so with the AWS CLI, run the following:
	//
	// `aws lambda add-permission --function-name "arn:aws:lambda:us-east-2:111122223333:function:my-function" --statement-id "appsync" --principal appsync.amazonaws.com --action lambda:InvokeFunction`
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html
	//
	// Experimental.
	Handler awslambda.IFunction `field:"required" json:"handler" yaml:"handler"`
	// How long the results are cached.
	//
	// Disable caching by setting this to 0.
	// Experimental.
	ResultsCacheTtl awscdk.Duration `field:"optional" json:"resultsCacheTtl" yaml:"resultsCacheTtl"`
	// A regular expression for validation of tokens before the Lambda function is called.
	// Experimental.
	ValidationRegex *string `field:"optional" json:"validationRegex" yaml:"validationRegex"`
}

