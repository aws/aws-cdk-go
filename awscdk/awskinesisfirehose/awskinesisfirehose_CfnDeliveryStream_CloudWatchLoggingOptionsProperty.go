package awskinesisfirehose


// The `CloudWatchLoggingOptions` property type specifies Amazon CloudWatch Logs (CloudWatch Logs) logging options that Amazon Kinesis Data Firehose (Kinesis Data Firehose) uses for the delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLoggingOptionsProperty := &cloudWatchLoggingOptionsProperty{
//   	enabled: jsii.Boolean(false),
//   	logGroupName: jsii.String("logGroupName"),
//   	logStreamName: jsii.String("logStreamName"),
//   }
//
type CfnDeliveryStream_CloudWatchLoggingOptionsProperty struct {
	// Indicates whether CloudWatch Logs logging is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The name of the CloudWatch Logs log group that contains the log stream that Kinesis Data Firehose will use.
	//
	// Conditional. If you enable logging, you must specify this property.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The name of the CloudWatch Logs log stream that Kinesis Data Firehose uses to send logs about data delivery.
	//
	// Conditional. If you enable logging, you must specify this property.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

