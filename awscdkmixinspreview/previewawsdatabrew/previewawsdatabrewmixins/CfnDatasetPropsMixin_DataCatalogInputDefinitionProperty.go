package previewawsdatabrewmixins


// Represents how metadata stored in the AWS Glue Data Catalog is defined in a DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataCatalogInputDefinitionProperty := &DataCatalogInputDefinitionProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   	TempDirectory: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		BucketOwner: jsii.String("bucketOwner"),
//   		Key: jsii.String("key"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datacataloginputdefinition.html
//
type CfnDatasetPropsMixin_DataCatalogInputDefinitionProperty struct {
	// The unique identifier of the AWS account that holds the Data Catalog that stores the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datacataloginputdefinition.html#cfn-databrew-dataset-datacataloginputdefinition-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of a database in the Data Catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datacataloginputdefinition.html#cfn-databrew-dataset-datacataloginputdefinition-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of a database table in the Data Catalog.
	//
	// This table corresponds to a DataBrew dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datacataloginputdefinition.html#cfn-databrew-dataset-datacataloginputdefinition-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// An Amazon location that AWS Glue Data Catalog can use as a temporary directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datacataloginputdefinition.html#cfn-databrew-dataset-datacataloginputdefinition-tempdirectory
	//
	TempDirectory interface{} `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

