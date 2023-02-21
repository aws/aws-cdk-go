package awsdynamodb


// The S3 bucket that is being imported from.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3BucketSourceProperty := &s3BucketSourceProperty{
//   	s3Bucket: jsii.String("s3Bucket"),
//
//   	// the properties below are optional
//   	s3BucketOwner: jsii.String("s3BucketOwner"),
//   	s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
type CfnTable_S3BucketSourceProperty struct {
	// The S3 bucket that is being imported from.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The account number of the S3 bucket that is being imported from.
	//
	// If the bucket is owned by the requester this is optional.
	S3BucketOwner *string `field:"optional" json:"s3BucketOwner" yaml:"s3BucketOwner"`
	// The key prefix shared by all S3 Objects that are being imported.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

