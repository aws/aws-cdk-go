package awsdevopsagent


// Transmission targets for agent notifications.
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
	// Slack channel configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-slacktransmissiontarget.html#cfn-devopsagent-association-slacktransmissiontarget-incidentresponsetarget
	//
	IncidentResponseTarget interface{} `field:"required" json:"incidentResponseTarget" yaml:"incidentResponseTarget"`
}

