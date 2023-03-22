package awscassandra


// Specifies the encryption at rest option selected for the table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionSpecificationProperty := &encryptionSpecificationProperty{
//   	encryptionType: jsii.String("encryptionType"),
//
//   	// the properties below are optional
//   	kmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   }
//
type CfnTable_EncryptionSpecificationProperty struct {
	// The encryption at rest options for the table.
	//
	// - *AWS owned key* (default) - `AWS_OWNED_KMS_KEY`
	// - *Customer managed key* - `CUSTOMER_MANAGED_KMS_KEY`
	//
	// > If you choose `CUSTOMER_MANAGED_KMS_KEY` , a `kms_key_identifier` in the format of a key ARN is required.
	//
	// Valid values: `CUSTOMER_MANAGED_KMS_KEY` | `AWS_OWNED_KMS_KEY` .
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// Requires a `kms_key_identifier` in the format of a key ARN.
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
}

