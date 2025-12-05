package awsdevopsagent


// Slack channel configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slackChannelProperty := &SlackChannelProperty{
//   	ChannelId: jsii.String("channelId"),
//
//   	// the properties below are optional
//   	ChannelName: jsii.String("channelName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackchannel.html
//
type CfnAssociation_SlackChannelProperty struct {
	// Slack channel ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackchannel.html#cfn-devopsagent-association-slackchannel-channelid
	//
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// Slack channel name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackchannel.html#cfn-devopsagent-association-slackchannel-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
}

