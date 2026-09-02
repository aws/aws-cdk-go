package awsmsk


// S3 log destination details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LogDestinationProperty := &S3LogDestinationProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	Bucket: jsii.String("bucket"),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3logdestination.html
//
type CfnChannel_S3LogDestinationProperty struct {
	// Whether S3 logging is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3logdestination.html#cfn-msk-channel-s3logdestination-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the S3 bucket for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3logdestination.html#cfn-msk-channel-s3logdestination-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3logdestination.html#cfn-msk-channel-s3logdestination-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

