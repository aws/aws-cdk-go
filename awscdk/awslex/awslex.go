package awslex

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslex/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Lex::Bot`.
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
//   						postFulfillmentStatusSpecification: &postFulfillmentStatusSpecificationProperty{
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
//
//   						// the properties below are optional
//   						isActive: jsii.Boolean(false),
//   					},
//   					intentConfirmationSetting: &intentConfirmationSettingProperty{
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
//   						},
//
//   						// the properties below are optional
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
//   								},
//   								sampleUtterances: []interface{}{
//   									&sampleUtteranceProperty{
//   										utterance: jsii.String("utterance"),
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
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The name of the field to filter the list of bots.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the IAM role used to build and run the bot.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
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
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
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

func (j *jsiiProxy_CfnBot) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnBot(scope constructs.Construct, id *string, props *CfnBotProps) CfnBot {
	_init_.Initialize()

	j := jsiiProxy_CfnBot{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBot",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::Bot`.
func NewCfnBot_Override(c CfnBot, scope constructs.Construct, id *string, props *CfnBotProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBot",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBot) SetAutoBuildBotLocales(val interface{}) {
	_jsii_.Set(
		j,
		"autoBuildBotLocales",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetBotFileS3Location(val interface{}) {
	_jsii_.Set(
		j,
		"botFileS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetBotLocales(val interface{}) {
	_jsii_.Set(
		j,
		"botLocales",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetBotTags(val interface{}) {
	_jsii_.Set(
		j,
		"botTags",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetDataPrivacy(val interface{}) {
	_jsii_.Set(
		j,
		"dataPrivacy",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetIdleSessionTtlInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"idleSessionTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetTestBotAliasSettings(val interface{}) {
	_jsii_.Set(
		j,
		"testBotAliasSettings",
		val,
	)
}

func (j *jsiiProxy_CfnBot) SetTestBotAliasTags(val interface{}) {
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
func CfnBot_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBot",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBot_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBot",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnBot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBot",
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
		"aws-cdk-lib.aws_lex.CfnBot",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBot) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBot) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBot) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBot) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBot) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBot) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBot) GetAtt(attributeName *string) awscdk.Reference {
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
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBot) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBot) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnBot) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies settings that enable advanced audio recognition for slot values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedRecognitionSettingProperty := &advancedRecognitionSettingProperty{
//   	audioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   }
//
type CfnBot_AdvancedRecognitionSettingProperty struct {
	// Specifies that Amazon Lex should use slot values as a custom vocabulary when recognizing user utterances.
	AudioRecognitionStrategy *string `field:"optional" json:"audioRecognitionStrategy" yaml:"audioRecognitionStrategy"`
}

// Specifies the location of audio log files collected when conversation logging is enabled for a bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioLogDestinationProperty := &audioLogDestinationProperty{
//   	s3Bucket: &s3BucketLogDestinationProperty{
//   		logPrefix: jsii.String("logPrefix"),
//   		s3BucketArn: jsii.String("s3BucketArn"),
//
//   		// the properties below are optional
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
type CfnBot_AudioLogDestinationProperty struct {
	// Specifies the Amazon S3 bucket where the audio files are stored.
	S3Bucket interface{} `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
}

// Specifies settings for logging the audio of conversations between Amazon Lex and a user.
//
// You specify whether to log audio and the Amazon S3 bucket where the audio file is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioLogSettingProperty := &audioLogSettingProperty{
//   	destination: &audioLogDestinationProperty{
//   		s3Bucket: &s3BucketLogDestinationProperty{
//   			logPrefix: jsii.String("logPrefix"),
//   			s3BucketArn: jsii.String("s3BucketArn"),
//
//   			// the properties below are optional
//   			kmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnBot_AudioLogSettingProperty struct {
	// Specifies the location of the audio log files collected when conversation logging is enabled for a bot.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether audio logging is enabled for the bot.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

// Specifies locale settings for a single locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botAliasLocaleSettingsItemProperty := &botAliasLocaleSettingsItemProperty{
//   	botAliasLocaleSetting: &botAliasLocaleSettingsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		codeHookSpecification: &codeHookSpecificationProperty{
//   			lambdaCodeHook: &lambdaCodeHookProperty{
//   				codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   				lambdaArn: jsii.String("lambdaArn"),
//   			},
//   		},
//   	},
//   	localeId: jsii.String("localeId"),
//   }
//
type CfnBot_BotAliasLocaleSettingsItemProperty struct {
	// Specifies locale settings for a locale.
	BotAliasLocaleSetting interface{} `field:"required" json:"botAliasLocaleSetting" yaml:"botAliasLocaleSetting"`
	// Specifies the locale that the settings apply to.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}

// Specifies settings that are unique to a locale.
//
// For example, you can use a different Lambda function for each locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botAliasLocaleSettingsProperty := &botAliasLocaleSettingsProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	codeHookSpecification: &codeHookSpecificationProperty{
//   		lambdaCodeHook: &lambdaCodeHookProperty{
//   			codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   			lambdaArn: jsii.String("lambdaArn"),
//   		},
//   	},
//   }
//
type CfnBot_BotAliasLocaleSettingsProperty struct {
	// Specifies whether the locale is enabled for the bot.
	//
	// If the value is false, the locale isn't available for use.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the Lambda function to use in this locale.
	CodeHookSpecification interface{} `field:"optional" json:"codeHookSpecification" yaml:"codeHookSpecification"`
}

// Provides configuration information for a locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   				postFulfillmentStatusSpecification: &postFulfillmentStatusSpecificationProperty{
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
//
//   				// the properties below are optional
//   				isActive: jsii.Boolean(false),
//   			},
//   			intentConfirmationSetting: &intentConfirmationSettingProperty{
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
//   				},
//
//   				// the properties below are optional
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
//   						},
//   						sampleUtterances: []interface{}{
//   							&sampleUtteranceProperty{
//   								utterance: jsii.String("utterance"),
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
//   	},
//   }
//
type CfnBot_BotLocaleProperty struct {
	// The identifier of the language and locale that the bot will be used in.
	//
	// The string must match one of the supported locales.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents. You must configure an AMAZON.FallbackIntent. AMAZON.KendraSearchIntent is only inserted if it is configured for the bot.
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
	// Identifies the Amazon Polly voice used for audio interaction with the user.
	VoiceSettings interface{} `field:"optional" json:"voiceSettings" yaml:"voiceSettings"`
}

// Describes a button to use on a response card used to gather slot values from a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   buttonProperty := &buttonProperty{
//   	text: jsii.String("text"),
//   	value: jsii.String("value"),
//   }
//
type CfnBot_ButtonProperty struct {
	// The text that appears on the button.
	//
	// Use this to tell the user the value that is returned when they choose this button.
	Text *string `field:"required" json:"text" yaml:"text"`
	// The value returned to Amazon Lex when the user chooses this button.
	//
	// This must be one of the slot values configured for the slot.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Specifies the Amazon CloudWatch Logs log group where text and metadata logs are delivered.
//
// The log group must exist before you enable logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogGroupLogDestinationProperty := &cloudWatchLogGroupLogDestinationProperty{
//   	cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   	logPrefix: jsii.String("logPrefix"),
//   }
//
type CfnBot_CloudWatchLogGroupLogDestinationProperty struct {
	// Specifies the Amazon Resource Name (ARN) of the log group where text and metadata logs are delivered.
	CloudWatchLogGroupArn *string `field:"required" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// Specifies the prefix of the log stream name within the log group that you specified.
	LogPrefix *string `field:"required" json:"logPrefix" yaml:"logPrefix"`
}

// Specifies information about code hooks that Amazon Lex calls during a conversation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeHookSpecificationProperty := &codeHookSpecificationProperty{
//   	lambdaCodeHook: &lambdaCodeHookProperty{
//   		codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   		lambdaArn: jsii.String("lambdaArn"),
//   	},
//   }
//
type CfnBot_CodeHookSpecificationProperty struct {
	// Specifies a Lambda function that verifies requests to a bot or fulfills the user's request to a bot.
	LambdaCodeHook interface{} `field:"required" json:"lambdaCodeHook" yaml:"lambdaCodeHook"`
}

// Specifies settings that manage logging that saves audio, text, and metadata for the conversations with your users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conversationLogSettingsProperty := &conversationLogSettingsProperty{
//   	audioLogSettings: []interface{}{
//   		&audioLogSettingProperty{
//   			destination: &audioLogDestinationProperty{
//   				s3Bucket: &s3BucketLogDestinationProperty{
//   					logPrefix: jsii.String("logPrefix"),
//   					s3BucketArn: jsii.String("s3BucketArn"),
//
//   					// the properties below are optional
//   					kmsKeyArn: jsii.String("kmsKeyArn"),
//   				},
//   			},
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	textLogSettings: []interface{}{
//   		&textLogSettingProperty{
//   			destination: &textLogDestinationProperty{
//   				cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   					cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   					logPrefix: jsii.String("logPrefix"),
//   				},
//   			},
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnBot_ConversationLogSettingsProperty struct {
	// Specifies the Amazon S3 settings for logging audio to an S3 bucket.
	AudioLogSettings interface{} `field:"optional" json:"audioLogSettings" yaml:"audioLogSettings"`
	// Specifies settings to enable text conversation logs.
	//
	// You specify the Amazon CloudWatch Logs log group and whether logs should be stored for an alias.
	TextLogSettings interface{} `field:"optional" json:"textLogSettings" yaml:"textLogSettings"`
}

// A custom response string that Amazon Lex sends to your application.
//
// You define the content and structure of the string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPayloadProperty := &customPayloadProperty{
//   	value: jsii.String("value"),
//   }
//
type CfnBot_CustomPayloadProperty struct {
	// The string that is sent to your application.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Specifies an entry in a custom vocabulary.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customVocabularyItemProperty := &customVocabularyItemProperty{
//   	phrase: jsii.String("phrase"),
//
//   	// the properties below are optional
//   	weight: jsii.Number(123),
//   }
//
type CfnBot_CustomVocabularyItemProperty struct {
	// Specifies 1 - 4 words that should be recognized.
	Phrase *string `field:"required" json:"phrase" yaml:"phrase"`
	// Specifies the degree to which the phrase recognition is boosted.
	//
	// The default value is 1.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// Specifies a custom vocabulary.
//
// A custom vocabulary is a list of words that you expect to be used during a conversation with your bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customVocabularyProperty := &customVocabularyProperty{
//   	customVocabularyItems: []interface{}{
//   		&customVocabularyItemProperty{
//   			phrase: jsii.String("phrase"),
//
//   			// the properties below are optional
//   			weight: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnBot_CustomVocabularyProperty struct {
	// Specifies a list of words that you expect to be used during a conversation with your bot.
	CustomVocabularyItems interface{} `field:"required" json:"customVocabularyItems" yaml:"customVocabularyItems"`
}

// Specifies whether an intent uses the dialog code hook during conversations with a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dialogCodeHookSettingProperty := &dialogCodeHookSettingProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnBot_DialogCodeHookSettingProperty struct {
	// Indicates whether an intent uses the dialog code hook during a conversation with a user.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

// Provides information about the external source of the slot type's definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   externalSourceSettingProperty := &externalSourceSettingProperty{
//   	grammarSlotTypeSetting: &grammarSlotTypeSettingProperty{
//   		source: &grammarSlotTypeSourceProperty{
//   			s3BucketName: jsii.String("s3BucketName"),
//   			s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   			// the properties below are optional
//   			kmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   }
//
type CfnBot_ExternalSourceSettingProperty struct {
	// Settings required for a slot type based on a grammar that you provide.
	GrammarSlotTypeSetting interface{} `field:"optional" json:"grammarSlotTypeSetting" yaml:"grammarSlotTypeSetting"`
}

// Determines if a Lambda function should be invoked for a specific intent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fulfillmentCodeHookSettingProperty := &fulfillmentCodeHookSettingProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	fulfillmentUpdatesSpecification: &fulfillmentUpdatesSpecificationProperty{
//   		active: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		startResponse: &fulfillmentStartResponseSpecificationProperty{
//   			delayInSeconds: jsii.Number(123),
//   			messageGroups: []interface{}{
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
//   		timeoutInSeconds: jsii.Number(123),
//   		updateResponse: &fulfillmentUpdateResponseSpecificationProperty{
//   			frequencyInSeconds: jsii.Number(123),
//   			messageGroups: []interface{}{
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
//   	postFulfillmentStatusSpecification: &postFulfillmentStatusSpecificationProperty{
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
//   }
//
type CfnBot_FulfillmentCodeHookSettingProperty struct {
	// Indicates whether a Lambda function should be invoked for fulfill a specific intent.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Provides settings for update messages sent to the user for long-running Lambda fulfillment functions.
	//
	// Fulfillment updates can be used only with streaming conversations.
	FulfillmentUpdatesSpecification interface{} `field:"optional" json:"fulfillmentUpdatesSpecification" yaml:"fulfillmentUpdatesSpecification"`
	// Provides settings for messages sent to the user for after the Lambda fulfillment function completes.
	//
	// Post-fulfillment messages can be sent for both streaming and non-streaming conversations.
	PostFulfillmentStatusSpecification interface{} `field:"optional" json:"postFulfillmentStatusSpecification" yaml:"postFulfillmentStatusSpecification"`
}

// Provides settings for a message that is sent to the user when a fulfillment Lambda function starts running.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fulfillmentStartResponseSpecificationProperty := &fulfillmentStartResponseSpecificationProperty{
//   	delayInSeconds: jsii.Number(123),
//   	messageGroups: []interface{}{
//   		&messageGroupProperty{
//   			message: &messageProperty{
//   				customPayload: &customPayloadProperty{
//   					value: jsii.String("value"),
//   				},
//   				imageResponseCard: &imageResponseCardProperty{
//   					title: jsii.String("title"),
//
//   					// the properties below are optional
//   					buttons: []interface{}{
//   						&buttonProperty{
//   							text: jsii.String("text"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					imageUrl: jsii.String("imageUrl"),
//   					subtitle: jsii.String("subtitle"),
//   				},
//   				plainTextMessage: &plainTextMessageProperty{
//   					value: jsii.String("value"),
//   				},
//   				ssmlMessage: &sSMLMessageProperty{
//   					value: jsii.String("value"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			variations: []interface{}{
//   				&messageProperty{
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
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	allowInterrupt: jsii.Boolean(false),
//   }
//
type CfnBot_FulfillmentStartResponseSpecificationProperty struct {
	// The delay between when the Lambda fulfillment function starts running and the start message is played.
	//
	// If the Lambda function returns before the delay is over, the start message isn't played.
	DelayInSeconds *float64 `field:"required" json:"delayInSeconds" yaml:"delayInSeconds"`
	// One to 5 message groups that contain start messages.
	//
	// Amazon Lex chooses one of the messages to play to the user.
	MessageGroups interface{} `field:"required" json:"messageGroups" yaml:"messageGroups"`
	// Determines whether the user can interrupt the start message while it is playing.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
}

// Provides information for updating the user on the progress of fulfilling an intent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fulfillmentUpdateResponseSpecificationProperty := &fulfillmentUpdateResponseSpecificationProperty{
//   	frequencyInSeconds: jsii.Number(123),
//   	messageGroups: []interface{}{
//   		&messageGroupProperty{
//   			message: &messageProperty{
//   				customPayload: &customPayloadProperty{
//   					value: jsii.String("value"),
//   				},
//   				imageResponseCard: &imageResponseCardProperty{
//   					title: jsii.String("title"),
//
//   					// the properties below are optional
//   					buttons: []interface{}{
//   						&buttonProperty{
//   							text: jsii.String("text"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					imageUrl: jsii.String("imageUrl"),
//   					subtitle: jsii.String("subtitle"),
//   				},
//   				plainTextMessage: &plainTextMessageProperty{
//   					value: jsii.String("value"),
//   				},
//   				ssmlMessage: &sSMLMessageProperty{
//   					value: jsii.String("value"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			variations: []interface{}{
//   				&messageProperty{
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
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	allowInterrupt: jsii.Boolean(false),
//   }
//
type CfnBot_FulfillmentUpdateResponseSpecificationProperty struct {
	// The frequency that a message is sent to the user.
	//
	// When the period ends, Amazon Lex chooses a message from the message groups and plays it to the user. If the fulfillment Lambda function returns before the first period ends, an update message is not played to the user.
	FrequencyInSeconds *float64 `field:"required" json:"frequencyInSeconds" yaml:"frequencyInSeconds"`
	// One to 5 message groups that contain update messages.
	//
	// Amazon Lex chooses one of the messages to play to the user.
	MessageGroups interface{} `field:"required" json:"messageGroups" yaml:"messageGroups"`
	// Determines whether the user can interrupt an update message while it is playing.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
}

// Provides information for updating the user on the progress of fulfilling an intent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fulfillmentUpdatesSpecificationProperty := &fulfillmentUpdatesSpecificationProperty{
//   	active: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	startResponse: &fulfillmentStartResponseSpecificationProperty{
//   		delayInSeconds: jsii.Number(123),
//   		messageGroups: []interface{}{
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
//   	timeoutInSeconds: jsii.Number(123),
//   	updateResponse: &fulfillmentUpdateResponseSpecificationProperty{
//   		frequencyInSeconds: jsii.Number(123),
//   		messageGroups: []interface{}{
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
type CfnBot_FulfillmentUpdatesSpecificationProperty struct {
	// Determines whether fulfillment updates are sent to the user. When this field is true, updates are sent.
	//
	// If the active field is set to true, the `startResponse` , `updateResponse` , and `timeoutInSeconds` fields are required.
	Active interface{} `field:"required" json:"active" yaml:"active"`
	// Provides configuration information for the message sent to users when the fulfillment Lambda functions starts running.
	StartResponse interface{} `field:"optional" json:"startResponse" yaml:"startResponse"`
	// The length of time that the fulfillment Lambda function should run before it times out.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// Provides configuration information for messages sent periodically to the user while the fulfillment Lambda function is running.
	UpdateResponse interface{} `field:"optional" json:"updateResponse" yaml:"updateResponse"`
}

// Settings required for a slot type based on a grammar that you provide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grammarSlotTypeSettingProperty := &grammarSlotTypeSettingProperty{
//   	source: &grammarSlotTypeSourceProperty{
//   		s3BucketName: jsii.String("s3BucketName"),
//   		s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   		// the properties below are optional
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
type CfnBot_GrammarSlotTypeSettingProperty struct {
	// The source of the grammar used to create the slot type.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

// Describes the Amazon S3 bucket name and location for the grammar that is the source of the slot type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grammarSlotTypeSourceProperty := &grammarSlotTypeSourceProperty{
//   	s3BucketName: jsii.String("s3BucketName"),
//   	s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   	// the properties below are optional
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnBot_GrammarSlotTypeSourceProperty struct {
	// The name of the S3 bucket that contains the grammar source.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// The path to the grammar in the S3 bucket.
	S3ObjectKey *string `field:"required" json:"s3ObjectKey" yaml:"s3ObjectKey"`
	// The AWS Key Management Service key required to decrypt the contents of the grammar, if any.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

// A card that is shown to the user by a messaging platform.
//
// You define the contents of the card, the card is displayed by the platform.
//
// When you use a response card, the response from the user is constrained to the text associated with a button on the card.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageResponseCardProperty := &imageResponseCardProperty{
//   	title: jsii.String("title"),
//
//   	// the properties below are optional
//   	buttons: []interface{}{
//   		&buttonProperty{
//   			text: jsii.String("text"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	imageUrl: jsii.String("imageUrl"),
//   	subtitle: jsii.String("subtitle"),
//   }
//
type CfnBot_ImageResponseCardProperty struct {
	// The title to display on the response card.
	//
	// The format of the title is determined by the platform displaying the response card.
	Title *string `field:"required" json:"title" yaml:"title"`
	// A list of buttons that should be displayed on the response card.
	//
	// The arrangement of the buttons is determined by the platform that displays the buttons.
	Buttons interface{} `field:"optional" json:"buttons" yaml:"buttons"`
	// The URL of an image to display on the response card.
	//
	// The image URL must be publicly available so that the platform displaying the response card has access to the image.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The subtitle to display on the response card.
	//
	// The format of the subtitle is determined by the platform displaying the response card.
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
}

// The name of a context that must be active for an intent to be selected by Amazon Lex .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputContextProperty := &inputContextProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnBot_InputContextProperty struct {
	// The name of the context.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Provides a statement the Amazon Lex conveys to the user when the intent is successfully fulfilled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intentClosingSettingProperty := &intentClosingSettingProperty{
//   	closingResponse: &responseSpecificationProperty{
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
//
//   	// the properties below are optional
//   	isActive: jsii.Boolean(false),
//   }
//
type CfnBot_IntentClosingSettingProperty struct {
	// The response that Amazon Lex sends to the user when the intent is complete.
	ClosingResponse interface{} `field:"required" json:"closingResponse" yaml:"closingResponse"`
	// Specifies whether an intent's closing response is used.
	//
	// When this field is false, the closing response isn't sent to the user and no closing input from the user is used. If the IsActive field isn't specified, the default is true.
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
}

// Provides a prompt for making sure that the user is ready for the intent to be fulfilled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intentConfirmationSettingProperty := &intentConfirmationSettingProperty{
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
//   	},
//
//   	// the properties below are optional
//   	isActive: jsii.Boolean(false),
//   }
//
type CfnBot_IntentConfirmationSettingProperty struct {
	// When the user answers "no" to the question defined in PromptSpecification, Amazon Lex responds with this response to acknowledge that the intent was canceled.
	DeclinationResponse interface{} `field:"required" json:"declinationResponse" yaml:"declinationResponse"`
	// Prompts the user to confirm the intent.
	//
	// This question should have a yes or no answer.
	PromptSpecification interface{} `field:"required" json:"promptSpecification" yaml:"promptSpecification"`
	// Specifies whether the intent's confirmation is sent to the user.
	//
	// When this field is false, confirmation and declination responses aren't sent and processing continues as if the responses aren't present. If the active field isn't specified, the default is true.
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
}

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

// Provides configuration information for the AMAZON.KendraSearchIntent intent. When you use this intent, Amazon Lex searches the specified Amazon Kendra index and returns documents from the index that match the user's utterance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kendraConfigurationProperty := &kendraConfigurationProperty{
//   	kendraIndex: jsii.String("kendraIndex"),
//
//   	// the properties below are optional
//   	queryFilterString: jsii.String("queryFilterString"),
//   	queryFilterStringEnabled: jsii.Boolean(false),
//   }
//
type CfnBot_KendraConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Kendra index that you want the AMAZON.KendraSearchIntent intent to search. The index must be in the same account and Region as the Amazon Lex bot.
	KendraIndex *string `field:"required" json:"kendraIndex" yaml:"kendraIndex"`
	// A query filter that Amazon Lex sends to Amazon Kendra to filter the response from a query.
	//
	// The filter is in the format defined by Amazon Kendra.
	QueryFilterString *string `field:"optional" json:"queryFilterString" yaml:"queryFilterString"`
	// Determines whether the AMAZON.KendraSearchIntent intent uses a custom query string to query the Amazon Kendra index.
	QueryFilterStringEnabled interface{} `field:"optional" json:"queryFilterStringEnabled" yaml:"queryFilterStringEnabled"`
}

// Specifies a Lambda function that verifies requests to a bot or fulfills the user's request to a bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaCodeHookProperty := &lambdaCodeHookProperty{
//   	codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   	lambdaArn: jsii.String("lambdaArn"),
//   }
//
type CfnBot_LambdaCodeHookProperty struct {
	// Specifies the version of the request-response that you want Amazon Lex to use to invoke your Lambda function.
	CodeHookInterfaceVersion *string `field:"required" json:"codeHookInterfaceVersion" yaml:"codeHookInterfaceVersion"`
	// Specifies the Amazon Resource Name (ARN) of the Lambda function.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
}

// Provides one or more messages that Amazon Lex should send to the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageGroupProperty := &messageGroupProperty{
//   	message: &messageProperty{
//   		customPayload: &customPayloadProperty{
//   			value: jsii.String("value"),
//   		},
//   		imageResponseCard: &imageResponseCardProperty{
//   			title: jsii.String("title"),
//
//   			// the properties below are optional
//   			buttons: []interface{}{
//   				&buttonProperty{
//   					text: jsii.String("text"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			imageUrl: jsii.String("imageUrl"),
//   			subtitle: jsii.String("subtitle"),
//   		},
//   		plainTextMessage: &plainTextMessageProperty{
//   			value: jsii.String("value"),
//   		},
//   		ssmlMessage: &sSMLMessageProperty{
//   			value: jsii.String("value"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	variations: []interface{}{
//   		&messageProperty{
//   			customPayload: &customPayloadProperty{
//   				value: jsii.String("value"),
//   			},
//   			imageResponseCard: &imageResponseCardProperty{
//   				title: jsii.String("title"),
//
//   				// the properties below are optional
//   				buttons: []interface{}{
//   					&buttonProperty{
//   						text: jsii.String("text"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				imageUrl: jsii.String("imageUrl"),
//   				subtitle: jsii.String("subtitle"),
//   			},
//   			plainTextMessage: &plainTextMessageProperty{
//   				value: jsii.String("value"),
//   			},
//   			ssmlMessage: &sSMLMessageProperty{
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnBot_MessageGroupProperty struct {
	// The primary message that Amazon Lex should send to the user.
	Message interface{} `field:"required" json:"message" yaml:"message"`
	// Message variations to send to the user.
	//
	// When variations are defined, Amazon Lex chooses the primary message or one of the variations to send to the user.
	Variations interface{} `field:"optional" json:"variations" yaml:"variations"`
}

// The object that provides message text and it's type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageProperty := &messageProperty{
//   	customPayload: &customPayloadProperty{
//   		value: jsii.String("value"),
//   	},
//   	imageResponseCard: &imageResponseCardProperty{
//   		title: jsii.String("title"),
//
//   		// the properties below are optional
//   		buttons: []interface{}{
//   			&buttonProperty{
//   				text: jsii.String("text"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		imageUrl: jsii.String("imageUrl"),
//   		subtitle: jsii.String("subtitle"),
//   	},
//   	plainTextMessage: &plainTextMessageProperty{
//   		value: jsii.String("value"),
//   	},
//   	ssmlMessage: &sSMLMessageProperty{
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnBot_MessageProperty struct {
	// A message in a custom format defined by the client application.
	CustomPayload interface{} `field:"optional" json:"customPayload" yaml:"customPayload"`
	// A message that defines a response card that the client application can show to the user.
	ImageResponseCard interface{} `field:"optional" json:"imageResponseCard" yaml:"imageResponseCard"`
	// A message in plain text format.
	PlainTextMessage interface{} `field:"optional" json:"plainTextMessage" yaml:"plainTextMessage"`
	// A message in Speech Synthesis Markup Language (SSML) format.
	SsmlMessage interface{} `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

// Indicates whether a slot can return multiple values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multipleValuesSettingProperty := &multipleValuesSettingProperty{
//   	allowMultipleValues: jsii.Boolean(false),
//   }
//
type CfnBot_MultipleValuesSettingProperty struct {
	// Indicates whether a slot can return multiple values.
	//
	// When true, the slot may return more than one value in a response. When false, the slot returns only a single value. If AllowMultipleValues is not set, the default value is false.
	//
	// Multi-value slots are only available in the en-US locale.
	AllowMultipleValues interface{} `field:"optional" json:"allowMultipleValues" yaml:"allowMultipleValues"`
}

// Determines whether Amazon Lex obscures slot values in conversation logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   obfuscationSettingProperty := &obfuscationSettingProperty{
//   	obfuscationSettingType: jsii.String("obfuscationSettingType"),
//   }
//
type CfnBot_ObfuscationSettingProperty struct {
	// Value that determines whether Amazon Lex obscures slot values in conversation logs.
	//
	// The default is to obscure the values.
	ObfuscationSettingType *string `field:"required" json:"obfuscationSettingType" yaml:"obfuscationSettingType"`
}

// Describes a session context that is activated when an intent is fulfilled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputContextProperty := &outputContextProperty{
//   	name: jsii.String("name"),
//   	timeToLiveInSeconds: jsii.Number(123),
//   	turnsToLive: jsii.Number(123),
//   }
//
type CfnBot_OutputContextProperty struct {
	// The name of the output context.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The amount of time, in seconds, that the output context should remain active.
	//
	// The time is figured from the first time the context is sent to the user.
	TimeToLiveInSeconds *float64 `field:"required" json:"timeToLiveInSeconds" yaml:"timeToLiveInSeconds"`
	// The number of conversation turns that the output context should remain active.
	//
	// The number of turns is counted from the first time that the context is sent to the user.
	TurnsToLive *float64 `field:"required" json:"turnsToLive" yaml:"turnsToLive"`
}

// Defines an ASCII text message to send to the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   plainTextMessageProperty := &plainTextMessageProperty{
//   	value: jsii.String("value"),
//   }
//
type CfnBot_PlainTextMessageProperty struct {
	// The message to send to the user.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Provides a setting that determines whether the post-fulfillment response is sent to the user.
//
// For more information, see [Post-fulfillment response](https://docs.aws.amazon.com/lex/latest/dg/streaming-progress.html#progress-complete) in the *Amazon Lex developer guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   postFulfillmentStatusSpecificationProperty := &postFulfillmentStatusSpecificationProperty{
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
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't successful.
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// Specifies a list of message groups that Amazon Lex uses to respond when the fulfillment is successful.
	SuccessResponse interface{} `field:"optional" json:"successResponse" yaml:"successResponse"`
	// Specifies a list of message groups that Amazon Lex uses to respond when fulfillment isn't completed within the timeout period.
	TimeoutResponse interface{} `field:"optional" json:"timeoutResponse" yaml:"timeoutResponse"`
}

// Specifies a list of message groups that Amazon Lex sends to a user to elicit a response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptSpecificationProperty := &promptSpecificationProperty{
//   	maxRetries: jsii.Number(123),
//   	messageGroupsList: []interface{}{
//   		&messageGroupProperty{
//   			message: &messageProperty{
//   				customPayload: &customPayloadProperty{
//   					value: jsii.String("value"),
//   				},
//   				imageResponseCard: &imageResponseCardProperty{
//   					title: jsii.String("title"),
//
//   					// the properties below are optional
//   					buttons: []interface{}{
//   						&buttonProperty{
//   							text: jsii.String("text"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					imageUrl: jsii.String("imageUrl"),
//   					subtitle: jsii.String("subtitle"),
//   				},
//   				plainTextMessage: &plainTextMessageProperty{
//   					value: jsii.String("value"),
//   				},
//   				ssmlMessage: &sSMLMessageProperty{
//   					value: jsii.String("value"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			variations: []interface{}{
//   				&messageProperty{
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
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	allowInterrupt: jsii.Boolean(false),
//   }
//
type CfnBot_PromptSpecificationProperty struct {
	// The maximum number of times the bot tries to elicit a response from the user using this prompt.
	MaxRetries *float64 `field:"required" json:"maxRetries" yaml:"maxRetries"`
	// A collection of responses that Amazon Lex can send to the user.
	//
	// Amazon Lex chooses the actual response to send at runtime.
	MessageGroupsList interface{} `field:"required" json:"messageGroupsList" yaml:"messageGroupsList"`
	// Indicates whether the user can interrupt a speech prompt from the bot.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
}

// Specifies a list of message groups that Amazon Lex uses to respond to user input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseSpecificationProperty := &responseSpecificationProperty{
//   	messageGroupsList: []interface{}{
//   		&messageGroupProperty{
//   			message: &messageProperty{
//   				customPayload: &customPayloadProperty{
//   					value: jsii.String("value"),
//   				},
//   				imageResponseCard: &imageResponseCardProperty{
//   					title: jsii.String("title"),
//
//   					// the properties below are optional
//   					buttons: []interface{}{
//   						&buttonProperty{
//   							text: jsii.String("text"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					imageUrl: jsii.String("imageUrl"),
//   					subtitle: jsii.String("subtitle"),
//   				},
//   				plainTextMessage: &plainTextMessageProperty{
//   					value: jsii.String("value"),
//   				},
//   				ssmlMessage: &sSMLMessageProperty{
//   					value: jsii.String("value"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			variations: []interface{}{
//   				&messageProperty{
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
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	allowInterrupt: jsii.Boolean(false),
//   }
//
type CfnBot_ResponseSpecificationProperty struct {
	// A collection of responses that Amazon Lex can send to the user.
	//
	// Amazon Lex chooses the actual response to send at runtime.
	MessageGroupsList interface{} `field:"required" json:"messageGroupsList" yaml:"messageGroupsList"`
	// Indicates whether the user can interrupt a speech response from Amazon Lex .
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
}

// Specifies an Amazon S3 bucket for logging audio conversations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3BucketLogDestinationProperty := &s3BucketLogDestinationProperty{
//   	logPrefix: jsii.String("logPrefix"),
//   	s3BucketArn: jsii.String("s3BucketArn"),
//
//   	// the properties below are optional
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnBot_S3BucketLogDestinationProperty struct {
	// Specifies the Amazon S3 prefix to assign to audio log files.
	LogPrefix *string `field:"required" json:"logPrefix" yaml:"logPrefix"`
	// Specifies the Amazon Resource Name (ARN) of the Amazon S3 bucket where audio files are stored.
	S3BucketArn *string `field:"required" json:"s3BucketArn" yaml:"s3BucketArn"`
	// Specifies the Amazon Resource Name (ARN) of an AWS Key Management Service key for encrypting audio log files stored in an Amazon S3 bucket.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

// Defines an Amazon S3 bucket location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	s3Bucket: jsii.String("s3Bucket"),
//   	s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   	// the properties below are optional
//   	s3ObjectVersion: jsii.String("s3ObjectVersion"),
//   }
//
type CfnBot_S3LocationProperty struct {
	// The S3 bucket name.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The path and file name to the object in the S3 bucket.
	S3ObjectKey *string `field:"required" json:"s3ObjectKey" yaml:"s3ObjectKey"`
	// The version of the object in the S3 bucket.
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
}

// Defines a Speech Synthesis Markup Language (SSML) prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sSMLMessageProperty := &sSMLMessageProperty{
//   	value: jsii.String("value"),
//   }
//
type CfnBot_SSMLMessageProperty struct {
	// The SSML text that defines the prompt.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// A sample utterance that invokes and intent or responds to a slot elicitation prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sampleUtteranceProperty := &sampleUtteranceProperty{
//   	utterance: jsii.String("utterance"),
//   }
//
type CfnBot_SampleUtteranceProperty struct {
	// The sample utterance that Amazon Lex uses to build its machine-learning model to recognize intents.
	Utterance *string `field:"required" json:"utterance" yaml:"utterance"`
}

// Defines one of the values for a slot type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sampleValueProperty := &sampleValueProperty{
//   	value: jsii.String("value"),
//   }
//
type CfnBot_SampleValueProperty struct {
	// The value that can be used for a slot type.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Specifies the default value to use when a user doesn't provide a value for a slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotDefaultValueProperty := &slotDefaultValueProperty{
//   	defaultValue: jsii.String("defaultValue"),
//   }
//
type CfnBot_SlotDefaultValueProperty struct {
	// The default value to use when a user doesn't provide a value for a slot.
	DefaultValue *string `field:"required" json:"defaultValue" yaml:"defaultValue"`
}

// Defines a list of values that Amazon Lex should use as the default value for a slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotDefaultValueSpecificationProperty := &slotDefaultValueSpecificationProperty{
//   	defaultValueList: []interface{}{
//   		&slotDefaultValueProperty{
//   			defaultValue: jsii.String("defaultValue"),
//   		},
//   	},
//   }
//
type CfnBot_SlotDefaultValueSpecificationProperty struct {
	// A list of default values.
	//
	// Amazon Lex chooses the default value to use in the order that they are presented in the list.
	DefaultValueList interface{} `field:"required" json:"defaultValueList" yaml:"defaultValueList"`
}

// Sets the priority that Amazon Lex should use when eliciting slots values from a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotPriorityProperty := &slotPriorityProperty{
//   	priority: jsii.Number(123),
//   	slotName: jsii.String("slotName"),
//   }
//
type CfnBot_SlotPriorityProperty struct {
	// The priority that Amazon Lex should apply to the slot.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The name of the slot.
	SlotName *string `field:"required" json:"slotName" yaml:"slotName"`
}

// Specifies the definition of a slot.
//
// Amazon Lex elicits slot values from uses to fulfill the user's intent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   		},
//   		sampleUtterances: []interface{}{
//   			&sampleUtteranceProperty{
//   				utterance: jsii.String("utterance"),
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
	// The name of the slot.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the slot type that this slot is based on.
	//
	// The slot type defines the acceptable values for the slot.
	SlotTypeName *string `field:"required" json:"slotTypeName" yaml:"slotTypeName"`
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - OriginalValue - Returns the value entered by the user, if the user value is similar to a slot value.
	// - TopResolution - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the valueSelectionStrategy, the default is OriginalValue.
	ValueElicitationSetting interface{} `field:"required" json:"valueElicitationSetting" yaml:"valueElicitationSetting"`
	// A description of the slot type.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Determines whether the slot can return multiple values to the application.
	MultipleValuesSetting interface{} `field:"optional" json:"multipleValuesSetting" yaml:"multipleValuesSetting"`
	// Determines whether the contents of the slot are obfuscated in Amazon CloudWatch Logs logs.
	//
	// Use obfuscated slots to protect information such as personally identifiable information (PII) in logs.
	ObfuscationSetting interface{} `field:"optional" json:"obfuscationSetting" yaml:"obfuscationSetting"`
}

// Describes a slot type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotTypeProperty := &slotTypeProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	externalSourceSetting: &externalSourceSettingProperty{
//   		grammarSlotTypeSetting: &grammarSlotTypeSettingProperty{
//   			source: &grammarSlotTypeSourceProperty{
//   				s3BucketName: jsii.String("s3BucketName"),
//   				s3ObjectKey: jsii.String("s3ObjectKey"),
//
//   				// the properties below are optional
//   				kmsKeyArn: jsii.String("kmsKeyArn"),
//   			},
//   		},
//   	},
//   	parentSlotTypeSignature: jsii.String("parentSlotTypeSignature"),
//   	slotTypeValues: []interface{}{
//   		&slotTypeValueProperty{
//   			sampleValue: &sampleValueProperty{
//   				value: jsii.String("value"),
//   			},
//
//   			// the properties below are optional
//   			synonyms: []interface{}{
//   				&sampleValueProperty{
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	valueSelectionSetting: &slotValueSelectionSettingProperty{
//   		resolutionStrategy: jsii.String("resolutionStrategy"),
//
//   		// the properties below are optional
//   		advancedRecognitionSetting: &advancedRecognitionSettingProperty{
//   			audioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   		},
//   		regexFilter: &slotValueRegexFilterProperty{
//   			pattern: jsii.String("pattern"),
//   		},
//   	},
//   }
//
type CfnBot_SlotTypeProperty struct {
	// The name of the slot type.
	//
	// A slot type name must be unique withing the account.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the slot type.
	//
	// Use the description to help identify the slot type in lists.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Sets the type of external information used to create the slot type.
	ExternalSourceSetting interface{} `field:"optional" json:"externalSourceSetting" yaml:"externalSourceSetting"`
	// The built-in slot type used as a parent of this slot type.
	//
	// When you define a parent slot type, the new slot type has the configuration of the parent lot type.
	//
	// Only AMAZON.AlphaNumeric is supported.
	ParentSlotTypeSignature *string `field:"optional" json:"parentSlotTypeSignature" yaml:"parentSlotTypeSignature"`
	// A list of SlotTypeValue objects that defines the values that the slot type can take.
	//
	// Each value can have a list of synonyms, additional values that help train the machine learning model about the values that it resolves for the slot.
	SlotTypeValues interface{} `field:"optional" json:"slotTypeValues" yaml:"slotTypeValues"`
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - OriginalValue - Returns the value entered by the user, if the user value is similar to a slot value.
	// - TopResolution - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the valueSelectionStrategy, the default is OriginalValue.
	ValueSelectionSetting interface{} `field:"optional" json:"valueSelectionSetting" yaml:"valueSelectionSetting"`
}

// Each slot type can have a set of values.
//
// The `SlotTypeValue` represents a value that the slot type can take.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotTypeValueProperty := &slotTypeValueProperty{
//   	sampleValue: &sampleValueProperty{
//   		value: jsii.String("value"),
//   	},
//
//   	// the properties below are optional
//   	synonyms: []interface{}{
//   		&sampleValueProperty{
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBot_SlotTypeValueProperty struct {
	// The value of the slot type entry.
	SampleValue interface{} `field:"required" json:"sampleValue" yaml:"sampleValue"`
	// Additional values related to the slot type entry.
	Synonyms interface{} `field:"optional" json:"synonyms" yaml:"synonyms"`
}

// Settings that you can use for eliciting a slot value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   	},
//   	sampleUtterances: []interface{}{
//   		&sampleUtteranceProperty{
//   			utterance: jsii.String("utterance"),
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
	// This is optional. In most cases Amazon Lex is capable of understanding user utterances.
	SampleUtterances interface{} `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// Specifies the prompts that Amazon Lex uses while a bot is waiting for customer input.
	WaitAndContinueSpecification interface{} `field:"optional" json:"waitAndContinueSpecification" yaml:"waitAndContinueSpecification"`
}

// Provides a regular expression used to validate the value of a slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotValueRegexFilterProperty := &slotValueRegexFilterProperty{
//   	pattern: jsii.String("pattern"),
//   }
//
type CfnBot_SlotValueRegexFilterProperty struct {
	// A regular expression used to validate the value of a slot.
	//
	// Use a standard regular expression. Amazon Lex supports the following characters in the regular expression:
	//
	// - A-Z, a-z
	// - 0-9
	// - Unicode characters ("\ u<Unicode>")
	//
	// Represent Unicode characters with four digits, for example "]u0041" or "\ u005A".
	//
	// The following regular expression operators are not supported:
	//
	// - Infinite repeaters: *, +, or {x,} with no upper bound
	// - Wild card (.)
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
}

// Contains settings used by Amazon Lex to select a slot value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotValueSelectionSettingProperty := &slotValueSelectionSettingProperty{
//   	resolutionStrategy: jsii.String("resolutionStrategy"),
//
//   	// the properties below are optional
//   	advancedRecognitionSetting: &advancedRecognitionSettingProperty{
//   		audioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   	},
//   	regexFilter: &slotValueRegexFilterProperty{
//   		pattern: jsii.String("pattern"),
//   	},
//   }
//
type CfnBot_SlotValueSelectionSettingProperty struct {
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - OriginalValue - Returns the value entered by the user, if the user value is similar to a slot value.
	// - TopResolution - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the valueSelectionStrategy, the default is OriginalValue.
	ResolutionStrategy *string `field:"required" json:"resolutionStrategy" yaml:"resolutionStrategy"`
	// Specifies settings that enable advanced recognition settings for slot values.
	//
	// You can use this to enable using slot values as a custom vocabulary for recognizing user utterances.
	AdvancedRecognitionSetting interface{} `field:"optional" json:"advancedRecognitionSetting" yaml:"advancedRecognitionSetting"`
	// A regular expression used to validate the value of a slot.
	RegexFilter interface{} `field:"optional" json:"regexFilter" yaml:"regexFilter"`
}

// Defines the messages that Amazon Lex sends to a user to remind them that the bot is waiting for a response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stillWaitingResponseSpecificationProperty := &stillWaitingResponseSpecificationProperty{
//   	frequencyInSeconds: jsii.Number(123),
//   	messageGroupsList: []interface{}{
//   		&messageGroupProperty{
//   			message: &messageProperty{
//   				customPayload: &customPayloadProperty{
//   					value: jsii.String("value"),
//   				},
//   				imageResponseCard: &imageResponseCardProperty{
//   					title: jsii.String("title"),
//
//   					// the properties below are optional
//   					buttons: []interface{}{
//   						&buttonProperty{
//   							text: jsii.String("text"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					imageUrl: jsii.String("imageUrl"),
//   					subtitle: jsii.String("subtitle"),
//   				},
//   				plainTextMessage: &plainTextMessageProperty{
//   					value: jsii.String("value"),
//   				},
//   				ssmlMessage: &sSMLMessageProperty{
//   					value: jsii.String("value"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			variations: []interface{}{
//   				&messageProperty{
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
//   			},
//   		},
//   	},
//   	timeoutInSeconds: jsii.Number(123),
//
//   	// the properties below are optional
//   	allowInterrupt: jsii.Boolean(false),
//   }
//
type CfnBot_StillWaitingResponseSpecificationProperty struct {
	// How often a message should be sent to the user.
	//
	// Minimum of 1 second, maximum of 5 minutes.
	FrequencyInSeconds *float64 `field:"required" json:"frequencyInSeconds" yaml:"frequencyInSeconds"`
	// A collection of responses that Amazon Lex can send to the user.
	//
	// Amazon Lex chooses the actual response to send at runtime.
	MessageGroupsList interface{} `field:"required" json:"messageGroupsList" yaml:"messageGroupsList"`
	// If Amazon Lex waits longer than this length of time for a response, it will stop sending messages.
	TimeoutInSeconds *float64 `field:"required" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// Indicates that the user can interrupt the response by speaking while the message is being played.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
}

// Specifies configuration settings for the alias used to test the bot.
//
// If the `TestBotAliasSettings` property is not specified, the settings are configured with default values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sentimentAnalysisSettings interface{}
//
//   testBotAliasSettingsProperty := &testBotAliasSettingsProperty{
//   	botAliasLocaleSettings: []interface{}{
//   		&botAliasLocaleSettingsItemProperty{
//   			botAliasLocaleSetting: &botAliasLocaleSettingsProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				codeHookSpecification: &codeHookSpecificationProperty{
//   					lambdaCodeHook: &lambdaCodeHookProperty{
//   						codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   						lambdaArn: jsii.String("lambdaArn"),
//   					},
//   				},
//   			},
//   			localeId: jsii.String("localeId"),
//   		},
//   	},
//   	conversationLogSettings: &conversationLogSettingsProperty{
//   		audioLogSettings: []interface{}{
//   			&audioLogSettingProperty{
//   				destination: &audioLogDestinationProperty{
//   					s3Bucket: &s3BucketLogDestinationProperty{
//   						logPrefix: jsii.String("logPrefix"),
//   						s3BucketArn: jsii.String("s3BucketArn"),
//
//   						// the properties below are optional
//   						kmsKeyArn: jsii.String("kmsKeyArn"),
//   					},
//   				},
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   		textLogSettings: []interface{}{
//   			&textLogSettingProperty{
//   				destination: &textLogDestinationProperty{
//   					cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   						cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   						logPrefix: jsii.String("logPrefix"),
//   					},
//   				},
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	sentimentAnalysisSettings: sentimentAnalysisSettings,
//   }
//
type CfnBot_TestBotAliasSettingsProperty struct {
	// Specifies settings that are unique to a locale.
	//
	// For example, you can use a different Lambda function depending on the bot's locale.
	BotAliasLocaleSettings interface{} `field:"optional" json:"botAliasLocaleSettings" yaml:"botAliasLocaleSettings"`
	// Specifies settings for conversation logs that save audio, text, and metadata information for conversations with your users.
	ConversationLogSettings interface{} `field:"optional" json:"conversationLogSettings" yaml:"conversationLogSettings"`
	// Specifies a description for the test bot alias.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
	SentimentAnalysisSettings interface{} `field:"optional" json:"sentimentAnalysisSettings" yaml:"sentimentAnalysisSettings"`
}

// Specifies the Amazon CloudWatch Logs destination log group for conversation text logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textLogDestinationProperty := &textLogDestinationProperty{
//   	cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   		cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		logPrefix: jsii.String("logPrefix"),
//   	},
//   }
//
type CfnBot_TextLogDestinationProperty struct {
	// Specifies the Amazon CloudWatch Logs log group where text and metadata logs are delivered.
	CloudWatch interface{} `field:"required" json:"cloudWatch" yaml:"cloudWatch"`
}

// Specifies settings to enable conversation logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textLogSettingProperty := &textLogSettingProperty{
//   	destination: &textLogDestinationProperty{
//   		cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   			cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			logPrefix: jsii.String("logPrefix"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnBot_TextLogSettingProperty struct {
	// Specifies the Amazon CloudWatch Logs destination log group for conversation text logs.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether conversation logs should be stored for an alias.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

// Identifies the Amazon Polly voice used for audio interaction with the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   voiceSettingsProperty := &voiceSettingsProperty{
//   	voiceId: jsii.String("voiceId"),
//   }
//
type CfnBot_VoiceSettingsProperty struct {
	// The Amazon Polly voice used for voice interaction with the user.
	VoiceId *string `field:"required" json:"voiceId" yaml:"voiceId"`
}

// Specifies the prompts that Amazon Lex uses while a bot is waiting for customer input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   waitAndContinueSpecificationProperty := &waitAndContinueSpecificationProperty{
//   	continueResponse: &responseSpecificationProperty{
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
//   	waitingResponse: &responseSpecificationProperty{
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
//
//   	// the properties below are optional
//   	isActive: jsii.Boolean(false),
//   	stillWaitingResponse: &stillWaitingResponseSpecificationProperty{
//   		frequencyInSeconds: jsii.Number(123),
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
//   		timeoutInSeconds: jsii.Number(123),
//
//   		// the properties below are optional
//   		allowInterrupt: jsii.Boolean(false),
//   	},
//   }
//
type CfnBot_WaitAndContinueSpecificationProperty struct {
	// The response that Amazon Lex sends to indicate that the bot is ready to continue the conversation.
	ContinueResponse interface{} `field:"required" json:"continueResponse" yaml:"continueResponse"`
	// The response that Amazon Lex sends to indicate that the bot is waiting for the conversation to continue.
	WaitingResponse interface{} `field:"required" json:"waitingResponse" yaml:"waitingResponse"`
	// Specifies whether the bot will wait for a user to respond.
	//
	// When this field is false, wait and continue responses for a slot aren't used and the bot expects an appropriate response within the configured timeout. If the IsActive field isn't specified, the default is true.
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// A response that Amazon Lex sends periodically to the user to indicate that the bot is still waiting for input from the user.
	StillWaitingResponse interface{} `field:"optional" json:"stillWaitingResponse" yaml:"stillWaitingResponse"`
}

// A CloudFormation `AWS::Lex::BotAlias`.
//
// Specifies an alias for the specified version of a bot. Use an alias to enable you to change the version of a bot without updating applications that use the bot.
//
// For example, you can specify an alias called "PROD" that your applications use to call the Amazon Lex bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sentimentAnalysisSettings interface{}
//
//   cfnBotAlias := awscdk.Aws_lex.NewCfnBotAlias(this, jsii.String("MyCfnBotAlias"), &cfnBotAliasProps{
//   	botAliasName: jsii.String("botAliasName"),
//   	botId: jsii.String("botId"),
//
//   	// the properties below are optional
//   	botAliasLocaleSettings: []interface{}{
//   		&botAliasLocaleSettingsItemProperty{
//   			botAliasLocaleSetting: &botAliasLocaleSettingsProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				codeHookSpecification: &codeHookSpecificationProperty{
//   					lambdaCodeHook: &lambdaCodeHookProperty{
//   						codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   						lambdaArn: jsii.String("lambdaArn"),
//   					},
//   				},
//   			},
//   			localeId: jsii.String("localeId"),
//   		},
//   	},
//   	botAliasTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	botVersion: jsii.String("botVersion"),
//   	conversationLogSettings: &conversationLogSettingsProperty{
//   		audioLogSettings: []interface{}{
//   			&audioLogSettingProperty{
//   				destination: &audioLogDestinationProperty{
//   					s3Bucket: &s3BucketLogDestinationProperty{
//   						logPrefix: jsii.String("logPrefix"),
//   						s3BucketArn: jsii.String("s3BucketArn"),
//
//   						// the properties below are optional
//   						kmsKeyArn: jsii.String("kmsKeyArn"),
//   					},
//   				},
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   		textLogSettings: []interface{}{
//   			&textLogSettingProperty{
//   				destination: &textLogDestinationProperty{
//   					cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   						cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   						logPrefix: jsii.String("logPrefix"),
//   					},
//   				},
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	sentimentAnalysisSettings: sentimentAnalysisSettings,
//   })
//
type CfnBotAlias interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the bot alias.
	AttrArn() *string
	// The unique identifier of the bot alias.
	AttrBotAliasId() *string
	// The current status of the bot alias.
	//
	// When the status is Available the alias is ready for use with your bot.
	AttrBotAliasStatus() *string
	// Maps configuration information to a specific locale.
	//
	// You can use this parameter to specify a specific Lambda function to run different functions in different locales.
	BotAliasLocaleSettings() interface{}
	SetBotAliasLocaleSettings(val interface{})
	// The name of the bot alias.
	BotAliasName() *string
	SetBotAliasName(val *string)
	// An array of key-value pairs to apply to this resource.
	//
	// You can only add tags when you specify an alias.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	BotAliasTags() interface{}
	SetBotAliasTags(val interface{})
	// The unique identifier of the bot.
	BotId() *string
	SetBotId(val *string)
	// The version of the bot that the bot alias references.
	BotVersion() *string
	SetBotVersion(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Specifies whether Amazon Lex logs text and audio for conversations with the bot.
	//
	// When you enable conversation logs, text logs store text input, transcripts of audio input, and associated metadata in Amazon CloudWatch logs. Audio logs store input in Amazon S3 .
	ConversationLogSettings() interface{}
	SetConversationLogSettings(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the bot alias.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
	SentimentAnalysisSettings() interface{}
	SetSentimentAnalysisSettings(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBotAlias
type jsiiProxy_CfnBotAlias struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBotAlias) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) AttrBotAliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBotAliasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) AttrBotAliasStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBotAliasStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) BotAliasLocaleSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botAliasLocaleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) BotAliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botAliasName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) BotAliasTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botAliasTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) BotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) BotVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) ConversationLogSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conversationLogSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) SentimentAnalysisSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sentimentAnalysisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotAlias) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lex::BotAlias`.
func NewCfnBotAlias(scope constructs.Construct, id *string, props *CfnBotAliasProps) CfnBotAlias {
	_init_.Initialize()

	j := jsiiProxy_CfnBotAlias{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::BotAlias`.
func NewCfnBotAlias_Override(c CfnBotAlias, scope constructs.Construct, id *string, props *CfnBotAliasProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetBotAliasLocaleSettings(val interface{}) {
	_jsii_.Set(
		j,
		"botAliasLocaleSettings",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetBotAliasName(val *string) {
	_jsii_.Set(
		j,
		"botAliasName",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetBotAliasTags(val interface{}) {
	_jsii_.Set(
		j,
		"botAliasTags",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetBotId(val *string) {
	_jsii_.Set(
		j,
		"botId",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetBotVersion(val *string) {
	_jsii_.Set(
		j,
		"botVersion",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetConversationLogSettings(val interface{}) {
	_jsii_.Set(
		j,
		"conversationLogSettings",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnBotAlias) SetSentimentAnalysisSettings(val interface{}) {
	_jsii_.Set(
		j,
		"sentimentAnalysisSettings",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnBotAlias_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBotAlias_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnBotAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBotAlias_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBotAlias) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBotAlias) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBotAlias) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBotAlias) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBotAlias) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBotAlias) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBotAlias) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBotAlias) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBotAlias) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBotAlias) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBotAlias) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBotAlias) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBotAlias) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBotAlias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBotAlias) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the S3 bucket location where audio logs are stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioLogDestinationProperty := &audioLogDestinationProperty{
//   	s3Bucket: &s3BucketLogDestinationProperty{
//   		logPrefix: jsii.String("logPrefix"),
//   		s3BucketArn: jsii.String("s3BucketArn"),
//
//   		// the properties below are optional
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
type CfnBotAlias_AudioLogDestinationProperty struct {
	// The S3 bucket location where audio logs are stored.
	S3Bucket interface{} `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
}

// Settings for logging audio of conversations between Amazon Lex and a user.
//
// You specify whether to log audio and the Amazon S3 bucket where the audio file is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioLogSettingProperty := &audioLogSettingProperty{
//   	destination: &audioLogDestinationProperty{
//   		s3Bucket: &s3BucketLogDestinationProperty{
//   			logPrefix: jsii.String("logPrefix"),
//   			s3BucketArn: jsii.String("s3BucketArn"),
//
//   			// the properties below are optional
//   			kmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnBotAlias_AudioLogSettingProperty struct {
	// The location of audio log files collected when conversation logging is enabled for a bot.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Determines whether audio logging in enabled for the bot.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

// Specifies settings that are unique to a locale.
//
// For example, you can use different Lambda function depending on the bot's locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botAliasLocaleSettingsItemProperty := &botAliasLocaleSettingsItemProperty{
//   	botAliasLocaleSetting: &botAliasLocaleSettingsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		codeHookSpecification: &codeHookSpecificationProperty{
//   			lambdaCodeHook: &lambdaCodeHookProperty{
//   				codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   				lambdaArn: jsii.String("lambdaArn"),
//   			},
//   		},
//   	},
//   	localeId: jsii.String("localeId"),
//   }
//
type CfnBotAlias_BotAliasLocaleSettingsItemProperty struct {
	// Specifies settings that are unique to a locale.
	BotAliasLocaleSetting interface{} `field:"required" json:"botAliasLocaleSetting" yaml:"botAliasLocaleSetting"`
	// The unique identifier of the locale.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}

// Specifies settings that are unique to a locale.
//
// For example, you can use different Lambda function depending on the bot's locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botAliasLocaleSettingsProperty := &botAliasLocaleSettingsProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	codeHookSpecification: &codeHookSpecificationProperty{
//   		lambdaCodeHook: &lambdaCodeHookProperty{
//   			codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   			lambdaArn: jsii.String("lambdaArn"),
//   		},
//   	},
//   }
//
type CfnBotAlias_BotAliasLocaleSettingsProperty struct {
	// Determines whether the locale is enabled for the bot.
	//
	// If the value is false, the locale isn't available for use.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the Lambda function that should be used in the locale.
	CodeHookSpecification interface{} `field:"optional" json:"codeHookSpecification" yaml:"codeHookSpecification"`
}

// The Amazon CloudWatch Logs log group where the text and metadata logs are delivered.
//
// The log group must exist before you enable logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogGroupLogDestinationProperty := &cloudWatchLogGroupLogDestinationProperty{
//   	cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   	logPrefix: jsii.String("logPrefix"),
//   }
//
type CfnBotAlias_CloudWatchLogGroupLogDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the log group where text and metadata logs are delivered.
	CloudWatchLogGroupArn *string `field:"required" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// The prefix of the log stream name within the log group that you specified.
	LogPrefix *string `field:"required" json:"logPrefix" yaml:"logPrefix"`
}

// Contains information about code hooks that Amazon Lex calls during a conversation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeHookSpecificationProperty := &codeHookSpecificationProperty{
//   	lambdaCodeHook: &lambdaCodeHookProperty{
//   		codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   		lambdaArn: jsii.String("lambdaArn"),
//   	},
//   }
//
type CfnBotAlias_CodeHookSpecificationProperty struct {
	// Specifies a Lambda function that verifies requests to a bot or fulfills the user's request to a bot.
	LambdaCodeHook interface{} `field:"required" json:"lambdaCodeHook" yaml:"lambdaCodeHook"`
}

// Configures conversation logging that saves audio, text, and metadata for the conversations with your users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conversationLogSettingsProperty := &conversationLogSettingsProperty{
//   	audioLogSettings: []interface{}{
//   		&audioLogSettingProperty{
//   			destination: &audioLogDestinationProperty{
//   				s3Bucket: &s3BucketLogDestinationProperty{
//   					logPrefix: jsii.String("logPrefix"),
//   					s3BucketArn: jsii.String("s3BucketArn"),
//
//   					// the properties below are optional
//   					kmsKeyArn: jsii.String("kmsKeyArn"),
//   				},
//   			},
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	textLogSettings: []interface{}{
//   		&textLogSettingProperty{
//   			destination: &textLogDestinationProperty{
//   				cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   					cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   					logPrefix: jsii.String("logPrefix"),
//   				},
//   			},
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnBotAlias_ConversationLogSettingsProperty struct {
	// The Amazon S3 settings for logging audio to an S3 bucket.
	AudioLogSettings interface{} `field:"optional" json:"audioLogSettings" yaml:"audioLogSettings"`
	// The Amazon CloudWatch Logs settings for logging text and metadata.
	TextLogSettings interface{} `field:"optional" json:"textLogSettings" yaml:"textLogSettings"`
}

// Specifies a Lambda function that verifies requests to a bot or fulfills the user's request to a bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaCodeHookProperty := &lambdaCodeHookProperty{
//   	codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   	lambdaArn: jsii.String("lambdaArn"),
//   }
//
type CfnBotAlias_LambdaCodeHookProperty struct {
	// The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.
	CodeHookInterfaceVersion *string `field:"required" json:"codeHookInterfaceVersion" yaml:"codeHookInterfaceVersion"`
	// The Amazon Resource Name (ARN) of the Lambda function.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
}

// Specifies an Amazon S3 bucket for logging audio conversations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3BucketLogDestinationProperty := &s3BucketLogDestinationProperty{
//   	logPrefix: jsii.String("logPrefix"),
//   	s3BucketArn: jsii.String("s3BucketArn"),
//
//   	// the properties below are optional
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnBotAlias_S3BucketLogDestinationProperty struct {
	// The S3 prefix to assign to audio log files.
	LogPrefix *string `field:"required" json:"logPrefix" yaml:"logPrefix"`
	// The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.
	S3BucketArn *string `field:"required" json:"s3BucketArn" yaml:"s3BucketArn"`
	// The Amazon Resource Name (ARN) of an AWS Key Management Service key for encrypting audio log files stored in an S3 bucket.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textLogDestinationProperty := &textLogDestinationProperty{
//   	cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   		cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		logPrefix: jsii.String("logPrefix"),
//   	},
//   }
//
type CfnBotAlias_TextLogDestinationProperty struct {
	// Defines the Amazon CloudWatch Logs log group where text and metadata logs are delivered.
	CloudWatch interface{} `field:"required" json:"cloudWatch" yaml:"cloudWatch"`
}

// Defines settings to enable conversation logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textLogSettingProperty := &textLogSettingProperty{
//   	destination: &textLogDestinationProperty{
//   		cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   			cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			logPrefix: jsii.String("logPrefix"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnBotAlias_TextLogSettingProperty struct {
	// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Determines whether conversation logs should be stored for an alias.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

// Properties for defining a `CfnBotAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sentimentAnalysisSettings interface{}
//
//   cfnBotAliasProps := &cfnBotAliasProps{
//   	botAliasName: jsii.String("botAliasName"),
//   	botId: jsii.String("botId"),
//
//   	// the properties below are optional
//   	botAliasLocaleSettings: []interface{}{
//   		&botAliasLocaleSettingsItemProperty{
//   			botAliasLocaleSetting: &botAliasLocaleSettingsProperty{
//   				enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				codeHookSpecification: &codeHookSpecificationProperty{
//   					lambdaCodeHook: &lambdaCodeHookProperty{
//   						codeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   						lambdaArn: jsii.String("lambdaArn"),
//   					},
//   				},
//   			},
//   			localeId: jsii.String("localeId"),
//   		},
//   	},
//   	botAliasTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	botVersion: jsii.String("botVersion"),
//   	conversationLogSettings: &conversationLogSettingsProperty{
//   		audioLogSettings: []interface{}{
//   			&audioLogSettingProperty{
//   				destination: &audioLogDestinationProperty{
//   					s3Bucket: &s3BucketLogDestinationProperty{
//   						logPrefix: jsii.String("logPrefix"),
//   						s3BucketArn: jsii.String("s3BucketArn"),
//
//   						// the properties below are optional
//   						kmsKeyArn: jsii.String("kmsKeyArn"),
//   					},
//   				},
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   		textLogSettings: []interface{}{
//   			&textLogSettingProperty{
//   				destination: &textLogDestinationProperty{
//   					cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   						cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   						logPrefix: jsii.String("logPrefix"),
//   					},
//   				},
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	sentimentAnalysisSettings: sentimentAnalysisSettings,
//   }
//
type CfnBotAliasProps struct {
	// The name of the bot alias.
	BotAliasName *string `field:"required" json:"botAliasName" yaml:"botAliasName"`
	// The unique identifier of the bot.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Maps configuration information to a specific locale.
	//
	// You can use this parameter to specify a specific Lambda function to run different functions in different locales.
	BotAliasLocaleSettings interface{} `field:"optional" json:"botAliasLocaleSettings" yaml:"botAliasLocaleSettings"`
	// An array of key-value pairs to apply to this resource.
	//
	// You can only add tags when you specify an alias.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	BotAliasTags interface{} `field:"optional" json:"botAliasTags" yaml:"botAliasTags"`
	// The version of the bot that the bot alias references.
	BotVersion *string `field:"optional" json:"botVersion" yaml:"botVersion"`
	// Specifies whether Amazon Lex logs text and audio for conversations with the bot.
	//
	// When you enable conversation logs, text logs store text input, transcripts of audio input, and associated metadata in Amazon CloudWatch logs. Audio logs store input in Amazon S3 .
	ConversationLogSettings interface{} `field:"optional" json:"conversationLogSettings" yaml:"conversationLogSettings"`
	// The description of the bot alias.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
	SentimentAnalysisSettings interface{} `field:"optional" json:"sentimentAnalysisSettings" yaml:"sentimentAnalysisSettings"`
}

// Properties for defining a `CfnBot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataPrivacy interface{}
//   var sentimentAnalysisSettings interface{}
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
//   						postFulfillmentStatusSpecification: &postFulfillmentStatusSpecificationProperty{
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
//
//   						// the properties below are optional
//   						isActive: jsii.Boolean(false),
//   					},
//   					intentConfirmationSetting: &intentConfirmationSettingProperty{
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
//   						},
//
//   						// the properties below are optional
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
//   								},
//   								sampleUtterances: []interface{}{
//   									&sampleUtteranceProperty{
//   										utterance: jsii.String("utterance"),
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
	// Provides information on additional privacy protections Amazon Lex should use with the bot's data.
	DataPrivacy interface{} `field:"required" json:"dataPrivacy" yaml:"dataPrivacy"`
	// The time, in seconds, that Amazon Lex should keep information about a user's conversation with the bot.
	//
	// A user interaction remains active for the amount of time specified. If no conversation occurs during this time, the session expires and Amazon Lex deletes any data provided before the timeout.
	//
	// You can specify between 60 (1 minute) and 86,400 (24 hours) seconds.
	IdleSessionTtlInSeconds *float64 `field:"required" json:"idleSessionTtlInSeconds" yaml:"idleSessionTtlInSeconds"`
	// The name of the field to filter the list of bots.
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

// A CloudFormation `AWS::Lex::BotVersion`.
//
// Specifies a new version of the bot based on the `DRAFT` version. If the `DRAFT` version of this resource hasn't changed since you created the last version, Amazon Lex doesn't create a new version, it returns the last created version.
//
// When you specify the first version of a bot, Amazon Lex sets the version to 1. Subsequent versions increment by 1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBotVersion := awscdk.Aws_lex.NewCfnBotVersion(this, jsii.String("MyCfnBotVersion"), &cfnBotVersionProps{
//   	botId: jsii.String("botId"),
//   	botVersionLocaleSpecification: []interface{}{
//   		&botVersionLocaleSpecificationProperty{
//   			botVersionLocaleDetails: &botVersionLocaleDetailsProperty{
//   				sourceBotVersion: jsii.String("sourceBotVersion"),
//   			},
//   			localeId: jsii.String("localeId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   })
//
type CfnBotVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The version of the bot.
	AttrBotVersion() *string
	// The unique identifier of the bot.
	BotId() *string
	SetBotId(val *string)
	// Specifies the locales that Amazon Lex adds to this version.
	//
	// You can choose the Draft version or any other previously published version for each locale. When you specify a source version, the locale data is copied from the source version to the new version.
	BotVersionLocaleSpecification() interface{}
	SetBotVersionLocaleSpecification(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the version.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBotVersion
type jsiiProxy_CfnBotVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBotVersion) AttrBotVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBotVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) BotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) BotVersionLocaleSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"botVersionLocaleSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBotVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lex::BotVersion`.
func NewCfnBotVersion(scope constructs.Construct, id *string, props *CfnBotVersionProps) CfnBotVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnBotVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::BotVersion`.
func NewCfnBotVersion_Override(c CfnBotVersion, scope constructs.Construct, id *string, props *CfnBotVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBotVersion) SetBotId(val *string) {
	_jsii_.Set(
		j,
		"botId",
		val,
	)
}

func (j *jsiiProxy_CfnBotVersion) SetBotVersionLocaleSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"botVersionLocaleSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnBotVersion) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnBotVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBotVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnBotVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBotVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBotVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBotVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBotVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBotVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBotVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBotVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBotVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBotVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBotVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBotVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBotVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBotVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBotVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBotVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBotVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The version of a bot used for a bot locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botVersionLocaleDetailsProperty := &botVersionLocaleDetailsProperty{
//   	sourceBotVersion: jsii.String("sourceBotVersion"),
//   }
//
type CfnBotVersion_BotVersionLocaleDetailsProperty struct {
	// The version of a bot used for a bot locale.
	SourceBotVersion *string `field:"required" json:"sourceBotVersion" yaml:"sourceBotVersion"`
}

// Specifies the locale that Amazon Lex adds to this version.
//
// You can choose the Draft version or any other previously published version for each locale. When you specify a source version, the locale data is copied from the source version to the new version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botVersionLocaleSpecificationProperty := &botVersionLocaleSpecificationProperty{
//   	botVersionLocaleDetails: &botVersionLocaleDetailsProperty{
//   		sourceBotVersion: jsii.String("sourceBotVersion"),
//   	},
//   	localeId: jsii.String("localeId"),
//   }
//
type CfnBotVersion_BotVersionLocaleSpecificationProperty struct {
	// The version of a bot used for a bot locale.
	BotVersionLocaleDetails interface{} `field:"required" json:"botVersionLocaleDetails" yaml:"botVersionLocaleDetails"`
	// The identifier of the locale to add to the version.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}

// Properties for defining a `CfnBotVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBotVersionProps := &cfnBotVersionProps{
//   	botId: jsii.String("botId"),
//   	botVersionLocaleSpecification: []interface{}{
//   		&botVersionLocaleSpecificationProperty{
//   			botVersionLocaleDetails: &botVersionLocaleDetailsProperty{
//   				sourceBotVersion: jsii.String("sourceBotVersion"),
//   			},
//   			localeId: jsii.String("localeId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnBotVersionProps struct {
	// The unique identifier of the bot.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Specifies the locales that Amazon Lex adds to this version.
	//
	// You can choose the Draft version or any other previously published version for each locale. When you specify a source version, the locale data is copied from the source version to the new version.
	BotVersionLocaleSpecification interface{} `field:"required" json:"botVersionLocaleSpecification" yaml:"botVersionLocaleSpecification"`
	// The description of the version.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

// A CloudFormation `AWS::Lex::ResourcePolicy`.
//
// Specifies a new resource policy with the specified policy statements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnResourcePolicy := awscdk.Aws_lex.NewCfnResourcePolicy(this, jsii.String("MyCfnResourcePolicy"), &cfnResourcePolicyProps{
//   	policy: policy,
//   	resourceArn: jsii.String("resourceArn"),
//   })
//
type CfnResourcePolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The identifier of the resource policy.
	AttrId() *string
	// Specifies the current revision of a resource policy.
	AttrRevisionId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// A resource policy to add to the resource.
	//
	// The policy is a JSON structure that contains one or more statements that define the policy. The policy must follow IAM syntax. If the policy isn't valid, Amazon Lex returns a validation exception.
	Policy() interface{}
	SetPolicy(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.
	ResourceArn() *string
	SetResourceArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnResourcePolicy
type jsiiProxy_CfnResourcePolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResourcePolicy) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) AttrRevisionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRevisionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Policy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lex::ResourcePolicy`.
func NewCfnResourcePolicy(scope constructs.Construct, id *string, props *CfnResourcePolicyProps) CfnResourcePolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnResourcePolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lex::ResourcePolicy`.
func NewCfnResourcePolicy_Override(c CfnResourcePolicy, scope constructs.Construct, id *string, props *CfnResourcePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResourcePolicy) SetPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_CfnResourcePolicy) SetResourceArn(val *string) {
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResourcePolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResourcePolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnResourcePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourcePolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResourcePolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResourcePolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResourcePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnResourcePolicyProps := &cfnResourcePolicyProps{
//   	policy: policy,
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnResourcePolicyProps struct {
	// A resource policy to add to the resource.
	//
	// The policy is a JSON structure that contains one or more statements that define the policy. The policy must follow IAM syntax. If the policy isn't valid, Amazon Lex returns a validation exception.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

