package awsmsk


// The details of the Amazon S3 destination for broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3Property := &s3Property{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	bucket: jsii.String("bucket"),
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnCluster_S3Property struct {
	// Specifies whether broker logs get sent to the specified Amazon S3 destination.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the S3 bucket that is the destination for broker logs.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix that is the destination for broker logs.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

