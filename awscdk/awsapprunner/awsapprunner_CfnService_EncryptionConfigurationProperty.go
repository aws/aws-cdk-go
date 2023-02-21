package awsapprunner


// Describes a custom encryption key that AWS App Runner uses to encrypt copies of the source repository and service logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &encryptionConfigurationProperty{
//   	kmsKey: jsii.String("kmsKey"),
//   }
//
type CfnService_EncryptionConfigurationProperty struct {
	// The ARN of the KMS key that's used for encryption.
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

