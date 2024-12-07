package awscdkpipesalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for `S3LogDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3LogDestinationProps := &S3LogDestinationProps{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	BucketOwner: jsii.String("bucketOwner"),
//   	OutputFormat: pipes_alpha.S3OutputFormat_PLAIN,
//   	Prefix: jsii.String("prefix"),
//   }
//
// Experimental.
type S3LogDestinationProps struct {
	// The S3 bucket to deliver the log records for the pipe.
	//
	// The bucket can be in the same or a different AWS Account. If the bucket is in
	// a different acccount, specify `bucketOwner`. You must also allow access to the
	// Pipes role in the bucket policy of the cross-account bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-s3logdestination.html#cfn-pipes-pipe-s3logdestination-bucketname
	//
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The AWS Account that owns the Amazon S3 bucket to which EventBridge delivers the log records for the pipe.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-s3logdestination.html#cfn-pipes-pipe-s3logdestination-bucketowner
	//
	// Default: - account ID derived from `bucket`.
	//
	// Experimental.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The format for the log records.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-s3logdestination.html#cfn-pipes-pipe-s3logdestination-outputformat
	//
	// Default: `S3OutputFormat.JSON`
	//
	// Experimental.
	OutputFormat S3OutputFormat `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// The prefix text with which to begin Amazon S3 log object names.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-s3logdestination.html#cfn-pipes-pipe-s3logdestination-prefix
	//
	// Default: - no prefix.
	//
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

