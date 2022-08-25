package awsqldb


// The configuration settings of the Amazon Kinesis Data Streams destination for an Amazon QLDB journal stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisConfigurationProperty := &kinesisConfigurationProperty{
//   	aggregationEnabled: jsii.Boolean(false),
//   	streamArn: jsii.String("streamArn"),
//   }
//
type CfnStream_KinesisConfigurationProperty struct {
	// Enables QLDB to publish multiple data records in a single Kinesis Data Streams record, increasing the number of records sent per API call.
	//
	// *This option is enabled by default.* Record aggregation has important implications for processing records and requires de-aggregation in your stream consumer. To learn more, see [KPL Key Concepts](https://docs.aws.amazon.com/streams/latest/dev/kinesis-kpl-concepts.html) and [Consumer De-aggregation](https://docs.aws.amazon.com/streams/latest/dev/kinesis-kpl-consumer-deaggregation.html) in the *Amazon Kinesis Data Streams Developer Guide* .
	AggregationEnabled interface{} `field:"optional" json:"aggregationEnabled" yaml:"aggregationEnabled"`
	// The Amazon Resource Name (ARN) of the Kinesis Data Streams resource.
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
}

