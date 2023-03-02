package awspinpoint


// Specifies the default behavior of a button that appears in an in-app message.
//
// You can optionally add button configurations that specifically apply to iOS, Android, or web browser users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultButtonConfigurationProperty := &defaultButtonConfigurationProperty{
//   	backgroundColor: jsii.String("backgroundColor"),
//   	borderRadius: jsii.Number(123),
//   	buttonAction: jsii.String("buttonAction"),
//   	link: jsii.String("link"),
//   	text: jsii.String("text"),
//   	textColor: jsii.String("textColor"),
//   }
//
type CfnInAppTemplate_DefaultButtonConfigurationProperty struct {
	// The background color of a button, expressed as a hex color code (such as #000000 for black).
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The border radius of a button.
	BorderRadius *float64 `field:"optional" json:"borderRadius" yaml:"borderRadius"`
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
	// The text that appears on a button in an in-app message.
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The color of the body text in a button, expressed as a hex color code (such as #000000 for black).
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

