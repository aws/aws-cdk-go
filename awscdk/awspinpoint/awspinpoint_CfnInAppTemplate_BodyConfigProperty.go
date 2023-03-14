package awspinpoint


// Specifies the configuration of the main body text of the in-app message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bodyConfigProperty := &bodyConfigProperty{
//   	alignment: jsii.String("alignment"),
//   	body: jsii.String("body"),
//   	textColor: jsii.String("textColor"),
//   }
//
type CfnInAppTemplate_BodyConfigProperty struct {
	// The text alignment of the main body text of the message.
	//
	// Acceptable values: `LEFT` , `CENTER` , `RIGHT` .
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// The main body text of the message.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The color of the body text, expressed as a hex color code (such as #000000 for black).
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

