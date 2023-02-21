package awslex


// Settings that specify the dialog code hook that is called by Amazon Lex at a step of the conversation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   dialogCodeHookInvocationSettingProperty := &dialogCodeHookInvocationSettingProperty{
//   	enableCodeHookInvocation: jsii.Boolean(false),
//   	isActive: jsii.Boolean(false),
//   	postCodeHookSpecification: &postDialogCodeHookInvocationSpecificationProperty{
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
//   		successConditional: &conditionalSpecificationProperty{
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
//   		successNextStep: &dialogStateProperty{
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
//   		successResponse: &responseSpecificationProperty{
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
//   		timeoutConditional: &conditionalSpecificationProperty{
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
//   		timeoutNextStep: &dialogStateProperty{
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
//   		timeoutResponse: &responseSpecificationProperty{
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
//
//   	// the properties below are optional
//   	invocationLabel: jsii.String("invocationLabel"),
//   }
//
type CfnBot_DialogCodeHookInvocationSettingProperty struct {
	// Indicates whether a Lambda function should be invoked for the dialog.
	EnableCodeHookInvocation interface{} `field:"required" json:"enableCodeHookInvocation" yaml:"enableCodeHookInvocation"`
	// Determines whether a dialog code hook is used when the intent is activated.
	IsActive interface{} `field:"required" json:"isActive" yaml:"isActive"`
	// Contains the responses and actions that Amazon Lex takes after the Lambda function is complete.
	PostCodeHookSpecification interface{} `field:"required" json:"postCodeHookSpecification" yaml:"postCodeHookSpecification"`
	// A label that indicates the dialog step from which the dialog code hook is happening.
	InvocationLabel *string `field:"optional" json:"invocationLabel" yaml:"invocationLabel"`
}

