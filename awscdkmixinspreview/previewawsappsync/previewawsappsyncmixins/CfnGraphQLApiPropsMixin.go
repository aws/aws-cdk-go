package previewawsappsyncmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappsync/previewawsappsyncmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppSync::GraphQLApi` resource creates a new AWS AppSync GraphQL API.
//
// This is the top-level construct for your application. For more information, see [Quick Start](https://docs.aws.amazon.com/appsync/latest/devguide/quickstart.html) in the *AWS AppSync Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGraphQLApiPropsMixin := awscdkmixinspreview.Mixins.NewCfnGraphQLApiPropsMixin(&CfnGraphQLApiMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html
//
type CfnGraphQLApiPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGraphQLApiMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGraphQLApiPropsMixin
type jsiiProxy_CfnGraphQLApiPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGraphQLApiPropsMixin) Props() *CfnGraphQLApiMixinProps {
	var returns *CfnGraphQLApiMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApiPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppSync::GraphQLApi`.
func NewCfnGraphQLApiPropsMixin(props *CfnGraphQLApiMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGraphQLApiPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGraphQLApiPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGraphQLApiPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnGraphQLApiPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppSync::GraphQLApi`.
func NewCfnGraphQLApiPropsMixin_Override(c CfnGraphQLApiPropsMixin, props *CfnGraphQLApiMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnGraphQLApiPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGraphQLApiPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGraphQLApiPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnGraphQLApiPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGraphQLApiPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnGraphQLApiPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGraphQLApiPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGraphQLApiPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

