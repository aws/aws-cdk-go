package awsmsk


// Firehose log destination details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseLogDestinationProperty := &FirehoseLogDestinationProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	DeliveryStream: jsii.String("deliveryStream"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-firehoselogdestination.html
//
type CfnChannel_FirehoseLogDestinationProperty struct {
	// Whether Firehose logging is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-firehoselogdestination.html#cfn-msk-channel-firehoselogdestination-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The Firehose delivery stream for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-firehoselogdestination.html#cfn-msk-channel-firehoselogdestination-deliverystream
	//
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

