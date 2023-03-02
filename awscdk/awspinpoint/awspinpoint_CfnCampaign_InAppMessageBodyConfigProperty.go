package awspinpoint


// Specifies the configuration of main body text of the in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppMessageBodyConfigProperty := &inAppMessageBodyConfigProperty{
//   	alignment: jsii.String("alignment"),
//   	body: jsii.String("body"),
//   	textColor: jsii.String("textColor"),
//   }
//
type CfnCampaign_InAppMessageBodyConfigProperty struct {
	// The text alignment of the main body text of the message.
	//
	// Acceptable values: `LEFT` , `CENTER` , `RIGHT` .
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// The main body text of the message.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The color of the body text, expressed as a string consisting of a hex color code (such as "#000000" for black).
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

