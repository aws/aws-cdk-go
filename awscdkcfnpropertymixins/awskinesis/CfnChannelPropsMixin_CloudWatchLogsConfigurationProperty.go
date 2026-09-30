package awskinesis


// CloudWatch Logs configuration for the channel.
//
// When Enabled is true and LogGroupName is omitted, the service uses a default group derived from the channel name and id; LogStreamName defaults to the literal string 'DestinationDelivery'.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cloudWatchLogsConfigurationProperty := &CloudWatchLogsConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogStreamName: jsii.String("logStreamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-cloudwatchlogsconfiguration.html
//
type CfnChannelPropsMixin_CloudWatchLogsConfigurationProperty struct {
	// Whether CloudWatch Logs delivery is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-cloudwatchlogsconfiguration.html#cfn-kinesis-channel-cloudwatchlogsconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The CloudWatch log group name.
	//
	// When Enabled is true and LogGroupName is omitted, the service uses the default '/aws/kinesis/<channelName>/<channelId>'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-cloudwatchlogsconfiguration.html#cfn-kinesis-channel-cloudwatchlogsconfiguration-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The CloudWatch log stream name.
	//
	// Defaults to the literal string 'DestinationDelivery' when omitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-cloudwatchlogsconfiguration.html#cfn-kinesis-channel-cloudwatchlogsconfiguration-logstreamname
	//
	// Default: - "DestinationDelivery".
	//
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

