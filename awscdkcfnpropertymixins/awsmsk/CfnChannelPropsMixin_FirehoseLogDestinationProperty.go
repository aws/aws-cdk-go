package awsmsk


// Firehose log destination details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   firehoseLogDestinationProperty := &FirehoseLogDestinationProperty{
//   	DeliveryStream: jsii.String("deliveryStream"),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-firehoselogdestination.html
//
type CfnChannelPropsMixin_FirehoseLogDestinationProperty struct {
	// The Firehose delivery stream for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-firehoselogdestination.html#cfn-msk-channel-firehoselogdestination-deliverystream
	//
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
	// Whether Firehose logging is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-firehoselogdestination.html#cfn-msk-channel-firehoselogdestination-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

