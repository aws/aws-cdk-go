package awspinpoint


// Specifies the configuration of a button with settings that are specific to a certain device type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   overrideButtonConfigurationProperty := &overrideButtonConfigurationProperty{
//   	buttonAction: jsii.String("buttonAction"),
//   	link: jsii.String("link"),
//   }
//
type CfnCampaign_OverrideButtonConfigurationProperty struct {
	// The action that occurs when a recipient chooses a button in an in-app message.
	//
	// You can specify one of the following:
	//
	// - `LINK` – A link to a web destination.
	// - `DEEP_LINK` – A link to a specific page in an application.
	// - `CLOSE` – Dismisses the message.
	ButtonAction *string `field:"optional" json:"buttonAction" yaml:"buttonAction"`
	// The destination (such as a URL) for a button.
	Link *string `field:"optional" json:"link" yaml:"link"`
}

