package awskafkaconnect


// The settings for delivering logs to Amazon Kinesis Data Firehose.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseLogDeliveryProperty := &firehoseLogDeliveryProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	deliveryStream: jsii.String("deliveryStream"),
//   }
//
type CfnConnector_FirehoseLogDeliveryProperty struct {
	// Specifies whether connector logs get delivered to Amazon Kinesis Data Firehose.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the Kinesis Data Firehose delivery stream that is the destination for log delivery.
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

