package awsdatabrew


// Represents secondary inputs in a UNION transform.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secondaryInputProperty := &SecondaryInputProperty{
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
type CfnRecipe_SecondaryInputProperty struct {
	// The AWS Glue Data Catalog parameters for the data.
	DataCatalogInputDefinition interface{} `field:"optional" json:"dataCatalogInputDefinition" yaml:"dataCatalogInputDefinition"`
	// The Amazon S3 location where the data is stored.
	S3InputDefinition interface{} `field:"optional" json:"s3InputDefinition" yaml:"s3InputDefinition"`
}

