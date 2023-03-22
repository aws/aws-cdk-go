package awsivschat


// The FirehoseDestinationConfiguration property type specifies a Kinesis Firehose location where chat logs will be stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseDestinationConfigurationProperty := &firehoseDestinationConfigurationProperty{
//   	deliveryStreamName: jsii.String("deliveryStreamName"),
//   }
//
type CfnLoggingConfiguration_FirehoseDestinationConfigurationProperty struct {
	// Name of the Amazon Kinesis Firehose delivery stream where chat activity will be logged.
	DeliveryStreamName *string `field:"required" json:"deliveryStreamName" yaml:"deliveryStreamName"`
}

