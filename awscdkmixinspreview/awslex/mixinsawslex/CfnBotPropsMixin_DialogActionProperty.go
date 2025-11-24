package mixinsawslex


// Defines the action that the bot executes at runtime when the conversation reaches this step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dialogActionProperty := &DialogActionProperty{
//   	SlotToElicit: jsii.String("slotToElicit"),
//   	SuppressNextMessage: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html
//
type CfnBotPropsMixin_DialogActionProperty struct {
	// If the dialog action is `ElicitSlot` , defines the slot to elicit from the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html#cfn-lex-bot-dialogaction-slottoelicit
	//
	SlotToElicit *string `field:"optional" json:"slotToElicit" yaml:"slotToElicit"`
	// When true the next message for the intent is not used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html#cfn-lex-bot-dialogaction-suppressnextmessage
	//
	SuppressNextMessage interface{} `field:"optional" json:"suppressNextMessage" yaml:"suppressNextMessage"`
	// The action that the bot should execute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html#cfn-lex-bot-dialogaction-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

