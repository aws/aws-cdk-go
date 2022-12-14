package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCatalogConfigProperty := &dataCatalogConfigProperty{
//   	catalog: jsii.String("catalog"),
//   	database: jsii.String("database"),
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnFeatureGroup_DataCatalogConfigProperty struct {
	// `CfnFeatureGroup.DataCatalogConfigProperty.Catalog`.
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// `CfnFeatureGroup.DataCatalogConfigProperty.Database`.
	Database *string `field:"required" json:"database" yaml:"database"`
	// `CfnFeatureGroup.DataCatalogConfigProperty.TableName`.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

