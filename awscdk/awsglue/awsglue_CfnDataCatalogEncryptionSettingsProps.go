package awsglue


// Properties for defining a `CfnDataCatalogEncryptionSettings`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataCatalogEncryptionSettingsProps := &cfnDataCatalogEncryptionSettingsProps{
//   	catalogId: jsii.String("catalogId"),
//   	dataCatalogEncryptionSettings: &dataCatalogEncryptionSettingsProperty{
//   		connectionPasswordEncryption: &connectionPasswordEncryptionProperty{
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   			returnConnectionPasswordEncrypted: jsii.Boolean(false),
//   		},
//   		encryptionAtRest: &encryptionAtRestProperty{
//   			catalogEncryptionMode: jsii.String("catalogEncryptionMode"),
//   			sseAwsKmsKeyId: jsii.String("sseAwsKmsKeyId"),
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

