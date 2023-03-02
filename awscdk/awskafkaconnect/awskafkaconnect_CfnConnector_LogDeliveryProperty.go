package awskafkaconnect


// Details about log delivery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDeliveryProperty := &logDeliveryProperty{
//   	workerLogDelivery: &workerLogDeliveryProperty{
//   		cloudWatchLogs: &cloudWatchLogsLogDeliveryProperty{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			logGroup: jsii.String("logGroup"),
//   		},
//   		firehose: &firehoseLogDeliveryProperty{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			deliveryStream: jsii.String("deliveryStream"),
//   		},
//   		s3: &s3LogDeliveryProperty{
//   			enabled: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			bucket: jsii.String("bucket"),
//   			prefix: jsii.String("prefix"),
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

