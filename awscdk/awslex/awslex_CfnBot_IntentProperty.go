package awslex


// Represents an action that the user wants to perform.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intentProperty := &intentProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	dialogCodeHook: &dialogCodeHookSettingProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	fulfillmentCodeHook: &fulfillmentCodeHookSettingProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		fulfillmentUpdatesSpecification: &fulfillmentUpdatesSpecificationProperty{
//   			active: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			startResponse: &fulfillmentStartResponseSpecificationProperty{
//   				delayInSeconds: jsii.Number(123),
//   				messageGroups: []interface{}{
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
//   			timeoutInSeconds: jsii.Number(123),
//   			updateResponse: &fulfillmentUpdateResponseSpecificationProperty{
//   				frequencyInSeconds: jsii.Number(123),
//   				messageGroups: []interface{}{
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
//   		postFulfillmentStatusSpecification: &postFulfillmentStatusSpecificationProperty{
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
//   	},
//   	inputContexts: []interface{}{
//   		&inputContextProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	intentClosingSetting: &intentClosingSettingProperty{
//   		closingResponse: &responseSpecificationProperty{
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
//   	},
//   	intentConfirmationSetting: &intentConfirmationSettingProperty{
//   		declinationResponse: &responseSpecificationProperty{
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
//
//   		// the properties below are optional
//   		isActive: jsii.Boolean(false),
//   	},
//   	kendraConfiguration: &kendraConfigurationProperty{
//   		kendraIndex: jsii.String("kendraIndex"),
//
//   		// the properties below are optional
//   		queryFilterString: jsii.String("queryFilterString"),
//   		queryFilterStringEnabled: jsii.Boolean(false),
//   	},
//   	outputContexts: []interface{}{
//   		&outputContextProperty{
//   			name: jsii.String("name"),
//   			timeToLiveInSeconds: jsii.Number(123),
//   			turnsToLive: jsii.Number(123),
//   		},
//   	},
//   	parentIntentSignature: jsii.String("parentIntentSignature"),
//   	sampleUtterances: []interface{}{
//   		&sampleUtteranceProperty{
//   			utterance: jsii.String("utterance"),
//   		},
//   	},
//   	slotPriorities: []interface{}{
//   		&slotPriorityProperty{
//   			priority: jsii.Number(123),
//   			slotName: jsii.String("slotName"),
//   		},
//   	},
//   	slots: []interface{}{
//   		&slotProperty{
//   			name: jsii.String("name"),
//   			slotTypeName: jsii.String("slotTypeName"),
//   			valueElicitationSetting: &slotValueElicitationSettingProperty{
//   				slotConstraint: jsii.String("slotConstraint"),
//
//   				// the properties below are optional
//   				defaultValueSpecification: &slotDefaultValueSpecificationProperty{
//   					defaultValueList: []interface{}{
//   						&slotDefaultValueProperty{
//   							defaultValue: jsii.String("defaultValue"),
//   						},
//   					},
//   				},
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
//   				sampleUtterances: []interface{}{
//   					&sampleUtteranceProperty{
//   						utterance: jsii.String("utterance"),
//   					},
//   				},
//   				waitAndContinueSpecification: &waitAndContinueSpecificationProperty{
//   					continueResponse: &responseSpecificationProperty{
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
//   					waitingResponse: &responseSpecificationProperty{
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
//
//   					// the properties below are optional
//   					isActive: jsii.Boolean(false),
//   					stillWaitingResponse: &stillWaitingResponseSpecificationProperty{
//   						frequencyInSeconds: jsii.Number(123),
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
//   						timeoutInSeconds: jsii.Number(123),
//
//   						// the properties below are optional
//   						allowInterrupt: jsii.Boolean(false),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			multipleValuesSetting: &multipleValuesSettingProperty{
//   				allowMultipleValues: jsii.Boolean(false),
//   			},
//   			obfuscationSetting: &obfuscationSettingProperty{
//   				obfuscationSettingType: jsii.String("obfuscationSettingType"),
//   			},
//   		},
//   	},
//   }
//
type CfnBot_IntentProperty struct {
	// The name of the intent.
	//
	// Intent names must be unique within the locale that contains the intent and can't match the name of any built-in intent.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the intent.
	//
	// Use the description to help identify the intent in lists.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies that Amazon Lex invokes the alias Lambda function for each user input.
	//
	// You can invoke this Lambda function to personalize user interaction.
	DialogCodeHook interface{} `field:"optional" json:"dialogCodeHook" yaml:"dialogCodeHook"`
	// Specifies that Amazon Lex invokes the alias Lambda function when the intent is ready for fulfillment.
	//
	// You can invoke this function to complete the bot's transaction with the user.
	FulfillmentCodeHook interface{} `field:"optional" json:"fulfillmentCodeHook" yaml:"fulfillmentCodeHook"`
	// A list of contexts that must be active for this intent to be considered by Amazon Lex .
	InputContexts interface{} `field:"optional" json:"inputContexts" yaml:"inputContexts"`
	// Sets the response that Amazon Lex sends to the user when the intent is closed.
	IntentClosingSetting interface{} `field:"optional" json:"intentClosingSetting" yaml:"intentClosingSetting"`
	// Provides prompts that Amazon Lex sends to the user to confirm the completion of an intent.
	//
	// If the user answers "no," the settings contain a statement that is sent to the user to end the intent.
	IntentConfirmationSetting interface{} `field:"optional" json:"intentConfirmationSetting" yaml:"intentConfirmationSetting"`
	// Configuration information required to use the AMAZON.KendraSearchIntent intent to connect to an Amazon Kendra index. The AMAZON.KendraSearchIntent intent is called with Amazon Lex can't determine another intent to invoke.
	KendraConfiguration interface{} `field:"optional" json:"kendraConfiguration" yaml:"kendraConfiguration"`
	// A list of contexts that the intent activates when it is fulfilled.
	OutputContexts interface{} `field:"optional" json:"outputContexts" yaml:"outputContexts"`
	// A unique identifier for the built-in intent to base this intent on.
	ParentIntentSignature *string `field:"optional" json:"parentIntentSignature" yaml:"parentIntentSignature"`
	// A list of utterances that a user might say to signal the intent.
	SampleUtterances interface{} `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// Indicates the priority for slots.
	//
	// Amazon Lex prompts the user for slot values in priority order.
	SlotPriorities interface{} `field:"optional" json:"slotPriorities" yaml:"slotPriorities"`
	// A list of slots that the intent requires for fulfillment.
	Slots interface{} `field:"optional" json:"slots" yaml:"slots"`
}

