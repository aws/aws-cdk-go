package awslex


// Provides a setting that determines whether the post-fulfillment response is sent to the user.
//
// For more information, see [](https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   postFulfillmentStatusSpecificationProperty := &PostFulfillmentStatusSpecificationProperty{
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
//   	SuccessConditional: &ConditionalSpecificationProperty{
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
//   	SuccessNextStep: &DialogStateProperty{
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
//   	SuccessResponse: &ResponseSpecificationProperty{
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
//   	TimeoutConditional: &ConditionalSpecificationProperty{
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
//   	TimeoutNextStep: &DialogStateProperty{
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
//   	TimeoutResponse: &ResponseSpecificationProperty{
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
type CfnBot_PostFulfillmentStatusSpecificationProperty struct {
	// A list of conditional branches to evaluate after the fulfillment code hook throws an exception or returns with the `State` field of the `Intent` object set to `Failed` .
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// Specifies the next step the bot runs after the fulfillment code hook throws an exception or returns with the `State` field of the `Intent` object set to `Failed` .
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't successful.
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// A list of conditional branches to evaluate after the fulfillment code hook finishes successfully.
	SuccessConditional interface{} `field:"optional" json:"successConditional" yaml:"successConditional"`
	// Specifies the next step in the conversation that Amazon Lex invokes when the fulfillment code hook completes successfully.
	SuccessNextStep interface{} `field:"optional" json:"successNextStep" yaml:"successNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when the fulfillment is successful.
	SuccessResponse interface{} `field:"optional" json:"successResponse" yaml:"successResponse"`
	// A list of conditional branches to evaluate if the fulfillment code hook times out.
	TimeoutConditional interface{} `field:"optional" json:"timeoutConditional" yaml:"timeoutConditional"`
	// Specifies the next step that the bot runs when the fulfillment code hook times out.
	TimeoutNextStep interface{} `field:"optional" json:"timeoutNextStep" yaml:"timeoutNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't completed within the timeout period.
	TimeoutResponse interface{} `field:"optional" json:"timeoutResponse" yaml:"timeoutResponse"`
}

