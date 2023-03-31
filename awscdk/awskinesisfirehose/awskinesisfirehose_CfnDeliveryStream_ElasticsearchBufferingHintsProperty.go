package awskinesisfirehose


// The `ElasticsearchBufferingHints` property type specifies how Amazon Kinesis Data Firehose (Kinesis Data Firehose) buffers incoming data while delivering it to the destination.
//
// The first buffer condition that is satisfied triggers Kinesis Data Firehose to deliver the data.
//
// ElasticsearchBufferingHints is the property type for the `BufferingHints` property of the [Amazon Kinesis Data Firehose DeliveryStream ElasticsearchDestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-elasticsearchdestinationconfiguration.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticsearchBufferingHintsProperty := &elasticsearchBufferingHintsProperty{
//   	intervalInSeconds: jsii.Number(123),
//   	sizeInMBs: jsii.Number(123),
//   }
//
type CfnDeliveryStream_ElasticsearchBufferingHintsProperty struct {
	// The length of time, in seconds, that Kinesis Data Firehose buffers incoming data before delivering it to the destination.
	//
	// For valid values, see the `IntervalInSeconds` content for the [BufferingHints](https://docs.aws.amazon.com/firehose/latest/APIReference/API_BufferingHints.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	IntervalInSeconds *float64 `field:"optional" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// The size of the buffer, in MBs, that Kinesis Data Firehose uses for incoming data before delivering it to the destination.
	//
	// For valid values, see the `SizeInMBs` content for the [BufferingHints](https://docs.aws.amazon.com/firehose/latest/APIReference/API_BufferingHints.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
	SizeInMBs *float64 `field:"optional" json:"sizeInMBs" yaml:"sizeInMBs"`
}

