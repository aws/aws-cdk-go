package awsdatabrew


// Represents information on how DataBrew can find data, in either the AWS Glue Data Catalog or Amazon S3.
//
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
	// The AWS Glue Data Catalog parameters for the data.
	DataCatalogInputDefinition interface{} `field:"optional" json:"dataCatalogInputDefinition" yaml:"dataCatalogInputDefinition"`
	// The Amazon S3 location where the data is stored.
	S3InputDefinition interface{} `field:"optional" json:"s3InputDefinition" yaml:"s3InputDefinition"`
}

