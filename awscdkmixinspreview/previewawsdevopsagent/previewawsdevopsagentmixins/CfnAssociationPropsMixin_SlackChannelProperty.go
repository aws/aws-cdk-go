package previewawsdevopsagentmixins


// Slack channel configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   slackChannelProperty := &SlackChannelProperty{
//   	ChannelId: jsii.String("channelId"),
//   	ChannelName: jsii.String("channelName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackchannel.html
//
type CfnAssociationPropsMixin_SlackChannelProperty struct {
	// Slack channel ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackchannel.html#cfn-devopsagent-association-slackchannel-channelid
	//
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// Slack channel name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackchannel.html#cfn-devopsagent-association-slackchannel-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
}

