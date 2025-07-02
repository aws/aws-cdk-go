package awslex


// Sets the priority that Amazon Lex should use when eliciting slot values from a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotPriorityProperty := &SlotPriorityProperty{
//   	Priority: jsii.Number(123),
//   	SlotName: jsii.String("slotName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotpriority.html
//
type CfnBot_SlotPriorityProperty struct {
	// The priority that Amazon Lex should apply to the slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotpriority.html#cfn-lex-bot-slotpriority-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The name of the slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotpriority.html#cfn-lex-bot-slotpriority-slotname
	//
	SlotName *string `field:"required" json:"slotName" yaml:"slotName"`
}

