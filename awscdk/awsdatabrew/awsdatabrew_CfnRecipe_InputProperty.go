package awsdatabrew


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProperty := &inputProperty{
//   	dataCatalogInputDefinition: &dataCatalogInputDefinitionProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//   		tableName: jsii.String("tableName"),
//   		tempDirectory: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   		},
//   	},
//   	s3InputDefinition: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnRecipe_InputProperty struct {
	// `CfnRecipe.InputProperty.DataCatalogInputDefinition`.
	DataCatalogInputDefinition interface{} `field:"optional" json:"dataCatalogInputDefinition" yaml:"dataCatalogInputDefinition"`
	// `CfnRecipe.InputProperty.S3InputDefinition`.
	S3InputDefinition interface{} `field:"optional" json:"s3InputDefinition" yaml:"s3InputDefinition"`
}

