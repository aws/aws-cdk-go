package awsevidently


// If the project stores evaluation events in an Amazon S3 bucket, this structure stores the bucket name and bucket prefix.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationProperty := &s3DestinationProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnProject_S3DestinationProperty struct {
	// The name of the bucket in which Evidently stores evaluation events.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The bucket prefix in which Evidently stores evaluation events.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

