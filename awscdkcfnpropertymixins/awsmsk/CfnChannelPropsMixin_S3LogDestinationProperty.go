package awsmsk


// S3 log destination details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3LogDestinationProperty := &S3LogDestinationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Enabled: jsii.Boolean(false),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3logdestination.html
//
type CfnChannelPropsMixin_S3LogDestinationProperty struct {
	// The name of the S3 bucket for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3logdestination.html#cfn-msk-channel-s3logdestination-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Whether S3 logging is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3logdestination.html#cfn-msk-channel-s3logdestination-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The S3 prefix for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3logdestination.html#cfn-msk-channel-s3logdestination-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

