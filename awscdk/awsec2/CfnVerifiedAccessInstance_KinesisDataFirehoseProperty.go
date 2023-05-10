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
type CfnVerifiedAccessInstance_KinesisDataFirehoseProperty struct {
	// The ID of the delivery stream.
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
	// Indicates whether logging is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

