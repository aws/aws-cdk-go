package awslex


// Defines the action that the bot executes at runtime when the conversation reaches this step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dialogActionProperty := &DialogActionProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	SlotToElicit: jsii.String("slotToElicit"),
//   	SuppressNextMessage: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html
//
type CfnBot_DialogActionProperty struct {
	// The action that the bot should execute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html#cfn-lex-bot-dialogaction-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// If the dialog action is `ElicitSlot` , defines the slot to elicit from the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html#cfn-lex-bot-dialogaction-slottoelicit
	//
	SlotToElicit *string `field:"optional" json:"slotToElicit" yaml:"slotToElicit"`
	// When true the next message for the intent is not used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html#cfn-lex-bot-dialogaction-suppressnextmessage
	//
	SuppressNextMessage interface{} `field:"optional" json:"suppressNextMessage" yaml:"suppressNextMessage"`
}

