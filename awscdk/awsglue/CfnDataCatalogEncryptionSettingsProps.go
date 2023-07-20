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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-datacatalogencryptionsettings.html
//
type CfnDataCatalogEncryptionSettingsProps struct {
	// The ID of the Data Catalog in which the settings are created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-catalogid
	//
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// Contains configuration information for maintaining Data Catalog security.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-datacatalogencryptionsettings
	//
	DataCatalogEncryptionSettings interface{} `field:"required" json:"dataCatalogEncryptionSettings" yaml:"dataCatalogEncryptionSettings"`
}

