package awsdatabrew


// Represents information on how DataBrew can find data, in either the AWS Glue Data Catalog or Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProperty := &inputProperty{
//   	databaseInputDefinition: &databaseInputDefinitionProperty{
//   		glueConnectionName: jsii.String("glueConnectionName"),
//
//   		// the properties below are optional
//   		databaseTableName: jsii.String("databaseTableName"),
//   		queryString: jsii.String("queryString"),
//   		tempDirectory: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   		},
//   	},
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
//   	metadata: &metadataProperty{
//   		sourceArn: jsii.String("sourceArn"),
//   	},
//   	s3InputDefinition: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnDataset_InputProperty struct {
	// Connection information for dataset input files stored in a database.
	DatabaseInputDefinition interface{} `field:"optional" json:"databaseInputDefinition" yaml:"databaseInputDefinition"`
	// The AWS Glue Data Catalog parameters for the data.
	DataCatalogInputDefinition interface{} `field:"optional" json:"dataCatalogInputDefinition" yaml:"dataCatalogInputDefinition"`
	// Contains additional resource information needed for specific datasets.
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The Amazon S3 location where the data is stored.
	S3InputDefinition interface{} `field:"optional" json:"s3InputDefinition" yaml:"s3InputDefinition"`
}

