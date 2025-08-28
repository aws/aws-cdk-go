package awsevidently


// If the project stores evaluation events in an Amazon S3 bucket, this structure stores the bucket name and bucket prefix.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationProperty := &S3DestinationProperty{
//   	BucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-project-s3destination.html
//
type CfnProject_S3DestinationProperty struct {
	// The name of the bucket in which Evidently stores evaluation events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-project-s3destination.html#cfn-evidently-project-s3destination-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The bucket prefix in which Evidently stores evaluation events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-project-s3destination.html#cfn-evidently-project-s3destination-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

