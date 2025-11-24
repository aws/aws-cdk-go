package mixinsawslex


// The current state of the conversation with the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var slotValueOverrideProperty_ SlotValueOverrideProperty
//
//   dialogStateProperty := &DialogStateProperty{
//   	DialogAction: &DialogActionProperty{
//   		SlotToElicit: jsii.String("slotToElicit"),
//   		SuppressNextMessage: jsii.Boolean(false),
//   		Type: jsii.String("type"),
//   	},
//   	Intent: &IntentOverrideProperty{
//   		Name: jsii.String("name"),
//   		Slots: []interface{}{
//   			&SlotValueOverrideMapProperty{
//   				SlotName: jsii.String("slotName"),
//   				SlotValueOverride: &SlotValueOverrideProperty{
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
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogstate.html
//
type CfnBotPropsMixin_DialogStateProperty struct {
	// Defines the action that the bot executes at runtime when the conversation reaches this step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogstate.html#cfn-lex-bot-dialogstate-dialogaction
	//
	DialogAction interface{} `field:"optional" json:"dialogAction" yaml:"dialogAction"`
	// Override settings to configure the intent state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogstate.html#cfn-lex-bot-dialogstate-intent
	//
	Intent interface{} `field:"optional" json:"intent" yaml:"intent"`
	// Map of key/value pairs representing session-specific context information.
	//
	// It contains application information passed between Amazon Lex and a client application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogstate.html#cfn-lex-bot-dialogstate-sessionattributes
	//
	SessionAttributes interface{} `field:"optional" json:"sessionAttributes" yaml:"sessionAttributes"`
}

