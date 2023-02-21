package awsiotwireless


// OTAA device object for v1.1 for create APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   otaaV11Property := &otaaV11Property{
//   	appKey: jsii.String("appKey"),
//   	joinEui: jsii.String("joinEui"),
//   	nwkKey: jsii.String("nwkKey"),
//   }
//
type CfnWirelessDevice_OtaaV11Property struct {
	// The AppKey is a secret key, which you should handle in a similar way as you would an application password.
	//
	// You can protect the AppKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	AppKey *string `field:"required" json:"appKey" yaml:"appKey"`
	// The JoinEUI value.
	JoinEui *string `field:"required" json:"joinEui" yaml:"joinEui"`
	// The NwkKey is a secret key, which you should handle in a similar way as you would an application password.
	//
	// You can protect the NwkKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	NwkKey *string `field:"required" json:"nwkKey" yaml:"nwkKey"`
}

