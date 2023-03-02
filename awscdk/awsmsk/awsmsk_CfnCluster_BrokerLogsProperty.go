package awsmsk


// You can configure your Amazon MSK cluster to send broker logs to different destination types.
//
// This configuration specifies the details of these destinations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   brokerLogsProperty := &brokerLogsProperty{
//   	cloudWatchLogs: &cloudWatchLogsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		logGroup: jsii.String("logGroup"),
//   	},
//   	firehose: &firehoseProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		deliveryStream: jsii.String("deliveryStream"),
//   	},
//   	s3: &s3Property{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		bucket: jsii.String("bucket"),
//   		prefix: jsii.String("prefix"),
//   	},
//   }
//
type CfnCluster_BrokerLogsProperty struct {
	// Details of the CloudWatch Logs destination for broker logs.
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Details of the Kinesis Data Firehose delivery stream that is the destination for broker logs.
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Details of the Amazon MSK destination for broker logs.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

