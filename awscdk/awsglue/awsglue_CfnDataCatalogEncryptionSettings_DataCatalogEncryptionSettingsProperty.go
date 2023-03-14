package awsglue


// Contains configuration information for maintaining Data Catalog security.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCatalogEncryptionSettingsProperty := &dataCatalogEncryptionSettingsProperty{
//   	connectionPasswordEncryption: &connectionPasswordEncryptionProperty{
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		returnConnectionPasswordEncrypted: jsii.Boolean(false),
//   	},
//   	encryptionAtRest: &encryptionAtRestProperty{
//   		catalogEncryptionMode: jsii.String("catalogEncryptionMode"),
//   		sseAwsKmsKeyId: jsii.String("sseAwsKmsKeyId"),
//   	},
//   }
//
type CfnDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsProperty struct {
	// When connection password protection is enabled, the Data Catalog uses a customer-provided key to encrypt the password as part of `CreateConnection` or `UpdateConnection` and store it in the `ENCRYPTED_PASSWORD` field in the connection properties.
	//
	// You can enable catalog encryption or only password encryption.
	ConnectionPasswordEncryption interface{} `field:"optional" json:"connectionPasswordEncryption" yaml:"connectionPasswordEncryption"`
	// Specifies the encryption-at-rest configuration for the Data Catalog.
	EncryptionAtRest interface{} `field:"optional" json:"encryptionAtRest" yaml:"encryptionAtRest"`
}

