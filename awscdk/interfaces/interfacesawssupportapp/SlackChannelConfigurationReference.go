package interfacesawssupportapp


// A reference to a SlackChannelConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slackChannelConfigurationReference := &SlackChannelConfigurationReference{
//   	ChannelId: jsii.String("channelId"),
//   	TeamId: jsii.String("teamId"),
//   }
//
type SlackChannelConfigurationReference struct {
	// The ChannelId of the SlackChannelConfiguration resource.
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// The TeamId of the SlackChannelConfiguration resource.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
}

