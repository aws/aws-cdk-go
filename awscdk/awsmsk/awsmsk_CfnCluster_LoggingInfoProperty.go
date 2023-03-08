package awsmsk


// You can configure your Amazon MSK cluster to send broker logs to different destination types.
//
// This is a container for the configuration details related to broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingInfoProperty := &loggingInfoProperty{
//   	brokerLogs: &brokerLogsProperty{
//   		cloudWatchLogs: &cloudWatchLogsProperty{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			logGroup: jsii.String("logGroup"),
//   		},
//   		firehose: &firehoseProperty{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			deliveryStream: jsii.String("deliveryStream"),
//   		},
//   		s3: &s3Property{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			bucket: jsii.String("bucket"),
//   			prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
type CfnCluster_LoggingInfoProperty struct {
	// You can configure your Amazon MSK cluster to send broker logs to different destination types.
	//
	// This configuration specifies the details of these destinations.
	BrokerLogs interface{} `field:"required" json:"brokerLogs" yaml:"brokerLogs"`
}

