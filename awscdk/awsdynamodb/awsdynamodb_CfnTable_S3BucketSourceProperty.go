package awsdynamodb


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
	// `CfnTable.S3BucketSourceProperty.S3Bucket`.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// `CfnTable.S3BucketSourceProperty.S3BucketOwner`.
	S3BucketOwner *string `field:"optional" json:"s3BucketOwner" yaml:"s3BucketOwner"`
	// `CfnTable.S3BucketSourceProperty.S3KeyPrefix`.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

