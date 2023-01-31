package awslex


// Provides a setting that determines whether the post-fulfillment response is sent to the user.
//
// For more information, see [Post-fulfillment response](https://docs.aws.amazon.com/lex/latest/dg/streaming-progress.html#progress-complete) in the *Amazon Lex developer guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   postFulfillmentStatusSpecificationProperty := &postFulfillmentStatusSpecificationProperty{
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
//   	successConditional: &conditionalSpecificationProperty{
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
//   	successNextStep: &dialogStateProperty{
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
//   	successResponse: &responseSpecificationProperty{
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
//   	timeoutConditional: &conditionalSpecificationProperty{
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
//   	timeoutNextStep: &dialogStateProperty{
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
//   	timeoutResponse: &responseSpecificationProperty{
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
//   }
//
type CfnBot_PostFulfillmentStatusSpecificationProperty struct {
	// `CfnBot.PostFulfillmentStatusSpecificationProperty.FailureConditional`.
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// `CfnBot.PostFulfillmentStatusSpecificationProperty.FailureNextStep`.
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't successful.
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// `CfnBot.PostFulfillmentStatusSpecificationProperty.SuccessConditional`.
	SuccessConditional interface{} `field:"optional" json:"successConditional" yaml:"successConditional"`
	// `CfnBot.PostFulfillmentStatusSpecificationProperty.SuccessNextStep`.
	SuccessNextStep interface{} `field:"optional" json:"successNextStep" yaml:"successNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when the fulfillment is successful.
	SuccessResponse interface{} `field:"optional" json:"successResponse" yaml:"successResponse"`
	// `CfnBot.PostFulfillmentStatusSpecificationProperty.TimeoutConditional`.
	TimeoutConditional interface{} `field:"optional" json:"timeoutConditional" yaml:"timeoutConditional"`
	// `CfnBot.PostFulfillmentStatusSpecificationProperty.TimeoutNextStep`.
	TimeoutNextStep interface{} `field:"optional" json:"timeoutNextStep" yaml:"timeoutNextStep"`
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't completed within the timeout period.
	TimeoutResponse interface{} `field:"optional" json:"timeoutResponse" yaml:"timeoutResponse"`
}

