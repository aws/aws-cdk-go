package awssupportapp


// Properties for defining a `CfnSlackChannelConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSlackChannelConfigurationProps := &CfnSlackChannelConfigurationProps{
//   	ChannelId: jsii.String("channelId"),
//   	ChannelRoleArn: jsii.String("channelRoleArn"),
//   	NotifyOnCaseSeverity: jsii.String("notifyOnCaseSeverity"),
//   	TeamId: jsii.String("teamId"),
//
//   	// the properties below are optional
//   	ChannelName: jsii.String("channelName"),
//   	NotifyOnAddCorrespondenceToCase: jsii.Boolean(false),
//   	NotifyOnCreateOrReopenCase: jsii.Boolean(false),
//   	NotifyOnResolveCase: jsii.Boolean(false),
//   }
//
type CfnSlackChannelConfigurationProps struct {
	// The channel ID in Slack.
	//
	// This ID identifies a channel within a Slack workspace.
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// The Amazon Resource Name (ARN) of the IAM role for this Slack channel configuration.
	//
	// The AWS Support App uses this role to perform AWS Support and Service Quotas actions on your behalf.
	ChannelRoleArn *string `field:"required" json:"channelRoleArn" yaml:"channelRoleArn"`
	// The case severity for your support cases that you want to receive notifications.
	//
	// You can specify `none` , `all` , or `high` .
	NotifyOnCaseSeverity *string `field:"required" json:"notifyOnCaseSeverity" yaml:"notifyOnCaseSeverity"`
	// The team ID in Slack.
	//
	// This ID uniquely identifies a Slack workspace.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// The channel name in Slack.
	//
	// This is the channel where you invite the AWS Support App .
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Whether to get notified when a correspondence is added to your support cases.
	NotifyOnAddCorrespondenceToCase interface{} `field:"optional" json:"notifyOnAddCorrespondenceToCase" yaml:"notifyOnAddCorrespondenceToCase"`
	// Whether to get notified when your support cases are created or reopened.
	NotifyOnCreateOrReopenCase interface{} `field:"optional" json:"notifyOnCreateOrReopenCase" yaml:"notifyOnCreateOrReopenCase"`
	// Whether to get notified when your support cases are resolved.
	NotifyOnResolveCase interface{} `field:"optional" json:"notifyOnResolveCase" yaml:"notifyOnResolveCase"`
}

