package awsmsk


// Details of the Kinesis Data Firehose delivery stream that is the destination for broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseProperty := &firehoseProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	deliveryStream: jsii.String("deliveryStream"),
//   }
//
type CfnCluster_FirehoseProperty struct {
	// Specifies whether broker logs get sent to the specified Kinesis Data Firehose delivery stream.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The Kinesis Data Firehose delivery stream that is the destination for broker logs.
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

