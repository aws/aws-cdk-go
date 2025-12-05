package previewawsdevopsagentmixins


// Transmission targets for agent notifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   slackTransmissionTargetProperty := &SlackTransmissionTargetProperty{
//   	IncidentResponseTarget: &SlackChannelProperty{
//   		ChannelId: jsii.String("channelId"),
//   		ChannelName: jsii.String("channelName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slacktransmissiontarget.html
//
type CfnAssociationPropsMixin_SlackTransmissionTargetProperty struct {
	// Slack channel configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slacktransmissiontarget.html#cfn-devopsagent-association-slacktransmissiontarget-incidentresponsetarget
	//
	IncidentResponseTarget interface{} `field:"optional" json:"incidentResponseTarget" yaml:"incidentResponseTarget"`
}

