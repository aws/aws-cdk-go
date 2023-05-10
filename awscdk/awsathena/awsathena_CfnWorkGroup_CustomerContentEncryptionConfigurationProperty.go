package awsathena


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerContentEncryptionConfigurationProperty := &CustomerContentEncryptionConfigurationProperty{
//   	KmsKey: jsii.String("kmsKey"),
//   }
//
type CfnWorkGroup_CustomerContentEncryptionConfigurationProperty struct {
	// `CfnWorkGroup.CustomerContentEncryptionConfigurationProperty.KmsKey`.
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

