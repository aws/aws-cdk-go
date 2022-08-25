package awsiotwireless


// OTAA device object for create APIs for v1.0.x.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   otaaV10xProperty := &otaaV10xProperty{
//   	appEui: jsii.String("appEui"),
//   	appKey: jsii.String("appKey"),
//   }
//
type CfnWirelessDevice_OtaaV10xProperty struct {
	// The AppEUI value, with pattern of `[a-fA-F0-9]{16}` .
	AppEui *string `field:"required" json:"appEui" yaml:"appEui"`
	// The AppKey is a secret key, which you should handle in a similar way as you would an application password.
	//
	// You can protect the AppKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	AppKey *string `field:"required" json:"appKey" yaml:"appKey"`
}

