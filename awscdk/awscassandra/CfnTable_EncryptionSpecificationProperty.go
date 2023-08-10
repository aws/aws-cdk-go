package awscassandra


// Specifies the encryption at rest option selected for the table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionSpecificationProperty := &EncryptionSpecificationProperty{
//   	EncryptionType: jsii.String("encryptionType"),
//
//   	// the properties below are optional
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-encryptionspecification.html
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-encryptionspecification.html#cfn-cassandra-table-encryptionspecification-encryptiontype
	//
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// Requires a `kms_key_identifier` in the format of a key ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cassandra-table-encryptionspecification.html#cfn-cassandra-table-encryptionspecification-kmskeyidentifier
	//
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
}

