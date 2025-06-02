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
//   var handler function
//
//
//   iamProvider := &AppSyncAuthProvider{
//   	AuthorizationType: appsync.AppSyncAuthorizationType_IAM,
//   }
//
//   apiKeyProvider := &AppSyncAuthProvider{
//   	AuthorizationType: appsync.AppSyncAuthorizationType_API_KEY,
//   }
//
//   lambdaProvider := &AppSyncAuthProvider{
//   	AuthorizationType: appsync.AppSyncAuthorizationType_LAMBDA,
//   	LambdaAuthorizerConfig: &AppSyncLambdaAuthorizerConfig{
//   		Handler: *Handler,
//   		ResultsCacheTtl: awscdk.Duration_Minutes(jsii.Number(6)),
//   		ValidationRegex: jsii.String("test"),
//   	},
//   }
//
//   api := appsync.NewEventApi(this, jsii.String("api"), &EventApiProps{
//   	ApiName: jsii.String("api"),
//   	AuthorizationConfig: &EventApiAuthConfig{
//   		// set auth providers
//   		AuthProviders: []appSyncAuthProvider{
//   			iamProvider,
//   			apiKeyProvider,
//   			lambdaProvider,
//   		},
//   		ConnectionAuthModeTypes: []appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_IAM,
//   		},
//   		DefaultPublishAuthModeTypes: []*appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   		DefaultSubscribeAuthModeTypes: []*appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_LAMBDA,
//   		},
//   	},
//   })
//
//   api.AddChannelNamespace(jsii.String("default"))
//
type AppSyncLambdaAuthorizerConfig struct {
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

