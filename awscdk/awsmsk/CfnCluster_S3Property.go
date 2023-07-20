package awsmsk


// The details of the Amazon S3 destination for broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3Property := &S3Property{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	Bucket: jsii.String("bucket"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-s3.html
//
type CfnCluster_S3Property struct {
	// Specifies whether broker logs get sent to the specified Amazon S3 destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-s3.html#cfn-msk-cluster-s3-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the S3 bucket that is the destination for broker logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-s3.html#cfn-msk-cluster-s3-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix that is the destination for broker logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-s3.html#cfn-msk-cluster-s3-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

