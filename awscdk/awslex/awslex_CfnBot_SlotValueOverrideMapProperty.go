package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   slotValueOverrideMapProperty := &slotValueOverrideMapProperty{
//   	slotName: jsii.String("slotName"),
//   	slotValueOverride: &slotValueOverrideProperty{
//   		shape: jsii.String("shape"),
//   		value: &slotValueProperty{
//   			interpretedValue: jsii.String("interpretedValue"),
//   		},
//   		values: []interface{}{
//   			slotValueOverrideProperty_,
//   		},
//   	},
//   }
//
type CfnBot_SlotValueOverrideMapProperty struct {
	// `CfnBot.SlotValueOverrideMapProperty.SlotName`.
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
	// `CfnBot.SlotValueOverrideMapProperty.SlotValueOverride`.
	SlotValueOverride interface{} `field:"optional" json:"slotValueOverride" yaml:"slotValueOverride"`
}

