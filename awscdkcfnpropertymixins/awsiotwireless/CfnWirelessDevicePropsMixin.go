package awsiotwireless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Provisions a wireless device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnWirelessDevicePropsMixin := awscdkcfnpropertymixins.Aws_iotwireless.NewCfnWirelessDevicePropsMixin(&CfnWirelessDeviceMixinProps{
//   	Description: jsii.String("description"),
//   	DestinationName: jsii.String("destinationName"),
//   	LastUplinkReceivedAt: jsii.String("lastUplinkReceivedAt"),
//   	LoRaWan: &LoRaWANDeviceProperty{
//   		AbpV10X: &AbpV10xProperty{
//   			DevAddr: jsii.String("devAddr"),
//   			SessionKeys: &SessionKeysAbpV10xProperty{
//   				AppSKey: jsii.String("appSKey"),
//   				NwkSKey: jsii.String("nwkSKey"),
//   			},
//   		},
//   		AbpV11: &AbpV11Property{
//   			DevAddr: jsii.String("devAddr"),
//   			SessionKeys: &SessionKeysAbpV11Property{
//   				AppSKey: jsii.String("appSKey"),
//   				FNwkSIntKey: jsii.String("fNwkSIntKey"),
//   				NwkSEncKey: jsii.String("nwkSEncKey"),
//   				SNwkSIntKey: jsii.String("sNwkSIntKey"),
//   			},
//   		},
//   		DevEui: jsii.String("devEui"),
//   		DeviceProfileId: jsii.String("deviceProfileId"),
//   		FPorts: &FPortsProperty{
//   			Applications: []interface{}{
//   				&ApplicationProperty{
//   					DestinationName: jsii.String("destinationName"),
//   					FPort: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		OtaaV10X: &OtaaV10xProperty{
//   			AppEui: jsii.String("appEui"),
//   			AppKey: jsii.String("appKey"),
//   		},
//   		OtaaV11: &OtaaV11Property{
//   			AppKey: jsii.String("appKey"),
//   			JoinEui: jsii.String("joinEui"),
//   			NwkKey: jsii.String("nwkKey"),
//   		},
//   		ServiceProfileId: jsii.String("serviceProfileId"),
//   	},
//   	Name: jsii.String("name"),
//   	Positioning: jsii.String("positioning"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThingArn: jsii.String("thingArn"),
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdevice.html
//
type CfnWirelessDevicePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnWirelessDeviceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWirelessDevicePropsMixin
type jsiiProxy_CfnWirelessDevicePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnWirelessDevicePropsMixin) Props() *CfnWirelessDeviceMixinProps {
	var returns *CfnWirelessDeviceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWirelessDevicePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTWireless::WirelessDevice`.
func NewCfnWirelessDevicePropsMixin(props *CfnWirelessDeviceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnWirelessDevicePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWirelessDevicePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWirelessDevicePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnWirelessDevicePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTWireless::WirelessDevice`.
func NewCfnWirelessDevicePropsMixin_Override(c CfnWirelessDevicePropsMixin, props *CfnWirelessDeviceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnWirelessDevicePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnWirelessDevicePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWirelessDevicePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnWirelessDevicePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWirelessDevicePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnWirelessDevicePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWirelessDevicePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWirelessDevicePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

