package awslex


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
	// `CfnBot.SlotValueOverrideProperty.Shape`.
	Shape *string `field:"optional" json:"shape" yaml:"shape"`
	// `CfnBot.SlotValueOverrideProperty.Value`.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
	// `CfnBot.SlotValueOverrideProperty.Values`.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

