package awsopensearch


// Configuration for an S3 Glue Data Catalog data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3GlueDataCatalogProperty := &S3GlueDataCatalogProperty{
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearch-datasource-s3gluedatacatalog.html
//
type CfnDataSource_S3GlueDataCatalogProperty struct {
	// The ARN of the IAM role that grants OpenSearch Service permission to access the Glue Data Catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearch-datasource-s3gluedatacatalog.html#cfn-opensearch-datasource-s3gluedatacatalog-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

