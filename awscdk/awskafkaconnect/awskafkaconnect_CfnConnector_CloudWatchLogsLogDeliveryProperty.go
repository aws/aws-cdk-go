package awskafkaconnect


// The settings for delivering connector logs to Amazon CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsLogDeliveryProperty := &cloudWatchLogsLogDeliveryProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	logGroup: jsii.String("logGroup"),
//   }
//
type CfnConnector_CloudWatchLogsLogDeliveryProperty struct {
	// Whether log delivery to Amazon CloudWatch Logs is enabled.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the CloudWatch log group that is the destination for log delivery.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

