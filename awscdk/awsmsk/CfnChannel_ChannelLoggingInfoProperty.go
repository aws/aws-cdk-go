package awsmsk


// Log configuration details for Channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelLoggingInfoProperty := &ChannelLoggingInfoProperty{
//   	CloudWatchLogs: &CloudWatchLogsLogDestinationProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		LogGroup: jsii.String("logGroup"),
//   	},
//   	Firehose: &FirehoseLogDestinationProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		DeliveryStream: jsii.String("deliveryStream"),
//   	},
//   	S3: &S3LogDestinationProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Bucket: jsii.String("bucket"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-channellogginginfo.html
//
type CfnChannel_ChannelLoggingInfoProperty struct {
	// CloudWatch Logs log destination details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-channellogginginfo.html#cfn-msk-channel-channellogginginfo-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Firehose log destination details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-channellogginginfo.html#cfn-msk-channel-channellogginginfo-firehose
	//
	Firehose interface{} `field:"optional" json:"firehose" yaml:"firehose"`
	// S3 log destination details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-channellogginginfo.html#cfn-msk-channel-channellogginginfo-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

