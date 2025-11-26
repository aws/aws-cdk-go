package previewawsgluemixins


// Contains configuration information for maintaining Data Catalog security.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataCatalogEncryptionSettingsProperty := &DataCatalogEncryptionSettingsProperty{
//   	ConnectionPasswordEncryption: &ConnectionPasswordEncryptionProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		ReturnConnectionPasswordEncrypted: jsii.Boolean(false),
//   	},
//   	EncryptionAtRest: &EncryptionAtRestProperty{
//   		CatalogEncryptionMode: jsii.String("catalogEncryptionMode"),
//   		CatalogEncryptionServiceRole: jsii.String("catalogEncryptionServiceRole"),
//   		SseAwsKmsKeyId: jsii.String("sseAwsKmsKeyId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html
//
type CfnDataCatalogEncryptionSettingsPropsMixin_DataCatalogEncryptionSettingsProperty struct {
	// When connection password protection is enabled, the Data Catalog uses a customer-provided key to encrypt the password as part of `CreateConnection` or `UpdateConnection` and store it in the `ENCRYPTED_PASSWORD` field in the connection properties.
	//
	// You can enable catalog encryption or only password encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-datacatalogencryptionsettings-connectionpasswordencryption
	//
	ConnectionPasswordEncryption interface{} `field:"optional" json:"connectionPasswordEncryption" yaml:"connectionPasswordEncryption"`
	// Specifies the encryption-at-rest configuration for the Data Catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-datacatalogencryptionsettings-encryptionatrest
	//
	EncryptionAtRest interface{} `field:"optional" json:"encryptionAtRest" yaml:"encryptionAtRest"`
}

