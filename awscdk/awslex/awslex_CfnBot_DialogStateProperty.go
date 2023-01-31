package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   dialogStateProperty := &dialogStateProperty{
//   	dialogAction: &dialogActionProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		slotToElicit: jsii.String("slotToElicit"),
//   		suppressNextMessage: jsii.Boolean(false),
//   	},
//   	intent: &intentOverrideProperty{
//   		name: jsii.String("name"),
//   		slots: []interface{}{
//   			&slotValueOverrideMapProperty{
//   				slotName: jsii.String("slotName"),
//   				slotValueOverride: &slotValueOverrideProperty{
//   					shape: jsii.String("shape"),
//   					value: &slotValueProperty{
//   						interpretedValue: jsii.String("interpretedValue"),
//   					},
//   					values: []interface{}{
//   						slotValueOverrideProperty_,
//   					},
//   				},
//   			},
//   		},
//   	},
//   	sessionAttributes: []interface{}{
//   		&sessionAttributeProperty{
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBot_DialogStateProperty struct {
	// `CfnBot.DialogStateProperty.DialogAction`.
	DialogAction interface{} `field:"optional" json:"dialogAction" yaml:"dialogAction"`
	// `CfnBot.DialogStateProperty.Intent`.
	Intent interface{} `field:"optional" json:"intent" yaml:"intent"`
	// `CfnBot.DialogStateProperty.SessionAttributes`.
	SessionAttributes interface{} `field:"optional" json:"sessionAttributes" yaml:"sessionAttributes"`
}

