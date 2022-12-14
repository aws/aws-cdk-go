package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigProperty := &encryptionConfigProperty{
//   	encryptionType: jsii.String("encryptionType"),
//   	keyId: jsii.String("keyId"),
//   }
//
type CfnInstanceStorageConfig_EncryptionConfigProperty struct {
	// `CfnInstanceStorageConfig.EncryptionConfigProperty.EncryptionType`.
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// `CfnInstanceStorageConfig.EncryptionConfigProperty.KeyId`.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

