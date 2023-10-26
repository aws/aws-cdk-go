package awsiotwireless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnWirelessDevice_AbpV10xProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-abpv10x.html#cfn-iotwireless-wirelessdevice-abpv10x-devaddr
	//
	DevAddr *string `field:"required" json:"devAddr" yaml:"devAddr"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-abpv10x.html#cfn-iotwireless-wirelessdevice-abpv10x-sessionkeys
	//
	SessionKeys interface{} `field:"required" json:"sessionKeys" yaml:"sessionKeys"`
}

