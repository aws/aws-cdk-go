package awsnimblestudio


// Configuration of the encryption method that is used for the studio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   studioEncryptionConfigurationProperty := &studioEncryptionConfigurationProperty{
//   	keyType: jsii.String("keyType"),
//
//   	// the properties below are optional
//   	keyArn: jsii.String("keyArn"),
//   }
//
type CfnStudio_StudioEncryptionConfigurationProperty struct {
	// The type of KMS key that is used to encrypt studio data.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The ARN for a KMS key that is used to encrypt studio data.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

