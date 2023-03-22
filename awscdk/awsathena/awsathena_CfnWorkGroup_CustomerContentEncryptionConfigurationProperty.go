package awsathena


// Specifies the KMS key that is used to encrypt the user's data stores in Athena.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerContentEncryptionConfigurationProperty := &customerContentEncryptionConfigurationProperty{
//   	kmsKey: jsii.String("kmsKey"),
//   }
//
type CfnWorkGroup_CustomerContentEncryptionConfigurationProperty struct {
	// The KMS key that is used to encrypt the user's data stores in Athena.
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

