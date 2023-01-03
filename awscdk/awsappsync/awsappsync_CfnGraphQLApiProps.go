package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnGraphQLApi`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphQLApiProps := &cfnGraphQLApiProps{
//   	authenticationType: jsii.String("authenticationType"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	additionalAuthenticationProviders: []interface{}{
//   		&additionalAuthenticationProviderProperty{
//   			authenticationType: jsii.String("authenticationType"),
//
//   			// the properties below are optional
//   			lambdaAuthorizerConfig: &lambdaAuthorizerConfigProperty{
//   				authorizerResultTtlInSeconds: jsii.Number(123),
//   				authorizerUri: jsii.String("authorizerUri"),
//   				identityValidationExpression: jsii.String("identityValidationExpression"),
//   			},
//   			openIdConnectConfig: &openIDConnectConfigProperty{
//   				authTtl: jsii.Number(123),
//   				clientId: jsii.String("clientId"),
//   				iatTtl: jsii.Number(123),
//   				issuer: jsii.String("issuer"),
//   			},
//   			userPoolConfig: &cognitoUserPoolConfigProperty{
//   				appIdClientRegex: jsii.String("appIdClientRegex"),
//   				awsRegion: jsii.String("awsRegion"),
//   				userPoolId: jsii.String("userPoolId"),
//   			},
//   		},
//   	},
//   	lambdaAuthorizerConfig: &lambdaAuthorizerConfigProperty{
//   		authorizerResultTtlInSeconds: jsii.Number(123),
//   		authorizerUri: jsii.String("authorizerUri"),
//   		identityValidationExpression: jsii.String("identityValidationExpression"),
//   	},
//   	logConfig: &logConfigProperty{
//   		cloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   		excludeVerboseContent: jsii.Boolean(false),
//   		fieldLogLevel: jsii.String("fieldLogLevel"),
//   	},
//   	openIdConnectConfig: &openIDConnectConfigProperty{
//   		authTtl: jsii.Number(123),
//   		clientId: jsii.String("clientId"),
//   		iatTtl: jsii.Number(123),
//   		issuer: jsii.String("issuer"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userPoolConfig: &userPoolConfigProperty{
//   		appIdClientRegex: jsii.String("appIdClientRegex"),
//   		awsRegion: jsii.String("awsRegion"),
//   		defaultAction: jsii.String("defaultAction"),
//   		userPoolId: jsii.String("userPoolId"),
//   	},
//   	xrayEnabled: jsii.Boolean(false),
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
	// A flag indicating whether to use AWS X-Ray tracing for this `GraphqlApi` .
	XrayEnabled interface{} `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

