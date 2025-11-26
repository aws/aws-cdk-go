package previewawsmediaconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmediaconnect/previewawsmediaconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds sources to an existing bridge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBridgeSourcePropsMixin := awscdkmixinspreview.Mixins.NewCfnBridgeSourcePropsMixin(&CfnBridgeSourceMixinProps{
//   	BridgeArn: jsii.String("bridgeArn"),
//   	FlowSource: &BridgeFlowSourceProperty{
//   		FlowArn: jsii.String("flowArn"),
//   		FlowVpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   			VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	NetworkSource: &BridgeNetworkSourceProperty{
//   		MulticastIp: jsii.String("multicastIp"),
//   		MulticastSourceSettings: &MulticastSourceSettingsProperty{
//   			MulticastSourceIp: jsii.String("multicastSourceIp"),
//   		},
//   		NetworkName: jsii.String("networkName"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html
//
type CfnBridgeSourcePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBridgeSourceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBridgeSourcePropsMixin
type jsiiProxy_CfnBridgeSourcePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBridgeSourcePropsMixin) Props() *CfnBridgeSourceMixinProps {
	var returns *CfnBridgeSourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBridgeSourcePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::BridgeSource`.
func NewCfnBridgeSourcePropsMixin(props *CfnBridgeSourceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBridgeSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBridgeSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBridgeSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnBridgeSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::BridgeSource`.
func NewCfnBridgeSourcePropsMixin_Override(c CfnBridgeSourcePropsMixin, props *CfnBridgeSourceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnBridgeSourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBridgeSourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBridgeSourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnBridgeSourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBridgeSourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnBridgeSourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBridgeSourcePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnBridgeSourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

