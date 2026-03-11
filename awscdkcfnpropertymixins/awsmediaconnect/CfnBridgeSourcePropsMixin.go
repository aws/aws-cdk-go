package awsmediaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds sources to an existing bridge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnBridgeSourcePropsMixin := awscdkcfnpropertymixins.Aws_mediaconnect.NewCfnBridgeSourcePropsMixin(&CfnBridgeSourceMixinProps{
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
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html
//
type CfnBridgeSourcePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnBridgeSourceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBridgeSourcePropsMixin
type jsiiProxy_CfnBridgeSourcePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnBridgeSourcePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::BridgeSource`.
func NewCfnBridgeSourcePropsMixin(props *CfnBridgeSourceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnBridgeSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBridgeSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBridgeSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::BridgeSource`.
func NewCfnBridgeSourcePropsMixin_Override(c CfnBridgeSourcePropsMixin, props *CfnBridgeSourceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeSourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnBridgeSourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBridgeSourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeSourcePropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnBridgeSourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBridgeSourcePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

