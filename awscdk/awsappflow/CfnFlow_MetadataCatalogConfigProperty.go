package awsappflow


// Specifies the configuration that Amazon AppFlow uses when it catalogs your data.
//
// When Amazon AppFlow catalogs your data, it stores metadata in a data catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataCatalogConfigProperty := &MetadataCatalogConfigProperty{
//   	GlueDataCatalog: &GlueDataCatalogProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		RoleArn: jsii.String("roleArn"),
//   		TablePrefix: jsii.String("tablePrefix"),
//   	},
//   }
//
type CfnFlow_MetadataCatalogConfigProperty struct {
	// Specifies the configuration that Amazon AppFlow uses when it catalogs your data with the AWS Glue Data Catalog .
	GlueDataCatalog interface{} `field:"optional" json:"glueDataCatalog" yaml:"glueDataCatalog"`
}

