package awslex


// The current state of the conversation with the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   dialogStateProperty := &DialogStateProperty{
//   	DialogAction: &DialogActionProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		SlotToElicit: jsii.String("slotToElicit"),
//   		SuppressNextMessage: jsii.Boolean(false),
//   	},
//   	Intent: &IntentOverrideProperty{
//   		Name: jsii.String("name"),
//   		Slots: []interface{}{
//   			&SlotValueOverrideMapProperty{
//   				SlotName: jsii.String("slotName"),
//   				SlotValueOverride: &slotValueOverrideProperty{
//   					Shape: jsii.String("shape"),
//   					Value: &SlotValueProperty{
//   						InterpretedValue: jsii.String("interpretedValue"),
//   					},
//   					Values: []interface{}{
//   						slotValueOverrideProperty_,
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SessionAttributes: []interface{}{
//   		&SessionAttributeProperty{
//   			Key: jsii.String("key"),
//
//   			// the properties below are optional
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBot_DialogStateProperty struct {
	// Defines the action that the bot executes at runtime when the conversation reaches this step.
	DialogAction interface{} `field:"optional" json:"dialogAction" yaml:"dialogAction"`
	// Override settings to configure the intent state.
	Intent interface{} `field:"optional" json:"intent" yaml:"intent"`
	// Map of key/value pairs representing session-specific context information.
	//
	// It contains application information passed between Amazon Lex and a client application.
	SessionAttributes interface{} `field:"optional" json:"sessionAttributes" yaml:"sessionAttributes"`
}

