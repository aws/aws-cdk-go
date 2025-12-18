package previewawsdevopsagentmixins


// Defines the Slack channels where different types of agent notifications will be sent.
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
	// Destination for AWS DevOps Agent Incident Response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slacktransmissiontarget.html#cfn-devopsagent-association-slacktransmissiontarget-incidentresponsetarget
	//
	IncidentResponseTarget interface{} `field:"optional" json:"incidentResponseTarget" yaml:"incidentResponseTarget"`
}

