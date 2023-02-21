package awselasticache


// Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationDetailsProperty := &destinationDetailsProperty{
//   	cloudWatchLogsDetails: &cloudWatchLogsDestinationDetailsProperty{
//   		logGroup: jsii.String("logGroup"),
//   	},
//   	kinesisFirehoseDetails: &kinesisFirehoseDestinationDetailsProperty{
//   		deliveryStream: jsii.String("deliveryStream"),
//   	},
//   }
//
type CfnReplicationGroup_DestinationDetailsProperty struct {
	// The configuration details of the CloudWatch Logs destination.
	//
	// Note that this field is marked as required but only if CloudWatch Logs was chosen as the destination.
	CloudWatchLogsDetails interface{} `field:"optional" json:"cloudWatchLogsDetails" yaml:"cloudWatchLogsDetails"`
	// The configuration details of the Kinesis Data Firehose destination.
	//
	// Note that this field is marked as required but only if Kinesis Data Firehose was chosen as the destination.
	KinesisFirehoseDetails interface{} `field:"optional" json:"kinesisFirehoseDetails" yaml:"kinesisFirehoseDetails"`
}

