package previewawsappsyncmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappsync/previewawsappsyncmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AppSync::Api` resource creates an AWS AppSync API that you can use for an AWS AppSync API with your preferred configuration, such as an Event API that provides real-time message publishing and message subscriptions over WebSockets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApiPropsMixin := awscdkmixinspreview.Mixins.NewCfnApiPropsMixin(&CfnApiMixinProps{
//   	EventConfig: &EventConfigProperty{
//   		AuthProviders: []interface{}{
//   			&AuthProviderProperty{
//   				AuthType: jsii.String("authType"),
//   				CognitoConfig: &CognitoConfigProperty{
//   					AppIdClientRegex: jsii.String("appIdClientRegex"),
//   					AwsRegion: jsii.String("awsRegion"),
//   					UserPoolId: jsii.String("userPoolId"),
//   				},
//   				LambdaAuthorizerConfig: &LambdaAuthorizerConfigProperty{
//   					AuthorizerResultTtlInSeconds: jsii.Number(123),
//   					AuthorizerUri: jsii.String("authorizerUri"),
//   					IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   				},
//   				OpenIdConnectConfig: &OpenIDConnectConfigProperty{
//   					AuthTtl: jsii.Number(123),
//   					ClientId: jsii.String("clientId"),
//   					IatTtl: jsii.Number(123),
//   					Issuer: jsii.String("issuer"),
//   				},
//   			},
//   		},
//   		ConnectionAuthModes: []interface{}{
//   			&AuthModeProperty{
//   				AuthType: jsii.String("authType"),
//   			},
//   		},
//   		DefaultPublishAuthModes: []interface{}{
//   			&AuthModeProperty{
//   				AuthType: jsii.String("authType"),
//   			},
//   		},
//   		DefaultSubscribeAuthModes: []interface{}{
//   			&AuthModeProperty{
//   				AuthType: jsii.String("authType"),
//   			},
//   		},
//   		LogConfig: &EventLogConfigProperty{
//   			CloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OwnerContact: jsii.String("ownerContact"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-api.html
//
type CfnApiPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApiMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApiPropsMixin
type jsiiProxy_CfnApiPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApiPropsMixin) Props() *CfnApiMixinProps {
	var returns *CfnApiMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppSync::Api`.
func NewCfnApiPropsMixin(props *CfnApiMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApiPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApiPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApiPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnApiPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppSync::Api`.
func NewCfnApiPropsMixin_Override(c CfnApiPropsMixin, props *CfnApiMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnApiPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApiPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApiPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnApiPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApiPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnApiPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApiPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnApiPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

