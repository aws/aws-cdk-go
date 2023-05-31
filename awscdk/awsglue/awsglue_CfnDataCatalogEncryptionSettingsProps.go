package awsglue


// Properties for defining a `CfnDataCatalogEncryptionSettings`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataCatalogEncryptionSettingsProps := &CfnDataCatalogEncryptionSettingsProps{
//   	CatalogId: jsii.String("catalogId"),
//   	DataCatalogEncryptionSettings: &DataCatalogEncryptionSettingsProperty{
//   		ConnectionPasswordEncryption: &ConnectionPasswordEncryptionProperty{
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			ReturnConnectionPasswordEncrypted: jsii.Boolean(false),
//   		},
//   		EncryptionAtRest: &EncryptionAtRestProperty{
//   			CatalogEncryptionMode: jsii.String("catalogEncryptionMode"),
//   			SseAwsKmsKeyId: jsii.String("sseAwsKmsKeyId"),
//   		},
//   	},
//   }
//
type CfnDataCatalogEncryptionSettingsProps struct {
	// The ID of the Data Catalog in which the settings are created.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// Contains configuration information for maintaining Data Catalog security.
	DataCatalogEncryptionSettings interface{} `field:"required" json:"dataCatalogEncryptionSettings" yaml:"dataCatalogEncryptionSettings"`
}

