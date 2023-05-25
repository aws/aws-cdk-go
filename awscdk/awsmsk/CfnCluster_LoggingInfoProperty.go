package awsmsk


// You can configure your MSK cluster to send broker logs to different destination types.
//
// This is a container for the configuration details related to broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingInfoProperty := &LoggingInfoProperty{
//   	BrokerLogs: &BrokerLogsProperty{
//   		CloudWatchLogs: &CloudWatchLogsProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		Firehose: &FirehoseProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			DeliveryStream: jsii.String("deliveryStream"),
//   		},
//   		S3: &S3Property{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			Bucket: jsii.String("bucket"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
type CfnCluster_LoggingInfoProperty struct {
	// You can configure your MSK cluster to send broker logs to different destination types.
	//
	// This configuration specifies the details of these destinations.
	BrokerLogs interface{} `field:"required" json:"brokerLogs" yaml:"brokerLogs"`
}

