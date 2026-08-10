package awsdynamodb


// Properties for defining a `CfnExport`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnExportProps := &CfnExportProps{
//   	S3Bucket: jsii.String("s3Bucket"),
//   	TableArn: jsii.String("tableArn"),
//
//   	// the properties below are optional
//   	ExportFormat: jsii.String("exportFormat"),
//   	ExportType: jsii.String("exportType"),
//   	S3BucketOwner: jsii.String("s3BucketOwner"),
//   	S3Prefix: jsii.String("s3Prefix"),
//   	S3SseAlgorithm: jsii.String("s3SseAlgorithm"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-export.html
//
type CfnExportProps struct {
	// The name of the Amazon S3 bucket containing the export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-export.html#cfn-dynamodb-export-s3bucket
	//
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The Amazon Resource Name (ARN) of the table that was exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-export.html#cfn-dynamodb-export-tablearn
	//
	TableArn *string `field:"required" json:"tableArn" yaml:"tableArn"`
	// The format of the exported data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-export.html#cfn-dynamodb-export-exportformat
	//
	ExportFormat *string `field:"optional" json:"exportFormat" yaml:"exportFormat"`
	// The type of export that was performed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-export.html#cfn-dynamodb-export-exporttype
	//
	ExportType *string `field:"optional" json:"exportType" yaml:"exportType"`
	// The ID of the Amazon Web Services account that owns the bucket containing the export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-export.html#cfn-dynamodb-export-s3bucketowner
	//
	S3BucketOwner *string `field:"optional" json:"s3BucketOwner" yaml:"s3BucketOwner"`
	// The Amazon S3 bucket prefix used as the file name and path of the exported snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-export.html#cfn-dynamodb-export-s3prefix
	//
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
	// Type of encryption used on the bucket where export data is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-export.html#cfn-dynamodb-export-s3ssealgorithm
	//
	S3SseAlgorithm *string `field:"optional" json:"s3SseAlgorithm" yaml:"s3SseAlgorithm"`
}

