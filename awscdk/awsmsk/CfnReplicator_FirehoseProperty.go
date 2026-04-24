package awsmsk


// Details about delivering logs to Firehose.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseProperty := &FirehoseProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	DeliveryStream: jsii.String("deliveryStream"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-firehose.html
//
type CfnReplicator_FirehoseProperty struct {
	// Whether log delivery to Firehose is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-firehose.html#cfn-msk-replicator-firehose-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The Firehose delivery stream that is the destination for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-firehose.html#cfn-msk-replicator-firehose-deliverystream
	//
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

