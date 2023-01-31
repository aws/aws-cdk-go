package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   intentOverrideProperty := &intentOverrideProperty{
//   	name: jsii.String("name"),
//   	slots: []interface{}{
//   		&slotValueOverrideMapProperty{
//   			slotName: jsii.String("slotName"),
//   			slotValueOverride: &slotValueOverrideProperty{
//   				shape: jsii.String("shape"),
//   				value: &slotValueProperty{
//   					interpretedValue: jsii.String("interpretedValue"),
//   				},
//   				values: []interface{}{
//   					slotValueOverrideProperty_,
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBot_IntentOverrideProperty struct {
	// `CfnBot.IntentOverrideProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnBot.IntentOverrideProperty.Slots`.
	Slots interface{} `field:"optional" json:"slots" yaml:"slots"`
}

