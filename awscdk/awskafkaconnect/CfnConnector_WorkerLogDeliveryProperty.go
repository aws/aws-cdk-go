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
//   workerLogDeliveryProperty := &WorkerLogDeliveryProperty{
//   	CloudWatchLogs: &CloudWatchLogsLogDeliveryProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		LogGroup: jsii.String("logGroup"),
//   	},
//   	Firehose: &FirehoseLogDeliveryProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		DeliveryStream: jsii.String("deliveryStream"),
//   	},
//   	S3: &S3LogDeliveryProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Bucket: jsii.String("bucket"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-workerlogdelivery.html
//
type CfnConnector_WorkerLogDeliveryProperty struct {
	// Details about delivering logs to Amazon CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-workerlogdelivery.html#cfn-kafkaconnect-connector-workerlogdelivery-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Details about delivering logs to Amazon Kinesis Data Firehose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-workerlogdelivery.html#cfn-kafkaconnect-connector-workerlogdelivery-firehose
	//
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// Details about delivering logs to Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-workerlogdelivery.html#cfn-kafkaconnect-connector-workerlogdelivery-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

