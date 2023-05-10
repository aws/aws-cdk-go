package awsec2


// Options for Amazon S3 as a logging destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3Property := &S3Property{
//   	BucketName: jsii.String("bucketName"),
//   	BucketOwner: jsii.String("bucketOwner"),
//   	Enabled: jsii.Boolean(false),
//   	Prefix: jsii.String("prefix"),
//   }
//
type CfnVerifiedAccessInstance_S3Property struct {
	// The bucket name.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The AWS account number that owns the bucket.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// Indicates whether logging is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The bucket prefix.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

