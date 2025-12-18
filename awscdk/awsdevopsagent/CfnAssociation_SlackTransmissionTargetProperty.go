package awsdevopsagent


// Defines the Slack channels where different types of agent notifications will be sent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slackTransmissionTargetProperty := &SlackTransmissionTargetProperty{
//   	IncidentResponseTarget: &SlackChannelProperty{
//   		ChannelId: jsii.String("channelId"),
//
//   		// the properties below are optional
//   		ChannelName: jsii.String("channelName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slacktransmissiontarget.html
//
type CfnAssociation_SlackTransmissionTargetProperty struct {
	// Destination for AWS DevOps Agent Incident Response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slacktransmissiontarget.html#cfn-devopsagent-association-slacktransmissiontarget-incidentresponsetarget
	//
	IncidentResponseTarget interface{} `field:"required" json:"incidentResponseTarget" yaml:"incidentResponseTarget"`
}

