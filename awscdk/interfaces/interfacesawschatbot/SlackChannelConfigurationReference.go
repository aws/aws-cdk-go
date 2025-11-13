package interfacesawschatbot


// A reference to a SlackChannelConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slackChannelConfigurationReference := &SlackChannelConfigurationReference{
//   	SlackChannelConfigurationArn: jsii.String("slackChannelConfigurationArn"),
//   }
//
type SlackChannelConfigurationReference struct {
	// The Arn of the SlackChannelConfiguration resource.
	SlackChannelConfigurationArn *string `field:"required" json:"slackChannelConfigurationArn" yaml:"slackChannelConfigurationArn"`
}

