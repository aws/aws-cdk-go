package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   offlineStoreConfigProperty := &offlineStoreConfigProperty{
//   	s3StorageConfig: &s3StorageConfigProperty{
//   		s3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//
//   	// the properties below are optional
//   	dataCatalogConfig: &dataCatalogConfigProperty{
//   		catalog: jsii.String("catalog"),
//   		database: jsii.String("database"),
//   		tableName: jsii.String("tableName"),
//   	},
//   	disableGlueTableCreation: jsii.Boolean(false),
//   	tableFormat: jsii.String("tableFormat"),
//   }
//
type CfnFeatureGroup_OfflineStoreConfigProperty struct {
	// `CfnFeatureGroup.OfflineStoreConfigProperty.S3StorageConfig`.
	S3StorageConfig interface{} `field:"required" json:"s3StorageConfig" yaml:"s3StorageConfig"`
	// `CfnFeatureGroup.OfflineStoreConfigProperty.DataCatalogConfig`.
	DataCatalogConfig interface{} `field:"optional" json:"dataCatalogConfig" yaml:"dataCatalogConfig"`
	// `CfnFeatureGroup.OfflineStoreConfigProperty.DisableGlueTableCreation`.
	DisableGlueTableCreation interface{} `field:"optional" json:"disableGlueTableCreation" yaml:"disableGlueTableCreation"`
	// `CfnFeatureGroup.OfflineStoreConfigProperty.TableFormat`.
	TableFormat *string `field:"optional" json:"tableFormat" yaml:"tableFormat"`
}

