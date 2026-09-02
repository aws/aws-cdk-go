package awsopensearch


// The type of data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceTypeProperty := &DataSourceTypeProperty{
//   	S3GlueDataCatalog: &S3GlueDataCatalogProperty{
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearch-datasource-datasourcetype.html
//
type CfnDataSource_DataSourceTypeProperty struct {
	// Configuration for an S3 Glue Data Catalog data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearch-datasource-datasourcetype.html#cfn-opensearch-datasource-datasourcetype-s3gluedatacatalog
	//
	S3GlueDataCatalog interface{} `field:"optional" json:"s3GlueDataCatalog" yaml:"s3GlueDataCatalog"`
}

