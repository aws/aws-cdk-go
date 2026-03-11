package awspinpoint

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A *channel* is a type of platform that you can deliver messages to.
//
// To send an SMS text message, you send the message through the SMS channel. Before you can use Amazon Pinpoint to send text messages, you have to enable the SMS channel for an Amazon Pinpoint application.
//
// The SMSChannel resource represents the status, sender ID, and other settings for the SMS channel for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnSMSChannelPropsMixin := awscdkcfnpropertymixins.Aws_pinpoint.NewCfnSMSChannelPropsMixin(&CfnSMSChannelMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	Enabled: jsii.Boolean(false),
//   	SenderId: jsii.String("senderId"),
//   	ShortCode: jsii.String("shortCode"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smschannel.html
//
type CfnSMSChannelPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSMSChannelMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSMSChannelPropsMixin
type jsiiProxy_CfnSMSChannelPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSMSChannelPropsMixin) Props() *CfnSMSChannelMixinProps {
	var returns *CfnSMSChannelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMSChannelPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Pinpoint::SMSChannel`.
func NewCfnSMSChannelPropsMixin(props *CfnSMSChannelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSMSChannelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSMSChannelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSMSChannelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_pinpoint.CfnSMSChannelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Pinpoint::SMSChannel`.
func NewCfnSMSChannelPropsMixin_Override(c CfnSMSChannelPropsMixin, props *CfnSMSChannelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_pinpoint.CfnSMSChannelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSMSChannelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSMSChannelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_pinpoint.CfnSMSChannelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSMSChannelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_pinpoint.CfnSMSChannelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSMSChannelPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSMSChannelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

