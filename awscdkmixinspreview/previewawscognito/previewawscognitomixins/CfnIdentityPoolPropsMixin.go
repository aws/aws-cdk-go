package previewawscognitomixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscognito/previewawscognitomixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Cognito::IdentityPool` resource creates an Amazon Cognito identity pool.
//
// To avoid deleting the resource accidentally from CloudFormation , use [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) and the [UpdateReplacePolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html) to retain the resource on deletion or replacement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var cognitoEvents interface{}
//   var supportedLoginProviders interface{}
//
//   cfnIdentityPoolPropsMixin := awscdkmixinspreview.Mixins.NewCfnIdentityPoolPropsMixin(&CfnIdentityPoolMixinProps{
//   	AllowClassicFlow: jsii.Boolean(false),
//   	AllowUnauthenticatedIdentities: jsii.Boolean(false),
//   	CognitoEvents: cognitoEvents,
//   	CognitoIdentityProviders: []interface{}{
//   		&CognitoIdentityProviderProperty{
//   			ClientId: jsii.String("clientId"),
//   			ProviderName: jsii.String("providerName"),
//   			ServerSideTokenCheck: jsii.Boolean(false),
//   		},
//   	},
//   	CognitoStreams: &CognitoStreamsProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		StreamingStatus: jsii.String("streamingStatus"),
//   		StreamName: jsii.String("streamName"),
//   	},
//   	DeveloperProviderName: jsii.String("developerProviderName"),
//   	IdentityPoolName: jsii.String("identityPoolName"),
//   	IdentityPoolTags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	OpenIdConnectProviderArns: []*string{
//   		jsii.String("openIdConnectProviderArns"),
//   	},
//   	PushSync: &PushSyncProperty{
//   		ApplicationArns: []*string{
//   			jsii.String("applicationArns"),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	SamlProviderArns: []*string{
//   		jsii.String("samlProviderArns"),
//   	},
//   	SupportedLoginProviders: supportedLoginProviders,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html
//
type CfnIdentityPoolPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIdentityPoolMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIdentityPoolPropsMixin
type jsiiProxy_CfnIdentityPoolPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIdentityPoolPropsMixin) Props() *CfnIdentityPoolMixinProps {
	var returns *CfnIdentityPoolMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdentityPoolPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::IdentityPool`.
func NewCfnIdentityPoolPropsMixin(props *CfnIdentityPoolMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIdentityPoolPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIdentityPoolPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIdentityPoolPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::IdentityPool`.
func NewCfnIdentityPoolPropsMixin_Override(c CfnIdentityPoolPropsMixin, props *CfnIdentityPoolMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIdentityPoolPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIdentityPoolPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdentityPoolPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnIdentityPoolPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdentityPoolPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIdentityPoolPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

