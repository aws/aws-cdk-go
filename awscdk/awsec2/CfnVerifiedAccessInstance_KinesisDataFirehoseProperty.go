package awsec2


// Options for Kinesis as a logging destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisDataFirehoseProperty := &KinesisDataFirehoseProperty{
//   	DeliveryStream: jsii.String("deliveryStream"),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-kinesisdatafirehose.html
//
type CfnVerifiedAccessInstance_KinesisDataFirehoseProperty struct {
	// The ID of the delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-kinesisdatafirehose.html#cfn-ec2-verifiedaccessinstance-kinesisdatafirehose-deliverystream
	//
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
	// Indicates whether logging is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-kinesisdatafirehose.html#cfn-ec2-verifiedaccessinstance-kinesisdatafirehose-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

