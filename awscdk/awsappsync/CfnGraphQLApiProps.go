package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGraphQLApi`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphQLApiProps := &CfnGraphQLApiProps{
//   	AuthenticationType: jsii.String("authenticationType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AdditionalAuthenticationProviders: []interface{}{
//   		&AdditionalAuthenticationProviderProperty{
//   			AuthenticationType: jsii.String("authenticationType"),
//
//   			// the properties below are optional
//   			LambdaAuthorizerConfig: &LambdaAuthorizerConfigProperty{
//   				AuthorizerResultTtlInSeconds: jsii.Number(123),
//   				AuthorizerUri: jsii.String("authorizerUri"),
//   				IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   			},
//   			OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   				AuthTtl: jsii.Number(123),
//   				ClientId: jsii.String("clientId"),
//   				IatTtl: jsii.Number(123),
//   				Issuer: jsii.String("issuer"),
//   			},
//   			UserPoolConfig: &CognitoUserPoolConfigProperty{
//   				AppIdClientRegex: jsii.String("appIdClientRegex"),
//   				AwsRegion: jsii.String("awsRegion"),
//   				UserPoolId: jsii.String("userPoolId"),
//   			},
//   		},
//   	},
//   	LambdaAuthorizerConfig: &LambdaAuthorizerConfigProperty{
//   		AuthorizerResultTtlInSeconds: jsii.Number(123),
//   		AuthorizerUri: jsii.String("authorizerUri"),
//   		IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   	},
//   	LogConfig: &LogConfigProperty{
//   		CloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   		ExcludeVerboseContent: jsii.Boolean(false),
//   		FieldLogLevel: jsii.String("fieldLogLevel"),
//   	},
//   	OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   		AuthTtl: jsii.Number(123),
//   		ClientId: jsii.String("clientId"),
//   		IatTtl: jsii.Number(123),
//   		Issuer: jsii.String("issuer"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserPoolConfig: &UserPoolConfigProperty{
//   		AppIdClientRegex: jsii.String("appIdClientRegex"),
//   		AwsRegion: jsii.String("awsRegion"),
//   		DefaultAction: jsii.String("defaultAction"),
//   		UserPoolId: jsii.String("userPoolId"),
//   	},
//   	Visibility: jsii.String("visibility"),
//   	XrayEnabled: jsii.Boolean(false),
//   }
//
type CfnGraphQLApiProps struct {
	// Security configuration for your GraphQL API.
	//
	// For allowed values (such as `API_KEY` , `AWS_IAM` , `AMAZON_COGNITO_USER_POOLS` , `OPENID_CONNECT` , or `AWS_LAMBDA` ), see [Security](https://docs.aws.amazon.com/appsync/latest/devguide/security.html) in the *AWS AppSync Developer Guide* .
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The API name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of additional authentication providers for the `GraphqlApi` API.
	AdditionalAuthenticationProviders interface{} `field:"optional" json:"additionalAuthenticationProviders" yaml:"additionalAuthenticationProviders"`
	// A `LambdaAuthorizerConfig` holds configuration on how to authorize AWS AppSync API access when using the `AWS_LAMBDA` authorizer mode.
	//
	// Be aware that an AWS AppSync API may have only one Lambda authorizer configured at a time.
	LambdaAuthorizerConfig interface{} `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// The Amazon CloudWatch Logs configuration.
	LogConfig interface{} `field:"optional" json:"logConfig" yaml:"logConfig"`
	// The OpenID Connect configuration.
	OpenIdConnectConfig interface{} `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
	// An arbitrary set of tags (key-value pairs) for this GraphQL API.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Optional authorization configuration for using Amazon Cognito user pools with your GraphQL endpoint.
	UserPoolConfig interface{} `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
	// `AWS::AppSync::GraphQLApi.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// A flag indicating whether to use AWS X-Ray tracing for this `GraphqlApi` .
	XrayEnabled interface{} `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

