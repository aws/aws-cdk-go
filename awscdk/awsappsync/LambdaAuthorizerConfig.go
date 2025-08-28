package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
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
//   appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("api"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.test.graphql"))),
//   	AuthorizationConfig: &AuthorizationConfig{
//   		DefaultAuthorization: &AuthorizationMode{
//   			AuthorizationType: appsync.AuthorizationType_LAMBDA,
//   			LambdaAuthorizerConfig: &LambdaAuthorizerConfig{
//   				Handler: authFunction,
//   			},
//   		},
//   	},
//   })
//
type LambdaAuthorizerConfig struct {
	// The authorizer lambda function.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html
	//
	Handler awslambda.IFunction `field:"required" json:"handler" yaml:"handler"`
	// How long the results are cached.
	//
	// Disable caching by setting this to 0.
	// Default: Duration.minutes(5)
	//
	ResultsCacheTtl awscdk.Duration `field:"optional" json:"resultsCacheTtl" yaml:"resultsCacheTtl"`
	// A regular expression for validation of tokens before the Lambda function is called.
	// Default: - no regex filter will be applied.
	//
	ValidationRegex *string `field:"optional" json:"validationRegex" yaml:"validationRegex"`
}

