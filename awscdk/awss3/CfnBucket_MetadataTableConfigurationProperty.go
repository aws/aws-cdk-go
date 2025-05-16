package awss3


// The metadata table configuration of an Amazon S3 general purpose bucket.
//
// For more information, see [Accelerating data discovery with S3 Metadata](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html) and [Setting up permissions for configuring metadata tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataTableConfigurationProperty := &MetadataTableConfigurationProperty{
//   	S3TablesDestination: &S3TablesDestinationProperty{
//   		TableBucketArn: jsii.String("tableBucketArn"),
//   		TableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		TableArn: jsii.String("tableArn"),
//   		TableNamespace: jsii.String("tableNamespace"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadatatableconfiguration.html
//
type CfnBucket_MetadataTableConfigurationProperty struct {
	// The destination information for the metadata table configuration.
	//
	// The destination table bucket must be in the same Region and AWS account as the general purpose bucket. The specified metadata table name must be unique within the `aws_s3_metadata` namespace in the destination table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadatatableconfiguration.html#cfn-s3-bucket-metadatatableconfiguration-s3tablesdestination
	//
	S3TablesDestination interface{} `field:"required" json:"s3TablesDestination" yaml:"s3TablesDestination"`
}

