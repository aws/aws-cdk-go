package awslex


// Defines an Amazon S3 bucket location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	s3Bucket: jsii.String("s3Bucket"),
//   	s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   	// the properties below are optional
//   	s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   }
//
type CfnBot_S3LocationProperty struct {
	// The S3 bucket name.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The path and file name to the object in the S3 bucket.
	S3ObjectKey *string `field:"required" json:"s3ObjectKey" yaml:"s3ObjectKey"`
	// The version of the object in the S3 bucket.
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
}

