package awsglue


// Specifies the encryption-at-rest configuration for the Data Catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionAtRestProperty := &EncryptionAtRestProperty{
//   	CatalogEncryptionMode: jsii.String("catalogEncryptionMode"),
//   	CatalogEncryptionServiceRole: jsii.String("catalogEncryptionServiceRole"),
//   	SseAwsKmsKeyId: jsii.String("sseAwsKmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-encryptionatrest.html
//
type CfnDataCatalogEncryptionSettings_EncryptionAtRestProperty struct {
	// The encryption-at-rest mode for encrypting Data Catalog data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-encryptionatrest.html#cfn-glue-datacatalogencryptionsettings-encryptionatrest-catalogencryptionmode
	//
	CatalogEncryptionMode *string `field:"optional" json:"catalogEncryptionMode" yaml:"catalogEncryptionMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-encryptionatrest.html#cfn-glue-datacatalogencryptionsettings-encryptionatrest-catalogencryptionservicerole
	//
	CatalogEncryptionServiceRole *string `field:"optional" json:"catalogEncryptionServiceRole" yaml:"catalogEncryptionServiceRole"`
	// The ID of the AWS KMS key to use for encryption at rest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-encryptionatrest.html#cfn-glue-datacatalogencryptionsettings-encryptionatrest-sseawskmskeyid
	//
	SseAwsKmsKeyId *string `field:"optional" json:"sseAwsKmsKeyId" yaml:"sseAwsKmsKeyId"`
}

