package awslex


// Maps a slot name to the [SlotValueOverride](https://docs.aws.amazon.com/lexv2/latest/APIReference/API_SlotValueOverride.html) object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   slotValueOverrideMapProperty := &SlotValueOverrideMapProperty{
//   	SlotName: jsii.String("slotName"),
//   	SlotValueOverride: &slotValueOverrideProperty{
//   		Shape: jsii.String("shape"),
//   		Value: &SlotValueProperty{
//   			InterpretedValue: jsii.String("interpretedValue"),
//   		},
//   		Values: []interface{}{
//   			slotValueOverrideProperty_,
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueoverridemap.html
//
type CfnBot_SlotValueOverrideMapProperty struct {
	// The name of the slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueoverridemap.html#cfn-lex-bot-slotvalueoverridemap-slotname
	//
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
	// The SlotValueOverride object to which the slot name will be mapped.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueoverridemap.html#cfn-lex-bot-slotvalueoverridemap-slotvalueoverride
	//
	SlotValueOverride interface{} `field:"optional" json:"slotValueOverride" yaml:"slotValueOverride"`
}

