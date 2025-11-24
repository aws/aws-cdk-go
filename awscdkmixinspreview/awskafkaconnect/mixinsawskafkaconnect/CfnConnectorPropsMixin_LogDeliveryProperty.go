package mixinsawskafkaconnect


// Details about log delivery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logDeliveryProperty := &LogDeliveryProperty{
//   	WorkerLogDelivery: &WorkerLogDeliveryProperty{
//   		CloudWatchLogs: &CloudWatchLogsLogDeliveryProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		Firehose: &FirehoseLogDeliveryProperty{
//   			DeliveryStream: jsii.String("deliveryStream"),
//   			Enabled: jsii.Boolean(false),
//   		},
//   		S3: &S3LogDeliveryProperty{
//   			Bucket: jsii.String("bucket"),
//   			Enabled: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-logdelivery.html
//
type CfnConnectorPropsMixin_LogDeliveryProperty struct {
	// The workers can send worker logs to different destination types.
	//
	// This configuration specifies the details of these destinations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-logdelivery.html#cfn-kafkaconnect-connector-logdelivery-workerlogdelivery
	//
	WorkerLogDelivery interface{} `field:"optional" json:"workerLogDelivery" yaml:"workerLogDelivery"`
}

