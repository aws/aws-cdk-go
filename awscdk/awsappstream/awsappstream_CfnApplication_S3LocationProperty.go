package awsappstream


// The S3 location of the application icon.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	s3Bucket: jsii.String("s3Bucket"),
//   	s3Key: jsii.String("s3Key"),
//   }
//
type CfnApplication_S3LocationProperty struct {
	// The S3 bucket of the S3 object.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The S3 key of the S3 object.
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}

