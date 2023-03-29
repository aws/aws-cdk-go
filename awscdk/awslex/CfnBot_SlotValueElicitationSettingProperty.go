package awslex


// Specifies the elicitation setting details eliciting a slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   slotValueElicitationSettingProperty := &SlotValueElicitationSettingProperty{
//   	SlotConstraint: jsii.String("slotConstraint"),
//
//   	// the properties below are optional
//   	DefaultValueSpecification: &SlotDefaultValueSpecificationProperty{
//   		DefaultValueList: []interface{}{
//   			&SlotDefaultValueProperty{
//   				DefaultValue: jsii.String("defaultValue"),
//   			},
//   		},
//   	},
//   	PromptSpecification: &PromptSpecificationProperty{
//   		MaxRetries: jsii.Number(123),
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
//   		MessageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   		PromptAttemptsSpecification: map[string]interface{}{
//   			"promptAttemptsSpecificationKey": &PromptAttemptSpecificationProperty{
//   				"allowedInputTypes": &AllowedInputTypesProperty{
//   					"allowAudioInput": jsii.Boolean(false),
//   					"allowDtmfInput": jsii.Boolean(false),
//   				},
//
//   				// the properties below are optional
//   				"allowInterrupt": jsii.Boolean(false),
//   				"audioAndDtmfInputSpecification": &AudioAndDTMFInputSpecificationProperty{
//   					"startTimeoutMs": jsii.Number(123),
//
//   					// the properties below are optional
//   					"audioSpecification": &AudioSpecificationProperty{
//   						"endTimeoutMs": jsii.Number(123),
//   						"maxLengthMs": jsii.Number(123),
//   					},
//   					"dtmfSpecification": &DTMFSpecificationProperty{
//   						"deletionCharacter": jsii.String("deletionCharacter"),
//   						"endCharacter": jsii.String("endCharacter"),
//   						"endTimeoutMs": jsii.Number(123),
//   						"maxLength": jsii.Number(123),
//   					},
//   				},
//   				"textInputSpecification": &TextInputSpecificationProperty{
//   					"startTimeoutMs": jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	SampleUtterances: []interface{}{
//   		&SampleUtteranceProperty{
//   			Utterance: jsii.String("utterance"),
//   		},
//   	},
//   	SlotCaptureSetting: &SlotCaptureSettingProperty{
//   		CaptureConditional: &ConditionalSpecificationProperty{
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
//   		CaptureNextStep: &DialogStateProperty{
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
//   		CaptureResponse: &ResponseSpecificationProperty{
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
//   	},
//   	WaitAndContinueSpecification: &WaitAndContinueSpecificationProperty{
//   		ContinueResponse: &ResponseSpecificationProperty{
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
//   		WaitingResponse: &ResponseSpecificationProperty{
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
//
//   		// the properties below are optional
//   		IsActive: jsii.Boolean(false),
//   		StillWaitingResponse: &StillWaitingResponseSpecificationProperty{
//   			FrequencyInSeconds: jsii.Number(123),
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
//   			TimeoutInSeconds: jsii.Number(123),
//
//   			// the properties below are optional
//   			AllowInterrupt: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnBot_SlotValueElicitationSettingProperty struct {
	// Specifies whether the slot is required or optional.
	SlotConstraint *string `field:"required" json:"slotConstraint" yaml:"slotConstraint"`
	// A list of default values for a slot.
	//
	// Default values are used when Amazon Lex hasn't determined a value for a slot. You can specify default values from context variables, session attributes, and defined values.
	DefaultValueSpecification interface{} `field:"optional" json:"defaultValueSpecification" yaml:"defaultValueSpecification"`
	// The prompt that Amazon Lex uses to elicit the slot value from the user.
	PromptSpecification interface{} `field:"optional" json:"promptSpecification" yaml:"promptSpecification"`
	// If you know a specific pattern that users might respond to an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy.
	//
	// This is optional. In most cases, Amazon Lex is capable of understanding user utterances.
	SampleUtterances interface{} `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// Specifies the settings that Amazon Lex uses when a slot value is successfully entered by a user.
	SlotCaptureSetting interface{} `field:"optional" json:"slotCaptureSetting" yaml:"slotCaptureSetting"`
	// Specifies the prompts that Amazon Lex uses while a bot is waiting for customer input.
	WaitAndContinueSpecification interface{} `field:"optional" json:"waitAndContinueSpecification" yaml:"waitAndContinueSpecification"`
}

