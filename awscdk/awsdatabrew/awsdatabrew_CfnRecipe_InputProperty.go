package awsdatabrew


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProperty := &InputProperty{
//   	DataCatalogInputDefinition: &DataCatalogInputDefinitionProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		DatabaseName: jsii.String("databaseName"),
//   		TableName: jsii.String("tableName"),
//   		TempDirectory: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	S3InputDefinition: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		Key: jsii.String("key"),
//   	},
//   }
//
type CfnRecipe_InputProperty struct {
	// `CfnRecipe.InputProperty.DataCatalogInputDefinition`.
	DataCatalogInputDefinition interface{} `field:"optional" json:"dataCatalogInputDefinition" yaml:"dataCatalogInputDefinition"`
	// `CfnRecipe.InputProperty.S3InputDefinition`.
	S3InputDefinition interface{} `field:"optional" json:"s3InputDefinition" yaml:"s3InputDefinition"`
}

