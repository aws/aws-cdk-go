package awslex

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslex/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::Lex::Bot`.
//
// > Amazon Lex V2 is the only supported version in AWS CloudFormation .
//
// Specifies an Amazon Lex conversational bot.
//
// You must configure an intent based on the AMAZON.FallbackIntent built-in intent. If you don't add one, creating the bot will fail.
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
//   cfnBot := awscdk.Aws_lex.NewCfnBot(this, jsii.String("MyCfnBot"), &cfnBotProps{
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
//   })
//
type CfnBot interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the bot.
	AttrArn() *string
	// The unique identifier of the bot.
	AttrId() *string
	// Indicates whether Amazon Lex V2 should automatically build the locales for the bot after a change.
	AutoBuildBotLocales() interface{}
	SetAutoBuildBotLocales(val interface{})
	// The Amazon S3 location of files used to import a bot.
	//
	// The files must be in the import format specified in [JSON format for importing and exporting](https://docs.aws.amazon.com/lexv2/latest/dg/import-export-format.html) in the *Amazon Lex developer guide.*
	BotFileS3Location() interface{}
	SetBotFileS3Location(val interface{})
	// A list of locales for the bot.
	BotLocales() interface{}
	SetBotLocales(val interface{})
	// A list of tags to add to the bot.
	//
	// You can only add tags when you import a bot. You can't use the `UpdateBot` operation to update tags. To update tags, use the `TagResource` operation.
	BotTags() interface{}
	SetBotTags(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Provides information on additional privacy protections Amazon Lex should use with the bot's data.
	DataPrivacy() interface{}
	SetDataPrivacy(val interface{})
	// The description of the version.
	Description() *string
	SetDescription(val *string)
	// The time, in seconds, that Amazon Lex should keep information about a user's conversation with the bot.
	//
	// A user interaction remains active for the amount of time specified. If no conversation occurs during this time, the session expires and Amazon Lex deletes any data provided before the timeout.
	//
	// You can specify between 60 (1 minute) and 86,400 (24 hours) seconds.
	IdleSessionTtlInSeconds() *float64
	SetIdleSessionTtlInSeconds(val *float64)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the field to filter the list of bots.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of the IAM role used to build and run the bot.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Specifies configuration settings for the alias used to test the bot.
	//
	// If the `TestBotAliasSettings` property is not specified, the settings are configured with default values.
	TestBotAliasSettings() interface{}
	SetTestBotAliasSettings(val interface{})
	// A list of tags to add to the test alias for a bot.
	//
	// You can only add tags when you import a bot. You can't use the `UpdateAlias` operation to update tags. To update tags on the test alias, use the `TagResource` operation.
	TestBotAliasTags() interface{}
	SetTestBotAliasTags(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBot
type jsiiProxy_CfnBot struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBot) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) AutoBuildBotLocales() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoBuildBotLocales",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) BotFileS3Location() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botFileS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) BotLocales() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botLocales",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) BotTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) DataPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) IdleSessionTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) TestBotAliasSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"testBotAliasSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) TestBotAliasTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"testBotAliasTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBot) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lex::Bot`.
func NewCfnBot(scope awscdk.Construct, id *string, props *CfnBotProps) CfnBot {
	_init_.Initialize()

	if err := validateNewCfnBotParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBot{}

	_jsii_.Create(
		"monocdk.aws_lex.CfnBot",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::Bot`.
func NewCfnBot_Override(c CfnBot, scope awscdk.Construct, id *string, props *CfnBotProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lex.CfnBot",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBot)SetAutoBuildBotLocales(val interface{}) {
	if err := j.validateSetAutoBuildBotLocalesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoBuildBotLocales",
		val,
	)
}

func (j *jsiiProxy_CfnBot)SetBotFileS3Location(val interface{}) {
	if err := j.validateSetBotFileS3LocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botFileS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnBot)SetBotLocales(val interface{}) {
	if err := j.validateSetBotLocalesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botLocales",
		val,
	)
}

func (j *jsiiProxy_CfnBot)SetBotTags(val interface{}) {
	if err := j.validateSetBotTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botTags",
		val,
	)
}

func (j *jsiiProxy_CfnBot)SetDataPrivacy(val interface{}) {
	if err := j.validateSetDataPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPrivacy",
		val,
	)
}

func (j *jsiiProxy_CfnBot)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnBot)SetIdleSessionTtlInSeconds(val *float64) {
	if err := j.validateSetIdleSessionTtlInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleSessionTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnBot)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnBot)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnBot)SetTestBotAliasSettings(val interface{}) {
	if err := j.validateSetTestBotAliasSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testBotAliasSettings",
		val,
	)
}

func (j *jsiiProxy_CfnBot)SetTestBotAliasTags(val interface{}) {
	if err := j.validateSetTestBotAliasTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testBotAliasTags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnBot_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBot_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBot",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnBot_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnBot_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBot",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnBot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lex.CfnBot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBot_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lex.CfnBot",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBot) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBot) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBot) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBot) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBot) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBot) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBot) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBot) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBot) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBot) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBot) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBot) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnBot) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBot) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBot) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnBot) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBot) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBot) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnBot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBot) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBot) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

