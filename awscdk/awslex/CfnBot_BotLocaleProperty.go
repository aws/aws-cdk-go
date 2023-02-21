package awslex


// Provides configuration information for a locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   botLocaleProperty := &BotLocaleProperty{
//   	LocaleId: jsii.String("localeId"),
//   	NluConfidenceThreshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	CustomVocabulary: &CustomVocabularyProperty{
//   		CustomVocabularyItems: []interface{}{
//   			&CustomVocabularyItemProperty{
//   				Phrase: jsii.String("phrase"),
//
//   				// the properties below are optional
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Intents: []interface{}{
//   		&IntentProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			DialogCodeHook: &DialogCodeHookSettingProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   			FulfillmentCodeHook: &FulfillmentCodeHookSettingProperty{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				FulfillmentUpdatesSpecification: &FulfillmentUpdatesSpecificationProperty{
//   					Active: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					StartResponse: &FulfillmentStartResponseSpecificationProperty{
//   						DelayInSeconds: jsii.Number(123),
//   						MessageGroups: []interface{}{
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
//   					TimeoutInSeconds: jsii.Number(123),
//   					UpdateResponse: &FulfillmentUpdateResponseSpecificationProperty{
//   						FrequencyInSeconds: jsii.Number(123),
//   						MessageGroups: []interface{}{
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
//   				PostFulfillmentStatusSpecification: &PostFulfillmentStatusSpecificationProperty{
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
//   					SuccessConditional: &ConditionalSpecificationProperty{
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
//   					SuccessNextStep: &DialogStateProperty{
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
//   					SuccessResponse: &ResponseSpecificationProperty{
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
//   					TimeoutConditional: &ConditionalSpecificationProperty{
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
//   					TimeoutNextStep: &DialogStateProperty{
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
//   					TimeoutResponse: &ResponseSpecificationProperty{
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
//   			InitialResponseSetting: &InitialResponseSettingProperty{
//   				CodeHook: &DialogCodeHookInvocationSettingProperty{
//   					EnableCodeHookInvocation: jsii.Boolean(false),
//   					IsActive: jsii.Boolean(false),
//   					PostCodeHookSpecification: &PostDialogCodeHookInvocationSpecificationProperty{
//   						FailureConditional: &ConditionalSpecificationProperty{
//   							ConditionalBranches: []interface{}{
//   								&ConditionalBranchProperty{
//   									Condition: &ConditionProperty{
//   										ExpressionString: jsii.String("expressionString"),
//   									},
//   									Name: jsii.String("name"),
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
//
//   									// the properties below are optional
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
//   							},
//   							DefaultBranch: &DefaultConditionalBranchProperty{
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
//   							IsActive: jsii.Boolean(false),
//   						},
//   						FailureNextStep: &DialogStateProperty{
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
//   						FailureResponse: &ResponseSpecificationProperty{
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
//   						SuccessConditional: &ConditionalSpecificationProperty{
//   							ConditionalBranches: []interface{}{
//   								&ConditionalBranchProperty{
//   									Condition: &ConditionProperty{
//   										ExpressionString: jsii.String("expressionString"),
//   									},
//   									Name: jsii.String("name"),
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
//
//   									// the properties below are optional
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
//   							},
//   							DefaultBranch: &DefaultConditionalBranchProperty{
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
//   							IsActive: jsii.Boolean(false),
//   						},
//   						SuccessNextStep: &DialogStateProperty{
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
//   						SuccessResponse: &ResponseSpecificationProperty{
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
//   						TimeoutConditional: &ConditionalSpecificationProperty{
//   							ConditionalBranches: []interface{}{
//   								&ConditionalBranchProperty{
//   									Condition: &ConditionProperty{
//   										ExpressionString: jsii.String("expressionString"),
//   									},
//   									Name: jsii.String("name"),
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
//
//   									// the properties below are optional
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
//   							},
//   							DefaultBranch: &DefaultConditionalBranchProperty{
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
//   							IsActive: jsii.Boolean(false),
//   						},
//   						TimeoutNextStep: &DialogStateProperty{
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
//   						TimeoutResponse: &ResponseSpecificationProperty{
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
//
//   					// the properties below are optional
//   					InvocationLabel: jsii.String("invocationLabel"),
//   				},
//   				Conditional: &ConditionalSpecificationProperty{
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
//   				InitialResponse: &ResponseSpecificationProperty{
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
//   			},
//   			InputContexts: []interface{}{
//   				&InputContextProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			IntentClosingSetting: &IntentClosingSettingProperty{
//   				ClosingResponse: &ResponseSpecificationProperty{
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
//   				Conditional: &ConditionalSpecificationProperty{
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
//   				IsActive: jsii.Boolean(false),
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
//   			},
//   			IntentConfirmationSetting: &IntentConfirmationSettingProperty{
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
//
//   				// the properties below are optional
//   				CodeHook: &DialogCodeHookInvocationSettingProperty{
//   					EnableCodeHookInvocation: jsii.Boolean(false),
//   					IsActive: jsii.Boolean(false),
//   					PostCodeHookSpecification: &PostDialogCodeHookInvocationSpecificationProperty{
//   						FailureConditional: &ConditionalSpecificationProperty{
//   							ConditionalBranches: []interface{}{
//   								&ConditionalBranchProperty{
//   									Condition: &ConditionProperty{
//   										ExpressionString: jsii.String("expressionString"),
//   									},
//   									Name: jsii.String("name"),
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
//
//   									// the properties below are optional
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
//   							},
//   							DefaultBranch: &DefaultConditionalBranchProperty{
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
//   							IsActive: jsii.Boolean(false),
//   						},
//   						FailureNextStep: &DialogStateProperty{
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
//   						FailureResponse: &ResponseSpecificationProperty{
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
//   						SuccessConditional: &ConditionalSpecificationProperty{
//   							ConditionalBranches: []interface{}{
//   								&ConditionalBranchProperty{
//   									Condition: &ConditionProperty{
//   										ExpressionString: jsii.String("expressionString"),
//   									},
//   									Name: jsii.String("name"),
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
//
//   									// the properties below are optional
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
//   							},
//   							DefaultBranch: &DefaultConditionalBranchProperty{
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
//   							IsActive: jsii.Boolean(false),
//   						},
//   						SuccessNextStep: &DialogStateProperty{
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
//   						SuccessResponse: &ResponseSpecificationProperty{
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
//   						TimeoutConditional: &ConditionalSpecificationProperty{
//   							ConditionalBranches: []interface{}{
//   								&ConditionalBranchProperty{
//   									Condition: &ConditionProperty{
//   										ExpressionString: jsii.String("expressionString"),
//   									},
//   									Name: jsii.String("name"),
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
//
//   									// the properties below are optional
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
//   							},
//   							DefaultBranch: &DefaultConditionalBranchProperty{
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
//   							IsActive: jsii.Boolean(false),
//   						},
//   						TimeoutNextStep: &DialogStateProperty{
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
//   						TimeoutResponse: &ResponseSpecificationProperty{
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
//
//   					// the properties below are optional
//   					InvocationLabel: jsii.String("invocationLabel"),
//   				},
//   				ConfirmationConditional: &ConditionalSpecificationProperty{
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
//   				ConfirmationNextStep: &DialogStateProperty{
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
//   				ConfirmationResponse: &ResponseSpecificationProperty{
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
//   				DeclinationConditional: &ConditionalSpecificationProperty{
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
//   				DeclinationNextStep: &DialogStateProperty{
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
//   				DeclinationResponse: &ResponseSpecificationProperty{
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
//   				ElicitationCodeHook: &ElicitationCodeHookInvocationSettingProperty{
//   					EnableCodeHookInvocation: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					InvocationLabel: jsii.String("invocationLabel"),
//   				},
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
//   				IsActive: jsii.Boolean(false),
//   			},
//   			KendraConfiguration: &KendraConfigurationProperty{
//   				KendraIndex: jsii.String("kendraIndex"),
//
//   				// the properties below are optional
//   				QueryFilterString: jsii.String("queryFilterString"),
//   				QueryFilterStringEnabled: jsii.Boolean(false),
//   			},
//   			OutputContexts: []interface{}{
//   				&OutputContextProperty{
//   					Name: jsii.String("name"),
//   					TimeToLiveInSeconds: jsii.Number(123),
//   					TurnsToLive: jsii.Number(123),
//   				},
//   			},
//   			ParentIntentSignature: jsii.String("parentIntentSignature"),
//   			SampleUtterances: []interface{}{
//   				&SampleUtteranceProperty{
//   					Utterance: jsii.String("utterance"),
//   				},
//   			},
//   			SlotPriorities: []interface{}{
//   				&SlotPriorityProperty{
//   					Priority: jsii.Number(123),
//   					SlotName: jsii.String("slotName"),
//   				},
//   			},
//   			Slots: []interface{}{
//   				&SlotProperty{
//   					Name: jsii.String("name"),
//   					SlotTypeName: jsii.String("slotTypeName"),
//   					ValueElicitationSetting: &SlotValueElicitationSettingProperty{
//   						SlotConstraint: jsii.String("slotConstraint"),
//
//   						// the properties below are optional
//   						DefaultValueSpecification: &SlotDefaultValueSpecificationProperty{
//   							DefaultValueList: []interface{}{
//   								&SlotDefaultValueProperty{
//   									DefaultValue: jsii.String("defaultValue"),
//   								},
//   							},
//   						},
//   						PromptSpecification: &PromptSpecificationProperty{
//   							MaxRetries: jsii.Number(123),
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
//   							MessageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   							PromptAttemptsSpecification: map[string]interface{}{
//   								"promptAttemptsSpecificationKey": &PromptAttemptSpecificationProperty{
//   									"allowedInputTypes": &AllowedInputTypesProperty{
//   										"allowAudioInput": jsii.Boolean(false),
//   										"allowDtmfInput": jsii.Boolean(false),
//   									},
//
//   									// the properties below are optional
//   									"allowInterrupt": jsii.Boolean(false),
//   									"audioAndDtmfInputSpecification": &AudioAndDTMFInputSpecificationProperty{
//   										"startTimeoutMs": jsii.Number(123),
//
//   										// the properties below are optional
//   										"audioSpecification": &AudioSpecificationProperty{
//   											"endTimeoutMs": jsii.Number(123),
//   											"maxLengthMs": jsii.Number(123),
//   										},
//   										"dtmfSpecification": &DTMFSpecificationProperty{
//   											"deletionCharacter": jsii.String("deletionCharacter"),
//   											"endCharacter": jsii.String("endCharacter"),
//   											"endTimeoutMs": jsii.Number(123),
//   											"maxLength": jsii.Number(123),
//   										},
//   									},
//   									"textInputSpecification": &TextInputSpecificationProperty{
//   										"startTimeoutMs": jsii.Number(123),
//   									},
//   								},
//   							},
//   						},
//   						SampleUtterances: []interface{}{
//   							&SampleUtteranceProperty{
//   								Utterance: jsii.String("utterance"),
//   							},
//   						},
//   						SlotCaptureSetting: &SlotCaptureSettingProperty{
//   							CaptureConditional: &ConditionalSpecificationProperty{
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
//   							CaptureNextStep: &DialogStateProperty{
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
//   							CaptureResponse: &ResponseSpecificationProperty{
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
//   							CodeHook: &DialogCodeHookInvocationSettingProperty{
//   								EnableCodeHookInvocation: jsii.Boolean(false),
//   								IsActive: jsii.Boolean(false),
//   								PostCodeHookSpecification: &PostDialogCodeHookInvocationSpecificationProperty{
//   									FailureConditional: &ConditionalSpecificationProperty{
//   										ConditionalBranches: []interface{}{
//   											&ConditionalBranchProperty{
//   												Condition: &ConditionProperty{
//   													ExpressionString: jsii.String("expressionString"),
//   												},
//   												Name: jsii.String("name"),
//   												NextStep: &DialogStateProperty{
//   													DialogAction: &DialogActionProperty{
//   														Type: jsii.String("type"),
//
//   														// the properties below are optional
//   														SlotToElicit: jsii.String("slotToElicit"),
//   														SuppressNextMessage: jsii.Boolean(false),
//   													},
//   													Intent: &IntentOverrideProperty{
//   														Name: jsii.String("name"),
//   														Slots: []interface{}{
//   															&SlotValueOverrideMapProperty{
//   																SlotName: jsii.String("slotName"),
//   																SlotValueOverride: &slotValueOverrideProperty{
//   																	Shape: jsii.String("shape"),
//   																	Value: &SlotValueProperty{
//   																		InterpretedValue: jsii.String("interpretedValue"),
//   																	},
//   																	Values: []interface{}{
//   																		slotValueOverrideProperty_,
//   																	},
//   																},
//   															},
//   														},
//   													},
//   													SessionAttributes: []interface{}{
//   														&SessionAttributeProperty{
//   															Key: jsii.String("key"),
//
//   															// the properties below are optional
//   															Value: jsii.String("value"),
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												Response: &ResponseSpecificationProperty{
//   													MessageGroupsList: []interface{}{
//   														&MessageGroupProperty{
//   															Message: &MessageProperty{
//   																CustomPayload: &CustomPayloadProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																ImageResponseCard: &ImageResponseCardProperty{
//   																	Title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	Buttons: []interface{}{
//   																		&ButtonProperty{
//   																			Text: jsii.String("text"),
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//   																	ImageUrl: jsii.String("imageUrl"),
//   																	Subtitle: jsii.String("subtitle"),
//   																},
//   																PlainTextMessage: &PlainTextMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																SsmlMessage: &SSMLMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   															},
//
//   															// the properties below are optional
//   															Variations: []interface{}{
//   																&MessageProperty{
//   																	CustomPayload: &CustomPayloadProperty{
//   																		Value: jsii.String("value"),
//   																	},
//   																	ImageResponseCard: &ImageResponseCardProperty{
//   																		Title: jsii.String("title"),
//
//   																		// the properties below are optional
//   																		Buttons: []interface{}{
//   																			&ButtonProperty{
//   																				Text: jsii.String("text"),
//   																				Value: jsii.String("value"),
//   																			},
//   																		},
//   																		ImageUrl: jsii.String("imageUrl"),
//   																		Subtitle: jsii.String("subtitle"),
//   																	},
//   																	PlainTextMessage: &PlainTextMessageProperty{
//   																		Value: jsii.String("value"),
//   																	},
//   																	SsmlMessage: &SSMLMessageProperty{
//   																		Value: jsii.String("value"),
//   																	},
//   																},
//   															},
//   														},
//   													},
//
//   													// the properties below are optional
//   													AllowInterrupt: jsii.Boolean(false),
//   												},
//   											},
//   										},
//   										DefaultBranch: &DefaultConditionalBranchProperty{
//   											NextStep: &DialogStateProperty{
//   												DialogAction: &DialogActionProperty{
//   													Type: jsii.String("type"),
//
//   													// the properties below are optional
//   													SlotToElicit: jsii.String("slotToElicit"),
//   													SuppressNextMessage: jsii.Boolean(false),
//   												},
//   												Intent: &IntentOverrideProperty{
//   													Name: jsii.String("name"),
//   													Slots: []interface{}{
//   														&SlotValueOverrideMapProperty{
//   															SlotName: jsii.String("slotName"),
//   															SlotValueOverride: &slotValueOverrideProperty{
//   																Shape: jsii.String("shape"),
//   																Value: &SlotValueProperty{
//   																	InterpretedValue: jsii.String("interpretedValue"),
//   																},
//   																Values: []interface{}{
//   																	slotValueOverrideProperty_,
//   																},
//   															},
//   														},
//   													},
//   												},
//   												SessionAttributes: []interface{}{
//   													&SessionAttributeProperty{
//   														Key: jsii.String("key"),
//
//   														// the properties below are optional
//   														Value: jsii.String("value"),
//   													},
//   												},
//   											},
//   											Response: &ResponseSpecificationProperty{
//   												MessageGroupsList: []interface{}{
//   													&MessageGroupProperty{
//   														Message: &MessageProperty{
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
//
//   														// the properties below are optional
//   														Variations: []interface{}{
//   															&MessageProperty{
//   																CustomPayload: &CustomPayloadProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																ImageResponseCard: &ImageResponseCardProperty{
//   																	Title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	Buttons: []interface{}{
//   																		&ButtonProperty{
//   																			Text: jsii.String("text"),
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//   																	ImageUrl: jsii.String("imageUrl"),
//   																	Subtitle: jsii.String("subtitle"),
//   																},
//   																PlainTextMessage: &PlainTextMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																SsmlMessage: &SSMLMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												AllowInterrupt: jsii.Boolean(false),
//   											},
//   										},
//   										IsActive: jsii.Boolean(false),
//   									},
//   									FailureNextStep: &DialogStateProperty{
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
//   									FailureResponse: &ResponseSpecificationProperty{
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
//   									SuccessConditional: &ConditionalSpecificationProperty{
//   										ConditionalBranches: []interface{}{
//   											&ConditionalBranchProperty{
//   												Condition: &ConditionProperty{
//   													ExpressionString: jsii.String("expressionString"),
//   												},
//   												Name: jsii.String("name"),
//   												NextStep: &DialogStateProperty{
//   													DialogAction: &DialogActionProperty{
//   														Type: jsii.String("type"),
//
//   														// the properties below are optional
//   														SlotToElicit: jsii.String("slotToElicit"),
//   														SuppressNextMessage: jsii.Boolean(false),
//   													},
//   													Intent: &IntentOverrideProperty{
//   														Name: jsii.String("name"),
//   														Slots: []interface{}{
//   															&SlotValueOverrideMapProperty{
//   																SlotName: jsii.String("slotName"),
//   																SlotValueOverride: &slotValueOverrideProperty{
//   																	Shape: jsii.String("shape"),
//   																	Value: &SlotValueProperty{
//   																		InterpretedValue: jsii.String("interpretedValue"),
//   																	},
//   																	Values: []interface{}{
//   																		slotValueOverrideProperty_,
//   																	},
//   																},
//   															},
//   														},
//   													},
//   													SessionAttributes: []interface{}{
//   														&SessionAttributeProperty{
//   															Key: jsii.String("key"),
//
//   															// the properties below are optional
//   															Value: jsii.String("value"),
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												Response: &ResponseSpecificationProperty{
//   													MessageGroupsList: []interface{}{
//   														&MessageGroupProperty{
//   															Message: &MessageProperty{
//   																CustomPayload: &CustomPayloadProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																ImageResponseCard: &ImageResponseCardProperty{
//   																	Title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	Buttons: []interface{}{
//   																		&ButtonProperty{
//   																			Text: jsii.String("text"),
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//   																	ImageUrl: jsii.String("imageUrl"),
//   																	Subtitle: jsii.String("subtitle"),
//   																},
//   																PlainTextMessage: &PlainTextMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																SsmlMessage: &SSMLMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   															},
//
//   															// the properties below are optional
//   															Variations: []interface{}{
//   																&MessageProperty{
//   																	CustomPayload: &CustomPayloadProperty{
//   																		Value: jsii.String("value"),
//   																	},
//   																	ImageResponseCard: &ImageResponseCardProperty{
//   																		Title: jsii.String("title"),
//
//   																		// the properties below are optional
//   																		Buttons: []interface{}{
//   																			&ButtonProperty{
//   																				Text: jsii.String("text"),
//   																				Value: jsii.String("value"),
//   																			},
//   																		},
//   																		ImageUrl: jsii.String("imageUrl"),
//   																		Subtitle: jsii.String("subtitle"),
//   																	},
//   																	PlainTextMessage: &PlainTextMessageProperty{
//   																		Value: jsii.String("value"),
//   																	},
//   																	SsmlMessage: &SSMLMessageProperty{
//   																		Value: jsii.String("value"),
//   																	},
//   																},
//   															},
//   														},
//   													},
//
//   													// the properties below are optional
//   													AllowInterrupt: jsii.Boolean(false),
//   												},
//   											},
//   										},
//   										DefaultBranch: &DefaultConditionalBranchProperty{
//   											NextStep: &DialogStateProperty{
//   												DialogAction: &DialogActionProperty{
//   													Type: jsii.String("type"),
//
//   													// the properties below are optional
//   													SlotToElicit: jsii.String("slotToElicit"),
//   													SuppressNextMessage: jsii.Boolean(false),
//   												},
//   												Intent: &IntentOverrideProperty{
//   													Name: jsii.String("name"),
//   													Slots: []interface{}{
//   														&SlotValueOverrideMapProperty{
//   															SlotName: jsii.String("slotName"),
//   															SlotValueOverride: &slotValueOverrideProperty{
//   																Shape: jsii.String("shape"),
//   																Value: &SlotValueProperty{
//   																	InterpretedValue: jsii.String("interpretedValue"),
//   																},
//   																Values: []interface{}{
//   																	slotValueOverrideProperty_,
//   																},
//   															},
//   														},
//   													},
//   												},
//   												SessionAttributes: []interface{}{
//   													&SessionAttributeProperty{
//   														Key: jsii.String("key"),
//
//   														// the properties below are optional
//   														Value: jsii.String("value"),
//   													},
//   												},
//   											},
//   											Response: &ResponseSpecificationProperty{
//   												MessageGroupsList: []interface{}{
//   													&MessageGroupProperty{
//   														Message: &MessageProperty{
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
//
//   														// the properties below are optional
//   														Variations: []interface{}{
//   															&MessageProperty{
//   																CustomPayload: &CustomPayloadProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																ImageResponseCard: &ImageResponseCardProperty{
//   																	Title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	Buttons: []interface{}{
//   																		&ButtonProperty{
//   																			Text: jsii.String("text"),
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//   																	ImageUrl: jsii.String("imageUrl"),
//   																	Subtitle: jsii.String("subtitle"),
//   																},
//   																PlainTextMessage: &PlainTextMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																SsmlMessage: &SSMLMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												AllowInterrupt: jsii.Boolean(false),
//   											},
//   										},
//   										IsActive: jsii.Boolean(false),
//   									},
//   									SuccessNextStep: &DialogStateProperty{
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
//   									SuccessResponse: &ResponseSpecificationProperty{
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
//   									TimeoutConditional: &ConditionalSpecificationProperty{
//   										ConditionalBranches: []interface{}{
//   											&ConditionalBranchProperty{
//   												Condition: &ConditionProperty{
//   													ExpressionString: jsii.String("expressionString"),
//   												},
//   												Name: jsii.String("name"),
//   												NextStep: &DialogStateProperty{
//   													DialogAction: &DialogActionProperty{
//   														Type: jsii.String("type"),
//
//   														// the properties below are optional
//   														SlotToElicit: jsii.String("slotToElicit"),
//   														SuppressNextMessage: jsii.Boolean(false),
//   													},
//   													Intent: &IntentOverrideProperty{
//   														Name: jsii.String("name"),
//   														Slots: []interface{}{
//   															&SlotValueOverrideMapProperty{
//   																SlotName: jsii.String("slotName"),
//   																SlotValueOverride: &slotValueOverrideProperty{
//   																	Shape: jsii.String("shape"),
//   																	Value: &SlotValueProperty{
//   																		InterpretedValue: jsii.String("interpretedValue"),
//   																	},
//   																	Values: []interface{}{
//   																		slotValueOverrideProperty_,
//   																	},
//   																},
//   															},
//   														},
//   													},
//   													SessionAttributes: []interface{}{
//   														&SessionAttributeProperty{
//   															Key: jsii.String("key"),
//
//   															// the properties below are optional
//   															Value: jsii.String("value"),
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												Response: &ResponseSpecificationProperty{
//   													MessageGroupsList: []interface{}{
//   														&MessageGroupProperty{
//   															Message: &MessageProperty{
//   																CustomPayload: &CustomPayloadProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																ImageResponseCard: &ImageResponseCardProperty{
//   																	Title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	Buttons: []interface{}{
//   																		&ButtonProperty{
//   																			Text: jsii.String("text"),
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//   																	ImageUrl: jsii.String("imageUrl"),
//   																	Subtitle: jsii.String("subtitle"),
//   																},
//   																PlainTextMessage: &PlainTextMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																SsmlMessage: &SSMLMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   															},
//
//   															// the properties below are optional
//   															Variations: []interface{}{
//   																&MessageProperty{
//   																	CustomPayload: &CustomPayloadProperty{
//   																		Value: jsii.String("value"),
//   																	},
//   																	ImageResponseCard: &ImageResponseCardProperty{
//   																		Title: jsii.String("title"),
//
//   																		// the properties below are optional
//   																		Buttons: []interface{}{
//   																			&ButtonProperty{
//   																				Text: jsii.String("text"),
//   																				Value: jsii.String("value"),
//   																			},
//   																		},
//   																		ImageUrl: jsii.String("imageUrl"),
//   																		Subtitle: jsii.String("subtitle"),
//   																	},
//   																	PlainTextMessage: &PlainTextMessageProperty{
//   																		Value: jsii.String("value"),
//   																	},
//   																	SsmlMessage: &SSMLMessageProperty{
//   																		Value: jsii.String("value"),
//   																	},
//   																},
//   															},
//   														},
//   													},
//
//   													// the properties below are optional
//   													AllowInterrupt: jsii.Boolean(false),
//   												},
//   											},
//   										},
//   										DefaultBranch: &DefaultConditionalBranchProperty{
//   											NextStep: &DialogStateProperty{
//   												DialogAction: &DialogActionProperty{
//   													Type: jsii.String("type"),
//
//   													// the properties below are optional
//   													SlotToElicit: jsii.String("slotToElicit"),
//   													SuppressNextMessage: jsii.Boolean(false),
//   												},
//   												Intent: &IntentOverrideProperty{
//   													Name: jsii.String("name"),
//   													Slots: []interface{}{
//   														&SlotValueOverrideMapProperty{
//   															SlotName: jsii.String("slotName"),
//   															SlotValueOverride: &slotValueOverrideProperty{
//   																Shape: jsii.String("shape"),
//   																Value: &SlotValueProperty{
//   																	InterpretedValue: jsii.String("interpretedValue"),
//   																},
//   																Values: []interface{}{
//   																	slotValueOverrideProperty_,
//   																},
//   															},
//   														},
//   													},
//   												},
//   												SessionAttributes: []interface{}{
//   													&SessionAttributeProperty{
//   														Key: jsii.String("key"),
//
//   														// the properties below are optional
//   														Value: jsii.String("value"),
//   													},
//   												},
//   											},
//   											Response: &ResponseSpecificationProperty{
//   												MessageGroupsList: []interface{}{
//   													&MessageGroupProperty{
//   														Message: &MessageProperty{
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
//
//   														// the properties below are optional
//   														Variations: []interface{}{
//   															&MessageProperty{
//   																CustomPayload: &CustomPayloadProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																ImageResponseCard: &ImageResponseCardProperty{
//   																	Title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	Buttons: []interface{}{
//   																		&ButtonProperty{
//   																			Text: jsii.String("text"),
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//   																	ImageUrl: jsii.String("imageUrl"),
//   																	Subtitle: jsii.String("subtitle"),
//   																},
//   																PlainTextMessage: &PlainTextMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   																SsmlMessage: &SSMLMessageProperty{
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												AllowInterrupt: jsii.Boolean(false),
//   											},
//   										},
//   										IsActive: jsii.Boolean(false),
//   									},
//   									TimeoutNextStep: &DialogStateProperty{
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
//   									TimeoutResponse: &ResponseSpecificationProperty{
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
//
//   								// the properties below are optional
//   								InvocationLabel: jsii.String("invocationLabel"),
//   							},
//   							ElicitationCodeHook: &ElicitationCodeHookInvocationSettingProperty{
//   								EnableCodeHookInvocation: jsii.Boolean(false),
//
//   								// the properties below are optional
//   								InvocationLabel: jsii.String("invocationLabel"),
//   							},
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
//   						},
//   						WaitAndContinueSpecification: &WaitAndContinueSpecificationProperty{
//   							ContinueResponse: &ResponseSpecificationProperty{
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
//   							WaitingResponse: &ResponseSpecificationProperty{
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
//
//   							// the properties below are optional
//   							IsActive: jsii.Boolean(false),
//   							StillWaitingResponse: &StillWaitingResponseSpecificationProperty{
//   								FrequencyInSeconds: jsii.Number(123),
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
//   								TimeoutInSeconds: jsii.Number(123),
//
//   								// the properties below are optional
//   								AllowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					Description: jsii.String("description"),
//   					MultipleValuesSetting: &MultipleValuesSettingProperty{
//   						AllowMultipleValues: jsii.Boolean(false),
//   					},
//   					ObfuscationSetting: &ObfuscationSettingProperty{
//   						ObfuscationSettingType: jsii.String("obfuscationSettingType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SlotTypes: []interface{}{
//   		&SlotTypeProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			ExternalSourceSetting: &ExternalSourceSettingProperty{
//   				GrammarSlotTypeSetting: &GrammarSlotTypeSettingProperty{
//   					Source: &GrammarSlotTypeSourceProperty{
//   						S3BucketName: jsii.String("s3BucketName"),
//   						S3ObjectKey: jsii.String("s3ObjectKey"),
//
//   						// the properties below are optional
//   						KmsKeyArn: jsii.String("kmsKeyArn"),
//   					},
//   				},
//   			},
//   			ParentSlotTypeSignature: jsii.String("parentSlotTypeSignature"),
//   			SlotTypeValues: []interface{}{
//   				&SlotTypeValueProperty{
//   					SampleValue: &SampleValueProperty{
//   						Value: jsii.String("value"),
//   					},
//
//   					// the properties below are optional
//   					Synonyms: []interface{}{
//   						&SampleValueProperty{
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			ValueSelectionSetting: &SlotValueSelectionSettingProperty{
//   				ResolutionStrategy: jsii.String("resolutionStrategy"),
//
//   				// the properties below are optional
//   				AdvancedRecognitionSetting: &AdvancedRecognitionSettingProperty{
//   					AudioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   				},
//   				RegexFilter: &SlotValueRegexFilterProperty{
//   					Pattern: jsii.String("pattern"),
//   				},
//   			},
//   		},
//   	},
//   	VoiceSettings: &VoiceSettingsProperty{
//   		VoiceId: jsii.String("voiceId"),
//
//   		// the properties below are optional
//   		Engine: jsii.String("engine"),
//   	},
//   }
//
type CfnBot_BotLocaleProperty struct {
	// The identifier of the language and locale that the bot will be used in.
	//
	// The string must match one of the supported locales.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Determines the threshold where Amazon Lex will insert the `AMAZON.FallbackIntent` , `AMAZON.KendraSearchIntent` , or both when returning alternative intents. You must configure an `AMAZON.FallbackIntent` . `AMAZON.KendraSearchIntent` is only inserted if it is configured for the bot.
	NluConfidenceThreshold *float64 `field:"required" json:"nluConfidenceThreshold" yaml:"nluConfidenceThreshold"`
	// Specifies a custom vocabulary to use with a specific locale.
	CustomVocabulary interface{} `field:"optional" json:"customVocabulary" yaml:"customVocabulary"`
	// A description of the bot locale.
	//
	// Use this to help identify the bot locale in lists.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// One or more intents defined for the locale.
	Intents interface{} `field:"optional" json:"intents" yaml:"intents"`
	// One or more slot types defined for the locale.
	SlotTypes interface{} `field:"optional" json:"slotTypes" yaml:"slotTypes"`
	// Defines settings for using an Amazon Polly voice to communicate with a user.
	VoiceSettings interface{} `field:"optional" json:"voiceSettings" yaml:"voiceSettings"`
}

