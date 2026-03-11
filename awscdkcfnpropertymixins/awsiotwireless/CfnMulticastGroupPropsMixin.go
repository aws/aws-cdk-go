package awsiotwireless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A multicast group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMulticastGroupPropsMixin := awscdkcfnpropertymixins.Aws_iotwireless.NewCfnMulticastGroupPropsMixin(&CfnMulticastGroupMixinProps{
//   	AssociateWirelessDevice: jsii.String("associateWirelessDevice"),
//   	Description: jsii.String("description"),
//   	DisassociateWirelessDevice: jsii.String("disassociateWirelessDevice"),
//   	LoRaWan: &LoRaWANProperty{
//   		DlClass: jsii.String("dlClass"),
//   		NumberOfDevicesInGroup: jsii.Number(123),
//   		NumberOfDevicesRequested: jsii.Number(123),
//   		RfRegion: jsii.String("rfRegion"),
//   	},
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-multicastgroup.html
//
type CfnMulticastGroupPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMulticastGroupMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMulticastGroupPropsMixin
type jsiiProxy_CfnMulticastGroupPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMulticastGroupPropsMixin) Props() *CfnMulticastGroupMixinProps {
	var returns *CfnMulticastGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMulticastGroupPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTWireless::MulticastGroup`.
func NewCfnMulticastGroupPropsMixin(props *CfnMulticastGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMulticastGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMulticastGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMulticastGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnMulticastGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTWireless::MulticastGroup`.
func NewCfnMulticastGroupPropsMixin_Override(c CfnMulticastGroupPropsMixin, props *CfnMulticastGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnMulticastGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMulticastGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMulticastGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnMulticastGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMulticastGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnMulticastGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMulticastGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMulticastGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

