package awss3


// Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &DestinationProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	Format: jsii.String("format"),
//
//   	// the properties below are optional
//   	BucketAccountId: jsii.String("bucketAccountId"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html
//
type CfnBucket_DestinationProperty struct {
	// The Amazon Resource Name (ARN) of the bucket to which data is exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Specifies the file format used when exporting data to Amazon S3.
	//
	// *Allowed values* : `CSV` | `ORC` | `Parquet`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
	// The account ID that owns the destination S3 bucket.
	//
	// If no account ID is provided, the owner is not validated before exporting data.
	//
	// > Although this value is optional, we strongly recommend that you set it to help prevent problems if the destination bucket ownership changes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-bucketaccountid
	//
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// The prefix to use when exporting data.
	//
	// The prefix is prepended to all results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

