package previewawsiotwirelessmixins


// OTAA device object for v1.1 for create APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   otaaV11Property := &OtaaV11Property{
//   	AppKey: jsii.String("appKey"),
//   	JoinEui: jsii.String("joinEui"),
//   	NwkKey: jsii.String("nwkKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-otaav11.html
//
type CfnWirelessDevicePropsMixin_OtaaV11Property struct {
	// The AppKey is a secret key, which you should handle in a similar way as you would an application password.
	//
	// You can protect the AppKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-otaav11.html#cfn-iotwireless-wirelessdevice-otaav11-appkey
	//
	AppKey *string `field:"optional" json:"appKey" yaml:"appKey"`
	// The JoinEUI value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-otaav11.html#cfn-iotwireless-wirelessdevice-otaav11-joineui
	//
	JoinEui *string `field:"optional" json:"joinEui" yaml:"joinEui"`
	// The NwkKey is a secret key, which you should handle in a similar way as you would an application password.
	//
	// You can protect the NwkKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-otaav11.html#cfn-iotwireless-wirelessdevice-otaav11-nwkkey
	//
	NwkKey *string `field:"optional" json:"nwkKey" yaml:"nwkKey"`
}

