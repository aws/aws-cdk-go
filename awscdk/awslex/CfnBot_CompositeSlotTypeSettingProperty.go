package awslex


// A composite slot is a combination of two or more slots that capture multiple pieces of information in a single user input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   compositeSlotTypeSettingProperty := &CompositeSlotTypeSettingProperty{
//   	SubSlots: []interface{}{
//   		&SubSlotTypeCompositionProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			SlotTypeId: jsii.String("slotTypeId"),
//   			SlotTypeName: jsii.String("slotTypeName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-compositeslottypesetting.html
//
type CfnBot_CompositeSlotTypeSettingProperty struct {
	// Subslots in the composite slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-compositeslottypesetting.html#cfn-lex-bot-compositeslottypesetting-subslots
	//
	SubSlots interface{} `field:"optional" json:"subSlots" yaml:"subSlots"`
}

