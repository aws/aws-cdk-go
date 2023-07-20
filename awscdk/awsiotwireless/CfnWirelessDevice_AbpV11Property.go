package awsiotwireless


// ABP device object for create APIs for v1.1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   abpV11Property := &AbpV11Property{
//   	DevAddr: jsii.String("devAddr"),
//   	SessionKeys: &SessionKeysAbpV11Property{
//   		AppSKey: jsii.String("appSKey"),
//   		FNwkSIntKey: jsii.String("fNwkSIntKey"),
//   		NwkSEncKey: jsii.String("nwkSEncKey"),
//   		SNwkSIntKey: jsii.String("sNwkSIntKey"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-abpv11.html
//
type CfnWirelessDevice_AbpV11Property struct {
	// The DevAddr value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-abpv11.html#cfn-iotwireless-wirelessdevice-abpv11-devaddr
	//
	DevAddr *string `field:"required" json:"devAddr" yaml:"devAddr"`
	// Session keys for ABP v1.1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-abpv11.html#cfn-iotwireless-wirelessdevice-abpv11-sessionkeys
	//
	SessionKeys interface{} `field:"required" json:"sessionKeys" yaml:"sessionKeys"`
}

