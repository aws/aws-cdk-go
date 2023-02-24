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
//   brokerLogsProperty := &BrokerLogsProperty{
//   	CloudWatchLogs: &CloudWatchLogsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		LogGroup: jsii.String("logGroup"),
//   	},
//   	Firehose: &FirehoseProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		DeliveryStream: jsii.String("deliveryStream"),
//   	},
//   	S3: &S3Property{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Bucket: jsii.String("bucket"),
//   		Prefix: jsii.String("prefix"),
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

