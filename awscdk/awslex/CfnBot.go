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
// > Amazon Lex V2 is the only supported version in AWS CloudFormation .
//
// Specifies an Amazon Lex conversational bot.
//
// You must configure an intent based on the `AMAZON.FallbackIntent` built-in intent. If you don't add one, creating the bot will fail.
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
//   cfnBot := awscdk.Aws_lex.NewCfnBot(this, jsii.String("MyCfnBot"), &CfnBotProps{
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
	// By default, data stored by Amazon Lex is encrypted.
	//
	// The `DataPrivacy` structure provides settings that determine how Amazon Lex handles special cases of securing the data for your bot.
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
	// The name of the bot locale.
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
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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

func (j *jsiiProxy_CfnBot) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lex::Bot`.
func NewCfnBot(scope constructs.Construct, id *string, props *CfnBotProps) CfnBot {
	_init_.Initialize()

	if err := validateNewCfnBotParameters(scope, id, props); err != nil {
		panic(err)
	}
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
func CfnBot_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBot_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateCfnBot_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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

	if err := validateCfnBot_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBot) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnBot) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
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

func (c *jsiiProxy_CfnBot) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBot) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
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

func (c *jsiiProxy_CfnBot) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnBot) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

