package awslex


// The slot values that Amazon Lex uses when it sets slot values in a dialog step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ SlotValueOverrideProperty
//
//   slotValueOverrideProperty := &SlotValueOverrideProperty{
//   	Shape: jsii.String("shape"),
//   	Value: &SlotValueProperty{
//   		InterpretedValue: jsii.String("interpretedValue"),
//   	},
//   	Values: []interface{}{
//   		slotValueOverrideProperty_,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueoverride.html
//
type CfnBot_SlotValueOverrideProperty struct {
	// When the shape value is `List` , it indicates that the `values` field contains a list of slot values.
	//
	// When the value is `Scalar` , it indicates that the `value` field contains a single value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueoverride.html#cfn-lex-bot-slotvalueoverride-shape
	//
	Shape *string `field:"optional" json:"shape" yaml:"shape"`
	// The current value of the slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueoverride.html#cfn-lex-bot-slotvalueoverride-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
	// A list of one or more values that the user provided for the slot.
	//
	// For example, for a slot that elicits pizza toppings, the values might be "pepperoni" and "pineapple."
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotvalueoverride.html#cfn-lex-bot-slotvalueoverride-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

