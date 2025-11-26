package previewawsiotwirelessmixins


// ABP device object for LoRaWAN specification v1.0.x.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   abpV10xProperty := &AbpV10xProperty{
//   	DevAddr: jsii.String("devAddr"),
//   	SessionKeys: &SessionKeysAbpV10xProperty{
//   		AppSKey: jsii.String("appSKey"),
//   		NwkSKey: jsii.String("nwkSKey"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-abpv10x.html
//
type CfnWirelessDevicePropsMixin_AbpV10xProperty struct {
	// The DevAddr value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-abpv10x.html#cfn-iotwireless-wirelessdevice-abpv10x-devaddr
	//
	DevAddr *string `field:"optional" json:"devAddr" yaml:"devAddr"`
	// Session keys for ABP v1.0.x.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-abpv10x.html#cfn-iotwireless-wirelessdevice-abpv10x-sessionkeys
	//
	SessionKeys interface{} `field:"optional" json:"sessionKeys" yaml:"sessionKeys"`
}

