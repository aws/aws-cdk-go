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
//   slotValueElicitationSettingProperty := &slotValueElicitationSettingProperty{
//   	slotConstraint: jsii.String("slotConstraint"),
//
//   	// the properties below are optional
//   	defaultValueSpecification: &slotDefaultValueSpecificationProperty{
//   		defaultValueList: []interface{}{
//   			&slotDefaultValueProperty{
//   				defaultValue: jsii.String("defaultValue"),
//   			},
//   		},
//   	},
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
//   	sampleUtterances: []interface{}{
//   		&sampleUtteranceProperty{
//   			utterance: jsii.String("utterance"),
//   		},
//   	},
//   	slotCaptureSetting: &slotCaptureSettingProperty{
//   		captureConditional: &conditionalSpecificationProperty{
//   			conditionalBranches: []interface{}{
//   				&conditionalBranchProperty{
//   					condition: &conditionProperty{
//   						expressionString: jsii.String("expressionString"),
//   					},
//   					name: jsii.String("name"),
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
//
//   					// the properties below are optional
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
//   			},
//   			defaultBranch: &defaultConditionalBranchProperty{
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
//   			isActive: jsii.Boolean(false),
//   		},
//   		captureNextStep: &dialogStateProperty{
//   			dialogAction: &dialogActionProperty{
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				slotToElicit: jsii.String("slotToElicit"),
//   				suppressNextMessage: jsii.Boolean(false),
//   			},
//   			intent: &intentOverrideProperty{
//   				name: jsii.String("name"),
//   				slots: []interface{}{
//   					&slotValueOverrideMapProperty{
//   						slotName: jsii.String("slotName"),
//   						slotValueOverride: &slotValueOverrideProperty{
//   							shape: jsii.String("shape"),
//   							value: &slotValueProperty{
//   								interpretedValue: jsii.String("interpretedValue"),
//   							},
//   							values: []interface{}{
//   								slotValueOverrideProperty_,
//   							},
//   						},
//   					},
//   				},
//   			},
//   			sessionAttributes: []interface{}{
//   				&sessionAttributeProperty{
//   					key: jsii.String("key"),
//
//   					// the properties below are optional
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		captureResponse: &responseSpecificationProperty{
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
//   		},
//   		codeHook: &dialogCodeHookInvocationSettingProperty{
//   			enableCodeHookInvocation: jsii.Boolean(false),
//   			isActive: jsii.Boolean(false),
//   			postCodeHookSpecification: &postDialogCodeHookInvocationSpecificationProperty{
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
//   				successConditional: &conditionalSpecificationProperty{
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
//   				successNextStep: &dialogStateProperty{
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
//   				successResponse: &responseSpecificationProperty{
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
//   				timeoutConditional: &conditionalSpecificationProperty{
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
//   				timeoutNextStep: &dialogStateProperty{
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
//   				timeoutResponse: &responseSpecificationProperty{
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
//
//   			// the properties below are optional
//   			invocationLabel: jsii.String("invocationLabel"),
//   		},
//   		elicitationCodeHook: &elicitationCodeHookInvocationSettingProperty{
//   			enableCodeHookInvocation: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			invocationLabel: jsii.String("invocationLabel"),
//   		},
//   		failureConditional: &conditionalSpecificationProperty{
//   			conditionalBranches: []interface{}{
//   				&conditionalBranchProperty{
//   					condition: &conditionProperty{
//   						expressionString: jsii.String("expressionString"),
//   					},
//   					name: jsii.String("name"),
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
//
//   					// the properties below are optional
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
//   			},
//   			defaultBranch: &defaultConditionalBranchProperty{
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
//   			isActive: jsii.Boolean(false),
//   		},
//   		failureNextStep: &dialogStateProperty{
//   			dialogAction: &dialogActionProperty{
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				slotToElicit: jsii.String("slotToElicit"),
//   				suppressNextMessage: jsii.Boolean(false),
//   			},
//   			intent: &intentOverrideProperty{
//   				name: jsii.String("name"),
//   				slots: []interface{}{
//   					&slotValueOverrideMapProperty{
//   						slotName: jsii.String("slotName"),
//   						slotValueOverride: &slotValueOverrideProperty{
//   							shape: jsii.String("shape"),
//   							value: &slotValueProperty{
//   								interpretedValue: jsii.String("interpretedValue"),
//   							},
//   							values: []interface{}{
//   								slotValueOverrideProperty_,
//   							},
//   						},
//   					},
//   				},
//   			},
//   			sessionAttributes: []interface{}{
//   				&sessionAttributeProperty{
//   					key: jsii.String("key"),
//
//   					// the properties below are optional
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		failureResponse: &responseSpecificationProperty{
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
//   		},
//   	},
//   	waitAndContinueSpecification: &waitAndContinueSpecificationProperty{
//   		continueResponse: &responseSpecificationProperty{
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
//   		},
//   		waitingResponse: &responseSpecificationProperty{
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
//   		},
//
//   		// the properties below are optional
//   		isActive: jsii.Boolean(false),
//   		stillWaitingResponse: &stillWaitingResponseSpecificationProperty{
//   			frequencyInSeconds: jsii.Number(123),
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
//   			timeoutInSeconds: jsii.Number(123),
//
//   			// the properties below are optional
//   			allowInterrupt: jsii.Boolean(false),
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

