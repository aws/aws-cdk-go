package awskafkaconnect


// Details about log delivery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDeliveryProperty := &LogDeliveryProperty{
//   	WorkerLogDelivery: &WorkerLogDeliveryProperty{
//   		CloudWatchLogs: &CloudWatchLogsLogDeliveryProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		Firehose: &FirehoseLogDeliveryProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			DeliveryStream: jsii.String("deliveryStream"),
//   		},
//   		S3: &S3LogDeliveryProperty{
//   			Enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			Bucket: jsii.String("bucket"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
type CfnConnector_LogDeliveryProperty struct {
	// The workers can send worker logs to different destination types.
	//
	// This configuration specifies the details of these destinations.
	WorkerLogDelivery interface{} `field:"required" json:"workerLogDelivery" yaml:"workerLogDelivery"`
}

