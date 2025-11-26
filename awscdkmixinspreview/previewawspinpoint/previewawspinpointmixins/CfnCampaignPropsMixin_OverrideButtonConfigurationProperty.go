package previewawspinpointmixins


// Specifies the configuration of a button with settings that are specific to a certain device type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   overrideButtonConfigurationProperty := &OverrideButtonConfigurationProperty{
//   	ButtonAction: jsii.String("buttonAction"),
//   	Link: jsii.String("link"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-overridebuttonconfiguration.html
//
type CfnCampaignPropsMixin_OverrideButtonConfigurationProperty struct {
	// The action that occurs when a recipient chooses a button in an in-app message.
	//
	// You can specify one of the following:
	//
	// - `LINK` – A link to a web destination.
	// - `DEEP_LINK` – A link to a specific page in an application.
	// - `CLOSE` – Dismisses the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-overridebuttonconfiguration.html#cfn-pinpoint-campaign-overridebuttonconfiguration-buttonaction
	//
	ButtonAction *string `field:"optional" json:"buttonAction" yaml:"buttonAction"`
	// The destination (such as a URL) for a button.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-overridebuttonconfiguration.html#cfn-pinpoint-campaign-overridebuttonconfiguration-link
	//
	Link *string `field:"optional" json:"link" yaml:"link"`
}

