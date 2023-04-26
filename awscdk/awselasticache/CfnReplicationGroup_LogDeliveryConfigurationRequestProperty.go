package awselasticache


// Specifies the destination, format and type of the logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDeliveryConfigurationRequestProperty := &LogDeliveryConfigurationRequestProperty{
//   	DestinationDetails: &DestinationDetailsProperty{
//   		CloudWatchLogsDetails: &CloudWatchLogsDestinationDetailsProperty{
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		KinesisFirehoseDetails: &KinesisFirehoseDestinationDetailsProperty{
//   			DeliveryStream: jsii.String("deliveryStream"),
//   		},
//   	},
//   	DestinationType: jsii.String("destinationType"),
//   	LogFormat: jsii.String("logFormat"),
//   	LogType: jsii.String("logType"),
//   }
//
type CfnReplicationGroup_LogDeliveryConfigurationRequestProperty struct {
	// Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.
	DestinationDetails interface{} `field:"required" json:"destinationDetails" yaml:"destinationDetails"`
	// Specify either CloudWatch Logs or Kinesis Data Firehose as the destination type.
	//
	// Valid values are either `cloudwatch-logs` or `kinesis-firehose` .
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
	// Valid values are either `json` or `text` .
	LogFormat *string `field:"required" json:"logFormat" yaml:"logFormat"`
	// Valid value is either `slow-log` , which refers to [slow-log](https://docs.aws.amazon.com/https://redis.io/commands/slowlog) or `engine-log` .
	LogType *string `field:"required" json:"logType" yaml:"logType"`
}

