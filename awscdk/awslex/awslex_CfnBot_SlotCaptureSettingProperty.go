package awslex


// Settings used when Amazon Lex successfully captures a slot value from a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   slotCaptureSettingProperty := &SlotCaptureSettingProperty{
//   	CaptureConditional: &ConditionalSpecificationProperty{
//   		ConditionalBranches: []interface{}{
//   			&ConditionalBranchProperty{
//   				Condition: &ConditionProperty{
//   					ExpressionString: jsii.String("expressionString"),
//   				},
//   				Name: jsii.String("name"),
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
//
//   				// the properties below are optional
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
//   		},
//   		DefaultBranch: &DefaultConditionalBranchProperty{
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
//   		IsActive: jsii.Boolean(false),
//   	},
//   	CaptureNextStep: &DialogStateProperty{
//   		DialogAction: &DialogActionProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			SlotToElicit: jsii.String("slotToElicit"),
//   			SuppressNextMessage: jsii.Boolean(false),
//   		},
//   		Intent: &IntentOverrideProperty{
//   			Name: jsii.String("name"),
//   			Slots: []interface{}{
//   				&SlotValueOverrideMapProperty{
//   					SlotName: jsii.String("slotName"),
//   					SlotValueOverride: &slotValueOverrideProperty{
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
//
//   				// the properties below are optional
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	CaptureResponse: &ResponseSpecificationProperty{
//   		MessageGroupsList: []interface{}{
//   			&MessageGroupProperty{
//   				Message: &MessageProperty{
//   					CustomPayload: &CustomPayloadProperty{
//   						Value: jsii.String("value"),
//   					},
//   					ImageResponseCard: &ImageResponseCardProperty{
//   						Title: jsii.String("title"),
//
//   						// the properties below are optional
//   						Buttons: []interface{}{
//   							&ButtonProperty{
//   								Text: jsii.String("text"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						ImageUrl: jsii.String("imageUrl"),
//   						Subtitle: jsii.String("subtitle"),
//   					},
//   					PlainTextMessage: &PlainTextMessageProperty{
//   						Value: jsii.String("value"),
//   					},
//   					SsmlMessage: &SSMLMessageProperty{
//   						Value: jsii.String("value"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				Variations: []interface{}{
//   					&MessageProperty{
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
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		AllowInterrupt: jsii.Boolean(false),
//   	},
//   	CodeHook: &DialogCodeHookInvocationSettingProperty{
//   		EnableCodeHookInvocation: jsii.Boolean(false),
//   		IsActive: jsii.Boolean(false),
//   		PostCodeHookSpecification: &PostDialogCodeHookInvocationSpecificationProperty{
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
//
//   		// the properties below are optional
//   		InvocationLabel: jsii.String("invocationLabel"),
//   	},
//   	ElicitationCodeHook: &ElicitationCodeHookInvocationSettingProperty{
//   		EnableCodeHookInvocation: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		InvocationLabel: jsii.String("invocationLabel"),
//   	},
//   	FailureConditional: &ConditionalSpecificationProperty{
//   		ConditionalBranches: []interface{}{
//   			&ConditionalBranchProperty{
//   				Condition: &ConditionProperty{
//   					ExpressionString: jsii.String("expressionString"),
//   				},
//   				Name: jsii.String("name"),
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
//
//   				// the properties below are optional
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
//   		},
//   		DefaultBranch: &DefaultConditionalBranchProperty{
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
//   		IsActive: jsii.Boolean(false),
//   	},
//   	FailureNextStep: &DialogStateProperty{
//   		DialogAction: &DialogActionProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			SlotToElicit: jsii.String("slotToElicit"),
//   			SuppressNextMessage: jsii.Boolean(false),
//   		},
//   		Intent: &IntentOverrideProperty{
//   			Name: jsii.String("name"),
//   			Slots: []interface{}{
//   				&SlotValueOverrideMapProperty{
//   					SlotName: jsii.String("slotName"),
//   					SlotValueOverride: &slotValueOverrideProperty{
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
//
//   				// the properties below are optional
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	FailureResponse: &ResponseSpecificationProperty{
//   		MessageGroupsList: []interface{}{
//   			&MessageGroupProperty{
//   				Message: &MessageProperty{
//   					CustomPayload: &CustomPayloadProperty{
//   						Value: jsii.String("value"),
//   					},
//   					ImageResponseCard: &ImageResponseCardProperty{
//   						Title: jsii.String("title"),
//
//   						// the properties below are optional
//   						Buttons: []interface{}{
//   							&ButtonProperty{
//   								Text: jsii.String("text"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						ImageUrl: jsii.String("imageUrl"),
//   						Subtitle: jsii.String("subtitle"),
//   					},
//   					PlainTextMessage: &PlainTextMessageProperty{
//   						Value: jsii.String("value"),
//   					},
//   					SsmlMessage: &SSMLMessageProperty{
//   						Value: jsii.String("value"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				Variations: []interface{}{
//   					&MessageProperty{
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
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		AllowInterrupt: jsii.Boolean(false),
//   	},
//   }
//
type CfnBot_SlotCaptureSettingProperty struct {
	// A list of conditional branches to evaluate after the slot value is captured.
	CaptureConditional interface{} `field:"optional" json:"captureConditional" yaml:"captureConditional"`
	// Specifies the next step that the bot runs when the slot value is captured before the code hook times out.
	CaptureNextStep interface{} `field:"optional" json:"captureNextStep" yaml:"captureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input.
	CaptureResponse interface{} `field:"optional" json:"captureResponse" yaml:"captureResponse"`
	// Code hook called after Amazon Lex successfully captures a slot value.
	CodeHook interface{} `field:"optional" json:"codeHook" yaml:"codeHook"`
	// Code hook called when Amazon Lex doesn't capture a slot value.
	ElicitationCodeHook interface{} `field:"optional" json:"elicitationCodeHook" yaml:"elicitationCodeHook"`
	// A list of conditional branches to evaluate when the slot value isn't captured.
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// Specifies the next step that the bot runs when the slot value code is not recognized.
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input when the slot fails to be captured.
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
}

