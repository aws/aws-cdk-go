package awsdatabrew


// Represents information on how DataBrew can find data, in either the AWS Glue Data Catalog or Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProperty := &InputProperty{
//   	DatabaseInputDefinition: &DatabaseInputDefinitionProperty{
//   		GlueConnectionName: jsii.String("glueConnectionName"),
//
//   		// the properties below are optional
//   		DatabaseTableName: jsii.String("databaseTableName"),
//   		QueryString: jsii.String("queryString"),
//   		TempDirectory: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			BucketOwner: jsii.String("bucketOwner"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	DataCatalogInputDefinition: &DataCatalogInputDefinitionProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		DatabaseName: jsii.String("databaseName"),
//   		TableName: jsii.String("tableName"),
//   		TempDirectory: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			BucketOwner: jsii.String("bucketOwner"),
//   			Key: jsii.String("key"),
//   		},
//   	},
//   	Metadata: &MetadataProperty{
//   		SourceArn: jsii.String("sourceArn"),
//   	},
//   	S3InputDefinition: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		BucketOwner: jsii.String("bucketOwner"),
//   		Key: jsii.String("key"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html
//
type CfnDataset_InputProperty struct {
	// Connection information for dataset input files stored in a database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-databaseinputdefinition
	//
	DatabaseInputDefinition interface{} `field:"optional" json:"databaseInputDefinition" yaml:"databaseInputDefinition"`
	// The AWS Glue Data Catalog parameters for the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-datacataloginputdefinition
	//
	DataCatalogInputDefinition interface{} `field:"optional" json:"dataCatalogInputDefinition" yaml:"dataCatalogInputDefinition"`
	// Contains additional resource information needed for specific datasets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The Amazon S3 location where the data is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-input.html#cfn-databrew-dataset-input-s3inputdefinition
	//
	S3InputDefinition interface{} `field:"optional" json:"s3InputDefinition" yaml:"s3InputDefinition"`
}

