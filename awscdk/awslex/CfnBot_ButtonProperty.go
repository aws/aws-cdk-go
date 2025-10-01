package awslex


// Describes a button to use on a response card used to gather slot values from a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   buttonProperty := &ButtonProperty{
//   	Text: jsii.String("text"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-button.html
//
type CfnBot_ButtonProperty struct {
	// The text that appears on the button.
	//
	// Use this to tell the user what value is returned when they choose this button.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-button.html#cfn-lex-bot-button-text
	//
	Text *string `field:"required" json:"text" yaml:"text"`
	// The value returned to Amazon Lex when the user chooses this button.
	//
	// This must be one of the slot values configured for the slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-button.html#cfn-lex-bot-button-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

