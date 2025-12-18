package awsdevopsagent


// Configuration for Slack workspace integration.
//
// Defines the workspace ID, workspace name, and transmission targets that specify which Slack channels receive notifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slackConfigurationProperty := &SlackConfigurationProperty{
//   	TransmissionTarget: &SlackTransmissionTargetProperty{
//   		IncidentResponseTarget: &SlackChannelProperty{
//   			ChannelId: jsii.String("channelId"),
//
//   			// the properties below are optional
//   			ChannelName: jsii.String("channelName"),
//   		},
//   	},
//   	WorkspaceId: jsii.String("workspaceId"),
//   	WorkspaceName: jsii.String("workspaceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackconfiguration.html
//
type CfnAssociation_SlackConfigurationProperty struct {
	// Transmission targets for agent notifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackconfiguration.html#cfn-devopsagent-association-slackconfiguration-transmissiontarget
	//
	TransmissionTarget interface{} `field:"required" json:"transmissionTarget" yaml:"transmissionTarget"`
	// Associated Slack workspace ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackconfiguration.html#cfn-devopsagent-association-slackconfiguration-workspaceid
	//
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Associated Slack workspace name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackconfiguration.html#cfn-devopsagent-association-slackconfiguration-workspacename
	//
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
}

