package awsiotwireless


// Session keys for ABP v1.1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionKeysAbpV11Property := &SessionKeysAbpV11Property{
//   	AppSKey: jsii.String("appSKey"),
//   	FNwkSIntKey: jsii.String("fNwkSIntKey"),
//   	NwkSEncKey: jsii.String("nwkSEncKey"),
//   	SNwkSIntKey: jsii.String("sNwkSIntKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv11.html
//
type CfnWirelessDevice_SessionKeysAbpV11Property struct {
	// The AppSKey is a secret key, which you should handle in a similar way as you would an application password.
	//
	// You can protect the AppSKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv11.html#cfn-iotwireless-wirelessdevice-sessionkeysabpv11-appskey
	//
	AppSKey *string `field:"required" json:"appSKey" yaml:"appSKey"`
	// The FNwkSIntKey is a secret key, which you should handle in a similar way as you would an application password.
	//
	// You can protect the FNwkSIntKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv11.html#cfn-iotwireless-wirelessdevice-sessionkeysabpv11-fnwksintkey
	//
	FNwkSIntKey *string `field:"required" json:"fNwkSIntKey" yaml:"fNwkSIntKey"`
	// The NwkSEncKey is a secret key, which you should handle in a similar way as you would an application password.
	//
	// You can protect the NwkSEncKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv11.html#cfn-iotwireless-wirelessdevice-sessionkeysabpv11-nwksenckey
	//
	NwkSEncKey *string `field:"required" json:"nwkSEncKey" yaml:"nwkSEncKey"`
	// The SNwkSIntKey is a secret key, which you should handle in a similar way as you would an application password.
	//
	// You can protect the SNwkSIntKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv11.html#cfn-iotwireless-wirelessdevice-sessionkeysabpv11-snwksintkey
	//
	SNwkSIntKey *string `field:"required" json:"sNwkSIntKey" yaml:"sNwkSIntKey"`
}

