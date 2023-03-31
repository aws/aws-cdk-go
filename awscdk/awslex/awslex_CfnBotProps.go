package awslex


// Properties for defining a `CfnBot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataPrivacy interface{}
//   var sentimentAnalysisSettings interface{}
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   cfnBotProps := &cfnBotProps{
//   	dataPrivacy: dataPrivacy,
//   	idleSessionTtlInSeconds: jsii.Number(123),
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	autoBuildBotLocales: jsii.Boolean(false),
//   	botFileS3Location: &s3LocationProperty{
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   		// the properties below are optional
//   		s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	},
//   	botLocales: []interface{}{
//   		&botLocaleProperty{
//   			localeId: jsii.String("localeId"),
//   			nluConfidenceThreshold: jsii.Number(123),
//
//   			// the properties below are optional
//   			customVocabulary: &customVocabularyProperty{
//   				customVocabularyItems: []interface{}{
//   					&customVocabularyItemProperty{
//   						phrase: jsii.String("phrase"),
//
//   						// the properties below are optional
//   						displayAs: jsii.String("displayAs"),
//   						weight: jsii.Number(123),
//   					},
//   				},
//   			},
//   			description: jsii.String("description"),
//   			intents: []interface{}{
//   				&intentProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					description: jsii.String("description"),
//   					dialogCodeHook: &dialogCodeHookSettingProperty{
//   						enabled: jsii.Boolean(false),
//   					},
//   					fulfillmentCodeHook: &fulfillmentCodeHookSettingProperty{
//   						enabled: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						fulfillmentUpdatesSpecification: &fulfillmentUpdatesSpecificationProperty{
//   							active: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							startResponse: &fulfillmentStartResponseSpecificationProperty{
//   								delayInSeconds: jsii.Number(123),
//   								messageGroups: []interface{}{
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
//   							timeoutInSeconds: jsii.Number(123),
//   							updateResponse: &fulfillmentUpdateResponseSpecificationProperty{
//   								frequencyInSeconds: jsii.Number(123),
//   								messageGroups: []interface{}{
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
//   						postFulfillmentStatusSpecification: &postFulfillmentStatusSpecificationProperty{
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
//   							successConditional: &conditionalSpecificationProperty{
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
//   							successNextStep: &dialogStateProperty{
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
//   							successResponse: &responseSpecificationProperty{
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
//   							timeoutConditional: &conditionalSpecificationProperty{
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
//   							timeoutNextStep: &dialogStateProperty{
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
//   							timeoutResponse: &responseSpecificationProperty{
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
//   					initialResponseSetting: &initialResponseSettingProperty{
//   						codeHook: &dialogCodeHookInvocationSettingProperty{
//   							enableCodeHookInvocation: jsii.Boolean(false),
//   							isActive: jsii.Boolean(false),
//   							postCodeHookSpecification: &postDialogCodeHookInvocationSpecificationProperty{
//   								failureConditional: &conditionalSpecificationProperty{
//   									conditionalBranches: []interface{}{
//   										&conditionalBranchProperty{
//   											condition: &conditionProperty{
//   												expressionString: jsii.String("expressionString"),
//   											},
//   											name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									defaultBranch: &defaultConditionalBranchProperty{
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
//   									isActive: jsii.Boolean(false),
//   								},
//   								failureNextStep: &dialogStateProperty{
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
//   								failureResponse: &responseSpecificationProperty{
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
//   								successConditional: &conditionalSpecificationProperty{
//   									conditionalBranches: []interface{}{
//   										&conditionalBranchProperty{
//   											condition: &conditionProperty{
//   												expressionString: jsii.String("expressionString"),
//   											},
//   											name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									defaultBranch: &defaultConditionalBranchProperty{
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
//   									isActive: jsii.Boolean(false),
//   								},
//   								successNextStep: &dialogStateProperty{
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
//   								successResponse: &responseSpecificationProperty{
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
//   								timeoutConditional: &conditionalSpecificationProperty{
//   									conditionalBranches: []interface{}{
//   										&conditionalBranchProperty{
//   											condition: &conditionProperty{
//   												expressionString: jsii.String("expressionString"),
//   											},
//   											name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									defaultBranch: &defaultConditionalBranchProperty{
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
//   									isActive: jsii.Boolean(false),
//   								},
//   								timeoutNextStep: &dialogStateProperty{
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
//   								timeoutResponse: &responseSpecificationProperty{
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
//
//   							// the properties below are optional
//   							invocationLabel: jsii.String("invocationLabel"),
//   						},
//   						conditional: &conditionalSpecificationProperty{
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
//   						initialResponse: &responseSpecificationProperty{
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
//   					},
//   					inputContexts: []interface{}{
//   						&inputContextProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					intentClosingSetting: &intentClosingSettingProperty{
//   						closingResponse: &responseSpecificationProperty{
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
//   						conditional: &conditionalSpecificationProperty{
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
//   						isActive: jsii.Boolean(false),
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
//   					},
//   					intentConfirmationSetting: &intentConfirmationSettingProperty{
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
//
//   						// the properties below are optional
//   						codeHook: &dialogCodeHookInvocationSettingProperty{
//   							enableCodeHookInvocation: jsii.Boolean(false),
//   							isActive: jsii.Boolean(false),
//   							postCodeHookSpecification: &postDialogCodeHookInvocationSpecificationProperty{
//   								failureConditional: &conditionalSpecificationProperty{
//   									conditionalBranches: []interface{}{
//   										&conditionalBranchProperty{
//   											condition: &conditionProperty{
//   												expressionString: jsii.String("expressionString"),
//   											},
//   											name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									defaultBranch: &defaultConditionalBranchProperty{
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
//   									isActive: jsii.Boolean(false),
//   								},
//   								failureNextStep: &dialogStateProperty{
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
//   								failureResponse: &responseSpecificationProperty{
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
//   								successConditional: &conditionalSpecificationProperty{
//   									conditionalBranches: []interface{}{
//   										&conditionalBranchProperty{
//   											condition: &conditionProperty{
//   												expressionString: jsii.String("expressionString"),
//   											},
//   											name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									defaultBranch: &defaultConditionalBranchProperty{
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
//   									isActive: jsii.Boolean(false),
//   								},
//   								successNextStep: &dialogStateProperty{
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
//   								successResponse: &responseSpecificationProperty{
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
//   								timeoutConditional: &conditionalSpecificationProperty{
//   									conditionalBranches: []interface{}{
//   										&conditionalBranchProperty{
//   											condition: &conditionProperty{
//   												expressionString: jsii.String("expressionString"),
//   											},
//   											name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									defaultBranch: &defaultConditionalBranchProperty{
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
//   									isActive: jsii.Boolean(false),
//   								},
//   								timeoutNextStep: &dialogStateProperty{
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
//   								timeoutResponse: &responseSpecificationProperty{
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
//
//   							// the properties below are optional
//   							invocationLabel: jsii.String("invocationLabel"),
//   						},
//   						confirmationConditional: &conditionalSpecificationProperty{
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
//   						confirmationNextStep: &dialogStateProperty{
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
//   						confirmationResponse: &responseSpecificationProperty{
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
//   						declinationConditional: &conditionalSpecificationProperty{
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
//   						declinationNextStep: &dialogStateProperty{
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
//   						declinationResponse: &responseSpecificationProperty{
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
//   						elicitationCodeHook: &elicitationCodeHookInvocationSettingProperty{
//   							enableCodeHookInvocation: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							invocationLabel: jsii.String("invocationLabel"),
//   						},
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
//   						isActive: jsii.Boolean(false),
//   					},
//   					kendraConfiguration: &kendraConfigurationProperty{
//   						kendraIndex: jsii.String("kendraIndex"),
//
//   						// the properties below are optional
//   						queryFilterString: jsii.String("queryFilterString"),
//   						queryFilterStringEnabled: jsii.Boolean(false),
//   					},
//   					outputContexts: []interface{}{
//   						&outputContextProperty{
//   							name: jsii.String("name"),
//   							timeToLiveInSeconds: jsii.Number(123),
//   							turnsToLive: jsii.Number(123),
//   						},
//   					},
//   					parentIntentSignature: jsii.String("parentIntentSignature"),
//   					sampleUtterances: []interface{}{
//   						&sampleUtteranceProperty{
//   							utterance: jsii.String("utterance"),
//   						},
//   					},
//   					slotPriorities: []interface{}{
//   						&slotPriorityProperty{
//   							priority: jsii.Number(123),
//   							slotName: jsii.String("slotName"),
//   						},
//   					},
//   					slots: []interface{}{
//   						&slotProperty{
//   							name: jsii.String("name"),
//   							slotTypeName: jsii.String("slotTypeName"),
//   							valueElicitationSetting: &slotValueElicitationSettingProperty{
//   								slotConstraint: jsii.String("slotConstraint"),
//
//   								// the properties below are optional
//   								defaultValueSpecification: &slotDefaultValueSpecificationProperty{
//   									defaultValueList: []interface{}{
//   										&slotDefaultValueProperty{
//   											defaultValue: jsii.String("defaultValue"),
//   										},
//   									},
//   								},
//   								promptSpecification: &promptSpecificationProperty{
//   									maxRetries: jsii.Number(123),
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
//   									messageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   									promptAttemptsSpecification: map[string]interface{}{
//   										"promptAttemptsSpecificationKey": &PromptAttemptSpecificationProperty{
//   											"allowedInputTypes": &AllowedInputTypesProperty{
//   												"allowAudioInput": jsii.Boolean(false),
//   												"allowDtmfInput": jsii.Boolean(false),
//   											},
//
//   											// the properties below are optional
//   											"allowInterrupt": jsii.Boolean(false),
//   											"audioAndDtmfInputSpecification": &AudioAndDTMFInputSpecificationProperty{
//   												"startTimeoutMs": jsii.Number(123),
//
//   												// the properties below are optional
//   												"audioSpecification": &AudioSpecificationProperty{
//   													"endTimeoutMs": jsii.Number(123),
//   													"maxLengthMs": jsii.Number(123),
//   												},
//   												"dtmfSpecification": &DTMFSpecificationProperty{
//   													"deletionCharacter": jsii.String("deletionCharacter"),
//   													"endCharacter": jsii.String("endCharacter"),
//   													"endTimeoutMs": jsii.Number(123),
//   													"maxLength": jsii.Number(123),
//   												},
//   											},
//   											"textInputSpecification": &TextInputSpecificationProperty{
//   												"startTimeoutMs": jsii.Number(123),
//   											},
//   										},
//   									},
//   								},
//   								sampleUtterances: []interface{}{
//   									&sampleUtteranceProperty{
//   										utterance: jsii.String("utterance"),
//   									},
//   								},
//   								slotCaptureSetting: &slotCaptureSettingProperty{
//   									captureConditional: &conditionalSpecificationProperty{
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
//   									captureNextStep: &dialogStateProperty{
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
//   									captureResponse: &responseSpecificationProperty{
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
//   									codeHook: &dialogCodeHookInvocationSettingProperty{
//   										enableCodeHookInvocation: jsii.Boolean(false),
//   										isActive: jsii.Boolean(false),
//   										postCodeHookSpecification: &postDialogCodeHookInvocationSpecificationProperty{
//   											failureConditional: &conditionalSpecificationProperty{
//   												conditionalBranches: []interface{}{
//   													&conditionalBranchProperty{
//   														condition: &conditionProperty{
//   															expressionString: jsii.String("expressionString"),
//   														},
//   														name: jsii.String("name"),
//   														nextStep: &dialogStateProperty{
//   															dialogAction: &dialogActionProperty{
//   																type: jsii.String("type"),
//
//   																// the properties below are optional
//   																slotToElicit: jsii.String("slotToElicit"),
//   																suppressNextMessage: jsii.Boolean(false),
//   															},
//   															intent: &intentOverrideProperty{
//   																name: jsii.String("name"),
//   																slots: []interface{}{
//   																	&slotValueOverrideMapProperty{
//   																		slotName: jsii.String("slotName"),
//   																		slotValueOverride: &slotValueOverrideProperty{
//   																			shape: jsii.String("shape"),
//   																			value: &slotValueProperty{
//   																				interpretedValue: jsii.String("interpretedValue"),
//   																			},
//   																			values: []interface{}{
//   																				slotValueOverrideProperty_,
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//   															sessionAttributes: []interface{}{
//   																&sessionAttributeProperty{
//   																	key: jsii.String("key"),
//
//   																	// the properties below are optional
//   																	value: jsii.String("value"),
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														response: &responseSpecificationProperty{
//   															messageGroupsList: []interface{}{
//   																&messageGroupProperty{
//   																	message: &messageProperty{
//   																		customPayload: &customPayloadProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		imageResponseCard: &imageResponseCardProperty{
//   																			title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			buttons: []interface{}{
//   																				&buttonProperty{
//   																					text: jsii.String("text"),
//   																					value: jsii.String("value"),
//   																				},
//   																			},
//   																			imageUrl: jsii.String("imageUrl"),
//   																			subtitle: jsii.String("subtitle"),
//   																		},
//   																		plainTextMessage: &plainTextMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		ssmlMessage: &sSMLMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//
//   																	// the properties below are optional
//   																	variations: []interface{}{
//   																		&messageProperty{
//   																			customPayload: &customPayloadProperty{
//   																				value: jsii.String("value"),
//   																			},
//   																			imageResponseCard: &imageResponseCardProperty{
//   																				title: jsii.String("title"),
//
//   																				// the properties below are optional
//   																				buttons: []interface{}{
//   																					&buttonProperty{
//   																						text: jsii.String("text"),
//   																						value: jsii.String("value"),
//   																					},
//   																				},
//   																				imageUrl: jsii.String("imageUrl"),
//   																				subtitle: jsii.String("subtitle"),
//   																			},
//   																			plainTextMessage: &plainTextMessageProperty{
//   																				value: jsii.String("value"),
//   																			},
//   																			ssmlMessage: &sSMLMessageProperty{
//   																				value: jsii.String("value"),
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//
//   															// the properties below are optional
//   															allowInterrupt: jsii.Boolean(false),
//   														},
//   													},
//   												},
//   												defaultBranch: &defaultConditionalBranchProperty{
//   													nextStep: &dialogStateProperty{
//   														dialogAction: &dialogActionProperty{
//   															type: jsii.String("type"),
//
//   															// the properties below are optional
//   															slotToElicit: jsii.String("slotToElicit"),
//   															suppressNextMessage: jsii.Boolean(false),
//   														},
//   														intent: &intentOverrideProperty{
//   															name: jsii.String("name"),
//   															slots: []interface{}{
//   																&slotValueOverrideMapProperty{
//   																	slotName: jsii.String("slotName"),
//   																	slotValueOverride: &slotValueOverrideProperty{
//   																		shape: jsii.String("shape"),
//   																		value: &slotValueProperty{
//   																			interpretedValue: jsii.String("interpretedValue"),
//   																		},
//   																		values: []interface{}{
//   																			slotValueOverrideProperty_,
//   																		},
//   																	},
//   																},
//   															},
//   														},
//   														sessionAttributes: []interface{}{
//   															&sessionAttributeProperty{
//   																key: jsii.String("key"),
//
//   																// the properties below are optional
//   																value: jsii.String("value"),
//   															},
//   														},
//   													},
//   													response: &responseSpecificationProperty{
//   														messageGroupsList: []interface{}{
//   															&messageGroupProperty{
//   																message: &messageProperty{
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
//
//   																// the properties below are optional
//   																variations: []interface{}{
//   																	&messageProperty{
//   																		customPayload: &customPayloadProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		imageResponseCard: &imageResponseCardProperty{
//   																			title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			buttons: []interface{}{
//   																				&buttonProperty{
//   																					text: jsii.String("text"),
//   																					value: jsii.String("value"),
//   																				},
//   																			},
//   																			imageUrl: jsii.String("imageUrl"),
//   																			subtitle: jsii.String("subtitle"),
//   																		},
//   																		plainTextMessage: &plainTextMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		ssmlMessage: &sSMLMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														allowInterrupt: jsii.Boolean(false),
//   													},
//   												},
//   												isActive: jsii.Boolean(false),
//   											},
//   											failureNextStep: &dialogStateProperty{
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
//   											failureResponse: &responseSpecificationProperty{
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
//   											successConditional: &conditionalSpecificationProperty{
//   												conditionalBranches: []interface{}{
//   													&conditionalBranchProperty{
//   														condition: &conditionProperty{
//   															expressionString: jsii.String("expressionString"),
//   														},
//   														name: jsii.String("name"),
//   														nextStep: &dialogStateProperty{
//   															dialogAction: &dialogActionProperty{
//   																type: jsii.String("type"),
//
//   																// the properties below are optional
//   																slotToElicit: jsii.String("slotToElicit"),
//   																suppressNextMessage: jsii.Boolean(false),
//   															},
//   															intent: &intentOverrideProperty{
//   																name: jsii.String("name"),
//   																slots: []interface{}{
//   																	&slotValueOverrideMapProperty{
//   																		slotName: jsii.String("slotName"),
//   																		slotValueOverride: &slotValueOverrideProperty{
//   																			shape: jsii.String("shape"),
//   																			value: &slotValueProperty{
//   																				interpretedValue: jsii.String("interpretedValue"),
//   																			},
//   																			values: []interface{}{
//   																				slotValueOverrideProperty_,
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//   															sessionAttributes: []interface{}{
//   																&sessionAttributeProperty{
//   																	key: jsii.String("key"),
//
//   																	// the properties below are optional
//   																	value: jsii.String("value"),
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														response: &responseSpecificationProperty{
//   															messageGroupsList: []interface{}{
//   																&messageGroupProperty{
//   																	message: &messageProperty{
//   																		customPayload: &customPayloadProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		imageResponseCard: &imageResponseCardProperty{
//   																			title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			buttons: []interface{}{
//   																				&buttonProperty{
//   																					text: jsii.String("text"),
//   																					value: jsii.String("value"),
//   																				},
//   																			},
//   																			imageUrl: jsii.String("imageUrl"),
//   																			subtitle: jsii.String("subtitle"),
//   																		},
//   																		plainTextMessage: &plainTextMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		ssmlMessage: &sSMLMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//
//   																	// the properties below are optional
//   																	variations: []interface{}{
//   																		&messageProperty{
//   																			customPayload: &customPayloadProperty{
//   																				value: jsii.String("value"),
//   																			},
//   																			imageResponseCard: &imageResponseCardProperty{
//   																				title: jsii.String("title"),
//
//   																				// the properties below are optional
//   																				buttons: []interface{}{
//   																					&buttonProperty{
//   																						text: jsii.String("text"),
//   																						value: jsii.String("value"),
//   																					},
//   																				},
//   																				imageUrl: jsii.String("imageUrl"),
//   																				subtitle: jsii.String("subtitle"),
//   																			},
//   																			plainTextMessage: &plainTextMessageProperty{
//   																				value: jsii.String("value"),
//   																			},
//   																			ssmlMessage: &sSMLMessageProperty{
//   																				value: jsii.String("value"),
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//
//   															// the properties below are optional
//   															allowInterrupt: jsii.Boolean(false),
//   														},
//   													},
//   												},
//   												defaultBranch: &defaultConditionalBranchProperty{
//   													nextStep: &dialogStateProperty{
//   														dialogAction: &dialogActionProperty{
//   															type: jsii.String("type"),
//
//   															// the properties below are optional
//   															slotToElicit: jsii.String("slotToElicit"),
//   															suppressNextMessage: jsii.Boolean(false),
//   														},
//   														intent: &intentOverrideProperty{
//   															name: jsii.String("name"),
//   															slots: []interface{}{
//   																&slotValueOverrideMapProperty{
//   																	slotName: jsii.String("slotName"),
//   																	slotValueOverride: &slotValueOverrideProperty{
//   																		shape: jsii.String("shape"),
//   																		value: &slotValueProperty{
//   																			interpretedValue: jsii.String("interpretedValue"),
//   																		},
//   																		values: []interface{}{
//   																			slotValueOverrideProperty_,
//   																		},
//   																	},
//   																},
//   															},
//   														},
//   														sessionAttributes: []interface{}{
//   															&sessionAttributeProperty{
//   																key: jsii.String("key"),
//
//   																// the properties below are optional
//   																value: jsii.String("value"),
//   															},
//   														},
//   													},
//   													response: &responseSpecificationProperty{
//   														messageGroupsList: []interface{}{
//   															&messageGroupProperty{
//   																message: &messageProperty{
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
//
//   																// the properties below are optional
//   																variations: []interface{}{
//   																	&messageProperty{
//   																		customPayload: &customPayloadProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		imageResponseCard: &imageResponseCardProperty{
//   																			title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			buttons: []interface{}{
//   																				&buttonProperty{
//   																					text: jsii.String("text"),
//   																					value: jsii.String("value"),
//   																				},
//   																			},
//   																			imageUrl: jsii.String("imageUrl"),
//   																			subtitle: jsii.String("subtitle"),
//   																		},
//   																		plainTextMessage: &plainTextMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		ssmlMessage: &sSMLMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														allowInterrupt: jsii.Boolean(false),
//   													},
//   												},
//   												isActive: jsii.Boolean(false),
//   											},
//   											successNextStep: &dialogStateProperty{
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
//   											successResponse: &responseSpecificationProperty{
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
//   											timeoutConditional: &conditionalSpecificationProperty{
//   												conditionalBranches: []interface{}{
//   													&conditionalBranchProperty{
//   														condition: &conditionProperty{
//   															expressionString: jsii.String("expressionString"),
//   														},
//   														name: jsii.String("name"),
//   														nextStep: &dialogStateProperty{
//   															dialogAction: &dialogActionProperty{
//   																type: jsii.String("type"),
//
//   																// the properties below are optional
//   																slotToElicit: jsii.String("slotToElicit"),
//   																suppressNextMessage: jsii.Boolean(false),
//   															},
//   															intent: &intentOverrideProperty{
//   																name: jsii.String("name"),
//   																slots: []interface{}{
//   																	&slotValueOverrideMapProperty{
//   																		slotName: jsii.String("slotName"),
//   																		slotValueOverride: &slotValueOverrideProperty{
//   																			shape: jsii.String("shape"),
//   																			value: &slotValueProperty{
//   																				interpretedValue: jsii.String("interpretedValue"),
//   																			},
//   																			values: []interface{}{
//   																				slotValueOverrideProperty_,
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//   															sessionAttributes: []interface{}{
//   																&sessionAttributeProperty{
//   																	key: jsii.String("key"),
//
//   																	// the properties below are optional
//   																	value: jsii.String("value"),
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														response: &responseSpecificationProperty{
//   															messageGroupsList: []interface{}{
//   																&messageGroupProperty{
//   																	message: &messageProperty{
//   																		customPayload: &customPayloadProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		imageResponseCard: &imageResponseCardProperty{
//   																			title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			buttons: []interface{}{
//   																				&buttonProperty{
//   																					text: jsii.String("text"),
//   																					value: jsii.String("value"),
//   																				},
//   																			},
//   																			imageUrl: jsii.String("imageUrl"),
//   																			subtitle: jsii.String("subtitle"),
//   																		},
//   																		plainTextMessage: &plainTextMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		ssmlMessage: &sSMLMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//
//   																	// the properties below are optional
//   																	variations: []interface{}{
//   																		&messageProperty{
//   																			customPayload: &customPayloadProperty{
//   																				value: jsii.String("value"),
//   																			},
//   																			imageResponseCard: &imageResponseCardProperty{
//   																				title: jsii.String("title"),
//
//   																				// the properties below are optional
//   																				buttons: []interface{}{
//   																					&buttonProperty{
//   																						text: jsii.String("text"),
//   																						value: jsii.String("value"),
//   																					},
//   																				},
//   																				imageUrl: jsii.String("imageUrl"),
//   																				subtitle: jsii.String("subtitle"),
//   																			},
//   																			plainTextMessage: &plainTextMessageProperty{
//   																				value: jsii.String("value"),
//   																			},
//   																			ssmlMessage: &sSMLMessageProperty{
//   																				value: jsii.String("value"),
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//
//   															// the properties below are optional
//   															allowInterrupt: jsii.Boolean(false),
//   														},
//   													},
//   												},
//   												defaultBranch: &defaultConditionalBranchProperty{
//   													nextStep: &dialogStateProperty{
//   														dialogAction: &dialogActionProperty{
//   															type: jsii.String("type"),
//
//   															// the properties below are optional
//   															slotToElicit: jsii.String("slotToElicit"),
//   															suppressNextMessage: jsii.Boolean(false),
//   														},
//   														intent: &intentOverrideProperty{
//   															name: jsii.String("name"),
//   															slots: []interface{}{
//   																&slotValueOverrideMapProperty{
//   																	slotName: jsii.String("slotName"),
//   																	slotValueOverride: &slotValueOverrideProperty{
//   																		shape: jsii.String("shape"),
//   																		value: &slotValueProperty{
//   																			interpretedValue: jsii.String("interpretedValue"),
//   																		},
//   																		values: []interface{}{
//   																			slotValueOverrideProperty_,
//   																		},
//   																	},
//   																},
//   															},
//   														},
//   														sessionAttributes: []interface{}{
//   															&sessionAttributeProperty{
//   																key: jsii.String("key"),
//
//   																// the properties below are optional
//   																value: jsii.String("value"),
//   															},
//   														},
//   													},
//   													response: &responseSpecificationProperty{
//   														messageGroupsList: []interface{}{
//   															&messageGroupProperty{
//   																message: &messageProperty{
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
//
//   																// the properties below are optional
//   																variations: []interface{}{
//   																	&messageProperty{
//   																		customPayload: &customPayloadProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		imageResponseCard: &imageResponseCardProperty{
//   																			title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			buttons: []interface{}{
//   																				&buttonProperty{
//   																					text: jsii.String("text"),
//   																					value: jsii.String("value"),
//   																				},
//   																			},
//   																			imageUrl: jsii.String("imageUrl"),
//   																			subtitle: jsii.String("subtitle"),
//   																		},
//   																		plainTextMessage: &plainTextMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																		ssmlMessage: &sSMLMessageProperty{
//   																			value: jsii.String("value"),
//   																		},
//   																	},
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														allowInterrupt: jsii.Boolean(false),
//   													},
//   												},
//   												isActive: jsii.Boolean(false),
//   											},
//   											timeoutNextStep: &dialogStateProperty{
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
//   											timeoutResponse: &responseSpecificationProperty{
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
//
//   										// the properties below are optional
//   										invocationLabel: jsii.String("invocationLabel"),
//   									},
//   									elicitationCodeHook: &elicitationCodeHookInvocationSettingProperty{
//   										enableCodeHookInvocation: jsii.Boolean(false),
//
//   										// the properties below are optional
//   										invocationLabel: jsii.String("invocationLabel"),
//   									},
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
//   								},
//   								waitAndContinueSpecification: &waitAndContinueSpecificationProperty{
//   									continueResponse: &responseSpecificationProperty{
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
//   									waitingResponse: &responseSpecificationProperty{
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
//
//   									// the properties below are optional
//   									isActive: jsii.Boolean(false),
//   									stillWaitingResponse: &stillWaitingResponseSpecificationProperty{
//   										frequencyInSeconds: jsii.Number(123),
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
//   										timeoutInSeconds: jsii.Number(123),
//
//   										// the properties below are optional
//   										allowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							description: jsii.String("description"),
//   							multipleValuesSetting: &multipleValuesSettingProperty{
//   								allowMultipleValues: jsii.Boolean(false),
//   							},
//   							obfuscationSetting: &obfuscationSettingProperty{
//   								obfuscationSettingType: jsii.String("obfuscationSettingType"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			slotTypes: []interface{}{
//   				&slotTypeProperty{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					description: jsii.String("description"),
//   					externalSourceSetting: &externalSourceSettingProperty{
//   						grammarSlotTypeSetting: &grammarSlotTypeSettingProperty{
//   							source: &grammarSlotTypeSourceProperty{
//   								s3BucketName: jsii.String("s3BucketName"),
//   								s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   								// the properties below are optional
//   								kmsKeyArn: jsii.String("kmsKeyArn"),
//   							},
//   						},
//   					},
//   					parentSlotTypeSignature: jsii.String("parentSlotTypeSignature"),
//   					slotTypeValues: []interface{}{
//   						&slotTypeValueProperty{
//   							sampleValue: &sampleValueProperty{
//   								value: jsii.String("value"),
//   							},
//
//   							// the properties below are optional
//   							synonyms: []interface{}{
//   								&sampleValueProperty{
//   									value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					valueSelectionSetting: &slotValueSelectionSettingProperty{
//   						resolutionStrategy: jsii.String("resolutionStrategy"),
//
//   						// the properties below are optional
//   						advancedRecognitionSetting: &advancedRecognitionSettingProperty{
//   							audioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   						},
//   						regexFilter: &slotValueRegexFilterProperty{
//   							pattern: jsii.String("pattern"),
//   						},
//   					},
//   				},
//   			},
//   			voiceSettings: &voiceSettingsProperty{
//   				voiceId: jsii.String("voiceId"),
//
//   				// the properties below are optional
//   				engine: jsii.String("engine"),
//   			},
//   		},
//   	},
//   	botTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	testBotAliasSettings: &testBotAliasSettingsProperty{
//   		botAliasLocaleSettings: []interface{}{
//   			&botAliasLocaleSettingsItemProperty{
//   				botAliasLocaleSetting: &botAliasLocaleSettingsProperty{
//   					enabled: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					codeHookSpecification: &codeHookSpecificationProperty{
//   						lambdaCodeHook: &lambdaCodeHookProperty{
//   							codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   							lambdaArn: jsii.String("lambdaArn"),
//   						},
//   					},
//   				},
//   				localeId: jsii.String("localeId"),
//   			},
//   		},
//   		conversationLogSettings: &conversationLogSettingsProperty{
//   			audioLogSettings: []interface{}{
//   				&audioLogSettingProperty{
//   					destination: &audioLogDestinationProperty{
//   						s3Bucket: &s3BucketLogDestinationProperty{
//   							logPrefix: jsii.String("logPrefix"),
//   							s3BucketArn: jsii.String("s3BucketArn"),
//
//   							// the properties below are optional
//   							kmsKeyArn: jsii.String("kmsKeyArn"),
//   						},
//   					},
//   					enabled: jsii.Boolean(false),
//   				},
//   			},
//   			textLogSettings: []interface{}{
//   				&textLogSettingProperty{
//   					destination: &textLogDestinationProperty{
//   						cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   							cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   							logPrefix: jsii.String("logPrefix"),
//   						},
//   					},
//   					enabled: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		description: jsii.String("description"),
//   		sentimentAnalysisSettings: sentimentAnalysisSettings,
//   	},
//   	testBotAliasTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBotProps struct {
	// `AWS::Lex::Bot.DataPrivacy`.
	DataPrivacy interface{} `field:"required" json:"dataPrivacy" yaml:"dataPrivacy"`
	// The time, in seconds, that Amazon Lex should keep information about a user's conversation with the bot.
	//
	// A user interaction remains active for the amount of time specified. If no conversation occurs during this time, the session expires and Amazon Lex deletes any data provided before the timeout.
	//
	// You can specify between 60 (1 minute) and 86,400 (24 hours) seconds.
	IdleSessionTtlInSeconds *float64 `field:"required" json:"idleSessionTtlInSeconds" yaml:"idleSessionTtlInSeconds"`
	// The name of the bot locale.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the IAM role used to build and run the bot.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Indicates whether Amazon Lex V2 should automatically build the locales for the bot after a change.
	AutoBuildBotLocales interface{} `field:"optional" json:"autoBuildBotLocales" yaml:"autoBuildBotLocales"`
	// The Amazon S3 location of files used to import a bot.
	//
	// The files must be in the import format specified in [JSON format for importing and exporting](https://docs.aws.amazon.com/lexv2/latest/dg/import-export-format.html) in the *Amazon Lex developer guide.*
	BotFileS3Location interface{} `field:"optional" json:"botFileS3Location" yaml:"botFileS3Location"`
	// A list of locales for the bot.
	BotLocales interface{} `field:"optional" json:"botLocales" yaml:"botLocales"`
	// A list of tags to add to the bot.
	//
	// You can only add tags when you import a bot. You can't use the `UpdateBot` operation to update tags. To update tags, use the `TagResource` operation.
	BotTags interface{} `field:"optional" json:"botTags" yaml:"botTags"`
	// The description of the version.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies configuration settings for the alias used to test the bot.
	//
	// If the `TestBotAliasSettings` property is not specified, the settings are configured with default values.
	TestBotAliasSettings interface{} `field:"optional" json:"testBotAliasSettings" yaml:"testBotAliasSettings"`
	// A list of tags to add to the test alias for a bot.
	//
	// You can only add tags when you import a bot. You can't use the `UpdateAlias` operation to update tags. To update tags on the test alias, use the `TagResource` operation.
	TestBotAliasTags interface{} `field:"optional" json:"testBotAliasTags" yaml:"testBotAliasTags"`
}

