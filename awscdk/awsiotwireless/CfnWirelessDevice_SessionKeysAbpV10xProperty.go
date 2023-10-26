package awsiotwireless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionKeysAbpV10xProperty := &SessionKeysAbpV10xProperty{
//   	AppSKey: jsii.String("appSKey"),
//   	NwkSKey: jsii.String("nwkSKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv10x.html
//
type CfnWirelessDevice_SessionKeysAbpV10xProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv10x.html#cfn-iotwireless-wirelessdevice-sessionkeysabpv10x-appskey
	//
	AppSKey *string `field:"required" json:"appSKey" yaml:"appSKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv10x.html#cfn-iotwireless-wirelessdevice-sessionkeysabpv10x-nwkskey
	//
	NwkSKey *string `field:"required" json:"nwkSKey" yaml:"nwkSKey"`
}

