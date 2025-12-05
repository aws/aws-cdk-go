package previewawsdevopsagentmixins


// Slack workspace integration configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   slackConfigurationProperty := &SlackConfigurationProperty{
//   	TransmissionTarget: &SlackTransmissionTargetProperty{
//   		IncidentResponseTarget: &SlackChannelProperty{
//   			ChannelId: jsii.String("channelId"),
//   			ChannelName: jsii.String("channelName"),
//   		},
//   	},
//   	WorkspaceId: jsii.String("workspaceId"),
//   	WorkspaceName: jsii.String("workspaceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackconfiguration.html
//
type CfnAssociationPropsMixin_SlackConfigurationProperty struct {
	// Transmission targets for agent notifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackconfiguration.html#cfn-devopsagent-association-slackconfiguration-transmissiontarget
	//
	TransmissionTarget interface{} `field:"optional" json:"transmissionTarget" yaml:"transmissionTarget"`
	// Associated Slack workspace ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackconfiguration.html#cfn-devopsagent-association-slackconfiguration-workspaceid
	//
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
	// Associated Slack workspace name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slackconfiguration.html#cfn-devopsagent-association-slackconfiguration-workspacename
	//
	WorkspaceName *string `field:"optional" json:"workspaceName" yaml:"workspaceName"`
}

