package awslex


// Represents an action that the user wants to perform.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   intentProperty := &IntentProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DialogCodeHook: &DialogCodeHookSettingProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	FulfillmentCodeHook: &FulfillmentCodeHookSettingProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		FulfillmentUpdatesSpecification: &FulfillmentUpdatesSpecificationProperty{
//   			Active: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			StartResponse: &FulfillmentStartResponseSpecificationProperty{
//   				DelayInSeconds: jsii.Number(123),
//   				MessageGroups: []interface{}{
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
//   			TimeoutInSeconds: jsii.Number(123),
//   			UpdateResponse: &FulfillmentUpdateResponseSpecificationProperty{
//   				FrequencyInSeconds: jsii.Number(123),
//   				MessageGroups: []interface{}{
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
//   		IsActive: jsii.Boolean(false),
//   		PostFulfillmentStatusSpecification: &PostFulfillmentStatusSpecificationProperty{
//   			FailureConditional: &ConditionalSpecificationProperty{
//   				ConditionalBranches: []interface{}{
//   					&ConditionalBranchProperty{
//   						Condition: &ConditionProperty{
//   							ExpressionString: jsii.String("expressionString"),
//   						},
//   						Name: jsii.String("name"),
//   						NextStep: &DialogStateProperty{
//   							DialogAction: &DialogActionProperty{
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								SlotToElicit: jsii.String("slotToElicit"),
//   								SuppressNextMessage: jsii.Boolean(false),
//   							},
//   							Intent: &IntentOverrideProperty{
//   								Name: jsii.String("name"),
//   								Slots: []interface{}{
//   									&SlotValueOverrideMapProperty{
//   										SlotName: jsii.String("slotName"),
//   										SlotValueOverride: &slotValueOverrideProperty{
//   											Shape: jsii.String("shape"),
//   											Value: &SlotValueProperty{
//   												InterpretedValue: jsii.String("interpretedValue"),
//   											},
//   											Values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							SessionAttributes: []interface{}{
//   								&SessionAttributeProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						Response: &ResponseSpecificationProperty{
//   							MessageGroupsList: []interface{}{
//   								&MessageGroupProperty{
//   									Message: &MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Variations: []interface{}{
//   										&MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							AllowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				DefaultBranch: &DefaultConditionalBranchProperty{
//   					NextStep: &DialogStateProperty{
//   						DialogAction: &DialogActionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							SlotToElicit: jsii.String("slotToElicit"),
//   							SuppressNextMessage: jsii.Boolean(false),
//   						},
//   						Intent: &IntentOverrideProperty{
//   							Name: jsii.String("name"),
//   							Slots: []interface{}{
//   								&SlotValueOverrideMapProperty{
//   									SlotName: jsii.String("slotName"),
//   									SlotValueOverride: &slotValueOverrideProperty{
//   										Shape: jsii.String("shape"),
//   										Value: &SlotValueProperty{
//   											InterpretedValue: jsii.String("interpretedValue"),
//   										},
//   										Values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						SessionAttributes: []interface{}{
//   							&SessionAttributeProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					Response: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   				IsActive: jsii.Boolean(false),
//   			},
//   			FailureNextStep: &DialogStateProperty{
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
//   							SlotValueOverride: &slotValueOverrideProperty{
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
//   			FailureResponse: &ResponseSpecificationProperty{
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
//   			SuccessConditional: &ConditionalSpecificationProperty{
//   				ConditionalBranches: []interface{}{
//   					&ConditionalBranchProperty{
//   						Condition: &ConditionProperty{
//   							ExpressionString: jsii.String("expressionString"),
//   						},
//   						Name: jsii.String("name"),
//   						NextStep: &DialogStateProperty{
//   							DialogAction: &DialogActionProperty{
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								SlotToElicit: jsii.String("slotToElicit"),
//   								SuppressNextMessage: jsii.Boolean(false),
//   							},
//   							Intent: &IntentOverrideProperty{
//   								Name: jsii.String("name"),
//   								Slots: []interface{}{
//   									&SlotValueOverrideMapProperty{
//   										SlotName: jsii.String("slotName"),
//   										SlotValueOverride: &slotValueOverrideProperty{
//   											Shape: jsii.String("shape"),
//   											Value: &SlotValueProperty{
//   												InterpretedValue: jsii.String("interpretedValue"),
//   											},
//   											Values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							SessionAttributes: []interface{}{
//   								&SessionAttributeProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						Response: &ResponseSpecificationProperty{
//   							MessageGroupsList: []interface{}{
//   								&MessageGroupProperty{
//   									Message: &MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Variations: []interface{}{
//   										&MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							AllowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				DefaultBranch: &DefaultConditionalBranchProperty{
//   					NextStep: &DialogStateProperty{
//   						DialogAction: &DialogActionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							SlotToElicit: jsii.String("slotToElicit"),
//   							SuppressNextMessage: jsii.Boolean(false),
//   						},
//   						Intent: &IntentOverrideProperty{
//   							Name: jsii.String("name"),
//   							Slots: []interface{}{
//   								&SlotValueOverrideMapProperty{
//   									SlotName: jsii.String("slotName"),
//   									SlotValueOverride: &slotValueOverrideProperty{
//   										Shape: jsii.String("shape"),
//   										Value: &SlotValueProperty{
//   											InterpretedValue: jsii.String("interpretedValue"),
//   										},
//   										Values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						SessionAttributes: []interface{}{
//   							&SessionAttributeProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					Response: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   				IsActive: jsii.Boolean(false),
//   			},
//   			SuccessNextStep: &DialogStateProperty{
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
//   							SlotValueOverride: &slotValueOverrideProperty{
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
//   			SuccessResponse: &ResponseSpecificationProperty{
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
//   			TimeoutConditional: &ConditionalSpecificationProperty{
//   				ConditionalBranches: []interface{}{
//   					&ConditionalBranchProperty{
//   						Condition: &ConditionProperty{
//   							ExpressionString: jsii.String("expressionString"),
//   						},
//   						Name: jsii.String("name"),
//   						NextStep: &DialogStateProperty{
//   							DialogAction: &DialogActionProperty{
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								SlotToElicit: jsii.String("slotToElicit"),
//   								SuppressNextMessage: jsii.Boolean(false),
//   							},
//   							Intent: &IntentOverrideProperty{
//   								Name: jsii.String("name"),
//   								Slots: []interface{}{
//   									&SlotValueOverrideMapProperty{
//   										SlotName: jsii.String("slotName"),
//   										SlotValueOverride: &slotValueOverrideProperty{
//   											Shape: jsii.String("shape"),
//   											Value: &SlotValueProperty{
//   												InterpretedValue: jsii.String("interpretedValue"),
//   											},
//   											Values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							SessionAttributes: []interface{}{
//   								&SessionAttributeProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						Response: &ResponseSpecificationProperty{
//   							MessageGroupsList: []interface{}{
//   								&MessageGroupProperty{
//   									Message: &MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Variations: []interface{}{
//   										&MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							AllowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				DefaultBranch: &DefaultConditionalBranchProperty{
//   					NextStep: &DialogStateProperty{
//   						DialogAction: &DialogActionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							SlotToElicit: jsii.String("slotToElicit"),
//   							SuppressNextMessage: jsii.Boolean(false),
//   						},
//   						Intent: &IntentOverrideProperty{
//   							Name: jsii.String("name"),
//   							Slots: []interface{}{
//   								&SlotValueOverrideMapProperty{
//   									SlotName: jsii.String("slotName"),
//   									SlotValueOverride: &slotValueOverrideProperty{
//   										Shape: jsii.String("shape"),
//   										Value: &SlotValueProperty{
//   											InterpretedValue: jsii.String("interpretedValue"),
//   										},
//   										Values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						SessionAttributes: []interface{}{
//   							&SessionAttributeProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					Response: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   				IsActive: jsii.Boolean(false),
//   			},
//   			TimeoutNextStep: &DialogStateProperty{
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
//   							SlotValueOverride: &slotValueOverrideProperty{
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
//   			TimeoutResponse: &ResponseSpecificationProperty{
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
//   	InitialResponseSetting: &InitialResponseSettingProperty{
//   		CodeHook: &DialogCodeHookInvocationSettingProperty{
//   			EnableCodeHookInvocation: jsii.Boolean(false),
//   			IsActive: jsii.Boolean(false),
//   			PostCodeHookSpecification: &PostDialogCodeHookInvocationSpecificationProperty{
//   				FailureConditional: &ConditionalSpecificationProperty{
//   					ConditionalBranches: []interface{}{
//   						&ConditionalBranchProperty{
//   							Condition: &ConditionProperty{
//   								ExpressionString: jsii.String("expressionString"),
//   							},
//   							Name: jsii.String("name"),
//   							NextStep: &DialogStateProperty{
//   								DialogAction: &DialogActionProperty{
//   									Type: jsii.String("type"),
//
//   									// the properties below are optional
//   									SlotToElicit: jsii.String("slotToElicit"),
//   									SuppressNextMessage: jsii.Boolean(false),
//   								},
//   								Intent: &IntentOverrideProperty{
//   									Name: jsii.String("name"),
//   									Slots: []interface{}{
//   										&SlotValueOverrideMapProperty{
//   											SlotName: jsii.String("slotName"),
//   											SlotValueOverride: &slotValueOverrideProperty{
//   												Shape: jsii.String("shape"),
//   												Value: &SlotValueProperty{
//   													InterpretedValue: jsii.String("interpretedValue"),
//   												},
//   												Values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								SessionAttributes: []interface{}{
//   									&SessionAttributeProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							Response: &ResponseSpecificationProperty{
//   								MessageGroupsList: []interface{}{
//   									&MessageGroupProperty{
//   										Message: &MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										Variations: []interface{}{
//   											&MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					DefaultBranch: &DefaultConditionalBranchProperty{
//   						NextStep: &DialogStateProperty{
//   							DialogAction: &DialogActionProperty{
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								SlotToElicit: jsii.String("slotToElicit"),
//   								SuppressNextMessage: jsii.Boolean(false),
//   							},
//   							Intent: &IntentOverrideProperty{
//   								Name: jsii.String("name"),
//   								Slots: []interface{}{
//   									&SlotValueOverrideMapProperty{
//   										SlotName: jsii.String("slotName"),
//   										SlotValueOverride: &slotValueOverrideProperty{
//   											Shape: jsii.String("shape"),
//   											Value: &SlotValueProperty{
//   												InterpretedValue: jsii.String("interpretedValue"),
//   											},
//   											Values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							SessionAttributes: []interface{}{
//   								&SessionAttributeProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						Response: &ResponseSpecificationProperty{
//   							MessageGroupsList: []interface{}{
//   								&MessageGroupProperty{
//   									Message: &MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Variations: []interface{}{
//   										&MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							AllowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   					IsActive: jsii.Boolean(false),
//   				},
//   				FailureNextStep: &DialogStateProperty{
//   					DialogAction: &DialogActionProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						SlotToElicit: jsii.String("slotToElicit"),
//   						SuppressNextMessage: jsii.Boolean(false),
//   					},
//   					Intent: &IntentOverrideProperty{
//   						Name: jsii.String("name"),
//   						Slots: []interface{}{
//   							&SlotValueOverrideMapProperty{
//   								SlotName: jsii.String("slotName"),
//   								SlotValueOverride: &slotValueOverrideProperty{
//   									Shape: jsii.String("shape"),
//   									Value: &SlotValueProperty{
//   										InterpretedValue: jsii.String("interpretedValue"),
//   									},
//   									Values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SessionAttributes: []interface{}{
//   						&SessionAttributeProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				FailureResponse: &ResponseSpecificationProperty{
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   				},
//   				SuccessConditional: &ConditionalSpecificationProperty{
//   					ConditionalBranches: []interface{}{
//   						&ConditionalBranchProperty{
//   							Condition: &ConditionProperty{
//   								ExpressionString: jsii.String("expressionString"),
//   							},
//   							Name: jsii.String("name"),
//   							NextStep: &DialogStateProperty{
//   								DialogAction: &DialogActionProperty{
//   									Type: jsii.String("type"),
//
//   									// the properties below are optional
//   									SlotToElicit: jsii.String("slotToElicit"),
//   									SuppressNextMessage: jsii.Boolean(false),
//   								},
//   								Intent: &IntentOverrideProperty{
//   									Name: jsii.String("name"),
//   									Slots: []interface{}{
//   										&SlotValueOverrideMapProperty{
//   											SlotName: jsii.String("slotName"),
//   											SlotValueOverride: &slotValueOverrideProperty{
//   												Shape: jsii.String("shape"),
//   												Value: &SlotValueProperty{
//   													InterpretedValue: jsii.String("interpretedValue"),
//   												},
//   												Values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								SessionAttributes: []interface{}{
//   									&SessionAttributeProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							Response: &ResponseSpecificationProperty{
//   								MessageGroupsList: []interface{}{
//   									&MessageGroupProperty{
//   										Message: &MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										Variations: []interface{}{
//   											&MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					DefaultBranch: &DefaultConditionalBranchProperty{
//   						NextStep: &DialogStateProperty{
//   							DialogAction: &DialogActionProperty{
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								SlotToElicit: jsii.String("slotToElicit"),
//   								SuppressNextMessage: jsii.Boolean(false),
//   							},
//   							Intent: &IntentOverrideProperty{
//   								Name: jsii.String("name"),
//   								Slots: []interface{}{
//   									&SlotValueOverrideMapProperty{
//   										SlotName: jsii.String("slotName"),
//   										SlotValueOverride: &slotValueOverrideProperty{
//   											Shape: jsii.String("shape"),
//   											Value: &SlotValueProperty{
//   												InterpretedValue: jsii.String("interpretedValue"),
//   											},
//   											Values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							SessionAttributes: []interface{}{
//   								&SessionAttributeProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						Response: &ResponseSpecificationProperty{
//   							MessageGroupsList: []interface{}{
//   								&MessageGroupProperty{
//   									Message: &MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Variations: []interface{}{
//   										&MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							AllowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   					IsActive: jsii.Boolean(false),
//   				},
//   				SuccessNextStep: &DialogStateProperty{
//   					DialogAction: &DialogActionProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						SlotToElicit: jsii.String("slotToElicit"),
//   						SuppressNextMessage: jsii.Boolean(false),
//   					},
//   					Intent: &IntentOverrideProperty{
//   						Name: jsii.String("name"),
//   						Slots: []interface{}{
//   							&SlotValueOverrideMapProperty{
//   								SlotName: jsii.String("slotName"),
//   								SlotValueOverride: &slotValueOverrideProperty{
//   									Shape: jsii.String("shape"),
//   									Value: &SlotValueProperty{
//   										InterpretedValue: jsii.String("interpretedValue"),
//   									},
//   									Values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SessionAttributes: []interface{}{
//   						&SessionAttributeProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				SuccessResponse: &ResponseSpecificationProperty{
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   				},
//   				TimeoutConditional: &ConditionalSpecificationProperty{
//   					ConditionalBranches: []interface{}{
//   						&ConditionalBranchProperty{
//   							Condition: &ConditionProperty{
//   								ExpressionString: jsii.String("expressionString"),
//   							},
//   							Name: jsii.String("name"),
//   							NextStep: &DialogStateProperty{
//   								DialogAction: &DialogActionProperty{
//   									Type: jsii.String("type"),
//
//   									// the properties below are optional
//   									SlotToElicit: jsii.String("slotToElicit"),
//   									SuppressNextMessage: jsii.Boolean(false),
//   								},
//   								Intent: &IntentOverrideProperty{
//   									Name: jsii.String("name"),
//   									Slots: []interface{}{
//   										&SlotValueOverrideMapProperty{
//   											SlotName: jsii.String("slotName"),
//   											SlotValueOverride: &slotValueOverrideProperty{
//   												Shape: jsii.String("shape"),
//   												Value: &SlotValueProperty{
//   													InterpretedValue: jsii.String("interpretedValue"),
//   												},
//   												Values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								SessionAttributes: []interface{}{
//   									&SessionAttributeProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							Response: &ResponseSpecificationProperty{
//   								MessageGroupsList: []interface{}{
//   									&MessageGroupProperty{
//   										Message: &MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										Variations: []interface{}{
//   											&MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					DefaultBranch: &DefaultConditionalBranchProperty{
//   						NextStep: &DialogStateProperty{
//   							DialogAction: &DialogActionProperty{
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								SlotToElicit: jsii.String("slotToElicit"),
//   								SuppressNextMessage: jsii.Boolean(false),
//   							},
//   							Intent: &IntentOverrideProperty{
//   								Name: jsii.String("name"),
//   								Slots: []interface{}{
//   									&SlotValueOverrideMapProperty{
//   										SlotName: jsii.String("slotName"),
//   										SlotValueOverride: &slotValueOverrideProperty{
//   											Shape: jsii.String("shape"),
//   											Value: &SlotValueProperty{
//   												InterpretedValue: jsii.String("interpretedValue"),
//   											},
//   											Values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							SessionAttributes: []interface{}{
//   								&SessionAttributeProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						Response: &ResponseSpecificationProperty{
//   							MessageGroupsList: []interface{}{
//   								&MessageGroupProperty{
//   									Message: &MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Variations: []interface{}{
//   										&MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							AllowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   					IsActive: jsii.Boolean(false),
//   				},
//   				TimeoutNextStep: &DialogStateProperty{
//   					DialogAction: &DialogActionProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						SlotToElicit: jsii.String("slotToElicit"),
//   						SuppressNextMessage: jsii.Boolean(false),
//   					},
//   					Intent: &IntentOverrideProperty{
//   						Name: jsii.String("name"),
//   						Slots: []interface{}{
//   							&SlotValueOverrideMapProperty{
//   								SlotName: jsii.String("slotName"),
//   								SlotValueOverride: &slotValueOverrideProperty{
//   									Shape: jsii.String("shape"),
//   									Value: &SlotValueProperty{
//   										InterpretedValue: jsii.String("interpretedValue"),
//   									},
//   									Values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SessionAttributes: []interface{}{
//   						&SessionAttributeProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				TimeoutResponse: &ResponseSpecificationProperty{
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   				},
//   			},
//
//   			// the properties below are optional
//   			InvocationLabel: jsii.String("invocationLabel"),
//   		},
//   		Conditional: &ConditionalSpecificationProperty{
//   			ConditionalBranches: []interface{}{
//   				&ConditionalBranchProperty{
//   					Condition: &ConditionProperty{
//   						ExpressionString: jsii.String("expressionString"),
//   					},
//   					Name: jsii.String("name"),
//   					NextStep: &DialogStateProperty{
//   						DialogAction: &DialogActionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							SlotToElicit: jsii.String("slotToElicit"),
//   							SuppressNextMessage: jsii.Boolean(false),
//   						},
//   						Intent: &IntentOverrideProperty{
//   							Name: jsii.String("name"),
//   							Slots: []interface{}{
//   								&SlotValueOverrideMapProperty{
//   									SlotName: jsii.String("slotName"),
//   									SlotValueOverride: &slotValueOverrideProperty{
//   										Shape: jsii.String("shape"),
//   										Value: &SlotValueProperty{
//   											InterpretedValue: jsii.String("interpretedValue"),
//   										},
//   										Values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						SessionAttributes: []interface{}{
//   							&SessionAttributeProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					Response: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			DefaultBranch: &DefaultConditionalBranchProperty{
//   				NextStep: &DialogStateProperty{
//   					DialogAction: &DialogActionProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						SlotToElicit: jsii.String("slotToElicit"),
//   						SuppressNextMessage: jsii.Boolean(false),
//   					},
//   					Intent: &IntentOverrideProperty{
//   						Name: jsii.String("name"),
//   						Slots: []interface{}{
//   							&SlotValueOverrideMapProperty{
//   								SlotName: jsii.String("slotName"),
//   								SlotValueOverride: &slotValueOverrideProperty{
//   									Shape: jsii.String("shape"),
//   									Value: &SlotValueProperty{
//   										InterpretedValue: jsii.String("interpretedValue"),
//   									},
//   									Values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SessionAttributes: []interface{}{
//   						&SessionAttributeProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				Response: &ResponseSpecificationProperty{
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   				},
//   			},
//   			IsActive: jsii.Boolean(false),
//   		},
//   		InitialResponse: &ResponseSpecificationProperty{
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
//   						SlotValueOverride: &slotValueOverrideProperty{
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
//   	},
//   	InputContexts: []interface{}{
//   		&InputContextProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	IntentClosingSetting: &IntentClosingSettingProperty{
//   		ClosingResponse: &ResponseSpecificationProperty{
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
//   		Conditional: &ConditionalSpecificationProperty{
//   			ConditionalBranches: []interface{}{
//   				&ConditionalBranchProperty{
//   					Condition: &ConditionProperty{
//   						ExpressionString: jsii.String("expressionString"),
//   					},
//   					Name: jsii.String("name"),
//   					NextStep: &DialogStateProperty{
//   						DialogAction: &DialogActionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							SlotToElicit: jsii.String("slotToElicit"),
//   							SuppressNextMessage: jsii.Boolean(false),
//   						},
//   						Intent: &IntentOverrideProperty{
//   							Name: jsii.String("name"),
//   							Slots: []interface{}{
//   								&SlotValueOverrideMapProperty{
//   									SlotName: jsii.String("slotName"),
//   									SlotValueOverride: &slotValueOverrideProperty{
//   										Shape: jsii.String("shape"),
//   										Value: &SlotValueProperty{
//   											InterpretedValue: jsii.String("interpretedValue"),
//   										},
//   										Values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						SessionAttributes: []interface{}{
//   							&SessionAttributeProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					Response: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			DefaultBranch: &DefaultConditionalBranchProperty{
//   				NextStep: &DialogStateProperty{
//   					DialogAction: &DialogActionProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						SlotToElicit: jsii.String("slotToElicit"),
//   						SuppressNextMessage: jsii.Boolean(false),
//   					},
//   					Intent: &IntentOverrideProperty{
//   						Name: jsii.String("name"),
//   						Slots: []interface{}{
//   							&SlotValueOverrideMapProperty{
//   								SlotName: jsii.String("slotName"),
//   								SlotValueOverride: &slotValueOverrideProperty{
//   									Shape: jsii.String("shape"),
//   									Value: &SlotValueProperty{
//   										InterpretedValue: jsii.String("interpretedValue"),
//   									},
//   									Values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SessionAttributes: []interface{}{
//   						&SessionAttributeProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				Response: &ResponseSpecificationProperty{
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   				},
//   			},
//   			IsActive: jsii.Boolean(false),
//   		},
//   		IsActive: jsii.Boolean(false),
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
//   						SlotValueOverride: &slotValueOverrideProperty{
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
//   	},
//   	IntentConfirmationSetting: &IntentConfirmationSettingProperty{
//   		PromptSpecification: &PromptSpecificationProperty{
//   			MaxRetries: jsii.Number(123),
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
//   			MessageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   			PromptAttemptsSpecification: map[string]interface{}{
//   				"promptAttemptsSpecificationKey": &PromptAttemptSpecificationProperty{
//   					"allowedInputTypes": &AllowedInputTypesProperty{
//   						"allowAudioInput": jsii.Boolean(false),
//   						"allowDtmfInput": jsii.Boolean(false),
//   					},
//
//   					// the properties below are optional
//   					"allowInterrupt": jsii.Boolean(false),
//   					"audioAndDtmfInputSpecification": &AudioAndDTMFInputSpecificationProperty{
//   						"startTimeoutMs": jsii.Number(123),
//
//   						// the properties below are optional
//   						"audioSpecification": &AudioSpecificationProperty{
//   							"endTimeoutMs": jsii.Number(123),
//   							"maxLengthMs": jsii.Number(123),
//   						},
//   						"dtmfSpecification": &DTMFSpecificationProperty{
//   							"deletionCharacter": jsii.String("deletionCharacter"),
//   							"endCharacter": jsii.String("endCharacter"),
//   							"endTimeoutMs": jsii.Number(123),
//   							"maxLength": jsii.Number(123),
//   						},
//   					},
//   					"textInputSpecification": &TextInputSpecificationProperty{
//   						"startTimeoutMs": jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		CodeHook: &DialogCodeHookInvocationSettingProperty{
//   			EnableCodeHookInvocation: jsii.Boolean(false),
//   			IsActive: jsii.Boolean(false),
//   			PostCodeHookSpecification: &PostDialogCodeHookInvocationSpecificationProperty{
//   				FailureConditional: &ConditionalSpecificationProperty{
//   					ConditionalBranches: []interface{}{
//   						&ConditionalBranchProperty{
//   							Condition: &ConditionProperty{
//   								ExpressionString: jsii.String("expressionString"),
//   							},
//   							Name: jsii.String("name"),
//   							NextStep: &DialogStateProperty{
//   								DialogAction: &DialogActionProperty{
//   									Type: jsii.String("type"),
//
//   									// the properties below are optional
//   									SlotToElicit: jsii.String("slotToElicit"),
//   									SuppressNextMessage: jsii.Boolean(false),
//   								},
//   								Intent: &IntentOverrideProperty{
//   									Name: jsii.String("name"),
//   									Slots: []interface{}{
//   										&SlotValueOverrideMapProperty{
//   											SlotName: jsii.String("slotName"),
//   											SlotValueOverride: &slotValueOverrideProperty{
//   												Shape: jsii.String("shape"),
//   												Value: &SlotValueProperty{
//   													InterpretedValue: jsii.String("interpretedValue"),
//   												},
//   												Values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								SessionAttributes: []interface{}{
//   									&SessionAttributeProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							Response: &ResponseSpecificationProperty{
//   								MessageGroupsList: []interface{}{
//   									&MessageGroupProperty{
//   										Message: &MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										Variations: []interface{}{
//   											&MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					DefaultBranch: &DefaultConditionalBranchProperty{
//   						NextStep: &DialogStateProperty{
//   							DialogAction: &DialogActionProperty{
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								SlotToElicit: jsii.String("slotToElicit"),
//   								SuppressNextMessage: jsii.Boolean(false),
//   							},
//   							Intent: &IntentOverrideProperty{
//   								Name: jsii.String("name"),
//   								Slots: []interface{}{
//   									&SlotValueOverrideMapProperty{
//   										SlotName: jsii.String("slotName"),
//   										SlotValueOverride: &slotValueOverrideProperty{
//   											Shape: jsii.String("shape"),
//   											Value: &SlotValueProperty{
//   												InterpretedValue: jsii.String("interpretedValue"),
//   											},
//   											Values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							SessionAttributes: []interface{}{
//   								&SessionAttributeProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						Response: &ResponseSpecificationProperty{
//   							MessageGroupsList: []interface{}{
//   								&MessageGroupProperty{
//   									Message: &MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Variations: []interface{}{
//   										&MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							AllowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   					IsActive: jsii.Boolean(false),
//   				},
//   				FailureNextStep: &DialogStateProperty{
//   					DialogAction: &DialogActionProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						SlotToElicit: jsii.String("slotToElicit"),
//   						SuppressNextMessage: jsii.Boolean(false),
//   					},
//   					Intent: &IntentOverrideProperty{
//   						Name: jsii.String("name"),
//   						Slots: []interface{}{
//   							&SlotValueOverrideMapProperty{
//   								SlotName: jsii.String("slotName"),
//   								SlotValueOverride: &slotValueOverrideProperty{
//   									Shape: jsii.String("shape"),
//   									Value: &SlotValueProperty{
//   										InterpretedValue: jsii.String("interpretedValue"),
//   									},
//   									Values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SessionAttributes: []interface{}{
//   						&SessionAttributeProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				FailureResponse: &ResponseSpecificationProperty{
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   				},
//   				SuccessConditional: &ConditionalSpecificationProperty{
//   					ConditionalBranches: []interface{}{
//   						&ConditionalBranchProperty{
//   							Condition: &ConditionProperty{
//   								ExpressionString: jsii.String("expressionString"),
//   							},
//   							Name: jsii.String("name"),
//   							NextStep: &DialogStateProperty{
//   								DialogAction: &DialogActionProperty{
//   									Type: jsii.String("type"),
//
//   									// the properties below are optional
//   									SlotToElicit: jsii.String("slotToElicit"),
//   									SuppressNextMessage: jsii.Boolean(false),
//   								},
//   								Intent: &IntentOverrideProperty{
//   									Name: jsii.String("name"),
//   									Slots: []interface{}{
//   										&SlotValueOverrideMapProperty{
//   											SlotName: jsii.String("slotName"),
//   											SlotValueOverride: &slotValueOverrideProperty{
//   												Shape: jsii.String("shape"),
//   												Value: &SlotValueProperty{
//   													InterpretedValue: jsii.String("interpretedValue"),
//   												},
//   												Values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								SessionAttributes: []interface{}{
//   									&SessionAttributeProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							Response: &ResponseSpecificationProperty{
//   								MessageGroupsList: []interface{}{
//   									&MessageGroupProperty{
//   										Message: &MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										Variations: []interface{}{
//   											&MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					DefaultBranch: &DefaultConditionalBranchProperty{
//   						NextStep: &DialogStateProperty{
//   							DialogAction: &DialogActionProperty{
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								SlotToElicit: jsii.String("slotToElicit"),
//   								SuppressNextMessage: jsii.Boolean(false),
//   							},
//   							Intent: &IntentOverrideProperty{
//   								Name: jsii.String("name"),
//   								Slots: []interface{}{
//   									&SlotValueOverrideMapProperty{
//   										SlotName: jsii.String("slotName"),
//   										SlotValueOverride: &slotValueOverrideProperty{
//   											Shape: jsii.String("shape"),
//   											Value: &SlotValueProperty{
//   												InterpretedValue: jsii.String("interpretedValue"),
//   											},
//   											Values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							SessionAttributes: []interface{}{
//   								&SessionAttributeProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						Response: &ResponseSpecificationProperty{
//   							MessageGroupsList: []interface{}{
//   								&MessageGroupProperty{
//   									Message: &MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Variations: []interface{}{
//   										&MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							AllowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   					IsActive: jsii.Boolean(false),
//   				},
//   				SuccessNextStep: &DialogStateProperty{
//   					DialogAction: &DialogActionProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						SlotToElicit: jsii.String("slotToElicit"),
//   						SuppressNextMessage: jsii.Boolean(false),
//   					},
//   					Intent: &IntentOverrideProperty{
//   						Name: jsii.String("name"),
//   						Slots: []interface{}{
//   							&SlotValueOverrideMapProperty{
//   								SlotName: jsii.String("slotName"),
//   								SlotValueOverride: &slotValueOverrideProperty{
//   									Shape: jsii.String("shape"),
//   									Value: &SlotValueProperty{
//   										InterpretedValue: jsii.String("interpretedValue"),
//   									},
//   									Values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SessionAttributes: []interface{}{
//   						&SessionAttributeProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				SuccessResponse: &ResponseSpecificationProperty{
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   				},
//   				TimeoutConditional: &ConditionalSpecificationProperty{
//   					ConditionalBranches: []interface{}{
//   						&ConditionalBranchProperty{
//   							Condition: &ConditionProperty{
//   								ExpressionString: jsii.String("expressionString"),
//   							},
//   							Name: jsii.String("name"),
//   							NextStep: &DialogStateProperty{
//   								DialogAction: &DialogActionProperty{
//   									Type: jsii.String("type"),
//
//   									// the properties below are optional
//   									SlotToElicit: jsii.String("slotToElicit"),
//   									SuppressNextMessage: jsii.Boolean(false),
//   								},
//   								Intent: &IntentOverrideProperty{
//   									Name: jsii.String("name"),
//   									Slots: []interface{}{
//   										&SlotValueOverrideMapProperty{
//   											SlotName: jsii.String("slotName"),
//   											SlotValueOverride: &slotValueOverrideProperty{
//   												Shape: jsii.String("shape"),
//   												Value: &SlotValueProperty{
//   													InterpretedValue: jsii.String("interpretedValue"),
//   												},
//   												Values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								SessionAttributes: []interface{}{
//   									&SessionAttributeProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							Response: &ResponseSpecificationProperty{
//   								MessageGroupsList: []interface{}{
//   									&MessageGroupProperty{
//   										Message: &MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										Variations: []interface{}{
//   											&MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					DefaultBranch: &DefaultConditionalBranchProperty{
//   						NextStep: &DialogStateProperty{
//   							DialogAction: &DialogActionProperty{
//   								Type: jsii.String("type"),
//
//   								// the properties below are optional
//   								SlotToElicit: jsii.String("slotToElicit"),
//   								SuppressNextMessage: jsii.Boolean(false),
//   							},
//   							Intent: &IntentOverrideProperty{
//   								Name: jsii.String("name"),
//   								Slots: []interface{}{
//   									&SlotValueOverrideMapProperty{
//   										SlotName: jsii.String("slotName"),
//   										SlotValueOverride: &slotValueOverrideProperty{
//   											Shape: jsii.String("shape"),
//   											Value: &SlotValueProperty{
//   												InterpretedValue: jsii.String("interpretedValue"),
//   											},
//   											Values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							SessionAttributes: []interface{}{
//   								&SessionAttributeProperty{
//   									Key: jsii.String("key"),
//
//   									// the properties below are optional
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						Response: &ResponseSpecificationProperty{
//   							MessageGroupsList: []interface{}{
//   								&MessageGroupProperty{
//   									Message: &MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									Variations: []interface{}{
//   										&MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							AllowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   					IsActive: jsii.Boolean(false),
//   				},
//   				TimeoutNextStep: &DialogStateProperty{
//   					DialogAction: &DialogActionProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						SlotToElicit: jsii.String("slotToElicit"),
//   						SuppressNextMessage: jsii.Boolean(false),
//   					},
//   					Intent: &IntentOverrideProperty{
//   						Name: jsii.String("name"),
//   						Slots: []interface{}{
//   							&SlotValueOverrideMapProperty{
//   								SlotName: jsii.String("slotName"),
//   								SlotValueOverride: &slotValueOverrideProperty{
//   									Shape: jsii.String("shape"),
//   									Value: &SlotValueProperty{
//   										InterpretedValue: jsii.String("interpretedValue"),
//   									},
//   									Values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SessionAttributes: []interface{}{
//   						&SessionAttributeProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				TimeoutResponse: &ResponseSpecificationProperty{
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   				},
//   			},
//
//   			// the properties below are optional
//   			InvocationLabel: jsii.String("invocationLabel"),
//   		},
//   		ConfirmationConditional: &ConditionalSpecificationProperty{
//   			ConditionalBranches: []interface{}{
//   				&ConditionalBranchProperty{
//   					Condition: &ConditionProperty{
//   						ExpressionString: jsii.String("expressionString"),
//   					},
//   					Name: jsii.String("name"),
//   					NextStep: &DialogStateProperty{
//   						DialogAction: &DialogActionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							SlotToElicit: jsii.String("slotToElicit"),
//   							SuppressNextMessage: jsii.Boolean(false),
//   						},
//   						Intent: &IntentOverrideProperty{
//   							Name: jsii.String("name"),
//   							Slots: []interface{}{
//   								&SlotValueOverrideMapProperty{
//   									SlotName: jsii.String("slotName"),
//   									SlotValueOverride: &slotValueOverrideProperty{
//   										Shape: jsii.String("shape"),
//   										Value: &SlotValueProperty{
//   											InterpretedValue: jsii.String("interpretedValue"),
//   										},
//   										Values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						SessionAttributes: []interface{}{
//   							&SessionAttributeProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					Response: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			DefaultBranch: &DefaultConditionalBranchProperty{
//   				NextStep: &DialogStateProperty{
//   					DialogAction: &DialogActionProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						SlotToElicit: jsii.String("slotToElicit"),
//   						SuppressNextMessage: jsii.Boolean(false),
//   					},
//   					Intent: &IntentOverrideProperty{
//   						Name: jsii.String("name"),
//   						Slots: []interface{}{
//   							&SlotValueOverrideMapProperty{
//   								SlotName: jsii.String("slotName"),
//   								SlotValueOverride: &slotValueOverrideProperty{
//   									Shape: jsii.String("shape"),
//   									Value: &SlotValueProperty{
//   										InterpretedValue: jsii.String("interpretedValue"),
//   									},
//   									Values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SessionAttributes: []interface{}{
//   						&SessionAttributeProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				Response: &ResponseSpecificationProperty{
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   				},
//   			},
//   			IsActive: jsii.Boolean(false),
//   		},
//   		ConfirmationNextStep: &DialogStateProperty{
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
//   						SlotValueOverride: &slotValueOverrideProperty{
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
//   		ConfirmationResponse: &ResponseSpecificationProperty{
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
//   		DeclinationConditional: &ConditionalSpecificationProperty{
//   			ConditionalBranches: []interface{}{
//   				&ConditionalBranchProperty{
//   					Condition: &ConditionProperty{
//   						ExpressionString: jsii.String("expressionString"),
//   					},
//   					Name: jsii.String("name"),
//   					NextStep: &DialogStateProperty{
//   						DialogAction: &DialogActionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							SlotToElicit: jsii.String("slotToElicit"),
//   							SuppressNextMessage: jsii.Boolean(false),
//   						},
//   						Intent: &IntentOverrideProperty{
//   							Name: jsii.String("name"),
//   							Slots: []interface{}{
//   								&SlotValueOverrideMapProperty{
//   									SlotName: jsii.String("slotName"),
//   									SlotValueOverride: &slotValueOverrideProperty{
//   										Shape: jsii.String("shape"),
//   										Value: &SlotValueProperty{
//   											InterpretedValue: jsii.String("interpretedValue"),
//   										},
//   										Values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						SessionAttributes: []interface{}{
//   							&SessionAttributeProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					Response: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			DefaultBranch: &DefaultConditionalBranchProperty{
//   				NextStep: &DialogStateProperty{
//   					DialogAction: &DialogActionProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						SlotToElicit: jsii.String("slotToElicit"),
//   						SuppressNextMessage: jsii.Boolean(false),
//   					},
//   					Intent: &IntentOverrideProperty{
//   						Name: jsii.String("name"),
//   						Slots: []interface{}{
//   							&SlotValueOverrideMapProperty{
//   								SlotName: jsii.String("slotName"),
//   								SlotValueOverride: &slotValueOverrideProperty{
//   									Shape: jsii.String("shape"),
//   									Value: &SlotValueProperty{
//   										InterpretedValue: jsii.String("interpretedValue"),
//   									},
//   									Values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SessionAttributes: []interface{}{
//   						&SessionAttributeProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				Response: &ResponseSpecificationProperty{
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   				},
//   			},
//   			IsActive: jsii.Boolean(false),
//   		},
//   		DeclinationNextStep: &DialogStateProperty{
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
//   						SlotValueOverride: &slotValueOverrideProperty{
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
//   		DeclinationResponse: &ResponseSpecificationProperty{
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
//   		ElicitationCodeHook: &ElicitationCodeHookInvocationSettingProperty{
//   			EnableCodeHookInvocation: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			InvocationLabel: jsii.String("invocationLabel"),
//   		},
//   		FailureConditional: &ConditionalSpecificationProperty{
//   			ConditionalBranches: []interface{}{
//   				&ConditionalBranchProperty{
//   					Condition: &ConditionProperty{
//   						ExpressionString: jsii.String("expressionString"),
//   					},
//   					Name: jsii.String("name"),
//   					NextStep: &DialogStateProperty{
//   						DialogAction: &DialogActionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							SlotToElicit: jsii.String("slotToElicit"),
//   							SuppressNextMessage: jsii.Boolean(false),
//   						},
//   						Intent: &IntentOverrideProperty{
//   							Name: jsii.String("name"),
//   							Slots: []interface{}{
//   								&SlotValueOverrideMapProperty{
//   									SlotName: jsii.String("slotName"),
//   									SlotValueOverride: &slotValueOverrideProperty{
//   										Shape: jsii.String("shape"),
//   										Value: &SlotValueProperty{
//   											InterpretedValue: jsii.String("interpretedValue"),
//   										},
//   										Values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						SessionAttributes: []interface{}{
//   							&SessionAttributeProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					Response: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			DefaultBranch: &DefaultConditionalBranchProperty{
//   				NextStep: &DialogStateProperty{
//   					DialogAction: &DialogActionProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						SlotToElicit: jsii.String("slotToElicit"),
//   						SuppressNextMessage: jsii.Boolean(false),
//   					},
//   					Intent: &IntentOverrideProperty{
//   						Name: jsii.String("name"),
//   						Slots: []interface{}{
//   							&SlotValueOverrideMapProperty{
//   								SlotName: jsii.String("slotName"),
//   								SlotValueOverride: &slotValueOverrideProperty{
//   									Shape: jsii.String("shape"),
//   									Value: &SlotValueProperty{
//   										InterpretedValue: jsii.String("interpretedValue"),
//   									},
//   									Values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SessionAttributes: []interface{}{
//   						&SessionAttributeProperty{
//   							Key: jsii.String("key"),
//
//   							// the properties below are optional
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				Response: &ResponseSpecificationProperty{
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   				},
//   			},
//   			IsActive: jsii.Boolean(false),
//   		},
//   		FailureNextStep: &DialogStateProperty{
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
//   						SlotValueOverride: &slotValueOverrideProperty{
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
//   		FailureResponse: &ResponseSpecificationProperty{
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
//   		IsActive: jsii.Boolean(false),
//   	},
//   	KendraConfiguration: &KendraConfigurationProperty{
//   		KendraIndex: jsii.String("kendraIndex"),
//
//   		// the properties below are optional
//   		QueryFilterString: jsii.String("queryFilterString"),
//   		QueryFilterStringEnabled: jsii.Boolean(false),
//   	},
//   	OutputContexts: []interface{}{
//   		&OutputContextProperty{
//   			Name: jsii.String("name"),
//   			TimeToLiveInSeconds: jsii.Number(123),
//   			TurnsToLive: jsii.Number(123),
//   		},
//   	},
//   	ParentIntentSignature: jsii.String("parentIntentSignature"),
//   	SampleUtterances: []interface{}{
//   		&SampleUtteranceProperty{
//   			Utterance: jsii.String("utterance"),
//   		},
//   	},
//   	SlotPriorities: []interface{}{
//   		&SlotPriorityProperty{
//   			Priority: jsii.Number(123),
//   			SlotName: jsii.String("slotName"),
//   		},
//   	},
//   	Slots: []interface{}{
//   		&SlotProperty{
//   			Name: jsii.String("name"),
//   			SlotTypeName: jsii.String("slotTypeName"),
//   			ValueElicitationSetting: &SlotValueElicitationSettingProperty{
//   				SlotConstraint: jsii.String("slotConstraint"),
//
//   				// the properties below are optional
//   				DefaultValueSpecification: &SlotDefaultValueSpecificationProperty{
//   					DefaultValueList: []interface{}{
//   						&SlotDefaultValueProperty{
//   							DefaultValue: jsii.String("defaultValue"),
//   						},
//   					},
//   				},
//   				PromptSpecification: &PromptSpecificationProperty{
//   					MaxRetries: jsii.Number(123),
//   					MessageGroupsList: []interface{}{
//   						&MessageGroupProperty{
//   							Message: &MessageProperty{
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
//
//   							// the properties below are optional
//   							Variations: []interface{}{
//   								&MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					AllowInterrupt: jsii.Boolean(false),
//   					MessageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   					PromptAttemptsSpecification: map[string]interface{}{
//   						"promptAttemptsSpecificationKey": &PromptAttemptSpecificationProperty{
//   							"allowedInputTypes": &AllowedInputTypesProperty{
//   								"allowAudioInput": jsii.Boolean(false),
//   								"allowDtmfInput": jsii.Boolean(false),
//   							},
//
//   							// the properties below are optional
//   							"allowInterrupt": jsii.Boolean(false),
//   							"audioAndDtmfInputSpecification": &AudioAndDTMFInputSpecificationProperty{
//   								"startTimeoutMs": jsii.Number(123),
//
//   								// the properties below are optional
//   								"audioSpecification": &AudioSpecificationProperty{
//   									"endTimeoutMs": jsii.Number(123),
//   									"maxLengthMs": jsii.Number(123),
//   								},
//   								"dtmfSpecification": &DTMFSpecificationProperty{
//   									"deletionCharacter": jsii.String("deletionCharacter"),
//   									"endCharacter": jsii.String("endCharacter"),
//   									"endTimeoutMs": jsii.Number(123),
//   									"maxLength": jsii.Number(123),
//   								},
//   							},
//   							"textInputSpecification": &TextInputSpecificationProperty{
//   								"startTimeoutMs": jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   				SampleUtterances: []interface{}{
//   					&SampleUtteranceProperty{
//   						Utterance: jsii.String("utterance"),
//   					},
//   				},
//   				SlotCaptureSetting: &SlotCaptureSettingProperty{
//   					CaptureConditional: &ConditionalSpecificationProperty{
//   						ConditionalBranches: []interface{}{
//   							&ConditionalBranchProperty{
//   								Condition: &ConditionProperty{
//   									ExpressionString: jsii.String("expressionString"),
//   								},
//   								Name: jsii.String("name"),
//   								NextStep: &DialogStateProperty{
//   									DialogAction: &DialogActionProperty{
//   										Type: jsii.String("type"),
//
//   										// the properties below are optional
//   										SlotToElicit: jsii.String("slotToElicit"),
//   										SuppressNextMessage: jsii.Boolean(false),
//   									},
//   									Intent: &IntentOverrideProperty{
//   										Name: jsii.String("name"),
//   										Slots: []interface{}{
//   											&SlotValueOverrideMapProperty{
//   												SlotName: jsii.String("slotName"),
//   												SlotValueOverride: &slotValueOverrideProperty{
//   													Shape: jsii.String("shape"),
//   													Value: &SlotValueProperty{
//   														InterpretedValue: jsii.String("interpretedValue"),
//   													},
//   													Values: []interface{}{
//   														slotValueOverrideProperty_,
//   													},
//   												},
//   											},
//   										},
//   									},
//   									SessionAttributes: []interface{}{
//   										&SessionAttributeProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								Response: &ResponseSpecificationProperty{
//   									MessageGroupsList: []interface{}{
//   										&MessageGroupProperty{
//   											Message: &MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//
//   											// the properties below are optional
//   											Variations: []interface{}{
//   												&MessageProperty{
//   													CustomPayload: &CustomPayloadProperty{
//   														Value: jsii.String("value"),
//   													},
//   													ImageResponseCard: &ImageResponseCardProperty{
//   														Title: jsii.String("title"),
//
//   														// the properties below are optional
//   														Buttons: []interface{}{
//   															&ButtonProperty{
//   																Text: jsii.String("text"),
//   																Value: jsii.String("value"),
//   															},
//   														},
//   														ImageUrl: jsii.String("imageUrl"),
//   														Subtitle: jsii.String("subtitle"),
//   													},
//   													PlainTextMessage: &PlainTextMessageProperty{
//   														Value: jsii.String("value"),
//   													},
//   													SsmlMessage: &SSMLMessageProperty{
//   														Value: jsii.String("value"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									AllowInterrupt: jsii.Boolean(false),
//   								},
//   							},
//   						},
//   						DefaultBranch: &DefaultConditionalBranchProperty{
//   							NextStep: &DialogStateProperty{
//   								DialogAction: &DialogActionProperty{
//   									Type: jsii.String("type"),
//
//   									// the properties below are optional
//   									SlotToElicit: jsii.String("slotToElicit"),
//   									SuppressNextMessage: jsii.Boolean(false),
//   								},
//   								Intent: &IntentOverrideProperty{
//   									Name: jsii.String("name"),
//   									Slots: []interface{}{
//   										&SlotValueOverrideMapProperty{
//   											SlotName: jsii.String("slotName"),
//   											SlotValueOverride: &slotValueOverrideProperty{
//   												Shape: jsii.String("shape"),
//   												Value: &SlotValueProperty{
//   													InterpretedValue: jsii.String("interpretedValue"),
//   												},
//   												Values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								SessionAttributes: []interface{}{
//   									&SessionAttributeProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							Response: &ResponseSpecificationProperty{
//   								MessageGroupsList: []interface{}{
//   									&MessageGroupProperty{
//   										Message: &MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										Variations: []interface{}{
//   											&MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   						IsActive: jsii.Boolean(false),
//   					},
//   					CaptureNextStep: &DialogStateProperty{
//   						DialogAction: &DialogActionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							SlotToElicit: jsii.String("slotToElicit"),
//   							SuppressNextMessage: jsii.Boolean(false),
//   						},
//   						Intent: &IntentOverrideProperty{
//   							Name: jsii.String("name"),
//   							Slots: []interface{}{
//   								&SlotValueOverrideMapProperty{
//   									SlotName: jsii.String("slotName"),
//   									SlotValueOverride: &slotValueOverrideProperty{
//   										Shape: jsii.String("shape"),
//   										Value: &SlotValueProperty{
//   											InterpretedValue: jsii.String("interpretedValue"),
//   										},
//   										Values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						SessionAttributes: []interface{}{
//   							&SessionAttributeProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					CaptureResponse: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   					CodeHook: &DialogCodeHookInvocationSettingProperty{
//   						EnableCodeHookInvocation: jsii.Boolean(false),
//   						IsActive: jsii.Boolean(false),
//   						PostCodeHookSpecification: &PostDialogCodeHookInvocationSpecificationProperty{
//   							FailureConditional: &ConditionalSpecificationProperty{
//   								ConditionalBranches: []interface{}{
//   									&ConditionalBranchProperty{
//   										Condition: &ConditionProperty{
//   											ExpressionString: jsii.String("expressionString"),
//   										},
//   										Name: jsii.String("name"),
//   										NextStep: &DialogStateProperty{
//   											DialogAction: &DialogActionProperty{
//   												Type: jsii.String("type"),
//
//   												// the properties below are optional
//   												SlotToElicit: jsii.String("slotToElicit"),
//   												SuppressNextMessage: jsii.Boolean(false),
//   											},
//   											Intent: &IntentOverrideProperty{
//   												Name: jsii.String("name"),
//   												Slots: []interface{}{
//   													&SlotValueOverrideMapProperty{
//   														SlotName: jsii.String("slotName"),
//   														SlotValueOverride: &slotValueOverrideProperty{
//   															Shape: jsii.String("shape"),
//   															Value: &SlotValueProperty{
//   																InterpretedValue: jsii.String("interpretedValue"),
//   															},
//   															Values: []interface{}{
//   																slotValueOverrideProperty_,
//   															},
//   														},
//   													},
//   												},
//   											},
//   											SessionAttributes: []interface{}{
//   												&SessionAttributeProperty{
//   													Key: jsii.String("key"),
//
//   													// the properties below are optional
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										Response: &ResponseSpecificationProperty{
//   											MessageGroupsList: []interface{}{
//   												&MessageGroupProperty{
//   													Message: &MessageProperty{
//   														CustomPayload: &CustomPayloadProperty{
//   															Value: jsii.String("value"),
//   														},
//   														ImageResponseCard: &ImageResponseCardProperty{
//   															Title: jsii.String("title"),
//
//   															// the properties below are optional
//   															Buttons: []interface{}{
//   																&ButtonProperty{
//   																	Text: jsii.String("text"),
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   															ImageUrl: jsii.String("imageUrl"),
//   															Subtitle: jsii.String("subtitle"),
//   														},
//   														PlainTextMessage: &PlainTextMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   														SsmlMessage: &SSMLMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   													},
//
//   													// the properties below are optional
//   													Variations: []interface{}{
//   														&MessageProperty{
//   															CustomPayload: &CustomPayloadProperty{
//   																Value: jsii.String("value"),
//   															},
//   															ImageResponseCard: &ImageResponseCardProperty{
//   																Title: jsii.String("title"),
//
//   																// the properties below are optional
//   																Buttons: []interface{}{
//   																	&ButtonProperty{
//   																		Text: jsii.String("text"),
//   																		Value: jsii.String("value"),
//   																	},
//   																},
//   																ImageUrl: jsii.String("imageUrl"),
//   																Subtitle: jsii.String("subtitle"),
//   															},
//   															PlainTextMessage: &PlainTextMessageProperty{
//   																Value: jsii.String("value"),
//   															},
//   															SsmlMessage: &SSMLMessageProperty{
//   																Value: jsii.String("value"),
//   															},
//   														},
//   													},
//   												},
//   											},
//
//   											// the properties below are optional
//   											AllowInterrupt: jsii.Boolean(false),
//   										},
//   									},
//   								},
//   								DefaultBranch: &DefaultConditionalBranchProperty{
//   									NextStep: &DialogStateProperty{
//   										DialogAction: &DialogActionProperty{
//   											Type: jsii.String("type"),
//
//   											// the properties below are optional
//   											SlotToElicit: jsii.String("slotToElicit"),
//   											SuppressNextMessage: jsii.Boolean(false),
//   										},
//   										Intent: &IntentOverrideProperty{
//   											Name: jsii.String("name"),
//   											Slots: []interface{}{
//   												&SlotValueOverrideMapProperty{
//   													SlotName: jsii.String("slotName"),
//   													SlotValueOverride: &slotValueOverrideProperty{
//   														Shape: jsii.String("shape"),
//   														Value: &SlotValueProperty{
//   															InterpretedValue: jsii.String("interpretedValue"),
//   														},
//   														Values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										SessionAttributes: []interface{}{
//   											&SessionAttributeProperty{
//   												Key: jsii.String("key"),
//
//   												// the properties below are optional
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   									Response: &ResponseSpecificationProperty{
//   										MessageGroupsList: []interface{}{
//   											&MessageGroupProperty{
//   												Message: &MessageProperty{
//   													CustomPayload: &CustomPayloadProperty{
//   														Value: jsii.String("value"),
//   													},
//   													ImageResponseCard: &ImageResponseCardProperty{
//   														Title: jsii.String("title"),
//
//   														// the properties below are optional
//   														Buttons: []interface{}{
//   															&ButtonProperty{
//   																Text: jsii.String("text"),
//   																Value: jsii.String("value"),
//   															},
//   														},
//   														ImageUrl: jsii.String("imageUrl"),
//   														Subtitle: jsii.String("subtitle"),
//   													},
//   													PlainTextMessage: &PlainTextMessageProperty{
//   														Value: jsii.String("value"),
//   													},
//   													SsmlMessage: &SSMLMessageProperty{
//   														Value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												Variations: []interface{}{
//   													&MessageProperty{
//   														CustomPayload: &CustomPayloadProperty{
//   															Value: jsii.String("value"),
//   														},
//   														ImageResponseCard: &ImageResponseCardProperty{
//   															Title: jsii.String("title"),
//
//   															// the properties below are optional
//   															Buttons: []interface{}{
//   																&ButtonProperty{
//   																	Text: jsii.String("text"),
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   															ImageUrl: jsii.String("imageUrl"),
//   															Subtitle: jsii.String("subtitle"),
//   														},
//   														PlainTextMessage: &PlainTextMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   														SsmlMessage: &SSMLMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										AllowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   								IsActive: jsii.Boolean(false),
//   							},
//   							FailureNextStep: &DialogStateProperty{
//   								DialogAction: &DialogActionProperty{
//   									Type: jsii.String("type"),
//
//   									// the properties below are optional
//   									SlotToElicit: jsii.String("slotToElicit"),
//   									SuppressNextMessage: jsii.Boolean(false),
//   								},
//   								Intent: &IntentOverrideProperty{
//   									Name: jsii.String("name"),
//   									Slots: []interface{}{
//   										&SlotValueOverrideMapProperty{
//   											SlotName: jsii.String("slotName"),
//   											SlotValueOverride: &slotValueOverrideProperty{
//   												Shape: jsii.String("shape"),
//   												Value: &SlotValueProperty{
//   													InterpretedValue: jsii.String("interpretedValue"),
//   												},
//   												Values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								SessionAttributes: []interface{}{
//   									&SessionAttributeProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							FailureResponse: &ResponseSpecificationProperty{
//   								MessageGroupsList: []interface{}{
//   									&MessageGroupProperty{
//   										Message: &MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										Variations: []interface{}{
//   											&MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   							SuccessConditional: &ConditionalSpecificationProperty{
//   								ConditionalBranches: []interface{}{
//   									&ConditionalBranchProperty{
//   										Condition: &ConditionProperty{
//   											ExpressionString: jsii.String("expressionString"),
//   										},
//   										Name: jsii.String("name"),
//   										NextStep: &DialogStateProperty{
//   											DialogAction: &DialogActionProperty{
//   												Type: jsii.String("type"),
//
//   												// the properties below are optional
//   												SlotToElicit: jsii.String("slotToElicit"),
//   												SuppressNextMessage: jsii.Boolean(false),
//   											},
//   											Intent: &IntentOverrideProperty{
//   												Name: jsii.String("name"),
//   												Slots: []interface{}{
//   													&SlotValueOverrideMapProperty{
//   														SlotName: jsii.String("slotName"),
//   														SlotValueOverride: &slotValueOverrideProperty{
//   															Shape: jsii.String("shape"),
//   															Value: &SlotValueProperty{
//   																InterpretedValue: jsii.String("interpretedValue"),
//   															},
//   															Values: []interface{}{
//   																slotValueOverrideProperty_,
//   															},
//   														},
//   													},
//   												},
//   											},
//   											SessionAttributes: []interface{}{
//   												&SessionAttributeProperty{
//   													Key: jsii.String("key"),
//
//   													// the properties below are optional
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										Response: &ResponseSpecificationProperty{
//   											MessageGroupsList: []interface{}{
//   												&MessageGroupProperty{
//   													Message: &MessageProperty{
//   														CustomPayload: &CustomPayloadProperty{
//   															Value: jsii.String("value"),
//   														},
//   														ImageResponseCard: &ImageResponseCardProperty{
//   															Title: jsii.String("title"),
//
//   															// the properties below are optional
//   															Buttons: []interface{}{
//   																&ButtonProperty{
//   																	Text: jsii.String("text"),
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   															ImageUrl: jsii.String("imageUrl"),
//   															Subtitle: jsii.String("subtitle"),
//   														},
//   														PlainTextMessage: &PlainTextMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   														SsmlMessage: &SSMLMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   													},
//
//   													// the properties below are optional
//   													Variations: []interface{}{
//   														&MessageProperty{
//   															CustomPayload: &CustomPayloadProperty{
//   																Value: jsii.String("value"),
//   															},
//   															ImageResponseCard: &ImageResponseCardProperty{
//   																Title: jsii.String("title"),
//
//   																// the properties below are optional
//   																Buttons: []interface{}{
//   																	&ButtonProperty{
//   																		Text: jsii.String("text"),
//   																		Value: jsii.String("value"),
//   																	},
//   																},
//   																ImageUrl: jsii.String("imageUrl"),
//   																Subtitle: jsii.String("subtitle"),
//   															},
//   															PlainTextMessage: &PlainTextMessageProperty{
//   																Value: jsii.String("value"),
//   															},
//   															SsmlMessage: &SSMLMessageProperty{
//   																Value: jsii.String("value"),
//   															},
//   														},
//   													},
//   												},
//   											},
//
//   											// the properties below are optional
//   											AllowInterrupt: jsii.Boolean(false),
//   										},
//   									},
//   								},
//   								DefaultBranch: &DefaultConditionalBranchProperty{
//   									NextStep: &DialogStateProperty{
//   										DialogAction: &DialogActionProperty{
//   											Type: jsii.String("type"),
//
//   											// the properties below are optional
//   											SlotToElicit: jsii.String("slotToElicit"),
//   											SuppressNextMessage: jsii.Boolean(false),
//   										},
//   										Intent: &IntentOverrideProperty{
//   											Name: jsii.String("name"),
//   											Slots: []interface{}{
//   												&SlotValueOverrideMapProperty{
//   													SlotName: jsii.String("slotName"),
//   													SlotValueOverride: &slotValueOverrideProperty{
//   														Shape: jsii.String("shape"),
//   														Value: &SlotValueProperty{
//   															InterpretedValue: jsii.String("interpretedValue"),
//   														},
//   														Values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										SessionAttributes: []interface{}{
//   											&SessionAttributeProperty{
//   												Key: jsii.String("key"),
//
//   												// the properties below are optional
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   									Response: &ResponseSpecificationProperty{
//   										MessageGroupsList: []interface{}{
//   											&MessageGroupProperty{
//   												Message: &MessageProperty{
//   													CustomPayload: &CustomPayloadProperty{
//   														Value: jsii.String("value"),
//   													},
//   													ImageResponseCard: &ImageResponseCardProperty{
//   														Title: jsii.String("title"),
//
//   														// the properties below are optional
//   														Buttons: []interface{}{
//   															&ButtonProperty{
//   																Text: jsii.String("text"),
//   																Value: jsii.String("value"),
//   															},
//   														},
//   														ImageUrl: jsii.String("imageUrl"),
//   														Subtitle: jsii.String("subtitle"),
//   													},
//   													PlainTextMessage: &PlainTextMessageProperty{
//   														Value: jsii.String("value"),
//   													},
//   													SsmlMessage: &SSMLMessageProperty{
//   														Value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												Variations: []interface{}{
//   													&MessageProperty{
//   														CustomPayload: &CustomPayloadProperty{
//   															Value: jsii.String("value"),
//   														},
//   														ImageResponseCard: &ImageResponseCardProperty{
//   															Title: jsii.String("title"),
//
//   															// the properties below are optional
//   															Buttons: []interface{}{
//   																&ButtonProperty{
//   																	Text: jsii.String("text"),
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   															ImageUrl: jsii.String("imageUrl"),
//   															Subtitle: jsii.String("subtitle"),
//   														},
//   														PlainTextMessage: &PlainTextMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   														SsmlMessage: &SSMLMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										AllowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   								IsActive: jsii.Boolean(false),
//   							},
//   							SuccessNextStep: &DialogStateProperty{
//   								DialogAction: &DialogActionProperty{
//   									Type: jsii.String("type"),
//
//   									// the properties below are optional
//   									SlotToElicit: jsii.String("slotToElicit"),
//   									SuppressNextMessage: jsii.Boolean(false),
//   								},
//   								Intent: &IntentOverrideProperty{
//   									Name: jsii.String("name"),
//   									Slots: []interface{}{
//   										&SlotValueOverrideMapProperty{
//   											SlotName: jsii.String("slotName"),
//   											SlotValueOverride: &slotValueOverrideProperty{
//   												Shape: jsii.String("shape"),
//   												Value: &SlotValueProperty{
//   													InterpretedValue: jsii.String("interpretedValue"),
//   												},
//   												Values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								SessionAttributes: []interface{}{
//   									&SessionAttributeProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							SuccessResponse: &ResponseSpecificationProperty{
//   								MessageGroupsList: []interface{}{
//   									&MessageGroupProperty{
//   										Message: &MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										Variations: []interface{}{
//   											&MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   							TimeoutConditional: &ConditionalSpecificationProperty{
//   								ConditionalBranches: []interface{}{
//   									&ConditionalBranchProperty{
//   										Condition: &ConditionProperty{
//   											ExpressionString: jsii.String("expressionString"),
//   										},
//   										Name: jsii.String("name"),
//   										NextStep: &DialogStateProperty{
//   											DialogAction: &DialogActionProperty{
//   												Type: jsii.String("type"),
//
//   												// the properties below are optional
//   												SlotToElicit: jsii.String("slotToElicit"),
//   												SuppressNextMessage: jsii.Boolean(false),
//   											},
//   											Intent: &IntentOverrideProperty{
//   												Name: jsii.String("name"),
//   												Slots: []interface{}{
//   													&SlotValueOverrideMapProperty{
//   														SlotName: jsii.String("slotName"),
//   														SlotValueOverride: &slotValueOverrideProperty{
//   															Shape: jsii.String("shape"),
//   															Value: &SlotValueProperty{
//   																InterpretedValue: jsii.String("interpretedValue"),
//   															},
//   															Values: []interface{}{
//   																slotValueOverrideProperty_,
//   															},
//   														},
//   													},
//   												},
//   											},
//   											SessionAttributes: []interface{}{
//   												&SessionAttributeProperty{
//   													Key: jsii.String("key"),
//
//   													// the properties below are optional
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										Response: &ResponseSpecificationProperty{
//   											MessageGroupsList: []interface{}{
//   												&MessageGroupProperty{
//   													Message: &MessageProperty{
//   														CustomPayload: &CustomPayloadProperty{
//   															Value: jsii.String("value"),
//   														},
//   														ImageResponseCard: &ImageResponseCardProperty{
//   															Title: jsii.String("title"),
//
//   															// the properties below are optional
//   															Buttons: []interface{}{
//   																&ButtonProperty{
//   																	Text: jsii.String("text"),
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   															ImageUrl: jsii.String("imageUrl"),
//   															Subtitle: jsii.String("subtitle"),
//   														},
//   														PlainTextMessage: &PlainTextMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   														SsmlMessage: &SSMLMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   													},
//
//   													// the properties below are optional
//   													Variations: []interface{}{
//   														&MessageProperty{
//   															CustomPayload: &CustomPayloadProperty{
//   																Value: jsii.String("value"),
//   															},
//   															ImageResponseCard: &ImageResponseCardProperty{
//   																Title: jsii.String("title"),
//
//   																// the properties below are optional
//   																Buttons: []interface{}{
//   																	&ButtonProperty{
//   																		Text: jsii.String("text"),
//   																		Value: jsii.String("value"),
//   																	},
//   																},
//   																ImageUrl: jsii.String("imageUrl"),
//   																Subtitle: jsii.String("subtitle"),
//   															},
//   															PlainTextMessage: &PlainTextMessageProperty{
//   																Value: jsii.String("value"),
//   															},
//   															SsmlMessage: &SSMLMessageProperty{
//   																Value: jsii.String("value"),
//   															},
//   														},
//   													},
//   												},
//   											},
//
//   											// the properties below are optional
//   											AllowInterrupt: jsii.Boolean(false),
//   										},
//   									},
//   								},
//   								DefaultBranch: &DefaultConditionalBranchProperty{
//   									NextStep: &DialogStateProperty{
//   										DialogAction: &DialogActionProperty{
//   											Type: jsii.String("type"),
//
//   											// the properties below are optional
//   											SlotToElicit: jsii.String("slotToElicit"),
//   											SuppressNextMessage: jsii.Boolean(false),
//   										},
//   										Intent: &IntentOverrideProperty{
//   											Name: jsii.String("name"),
//   											Slots: []interface{}{
//   												&SlotValueOverrideMapProperty{
//   													SlotName: jsii.String("slotName"),
//   													SlotValueOverride: &slotValueOverrideProperty{
//   														Shape: jsii.String("shape"),
//   														Value: &SlotValueProperty{
//   															InterpretedValue: jsii.String("interpretedValue"),
//   														},
//   														Values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										SessionAttributes: []interface{}{
//   											&SessionAttributeProperty{
//   												Key: jsii.String("key"),
//
//   												// the properties below are optional
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   									Response: &ResponseSpecificationProperty{
//   										MessageGroupsList: []interface{}{
//   											&MessageGroupProperty{
//   												Message: &MessageProperty{
//   													CustomPayload: &CustomPayloadProperty{
//   														Value: jsii.String("value"),
//   													},
//   													ImageResponseCard: &ImageResponseCardProperty{
//   														Title: jsii.String("title"),
//
//   														// the properties below are optional
//   														Buttons: []interface{}{
//   															&ButtonProperty{
//   																Text: jsii.String("text"),
//   																Value: jsii.String("value"),
//   															},
//   														},
//   														ImageUrl: jsii.String("imageUrl"),
//   														Subtitle: jsii.String("subtitle"),
//   													},
//   													PlainTextMessage: &PlainTextMessageProperty{
//   														Value: jsii.String("value"),
//   													},
//   													SsmlMessage: &SSMLMessageProperty{
//   														Value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												Variations: []interface{}{
//   													&MessageProperty{
//   														CustomPayload: &CustomPayloadProperty{
//   															Value: jsii.String("value"),
//   														},
//   														ImageResponseCard: &ImageResponseCardProperty{
//   															Title: jsii.String("title"),
//
//   															// the properties below are optional
//   															Buttons: []interface{}{
//   																&ButtonProperty{
//   																	Text: jsii.String("text"),
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   															ImageUrl: jsii.String("imageUrl"),
//   															Subtitle: jsii.String("subtitle"),
//   														},
//   														PlainTextMessage: &PlainTextMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   														SsmlMessage: &SSMLMessageProperty{
//   															Value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										AllowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   								IsActive: jsii.Boolean(false),
//   							},
//   							TimeoutNextStep: &DialogStateProperty{
//   								DialogAction: &DialogActionProperty{
//   									Type: jsii.String("type"),
//
//   									// the properties below are optional
//   									SlotToElicit: jsii.String("slotToElicit"),
//   									SuppressNextMessage: jsii.Boolean(false),
//   								},
//   								Intent: &IntentOverrideProperty{
//   									Name: jsii.String("name"),
//   									Slots: []interface{}{
//   										&SlotValueOverrideMapProperty{
//   											SlotName: jsii.String("slotName"),
//   											SlotValueOverride: &slotValueOverrideProperty{
//   												Shape: jsii.String("shape"),
//   												Value: &SlotValueProperty{
//   													InterpretedValue: jsii.String("interpretedValue"),
//   												},
//   												Values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								SessionAttributes: []interface{}{
//   									&SessionAttributeProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							TimeoutResponse: &ResponseSpecificationProperty{
//   								MessageGroupsList: []interface{}{
//   									&MessageGroupProperty{
//   										Message: &MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										Variations: []interface{}{
//   											&MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//
//   						// the properties below are optional
//   						InvocationLabel: jsii.String("invocationLabel"),
//   					},
//   					ElicitationCodeHook: &ElicitationCodeHookInvocationSettingProperty{
//   						EnableCodeHookInvocation: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						InvocationLabel: jsii.String("invocationLabel"),
//   					},
//   					FailureConditional: &ConditionalSpecificationProperty{
//   						ConditionalBranches: []interface{}{
//   							&ConditionalBranchProperty{
//   								Condition: &ConditionProperty{
//   									ExpressionString: jsii.String("expressionString"),
//   								},
//   								Name: jsii.String("name"),
//   								NextStep: &DialogStateProperty{
//   									DialogAction: &DialogActionProperty{
//   										Type: jsii.String("type"),
//
//   										// the properties below are optional
//   										SlotToElicit: jsii.String("slotToElicit"),
//   										SuppressNextMessage: jsii.Boolean(false),
//   									},
//   									Intent: &IntentOverrideProperty{
//   										Name: jsii.String("name"),
//   										Slots: []interface{}{
//   											&SlotValueOverrideMapProperty{
//   												SlotName: jsii.String("slotName"),
//   												SlotValueOverride: &slotValueOverrideProperty{
//   													Shape: jsii.String("shape"),
//   													Value: &SlotValueProperty{
//   														InterpretedValue: jsii.String("interpretedValue"),
//   													},
//   													Values: []interface{}{
//   														slotValueOverrideProperty_,
//   													},
//   												},
//   											},
//   										},
//   									},
//   									SessionAttributes: []interface{}{
//   										&SessionAttributeProperty{
//   											Key: jsii.String("key"),
//
//   											// the properties below are optional
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								Response: &ResponseSpecificationProperty{
//   									MessageGroupsList: []interface{}{
//   										&MessageGroupProperty{
//   											Message: &MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//
//   											// the properties below are optional
//   											Variations: []interface{}{
//   												&MessageProperty{
//   													CustomPayload: &CustomPayloadProperty{
//   														Value: jsii.String("value"),
//   													},
//   													ImageResponseCard: &ImageResponseCardProperty{
//   														Title: jsii.String("title"),
//
//   														// the properties below are optional
//   														Buttons: []interface{}{
//   															&ButtonProperty{
//   																Text: jsii.String("text"),
//   																Value: jsii.String("value"),
//   															},
//   														},
//   														ImageUrl: jsii.String("imageUrl"),
//   														Subtitle: jsii.String("subtitle"),
//   													},
//   													PlainTextMessage: &PlainTextMessageProperty{
//   														Value: jsii.String("value"),
//   													},
//   													SsmlMessage: &SSMLMessageProperty{
//   														Value: jsii.String("value"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									AllowInterrupt: jsii.Boolean(false),
//   								},
//   							},
//   						},
//   						DefaultBranch: &DefaultConditionalBranchProperty{
//   							NextStep: &DialogStateProperty{
//   								DialogAction: &DialogActionProperty{
//   									Type: jsii.String("type"),
//
//   									// the properties below are optional
//   									SlotToElicit: jsii.String("slotToElicit"),
//   									SuppressNextMessage: jsii.Boolean(false),
//   								},
//   								Intent: &IntentOverrideProperty{
//   									Name: jsii.String("name"),
//   									Slots: []interface{}{
//   										&SlotValueOverrideMapProperty{
//   											SlotName: jsii.String("slotName"),
//   											SlotValueOverride: &slotValueOverrideProperty{
//   												Shape: jsii.String("shape"),
//   												Value: &SlotValueProperty{
//   													InterpretedValue: jsii.String("interpretedValue"),
//   												},
//   												Values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								SessionAttributes: []interface{}{
//   									&SessionAttributeProperty{
//   										Key: jsii.String("key"),
//
//   										// the properties below are optional
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							Response: &ResponseSpecificationProperty{
//   								MessageGroupsList: []interface{}{
//   									&MessageGroupProperty{
//   										Message: &MessageProperty{
//   											CustomPayload: &CustomPayloadProperty{
//   												Value: jsii.String("value"),
//   											},
//   											ImageResponseCard: &ImageResponseCardProperty{
//   												Title: jsii.String("title"),
//
//   												// the properties below are optional
//   												Buttons: []interface{}{
//   													&ButtonProperty{
//   														Text: jsii.String("text"),
//   														Value: jsii.String("value"),
//   													},
//   												},
//   												ImageUrl: jsii.String("imageUrl"),
//   												Subtitle: jsii.String("subtitle"),
//   											},
//   											PlainTextMessage: &PlainTextMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   											SsmlMessage: &SSMLMessageProperty{
//   												Value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										Variations: []interface{}{
//   											&MessageProperty{
//   												CustomPayload: &CustomPayloadProperty{
//   													Value: jsii.String("value"),
//   												},
//   												ImageResponseCard: &ImageResponseCardProperty{
//   													Title: jsii.String("title"),
//
//   													// the properties below are optional
//   													Buttons: []interface{}{
//   														&ButtonProperty{
//   															Text: jsii.String("text"),
//   															Value: jsii.String("value"),
//   														},
//   													},
//   													ImageUrl: jsii.String("imageUrl"),
//   													Subtitle: jsii.String("subtitle"),
//   												},
//   												PlainTextMessage: &PlainTextMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   												SsmlMessage: &SSMLMessageProperty{
//   													Value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   						IsActive: jsii.Boolean(false),
//   					},
//   					FailureNextStep: &DialogStateProperty{
//   						DialogAction: &DialogActionProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							SlotToElicit: jsii.String("slotToElicit"),
//   							SuppressNextMessage: jsii.Boolean(false),
//   						},
//   						Intent: &IntentOverrideProperty{
//   							Name: jsii.String("name"),
//   							Slots: []interface{}{
//   								&SlotValueOverrideMapProperty{
//   									SlotName: jsii.String("slotName"),
//   									SlotValueOverride: &slotValueOverrideProperty{
//   										Shape: jsii.String("shape"),
//   										Value: &SlotValueProperty{
//   											InterpretedValue: jsii.String("interpretedValue"),
//   										},
//   										Values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						SessionAttributes: []interface{}{
//   							&SessionAttributeProperty{
//   								Key: jsii.String("key"),
//
//   								// the properties below are optional
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					FailureResponse: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   				WaitAndContinueSpecification: &WaitAndContinueSpecificationProperty{
//   					ContinueResponse: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   					WaitingResponse: &ResponseSpecificationProperty{
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//
//   					// the properties below are optional
//   					IsActive: jsii.Boolean(false),
//   					StillWaitingResponse: &StillWaitingResponseSpecificationProperty{
//   						FrequencyInSeconds: jsii.Number(123),
//   						MessageGroupsList: []interface{}{
//   							&MessageGroupProperty{
//   								Message: &MessageProperty{
//   									CustomPayload: &CustomPayloadProperty{
//   										Value: jsii.String("value"),
//   									},
//   									ImageResponseCard: &ImageResponseCardProperty{
//   										Title: jsii.String("title"),
//
//   										// the properties below are optional
//   										Buttons: []interface{}{
//   											&ButtonProperty{
//   												Text: jsii.String("text"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   										ImageUrl: jsii.String("imageUrl"),
//   										Subtitle: jsii.String("subtitle"),
//   									},
//   									PlainTextMessage: &PlainTextMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   									SsmlMessage: &SSMLMessageProperty{
//   										Value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								Variations: []interface{}{
//   									&MessageProperty{
//   										CustomPayload: &CustomPayloadProperty{
//   											Value: jsii.String("value"),
//   										},
//   										ImageResponseCard: &ImageResponseCardProperty{
//   											Title: jsii.String("title"),
//
//   											// the properties below are optional
//   											Buttons: []interface{}{
//   												&ButtonProperty{
//   													Text: jsii.String("text"),
//   													Value: jsii.String("value"),
//   												},
//   											},
//   											ImageUrl: jsii.String("imageUrl"),
//   											Subtitle: jsii.String("subtitle"),
//   										},
//   										PlainTextMessage: &PlainTextMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   										SsmlMessage: &SSMLMessageProperty{
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						TimeoutInSeconds: jsii.Number(123),
//
//   						// the properties below are optional
//   						AllowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			MultipleValuesSetting: &MultipleValuesSettingProperty{
//   				AllowMultipleValues: jsii.Boolean(false),
//   			},
//   			ObfuscationSetting: &ObfuscationSettingProperty{
//   				ObfuscationSettingType: jsii.String("obfuscationSettingType"),
//   			},
//   		},
//   	},
//   }
//
type CfnBot_IntentProperty struct {
	// The name of the intent.
	//
	// Intent names must be unique within the locale that contains the intent and can't match the name of any built-in intent.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the intent.
	//
	// Use the description to help identify the intent in lists.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies that Amazon Lex invokes the alias Lambda function for each user input.
	//
	// You can invoke this Lambda function to personalize user interaction.
	DialogCodeHook interface{} `field:"optional" json:"dialogCodeHook" yaml:"dialogCodeHook"`
	// Specifies that Amazon Lex invokes the alias Lambda function when the intent is ready for fulfillment.
	//
	// You can invoke this function to complete the bot's transaction with the user.
	FulfillmentCodeHook interface{} `field:"optional" json:"fulfillmentCodeHook" yaml:"fulfillmentCodeHook"`
	// Configuration setting for a response sent to the user before Amazon Lex starts eliciting slots.
	InitialResponseSetting interface{} `field:"optional" json:"initialResponseSetting" yaml:"initialResponseSetting"`
	// A list of contexts that must be active for this intent to be considered by Amazon Lex .
	InputContexts interface{} `field:"optional" json:"inputContexts" yaml:"inputContexts"`
	// Sets the response that Amazon Lex sends to the user when the intent is closed.
	IntentClosingSetting interface{} `field:"optional" json:"intentClosingSetting" yaml:"intentClosingSetting"`
	// Provides prompts that Amazon Lex sends to the user to confirm the completion of an intent.
	//
	// If the user answers "no," the settings contain a statement that is sent to the user to end the intent.
	IntentConfirmationSetting interface{} `field:"optional" json:"intentConfirmationSetting" yaml:"intentConfirmationSetting"`
	// Provides configuration information for the `AMAZON.KendraSearchIntent` intent. When you use this intent, Amazon Lex searches the specified Amazon Kendra index and returns documents from the index that match the user's utterance.
	KendraConfiguration interface{} `field:"optional" json:"kendraConfiguration" yaml:"kendraConfiguration"`
	// A list of contexts that the intent activates when it is fulfilled.
	OutputContexts interface{} `field:"optional" json:"outputContexts" yaml:"outputContexts"`
	// A unique identifier for the built-in intent to base this intent on.
	ParentIntentSignature *string `field:"optional" json:"parentIntentSignature" yaml:"parentIntentSignature"`
	// A list of utterances that a user might say to signal the intent.
	SampleUtterances interface{} `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// Indicates the priority for slots.
	//
	// Amazon Lex prompts the user for slot values in priority order.
	SlotPriorities interface{} `field:"optional" json:"slotPriorities" yaml:"slotPriorities"`
	// A list of slots that the intent requires for fulfillment.
	Slots interface{} `field:"optional" json:"slots" yaml:"slots"`
}

