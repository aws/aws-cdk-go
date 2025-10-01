package awss3


// The destination information for a V1 S3 Metadata configuration.
//
// The destination table bucket must be in the same Region and AWS account as the general purpose bucket. The specified metadata table name must be unique within the `aws_s3_metadata` namespace in the destination table bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3TablesDestinationProperty := &S3TablesDestinationProperty{
//   	TableBucketArn: jsii.String("tableBucketArn"),
//   	TableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	TableArn: jsii.String("tableArn"),
//   	TableNamespace: jsii.String("tableNamespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-s3tablesdestination.html
//
type CfnBucket_S3TablesDestinationProperty struct {
	// The Amazon Resource Name (ARN) for the table bucket that's specified as the destination in the metadata table configuration.
	//
	// The destination table bucket must be in the same Region and AWS account as the general purpose bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-s3tablesdestination.html#cfn-s3-bucket-s3tablesdestination-tablebucketarn
	//
	TableBucketArn *string `field:"required" json:"tableBucketArn" yaml:"tableBucketArn"`
	// The name for the metadata table in your metadata table configuration.
	//
	// The specified metadata table name must be unique within the `aws_s3_metadata` namespace in the destination table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-s3tablesdestination.html#cfn-s3-bucket-s3tablesdestination-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The Amazon Resource Name (ARN) for the metadata table in the metadata table configuration.
	//
	// The specified metadata table name must be unique within the `aws_s3_metadata` namespace in the destination table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-s3tablesdestination.html#cfn-s3-bucket-s3tablesdestination-tablearn
	//
	TableArn *string `field:"optional" json:"tableArn" yaml:"tableArn"`
	// The table bucket namespace for the metadata table in your metadata table configuration.
	//
	// This value is always `aws_s3_metadata` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-s3tablesdestination.html#cfn-s3-bucket-s3tablesdestination-tablenamespace
	//
	TableNamespace *string `field:"optional" json:"tableNamespace" yaml:"tableNamespace"`
}

