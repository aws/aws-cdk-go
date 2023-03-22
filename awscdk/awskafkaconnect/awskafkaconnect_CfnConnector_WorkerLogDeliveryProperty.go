package awskafkaconnect


// Workers can send worker logs to different destination types.
//
// This configuration specifies the details of these destinations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerLogDeliveryProperty := &workerLogDeliveryProperty{
//   	cloudWatchLogs: &cloudWatchLogsLogDeliveryProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		logGroup: jsii.String("logGroup"),
//   	},
//   	firehose: &firehoseLogDeliveryProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		deliveryStream: jsii.String("deliveryStream"),
//   	},
//   	s3: &s3LogDeliveryProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		bucket: jsii.String("bucket"),
//   		prefix: jsii.String("prefix"),
//   	},
//   }
//
type CfnConnector_WorkerLogDeliveryProperty struct {
	// Details about delivering logs to Amazon CloudWatch Logs.
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Details about delivering logs to Amazon Kinesis Data Firehose.
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Details about delivering logs to Amazon S3.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

