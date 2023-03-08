package awsglue


// Specifies the encryption-at-rest configuration for the Data Catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionAtRestProperty := &encryptionAtRestProperty{
//   	catalogEncryptionMode: jsii.String("catalogEncryptionMode"),
//   	sseAwsKmsKeyId: jsii.String("sseAwsKmsKeyId"),
//   }
//
type CfnDataCatalogEncryptionSettings_EncryptionAtRestProperty struct {
	// The encryption-at-rest mode for encrypting Data Catalog data.
	CatalogEncryptionMode *string `field:"optional" json:"catalogEncryptionMode" yaml:"catalogEncryptionMode"`
	// The ID of the AWS KMS key to use for encryption at rest.
	SseAwsKmsKeyId *string `field:"optional" json:"sseAwsKmsKeyId" yaml:"sseAwsKmsKeyId"`
}

