package awscognito


// The configuration of user event logs to an external AWS service like Amazon Data Firehose, Amazon S3, or Amazon CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	CloudWatchLogsConfiguration: &CloudWatchLogsConfigurationProperty{
//   		LogGroupArn: jsii.String("logGroupArn"),
//   	},
//   	EventSource: jsii.String("eventSource"),
//   	FirehoseConfiguration: &FirehoseConfigurationProperty{
//   		StreamArn: jsii.String("streamArn"),
//   	},
//   	LogLevel: jsii.String("logLevel"),
//   	S3Configuration: &S3ConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-logdeliveryconfiguration-logconfiguration.html
//
type CfnLogDeliveryConfiguration_LogConfigurationProperty struct {
	// Configuration for the CloudWatch log group destination of user pool detailed activity logging, or of user activity log export with advanced security features.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-logdeliveryconfiguration-logconfiguration.html#cfn-cognito-logdeliveryconfiguration-logconfiguration-cloudwatchlogsconfiguration
	//
	CloudWatchLogsConfiguration interface{} `field:"optional" json:"cloudWatchLogsConfiguration" yaml:"cloudWatchLogsConfiguration"`
	// The source of events that your user pool sends for logging.
	//
	// To send error-level logs about user notification activity, set to `userNotification` . To send info-level logs about threat-protection user activity in user pools with the Plus feature plan, set to `userAuthEvents` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-logdeliveryconfiguration-logconfiguration.html#cfn-cognito-logdeliveryconfiguration-logconfiguration-eventsource
	//
	EventSource *string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// Configuration for the Amazon Data Firehose stream destination of user activity log export with threat protection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-logdeliveryconfiguration-logconfiguration.html#cfn-cognito-logdeliveryconfiguration-logconfiguration-firehoseconfiguration
	//
	FirehoseConfiguration interface{} `field:"optional" json:"firehoseConfiguration" yaml:"firehoseConfiguration"`
	// The `errorlevel` selection of logs that a user pool sends for detailed activity logging.
	//
	// To send `userNotification` activity with [information about message delivery](https://docs.aws.amazon.com/cognito/latest/developerguide/exporting-quotas-and-usage.html) , choose `ERROR` with `CloudWatchLogsConfiguration` . To send `userAuthEvents` activity with user logs from threat protection with the Plus feature plan, choose `INFO` with one of `CloudWatchLogsConfiguration` , `FirehoseConfiguration` , or `S3Configuration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-logdeliveryconfiguration-logconfiguration.html#cfn-cognito-logdeliveryconfiguration-logconfiguration-loglevel
	//
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Configuration for the Amazon S3 bucket destination of user activity log export with threat protection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-logdeliveryconfiguration-logconfiguration.html#cfn-cognito-logdeliveryconfiguration-logconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

