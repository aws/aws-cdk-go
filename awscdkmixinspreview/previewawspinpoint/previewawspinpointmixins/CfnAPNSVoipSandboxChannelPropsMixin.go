package previewawspinpointmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspinpoint/previewawspinpointmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A *channel* is a type of platform that you can deliver messages to.
//
// You can use the APNs VoIP sandbox channel to send VoIP notification messages to the sandbox environment of the Apple Push Notification service (APNs). Before you can use Amazon Pinpoint to send VoIP notifications to the APNs sandbox environment, you have to enable the APNs VoIP sandbox channel for an Amazon Pinpoint application.
//
// The APNSVoipSandboxChannel resource represents the status and authentication settings of the APNs VoIP sandbox channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAPNSVoipSandboxChannelPropsMixin := awscdkmixinspreview.Mixins.NewCfnAPNSVoipSandboxChannelPropsMixin(&CfnAPNSVoipSandboxChannelMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	BundleId: jsii.String("bundleId"),
//   	Certificate: jsii.String("certificate"),
//   	DefaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	Enabled: jsii.Boolean(false),
//   	PrivateKey: jsii.String("privateKey"),
//   	TeamId: jsii.String("teamId"),
//   	TokenKey: jsii.String("tokenKey"),
//   	TokenKeyId: jsii.String("tokenKeyId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnsvoipsandboxchannel.html
//
type CfnAPNSVoipSandboxChannelPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAPNSVoipSandboxChannelMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAPNSVoipSandboxChannelPropsMixin
type jsiiProxy_CfnAPNSVoipSandboxChannelPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannelPropsMixin) Props() *CfnAPNSVoipSandboxChannelMixinProps {
	var returns *CfnAPNSVoipSandboxChannelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAPNSVoipSandboxChannelPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Pinpoint::APNSVoipSandboxChannel`.
func NewCfnAPNSVoipSandboxChannelPropsMixin(props *CfnAPNSVoipSandboxChannelMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAPNSVoipSandboxChannelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAPNSVoipSandboxChannelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAPNSVoipSandboxChannelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnAPNSVoipSandboxChannelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Pinpoint::APNSVoipSandboxChannel`.
func NewCfnAPNSVoipSandboxChannelPropsMixin_Override(c CfnAPNSVoipSandboxChannelPropsMixin, props *CfnAPNSVoipSandboxChannelMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnAPNSVoipSandboxChannelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAPNSVoipSandboxChannelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAPNSVoipSandboxChannelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnAPNSVoipSandboxChannelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAPNSVoipSandboxChannelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pinpoint.mixins.CfnAPNSVoipSandboxChannelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAPNSVoipSandboxChannelPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAPNSVoipSandboxChannelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

