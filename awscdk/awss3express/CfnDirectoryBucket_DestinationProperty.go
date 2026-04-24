package awss3express


// Specifies information about where to publish inventory reports for an Amazon S3 Express bucket.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-destination.html
//
type CfnDirectoryBucket_DestinationProperty struct {
	// The Amazon Resource Name (ARN) of the destination Amazon S3 bucket to which data is exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-destination.html#cfn-s3express-directorybucket-destination-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Specifies the file format used when exporting data to Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-destination.html#cfn-s3express-directorybucket-destination-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
	// The account ID that owns the destination S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-destination.html#cfn-s3express-directorybucket-destination-bucketaccountid
	//
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// The prefix to use when exporting data.
	//
	// The prefix is prepended to all results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-destination.html#cfn-s3express-directorybucket-destination-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

