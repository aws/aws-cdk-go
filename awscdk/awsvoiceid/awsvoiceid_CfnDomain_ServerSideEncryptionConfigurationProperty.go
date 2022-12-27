package awsvoiceid


// The configuration containing information about the customer-managed KMS Key used for encrypting customer data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSideEncryptionConfigurationProperty := &serverSideEncryptionConfigurationProperty{
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnDomain_ServerSideEncryptionConfigurationProperty struct {
	// The identifier of the KMS Key you want Voice ID to use to encrypt your data.
	KmsKeyId *string `field:"required" json:"kmsKeyId" yaml:"kmsKeyId"`
}

