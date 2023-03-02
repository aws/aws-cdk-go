package awsappflow


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
	// `CfnFlow.MetadataCatalogConfigProperty.GlueDataCatalog`.
	GlueDataCatalog interface{} `field:"optional" json:"glueDataCatalog" yaml:"glueDataCatalog"`
}

