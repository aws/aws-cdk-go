package previewawsiotwirelessmixins


// Session keys for ABP v1.0.x.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sessionKeysAbpV10xProperty := &SessionKeysAbpV10xProperty{
//   	AppSKey: jsii.String("appSKey"),
//   	NwkSKey: jsii.String("nwkSKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv10x.html
//
type CfnWirelessDevicePropsMixin_SessionKeysAbpV10xProperty struct {
	// The AppSKey value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv10x.html#cfn-iotwireless-wirelessdevice-sessionkeysabpv10x-appskey
	//
	AppSKey *string `field:"optional" json:"appSKey" yaml:"appSKey"`
	// The NwkKey value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv10x.html#cfn-iotwireless-wirelessdevice-sessionkeysabpv10x-nwkskey
	//
	NwkSKey *string `field:"optional" json:"nwkSKey" yaml:"nwkSKey"`
}

