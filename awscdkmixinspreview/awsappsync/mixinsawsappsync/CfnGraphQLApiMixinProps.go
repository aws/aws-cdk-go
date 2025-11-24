package mixinsawsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnGraphQLApiPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGraphQLApiMixinProps := &CfnGraphQLApiMixinProps{
//   	AdditionalAuthenticationProviders: []interface{}{
//   		&AdditionalAuthenticationProviderProperty{
//   			AuthenticationType: jsii.String("authenticationType"),
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
//   	AuthenticationType: jsii.String("authenticationType"),
//   	EnhancedMetricsConfig: &EnhancedMetricsConfigProperty{
//   		DataSourceLevelMetricsBehavior: jsii.String("dataSourceLevelMetricsBehavior"),
//   		OperationLevelMetricsConfig: jsii.String("operationLevelMetricsConfig"),
//   		ResolverLevelMetricsBehavior: jsii.String("resolverLevelMetricsBehavior"),
//   	},
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
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
//   	Name: jsii.String("name"),
//   	OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   		AuthTtl: jsii.Number(123),
//   		ClientId: jsii.String("clientId"),
//   		IatTtl: jsii.Number(123),
//   		Issuer: jsii.String("issuer"),
//   	},
//   	OwnerContact: jsii.String("ownerContact"),
//   	QueryDepthLimit: jsii.Number(123),
//   	ResolverCountLimit: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
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
type CfnGraphQLApiMixinProps struct {
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
	// Security configuration for your GraphQL API.
	//
	// For allowed values (such as `API_KEY` , `AWS_IAM` , `AMAZON_COGNITO_USER_POOLS` , `OPENID_CONNECT` , or `AWS_LAMBDA` ), see [Security](https://docs.aws.amazon.com/appsync/latest/devguide/security.html) in the *AWS AppSync Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Enables and controls the enhanced metrics feature.
	//
	// Enhanced metrics emit granular data on API usage and performance such as AppSync request and error counts, latency, and cache hits/misses. All enhanced metric data is sent to your CloudWatch account, and you can configure the types of data that will be sent.
	//
	// Enhanced metrics can be configured at the resolver, data source, and operation levels. For more information, see [Monitoring and logging](https://docs.aws.amazon.com//appsync/latest/devguide/monitoring.html#cw-metrics) in the *AWS AppSync User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-enhancedmetricsconfig
	//
	EnhancedMetricsConfig interface{} `field:"optional" json:"enhancedMetricsConfig" yaml:"enhancedMetricsConfig"`
	// A map containing the list of resources with their properties and environment variables.
	//
	// For more information, see [Environmental variables](https://docs.aws.amazon.com/appsync/latest/devguide/environmental-variables.html) .
	//
	// *Pattern* : `^[A-Za-z]+\\w*$\\`
	//
	// *Minimum* : 2
	//
	// *Maximum* : 64.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Sets the value of the GraphQL API to enable ( `ENABLED` ) or disable ( `DISABLED` ) introspection.
	//
	// If no value is provided, the introspection configuration will be set to `ENABLED` by default. This field will produce an error if the operation attempts to use the introspection feature while this field is disabled.
	//
	// For more information about introspection, see [GraphQL introspection](https://docs.aws.amazon.com/https://graphql.org/learn/introspection/) .
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
	// The API name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
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
	// The maximum depth a query can have in a single request.
	//
	// Depth refers to the amount of nested levels allowed in the body of query. The default value is `0` (or unspecified), which indicates there's no depth limit. If you set a limit, it can be between `1` and `75` nested levels. This field will produce a limit error if the operation falls out of bounds. Note that fields can still be set to nullable or non-nullable. If a non-nullable field produces an error, the error will be thrown upwards to the first nullable field available.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html#cfn-appsync-graphqlapi-querydepthlimit
	//
	QueryDepthLimit *float64 `field:"optional" json:"queryDepthLimit" yaml:"queryDepthLimit"`
	// The maximum number of resolvers that can be invoked in a single request.
	//
	// The default value is `0` (or unspecified), which will set the limit to `10000` . When specified, the limit value can be between `1` and `10000` . This field will produce a limit error if the operation falls out of bounds.
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

