package previewawslexmixins


// A set of actions that Amazon Lex should run if none of the other conditions are met.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var slotValueOverrideProperty_ SlotValueOverrideProperty
//
//   defaultConditionalBranchProperty := &DefaultConditionalBranchProperty{
//   	NextStep: &DialogStateProperty{
//   		DialogAction: &DialogActionProperty{
//   			SlotToElicit: jsii.String("slotToElicit"),
//   			SuppressNextMessage: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   		},
//   		Intent: &IntentOverrideProperty{
//   			Name: jsii.String("name"),
//   			Slots: []interface{}{
//   				&SlotValueOverrideMapProperty{
//   					SlotName: jsii.String("slotName"),
//   					SlotValueOverride: &SlotValueOverrideProperty{
//   						Shape: jsii.String("shape"),
//   						Value: &SlotValueProperty{
//   							InterpretedValue: jsii.String("interpretedValue"),
//   						},
//   						Values: []interface{}{
//   							slotValueOverrideProperty_,
//   						},
//   					},
//   				},
//   			},
//   		},
//   		SessionAttributes: []interface{}{
//   			&SessionAttributeProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Response: &ResponseSpecificationProperty{
//   		AllowInterrupt: jsii.Boolean(false),
//   		MessageGroupsList: []interface{}{
//   			&MessageGroupProperty{
//   				Message: &MessageProperty{
//   					CustomPayload: &CustomPayloadProperty{
//   						Value: jsii.String("value"),
//   					},
//   					ImageResponseCard: &ImageResponseCardProperty{
//   						Buttons: []interface{}{
//   							&ButtonProperty{
//   								Text: jsii.String("text"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						ImageUrl: jsii.String("imageUrl"),
//   						Subtitle: jsii.String("subtitle"),
//   						Title: jsii.String("title"),
//   					},
//   					PlainTextMessage: &PlainTextMessageProperty{
//   						Value: jsii.String("value"),
//   					},
//   					SsmlMessage: &SSMLMessageProperty{
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Variations: []interface{}{
//   					&MessageProperty{
//   						CustomPayload: &CustomPayloadProperty{
//   							Value: jsii.String("value"),
//   						},
//   						ImageResponseCard: &ImageResponseCardProperty{
//   							Buttons: []interface{}{
//   								&ButtonProperty{
//   									Text: jsii.String("text"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							ImageUrl: jsii.String("imageUrl"),
//   							Subtitle: jsii.String("subtitle"),
//   							Title: jsii.String("title"),
//   						},
//   						PlainTextMessage: &PlainTextMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   						SsmlMessage: &SSMLMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-defaultconditionalbranch.html
//
type CfnBotPropsMixin_DefaultConditionalBranchProperty struct {
	// The next step in the conversation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-defaultconditionalbranch.html#cfn-lex-bot-defaultconditionalbranch-nextstep
	//
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-defaultconditionalbranch.html#cfn-lex-bot-defaultconditionalbranch-response
	//
	Response interface{} `field:"optional" json:"response" yaml:"response"`
}

