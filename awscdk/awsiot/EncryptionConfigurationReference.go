package awsiot


// A reference to a EncryptionConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationReference := &EncryptionConfigurationReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type EncryptionConfigurationReference struct {
	// The AccountId of the EncryptionConfiguration resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

