package mixinsawsiotwireless


// OTAA device object for v1.0.x.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   otaaV10xProperty := &OtaaV10xProperty{
//   	AppEui: jsii.String("appEui"),
//   	AppKey: jsii.String("appKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-otaav10x.html
//
type CfnWirelessDevicePropsMixin_OtaaV10xProperty struct {
	// The AppEUI value.
	//
	// You specify this value when using LoRaWAN versions v1.0.2 or v1.0.3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-otaav10x.html#cfn-iotwireless-wirelessdevice-otaav10x-appeui
	//
	AppEui *string `field:"optional" json:"appEui" yaml:"appEui"`
	// The AppKey value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-otaav10x.html#cfn-iotwireless-wirelessdevice-otaav10x-appkey
	//
	AppKey *string `field:"optional" json:"appKey" yaml:"appKey"`
}

