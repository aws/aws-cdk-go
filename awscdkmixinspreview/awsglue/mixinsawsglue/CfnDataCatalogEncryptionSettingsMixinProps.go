package mixinsawsglue


// Properties for CfnDataCatalogEncryptionSettingsPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataCatalogEncryptionSettingsMixinProps := &CfnDataCatalogEncryptionSettingsMixinProps{
//   	CatalogId: jsii.String("catalogId"),
//   	DataCatalogEncryptionSettings: &DataCatalogEncryptionSettingsProperty{
//   		ConnectionPasswordEncryption: &ConnectionPasswordEncryptionProperty{
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			ReturnConnectionPasswordEncrypted: jsii.Boolean(false),
//   		},
//   		EncryptionAtRest: &EncryptionAtRestProperty{
//   			CatalogEncryptionMode: jsii.String("catalogEncryptionMode"),
//   			CatalogEncryptionServiceRole: jsii.String("catalogEncryptionServiceRole"),
//   			SseAwsKmsKeyId: jsii.String("sseAwsKmsKeyId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-datacatalogencryptionsettings.html
//
type CfnDataCatalogEncryptionSettingsMixinProps struct {
	// The ID of the Data Catalog in which the settings are created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Contains configuration information for maintaining Data Catalog security.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-datacatalogencryptionsettings
	//
	DataCatalogEncryptionSettings interface{} `field:"optional" json:"dataCatalogEncryptionSettings" yaml:"dataCatalogEncryptionSettings"`
}

