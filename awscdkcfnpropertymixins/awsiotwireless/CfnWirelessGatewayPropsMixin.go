package awsiotwireless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Provisions a wireless gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnWirelessGatewayPropsMixin := awscdkcfnpropertymixins.Aws_iotwireless.NewCfnWirelessGatewayPropsMixin(&CfnWirelessGatewayMixinProps{
//   	Description: jsii.String("description"),
//   	LastUplinkReceivedAt: jsii.String("lastUplinkReceivedAt"),
//   	LoRaWan: &LoRaWANGatewayProperty{
//   		GatewayEui: jsii.String("gatewayEui"),
//   		RfRegion: jsii.String("rfRegion"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThingArn: jsii.String("thingArn"),
//   	ThingName: jsii.String("thingName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessgateway.html
//
type CfnWirelessGatewayPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnWirelessGatewayMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWirelessGatewayPropsMixin
type jsiiProxy_CfnWirelessGatewayPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnWirelessGatewayPropsMixin) Props() *CfnWirelessGatewayMixinProps {
	var returns *CfnWirelessGatewayMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWirelessGatewayPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTWireless::WirelessGateway`.
func NewCfnWirelessGatewayPropsMixin(props *CfnWirelessGatewayMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnWirelessGatewayPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWirelessGatewayPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWirelessGatewayPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnWirelessGatewayPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTWireless::WirelessGateway`.
func NewCfnWirelessGatewayPropsMixin_Override(c CfnWirelessGatewayPropsMixin, props *CfnWirelessGatewayMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnWirelessGatewayPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnWirelessGatewayPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWirelessGatewayPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnWirelessGatewayPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWirelessGatewayPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnWirelessGatewayPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWirelessGatewayPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWirelessGatewayPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

