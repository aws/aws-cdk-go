package previewawslexmixins


// Provides a list of conditional branches.
//
// Branches are evaluated in the order that they are entered in the list. The first branch with a condition that evaluates to true is executed. The last branch in the list is the default branch. The default branch should not have any condition expression. The default branch is executed if no other branch has a matching condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var slotValueOverrideProperty_ SlotValueOverrideProperty
//
//   conditionalSpecificationProperty := &ConditionalSpecificationProperty{
//   	ConditionalBranches: []interface{}{
//   		&ConditionalBranchProperty{
//   			Condition: &ConditionProperty{
//   				ExpressionString: jsii.String("expressionString"),
//   			},
//   			Name: jsii.String("name"),
//   			NextStep: &DialogStateProperty{
//   				DialogAction: &DialogActionProperty{
//   					SlotToElicit: jsii.String("slotToElicit"),
//   					SuppressNextMessage: jsii.Boolean(false),
//   					Type: jsii.String("type"),
//   				},
//   				Intent: &IntentOverrideProperty{
//   					Name: jsii.String("name"),
//   					Slots: []interface{}{
//   						&SlotValueOverrideMapProperty{
//   							SlotName: jsii.String("slotName"),
//   							SlotValueOverride: &SlotValueOverrideProperty{
//   								Shape: jsii.String("shape"),
//   								Value: &SlotValueProperty{
//   									InterpretedValue: jsii.String("interpretedValue"),
//   								},
//   								Values: []interface{}{
//   									slotValueOverrideProperty_,
//   								},
//   							},
//   						},
//   					},
//   				},
//   				SessionAttributes: []interface{}{
//   					&SessionAttributeProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			Response: &ResponseSpecificationProperty{
//   				AllowInterrupt: jsii.Boolean(false),
//   				MessageGroupsList: []interface{}{
//   					&MessageGroupProperty{
//   						Message: &MessageProperty{
//   							CustomPayload: &CustomPayloadProperty{
//   								Value: jsii.String("value"),
//   							},
//   							ImageResponseCard: &ImageResponseCardProperty{
//   								Buttons: []interface{}{
//   									&ButtonProperty{
//   										Text: jsii.String("text"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								ImageUrl: jsii.String("imageUrl"),
//   								Subtitle: jsii.String("subtitle"),
//   								Title: jsii.String("title"),
//   							},
//   							PlainTextMessage: &PlainTextMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   							SsmlMessage: &SSMLMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						Variations: []interface{}{
//   							&MessageProperty{
//   								CustomPayload: &CustomPayloadProperty{
//   									Value: jsii.String("value"),
//   								},
//   								ImageResponseCard: &ImageResponseCardProperty{
//   									Buttons: []interface{}{
//   										&ButtonProperty{
//   											Text: jsii.String("text"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									ImageUrl: jsii.String("imageUrl"),
//   									Subtitle: jsii.String("subtitle"),
//   									Title: jsii.String("title"),
//   								},
//   								PlainTextMessage: &PlainTextMessageProperty{
//   									Value: jsii.String("value"),
//   								},
//   								SsmlMessage: &SSMLMessageProperty{
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	DefaultBranch: &DefaultConditionalBranchProperty{
//   		NextStep: &DialogStateProperty{
//   			DialogAction: &DialogActionProperty{
//   				SlotToElicit: jsii.String("slotToElicit"),
//   				SuppressNextMessage: jsii.Boolean(false),
//   				Type: jsii.String("type"),
//   			},
//   			Intent: &IntentOverrideProperty{
//   				Name: jsii.String("name"),
//   				Slots: []interface{}{
//   					&SlotValueOverrideMapProperty{
//   						SlotName: jsii.String("slotName"),
//   						SlotValueOverride: &SlotValueOverrideProperty{
//   							Shape: jsii.String("shape"),
//   							Value: &SlotValueProperty{
//   								InterpretedValue: jsii.String("interpretedValue"),
//   							},
//   							Values: []interface{}{
//   								slotValueOverrideProperty_,
//   							},
//   						},
//   					},
//   				},
//   			},
//   			SessionAttributes: []interface{}{
//   				&SessionAttributeProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		Response: &ResponseSpecificationProperty{
//   			AllowInterrupt: jsii.Boolean(false),
//   			MessageGroupsList: []interface{}{
//   				&MessageGroupProperty{
//   					Message: &MessageProperty{
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
//   					Variations: []interface{}{
//   						&MessageProperty{
//   							CustomPayload: &CustomPayloadProperty{
//   								Value: jsii.String("value"),
//   							},
//   							ImageResponseCard: &ImageResponseCardProperty{
//   								Buttons: []interface{}{
//   									&ButtonProperty{
//   										Text: jsii.String("text"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								ImageUrl: jsii.String("imageUrl"),
//   								Subtitle: jsii.String("subtitle"),
//   								Title: jsii.String("title"),
//   							},
//   							PlainTextMessage: &PlainTextMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   							SsmlMessage: &SSMLMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	IsActive: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-conditionalspecification.html
//
type CfnBotPropsMixin_ConditionalSpecificationProperty struct {
	// A list of conditional branches.
	//
	// A conditional branch is made up of a condition, a response and a next step. The response and next step are executed when the condition is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-conditionalspecification.html#cfn-lex-bot-conditionalspecification-conditionalbranches
	//
	ConditionalBranches interface{} `field:"optional" json:"conditionalBranches" yaml:"conditionalBranches"`
	// The conditional branch that should be followed when the conditions for other branches are not satisfied.
	//
	// A conditional branch is made up of a condition, a response and a next step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-conditionalspecification.html#cfn-lex-bot-conditionalspecification-defaultbranch
	//
	DefaultBranch interface{} `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// Determines whether a conditional branch is active.
	//
	// When `IsActive` is false, the conditions are not evaluated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-conditionalspecification.html#cfn-lex-bot-conditionalspecification-isactive
	//
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
}

