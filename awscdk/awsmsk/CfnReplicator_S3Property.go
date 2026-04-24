package awsmsk


// Details about delivering logs to S3.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-s3.html
//
type CfnReplicator_S3Property struct {
	// Whether log delivery to S3 is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-s3.html#cfn-msk-replicator-s3-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The S3 bucket that is the destination for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-s3.html#cfn-msk-replicator-s3-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix that is the destination for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-s3.html#cfn-msk-replicator-s3-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

