package awssupportapp


// Properties for defining a `CfnSlackChannelConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSlackChannelConfigurationProps := &cfnSlackChannelConfigurationProps{
//   	channelId: jsii.String("channelId"),
//   	channelRoleArn: jsii.String("channelRoleArn"),
//   	notifyOnCaseSeverity: jsii.String("notifyOnCaseSeverity"),
//   	teamId: jsii.String("teamId"),
//
//   	// the properties below are optional
//   	channelName: jsii.String("channelName"),
//   	notifyOnAddCorrespondenceToCase: jsii.Boolean(false),
//   	notifyOnCreateOrReopenCase: jsii.Boolean(false),
//   	notifyOnResolveCase: jsii.Boolean(false),
//   }
//
type CfnSlackChannelConfigurationProps struct {
	// `AWS::SupportApp::SlackChannelConfiguration.ChannelId`.
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// `AWS::SupportApp::SlackChannelConfiguration.ChannelRoleArn`.
	ChannelRoleArn *string `field:"required" json:"channelRoleArn" yaml:"channelRoleArn"`
	// `AWS::SupportApp::SlackChannelConfiguration.NotifyOnCaseSeverity`.
	NotifyOnCaseSeverity *string `field:"required" json:"notifyOnCaseSeverity" yaml:"notifyOnCaseSeverity"`
	// `AWS::SupportApp::SlackChannelConfiguration.TeamId`.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// `AWS::SupportApp::SlackChannelConfiguration.ChannelName`.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// `AWS::SupportApp::SlackChannelConfiguration.NotifyOnAddCorrespondenceToCase`.
	NotifyOnAddCorrespondenceToCase interface{} `field:"optional" json:"notifyOnAddCorrespondenceToCase" yaml:"notifyOnAddCorrespondenceToCase"`
	// `AWS::SupportApp::SlackChannelConfiguration.NotifyOnCreateOrReopenCase`.
	NotifyOnCreateOrReopenCase interface{} `field:"optional" json:"notifyOnCreateOrReopenCase" yaml:"notifyOnCreateOrReopenCase"`
	// `AWS::SupportApp::SlackChannelConfiguration.NotifyOnResolveCase`.
	NotifyOnResolveCase interface{} `field:"optional" json:"notifyOnResolveCase" yaml:"notifyOnResolveCase"`
}

