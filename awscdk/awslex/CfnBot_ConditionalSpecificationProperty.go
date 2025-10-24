package awslex


// Provides a list of conditional branches.
//
// Branches are evaluated in the order that they are entered in the list. The first branch with a condition that evaluates to true is executed. The last branch in the list is the default branch. The default branch should not have any condition expression. The default branch is executed if no other branch has a matching condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					SlotToElicit: jsii.String("slotToElicit"),
//   					SuppressNextMessage: jsii.Boolean(false),
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
//
//   						// the properties below are optional
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			Response: &ResponseSpecificationProperty{
//   				MessageGroupsList: []interface{}{
//   					&MessageGroupProperty{
//   						Message: &MessageProperty{
//   							CustomPayload: &CustomPayloadProperty{
//   								Value: jsii.String("value"),
//   							},
//   							ImageResponseCard: &ImageResponseCardProperty{
//   								Title: jsii.String("title"),
//
//   								// the properties below are optional
//   								Buttons: []interface{}{
//   									&ButtonProperty{
//   										Text: jsii.String("text"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								ImageUrl: jsii.String("imageUrl"),
//   								Subtitle: jsii.String("subtitle"),
//   							},
//   							PlainTextMessage: &PlainTextMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   							SsmlMessage: &SSMLMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						Variations: []interface{}{
//   							&MessageProperty{
//   								CustomPayload: &CustomPayloadProperty{
//   									Value: jsii.String("value"),
//   								},
//   								ImageResponseCard: &ImageResponseCardProperty{
//   									Title: jsii.String("title"),
//
//   									// the properties below are optional
//   									Buttons: []interface{}{
//   										&ButtonProperty{
//   											Text: jsii.String("text"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									ImageUrl: jsii.String("imageUrl"),
//   									Subtitle: jsii.String("subtitle"),
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
//
//   				// the properties below are optional
//   				AllowInterrupt: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	DefaultBranch: &DefaultConditionalBranchProperty{
//   		NextStep: &DialogStateProperty{
//   			DialogAction: &DialogActionProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				SlotToElicit: jsii.String("slotToElicit"),
//   				SuppressNextMessage: jsii.Boolean(false),
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
//
//   					// the properties below are optional
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		Response: &ResponseSpecificationProperty{
//   			MessageGroupsList: []interface{}{
//   				&MessageGroupProperty{
//   					Message: &MessageProperty{
//   						CustomPayload: &CustomPayloadProperty{
//   							Value: jsii.String("value"),
//   						},
//   						ImageResponseCard: &ImageResponseCardProperty{
//   							Title: jsii.String("title"),
//
//   							// the properties below are optional
//   							Buttons: []interface{}{
//   								&ButtonProperty{
//   									Text: jsii.String("text"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							ImageUrl: jsii.String("imageUrl"),
//   							Subtitle: jsii.String("subtitle"),
//   						},
//   						PlainTextMessage: &PlainTextMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   						SsmlMessage: &SSMLMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Variations: []interface{}{
//   						&MessageProperty{
//   							CustomPayload: &CustomPayloadProperty{
//   								Value: jsii.String("value"),
//   							},
//   							ImageResponseCard: &ImageResponseCardProperty{
//   								Title: jsii.String("title"),
//
//   								// the properties below are optional
//   								Buttons: []interface{}{
//   									&ButtonProperty{
//   										Text: jsii.String("text"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								ImageUrl: jsii.String("imageUrl"),
//   								Subtitle: jsii.String("subtitle"),
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
//
//   			// the properties below are optional
//   			AllowInterrupt: jsii.Boolean(false),
//   		},
//   	},
//   	IsActive: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-conditionalspecification.html
//
type CfnBot_ConditionalSpecificationProperty struct {
	// A list of conditional branches.
	//
	// A conditional branch is made up of a condition, a response and a next step. The response and next step are executed when the condition is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-conditionalspecification.html#cfn-lex-bot-conditionalspecification-conditionalbranches
	//
	ConditionalBranches interface{} `field:"required" json:"conditionalBranches" yaml:"conditionalBranches"`
	// The conditional branch that should be followed when the conditions for other branches are not satisfied.
	//
	// A conditional branch is made up of a condition, a response and a next step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-conditionalspecification.html#cfn-lex-bot-conditionalspecification-defaultbranch
	//
	DefaultBranch interface{} `field:"required" json:"defaultBranch" yaml:"defaultBranch"`
	// Determines whether a conditional branch is active.
	//
	// When `IsActive` is false, the conditions are not evaluated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-conditionalspecification.html#cfn-lex-bot-conditionalspecification-isactive
	//
	IsActive interface{} `field:"required" json:"isActive" yaml:"isActive"`
}

