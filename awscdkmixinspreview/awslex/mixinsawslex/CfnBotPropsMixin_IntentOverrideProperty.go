package mixinsawslex


// Override settings to configure the intent state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var slotValueOverrideProperty_ SlotValueOverrideProperty
//
//   intentOverrideProperty := &IntentOverrideProperty{
//   	Name: jsii.String("name"),
//   	Slots: []interface{}{
//   		&SlotValueOverrideMapProperty{
//   			SlotName: jsii.String("slotName"),
//   			SlotValueOverride: &SlotValueOverrideProperty{
//   				Shape: jsii.String("shape"),
//   				Value: &SlotValueProperty{
//   					InterpretedValue: jsii.String("interpretedValue"),
//   				},
//   				Values: []interface{}{
//   					slotValueOverrideProperty_,
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intentoverride.html
//
type CfnBotPropsMixin_IntentOverrideProperty struct {
	// The name of the intent.
	//
	// Only required when you're switching intents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intentoverride.html#cfn-lex-bot-intentoverride-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A map of all of the slot value overrides for the intent.
	//
	// The name of the slot maps to the value of the slot. Slots that are not included in the map aren't overridden.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-intentoverride.html#cfn-lex-bot-intentoverride-slots
	//
	Slots interface{} `field:"optional" json:"slots" yaml:"slots"`
}

