package awslex


// Provides a prompt for making sure that the user is ready for the intent to be fulfilled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   intentConfirmationSettingProperty := &intentConfirmationSettingProperty{
//   	promptSpecification: &promptSpecificationProperty{
//   		maxRetries: jsii.Number(123),
//   		messageGroupsList: []interface{}{
//   			&messageGroupProperty{
//   				message: &messageProperty{
//   					customPayload: &customPayloadProperty{
//   						value: jsii.String("value"),
//   					},
//   					imageResponseCard: &imageResponseCardProperty{
//   						title: jsii.String("title"),
//
//   						// the properties below are optional
//   						buttons: []interface{}{
//   							&buttonProperty{
//   								text: jsii.String("text"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						imageUrl: jsii.String("imageUrl"),
//   						subtitle: jsii.String("subtitle"),
//   					},
//   					plainTextMessage: &plainTextMessageProperty{
//   						value: jsii.String("value"),
//   					},
//   					ssmlMessage: &sSMLMessageProperty{
//   						value: jsii.String("value"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				variations: []interface{}{
//   					&messageProperty{
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
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		allowInterrupt: jsii.Boolean(false),
//   		messageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   		promptAttemptsSpecification: map[string]interface{}{
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
//
//   	// the properties below are optional
//   	codeHook: &dialogCodeHookInvocationSettingProperty{
//   		enableCodeHookInvocation: jsii.Boolean(false),
//   		isActive: jsii.Boolean(false),
//   		postCodeHookSpecification: &postDialogCodeHookInvocationSpecificationProperty{
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
//   			successConditional: &conditionalSpecificationProperty{
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
//   			successNextStep: &dialogStateProperty{
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
//   			successResponse: &responseSpecificationProperty{
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
//   			timeoutConditional: &conditionalSpecificationProperty{
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
//   			timeoutNextStep: &dialogStateProperty{
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
//   			timeoutResponse: &responseSpecificationProperty{
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
//
//   		// the properties below are optional
//   		invocationLabel: jsii.String("invocationLabel"),
//   	},
//   	confirmationConditional: &conditionalSpecificationProperty{
//   		conditionalBranches: []interface{}{
//   			&conditionalBranchProperty{
//   				condition: &conditionProperty{
//   					expressionString: jsii.String("expressionString"),
//   				},
//   				name: jsii.String("name"),
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
//
//   				// the properties below are optional
//   				response: &responseSpecificationProperty{
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
//   			},
//   		},
//   		defaultBranch: &defaultConditionalBranchProperty{
//   			nextStep: &dialogStateProperty{
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
//   			response: &responseSpecificationProperty{
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
//   		isActive: jsii.Boolean(false),
//   	},
//   	confirmationNextStep: &dialogStateProperty{
//   		dialogAction: &dialogActionProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			slotToElicit: jsii.String("slotToElicit"),
//   			suppressNextMessage: jsii.Boolean(false),
//   		},
//   		intent: &intentOverrideProperty{
//   			name: jsii.String("name"),
//   			slots: []interface{}{
//   				&slotValueOverrideMapProperty{
//   					slotName: jsii.String("slotName"),
//   					slotValueOverride: &slotValueOverrideProperty{
//   						shape: jsii.String("shape"),
//   						value: &slotValueProperty{
//   							interpretedValue: jsii.String("interpretedValue"),
//   						},
//   						values: []interface{}{
//   							slotValueOverrideProperty_,
//   						},
//   					},
//   				},
//   			},
//   		},
//   		sessionAttributes: []interface{}{
//   			&sessionAttributeProperty{
//   				key: jsii.String("key"),
//
//   				// the properties below are optional
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	confirmationResponse: &responseSpecificationProperty{
//   		messageGroupsList: []interface{}{
//   			&messageGroupProperty{
//   				message: &messageProperty{
//   					customPayload: &customPayloadProperty{
//   						value: jsii.String("value"),
//   					},
//   					imageResponseCard: &imageResponseCardProperty{
//   						title: jsii.String("title"),
//
//   						// the properties below are optional
//   						buttons: []interface{}{
//   							&buttonProperty{
//   								text: jsii.String("text"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						imageUrl: jsii.String("imageUrl"),
//   						subtitle: jsii.String("subtitle"),
//   					},
//   					plainTextMessage: &plainTextMessageProperty{
//   						value: jsii.String("value"),
//   					},
//   					ssmlMessage: &sSMLMessageProperty{
//   						value: jsii.String("value"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				variations: []interface{}{
//   					&messageProperty{
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
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		allowInterrupt: jsii.Boolean(false),
//   	},
//   	declinationConditional: &conditionalSpecificationProperty{
//   		conditionalBranches: []interface{}{
//   			&conditionalBranchProperty{
//   				condition: &conditionProperty{
//   					expressionString: jsii.String("expressionString"),
//   				},
//   				name: jsii.String("name"),
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
//
//   				// the properties below are optional
//   				response: &responseSpecificationProperty{
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
//   			},
//   		},
//   		defaultBranch: &defaultConditionalBranchProperty{
//   			nextStep: &dialogStateProperty{
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
//   			response: &responseSpecificationProperty{
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
//   		isActive: jsii.Boolean(false),
//   	},
//   	declinationNextStep: &dialogStateProperty{
//   		dialogAction: &dialogActionProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			slotToElicit: jsii.String("slotToElicit"),
//   			suppressNextMessage: jsii.Boolean(false),
//   		},
//   		intent: &intentOverrideProperty{
//   			name: jsii.String("name"),
//   			slots: []interface{}{
//   				&slotValueOverrideMapProperty{
//   					slotName: jsii.String("slotName"),
//   					slotValueOverride: &slotValueOverrideProperty{
//   						shape: jsii.String("shape"),
//   						value: &slotValueProperty{
//   							interpretedValue: jsii.String("interpretedValue"),
//   						},
//   						values: []interface{}{
//   							slotValueOverrideProperty_,
//   						},
//   					},
//   				},
//   			},
//   		},
//   		sessionAttributes: []interface{}{
//   			&sessionAttributeProperty{
//   				key: jsii.String("key"),
//
//   				// the properties below are optional
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	declinationResponse: &responseSpecificationProperty{
//   		messageGroupsList: []interface{}{
//   			&messageGroupProperty{
//   				message: &messageProperty{
//   					customPayload: &customPayloadProperty{
//   						value: jsii.String("value"),
//   					},
//   					imageResponseCard: &imageResponseCardProperty{
//   						title: jsii.String("title"),
//
//   						// the properties below are optional
//   						buttons: []interface{}{
//   							&buttonProperty{
//   								text: jsii.String("text"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						imageUrl: jsii.String("imageUrl"),
//   						subtitle: jsii.String("subtitle"),
//   					},
//   					plainTextMessage: &plainTextMessageProperty{
//   						value: jsii.String("value"),
//   					},
//   					ssmlMessage: &sSMLMessageProperty{
//   						value: jsii.String("value"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				variations: []interface{}{
//   					&messageProperty{
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
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		allowInterrupt: jsii.Boolean(false),
//   	},
//   	elicitationCodeHook: &elicitationCodeHookInvocationSettingProperty{
//   		enableCodeHookInvocation: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		invocationLabel: jsii.String("invocationLabel"),
//   	},
//   	failureConditional: &conditionalSpecificationProperty{
//   		conditionalBranches: []interface{}{
//   			&conditionalBranchProperty{
//   				condition: &conditionProperty{
//   					expressionString: jsii.String("expressionString"),
//   				},
//   				name: jsii.String("name"),
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
//
//   				// the properties below are optional
//   				response: &responseSpecificationProperty{
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
//   			},
//   		},
//   		defaultBranch: &defaultConditionalBranchProperty{
//   			nextStep: &dialogStateProperty{
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
//   			response: &responseSpecificationProperty{
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
//   		isActive: jsii.Boolean(false),
//   	},
//   	failureNextStep: &dialogStateProperty{
//   		dialogAction: &dialogActionProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			slotToElicit: jsii.String("slotToElicit"),
//   			suppressNextMessage: jsii.Boolean(false),
//   		},
//   		intent: &intentOverrideProperty{
//   			name: jsii.String("name"),
//   			slots: []interface{}{
//   				&slotValueOverrideMapProperty{
//   					slotName: jsii.String("slotName"),
//   					slotValueOverride: &slotValueOverrideProperty{
//   						shape: jsii.String("shape"),
//   						value: &slotValueProperty{
//   							interpretedValue: jsii.String("interpretedValue"),
//   						},
//   						values: []interface{}{
//   							slotValueOverrideProperty_,
//   						},
//   					},
//   				},
//   			},
//   		},
//   		sessionAttributes: []interface{}{
//   			&sessionAttributeProperty{
//   				key: jsii.String("key"),
//
//   				// the properties below are optional
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	failureResponse: &responseSpecificationProperty{
//   		messageGroupsList: []interface{}{
//   			&messageGroupProperty{
//   				message: &messageProperty{
//   					customPayload: &customPayloadProperty{
//   						value: jsii.String("value"),
//   					},
//   					imageResponseCard: &imageResponseCardProperty{
//   						title: jsii.String("title"),
//
//   						// the properties below are optional
//   						buttons: []interface{}{
//   							&buttonProperty{
//   								text: jsii.String("text"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						imageUrl: jsii.String("imageUrl"),
//   						subtitle: jsii.String("subtitle"),
//   					},
//   					plainTextMessage: &plainTextMessageProperty{
//   						value: jsii.String("value"),
//   					},
//   					ssmlMessage: &sSMLMessageProperty{
//   						value: jsii.String("value"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				variations: []interface{}{
//   					&messageProperty{
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
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		allowInterrupt: jsii.Boolean(false),
//   	},
//   	isActive: jsii.Boolean(false),
//   }
//
type CfnBot_IntentConfirmationSettingProperty struct {
	// Prompts the user to confirm the intent. This question should have a yes or no answer.
	//
	// Amazon Lex uses this prompt to ensure that the user acknowledges that the intent is ready for fulfillment. For example, with the `OrderPizza` intent, you might want to confirm that the order is correct before placing it. For other intents, such as intents that simply respond to user questions, you might not need to ask the user for confirmation before providing the information.
	PromptSpecification interface{} `field:"required" json:"promptSpecification" yaml:"promptSpecification"`
	// The `DialogCodeHookInvocationSetting` object associated with intent's confirmation step.
	//
	// The dialog code hook is triggered based on these invocation settings when the confirmation next step or declination next step or failure next step is `InvokeDialogCodeHook` .
	CodeHook interface{} `field:"optional" json:"codeHook" yaml:"codeHook"`
	// A list of conditional branches to evaluate after the intent is closed.
	ConfirmationConditional interface{} `field:"optional" json:"confirmationConditional" yaml:"confirmationConditional"`
	// Specifies the next step that the bot executes when the customer confirms the intent.
	ConfirmationNextStep interface{} `field:"optional" json:"confirmationNextStep" yaml:"confirmationNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input.
	ConfirmationResponse interface{} `field:"optional" json:"confirmationResponse" yaml:"confirmationResponse"`
	// A list of conditional branches to evaluate after the intent is declined.
	DeclinationConditional interface{} `field:"optional" json:"declinationConditional" yaml:"declinationConditional"`
	// Specifies the next step that the bot executes when the customer declines the intent.
	DeclinationNextStep interface{} `field:"optional" json:"declinationNextStep" yaml:"declinationNextStep"`
	// When the user answers "no" to the question defined in `promptSpecification` , Amazon Lex responds with this response to acknowledge that the intent was canceled.
	DeclinationResponse interface{} `field:"optional" json:"declinationResponse" yaml:"declinationResponse"`
	// The `DialogCodeHookInvocationSetting` used when the code hook is invoked during confirmation prompt retries.
	ElicitationCodeHook interface{} `field:"optional" json:"elicitationCodeHook" yaml:"elicitationCodeHook"`
	// Provides a list of conditional branches.
	//
	// Branches are evaluated in the order that they are entered in the list. The first branch with a condition that evaluates to true is executed. The last branch in the list is the default branch. The default branch should not have any condition expression. The default branch is executed if no other branch has a matching condition.
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// The next step to take in the conversation if the confirmation step fails.
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond the user input when the intent confirmation fails.
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// Specifies whether the intent's confirmation is sent to the user.
	//
	// When this field is false, confirmation and declination responses aren't sent. If the `IsActive` field isn't specified, the default is true.
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
}

