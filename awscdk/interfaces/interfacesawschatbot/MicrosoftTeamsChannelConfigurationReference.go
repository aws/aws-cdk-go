package interfacesawschatbot


// A reference to a MicrosoftTeamsChannelConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   microsoftTeamsChannelConfigurationReference := &MicrosoftTeamsChannelConfigurationReference{
//   	MicrosoftTeamsChannelConfigurationArn: jsii.String("microsoftTeamsChannelConfigurationArn"),
//   }
//
type MicrosoftTeamsChannelConfigurationReference struct {
	// The Arn of the MicrosoftTeamsChannelConfiguration resource.
	MicrosoftTeamsChannelConfigurationArn *string `field:"required" json:"microsoftTeamsChannelConfigurationArn" yaml:"microsoftTeamsChannelConfigurationArn"`
}

