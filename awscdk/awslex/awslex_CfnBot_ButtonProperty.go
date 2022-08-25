package awslex


// Describes a button to use on a response card used to gather slot values from a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   buttonProperty := &buttonProperty{
//   	text: jsii.String("text"),
//   	value: jsii.String("value"),
//   }
//
type CfnBot_ButtonProperty struct {
	// The text that appears on the button.
	//
	// Use this to tell the user the value that is returned when they choose this button.
	Text *string `field:"required" json:"text" yaml:"text"`
	// The value returned to Amazon Lex when the user chooses this button.
	//
	// This must be one of the slot values configured for the slot.
	Value *string `field:"required" json:"value" yaml:"value"`
}

