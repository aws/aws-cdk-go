package mixinsawslex


// Subslot type composition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subSlotTypeCompositionProperty := &SubSlotTypeCompositionProperty{
//   	Name: jsii.String("name"),
//   	SlotTypeId: jsii.String("slotTypeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslottypecomposition.html
//
type CfnBotPropsMixin_SubSlotTypeCompositionProperty struct {
	// Name of a constituent sub slot inside a composite slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslottypecomposition.html#cfn-lex-bot-subslottypecomposition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The unique identifier assigned to a slot type.
	//
	// This refers to either a built-in slot type or the unique slotTypeId of a custom slot type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslottypecomposition.html#cfn-lex-bot-subslottypecomposition-slottypeid
	//
	SlotTypeId *string `field:"optional" json:"slotTypeId" yaml:"slotTypeId"`
}

