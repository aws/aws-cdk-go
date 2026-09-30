package awskinesis


// Configuration for delivering channel operational logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   loggingConfigurationProperty := &LoggingConfigurationProperty{
//   	CloudWatchLogs: &CloudWatchLogsConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogStreamName: jsii.String("logStreamName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-loggingconfiguration.html
//
type CfnChannelPropsMixin_LoggingConfigurationProperty struct {
	// CloudWatch Logs configuration for the channel.
	//
	// When Enabled is true and LogGroupName is omitted, the service uses a default group derived from the channel name and id; LogStreamName defaults to the literal string 'DestinationDelivery'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-loggingconfiguration.html#cfn-kinesis-channel-loggingconfiguration-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
}

