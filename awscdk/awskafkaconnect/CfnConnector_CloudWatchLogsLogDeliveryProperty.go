package awskafkaconnect


// The settings for delivering connector logs to Amazon CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsLogDeliveryProperty := &CloudWatchLogsLogDeliveryProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	LogGroup: jsii.String("logGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-cloudwatchlogslogdelivery.html
//
type CfnConnector_CloudWatchLogsLogDeliveryProperty struct {
	// Whether log delivery to Amazon CloudWatch Logs is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-cloudwatchlogslogdelivery.html#cfn-kafkaconnect-connector-cloudwatchlogslogdelivery-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the CloudWatch log group that is the destination for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-connector-cloudwatchlogslogdelivery.html#cfn-kafkaconnect-connector-cloudwatchlogslogdelivery-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

