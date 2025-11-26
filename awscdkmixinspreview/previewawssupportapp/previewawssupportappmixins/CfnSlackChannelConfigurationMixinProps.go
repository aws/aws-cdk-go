package previewawssupportappmixins


// Properties for CfnSlackChannelConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSlackChannelConfigurationMixinProps := &CfnSlackChannelConfigurationMixinProps{
//   	ChannelId: jsii.String("channelId"),
//   	ChannelName: jsii.String("channelName"),
//   	ChannelRoleArn: jsii.String("channelRoleArn"),
//   	NotifyOnAddCorrespondenceToCase: jsii.Boolean(false),
//   	NotifyOnCaseSeverity: jsii.String("notifyOnCaseSeverity"),
//   	NotifyOnCreateOrReopenCase: jsii.Boolean(false),
//   	NotifyOnResolveCase: jsii.Boolean(false),
//   	TeamId: jsii.String("teamId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackchannelconfiguration.html
//
type CfnSlackChannelConfigurationMixinProps struct {
	// The channel ID in Slack.
	//
	// This ID identifies a channel within a Slack workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackchannelconfiguration.html#cfn-supportapp-slackchannelconfiguration-channelid
	//
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// The channel name in Slack.
	//
	// This is the channel where you invite the AWS Support App .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackchannelconfiguration.html#cfn-supportapp-slackchannelconfiguration-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The Amazon Resource Name (ARN) of the IAM role for this Slack channel configuration.
	//
	// The  App uses this role to perform  and Service Quotas actions on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackchannelconfiguration.html#cfn-supportapp-slackchannelconfiguration-channelrolearn
	//
	ChannelRoleArn *string `field:"optional" json:"channelRoleArn" yaml:"channelRoleArn"`
	// Whether to get notified when a correspondence is added to your support cases.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackchannelconfiguration.html#cfn-supportapp-slackchannelconfiguration-notifyonaddcorrespondencetocase
	//
	NotifyOnAddCorrespondenceToCase interface{} `field:"optional" json:"notifyOnAddCorrespondenceToCase" yaml:"notifyOnAddCorrespondenceToCase"`
	// The case severity for your support cases that you want to receive notifications.
	//
	// You can specify `none` , `all` , or `high` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackchannelconfiguration.html#cfn-supportapp-slackchannelconfiguration-notifyoncaseseverity
	//
	NotifyOnCaseSeverity *string `field:"optional" json:"notifyOnCaseSeverity" yaml:"notifyOnCaseSeverity"`
	// Whether to get notified when your support cases are created or reopened.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackchannelconfiguration.html#cfn-supportapp-slackchannelconfiguration-notifyoncreateorreopencase
	//
	NotifyOnCreateOrReopenCase interface{} `field:"optional" json:"notifyOnCreateOrReopenCase" yaml:"notifyOnCreateOrReopenCase"`
	// Whether to get notified when your support cases are resolved.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackchannelconfiguration.html#cfn-supportapp-slackchannelconfiguration-notifyonresolvecase
	//
	NotifyOnResolveCase interface{} `field:"optional" json:"notifyOnResolveCase" yaml:"notifyOnResolveCase"`
	// The team ID in Slack.
	//
	// This ID uniquely identifies a Slack workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-supportapp-slackchannelconfiguration.html#cfn-supportapp-slackchannelconfiguration-teamid
	//
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
}

