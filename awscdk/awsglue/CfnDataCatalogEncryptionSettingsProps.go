package awsglue


// Properties for defining a `CfnDataCatalogEncryptionSettings`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   awscdk.NewCfnDataCatalogEncryptionSettings(this, jsii.String("Encryption"), &CfnDataCatalogEncryptionSettingsProps{
//   	CatalogId: jsii.String("my-catalog-id"),
//   	DataCatalogEncryptionSettings: &DataCatalogEncryptionSettingsProperty{
//   		EncryptionAtRest: &EncryptionAtRestProperty{
//   			CatalogEncryptionMode: jsii.String("SSE-KMS"),
//   		},
//   	},
//   })
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

