package awsmediapackagev2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediapackagev2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a channel to receive content.
//
// After it's created, a channel provides static input URLs. These URLs remain the same throughout the lifetime of the channel, regardless of any failures or upgrades that might occur. Use these URLs to configure the outputs of your upstream encoder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnChannelPropsMixin := awscdkcfnpropertymixins.Aws_mediapackagev2.NewCfnChannelPropsMixin(&CfnChannelMixinProps{
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//   	Description: jsii.String("description"),
//   	InputSwitchConfiguration: &InputSwitchConfigurationProperty{
//   		MqcsInputSwitching: jsii.Boolean(false),
//   		PreferredInput: jsii.Number(123),
//   	},
//   	InputType: jsii.String("inputType"),
//   	OutputHeaderConfiguration: &OutputHeaderConfigurationProperty{
//   		PublishMqcs: jsii.Boolean(false),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channel.html
//
type CfnChannelPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnChannelMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnChannelPropsMixin
type jsiiProxy_CfnChannelPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnChannelPropsMixin) Props() *CfnChannelMixinProps {
	var returns *CfnChannelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannelPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaPackageV2::Channel`.
func NewCfnChannelPropsMixin(props *CfnChannelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnChannelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnChannelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnChannelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnChannelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaPackageV2::Channel`.
func NewCfnChannelPropsMixin_Override(c CfnChannelPropsMixin, props *CfnChannelMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnChannelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnChannelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnChannelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnChannelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnChannelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnChannelPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnChannelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

