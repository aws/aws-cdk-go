package awspinpoint


// Specifies the configuration and content of the header or title text of the in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppMessageHeaderConfigProperty := &inAppMessageHeaderConfigProperty{
//   	alignment: jsii.String("alignment"),
//   	header: jsii.String("header"),
//   	textColor: jsii.String("textColor"),
//   }
//
type CfnCampaign_InAppMessageHeaderConfigProperty struct {
	// The text alignment of the title of the message.
	//
	// Acceptable values: `LEFT` , `CENTER` , `RIGHT` .
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// The header or title text of the in-app message.
	Header *string `field:"optional" json:"header" yaml:"header"`
	// The color of the body text, expressed as a string consisting of a hex color code (such as "#000000" for black).
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

