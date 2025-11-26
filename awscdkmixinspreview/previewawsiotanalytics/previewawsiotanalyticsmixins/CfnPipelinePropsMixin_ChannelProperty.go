package previewawsiotanalyticsmixins


// Determines the source of the messages to be processed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   channelProperty := &ChannelProperty{
//   	ChannelName: jsii.String("channelName"),
//   	Name: jsii.String("name"),
//   	Next: jsii.String("next"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-channel.html
//
type CfnPipelinePropsMixin_ChannelProperty struct {
	// The name of the channel from which the messages are processed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-channel.html#cfn-iotanalytics-pipeline-channel-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The name of the 'channel' activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-channel.html#cfn-iotanalytics-pipeline-channel-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-channel.html#cfn-iotanalytics-pipeline-channel-next
	//
	Next *string `field:"optional" json:"next" yaml:"next"`
}

