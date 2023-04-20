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
//   cfnBotProps := &CfnBotProps{
//   	DataPrivacy: dataPrivacy,
//   	IdleSessionTtlInSeconds: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AutoBuildBotLocales: jsii.Boolean(false),
//   	BotFileS3Location: &S3LocationProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3ObjectKey: jsii.String("s3ObjectKey"),
//
//   		// the properties below are optional
//   		S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   	},
//   	BotLocales: []interface{}{
//   		&BotLocaleProperty{
//   			LocaleId: jsii.String("localeId"),
//   			NluConfidenceThreshold: jsii.Number(123),
//
//   			// the properties below are optional
//   			CustomVocabulary: &CustomVocabularyProperty{
//   				CustomVocabularyItems: []interface{}{
//   					&CustomVocabularyItemProperty{
//   						Phrase: jsii.String("phrase"),
//
//   						// the properties below are optional
//   						DisplayAs: jsii.String("displayAs"),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			Intents: []interface{}{
//   				&IntentProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Description: jsii.String("description"),
//   					DialogCodeHook: &DialogCodeHookSettingProperty{
//   						Enabled: jsii.Boolean(false),
//   					},
//   					FulfillmentCodeHook: &FulfillmentCodeHookSettingProperty{
//   						Enabled: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						FulfillmentUpdatesSpecification: &FulfillmentUpdatesSpecificationProperty{
//   							Active: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							StartResponse: &FulfillmentStartResponseSpecificationProperty{
//   								DelayInSeconds: jsii.Number(123),
//   								MessageGroups: []interface{}{
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
//   							TimeoutInSeconds: jsii.Number(123),
//   							UpdateResponse: &FulfillmentUpdateResponseSpecificationProperty{
//   								FrequencyInSeconds: jsii.Number(123),
//   								MessageGroups: []interface{}{
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
//   						PostFulfillmentStatusSpecification: &PostFulfillmentStatusSpecificationProperty{
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
//   							SuccessConditional: &ConditionalSpecificationProperty{
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
//   							SuccessNextStep: &DialogStateProperty{
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
//   							SuccessResponse: &ResponseSpecificationProperty{
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
//   							TimeoutConditional: &ConditionalSpecificationProperty{
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
//   							TimeoutNextStep: &DialogStateProperty{
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
//   							TimeoutResponse: &ResponseSpecificationProperty{
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
//   					InitialResponseSetting: &InitialResponseSettingProperty{
//   						CodeHook: &DialogCodeHookInvocationSettingProperty{
//   							EnableCodeHookInvocation: jsii.Boolean(false),
//   							IsActive: jsii.Boolean(false),
//   							PostCodeHookSpecification: &PostDialogCodeHookInvocationSpecificationProperty{
//   								FailureConditional: &ConditionalSpecificationProperty{
//   									ConditionalBranches: []interface{}{
//   										&ConditionalBranchProperty{
//   											Condition: &ConditionProperty{
//   												ExpressionString: jsii.String("expressionString"),
//   											},
//   											Name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									DefaultBranch: &DefaultConditionalBranchProperty{
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
//   									IsActive: jsii.Boolean(false),
//   								},
//   								FailureNextStep: &DialogStateProperty{
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
//   								FailureResponse: &ResponseSpecificationProperty{
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
//   								SuccessConditional: &ConditionalSpecificationProperty{
//   									ConditionalBranches: []interface{}{
//   										&ConditionalBranchProperty{
//   											Condition: &ConditionProperty{
//   												ExpressionString: jsii.String("expressionString"),
//   											},
//   											Name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									DefaultBranch: &DefaultConditionalBranchProperty{
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
//   									IsActive: jsii.Boolean(false),
//   								},
//   								SuccessNextStep: &DialogStateProperty{
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
//   								SuccessResponse: &ResponseSpecificationProperty{
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
//   								TimeoutConditional: &ConditionalSpecificationProperty{
//   									ConditionalBranches: []interface{}{
//   										&ConditionalBranchProperty{
//   											Condition: &ConditionProperty{
//   												ExpressionString: jsii.String("expressionString"),
//   											},
//   											Name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									DefaultBranch: &DefaultConditionalBranchProperty{
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
//   									IsActive: jsii.Boolean(false),
//   								},
//   								TimeoutNextStep: &DialogStateProperty{
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
//   								TimeoutResponse: &ResponseSpecificationProperty{
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
//
//   							// the properties below are optional
//   							InvocationLabel: jsii.String("invocationLabel"),
//   						},
//   						Conditional: &ConditionalSpecificationProperty{
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
//   						InitialResponse: &ResponseSpecificationProperty{
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
//   					},
//   					InputContexts: []interface{}{
//   						&InputContextProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					IntentClosingSetting: &IntentClosingSettingProperty{
//   						ClosingResponse: &ResponseSpecificationProperty{
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
//   						Conditional: &ConditionalSpecificationProperty{
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
//   						IsActive: jsii.Boolean(false),
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
//   					},
//   					IntentConfirmationSetting: &IntentConfirmationSettingProperty{
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
//
//   						// the properties below are optional
//   						CodeHook: &DialogCodeHookInvocationSettingProperty{
//   							EnableCodeHookInvocation: jsii.Boolean(false),
//   							IsActive: jsii.Boolean(false),
//   							PostCodeHookSpecification: &PostDialogCodeHookInvocationSpecificationProperty{
//   								FailureConditional: &ConditionalSpecificationProperty{
//   									ConditionalBranches: []interface{}{
//   										&ConditionalBranchProperty{
//   											Condition: &ConditionProperty{
//   												ExpressionString: jsii.String("expressionString"),
//   											},
//   											Name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									DefaultBranch: &DefaultConditionalBranchProperty{
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
//   									IsActive: jsii.Boolean(false),
//   								},
//   								FailureNextStep: &DialogStateProperty{
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
//   								FailureResponse: &ResponseSpecificationProperty{
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
//   								SuccessConditional: &ConditionalSpecificationProperty{
//   									ConditionalBranches: []interface{}{
//   										&ConditionalBranchProperty{
//   											Condition: &ConditionProperty{
//   												ExpressionString: jsii.String("expressionString"),
//   											},
//   											Name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									DefaultBranch: &DefaultConditionalBranchProperty{
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
//   									IsActive: jsii.Boolean(false),
//   								},
//   								SuccessNextStep: &DialogStateProperty{
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
//   								SuccessResponse: &ResponseSpecificationProperty{
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
//   								TimeoutConditional: &ConditionalSpecificationProperty{
//   									ConditionalBranches: []interface{}{
//   										&ConditionalBranchProperty{
//   											Condition: &ConditionProperty{
//   												ExpressionString: jsii.String("expressionString"),
//   											},
//   											Name: jsii.String("name"),
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
//
//   											// the properties below are optional
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
//   									},
//   									DefaultBranch: &DefaultConditionalBranchProperty{
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
//   									IsActive: jsii.Boolean(false),
//   								},
//   								TimeoutNextStep: &DialogStateProperty{
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
//   								TimeoutResponse: &ResponseSpecificationProperty{
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
//
//   							// the properties below are optional
//   							InvocationLabel: jsii.String("invocationLabel"),
//   						},
//   						ConfirmationConditional: &ConditionalSpecificationProperty{
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
//   						ConfirmationNextStep: &DialogStateProperty{
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
//   						ConfirmationResponse: &ResponseSpecificationProperty{
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
//   						DeclinationConditional: &ConditionalSpecificationProperty{
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
//   						DeclinationNextStep: &DialogStateProperty{
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
//   						DeclinationResponse: &ResponseSpecificationProperty{
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
//   						ElicitationCodeHook: &ElicitationCodeHookInvocationSettingProperty{
//   							EnableCodeHookInvocation: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							InvocationLabel: jsii.String("invocationLabel"),
//   						},
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
//   						IsActive: jsii.Boolean(false),
//   					},
//   					KendraConfiguration: &KendraConfigurationProperty{
//   						KendraIndex: jsii.String("kendraIndex"),
//
//   						// the properties below are optional
//   						QueryFilterString: jsii.String("queryFilterString"),
//   						QueryFilterStringEnabled: jsii.Boolean(false),
//   					},
//   					OutputContexts: []interface{}{
//   						&OutputContextProperty{
//   							Name: jsii.String("name"),
//   							TimeToLiveInSeconds: jsii.Number(123),
//   							TurnsToLive: jsii.Number(123),
//   						},
//   					},
//   					ParentIntentSignature: jsii.String("parentIntentSignature"),
//   					SampleUtterances: []interface{}{
//   						&SampleUtteranceProperty{
//   							Utterance: jsii.String("utterance"),
//   						},
//   					},
//   					SlotPriorities: []interface{}{
//   						&SlotPriorityProperty{
//   							Priority: jsii.Number(123),
//   							SlotName: jsii.String("slotName"),
//   						},
//   					},
//   					Slots: []interface{}{
//   						&SlotProperty{
//   							Name: jsii.String("name"),
//   							SlotTypeName: jsii.String("slotTypeName"),
//   							ValueElicitationSetting: &SlotValueElicitationSettingProperty{
//   								SlotConstraint: jsii.String("slotConstraint"),
//
//   								// the properties below are optional
//   								DefaultValueSpecification: &SlotDefaultValueSpecificationProperty{
//   									DefaultValueList: []interface{}{
//   										&SlotDefaultValueProperty{
//   											DefaultValue: jsii.String("defaultValue"),
//   										},
//   									},
//   								},
//   								PromptSpecification: &PromptSpecificationProperty{
//   									MaxRetries: jsii.Number(123),
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
//   									MessageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   									PromptAttemptsSpecification: map[string]interface{}{
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
//   								SampleUtterances: []interface{}{
//   									&SampleUtteranceProperty{
//   										Utterance: jsii.String("utterance"),
//   									},
//   								},
//   								SlotCaptureSetting: &SlotCaptureSettingProperty{
//   									CaptureConditional: &ConditionalSpecificationProperty{
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
//   									CaptureNextStep: &DialogStateProperty{
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
//   									CaptureResponse: &ResponseSpecificationProperty{
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
//   									CodeHook: &DialogCodeHookInvocationSettingProperty{
//   										EnableCodeHookInvocation: jsii.Boolean(false),
//   										IsActive: jsii.Boolean(false),
//   										PostCodeHookSpecification: &PostDialogCodeHookInvocationSpecificationProperty{
//   											FailureConditional: &ConditionalSpecificationProperty{
//   												ConditionalBranches: []interface{}{
//   													&ConditionalBranchProperty{
//   														Condition: &ConditionProperty{
//   															ExpressionString: jsii.String("expressionString"),
//   														},
//   														Name: jsii.String("name"),
//   														NextStep: &DialogStateProperty{
//   															DialogAction: &DialogActionProperty{
//   																Type: jsii.String("type"),
//
//   																// the properties below are optional
//   																SlotToElicit: jsii.String("slotToElicit"),
//   																SuppressNextMessage: jsii.Boolean(false),
//   															},
//   															Intent: &IntentOverrideProperty{
//   																Name: jsii.String("name"),
//   																Slots: []interface{}{
//   																	&SlotValueOverrideMapProperty{
//   																		SlotName: jsii.String("slotName"),
//   																		SlotValueOverride: &slotValueOverrideProperty{
//   																			Shape: jsii.String("shape"),
//   																			Value: &SlotValueProperty{
//   																				InterpretedValue: jsii.String("interpretedValue"),
//   																			},
//   																			Values: []interface{}{
//   																				slotValueOverrideProperty_,
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//   															SessionAttributes: []interface{}{
//   																&SessionAttributeProperty{
//   																	Key: jsii.String("key"),
//
//   																	// the properties below are optional
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														Response: &ResponseSpecificationProperty{
//   															MessageGroupsList: []interface{}{
//   																&MessageGroupProperty{
//   																	Message: &MessageProperty{
//   																		CustomPayload: &CustomPayloadProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		ImageResponseCard: &ImageResponseCardProperty{
//   																			Title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			Buttons: []interface{}{
//   																				&ButtonProperty{
//   																					Text: jsii.String("text"),
//   																					Value: jsii.String("value"),
//   																				},
//   																			},
//   																			ImageUrl: jsii.String("imageUrl"),
//   																			Subtitle: jsii.String("subtitle"),
//   																		},
//   																		PlainTextMessage: &PlainTextMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		SsmlMessage: &SSMLMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//
//   																	// the properties below are optional
//   																	Variations: []interface{}{
//   																		&MessageProperty{
//   																			CustomPayload: &CustomPayloadProperty{
//   																				Value: jsii.String("value"),
//   																			},
//   																			ImageResponseCard: &ImageResponseCardProperty{
//   																				Title: jsii.String("title"),
//
//   																				// the properties below are optional
//   																				Buttons: []interface{}{
//   																					&ButtonProperty{
//   																						Text: jsii.String("text"),
//   																						Value: jsii.String("value"),
//   																					},
//   																				},
//   																				ImageUrl: jsii.String("imageUrl"),
//   																				Subtitle: jsii.String("subtitle"),
//   																			},
//   																			PlainTextMessage: &PlainTextMessageProperty{
//   																				Value: jsii.String("value"),
//   																			},
//   																			SsmlMessage: &SSMLMessageProperty{
//   																				Value: jsii.String("value"),
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//
//   															// the properties below are optional
//   															AllowInterrupt: jsii.Boolean(false),
//   														},
//   													},
//   												},
//   												DefaultBranch: &DefaultConditionalBranchProperty{
//   													NextStep: &DialogStateProperty{
//   														DialogAction: &DialogActionProperty{
//   															Type: jsii.String("type"),
//
//   															// the properties below are optional
//   															SlotToElicit: jsii.String("slotToElicit"),
//   															SuppressNextMessage: jsii.Boolean(false),
//   														},
//   														Intent: &IntentOverrideProperty{
//   															Name: jsii.String("name"),
//   															Slots: []interface{}{
//   																&SlotValueOverrideMapProperty{
//   																	SlotName: jsii.String("slotName"),
//   																	SlotValueOverride: &slotValueOverrideProperty{
//   																		Shape: jsii.String("shape"),
//   																		Value: &SlotValueProperty{
//   																			InterpretedValue: jsii.String("interpretedValue"),
//   																		},
//   																		Values: []interface{}{
//   																			slotValueOverrideProperty_,
//   																		},
//   																	},
//   																},
//   															},
//   														},
//   														SessionAttributes: []interface{}{
//   															&SessionAttributeProperty{
//   																Key: jsii.String("key"),
//
//   																// the properties below are optional
//   																Value: jsii.String("value"),
//   															},
//   														},
//   													},
//   													Response: &ResponseSpecificationProperty{
//   														MessageGroupsList: []interface{}{
//   															&MessageGroupProperty{
//   																Message: &MessageProperty{
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
//
//   																// the properties below are optional
//   																Variations: []interface{}{
//   																	&MessageProperty{
//   																		CustomPayload: &CustomPayloadProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		ImageResponseCard: &ImageResponseCardProperty{
//   																			Title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			Buttons: []interface{}{
//   																				&ButtonProperty{
//   																					Text: jsii.String("text"),
//   																					Value: jsii.String("value"),
//   																				},
//   																			},
//   																			ImageUrl: jsii.String("imageUrl"),
//   																			Subtitle: jsii.String("subtitle"),
//   																		},
//   																		PlainTextMessage: &PlainTextMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		SsmlMessage: &SSMLMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														AllowInterrupt: jsii.Boolean(false),
//   													},
//   												},
//   												IsActive: jsii.Boolean(false),
//   											},
//   											FailureNextStep: &DialogStateProperty{
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
//   											FailureResponse: &ResponseSpecificationProperty{
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
//   											SuccessConditional: &ConditionalSpecificationProperty{
//   												ConditionalBranches: []interface{}{
//   													&ConditionalBranchProperty{
//   														Condition: &ConditionProperty{
//   															ExpressionString: jsii.String("expressionString"),
//   														},
//   														Name: jsii.String("name"),
//   														NextStep: &DialogStateProperty{
//   															DialogAction: &DialogActionProperty{
//   																Type: jsii.String("type"),
//
//   																// the properties below are optional
//   																SlotToElicit: jsii.String("slotToElicit"),
//   																SuppressNextMessage: jsii.Boolean(false),
//   															},
//   															Intent: &IntentOverrideProperty{
//   																Name: jsii.String("name"),
//   																Slots: []interface{}{
//   																	&SlotValueOverrideMapProperty{
//   																		SlotName: jsii.String("slotName"),
//   																		SlotValueOverride: &slotValueOverrideProperty{
//   																			Shape: jsii.String("shape"),
//   																			Value: &SlotValueProperty{
//   																				InterpretedValue: jsii.String("interpretedValue"),
//   																			},
//   																			Values: []interface{}{
//   																				slotValueOverrideProperty_,
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//   															SessionAttributes: []interface{}{
//   																&SessionAttributeProperty{
//   																	Key: jsii.String("key"),
//
//   																	// the properties below are optional
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														Response: &ResponseSpecificationProperty{
//   															MessageGroupsList: []interface{}{
//   																&MessageGroupProperty{
//   																	Message: &MessageProperty{
//   																		CustomPayload: &CustomPayloadProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		ImageResponseCard: &ImageResponseCardProperty{
//   																			Title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			Buttons: []interface{}{
//   																				&ButtonProperty{
//   																					Text: jsii.String("text"),
//   																					Value: jsii.String("value"),
//   																				},
//   																			},
//   																			ImageUrl: jsii.String("imageUrl"),
//   																			Subtitle: jsii.String("subtitle"),
//   																		},
//   																		PlainTextMessage: &PlainTextMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		SsmlMessage: &SSMLMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//
//   																	// the properties below are optional
//   																	Variations: []interface{}{
//   																		&MessageProperty{
//   																			CustomPayload: &CustomPayloadProperty{
//   																				Value: jsii.String("value"),
//   																			},
//   																			ImageResponseCard: &ImageResponseCardProperty{
//   																				Title: jsii.String("title"),
//
//   																				// the properties below are optional
//   																				Buttons: []interface{}{
//   																					&ButtonProperty{
//   																						Text: jsii.String("text"),
//   																						Value: jsii.String("value"),
//   																					},
//   																				},
//   																				ImageUrl: jsii.String("imageUrl"),
//   																				Subtitle: jsii.String("subtitle"),
//   																			},
//   																			PlainTextMessage: &PlainTextMessageProperty{
//   																				Value: jsii.String("value"),
//   																			},
//   																			SsmlMessage: &SSMLMessageProperty{
//   																				Value: jsii.String("value"),
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//
//   															// the properties below are optional
//   															AllowInterrupt: jsii.Boolean(false),
//   														},
//   													},
//   												},
//   												DefaultBranch: &DefaultConditionalBranchProperty{
//   													NextStep: &DialogStateProperty{
//   														DialogAction: &DialogActionProperty{
//   															Type: jsii.String("type"),
//
//   															// the properties below are optional
//   															SlotToElicit: jsii.String("slotToElicit"),
//   															SuppressNextMessage: jsii.Boolean(false),
//   														},
//   														Intent: &IntentOverrideProperty{
//   															Name: jsii.String("name"),
//   															Slots: []interface{}{
//   																&SlotValueOverrideMapProperty{
//   																	SlotName: jsii.String("slotName"),
//   																	SlotValueOverride: &slotValueOverrideProperty{
//   																		Shape: jsii.String("shape"),
//   																		Value: &SlotValueProperty{
//   																			InterpretedValue: jsii.String("interpretedValue"),
//   																		},
//   																		Values: []interface{}{
//   																			slotValueOverrideProperty_,
//   																		},
//   																	},
//   																},
//   															},
//   														},
//   														SessionAttributes: []interface{}{
//   															&SessionAttributeProperty{
//   																Key: jsii.String("key"),
//
//   																// the properties below are optional
//   																Value: jsii.String("value"),
//   															},
//   														},
//   													},
//   													Response: &ResponseSpecificationProperty{
//   														MessageGroupsList: []interface{}{
//   															&MessageGroupProperty{
//   																Message: &MessageProperty{
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
//
//   																// the properties below are optional
//   																Variations: []interface{}{
//   																	&MessageProperty{
//   																		CustomPayload: &CustomPayloadProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		ImageResponseCard: &ImageResponseCardProperty{
//   																			Title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			Buttons: []interface{}{
//   																				&ButtonProperty{
//   																					Text: jsii.String("text"),
//   																					Value: jsii.String("value"),
//   																				},
//   																			},
//   																			ImageUrl: jsii.String("imageUrl"),
//   																			Subtitle: jsii.String("subtitle"),
//   																		},
//   																		PlainTextMessage: &PlainTextMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		SsmlMessage: &SSMLMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														AllowInterrupt: jsii.Boolean(false),
//   													},
//   												},
//   												IsActive: jsii.Boolean(false),
//   											},
//   											SuccessNextStep: &DialogStateProperty{
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
//   											SuccessResponse: &ResponseSpecificationProperty{
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
//   											TimeoutConditional: &ConditionalSpecificationProperty{
//   												ConditionalBranches: []interface{}{
//   													&ConditionalBranchProperty{
//   														Condition: &ConditionProperty{
//   															ExpressionString: jsii.String("expressionString"),
//   														},
//   														Name: jsii.String("name"),
//   														NextStep: &DialogStateProperty{
//   															DialogAction: &DialogActionProperty{
//   																Type: jsii.String("type"),
//
//   																// the properties below are optional
//   																SlotToElicit: jsii.String("slotToElicit"),
//   																SuppressNextMessage: jsii.Boolean(false),
//   															},
//   															Intent: &IntentOverrideProperty{
//   																Name: jsii.String("name"),
//   																Slots: []interface{}{
//   																	&SlotValueOverrideMapProperty{
//   																		SlotName: jsii.String("slotName"),
//   																		SlotValueOverride: &slotValueOverrideProperty{
//   																			Shape: jsii.String("shape"),
//   																			Value: &SlotValueProperty{
//   																				InterpretedValue: jsii.String("interpretedValue"),
//   																			},
//   																			Values: []interface{}{
//   																				slotValueOverrideProperty_,
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//   															SessionAttributes: []interface{}{
//   																&SessionAttributeProperty{
//   																	Key: jsii.String("key"),
//
//   																	// the properties below are optional
//   																	Value: jsii.String("value"),
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														Response: &ResponseSpecificationProperty{
//   															MessageGroupsList: []interface{}{
//   																&MessageGroupProperty{
//   																	Message: &MessageProperty{
//   																		CustomPayload: &CustomPayloadProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		ImageResponseCard: &ImageResponseCardProperty{
//   																			Title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			Buttons: []interface{}{
//   																				&ButtonProperty{
//   																					Text: jsii.String("text"),
//   																					Value: jsii.String("value"),
//   																				},
//   																			},
//   																			ImageUrl: jsii.String("imageUrl"),
//   																			Subtitle: jsii.String("subtitle"),
//   																		},
//   																		PlainTextMessage: &PlainTextMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		SsmlMessage: &SSMLMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//
//   																	// the properties below are optional
//   																	Variations: []interface{}{
//   																		&MessageProperty{
//   																			CustomPayload: &CustomPayloadProperty{
//   																				Value: jsii.String("value"),
//   																			},
//   																			ImageResponseCard: &ImageResponseCardProperty{
//   																				Title: jsii.String("title"),
//
//   																				// the properties below are optional
//   																				Buttons: []interface{}{
//   																					&ButtonProperty{
//   																						Text: jsii.String("text"),
//   																						Value: jsii.String("value"),
//   																					},
//   																				},
//   																				ImageUrl: jsii.String("imageUrl"),
//   																				Subtitle: jsii.String("subtitle"),
//   																			},
//   																			PlainTextMessage: &PlainTextMessageProperty{
//   																				Value: jsii.String("value"),
//   																			},
//   																			SsmlMessage: &SSMLMessageProperty{
//   																				Value: jsii.String("value"),
//   																			},
//   																		},
//   																	},
//   																},
//   															},
//
//   															// the properties below are optional
//   															AllowInterrupt: jsii.Boolean(false),
//   														},
//   													},
//   												},
//   												DefaultBranch: &DefaultConditionalBranchProperty{
//   													NextStep: &DialogStateProperty{
//   														DialogAction: &DialogActionProperty{
//   															Type: jsii.String("type"),
//
//   															// the properties below are optional
//   															SlotToElicit: jsii.String("slotToElicit"),
//   															SuppressNextMessage: jsii.Boolean(false),
//   														},
//   														Intent: &IntentOverrideProperty{
//   															Name: jsii.String("name"),
//   															Slots: []interface{}{
//   																&SlotValueOverrideMapProperty{
//   																	SlotName: jsii.String("slotName"),
//   																	SlotValueOverride: &slotValueOverrideProperty{
//   																		Shape: jsii.String("shape"),
//   																		Value: &SlotValueProperty{
//   																			InterpretedValue: jsii.String("interpretedValue"),
//   																		},
//   																		Values: []interface{}{
//   																			slotValueOverrideProperty_,
//   																		},
//   																	},
//   																},
//   															},
//   														},
//   														SessionAttributes: []interface{}{
//   															&SessionAttributeProperty{
//   																Key: jsii.String("key"),
//
//   																// the properties below are optional
//   																Value: jsii.String("value"),
//   															},
//   														},
//   													},
//   													Response: &ResponseSpecificationProperty{
//   														MessageGroupsList: []interface{}{
//   															&MessageGroupProperty{
//   																Message: &MessageProperty{
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
//
//   																// the properties below are optional
//   																Variations: []interface{}{
//   																	&MessageProperty{
//   																		CustomPayload: &CustomPayloadProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		ImageResponseCard: &ImageResponseCardProperty{
//   																			Title: jsii.String("title"),
//
//   																			// the properties below are optional
//   																			Buttons: []interface{}{
//   																				&ButtonProperty{
//   																					Text: jsii.String("text"),
//   																					Value: jsii.String("value"),
//   																				},
//   																			},
//   																			ImageUrl: jsii.String("imageUrl"),
//   																			Subtitle: jsii.String("subtitle"),
//   																		},
//   																		PlainTextMessage: &PlainTextMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																		SsmlMessage: &SSMLMessageProperty{
//   																			Value: jsii.String("value"),
//   																		},
//   																	},
//   																},
//   															},
//   														},
//
//   														// the properties below are optional
//   														AllowInterrupt: jsii.Boolean(false),
//   													},
//   												},
//   												IsActive: jsii.Boolean(false),
//   											},
//   											TimeoutNextStep: &DialogStateProperty{
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
//   											TimeoutResponse: &ResponseSpecificationProperty{
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
//
//   										// the properties below are optional
//   										InvocationLabel: jsii.String("invocationLabel"),
//   									},
//   									ElicitationCodeHook: &ElicitationCodeHookInvocationSettingProperty{
//   										EnableCodeHookInvocation: jsii.Boolean(false),
//
//   										// the properties below are optional
//   										InvocationLabel: jsii.String("invocationLabel"),
//   									},
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
//   								},
//   								WaitAndContinueSpecification: &WaitAndContinueSpecificationProperty{
//   									ContinueResponse: &ResponseSpecificationProperty{
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
//   									WaitingResponse: &ResponseSpecificationProperty{
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
//
//   									// the properties below are optional
//   									IsActive: jsii.Boolean(false),
//   									StillWaitingResponse: &StillWaitingResponseSpecificationProperty{
//   										FrequencyInSeconds: jsii.Number(123),
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
//   										TimeoutInSeconds: jsii.Number(123),
//
//   										// the properties below are optional
//   										AllowInterrupt: jsii.Boolean(false),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							Description: jsii.String("description"),
//   							MultipleValuesSetting: &MultipleValuesSettingProperty{
//   								AllowMultipleValues: jsii.Boolean(false),
//   							},
//   							ObfuscationSetting: &ObfuscationSettingProperty{
//   								ObfuscationSettingType: jsii.String("obfuscationSettingType"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			SlotTypes: []interface{}{
//   				&SlotTypeProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Description: jsii.String("description"),
//   					ExternalSourceSetting: &ExternalSourceSettingProperty{
//   						GrammarSlotTypeSetting: &GrammarSlotTypeSettingProperty{
//   							Source: &GrammarSlotTypeSourceProperty{
//   								S3BucketName: jsii.String("s3BucketName"),
//   								S3ObjectKey: jsii.String("s3ObjectKey"),
//
//   								// the properties below are optional
//   								KmsKeyArn: jsii.String("kmsKeyArn"),
//   							},
//   						},
//   					},
//   					ParentSlotTypeSignature: jsii.String("parentSlotTypeSignature"),
//   					SlotTypeValues: []interface{}{
//   						&SlotTypeValueProperty{
//   							SampleValue: &SampleValueProperty{
//   								Value: jsii.String("value"),
//   							},
//
//   							// the properties below are optional
//   							Synonyms: []interface{}{
//   								&SampleValueProperty{
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					ValueSelectionSetting: &SlotValueSelectionSettingProperty{
//   						ResolutionStrategy: jsii.String("resolutionStrategy"),
//
//   						// the properties below are optional
//   						AdvancedRecognitionSetting: &AdvancedRecognitionSettingProperty{
//   							AudioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   						},
//   						RegexFilter: &SlotValueRegexFilterProperty{
//   							Pattern: jsii.String("pattern"),
//   						},
//   					},
//   				},
//   			},
//   			VoiceSettings: &VoiceSettingsProperty{
//   				VoiceId: jsii.String("voiceId"),
//
//   				// the properties below are optional
//   				Engine: jsii.String("engine"),
//   			},
//   		},
//   	},
//   	BotTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	TestBotAliasSettings: &TestBotAliasSettingsProperty{
//   		BotAliasLocaleSettings: []interface{}{
//   			&BotAliasLocaleSettingsItemProperty{
//   				BotAliasLocaleSetting: &BotAliasLocaleSettingsProperty{
//   					Enabled: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					CodeHookSpecification: &CodeHookSpecificationProperty{
//   						LambdaCodeHook: &LambdaCodeHookProperty{
//   							CodeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   							LambdaArn: jsii.String("lambdaArn"),
//   						},
//   					},
//   				},
//   				LocaleId: jsii.String("localeId"),
//   			},
//   		},
//   		ConversationLogSettings: &ConversationLogSettingsProperty{
//   			AudioLogSettings: []interface{}{
//   				&AudioLogSettingProperty{
//   					Destination: &AudioLogDestinationProperty{
//   						S3Bucket: &S3BucketLogDestinationProperty{
//   							LogPrefix: jsii.String("logPrefix"),
//   							S3BucketArn: jsii.String("s3BucketArn"),
//
//   							// the properties below are optional
//   							KmsKeyArn: jsii.String("kmsKeyArn"),
//   						},
//   					},
//   					Enabled: jsii.Boolean(false),
//   				},
//   			},
//   			TextLogSettings: []interface{}{
//   				&TextLogSettingProperty{
//   					Destination: &TextLogDestinationProperty{
//   						CloudWatch: &CloudWatchLogGroupLogDestinationProperty{
//   							CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   							LogPrefix: jsii.String("logPrefix"),
//   						},
//   					},
//   					Enabled: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		Description: jsii.String("description"),
//   		SentimentAnalysisSettings: sentimentAnalysisSettings,
//   	},
//   	TestBotAliasTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

