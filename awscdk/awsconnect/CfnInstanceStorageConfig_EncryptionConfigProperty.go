package awsconnect


// The encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigProperty := &EncryptionConfigProperty{
//   	EncryptionType: jsii.String("encryptionType"),
//   	KeyId: jsii.String("keyId"),
//   }
//
type CfnInstanceStorageConfig_EncryptionConfigProperty struct {
	// The type of encryption.
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// The full ARN of the encryption key.
	//
	// > Be sure to provide the full ARN of the encryption key, not just the ID.
	// >
	// > Amazon Connect supports only KMS keys with the default key spec of [`SYMMETRIC_DEFAULT`](https://docs.aws.amazon.com/kms/latest/developerguide/asymmetric-key-specs.html#key-spec-symmetric-default) .
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

