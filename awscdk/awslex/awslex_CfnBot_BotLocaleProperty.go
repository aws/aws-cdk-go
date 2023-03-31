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
//   botLocaleProperty := &botLocaleProperty{
//   	localeId: jsii.String("localeId"),
//   	nluConfidenceThreshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	customVocabulary: &customVocabularyProperty{
//   		customVocabularyItems: []interface{}{
//   			&customVocabularyItemProperty{
//   				phrase: jsii.String("phrase"),
//
//   				// the properties below are optional
//   				displayAs: jsii.String("displayAs"),
//   				weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	intents: []interface{}{
//   		&intentProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			dialogCodeHook: &dialogCodeHookSettingProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   			fulfillmentCodeHook: &fulfillmentCodeHookSettingProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				fulfillmentUpdatesSpecification: &fulfillmentUpdatesSpecificationProperty{
//   					active: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					startResponse: &fulfillmentStartResponseSpecificationProperty{
//   						delayInSeconds: jsii.Number(123),
//   						messageGroups: []interface{}{
//   							&messageGroupProperty{
//   								message: &messageProperty{
//   									customPayload: &customPayloadProperty{
//   										value: jsii.String("value"),
//   									},
//   									imageResponseCard: &imageResponseCardProperty{
//   										title: jsii.String("title"),
//
//   										// the properties below are optional
//   										buttons: []interface{}{
//   											&buttonProperty{
//   												text: jsii.String("text"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   										imageUrl: jsii.String("imageUrl"),
//   										subtitle: jsii.String("subtitle"),
//   									},
//   									plainTextMessage: &plainTextMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   									ssmlMessage: &sSMLMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								variations: []interface{}{
//   									&messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						allowInterrupt: jsii.Boolean(false),
//   					},
//   					timeoutInSeconds: jsii.Number(123),
//   					updateResponse: &fulfillmentUpdateResponseSpecificationProperty{
//   						frequencyInSeconds: jsii.Number(123),
//   						messageGroups: []interface{}{
//   							&messageGroupProperty{
//   								message: &messageProperty{
//   									customPayload: &customPayloadProperty{
//   										value: jsii.String("value"),
//   									},
//   									imageResponseCard: &imageResponseCardProperty{
//   										title: jsii.String("title"),
//
//   										// the properties below are optional
//   										buttons: []interface{}{
//   											&buttonProperty{
//   												text: jsii.String("text"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   										imageUrl: jsii.String("imageUrl"),
//   										subtitle: jsii.String("subtitle"),
//   									},
//   									plainTextMessage: &plainTextMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   									ssmlMessage: &sSMLMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								variations: []interface{}{
//   									&messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						allowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   				isActive: jsii.Boolean(false),
//   				postFulfillmentStatusSpecification: &postFulfillmentStatusSpecificationProperty{
//   					failureConditional: &conditionalSpecificationProperty{
//   						conditionalBranches: []interface{}{
//   							&conditionalBranchProperty{
//   								condition: &conditionProperty{
//   									expressionString: jsii.String("expressionString"),
//   								},
//   								name: jsii.String("name"),
//   								nextStep: &dialogStateProperty{
//   									dialogAction: &dialogActionProperty{
//   										type: jsii.String("type"),
//
//   										// the properties below are optional
//   										slotToElicit: jsii.String("slotToElicit"),
//   										suppressNextMessage: jsii.Boolean(false),
//   									},
//   									intent: &intentOverrideProperty{
//   										name: jsii.String("name"),
//   										slots: []interface{}{
//   											&slotValueOverrideMapProperty{
//   												slotName: jsii.String("slotName"),
//   												slotValueOverride: &slotValueOverrideProperty{
//   													shape: jsii.String("shape"),
//   													value: &slotValueProperty{
//   														interpretedValue: jsii.String("interpretedValue"),
//   													},
//   													values: []interface{}{
//   														slotValueOverrideProperty_,
//   													},
//   												},
//   											},
//   										},
//   									},
//   									sessionAttributes: []interface{}{
//   										&sessionAttributeProperty{
//   											key: jsii.String("key"),
//
//   											// the properties below are optional
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								response: &responseSpecificationProperty{
//   									messageGroupsList: []interface{}{
//   										&messageGroupProperty{
//   											message: &messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//
//   											// the properties below are optional
//   											variations: []interface{}{
//   												&messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									allowInterrupt: jsii.Boolean(false),
//   								},
//   							},
//   						},
//   						defaultBranch: &defaultConditionalBranchProperty{
//   							nextStep: &dialogStateProperty{
//   								dialogAction: &dialogActionProperty{
//   									type: jsii.String("type"),
//
//   									// the properties below are optional
//   									slotToElicit: jsii.String("slotToElicit"),
//   									suppressNextMessage: jsii.Boolean(false),
//   								},
//   								intent: &intentOverrideProperty{
//   									name: jsii.String("name"),
//   									slots: []interface{}{
//   										&slotValueOverrideMapProperty{
//   											slotName: jsii.String("slotName"),
//   											slotValueOverride: &slotValueOverrideProperty{
//   												shape: jsii.String("shape"),
//   												value: &slotValueProperty{
//   													interpretedValue: jsii.String("interpretedValue"),
//   												},
//   												values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								sessionAttributes: []interface{}{
//   									&sessionAttributeProperty{
//   										key: jsii.String("key"),
//
//   										// the properties below are optional
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							response: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   						isActive: jsii.Boolean(false),
//   					},
//   					failureNextStep: &dialogStateProperty{
//   						dialogAction: &dialogActionProperty{
//   							type: jsii.String("type"),
//
//   							// the properties below are optional
//   							slotToElicit: jsii.String("slotToElicit"),
//   							suppressNextMessage: jsii.Boolean(false),
//   						},
//   						intent: &intentOverrideProperty{
//   							name: jsii.String("name"),
//   							slots: []interface{}{
//   								&slotValueOverrideMapProperty{
//   									slotName: jsii.String("slotName"),
//   									slotValueOverride: &slotValueOverrideProperty{
//   										shape: jsii.String("shape"),
//   										value: &slotValueProperty{
//   											interpretedValue: jsii.String("interpretedValue"),
//   										},
//   										values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						sessionAttributes: []interface{}{
//   							&sessionAttributeProperty{
//   								key: jsii.String("key"),
//
//   								// the properties below are optional
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					failureResponse: &responseSpecificationProperty{
//   						messageGroupsList: []interface{}{
//   							&messageGroupProperty{
//   								message: &messageProperty{
//   									customPayload: &customPayloadProperty{
//   										value: jsii.String("value"),
//   									},
//   									imageResponseCard: &imageResponseCardProperty{
//   										title: jsii.String("title"),
//
//   										// the properties below are optional
//   										buttons: []interface{}{
//   											&buttonProperty{
//   												text: jsii.String("text"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   										imageUrl: jsii.String("imageUrl"),
//   										subtitle: jsii.String("subtitle"),
//   									},
//   									plainTextMessage: &plainTextMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   									ssmlMessage: &sSMLMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								variations: []interface{}{
//   									&messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						allowInterrupt: jsii.Boolean(false),
//   					},
//   					successConditional: &conditionalSpecificationProperty{
//   						conditionalBranches: []interface{}{
//   							&conditionalBranchProperty{
//   								condition: &conditionProperty{
//   									expressionString: jsii.String("expressionString"),
//   								},
//   								name: jsii.String("name"),
//   								nextStep: &dialogStateProperty{
//   									dialogAction: &dialogActionProperty{
//   										type: jsii.String("type"),
//
//   										// the properties below are optional
//   										slotToElicit: jsii.String("slotToElicit"),
//   										suppressNextMessage: jsii.Boolean(false),
//   									},
//   									intent: &intentOverrideProperty{
//   										name: jsii.String("name"),
//   										slots: []interface{}{
//   											&slotValueOverrideMapProperty{
//   												slotName: jsii.String("slotName"),
//   												slotValueOverride: &slotValueOverrideProperty{
//   													shape: jsii.String("shape"),
//   													value: &slotValueProperty{
//   														interpretedValue: jsii.String("interpretedValue"),
//   													},
//   													values: []interface{}{
//   														slotValueOverrideProperty_,
//   													},
//   												},
//   											},
//   										},
//   									},
//   									sessionAttributes: []interface{}{
//   										&sessionAttributeProperty{
//   											key: jsii.String("key"),
//
//   											// the properties below are optional
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								response: &responseSpecificationProperty{
//   									messageGroupsList: []interface{}{
//   										&messageGroupProperty{
//   											message: &messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//
//   											// the properties below are optional
//   											variations: []interface{}{
//   												&messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									allowInterrupt: jsii.Boolean(false),
//   								},
//   							},
//   						},
//   						defaultBranch: &defaultConditionalBranchProperty{
//   							nextStep: &dialogStateProperty{
//   								dialogAction: &dialogActionProperty{
//   									type: jsii.String("type"),
//
//   									// the properties below are optional
//   									slotToElicit: jsii.String("slotToElicit"),
//   									suppressNextMessage: jsii.Boolean(false),
//   								},
//   								intent: &intentOverrideProperty{
//   									name: jsii.String("name"),
//   									slots: []interface{}{
//   										&slotValueOverrideMapProperty{
//   											slotName: jsii.String("slotName"),
//   											slotValueOverride: &slotValueOverrideProperty{
//   												shape: jsii.String("shape"),
//   												value: &slotValueProperty{
//   													interpretedValue: jsii.String("interpretedValue"),
//   												},
//   												values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								sessionAttributes: []interface{}{
//   									&sessionAttributeProperty{
//   										key: jsii.String("key"),
//
//   										// the properties below are optional
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							response: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   						isActive: jsii.Boolean(false),
//   					},
//   					successNextStep: &dialogStateProperty{
//   						dialogAction: &dialogActionProperty{
//   							type: jsii.String("type"),
//
//   							// the properties below are optional
//   							slotToElicit: jsii.String("slotToElicit"),
//   							suppressNextMessage: jsii.Boolean(false),
//   						},
//   						intent: &intentOverrideProperty{
//   							name: jsii.String("name"),
//   							slots: []interface{}{
//   								&slotValueOverrideMapProperty{
//   									slotName: jsii.String("slotName"),
//   									slotValueOverride: &slotValueOverrideProperty{
//   										shape: jsii.String("shape"),
//   										value: &slotValueProperty{
//   											interpretedValue: jsii.String("interpretedValue"),
//   										},
//   										values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						sessionAttributes: []interface{}{
//   							&sessionAttributeProperty{
//   								key: jsii.String("key"),
//
//   								// the properties below are optional
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					successResponse: &responseSpecificationProperty{
//   						messageGroupsList: []interface{}{
//   							&messageGroupProperty{
//   								message: &messageProperty{
//   									customPayload: &customPayloadProperty{
//   										value: jsii.String("value"),
//   									},
//   									imageResponseCard: &imageResponseCardProperty{
//   										title: jsii.String("title"),
//
//   										// the properties below are optional
//   										buttons: []interface{}{
//   											&buttonProperty{
//   												text: jsii.String("text"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   										imageUrl: jsii.String("imageUrl"),
//   										subtitle: jsii.String("subtitle"),
//   									},
//   									plainTextMessage: &plainTextMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   									ssmlMessage: &sSMLMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								variations: []interface{}{
//   									&messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						allowInterrupt: jsii.Boolean(false),
//   					},
//   					timeoutConditional: &conditionalSpecificationProperty{
//   						conditionalBranches: []interface{}{
//   							&conditionalBranchProperty{
//   								condition: &conditionProperty{
//   									expressionString: jsii.String("expressionString"),
//   								},
//   								name: jsii.String("name"),
//   								nextStep: &dialogStateProperty{
//   									dialogAction: &dialogActionProperty{
//   										type: jsii.String("type"),
//
//   										// the properties below are optional
//   										slotToElicit: jsii.String("slotToElicit"),
//   										suppressNextMessage: jsii.Boolean(false),
//   									},
//   									intent: &intentOverrideProperty{
//   										name: jsii.String("name"),
//   										slots: []interface{}{
//   											&slotValueOverrideMapProperty{
//   												slotName: jsii.String("slotName"),
//   												slotValueOverride: &slotValueOverrideProperty{
//   													shape: jsii.String("shape"),
//   													value: &slotValueProperty{
//   														interpretedValue: jsii.String("interpretedValue"),
//   													},
//   													values: []interface{}{
//   														slotValueOverrideProperty_,
//   													},
//   												},
//   											},
//   										},
//   									},
//   									sessionAttributes: []interface{}{
//   										&sessionAttributeProperty{
//   											key: jsii.String("key"),
//
//   											// the properties below are optional
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								response: &responseSpecificationProperty{
//   									messageGroupsList: []interface{}{
//   										&messageGroupProperty{
//   											message: &messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//
//   											// the properties below are optional
//   											variations: []interface{}{
//   												&messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									allowInterrupt: jsii.Boolean(false),
//   								},
//   							},
//   						},
//   						defaultBranch: &defaultConditionalBranchProperty{
//   							nextStep: &dialogStateProperty{
//   								dialogAction: &dialogActionProperty{
//   									type: jsii.String("type"),
//
//   									// the properties below are optional
//   									slotToElicit: jsii.String("slotToElicit"),
//   									suppressNextMessage: jsii.Boolean(false),
//   								},
//   								intent: &intentOverrideProperty{
//   									name: jsii.String("name"),
//   									slots: []interface{}{
//   										&slotValueOverrideMapProperty{
//   											slotName: jsii.String("slotName"),
//   											slotValueOverride: &slotValueOverrideProperty{
//   												shape: jsii.String("shape"),
//   												value: &slotValueProperty{
//   													interpretedValue: jsii.String("interpretedValue"),
//   												},
//   												values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								sessionAttributes: []interface{}{
//   									&sessionAttributeProperty{
//   										key: jsii.String("key"),
//
//   										// the properties below are optional
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							response: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   						isActive: jsii.Boolean(false),
//   					},
//   					timeoutNextStep: &dialogStateProperty{
//   						dialogAction: &dialogActionProperty{
//   							type: jsii.String("type"),
//
//   							// the properties below are optional
//   							slotToElicit: jsii.String("slotToElicit"),
//   							suppressNextMessage: jsii.Boolean(false),
//   						},
//   						intent: &intentOverrideProperty{
//   							name: jsii.String("name"),
//   							slots: []interface{}{
//   								&slotValueOverrideMapProperty{
//   									slotName: jsii.String("slotName"),
//   									slotValueOverride: &slotValueOverrideProperty{
//   										shape: jsii.String("shape"),
//   										value: &slotValueProperty{
//   											interpretedValue: jsii.String("interpretedValue"),
//   										},
//   										values: []interface{}{
//   											slotValueOverrideProperty_,
//   										},
//   									},
//   								},
//   							},
//   						},
//   						sessionAttributes: []interface{}{
//   							&sessionAttributeProperty{
//   								key: jsii.String("key"),
//
//   								// the properties below are optional
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					timeoutResponse: &responseSpecificationProperty{
//   						messageGroupsList: []interface{}{
//   							&messageGroupProperty{
//   								message: &messageProperty{
//   									customPayload: &customPayloadProperty{
//   										value: jsii.String("value"),
//   									},
//   									imageResponseCard: &imageResponseCardProperty{
//   										title: jsii.String("title"),
//
//   										// the properties below are optional
//   										buttons: []interface{}{
//   											&buttonProperty{
//   												text: jsii.String("text"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   										imageUrl: jsii.String("imageUrl"),
//   										subtitle: jsii.String("subtitle"),
//   									},
//   									plainTextMessage: &plainTextMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   									ssmlMessage: &sSMLMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								variations: []interface{}{
//   									&messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						allowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			initialResponseSetting: &initialResponseSettingProperty{
//   				codeHook: &dialogCodeHookInvocationSettingProperty{
//   					enableCodeHookInvocation: jsii.Boolean(false),
//   					isActive: jsii.Boolean(false),
//   					postCodeHookSpecification: &postDialogCodeHookInvocationSpecificationProperty{
//   						failureConditional: &conditionalSpecificationProperty{
//   							conditionalBranches: []interface{}{
//   								&conditionalBranchProperty{
//   									condition: &conditionProperty{
//   										expressionString: jsii.String("expressionString"),
//   									},
//   									name: jsii.String("name"),
//   									nextStep: &dialogStateProperty{
//   										dialogAction: &dialogActionProperty{
//   											type: jsii.String("type"),
//
//   											// the properties below are optional
//   											slotToElicit: jsii.String("slotToElicit"),
//   											suppressNextMessage: jsii.Boolean(false),
//   										},
//   										intent: &intentOverrideProperty{
//   											name: jsii.String("name"),
//   											slots: []interface{}{
//   												&slotValueOverrideMapProperty{
//   													slotName: jsii.String("slotName"),
//   													slotValueOverride: &slotValueOverrideProperty{
//   														shape: jsii.String("shape"),
//   														value: &slotValueProperty{
//   															interpretedValue: jsii.String("interpretedValue"),
//   														},
//   														values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										sessionAttributes: []interface{}{
//   											&sessionAttributeProperty{
//   												key: jsii.String("key"),
//
//   												// the properties below are optional
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									response: &responseSpecificationProperty{
//   										messageGroupsList: []interface{}{
//   											&messageGroupProperty{
//   												message: &messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												variations: []interface{}{
//   													&messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							defaultBranch: &defaultConditionalBranchProperty{
//   								nextStep: &dialogStateProperty{
//   									dialogAction: &dialogActionProperty{
//   										type: jsii.String("type"),
//
//   										// the properties below are optional
//   										slotToElicit: jsii.String("slotToElicit"),
//   										suppressNextMessage: jsii.Boolean(false),
//   									},
//   									intent: &intentOverrideProperty{
//   										name: jsii.String("name"),
//   										slots: []interface{}{
//   											&slotValueOverrideMapProperty{
//   												slotName: jsii.String("slotName"),
//   												slotValueOverride: &slotValueOverrideProperty{
//   													shape: jsii.String("shape"),
//   													value: &slotValueProperty{
//   														interpretedValue: jsii.String("interpretedValue"),
//   													},
//   													values: []interface{}{
//   														slotValueOverrideProperty_,
//   													},
//   												},
//   											},
//   										},
//   									},
//   									sessionAttributes: []interface{}{
//   										&sessionAttributeProperty{
//   											key: jsii.String("key"),
//
//   											// the properties below are optional
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								response: &responseSpecificationProperty{
//   									messageGroupsList: []interface{}{
//   										&messageGroupProperty{
//   											message: &messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//
//   											// the properties below are optional
//   											variations: []interface{}{
//   												&messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									allowInterrupt: jsii.Boolean(false),
//   								},
//   							},
//   							isActive: jsii.Boolean(false),
//   						},
//   						failureNextStep: &dialogStateProperty{
//   							dialogAction: &dialogActionProperty{
//   								type: jsii.String("type"),
//
//   								// the properties below are optional
//   								slotToElicit: jsii.String("slotToElicit"),
//   								suppressNextMessage: jsii.Boolean(false),
//   							},
//   							intent: &intentOverrideProperty{
//   								name: jsii.String("name"),
//   								slots: []interface{}{
//   									&slotValueOverrideMapProperty{
//   										slotName: jsii.String("slotName"),
//   										slotValueOverride: &slotValueOverrideProperty{
//   											shape: jsii.String("shape"),
//   											value: &slotValueProperty{
//   												interpretedValue: jsii.String("interpretedValue"),
//   											},
//   											values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							sessionAttributes: []interface{}{
//   								&sessionAttributeProperty{
//   									key: jsii.String("key"),
//
//   									// the properties below are optional
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						failureResponse: &responseSpecificationProperty{
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   						},
//   						successConditional: &conditionalSpecificationProperty{
//   							conditionalBranches: []interface{}{
//   								&conditionalBranchProperty{
//   									condition: &conditionProperty{
//   										expressionString: jsii.String("expressionString"),
//   									},
//   									name: jsii.String("name"),
//   									nextStep: &dialogStateProperty{
//   										dialogAction: &dialogActionProperty{
//   											type: jsii.String("type"),
//
//   											// the properties below are optional
//   											slotToElicit: jsii.String("slotToElicit"),
//   											suppressNextMessage: jsii.Boolean(false),
//   										},
//   										intent: &intentOverrideProperty{
//   											name: jsii.String("name"),
//   											slots: []interface{}{
//   												&slotValueOverrideMapProperty{
//   													slotName: jsii.String("slotName"),
//   													slotValueOverride: &slotValueOverrideProperty{
//   														shape: jsii.String("shape"),
//   														value: &slotValueProperty{
//   															interpretedValue: jsii.String("interpretedValue"),
//   														},
//   														values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										sessionAttributes: []interface{}{
//   											&sessionAttributeProperty{
//   												key: jsii.String("key"),
//
//   												// the properties below are optional
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									response: &responseSpecificationProperty{
//   										messageGroupsList: []interface{}{
//   											&messageGroupProperty{
//   												message: &messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												variations: []interface{}{
//   													&messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							defaultBranch: &defaultConditionalBranchProperty{
//   								nextStep: &dialogStateProperty{
//   									dialogAction: &dialogActionProperty{
//   										type: jsii.String("type"),
//
//   										// the properties below are optional
//   										slotToElicit: jsii.String("slotToElicit"),
//   										suppressNextMessage: jsii.Boolean(false),
//   									},
//   									intent: &intentOverrideProperty{
//   										name: jsii.String("name"),
//   										slots: []interface{}{
//   											&slotValueOverrideMapProperty{
//   												slotName: jsii.String("slotName"),
//   												slotValueOverride: &slotValueOverrideProperty{
//   													shape: jsii.String("shape"),
//   													value: &slotValueProperty{
//   														interpretedValue: jsii.String("interpretedValue"),
//   													},
//   													values: []interface{}{
//   														slotValueOverrideProperty_,
//   													},
//   												},
//   											},
//   										},
//   									},
//   									sessionAttributes: []interface{}{
//   										&sessionAttributeProperty{
//   											key: jsii.String("key"),
//
//   											// the properties below are optional
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								response: &responseSpecificationProperty{
//   									messageGroupsList: []interface{}{
//   										&messageGroupProperty{
//   											message: &messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//
//   											// the properties below are optional
//   											variations: []interface{}{
//   												&messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									allowInterrupt: jsii.Boolean(false),
//   								},
//   							},
//   							isActive: jsii.Boolean(false),
//   						},
//   						successNextStep: &dialogStateProperty{
//   							dialogAction: &dialogActionProperty{
//   								type: jsii.String("type"),
//
//   								// the properties below are optional
//   								slotToElicit: jsii.String("slotToElicit"),
//   								suppressNextMessage: jsii.Boolean(false),
//   							},
//   							intent: &intentOverrideProperty{
//   								name: jsii.String("name"),
//   								slots: []interface{}{
//   									&slotValueOverrideMapProperty{
//   										slotName: jsii.String("slotName"),
//   										slotValueOverride: &slotValueOverrideProperty{
//   											shape: jsii.String("shape"),
//   											value: &slotValueProperty{
//   												interpretedValue: jsii.String("interpretedValue"),
//   											},
//   											values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							sessionAttributes: []interface{}{
//   								&sessionAttributeProperty{
//   									key: jsii.String("key"),
//
//   									// the properties below are optional
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						successResponse: &responseSpecificationProperty{
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   						},
//   						timeoutConditional: &conditionalSpecificationProperty{
//   							conditionalBranches: []interface{}{
//   								&conditionalBranchProperty{
//   									condition: &conditionProperty{
//   										expressionString: jsii.String("expressionString"),
//   									},
//   									name: jsii.String("name"),
//   									nextStep: &dialogStateProperty{
//   										dialogAction: &dialogActionProperty{
//   											type: jsii.String("type"),
//
//   											// the properties below are optional
//   											slotToElicit: jsii.String("slotToElicit"),
//   											suppressNextMessage: jsii.Boolean(false),
//   										},
//   										intent: &intentOverrideProperty{
//   											name: jsii.String("name"),
//   											slots: []interface{}{
//   												&slotValueOverrideMapProperty{
//   													slotName: jsii.String("slotName"),
//   													slotValueOverride: &slotValueOverrideProperty{
//   														shape: jsii.String("shape"),
//   														value: &slotValueProperty{
//   															interpretedValue: jsii.String("interpretedValue"),
//   														},
//   														values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										sessionAttributes: []interface{}{
//   											&sessionAttributeProperty{
//   												key: jsii.String("key"),
//
//   												// the properties below are optional
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									response: &responseSpecificationProperty{
//   										messageGroupsList: []interface{}{
//   											&messageGroupProperty{
//   												message: &messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												variations: []interface{}{
//   													&messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							defaultBranch: &defaultConditionalBranchProperty{
//   								nextStep: &dialogStateProperty{
//   									dialogAction: &dialogActionProperty{
//   										type: jsii.String("type"),
//
//   										// the properties below are optional
//   										slotToElicit: jsii.String("slotToElicit"),
//   										suppressNextMessage: jsii.Boolean(false),
//   									},
//   									intent: &intentOverrideProperty{
//   										name: jsii.String("name"),
//   										slots: []interface{}{
//   											&slotValueOverrideMapProperty{
//   												slotName: jsii.String("slotName"),
//   												slotValueOverride: &slotValueOverrideProperty{
//   													shape: jsii.String("shape"),
//   													value: &slotValueProperty{
//   														interpretedValue: jsii.String("interpretedValue"),
//   													},
//   													values: []interface{}{
//   														slotValueOverrideProperty_,
//   													},
//   												},
//   											},
//   										},
//   									},
//   									sessionAttributes: []interface{}{
//   										&sessionAttributeProperty{
//   											key: jsii.String("key"),
//
//   											// the properties below are optional
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								response: &responseSpecificationProperty{
//   									messageGroupsList: []interface{}{
//   										&messageGroupProperty{
//   											message: &messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//
//   											// the properties below are optional
//   											variations: []interface{}{
//   												&messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									allowInterrupt: jsii.Boolean(false),
//   								},
//   							},
//   							isActive: jsii.Boolean(false),
//   						},
//   						timeoutNextStep: &dialogStateProperty{
//   							dialogAction: &dialogActionProperty{
//   								type: jsii.String("type"),
//
//   								// the properties below are optional
//   								slotToElicit: jsii.String("slotToElicit"),
//   								suppressNextMessage: jsii.Boolean(false),
//   							},
//   							intent: &intentOverrideProperty{
//   								name: jsii.String("name"),
//   								slots: []interface{}{
//   									&slotValueOverrideMapProperty{
//   										slotName: jsii.String("slotName"),
//   										slotValueOverride: &slotValueOverrideProperty{
//   											shape: jsii.String("shape"),
//   											value: &slotValueProperty{
//   												interpretedValue: jsii.String("interpretedValue"),
//   											},
//   											values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							sessionAttributes: []interface{}{
//   								&sessionAttributeProperty{
//   									key: jsii.String("key"),
//
//   									// the properties below are optional
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						timeoutResponse: &responseSpecificationProperty{
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//
//   					// the properties below are optional
//   					invocationLabel: jsii.String("invocationLabel"),
//   				},
//   				conditional: &conditionalSpecificationProperty{
//   					conditionalBranches: []interface{}{
//   						&conditionalBranchProperty{
//   							condition: &conditionProperty{
//   								expressionString: jsii.String("expressionString"),
//   							},
//   							name: jsii.String("name"),
//   							nextStep: &dialogStateProperty{
//   								dialogAction: &dialogActionProperty{
//   									type: jsii.String("type"),
//
//   									// the properties below are optional
//   									slotToElicit: jsii.String("slotToElicit"),
//   									suppressNextMessage: jsii.Boolean(false),
//   								},
//   								intent: &intentOverrideProperty{
//   									name: jsii.String("name"),
//   									slots: []interface{}{
//   										&slotValueOverrideMapProperty{
//   											slotName: jsii.String("slotName"),
//   											slotValueOverride: &slotValueOverrideProperty{
//   												shape: jsii.String("shape"),
//   												value: &slotValueProperty{
//   													interpretedValue: jsii.String("interpretedValue"),
//   												},
//   												values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								sessionAttributes: []interface{}{
//   									&sessionAttributeProperty{
//   										key: jsii.String("key"),
//
//   										// the properties below are optional
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							response: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					defaultBranch: &defaultConditionalBranchProperty{
//   						nextStep: &dialogStateProperty{
//   							dialogAction: &dialogActionProperty{
//   								type: jsii.String("type"),
//
//   								// the properties below are optional
//   								slotToElicit: jsii.String("slotToElicit"),
//   								suppressNextMessage: jsii.Boolean(false),
//   							},
//   							intent: &intentOverrideProperty{
//   								name: jsii.String("name"),
//   								slots: []interface{}{
//   									&slotValueOverrideMapProperty{
//   										slotName: jsii.String("slotName"),
//   										slotValueOverride: &slotValueOverrideProperty{
//   											shape: jsii.String("shape"),
//   											value: &slotValueProperty{
//   												interpretedValue: jsii.String("interpretedValue"),
//   											},
//   											values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							sessionAttributes: []interface{}{
//   								&sessionAttributeProperty{
//   									key: jsii.String("key"),
//
//   									// the properties below are optional
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						response: &responseSpecificationProperty{
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   					isActive: jsii.Boolean(false),
//   				},
//   				initialResponse: &responseSpecificationProperty{
//   					messageGroupsList: []interface{}{
//   						&messageGroupProperty{
//   							message: &messageProperty{
//   								customPayload: &customPayloadProperty{
//   									value: jsii.String("value"),
//   								},
//   								imageResponseCard: &imageResponseCardProperty{
//   									title: jsii.String("title"),
//
//   									// the properties below are optional
//   									buttons: []interface{}{
//   										&buttonProperty{
//   											text: jsii.String("text"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   									imageUrl: jsii.String("imageUrl"),
//   									subtitle: jsii.String("subtitle"),
//   								},
//   								plainTextMessage: &plainTextMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   								ssmlMessage: &sSMLMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							variations: []interface{}{
//   								&messageProperty{
//   									customPayload: &customPayloadProperty{
//   										value: jsii.String("value"),
//   									},
//   									imageResponseCard: &imageResponseCardProperty{
//   										title: jsii.String("title"),
//
//   										// the properties below are optional
//   										buttons: []interface{}{
//   											&buttonProperty{
//   												text: jsii.String("text"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   										imageUrl: jsii.String("imageUrl"),
//   										subtitle: jsii.String("subtitle"),
//   									},
//   									plainTextMessage: &plainTextMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   									ssmlMessage: &sSMLMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					allowInterrupt: jsii.Boolean(false),
//   				},
//   				nextStep: &dialogStateProperty{
//   					dialogAction: &dialogActionProperty{
//   						type: jsii.String("type"),
//
//   						// the properties below are optional
//   						slotToElicit: jsii.String("slotToElicit"),
//   						suppressNextMessage: jsii.Boolean(false),
//   					},
//   					intent: &intentOverrideProperty{
//   						name: jsii.String("name"),
//   						slots: []interface{}{
//   							&slotValueOverrideMapProperty{
//   								slotName: jsii.String("slotName"),
//   								slotValueOverride: &slotValueOverrideProperty{
//   									shape: jsii.String("shape"),
//   									value: &slotValueProperty{
//   										interpretedValue: jsii.String("interpretedValue"),
//   									},
//   									values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					sessionAttributes: []interface{}{
//   						&sessionAttributeProperty{
//   							key: jsii.String("key"),
//
//   							// the properties below are optional
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			inputContexts: []interface{}{
//   				&inputContextProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			intentClosingSetting: &intentClosingSettingProperty{
//   				closingResponse: &responseSpecificationProperty{
//   					messageGroupsList: []interface{}{
//   						&messageGroupProperty{
//   							message: &messageProperty{
//   								customPayload: &customPayloadProperty{
//   									value: jsii.String("value"),
//   								},
//   								imageResponseCard: &imageResponseCardProperty{
//   									title: jsii.String("title"),
//
//   									// the properties below are optional
//   									buttons: []interface{}{
//   										&buttonProperty{
//   											text: jsii.String("text"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   									imageUrl: jsii.String("imageUrl"),
//   									subtitle: jsii.String("subtitle"),
//   								},
//   								plainTextMessage: &plainTextMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   								ssmlMessage: &sSMLMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							variations: []interface{}{
//   								&messageProperty{
//   									customPayload: &customPayloadProperty{
//   										value: jsii.String("value"),
//   									},
//   									imageResponseCard: &imageResponseCardProperty{
//   										title: jsii.String("title"),
//
//   										// the properties below are optional
//   										buttons: []interface{}{
//   											&buttonProperty{
//   												text: jsii.String("text"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   										imageUrl: jsii.String("imageUrl"),
//   										subtitle: jsii.String("subtitle"),
//   									},
//   									plainTextMessage: &plainTextMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   									ssmlMessage: &sSMLMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					allowInterrupt: jsii.Boolean(false),
//   				},
//   				conditional: &conditionalSpecificationProperty{
//   					conditionalBranches: []interface{}{
//   						&conditionalBranchProperty{
//   							condition: &conditionProperty{
//   								expressionString: jsii.String("expressionString"),
//   							},
//   							name: jsii.String("name"),
//   							nextStep: &dialogStateProperty{
//   								dialogAction: &dialogActionProperty{
//   									type: jsii.String("type"),
//
//   									// the properties below are optional
//   									slotToElicit: jsii.String("slotToElicit"),
//   									suppressNextMessage: jsii.Boolean(false),
//   								},
//   								intent: &intentOverrideProperty{
//   									name: jsii.String("name"),
//   									slots: []interface{}{
//   										&slotValueOverrideMapProperty{
//   											slotName: jsii.String("slotName"),
//   											slotValueOverride: &slotValueOverrideProperty{
//   												shape: jsii.String("shape"),
//   												value: &slotValueProperty{
//   													interpretedValue: jsii.String("interpretedValue"),
//   												},
//   												values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								sessionAttributes: []interface{}{
//   									&sessionAttributeProperty{
//   										key: jsii.String("key"),
//
//   										// the properties below are optional
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							response: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					defaultBranch: &defaultConditionalBranchProperty{
//   						nextStep: &dialogStateProperty{
//   							dialogAction: &dialogActionProperty{
//   								type: jsii.String("type"),
//
//   								// the properties below are optional
//   								slotToElicit: jsii.String("slotToElicit"),
//   								suppressNextMessage: jsii.Boolean(false),
//   							},
//   							intent: &intentOverrideProperty{
//   								name: jsii.String("name"),
//   								slots: []interface{}{
//   									&slotValueOverrideMapProperty{
//   										slotName: jsii.String("slotName"),
//   										slotValueOverride: &slotValueOverrideProperty{
//   											shape: jsii.String("shape"),
//   											value: &slotValueProperty{
//   												interpretedValue: jsii.String("interpretedValue"),
//   											},
//   											values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							sessionAttributes: []interface{}{
//   								&sessionAttributeProperty{
//   									key: jsii.String("key"),
//
//   									// the properties below are optional
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						response: &responseSpecificationProperty{
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   					isActive: jsii.Boolean(false),
//   				},
//   				isActive: jsii.Boolean(false),
//   				nextStep: &dialogStateProperty{
//   					dialogAction: &dialogActionProperty{
//   						type: jsii.String("type"),
//
//   						// the properties below are optional
//   						slotToElicit: jsii.String("slotToElicit"),
//   						suppressNextMessage: jsii.Boolean(false),
//   					},
//   					intent: &intentOverrideProperty{
//   						name: jsii.String("name"),
//   						slots: []interface{}{
//   							&slotValueOverrideMapProperty{
//   								slotName: jsii.String("slotName"),
//   								slotValueOverride: &slotValueOverrideProperty{
//   									shape: jsii.String("shape"),
//   									value: &slotValueProperty{
//   										interpretedValue: jsii.String("interpretedValue"),
//   									},
//   									values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					sessionAttributes: []interface{}{
//   						&sessionAttributeProperty{
//   							key: jsii.String("key"),
//
//   							// the properties below are optional
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			intentConfirmationSetting: &intentConfirmationSettingProperty{
//   				promptSpecification: &promptSpecificationProperty{
//   					maxRetries: jsii.Number(123),
//   					messageGroupsList: []interface{}{
//   						&messageGroupProperty{
//   							message: &messageProperty{
//   								customPayload: &customPayloadProperty{
//   									value: jsii.String("value"),
//   								},
//   								imageResponseCard: &imageResponseCardProperty{
//   									title: jsii.String("title"),
//
//   									// the properties below are optional
//   									buttons: []interface{}{
//   										&buttonProperty{
//   											text: jsii.String("text"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   									imageUrl: jsii.String("imageUrl"),
//   									subtitle: jsii.String("subtitle"),
//   								},
//   								plainTextMessage: &plainTextMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   								ssmlMessage: &sSMLMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							variations: []interface{}{
//   								&messageProperty{
//   									customPayload: &customPayloadProperty{
//   										value: jsii.String("value"),
//   									},
//   									imageResponseCard: &imageResponseCardProperty{
//   										title: jsii.String("title"),
//
//   										// the properties below are optional
//   										buttons: []interface{}{
//   											&buttonProperty{
//   												text: jsii.String("text"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   										imageUrl: jsii.String("imageUrl"),
//   										subtitle: jsii.String("subtitle"),
//   									},
//   									plainTextMessage: &plainTextMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   									ssmlMessage: &sSMLMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					allowInterrupt: jsii.Boolean(false),
//   					messageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   					promptAttemptsSpecification: map[string]interface{}{
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
//   				codeHook: &dialogCodeHookInvocationSettingProperty{
//   					enableCodeHookInvocation: jsii.Boolean(false),
//   					isActive: jsii.Boolean(false),
//   					postCodeHookSpecification: &postDialogCodeHookInvocationSpecificationProperty{
//   						failureConditional: &conditionalSpecificationProperty{
//   							conditionalBranches: []interface{}{
//   								&conditionalBranchProperty{
//   									condition: &conditionProperty{
//   										expressionString: jsii.String("expressionString"),
//   									},
//   									name: jsii.String("name"),
//   									nextStep: &dialogStateProperty{
//   										dialogAction: &dialogActionProperty{
//   											type: jsii.String("type"),
//
//   											// the properties below are optional
//   											slotToElicit: jsii.String("slotToElicit"),
//   											suppressNextMessage: jsii.Boolean(false),
//   										},
//   										intent: &intentOverrideProperty{
//   											name: jsii.String("name"),
//   											slots: []interface{}{
//   												&slotValueOverrideMapProperty{
//   													slotName: jsii.String("slotName"),
//   													slotValueOverride: &slotValueOverrideProperty{
//   														shape: jsii.String("shape"),
//   														value: &slotValueProperty{
//   															interpretedValue: jsii.String("interpretedValue"),
//   														},
//   														values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										sessionAttributes: []interface{}{
//   											&sessionAttributeProperty{
//   												key: jsii.String("key"),
//
//   												// the properties below are optional
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									response: &responseSpecificationProperty{
//   										messageGroupsList: []interface{}{
//   											&messageGroupProperty{
//   												message: &messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												variations: []interface{}{
//   													&messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							defaultBranch: &defaultConditionalBranchProperty{
//   								nextStep: &dialogStateProperty{
//   									dialogAction: &dialogActionProperty{
//   										type: jsii.String("type"),
//
//   										// the properties below are optional
//   										slotToElicit: jsii.String("slotToElicit"),
//   										suppressNextMessage: jsii.Boolean(false),
//   									},
//   									intent: &intentOverrideProperty{
//   										name: jsii.String("name"),
//   										slots: []interface{}{
//   											&slotValueOverrideMapProperty{
//   												slotName: jsii.String("slotName"),
//   												slotValueOverride: &slotValueOverrideProperty{
//   													shape: jsii.String("shape"),
//   													value: &slotValueProperty{
//   														interpretedValue: jsii.String("interpretedValue"),
//   													},
//   													values: []interface{}{
//   														slotValueOverrideProperty_,
//   													},
//   												},
//   											},
//   										},
//   									},
//   									sessionAttributes: []interface{}{
//   										&sessionAttributeProperty{
//   											key: jsii.String("key"),
//
//   											// the properties below are optional
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								response: &responseSpecificationProperty{
//   									messageGroupsList: []interface{}{
//   										&messageGroupProperty{
//   											message: &messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//
//   											// the properties below are optional
//   											variations: []interface{}{
//   												&messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									allowInterrupt: jsii.Boolean(false),
//   								},
//   							},
//   							isActive: jsii.Boolean(false),
//   						},
//   						failureNextStep: &dialogStateProperty{
//   							dialogAction: &dialogActionProperty{
//   								type: jsii.String("type"),
//
//   								// the properties below are optional
//   								slotToElicit: jsii.String("slotToElicit"),
//   								suppressNextMessage: jsii.Boolean(false),
//   							},
//   							intent: &intentOverrideProperty{
//   								name: jsii.String("name"),
//   								slots: []interface{}{
//   									&slotValueOverrideMapProperty{
//   										slotName: jsii.String("slotName"),
//   										slotValueOverride: &slotValueOverrideProperty{
//   											shape: jsii.String("shape"),
//   											value: &slotValueProperty{
//   												interpretedValue: jsii.String("interpretedValue"),
//   											},
//   											values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							sessionAttributes: []interface{}{
//   								&sessionAttributeProperty{
//   									key: jsii.String("key"),
//
//   									// the properties below are optional
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						failureResponse: &responseSpecificationProperty{
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   						},
//   						successConditional: &conditionalSpecificationProperty{
//   							conditionalBranches: []interface{}{
//   								&conditionalBranchProperty{
//   									condition: &conditionProperty{
//   										expressionString: jsii.String("expressionString"),
//   									},
//   									name: jsii.String("name"),
//   									nextStep: &dialogStateProperty{
//   										dialogAction: &dialogActionProperty{
//   											type: jsii.String("type"),
//
//   											// the properties below are optional
//   											slotToElicit: jsii.String("slotToElicit"),
//   											suppressNextMessage: jsii.Boolean(false),
//   										},
//   										intent: &intentOverrideProperty{
//   											name: jsii.String("name"),
//   											slots: []interface{}{
//   												&slotValueOverrideMapProperty{
//   													slotName: jsii.String("slotName"),
//   													slotValueOverride: &slotValueOverrideProperty{
//   														shape: jsii.String("shape"),
//   														value: &slotValueProperty{
//   															interpretedValue: jsii.String("interpretedValue"),
//   														},
//   														values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										sessionAttributes: []interface{}{
//   											&sessionAttributeProperty{
//   												key: jsii.String("key"),
//
//   												// the properties below are optional
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									response: &responseSpecificationProperty{
//   										messageGroupsList: []interface{}{
//   											&messageGroupProperty{
//   												message: &messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												variations: []interface{}{
//   													&messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							defaultBranch: &defaultConditionalBranchProperty{
//   								nextStep: &dialogStateProperty{
//   									dialogAction: &dialogActionProperty{
//   										type: jsii.String("type"),
//
//   										// the properties below are optional
//   										slotToElicit: jsii.String("slotToElicit"),
//   										suppressNextMessage: jsii.Boolean(false),
//   									},
//   									intent: &intentOverrideProperty{
//   										name: jsii.String("name"),
//   										slots: []interface{}{
//   											&slotValueOverrideMapProperty{
//   												slotName: jsii.String("slotName"),
//   												slotValueOverride: &slotValueOverrideProperty{
//   													shape: jsii.String("shape"),
//   													value: &slotValueProperty{
//   														interpretedValue: jsii.String("interpretedValue"),
//   													},
//   													values: []interface{}{
//   														slotValueOverrideProperty_,
//   													},
//   												},
//   											},
//   										},
//   									},
//   									sessionAttributes: []interface{}{
//   										&sessionAttributeProperty{
//   											key: jsii.String("key"),
//
//   											// the properties below are optional
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								response: &responseSpecificationProperty{
//   									messageGroupsList: []interface{}{
//   										&messageGroupProperty{
//   											message: &messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//
//   											// the properties below are optional
//   											variations: []interface{}{
//   												&messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									allowInterrupt: jsii.Boolean(false),
//   								},
//   							},
//   							isActive: jsii.Boolean(false),
//   						},
//   						successNextStep: &dialogStateProperty{
//   							dialogAction: &dialogActionProperty{
//   								type: jsii.String("type"),
//
//   								// the properties below are optional
//   								slotToElicit: jsii.String("slotToElicit"),
//   								suppressNextMessage: jsii.Boolean(false),
//   							},
//   							intent: &intentOverrideProperty{
//   								name: jsii.String("name"),
//   								slots: []interface{}{
//   									&slotValueOverrideMapProperty{
//   										slotName: jsii.String("slotName"),
//   										slotValueOverride: &slotValueOverrideProperty{
//   											shape: jsii.String("shape"),
//   											value: &slotValueProperty{
//   												interpretedValue: jsii.String("interpretedValue"),
//   											},
//   											values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							sessionAttributes: []interface{}{
//   								&sessionAttributeProperty{
//   									key: jsii.String("key"),
//
//   									// the properties below are optional
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						successResponse: &responseSpecificationProperty{
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   						},
//   						timeoutConditional: &conditionalSpecificationProperty{
//   							conditionalBranches: []interface{}{
//   								&conditionalBranchProperty{
//   									condition: &conditionProperty{
//   										expressionString: jsii.String("expressionString"),
//   									},
//   									name: jsii.String("name"),
//   									nextStep: &dialogStateProperty{
//   										dialogAction: &dialogActionProperty{
//   											type: jsii.String("type"),
//
//   											// the properties below are optional
//   											slotToElicit: jsii.String("slotToElicit"),
//   											suppressNextMessage: jsii.Boolean(false),
//   										},
//   										intent: &intentOverrideProperty{
//   											name: jsii.String("name"),
//   											slots: []interface{}{
//   												&slotValueOverrideMapProperty{
//   													slotName: jsii.String("slotName"),
//   													slotValueOverride: &slotValueOverrideProperty{
//   														shape: jsii.String("shape"),
//   														value: &slotValueProperty{
//   															interpretedValue: jsii.String("interpretedValue"),
//   														},
//   														values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										sessionAttributes: []interface{}{
//   											&sessionAttributeProperty{
//   												key: jsii.String("key"),
//
//   												// the properties below are optional
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									response: &responseSpecificationProperty{
//   										messageGroupsList: []interface{}{
//   											&messageGroupProperty{
//   												message: &messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												variations: []interface{}{
//   													&messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   							},
//   							defaultBranch: &defaultConditionalBranchProperty{
//   								nextStep: &dialogStateProperty{
//   									dialogAction: &dialogActionProperty{
//   										type: jsii.String("type"),
//
//   										// the properties below are optional
//   										slotToElicit: jsii.String("slotToElicit"),
//   										suppressNextMessage: jsii.Boolean(false),
//   									},
//   									intent: &intentOverrideProperty{
//   										name: jsii.String("name"),
//   										slots: []interface{}{
//   											&slotValueOverrideMapProperty{
//   												slotName: jsii.String("slotName"),
//   												slotValueOverride: &slotValueOverrideProperty{
//   													shape: jsii.String("shape"),
//   													value: &slotValueProperty{
//   														interpretedValue: jsii.String("interpretedValue"),
//   													},
//   													values: []interface{}{
//   														slotValueOverrideProperty_,
//   													},
//   												},
//   											},
//   										},
//   									},
//   									sessionAttributes: []interface{}{
//   										&sessionAttributeProperty{
//   											key: jsii.String("key"),
//
//   											// the properties below are optional
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								response: &responseSpecificationProperty{
//   									messageGroupsList: []interface{}{
//   										&messageGroupProperty{
//   											message: &messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//
//   											// the properties below are optional
//   											variations: []interface{}{
//   												&messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   										},
//   									},
//
//   									// the properties below are optional
//   									allowInterrupt: jsii.Boolean(false),
//   								},
//   							},
//   							isActive: jsii.Boolean(false),
//   						},
//   						timeoutNextStep: &dialogStateProperty{
//   							dialogAction: &dialogActionProperty{
//   								type: jsii.String("type"),
//
//   								// the properties below are optional
//   								slotToElicit: jsii.String("slotToElicit"),
//   								suppressNextMessage: jsii.Boolean(false),
//   							},
//   							intent: &intentOverrideProperty{
//   								name: jsii.String("name"),
//   								slots: []interface{}{
//   									&slotValueOverrideMapProperty{
//   										slotName: jsii.String("slotName"),
//   										slotValueOverride: &slotValueOverrideProperty{
//   											shape: jsii.String("shape"),
//   											value: &slotValueProperty{
//   												interpretedValue: jsii.String("interpretedValue"),
//   											},
//   											values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							sessionAttributes: []interface{}{
//   								&sessionAttributeProperty{
//   									key: jsii.String("key"),
//
//   									// the properties below are optional
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						timeoutResponse: &responseSpecificationProperty{
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//
//   					// the properties below are optional
//   					invocationLabel: jsii.String("invocationLabel"),
//   				},
//   				confirmationConditional: &conditionalSpecificationProperty{
//   					conditionalBranches: []interface{}{
//   						&conditionalBranchProperty{
//   							condition: &conditionProperty{
//   								expressionString: jsii.String("expressionString"),
//   							},
//   							name: jsii.String("name"),
//   							nextStep: &dialogStateProperty{
//   								dialogAction: &dialogActionProperty{
//   									type: jsii.String("type"),
//
//   									// the properties below are optional
//   									slotToElicit: jsii.String("slotToElicit"),
//   									suppressNextMessage: jsii.Boolean(false),
//   								},
//   								intent: &intentOverrideProperty{
//   									name: jsii.String("name"),
//   									slots: []interface{}{
//   										&slotValueOverrideMapProperty{
//   											slotName: jsii.String("slotName"),
//   											slotValueOverride: &slotValueOverrideProperty{
//   												shape: jsii.String("shape"),
//   												value: &slotValueProperty{
//   													interpretedValue: jsii.String("interpretedValue"),
//   												},
//   												values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								sessionAttributes: []interface{}{
//   									&sessionAttributeProperty{
//   										key: jsii.String("key"),
//
//   										// the properties below are optional
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							response: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					defaultBranch: &defaultConditionalBranchProperty{
//   						nextStep: &dialogStateProperty{
//   							dialogAction: &dialogActionProperty{
//   								type: jsii.String("type"),
//
//   								// the properties below are optional
//   								slotToElicit: jsii.String("slotToElicit"),
//   								suppressNextMessage: jsii.Boolean(false),
//   							},
//   							intent: &intentOverrideProperty{
//   								name: jsii.String("name"),
//   								slots: []interface{}{
//   									&slotValueOverrideMapProperty{
//   										slotName: jsii.String("slotName"),
//   										slotValueOverride: &slotValueOverrideProperty{
//   											shape: jsii.String("shape"),
//   											value: &slotValueProperty{
//   												interpretedValue: jsii.String("interpretedValue"),
//   											},
//   											values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							sessionAttributes: []interface{}{
//   								&sessionAttributeProperty{
//   									key: jsii.String("key"),
//
//   									// the properties below are optional
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						response: &responseSpecificationProperty{
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   					isActive: jsii.Boolean(false),
//   				},
//   				confirmationNextStep: &dialogStateProperty{
//   					dialogAction: &dialogActionProperty{
//   						type: jsii.String("type"),
//
//   						// the properties below are optional
//   						slotToElicit: jsii.String("slotToElicit"),
//   						suppressNextMessage: jsii.Boolean(false),
//   					},
//   					intent: &intentOverrideProperty{
//   						name: jsii.String("name"),
//   						slots: []interface{}{
//   							&slotValueOverrideMapProperty{
//   								slotName: jsii.String("slotName"),
//   								slotValueOverride: &slotValueOverrideProperty{
//   									shape: jsii.String("shape"),
//   									value: &slotValueProperty{
//   										interpretedValue: jsii.String("interpretedValue"),
//   									},
//   									values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					sessionAttributes: []interface{}{
//   						&sessionAttributeProperty{
//   							key: jsii.String("key"),
//
//   							// the properties below are optional
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				confirmationResponse: &responseSpecificationProperty{
//   					messageGroupsList: []interface{}{
//   						&messageGroupProperty{
//   							message: &messageProperty{
//   								customPayload: &customPayloadProperty{
//   									value: jsii.String("value"),
//   								},
//   								imageResponseCard: &imageResponseCardProperty{
//   									title: jsii.String("title"),
//
//   									// the properties below are optional
//   									buttons: []interface{}{
//   										&buttonProperty{
//   											text: jsii.String("text"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   									imageUrl: jsii.String("imageUrl"),
//   									subtitle: jsii.String("subtitle"),
//   								},
//   								plainTextMessage: &plainTextMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   								ssmlMessage: &sSMLMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							variations: []interface{}{
//   								&messageProperty{
//   									customPayload: &customPayloadProperty{
//   										value: jsii.String("value"),
//   									},
//   									imageResponseCard: &imageResponseCardProperty{
//   										title: jsii.String("title"),
//
//   										// the properties below are optional
//   										buttons: []interface{}{
//   											&buttonProperty{
//   												text: jsii.String("text"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   										imageUrl: jsii.String("imageUrl"),
//   										subtitle: jsii.String("subtitle"),
//   									},
//   									plainTextMessage: &plainTextMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   									ssmlMessage: &sSMLMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					allowInterrupt: jsii.Boolean(false),
//   				},
//   				declinationConditional: &conditionalSpecificationProperty{
//   					conditionalBranches: []interface{}{
//   						&conditionalBranchProperty{
//   							condition: &conditionProperty{
//   								expressionString: jsii.String("expressionString"),
//   							},
//   							name: jsii.String("name"),
//   							nextStep: &dialogStateProperty{
//   								dialogAction: &dialogActionProperty{
//   									type: jsii.String("type"),
//
//   									// the properties below are optional
//   									slotToElicit: jsii.String("slotToElicit"),
//   									suppressNextMessage: jsii.Boolean(false),
//   								},
//   								intent: &intentOverrideProperty{
//   									name: jsii.String("name"),
//   									slots: []interface{}{
//   										&slotValueOverrideMapProperty{
//   											slotName: jsii.String("slotName"),
//   											slotValueOverride: &slotValueOverrideProperty{
//   												shape: jsii.String("shape"),
//   												value: &slotValueProperty{
//   													interpretedValue: jsii.String("interpretedValue"),
//   												},
//   												values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								sessionAttributes: []interface{}{
//   									&sessionAttributeProperty{
//   										key: jsii.String("key"),
//
//   										// the properties below are optional
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							response: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					defaultBranch: &defaultConditionalBranchProperty{
//   						nextStep: &dialogStateProperty{
//   							dialogAction: &dialogActionProperty{
//   								type: jsii.String("type"),
//
//   								// the properties below are optional
//   								slotToElicit: jsii.String("slotToElicit"),
//   								suppressNextMessage: jsii.Boolean(false),
//   							},
//   							intent: &intentOverrideProperty{
//   								name: jsii.String("name"),
//   								slots: []interface{}{
//   									&slotValueOverrideMapProperty{
//   										slotName: jsii.String("slotName"),
//   										slotValueOverride: &slotValueOverrideProperty{
//   											shape: jsii.String("shape"),
//   											value: &slotValueProperty{
//   												interpretedValue: jsii.String("interpretedValue"),
//   											},
//   											values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							sessionAttributes: []interface{}{
//   								&sessionAttributeProperty{
//   									key: jsii.String("key"),
//
//   									// the properties below are optional
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						response: &responseSpecificationProperty{
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   					isActive: jsii.Boolean(false),
//   				},
//   				declinationNextStep: &dialogStateProperty{
//   					dialogAction: &dialogActionProperty{
//   						type: jsii.String("type"),
//
//   						// the properties below are optional
//   						slotToElicit: jsii.String("slotToElicit"),
//   						suppressNextMessage: jsii.Boolean(false),
//   					},
//   					intent: &intentOverrideProperty{
//   						name: jsii.String("name"),
//   						slots: []interface{}{
//   							&slotValueOverrideMapProperty{
//   								slotName: jsii.String("slotName"),
//   								slotValueOverride: &slotValueOverrideProperty{
//   									shape: jsii.String("shape"),
//   									value: &slotValueProperty{
//   										interpretedValue: jsii.String("interpretedValue"),
//   									},
//   									values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					sessionAttributes: []interface{}{
//   						&sessionAttributeProperty{
//   							key: jsii.String("key"),
//
//   							// the properties below are optional
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				declinationResponse: &responseSpecificationProperty{
//   					messageGroupsList: []interface{}{
//   						&messageGroupProperty{
//   							message: &messageProperty{
//   								customPayload: &customPayloadProperty{
//   									value: jsii.String("value"),
//   								},
//   								imageResponseCard: &imageResponseCardProperty{
//   									title: jsii.String("title"),
//
//   									// the properties below are optional
//   									buttons: []interface{}{
//   										&buttonProperty{
//   											text: jsii.String("text"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   									imageUrl: jsii.String("imageUrl"),
//   									subtitle: jsii.String("subtitle"),
//   								},
//   								plainTextMessage: &plainTextMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   								ssmlMessage: &sSMLMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							variations: []interface{}{
//   								&messageProperty{
//   									customPayload: &customPayloadProperty{
//   										value: jsii.String("value"),
//   									},
//   									imageResponseCard: &imageResponseCardProperty{
//   										title: jsii.String("title"),
//
//   										// the properties below are optional
//   										buttons: []interface{}{
//   											&buttonProperty{
//   												text: jsii.String("text"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   										imageUrl: jsii.String("imageUrl"),
//   										subtitle: jsii.String("subtitle"),
//   									},
//   									plainTextMessage: &plainTextMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   									ssmlMessage: &sSMLMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					allowInterrupt: jsii.Boolean(false),
//   				},
//   				elicitationCodeHook: &elicitationCodeHookInvocationSettingProperty{
//   					enableCodeHookInvocation: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					invocationLabel: jsii.String("invocationLabel"),
//   				},
//   				failureConditional: &conditionalSpecificationProperty{
//   					conditionalBranches: []interface{}{
//   						&conditionalBranchProperty{
//   							condition: &conditionProperty{
//   								expressionString: jsii.String("expressionString"),
//   							},
//   							name: jsii.String("name"),
//   							nextStep: &dialogStateProperty{
//   								dialogAction: &dialogActionProperty{
//   									type: jsii.String("type"),
//
//   									// the properties below are optional
//   									slotToElicit: jsii.String("slotToElicit"),
//   									suppressNextMessage: jsii.Boolean(false),
//   								},
//   								intent: &intentOverrideProperty{
//   									name: jsii.String("name"),
//   									slots: []interface{}{
//   										&slotValueOverrideMapProperty{
//   											slotName: jsii.String("slotName"),
//   											slotValueOverride: &slotValueOverrideProperty{
//   												shape: jsii.String("shape"),
//   												value: &slotValueProperty{
//   													interpretedValue: jsii.String("interpretedValue"),
//   												},
//   												values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								sessionAttributes: []interface{}{
//   									&sessionAttributeProperty{
//   										key: jsii.String("key"),
//
//   										// the properties below are optional
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							response: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//   					defaultBranch: &defaultConditionalBranchProperty{
//   						nextStep: &dialogStateProperty{
//   							dialogAction: &dialogActionProperty{
//   								type: jsii.String("type"),
//
//   								// the properties below are optional
//   								slotToElicit: jsii.String("slotToElicit"),
//   								suppressNextMessage: jsii.Boolean(false),
//   							},
//   							intent: &intentOverrideProperty{
//   								name: jsii.String("name"),
//   								slots: []interface{}{
//   									&slotValueOverrideMapProperty{
//   										slotName: jsii.String("slotName"),
//   										slotValueOverride: &slotValueOverrideProperty{
//   											shape: jsii.String("shape"),
//   											value: &slotValueProperty{
//   												interpretedValue: jsii.String("interpretedValue"),
//   											},
//   											values: []interface{}{
//   												slotValueOverrideProperty_,
//   											},
//   										},
//   									},
//   								},
//   							},
//   							sessionAttributes: []interface{}{
//   								&sessionAttributeProperty{
//   									key: jsii.String("key"),
//
//   									// the properties below are optional
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   						response: &responseSpecificationProperty{
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   						},
//   					},
//   					isActive: jsii.Boolean(false),
//   				},
//   				failureNextStep: &dialogStateProperty{
//   					dialogAction: &dialogActionProperty{
//   						type: jsii.String("type"),
//
//   						// the properties below are optional
//   						slotToElicit: jsii.String("slotToElicit"),
//   						suppressNextMessage: jsii.Boolean(false),
//   					},
//   					intent: &intentOverrideProperty{
//   						name: jsii.String("name"),
//   						slots: []interface{}{
//   							&slotValueOverrideMapProperty{
//   								slotName: jsii.String("slotName"),
//   								slotValueOverride: &slotValueOverrideProperty{
//   									shape: jsii.String("shape"),
//   									value: &slotValueProperty{
//   										interpretedValue: jsii.String("interpretedValue"),
//   									},
//   									values: []interface{}{
//   										slotValueOverrideProperty_,
//   									},
//   								},
//   							},
//   						},
//   					},
//   					sessionAttributes: []interface{}{
//   						&sessionAttributeProperty{
//   							key: jsii.String("key"),
//
//   							// the properties below are optional
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				failureResponse: &responseSpecificationProperty{
//   					messageGroupsList: []interface{}{
//   						&messageGroupProperty{
//   							message: &messageProperty{
//   								customPayload: &customPayloadProperty{
//   									value: jsii.String("value"),
//   								},
//   								imageResponseCard: &imageResponseCardProperty{
//   									title: jsii.String("title"),
//
//   									// the properties below are optional
//   									buttons: []interface{}{
//   										&buttonProperty{
//   											text: jsii.String("text"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   									imageUrl: jsii.String("imageUrl"),
//   									subtitle: jsii.String("subtitle"),
//   								},
//   								plainTextMessage: &plainTextMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   								ssmlMessage: &sSMLMessageProperty{
//   									value: jsii.String("value"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							variations: []interface{}{
//   								&messageProperty{
//   									customPayload: &customPayloadProperty{
//   										value: jsii.String("value"),
//   									},
//   									imageResponseCard: &imageResponseCardProperty{
//   										title: jsii.String("title"),
//
//   										// the properties below are optional
//   										buttons: []interface{}{
//   											&buttonProperty{
//   												text: jsii.String("text"),
//   												value: jsii.String("value"),
//   											},
//   										},
//   										imageUrl: jsii.String("imageUrl"),
//   										subtitle: jsii.String("subtitle"),
//   									},
//   									plainTextMessage: &plainTextMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   									ssmlMessage: &sSMLMessageProperty{
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					allowInterrupt: jsii.Boolean(false),
//   				},
//   				isActive: jsii.Boolean(false),
//   			},
//   			kendraConfiguration: &kendraConfigurationProperty{
//   				kendraIndex: jsii.String("kendraIndex"),
//
//   				// the properties below are optional
//   				queryFilterString: jsii.String("queryFilterString"),
//   				queryFilterStringEnabled: jsii.Boolean(false),
//   			},
//   			outputContexts: []interface{}{
//   				&outputContextProperty{
//   					name: jsii.String("name"),
//   					timeToLiveInSeconds: jsii.Number(123),
//   					turnsToLive: jsii.Number(123),
//   				},
//   			},
//   			parentIntentSignature: jsii.String("parentIntentSignature"),
//   			sampleUtterances: []interface{}{
//   				&sampleUtteranceProperty{
//   					utterance: jsii.String("utterance"),
//   				},
//   			},
//   			slotPriorities: []interface{}{
//   				&slotPriorityProperty{
//   					priority: jsii.Number(123),
//   					slotName: jsii.String("slotName"),
//   				},
//   			},
//   			slots: []interface{}{
//   				&slotProperty{
//   					name: jsii.String("name"),
//   					slotTypeName: jsii.String("slotTypeName"),
//   					valueElicitationSetting: &slotValueElicitationSettingProperty{
//   						slotConstraint: jsii.String("slotConstraint"),
//
//   						// the properties below are optional
//   						defaultValueSpecification: &slotDefaultValueSpecificationProperty{
//   							defaultValueList: []interface{}{
//   								&slotDefaultValueProperty{
//   									defaultValue: jsii.String("defaultValue"),
//   								},
//   							},
//   						},
//   						promptSpecification: &promptSpecificationProperty{
//   							maxRetries: jsii.Number(123),
//   							messageGroupsList: []interface{}{
//   								&messageGroupProperty{
//   									message: &messageProperty{
//   										customPayload: &customPayloadProperty{
//   											value: jsii.String("value"),
//   										},
//   										imageResponseCard: &imageResponseCardProperty{
//   											title: jsii.String("title"),
//
//   											// the properties below are optional
//   											buttons: []interface{}{
//   												&buttonProperty{
//   													text: jsii.String("text"),
//   													value: jsii.String("value"),
//   												},
//   											},
//   											imageUrl: jsii.String("imageUrl"),
//   											subtitle: jsii.String("subtitle"),
//   										},
//   										plainTextMessage: &plainTextMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   										ssmlMessage: &sSMLMessageProperty{
//   											value: jsii.String("value"),
//   										},
//   									},
//
//   									// the properties below are optional
//   									variations: []interface{}{
//   										&messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							allowInterrupt: jsii.Boolean(false),
//   							messageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   							promptAttemptsSpecification: map[string]interface{}{
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
//   						sampleUtterances: []interface{}{
//   							&sampleUtteranceProperty{
//   								utterance: jsii.String("utterance"),
//   							},
//   						},
//   						slotCaptureSetting: &slotCaptureSettingProperty{
//   							captureConditional: &conditionalSpecificationProperty{
//   								conditionalBranches: []interface{}{
//   									&conditionalBranchProperty{
//   										condition: &conditionProperty{
//   											expressionString: jsii.String("expressionString"),
//   										},
//   										name: jsii.String("name"),
//   										nextStep: &dialogStateProperty{
//   											dialogAction: &dialogActionProperty{
//   												type: jsii.String("type"),
//
//   												// the properties below are optional
//   												slotToElicit: jsii.String("slotToElicit"),
//   												suppressNextMessage: jsii.Boolean(false),
//   											},
//   											intent: &intentOverrideProperty{
//   												name: jsii.String("name"),
//   												slots: []interface{}{
//   													&slotValueOverrideMapProperty{
//   														slotName: jsii.String("slotName"),
//   														slotValueOverride: &slotValueOverrideProperty{
//   															shape: jsii.String("shape"),
//   															value: &slotValueProperty{
//   																interpretedValue: jsii.String("interpretedValue"),
//   															},
//   															values: []interface{}{
//   																slotValueOverrideProperty_,
//   															},
//   														},
//   													},
//   												},
//   											},
//   											sessionAttributes: []interface{}{
//   												&sessionAttributeProperty{
//   													key: jsii.String("key"),
//
//   													// the properties below are optional
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										response: &responseSpecificationProperty{
//   											messageGroupsList: []interface{}{
//   												&messageGroupProperty{
//   													message: &messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//
//   													// the properties below are optional
//   													variations: []interface{}{
//   														&messageProperty{
//   															customPayload: &customPayloadProperty{
//   																value: jsii.String("value"),
//   															},
//   															imageResponseCard: &imageResponseCardProperty{
//   																title: jsii.String("title"),
//
//   																// the properties below are optional
//   																buttons: []interface{}{
//   																	&buttonProperty{
//   																		text: jsii.String("text"),
//   																		value: jsii.String("value"),
//   																	},
//   																},
//   																imageUrl: jsii.String("imageUrl"),
//   																subtitle: jsii.String("subtitle"),
//   															},
//   															plainTextMessage: &plainTextMessageProperty{
//   																value: jsii.String("value"),
//   															},
//   															ssmlMessage: &sSMLMessageProperty{
//   																value: jsii.String("value"),
//   															},
//   														},
//   													},
//   												},
//   											},
//
//   											// the properties below are optional
//   											allowInterrupt: jsii.Boolean(false),
//   										},
//   									},
//   								},
//   								defaultBranch: &defaultConditionalBranchProperty{
//   									nextStep: &dialogStateProperty{
//   										dialogAction: &dialogActionProperty{
//   											type: jsii.String("type"),
//
//   											// the properties below are optional
//   											slotToElicit: jsii.String("slotToElicit"),
//   											suppressNextMessage: jsii.Boolean(false),
//   										},
//   										intent: &intentOverrideProperty{
//   											name: jsii.String("name"),
//   											slots: []interface{}{
//   												&slotValueOverrideMapProperty{
//   													slotName: jsii.String("slotName"),
//   													slotValueOverride: &slotValueOverrideProperty{
//   														shape: jsii.String("shape"),
//   														value: &slotValueProperty{
//   															interpretedValue: jsii.String("interpretedValue"),
//   														},
//   														values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										sessionAttributes: []interface{}{
//   											&sessionAttributeProperty{
//   												key: jsii.String("key"),
//
//   												// the properties below are optional
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   									response: &responseSpecificationProperty{
//   										messageGroupsList: []interface{}{
//   											&messageGroupProperty{
//   												message: &messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												variations: []interface{}{
//   													&messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   								isActive: jsii.Boolean(false),
//   							},
//   							captureNextStep: &dialogStateProperty{
//   								dialogAction: &dialogActionProperty{
//   									type: jsii.String("type"),
//
//   									// the properties below are optional
//   									slotToElicit: jsii.String("slotToElicit"),
//   									suppressNextMessage: jsii.Boolean(false),
//   								},
//   								intent: &intentOverrideProperty{
//   									name: jsii.String("name"),
//   									slots: []interface{}{
//   										&slotValueOverrideMapProperty{
//   											slotName: jsii.String("slotName"),
//   											slotValueOverride: &slotValueOverrideProperty{
//   												shape: jsii.String("shape"),
//   												value: &slotValueProperty{
//   													interpretedValue: jsii.String("interpretedValue"),
//   												},
//   												values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								sessionAttributes: []interface{}{
//   									&sessionAttributeProperty{
//   										key: jsii.String("key"),
//
//   										// the properties below are optional
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							captureResponse: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   							codeHook: &dialogCodeHookInvocationSettingProperty{
//   								enableCodeHookInvocation: jsii.Boolean(false),
//   								isActive: jsii.Boolean(false),
//   								postCodeHookSpecification: &postDialogCodeHookInvocationSpecificationProperty{
//   									failureConditional: &conditionalSpecificationProperty{
//   										conditionalBranches: []interface{}{
//   											&conditionalBranchProperty{
//   												condition: &conditionProperty{
//   													expressionString: jsii.String("expressionString"),
//   												},
//   												name: jsii.String("name"),
//   												nextStep: &dialogStateProperty{
//   													dialogAction: &dialogActionProperty{
//   														type: jsii.String("type"),
//
//   														// the properties below are optional
//   														slotToElicit: jsii.String("slotToElicit"),
//   														suppressNextMessage: jsii.Boolean(false),
//   													},
//   													intent: &intentOverrideProperty{
//   														name: jsii.String("name"),
//   														slots: []interface{}{
//   															&slotValueOverrideMapProperty{
//   																slotName: jsii.String("slotName"),
//   																slotValueOverride: &slotValueOverrideProperty{
//   																	shape: jsii.String("shape"),
//   																	value: &slotValueProperty{
//   																		interpretedValue: jsii.String("interpretedValue"),
//   																	},
//   																	values: []interface{}{
//   																		slotValueOverrideProperty_,
//   																	},
//   																},
//   															},
//   														},
//   													},
//   													sessionAttributes: []interface{}{
//   														&sessionAttributeProperty{
//   															key: jsii.String("key"),
//
//   															// the properties below are optional
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												response: &responseSpecificationProperty{
//   													messageGroupsList: []interface{}{
//   														&messageGroupProperty{
//   															message: &messageProperty{
//   																customPayload: &customPayloadProperty{
//   																	value: jsii.String("value"),
//   																},
//   																imageResponseCard: &imageResponseCardProperty{
//   																	title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	buttons: []interface{}{
//   																		&buttonProperty{
//   																			text: jsii.String("text"),
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//   																	imageUrl: jsii.String("imageUrl"),
//   																	subtitle: jsii.String("subtitle"),
//   																},
//   																plainTextMessage: &plainTextMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   																ssmlMessage: &sSMLMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   															},
//
//   															// the properties below are optional
//   															variations: []interface{}{
//   																&messageProperty{
//   																	customPayload: &customPayloadProperty{
//   																		value: jsii.String("value"),
//   																	},
//   																	imageResponseCard: &imageResponseCardProperty{
//   																		title: jsii.String("title"),
//
//   																		// the properties below are optional
//   																		buttons: []interface{}{
//   																			&buttonProperty{
//   																				text: jsii.String("text"),
//   																				value: jsii.String("value"),
//   																			},
//   																		},
//   																		imageUrl: jsii.String("imageUrl"),
//   																		subtitle: jsii.String("subtitle"),
//   																	},
//   																	plainTextMessage: &plainTextMessageProperty{
//   																		value: jsii.String("value"),
//   																	},
//   																	ssmlMessage: &sSMLMessageProperty{
//   																		value: jsii.String("value"),
//   																	},
//   																},
//   															},
//   														},
//   													},
//
//   													// the properties below are optional
//   													allowInterrupt: jsii.Boolean(false),
//   												},
//   											},
//   										},
//   										defaultBranch: &defaultConditionalBranchProperty{
//   											nextStep: &dialogStateProperty{
//   												dialogAction: &dialogActionProperty{
//   													type: jsii.String("type"),
//
//   													// the properties below are optional
//   													slotToElicit: jsii.String("slotToElicit"),
//   													suppressNextMessage: jsii.Boolean(false),
//   												},
//   												intent: &intentOverrideProperty{
//   													name: jsii.String("name"),
//   													slots: []interface{}{
//   														&slotValueOverrideMapProperty{
//   															slotName: jsii.String("slotName"),
//   															slotValueOverride: &slotValueOverrideProperty{
//   																shape: jsii.String("shape"),
//   																value: &slotValueProperty{
//   																	interpretedValue: jsii.String("interpretedValue"),
//   																},
//   																values: []interface{}{
//   																	slotValueOverrideProperty_,
//   																},
//   															},
//   														},
//   													},
//   												},
//   												sessionAttributes: []interface{}{
//   													&sessionAttributeProperty{
//   														key: jsii.String("key"),
//
//   														// the properties below are optional
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   											response: &responseSpecificationProperty{
//   												messageGroupsList: []interface{}{
//   													&messageGroupProperty{
//   														message: &messageProperty{
//   															customPayload: &customPayloadProperty{
//   																value: jsii.String("value"),
//   															},
//   															imageResponseCard: &imageResponseCardProperty{
//   																title: jsii.String("title"),
//
//   																// the properties below are optional
//   																buttons: []interface{}{
//   																	&buttonProperty{
//   																		text: jsii.String("text"),
//   																		value: jsii.String("value"),
//   																	},
//   																},
//   																imageUrl: jsii.String("imageUrl"),
//   																subtitle: jsii.String("subtitle"),
//   															},
//   															plainTextMessage: &plainTextMessageProperty{
//   																value: jsii.String("value"),
//   															},
//   															ssmlMessage: &sSMLMessageProperty{
//   																value: jsii.String("value"),
//   															},
//   														},
//
//   														// the properties below are optional
//   														variations: []interface{}{
//   															&messageProperty{
//   																customPayload: &customPayloadProperty{
//   																	value: jsii.String("value"),
//   																},
//   																imageResponseCard: &imageResponseCardProperty{
//   																	title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	buttons: []interface{}{
//   																		&buttonProperty{
//   																			text: jsii.String("text"),
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//   																	imageUrl: jsii.String("imageUrl"),
//   																	subtitle: jsii.String("subtitle"),
//   																},
//   																plainTextMessage: &plainTextMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   																ssmlMessage: &sSMLMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   															},
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												allowInterrupt: jsii.Boolean(false),
//   											},
//   										},
//   										isActive: jsii.Boolean(false),
//   									},
//   									failureNextStep: &dialogStateProperty{
//   										dialogAction: &dialogActionProperty{
//   											type: jsii.String("type"),
//
//   											// the properties below are optional
//   											slotToElicit: jsii.String("slotToElicit"),
//   											suppressNextMessage: jsii.Boolean(false),
//   										},
//   										intent: &intentOverrideProperty{
//   											name: jsii.String("name"),
//   											slots: []interface{}{
//   												&slotValueOverrideMapProperty{
//   													slotName: jsii.String("slotName"),
//   													slotValueOverride: &slotValueOverrideProperty{
//   														shape: jsii.String("shape"),
//   														value: &slotValueProperty{
//   															interpretedValue: jsii.String("interpretedValue"),
//   														},
//   														values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										sessionAttributes: []interface{}{
//   											&sessionAttributeProperty{
//   												key: jsii.String("key"),
//
//   												// the properties below are optional
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   									failureResponse: &responseSpecificationProperty{
//   										messageGroupsList: []interface{}{
//   											&messageGroupProperty{
//   												message: &messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												variations: []interface{}{
//   													&messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   									successConditional: &conditionalSpecificationProperty{
//   										conditionalBranches: []interface{}{
//   											&conditionalBranchProperty{
//   												condition: &conditionProperty{
//   													expressionString: jsii.String("expressionString"),
//   												},
//   												name: jsii.String("name"),
//   												nextStep: &dialogStateProperty{
//   													dialogAction: &dialogActionProperty{
//   														type: jsii.String("type"),
//
//   														// the properties below are optional
//   														slotToElicit: jsii.String("slotToElicit"),
//   														suppressNextMessage: jsii.Boolean(false),
//   													},
//   													intent: &intentOverrideProperty{
//   														name: jsii.String("name"),
//   														slots: []interface{}{
//   															&slotValueOverrideMapProperty{
//   																slotName: jsii.String("slotName"),
//   																slotValueOverride: &slotValueOverrideProperty{
//   																	shape: jsii.String("shape"),
//   																	value: &slotValueProperty{
//   																		interpretedValue: jsii.String("interpretedValue"),
//   																	},
//   																	values: []interface{}{
//   																		slotValueOverrideProperty_,
//   																	},
//   																},
//   															},
//   														},
//   													},
//   													sessionAttributes: []interface{}{
//   														&sessionAttributeProperty{
//   															key: jsii.String("key"),
//
//   															// the properties below are optional
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												response: &responseSpecificationProperty{
//   													messageGroupsList: []interface{}{
//   														&messageGroupProperty{
//   															message: &messageProperty{
//   																customPayload: &customPayloadProperty{
//   																	value: jsii.String("value"),
//   																},
//   																imageResponseCard: &imageResponseCardProperty{
//   																	title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	buttons: []interface{}{
//   																		&buttonProperty{
//   																			text: jsii.String("text"),
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//   																	imageUrl: jsii.String("imageUrl"),
//   																	subtitle: jsii.String("subtitle"),
//   																},
//   																plainTextMessage: &plainTextMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   																ssmlMessage: &sSMLMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   															},
//
//   															// the properties below are optional
//   															variations: []interface{}{
//   																&messageProperty{
//   																	customPayload: &customPayloadProperty{
//   																		value: jsii.String("value"),
//   																	},
//   																	imageResponseCard: &imageResponseCardProperty{
//   																		title: jsii.String("title"),
//
//   																		// the properties below are optional
//   																		buttons: []interface{}{
//   																			&buttonProperty{
//   																				text: jsii.String("text"),
//   																				value: jsii.String("value"),
//   																			},
//   																		},
//   																		imageUrl: jsii.String("imageUrl"),
//   																		subtitle: jsii.String("subtitle"),
//   																	},
//   																	plainTextMessage: &plainTextMessageProperty{
//   																		value: jsii.String("value"),
//   																	},
//   																	ssmlMessage: &sSMLMessageProperty{
//   																		value: jsii.String("value"),
//   																	},
//   																},
//   															},
//   														},
//   													},
//
//   													// the properties below are optional
//   													allowInterrupt: jsii.Boolean(false),
//   												},
//   											},
//   										},
//   										defaultBranch: &defaultConditionalBranchProperty{
//   											nextStep: &dialogStateProperty{
//   												dialogAction: &dialogActionProperty{
//   													type: jsii.String("type"),
//
//   													// the properties below are optional
//   													slotToElicit: jsii.String("slotToElicit"),
//   													suppressNextMessage: jsii.Boolean(false),
//   												},
//   												intent: &intentOverrideProperty{
//   													name: jsii.String("name"),
//   													slots: []interface{}{
//   														&slotValueOverrideMapProperty{
//   															slotName: jsii.String("slotName"),
//   															slotValueOverride: &slotValueOverrideProperty{
//   																shape: jsii.String("shape"),
//   																value: &slotValueProperty{
//   																	interpretedValue: jsii.String("interpretedValue"),
//   																},
//   																values: []interface{}{
//   																	slotValueOverrideProperty_,
//   																},
//   															},
//   														},
//   													},
//   												},
//   												sessionAttributes: []interface{}{
//   													&sessionAttributeProperty{
//   														key: jsii.String("key"),
//
//   														// the properties below are optional
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   											response: &responseSpecificationProperty{
//   												messageGroupsList: []interface{}{
//   													&messageGroupProperty{
//   														message: &messageProperty{
//   															customPayload: &customPayloadProperty{
//   																value: jsii.String("value"),
//   															},
//   															imageResponseCard: &imageResponseCardProperty{
//   																title: jsii.String("title"),
//
//   																// the properties below are optional
//   																buttons: []interface{}{
//   																	&buttonProperty{
//   																		text: jsii.String("text"),
//   																		value: jsii.String("value"),
//   																	},
//   																},
//   																imageUrl: jsii.String("imageUrl"),
//   																subtitle: jsii.String("subtitle"),
//   															},
//   															plainTextMessage: &plainTextMessageProperty{
//   																value: jsii.String("value"),
//   															},
//   															ssmlMessage: &sSMLMessageProperty{
//   																value: jsii.String("value"),
//   															},
//   														},
//
//   														// the properties below are optional
//   														variations: []interface{}{
//   															&messageProperty{
//   																customPayload: &customPayloadProperty{
//   																	value: jsii.String("value"),
//   																},
//   																imageResponseCard: &imageResponseCardProperty{
//   																	title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	buttons: []interface{}{
//   																		&buttonProperty{
//   																			text: jsii.String("text"),
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//   																	imageUrl: jsii.String("imageUrl"),
//   																	subtitle: jsii.String("subtitle"),
//   																},
//   																plainTextMessage: &plainTextMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   																ssmlMessage: &sSMLMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   															},
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												allowInterrupt: jsii.Boolean(false),
//   											},
//   										},
//   										isActive: jsii.Boolean(false),
//   									},
//   									successNextStep: &dialogStateProperty{
//   										dialogAction: &dialogActionProperty{
//   											type: jsii.String("type"),
//
//   											// the properties below are optional
//   											slotToElicit: jsii.String("slotToElicit"),
//   											suppressNextMessage: jsii.Boolean(false),
//   										},
//   										intent: &intentOverrideProperty{
//   											name: jsii.String("name"),
//   											slots: []interface{}{
//   												&slotValueOverrideMapProperty{
//   													slotName: jsii.String("slotName"),
//   													slotValueOverride: &slotValueOverrideProperty{
//   														shape: jsii.String("shape"),
//   														value: &slotValueProperty{
//   															interpretedValue: jsii.String("interpretedValue"),
//   														},
//   														values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										sessionAttributes: []interface{}{
//   											&sessionAttributeProperty{
//   												key: jsii.String("key"),
//
//   												// the properties below are optional
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   									successResponse: &responseSpecificationProperty{
//   										messageGroupsList: []interface{}{
//   											&messageGroupProperty{
//   												message: &messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												variations: []interface{}{
//   													&messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   									timeoutConditional: &conditionalSpecificationProperty{
//   										conditionalBranches: []interface{}{
//   											&conditionalBranchProperty{
//   												condition: &conditionProperty{
//   													expressionString: jsii.String("expressionString"),
//   												},
//   												name: jsii.String("name"),
//   												nextStep: &dialogStateProperty{
//   													dialogAction: &dialogActionProperty{
//   														type: jsii.String("type"),
//
//   														// the properties below are optional
//   														slotToElicit: jsii.String("slotToElicit"),
//   														suppressNextMessage: jsii.Boolean(false),
//   													},
//   													intent: &intentOverrideProperty{
//   														name: jsii.String("name"),
//   														slots: []interface{}{
//   															&slotValueOverrideMapProperty{
//   																slotName: jsii.String("slotName"),
//   																slotValueOverride: &slotValueOverrideProperty{
//   																	shape: jsii.String("shape"),
//   																	value: &slotValueProperty{
//   																		interpretedValue: jsii.String("interpretedValue"),
//   																	},
//   																	values: []interface{}{
//   																		slotValueOverrideProperty_,
//   																	},
//   																},
//   															},
//   														},
//   													},
//   													sessionAttributes: []interface{}{
//   														&sessionAttributeProperty{
//   															key: jsii.String("key"),
//
//   															// the properties below are optional
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												response: &responseSpecificationProperty{
//   													messageGroupsList: []interface{}{
//   														&messageGroupProperty{
//   															message: &messageProperty{
//   																customPayload: &customPayloadProperty{
//   																	value: jsii.String("value"),
//   																},
//   																imageResponseCard: &imageResponseCardProperty{
//   																	title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	buttons: []interface{}{
//   																		&buttonProperty{
//   																			text: jsii.String("text"),
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//   																	imageUrl: jsii.String("imageUrl"),
//   																	subtitle: jsii.String("subtitle"),
//   																},
//   																plainTextMessage: &plainTextMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   																ssmlMessage: &sSMLMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   															},
//
//   															// the properties below are optional
//   															variations: []interface{}{
//   																&messageProperty{
//   																	customPayload: &customPayloadProperty{
//   																		value: jsii.String("value"),
//   																	},
//   																	imageResponseCard: &imageResponseCardProperty{
//   																		title: jsii.String("title"),
//
//   																		// the properties below are optional
//   																		buttons: []interface{}{
//   																			&buttonProperty{
//   																				text: jsii.String("text"),
//   																				value: jsii.String("value"),
//   																			},
//   																		},
//   																		imageUrl: jsii.String("imageUrl"),
//   																		subtitle: jsii.String("subtitle"),
//   																	},
//   																	plainTextMessage: &plainTextMessageProperty{
//   																		value: jsii.String("value"),
//   																	},
//   																	ssmlMessage: &sSMLMessageProperty{
//   																		value: jsii.String("value"),
//   																	},
//   																},
//   															},
//   														},
//   													},
//
//   													// the properties below are optional
//   													allowInterrupt: jsii.Boolean(false),
//   												},
//   											},
//   										},
//   										defaultBranch: &defaultConditionalBranchProperty{
//   											nextStep: &dialogStateProperty{
//   												dialogAction: &dialogActionProperty{
//   													type: jsii.String("type"),
//
//   													// the properties below are optional
//   													slotToElicit: jsii.String("slotToElicit"),
//   													suppressNextMessage: jsii.Boolean(false),
//   												},
//   												intent: &intentOverrideProperty{
//   													name: jsii.String("name"),
//   													slots: []interface{}{
//   														&slotValueOverrideMapProperty{
//   															slotName: jsii.String("slotName"),
//   															slotValueOverride: &slotValueOverrideProperty{
//   																shape: jsii.String("shape"),
//   																value: &slotValueProperty{
//   																	interpretedValue: jsii.String("interpretedValue"),
//   																},
//   																values: []interface{}{
//   																	slotValueOverrideProperty_,
//   																},
//   															},
//   														},
//   													},
//   												},
//   												sessionAttributes: []interface{}{
//   													&sessionAttributeProperty{
//   														key: jsii.String("key"),
//
//   														// the properties below are optional
//   														value: jsii.String("value"),
//   													},
//   												},
//   											},
//   											response: &responseSpecificationProperty{
//   												messageGroupsList: []interface{}{
//   													&messageGroupProperty{
//   														message: &messageProperty{
//   															customPayload: &customPayloadProperty{
//   																value: jsii.String("value"),
//   															},
//   															imageResponseCard: &imageResponseCardProperty{
//   																title: jsii.String("title"),
//
//   																// the properties below are optional
//   																buttons: []interface{}{
//   																	&buttonProperty{
//   																		text: jsii.String("text"),
//   																		value: jsii.String("value"),
//   																	},
//   																},
//   																imageUrl: jsii.String("imageUrl"),
//   																subtitle: jsii.String("subtitle"),
//   															},
//   															plainTextMessage: &plainTextMessageProperty{
//   																value: jsii.String("value"),
//   															},
//   															ssmlMessage: &sSMLMessageProperty{
//   																value: jsii.String("value"),
//   															},
//   														},
//
//   														// the properties below are optional
//   														variations: []interface{}{
//   															&messageProperty{
//   																customPayload: &customPayloadProperty{
//   																	value: jsii.String("value"),
//   																},
//   																imageResponseCard: &imageResponseCardProperty{
//   																	title: jsii.String("title"),
//
//   																	// the properties below are optional
//   																	buttons: []interface{}{
//   																		&buttonProperty{
//   																			text: jsii.String("text"),
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//   																	imageUrl: jsii.String("imageUrl"),
//   																	subtitle: jsii.String("subtitle"),
//   																},
//   																plainTextMessage: &plainTextMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   																ssmlMessage: &sSMLMessageProperty{
//   																	value: jsii.String("value"),
//   																},
//   															},
//   														},
//   													},
//   												},
//
//   												// the properties below are optional
//   												allowInterrupt: jsii.Boolean(false),
//   											},
//   										},
//   										isActive: jsii.Boolean(false),
//   									},
//   									timeoutNextStep: &dialogStateProperty{
//   										dialogAction: &dialogActionProperty{
//   											type: jsii.String("type"),
//
//   											// the properties below are optional
//   											slotToElicit: jsii.String("slotToElicit"),
//   											suppressNextMessage: jsii.Boolean(false),
//   										},
//   										intent: &intentOverrideProperty{
//   											name: jsii.String("name"),
//   											slots: []interface{}{
//   												&slotValueOverrideMapProperty{
//   													slotName: jsii.String("slotName"),
//   													slotValueOverride: &slotValueOverrideProperty{
//   														shape: jsii.String("shape"),
//   														value: &slotValueProperty{
//   															interpretedValue: jsii.String("interpretedValue"),
//   														},
//   														values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										sessionAttributes: []interface{}{
//   											&sessionAttributeProperty{
//   												key: jsii.String("key"),
//
//   												// the properties below are optional
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   									timeoutResponse: &responseSpecificationProperty{
//   										messageGroupsList: []interface{}{
//   											&messageGroupProperty{
//   												message: &messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												variations: []interface{}{
//   													&messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//
//   								// the properties below are optional
//   								invocationLabel: jsii.String("invocationLabel"),
//   							},
//   							elicitationCodeHook: &elicitationCodeHookInvocationSettingProperty{
//   								enableCodeHookInvocation: jsii.Boolean(false),
//
//   								// the properties below are optional
//   								invocationLabel: jsii.String("invocationLabel"),
//   							},
//   							failureConditional: &conditionalSpecificationProperty{
//   								conditionalBranches: []interface{}{
//   									&conditionalBranchProperty{
//   										condition: &conditionProperty{
//   											expressionString: jsii.String("expressionString"),
//   										},
//   										name: jsii.String("name"),
//   										nextStep: &dialogStateProperty{
//   											dialogAction: &dialogActionProperty{
//   												type: jsii.String("type"),
//
//   												// the properties below are optional
//   												slotToElicit: jsii.String("slotToElicit"),
//   												suppressNextMessage: jsii.Boolean(false),
//   											},
//   											intent: &intentOverrideProperty{
//   												name: jsii.String("name"),
//   												slots: []interface{}{
//   													&slotValueOverrideMapProperty{
//   														slotName: jsii.String("slotName"),
//   														slotValueOverride: &slotValueOverrideProperty{
//   															shape: jsii.String("shape"),
//   															value: &slotValueProperty{
//   																interpretedValue: jsii.String("interpretedValue"),
//   															},
//   															values: []interface{}{
//   																slotValueOverrideProperty_,
//   															},
//   														},
//   													},
//   												},
//   											},
//   											sessionAttributes: []interface{}{
//   												&sessionAttributeProperty{
//   													key: jsii.String("key"),
//
//   													// the properties below are optional
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										response: &responseSpecificationProperty{
//   											messageGroupsList: []interface{}{
//   												&messageGroupProperty{
//   													message: &messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//
//   													// the properties below are optional
//   													variations: []interface{}{
//   														&messageProperty{
//   															customPayload: &customPayloadProperty{
//   																value: jsii.String("value"),
//   															},
//   															imageResponseCard: &imageResponseCardProperty{
//   																title: jsii.String("title"),
//
//   																// the properties below are optional
//   																buttons: []interface{}{
//   																	&buttonProperty{
//   																		text: jsii.String("text"),
//   																		value: jsii.String("value"),
//   																	},
//   																},
//   																imageUrl: jsii.String("imageUrl"),
//   																subtitle: jsii.String("subtitle"),
//   															},
//   															plainTextMessage: &plainTextMessageProperty{
//   																value: jsii.String("value"),
//   															},
//   															ssmlMessage: &sSMLMessageProperty{
//   																value: jsii.String("value"),
//   															},
//   														},
//   													},
//   												},
//   											},
//
//   											// the properties below are optional
//   											allowInterrupt: jsii.Boolean(false),
//   										},
//   									},
//   								},
//   								defaultBranch: &defaultConditionalBranchProperty{
//   									nextStep: &dialogStateProperty{
//   										dialogAction: &dialogActionProperty{
//   											type: jsii.String("type"),
//
//   											// the properties below are optional
//   											slotToElicit: jsii.String("slotToElicit"),
//   											suppressNextMessage: jsii.Boolean(false),
//   										},
//   										intent: &intentOverrideProperty{
//   											name: jsii.String("name"),
//   											slots: []interface{}{
//   												&slotValueOverrideMapProperty{
//   													slotName: jsii.String("slotName"),
//   													slotValueOverride: &slotValueOverrideProperty{
//   														shape: jsii.String("shape"),
//   														value: &slotValueProperty{
//   															interpretedValue: jsii.String("interpretedValue"),
//   														},
//   														values: []interface{}{
//   															slotValueOverrideProperty_,
//   														},
//   													},
//   												},
//   											},
//   										},
//   										sessionAttributes: []interface{}{
//   											&sessionAttributeProperty{
//   												key: jsii.String("key"),
//
//   												// the properties below are optional
//   												value: jsii.String("value"),
//   											},
//   										},
//   									},
//   									response: &responseSpecificationProperty{
//   										messageGroupsList: []interface{}{
//   											&messageGroupProperty{
//   												message: &messageProperty{
//   													customPayload: &customPayloadProperty{
//   														value: jsii.String("value"),
//   													},
//   													imageResponseCard: &imageResponseCardProperty{
//   														title: jsii.String("title"),
//
//   														// the properties below are optional
//   														buttons: []interface{}{
//   															&buttonProperty{
//   																text: jsii.String("text"),
//   																value: jsii.String("value"),
//   															},
//   														},
//   														imageUrl: jsii.String("imageUrl"),
//   														subtitle: jsii.String("subtitle"),
//   													},
//   													plainTextMessage: &plainTextMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   													ssmlMessage: &sSMLMessageProperty{
//   														value: jsii.String("value"),
//   													},
//   												},
//
//   												// the properties below are optional
//   												variations: []interface{}{
//   													&messageProperty{
//   														customPayload: &customPayloadProperty{
//   															value: jsii.String("value"),
//   														},
//   														imageResponseCard: &imageResponseCardProperty{
//   															title: jsii.String("title"),
//
//   															// the properties below are optional
//   															buttons: []interface{}{
//   																&buttonProperty{
//   																	text: jsii.String("text"),
//   																	value: jsii.String("value"),
//   																},
//   															},
//   															imageUrl: jsii.String("imageUrl"),
//   															subtitle: jsii.String("subtitle"),
//   														},
//   														plainTextMessage: &plainTextMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   														ssmlMessage: &sSMLMessageProperty{
//   															value: jsii.String("value"),
//   														},
//   													},
//   												},
//   											},
//   										},
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   								isActive: jsii.Boolean(false),
//   							},
//   							failureNextStep: &dialogStateProperty{
//   								dialogAction: &dialogActionProperty{
//   									type: jsii.String("type"),
//
//   									// the properties below are optional
//   									slotToElicit: jsii.String("slotToElicit"),
//   									suppressNextMessage: jsii.Boolean(false),
//   								},
//   								intent: &intentOverrideProperty{
//   									name: jsii.String("name"),
//   									slots: []interface{}{
//   										&slotValueOverrideMapProperty{
//   											slotName: jsii.String("slotName"),
//   											slotValueOverride: &slotValueOverrideProperty{
//   												shape: jsii.String("shape"),
//   												value: &slotValueProperty{
//   													interpretedValue: jsii.String("interpretedValue"),
//   												},
//   												values: []interface{}{
//   													slotValueOverrideProperty_,
//   												},
//   											},
//   										},
//   									},
//   								},
//   								sessionAttributes: []interface{}{
//   									&sessionAttributeProperty{
//   										key: jsii.String("key"),
//
//   										// the properties below are optional
//   										value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							failureResponse: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   						waitAndContinueSpecification: &waitAndContinueSpecificationProperty{
//   							continueResponse: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   							waitingResponse: &responseSpecificationProperty{
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//
//   							// the properties below are optional
//   							isActive: jsii.Boolean(false),
//   							stillWaitingResponse: &stillWaitingResponseSpecificationProperty{
//   								frequencyInSeconds: jsii.Number(123),
//   								messageGroupsList: []interface{}{
//   									&messageGroupProperty{
//   										message: &messageProperty{
//   											customPayload: &customPayloadProperty{
//   												value: jsii.String("value"),
//   											},
//   											imageResponseCard: &imageResponseCardProperty{
//   												title: jsii.String("title"),
//
//   												// the properties below are optional
//   												buttons: []interface{}{
//   													&buttonProperty{
//   														text: jsii.String("text"),
//   														value: jsii.String("value"),
//   													},
//   												},
//   												imageUrl: jsii.String("imageUrl"),
//   												subtitle: jsii.String("subtitle"),
//   											},
//   											plainTextMessage: &plainTextMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   											ssmlMessage: &sSMLMessageProperty{
//   												value: jsii.String("value"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										variations: []interface{}{
//   											&messageProperty{
//   												customPayload: &customPayloadProperty{
//   													value: jsii.String("value"),
//   												},
//   												imageResponseCard: &imageResponseCardProperty{
//   													title: jsii.String("title"),
//
//   													// the properties below are optional
//   													buttons: []interface{}{
//   														&buttonProperty{
//   															text: jsii.String("text"),
//   															value: jsii.String("value"),
//   														},
//   													},
//   													imageUrl: jsii.String("imageUrl"),
//   													subtitle: jsii.String("subtitle"),
//   												},
//   												plainTextMessage: &plainTextMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   												ssmlMessage: &sSMLMessageProperty{
//   													value: jsii.String("value"),
//   												},
//   											},
//   										},
//   									},
//   								},
//   								timeoutInSeconds: jsii.Number(123),
//
//   								// the properties below are optional
//   								allowInterrupt: jsii.Boolean(false),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					description: jsii.String("description"),
//   					multipleValuesSetting: &multipleValuesSettingProperty{
//   						allowMultipleValues: jsii.Boolean(false),
//   					},
//   					obfuscationSetting: &obfuscationSettingProperty{
//   						obfuscationSettingType: jsii.String("obfuscationSettingType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	slotTypes: []interface{}{
//   		&slotTypeProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			externalSourceSetting: &externalSourceSettingProperty{
//   				grammarSlotTypeSetting: &grammarSlotTypeSettingProperty{
//   					source: &grammarSlotTypeSourceProperty{
//   						s3BucketName: jsii.String("s3BucketName"),
//   						s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   						// the properties below are optional
//   						kmsKeyArn: jsii.String("kmsKeyArn"),
//   					},
//   				},
//   			},
//   			parentSlotTypeSignature: jsii.String("parentSlotTypeSignature"),
//   			slotTypeValues: []interface{}{
//   				&slotTypeValueProperty{
//   					sampleValue: &sampleValueProperty{
//   						value: jsii.String("value"),
//   					},
//
//   					// the properties below are optional
//   					synonyms: []interface{}{
//   						&sampleValueProperty{
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			valueSelectionSetting: &slotValueSelectionSettingProperty{
//   				resolutionStrategy: jsii.String("resolutionStrategy"),
//
//   				// the properties below are optional
//   				advancedRecognitionSetting: &advancedRecognitionSettingProperty{
//   					audioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   				},
//   				regexFilter: &slotValueRegexFilterProperty{
//   					pattern: jsii.String("pattern"),
//   				},
//   			},
//   		},
//   	},
//   	voiceSettings: &voiceSettingsProperty{
//   		voiceId: jsii.String("voiceId"),
//
//   		// the properties below are optional
//   		engine: jsii.String("engine"),
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

