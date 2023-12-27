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
//   	ApiType: jsii.String("apiType"),
//   	IntrospectionConfig: jsii.String("introspectionConfig"),
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
//   	MergedApiExecutionRoleArn: jsii.String("mergedApiExecutionRoleArn"),
//   	OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   		AuthTtl: jsii.Number(123),
//   		ClientId: jsii.String("clientId"),
//   		IatTtl: jsii.Number(123),
//   		Issuer: jsii.String("issuer"),
//   	},
//   	OwnerContact: jsii.String("ownerContact"),
//   	QueryDepthLimit: jsii.Number(123),
//   	ResolverCountLimit: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html
//
type CfnGraphQLApiProps struct {
	// Security configuration for your GraphQL API.
	//
	// For allowed values (such as `API_KEY` , `AWS_IAM` , `AMAZON_COGNITO_USER_POOLS` , `OPENID_CONNECT` , or `AWS_LAMBDA` ), see [Security](https://docs.aws.amazon.com/appsync/latest/devguide/security.html) in the *AWS AppSync Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-authenticationtype
	//
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// The API name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of additional authentication providers for the `GraphqlApi` API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-additionalauthenticationproviders
	//
	AdditionalAuthenticationProviders interface{} `field:"optional" json:"additionalAuthenticationProviders" yaml:"additionalAuthenticationProviders"`
	// The value that indicates whether the GraphQL API is a standard API ( `GRAPHQL` ) or merged API ( `MERGED` ).
	//
	// *WARNING* : If the `ApiType` has not been defined, *explicitly* setting it to `GRAPHQL` in a template/stack update will result in an API replacement and new DNS values.
	//
	// The following values are valid:
	//
	// `GRAPHQL | MERGED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-apitype
	//
	ApiType *string `field:"optional" json:"apiType" yaml:"apiType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-introspectionconfig
	//
	IntrospectionConfig *string `field:"optional" json:"introspectionConfig" yaml:"introspectionConfig"`
	// A `LambdaAuthorizerConfig` holds configuration on how to authorize AWS AppSync API access when using the `AWS_LAMBDA` authorizer mode.
	//
	// Be aware that an AWS AppSync API may have only one Lambda authorizer configured at a time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-lambdaauthorizerconfig
	//
	LambdaAuthorizerConfig interface{} `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// The Amazon CloudWatch Logs configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-logconfig
	//
	LogConfig interface{} `field:"optional" json:"logConfig" yaml:"logConfig"`
	// The AWS Identity and Access Management service role ARN for a merged API.
	//
	// The AppSync service assumes this role on behalf of the Merged API to validate access to source APIs at runtime and to prompt the `AUTO_MERGE` to update the merged API endpoint with the source API changes automatically.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-mergedapiexecutionrolearn
	//
	MergedApiExecutionRoleArn *string `field:"optional" json:"mergedApiExecutionRoleArn" yaml:"mergedApiExecutionRoleArn"`
	// The OpenID Connect configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-openidconnectconfig
	//
	OpenIdConnectConfig interface{} `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
	// The owner contact information for an API resource.
	//
	// This field accepts any string input with a length of 0 - 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-ownercontact
	//
	OwnerContact *string `field:"optional" json:"ownerContact" yaml:"ownerContact"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-querydepthlimit
	//
	QueryDepthLimit *float64 `field:"optional" json:"queryDepthLimit" yaml:"queryDepthLimit"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-resolvercountlimit
	//
	ResolverCountLimit *float64 `field:"optional" json:"resolverCountLimit" yaml:"resolverCountLimit"`
	// An arbitrary set of tags (key-value pairs) for this GraphQL API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Optional authorization configuration for using Amazon Cognito user pools with your GraphQL endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-userpoolconfig
	//
	UserPoolConfig interface{} `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
	// Sets the scope of the GraphQL API to public ( `GLOBAL` ) or private ( `PRIVATE` ).
	//
	// By default, the scope is set to `Global` if no value is provided.
	//
	// *WARNING* : If `Visibility` has not been defined, *explicitly* setting it to `GLOBAL` in a template/stack update will result in an API replacement and new DNS values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// A flag indicating whether to use AWS X-Ray tracing for this `GraphqlApi` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-xrayenabled
	//
	XrayEnabled interface{} `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

