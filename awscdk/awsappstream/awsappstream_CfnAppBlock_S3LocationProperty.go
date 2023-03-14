package awsappstream


// The S3 location of the app block.
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
type CfnAppBlock_S3LocationProperty struct {
	// The S3 bucket of the app block.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The S3 key of the S3 object of the virtual hard disk.
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}

