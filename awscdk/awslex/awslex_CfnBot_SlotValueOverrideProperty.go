package awslex


// The slot values that Amazon Lex uses when it sets slot values in a dialog step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   slotValueOverrideProperty := &slotValueOverrideProperty{
//   	shape: jsii.String("shape"),
//   	value: &slotValueProperty{
//   		interpretedValue: jsii.String("interpretedValue"),
//   	},
//   	values: []interface{}{
//   		&slotValueOverrideProperty{
//   			shape: jsii.String("shape"),
//   			value: &slotValueProperty{
//   				interpretedValue: jsii.String("interpretedValue"),
//   			},
//   			values: []interface{}{
//   				slotValueOverrideProperty_,
//   			},
//   		},
//   	},
//   }
//
type CfnBot_SlotValueOverrideProperty struct {
	// When the shape value is `List` , it indicates that the `values` field contains a list of slot values.
	//
	// When the value is `Scalar` , it indicates that the `value` field contains a single value.
	Shape *string `field:"optional" json:"shape" yaml:"shape"`
	// The current value of the slot.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
	// A list of one or more values that the user provided for the slot.
	//
	// For example, for a slot that elicits pizza toppings, the values might be "pepperoni" and "pineapple."
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

