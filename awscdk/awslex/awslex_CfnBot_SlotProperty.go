package awslex


// Specifies the definition of a slot.
//
// Amazon Lex elicits slot values from uses to fulfill the user's intent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   slotProperty := &slotProperty{
//   	name: jsii.String("name"),
//   	slotTypeName: jsii.String("slotTypeName"),
//   	valueElicitationSetting: &slotValueElicitationSettingProperty{
//   		slotConstraint: jsii.String("slotConstraint"),
//
//   		// the properties below are optional
//   		defaultValueSpecification: &slotDefaultValueSpecificationProperty{
//   			defaultValueList: []interface{}{
//   				&slotDefaultValueProperty{
//   					defaultValue: jsii.String("defaultValue"),
//   				},
//   			},
//   		},
//   		promptSpecification: &promptSpecificationProperty{
//   			maxRetries: jsii.Number(123),
//   			messageGroupsList: []interface{}{
//   				&messageGroupProperty{
//   					message: &messageProperty{
//   						customPayload: &customPayloadProperty{
//   							value: jsii.String("value"),
//   						},
//   						imageResponseCard: &imageResponseCardProperty{
//   							title: jsii.String("title"),
//
//   							// the properties below are optional
//   							buttons: []interface{}{
//   								&buttonProperty{
//   									text: jsii.String("text"),
//   									value: jsii.String("value"),
//   								},
//   							},
//   							imageUrl: jsii.String("imageUrl"),
//   							subtitle: jsii.String("subtitle"),
//   						},
//   						plainTextMessage: &plainTextMessageProperty{
//   							value: jsii.String("value"),
//   						},
//   						ssmlMessage: &sSMLMessageProperty{
//   							value: jsii.String("value"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					variations: []interface{}{
//   						&messageProperty{
//   							customPayload: &customPayloadProperty{
//   								value: jsii.String("value"),
//   							},
//   							imageResponseCard: &imageResponseCardProperty{
//   								title: jsii.String("title"),
//
//   								// the properties below are optional
//   								buttons: []interface{}{
//   									&buttonProperty{
//   										text: jsii.String("text"),
//   										value: jsii.String("value"),
//   									},
//   								},
//   								imageUrl: jsii.String("imageUrl"),
//   								subtitle: jsii.String("subtitle"),
//   							},
//   							plainTextMessage: &plainTextMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   							ssmlMessage: &sSMLMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			allowInterrupt: jsii.Boolean(false),
//   			messageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   			promptAttemptsSpecification: map[string]interface{}{
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
//   		sampleUtterances: []interface{}{
//   			&sampleUtteranceProperty{
//   				utterance: jsii.String("utterance"),
//   			},
//   		},
//   		slotCaptureSetting: &slotCaptureSettingProperty{
//   			captureConditional: &conditionalSpecificationProperty{
//   				conditionalBranches: []interface{}{
//   					&conditionalBranchProperty{
//   						condition: &conditionProperty{
//   							expressionString: jsii.String("expressionString"),
//   						},
//   						name: jsii.String("name"),
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
//
//   						// the properties below are optional
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
//   				},
//   				defaultBranch: &defaultConditionalBranchProperty{
//   					nextStep: &dialogStateProperty{
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
//   					response: &responseSpecificationProperty{
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
//   				isActive: jsii.Boolean(false),
//   			},
//   			captureNextStep: &dialogStateProperty{
//   				dialogAction: &dialogActionProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					slotToElicit: jsii.String("slotToElicit"),
//   					suppressNextMessage: jsii.Boolean(false),
//   				},
//   				intent: &intentOverrideProperty{
//   					name: jsii.String("name"),
//   					slots: []interface{}{
//   						&slotValueOverrideMapProperty{
//   							slotName: jsii.String("slotName"),
//   							slotValueOverride: &slotValueOverrideProperty{
//   								shape: jsii.String("shape"),
//   								value: &slotValueProperty{
//   									interpretedValue: jsii.String("interpretedValue"),
//   								},
//   								values: []interface{}{
//   									slotValueOverrideProperty_,
//   								},
//   							},
//   						},
//   					},
//   				},
//   				sessionAttributes: []interface{}{
//   					&sessionAttributeProperty{
//   						key: jsii.String("key"),
//
//   						// the properties below are optional
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			captureResponse: &responseSpecificationProperty{
//   				messageGroupsList: []interface{}{
//   					&messageGroupProperty{
//   						message: &messageProperty{
//   							customPayload: &customPayloadProperty{
//   								value: jsii.String("value"),
//   							},
//   							imageResponseCard: &imageResponseCardProperty{
//   								title: jsii.String("title"),
//
//   								// the properties below are optional
//   								buttons: []interface{}{
//   									&buttonProperty{
//   										text: jsii.String("text"),
//   										value: jsii.String("value"),
//   									},
//   								},
//   								imageUrl: jsii.String("imageUrl"),
//   								subtitle: jsii.String("subtitle"),
//   							},
//   							plainTextMessage: &plainTextMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   							ssmlMessage: &sSMLMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						variations: []interface{}{
//   							&messageProperty{
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
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				allowInterrupt: jsii.Boolean(false),
//   			},
//   			codeHook: &dialogCodeHookInvocationSettingProperty{
//   				enableCodeHookInvocation: jsii.Boolean(false),
//   				isActive: jsii.Boolean(false),
//   				postCodeHookSpecification: &postDialogCodeHookInvocationSpecificationProperty{
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
//
//   				// the properties below are optional
//   				invocationLabel: jsii.String("invocationLabel"),
//   			},
//   			elicitationCodeHook: &elicitationCodeHookInvocationSettingProperty{
//   				enableCodeHookInvocation: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				invocationLabel: jsii.String("invocationLabel"),
//   			},
//   			failureConditional: &conditionalSpecificationProperty{
//   				conditionalBranches: []interface{}{
//   					&conditionalBranchProperty{
//   						condition: &conditionProperty{
//   							expressionString: jsii.String("expressionString"),
//   						},
//   						name: jsii.String("name"),
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
//
//   						// the properties below are optional
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
//   				},
//   				defaultBranch: &defaultConditionalBranchProperty{
//   					nextStep: &dialogStateProperty{
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
//   					response: &responseSpecificationProperty{
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
//   				isActive: jsii.Boolean(false),
//   			},
//   			failureNextStep: &dialogStateProperty{
//   				dialogAction: &dialogActionProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					slotToElicit: jsii.String("slotToElicit"),
//   					suppressNextMessage: jsii.Boolean(false),
//   				},
//   				intent: &intentOverrideProperty{
//   					name: jsii.String("name"),
//   					slots: []interface{}{
//   						&slotValueOverrideMapProperty{
//   							slotName: jsii.String("slotName"),
//   							slotValueOverride: &slotValueOverrideProperty{
//   								shape: jsii.String("shape"),
//   								value: &slotValueProperty{
//   									interpretedValue: jsii.String("interpretedValue"),
//   								},
//   								values: []interface{}{
//   									slotValueOverrideProperty_,
//   								},
//   							},
//   						},
//   					},
//   				},
//   				sessionAttributes: []interface{}{
//   					&sessionAttributeProperty{
//   						key: jsii.String("key"),
//
//   						// the properties below are optional
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			failureResponse: &responseSpecificationProperty{
//   				messageGroupsList: []interface{}{
//   					&messageGroupProperty{
//   						message: &messageProperty{
//   							customPayload: &customPayloadProperty{
//   								value: jsii.String("value"),
//   							},
//   							imageResponseCard: &imageResponseCardProperty{
//   								title: jsii.String("title"),
//
//   								// the properties below are optional
//   								buttons: []interface{}{
//   									&buttonProperty{
//   										text: jsii.String("text"),
//   										value: jsii.String("value"),
//   									},
//   								},
//   								imageUrl: jsii.String("imageUrl"),
//   								subtitle: jsii.String("subtitle"),
//   							},
//   							plainTextMessage: &plainTextMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   							ssmlMessage: &sSMLMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						variations: []interface{}{
//   							&messageProperty{
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
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				allowInterrupt: jsii.Boolean(false),
//   			},
//   		},
//   		waitAndContinueSpecification: &waitAndContinueSpecificationProperty{
//   			continueResponse: &responseSpecificationProperty{
//   				messageGroupsList: []interface{}{
//   					&messageGroupProperty{
//   						message: &messageProperty{
//   							customPayload: &customPayloadProperty{
//   								value: jsii.String("value"),
//   							},
//   							imageResponseCard: &imageResponseCardProperty{
//   								title: jsii.String("title"),
//
//   								// the properties below are optional
//   								buttons: []interface{}{
//   									&buttonProperty{
//   										text: jsii.String("text"),
//   										value: jsii.String("value"),
//   									},
//   								},
//   								imageUrl: jsii.String("imageUrl"),
//   								subtitle: jsii.String("subtitle"),
//   							},
//   							plainTextMessage: &plainTextMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   							ssmlMessage: &sSMLMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						variations: []interface{}{
//   							&messageProperty{
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
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				allowInterrupt: jsii.Boolean(false),
//   			},
//   			waitingResponse: &responseSpecificationProperty{
//   				messageGroupsList: []interface{}{
//   					&messageGroupProperty{
//   						message: &messageProperty{
//   							customPayload: &customPayloadProperty{
//   								value: jsii.String("value"),
//   							},
//   							imageResponseCard: &imageResponseCardProperty{
//   								title: jsii.String("title"),
//
//   								// the properties below are optional
//   								buttons: []interface{}{
//   									&buttonProperty{
//   										text: jsii.String("text"),
//   										value: jsii.String("value"),
//   									},
//   								},
//   								imageUrl: jsii.String("imageUrl"),
//   								subtitle: jsii.String("subtitle"),
//   							},
//   							plainTextMessage: &plainTextMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   							ssmlMessage: &sSMLMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						variations: []interface{}{
//   							&messageProperty{
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
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				allowInterrupt: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			isActive: jsii.Boolean(false),
//   			stillWaitingResponse: &stillWaitingResponseSpecificationProperty{
//   				frequencyInSeconds: jsii.Number(123),
//   				messageGroupsList: []interface{}{
//   					&messageGroupProperty{
//   						message: &messageProperty{
//   							customPayload: &customPayloadProperty{
//   								value: jsii.String("value"),
//   							},
//   							imageResponseCard: &imageResponseCardProperty{
//   								title: jsii.String("title"),
//
//   								// the properties below are optional
//   								buttons: []interface{}{
//   									&buttonProperty{
//   										text: jsii.String("text"),
//   										value: jsii.String("value"),
//   									},
//   								},
//   								imageUrl: jsii.String("imageUrl"),
//   								subtitle: jsii.String("subtitle"),
//   							},
//   							plainTextMessage: &plainTextMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   							ssmlMessage: &sSMLMessageProperty{
//   								value: jsii.String("value"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						variations: []interface{}{
//   							&messageProperty{
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
//   						},
//   					},
//   				},
//   				timeoutInSeconds: jsii.Number(123),
//
//   				// the properties below are optional
//   				allowInterrupt: jsii.Boolean(false),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	multipleValuesSetting: &multipleValuesSettingProperty{
//   		allowMultipleValues: jsii.Boolean(false),
//   	},
//   	obfuscationSetting: &obfuscationSettingProperty{
//   		obfuscationSettingType: jsii.String("obfuscationSettingType"),
//   	},
//   }
//
type CfnBot_SlotProperty struct {
	// The name given to the slot.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the slot type that this slot is based on.
	//
	// The slot type defines the acceptable values for the slot.
	SlotTypeName *string `field:"required" json:"slotTypeName" yaml:"slotTypeName"`
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - ORIGINAL_VALUE - Returns the value entered by the user, if the user value is similar to a slot value.
	// - TOP_RESOLUTION - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the `valueSelectionStrategy` , the default is `ORIGINAL_VALUE` .
	ValueElicitationSetting interface{} `field:"required" json:"valueElicitationSetting" yaml:"valueElicitationSetting"`
	// The description of the slot.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether a slot can return multiple values.
	MultipleValuesSetting interface{} `field:"optional" json:"multipleValuesSetting" yaml:"multipleValuesSetting"`
	// Determines whether the contents of the slot are obfuscated in Amazon CloudWatch Logs logs.
	//
	// Use obfuscated slots to protect information such as personally identifiable information (PII) in logs.
	ObfuscationSetting interface{} `field:"optional" json:"obfuscationSetting" yaml:"obfuscationSetting"`
}

