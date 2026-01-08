package awslex


// Specifications for the constituent sub slots and the expression for the composite slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subSlotSettingProperty := &SubSlotSettingProperty{
//   	Expression: jsii.String("expression"),
//   	SlotSpecifications: map[string]interface{}{
//   		"slotSpecificationsKey": &SpecificationsProperty{
//   			"valueElicitationSetting": &SubSlotValueElicitationSettingProperty{
//   				"defaultValueSpecification": &SlotDefaultValueSpecificationProperty{
//   					"defaultValueList": []interface{}{
//   						&SlotDefaultValueProperty{
//   							"defaultValue": jsii.String("defaultValue"),
//   						},
//   					},
//   				},
//   				"promptSpecification": &PromptSpecificationProperty{
//   					"maxRetries": jsii.Number(123),
//   					"messageGroupsList": []interface{}{
//   						&MessageGroupProperty{
//   							"message": &MessageProperty{
//   								"customPayload": &CustomPayloadProperty{
//   									"value": jsii.String("value"),
//   								},
//   								"imageResponseCard": &ImageResponseCardProperty{
//   									"title": jsii.String("title"),
//
//   									// the properties below are optional
//   									"buttons": []interface{}{
//   										&ButtonProperty{
//   											"text": jsii.String("text"),
//   											"value": jsii.String("value"),
//   										},
//   									},
//   									"imageUrl": jsii.String("imageUrl"),
//   									"subtitle": jsii.String("subtitle"),
//   								},
//   								"plainTextMessage": &PlainTextMessageProperty{
//   									"value": jsii.String("value"),
//   								},
//   								"ssmlMessage": &SSMLMessageProperty{
//   									"value": jsii.String("value"),
//   								},
//   							},
//
//   							// the properties below are optional
//   							"variations": []interface{}{
//   								&MessageProperty{
//   									"customPayload": &CustomPayloadProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"imageResponseCard": &ImageResponseCardProperty{
//   										"title": jsii.String("title"),
//
//   										// the properties below are optional
//   										"buttons": []interface{}{
//   											&ButtonProperty{
//   												"text": jsii.String("text"),
//   												"value": jsii.String("value"),
//   											},
//   										},
//   										"imageUrl": jsii.String("imageUrl"),
//   										"subtitle": jsii.String("subtitle"),
//   									},
//   									"plainTextMessage": &PlainTextMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"ssmlMessage": &SSMLMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					"allowInterrupt": jsii.Boolean(false),
//   					"messageSelectionStrategy": jsii.String("messageSelectionStrategy"),
//   					"promptAttemptsSpecification": map[string]interface{}{
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
//   				"sampleUtterances": []interface{}{
//   					&SampleUtteranceProperty{
//   						"utterance": jsii.String("utterance"),
//   					},
//   				},
//   				"waitAndContinueSpecification": &WaitAndContinueSpecificationProperty{
//   					"continueResponse": &ResponseSpecificationProperty{
//   						"messageGroupsList": []interface{}{
//   							&MessageGroupProperty{
//   								"message": &MessageProperty{
//   									"customPayload": &CustomPayloadProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"imageResponseCard": &ImageResponseCardProperty{
//   										"title": jsii.String("title"),
//
//   										// the properties below are optional
//   										"buttons": []interface{}{
//   											&ButtonProperty{
//   												"text": jsii.String("text"),
//   												"value": jsii.String("value"),
//   											},
//   										},
//   										"imageUrl": jsii.String("imageUrl"),
//   										"subtitle": jsii.String("subtitle"),
//   									},
//   									"plainTextMessage": &PlainTextMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"ssmlMessage": &SSMLMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								"variations": []interface{}{
//   									&MessageProperty{
//   										"customPayload": &CustomPayloadProperty{
//   											"value": jsii.String("value"),
//   										},
//   										"imageResponseCard": &ImageResponseCardProperty{
//   											"title": jsii.String("title"),
//
//   											// the properties below are optional
//   											"buttons": []interface{}{
//   												&ButtonProperty{
//   													"text": jsii.String("text"),
//   													"value": jsii.String("value"),
//   												},
//   											},
//   											"imageUrl": jsii.String("imageUrl"),
//   											"subtitle": jsii.String("subtitle"),
//   										},
//   										"plainTextMessage": &PlainTextMessageProperty{
//   											"value": jsii.String("value"),
//   										},
//   										"ssmlMessage": &SSMLMessageProperty{
//   											"value": jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						"allowInterrupt": jsii.Boolean(false),
//   					},
//   					"waitingResponse": &ResponseSpecificationProperty{
//   						"messageGroupsList": []interface{}{
//   							&MessageGroupProperty{
//   								"message": &MessageProperty{
//   									"customPayload": &CustomPayloadProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"imageResponseCard": &ImageResponseCardProperty{
//   										"title": jsii.String("title"),
//
//   										// the properties below are optional
//   										"buttons": []interface{}{
//   											&ButtonProperty{
//   												"text": jsii.String("text"),
//   												"value": jsii.String("value"),
//   											},
//   										},
//   										"imageUrl": jsii.String("imageUrl"),
//   										"subtitle": jsii.String("subtitle"),
//   									},
//   									"plainTextMessage": &PlainTextMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"ssmlMessage": &SSMLMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								"variations": []interface{}{
//   									&MessageProperty{
//   										"customPayload": &CustomPayloadProperty{
//   											"value": jsii.String("value"),
//   										},
//   										"imageResponseCard": &ImageResponseCardProperty{
//   											"title": jsii.String("title"),
//
//   											// the properties below are optional
//   											"buttons": []interface{}{
//   												&ButtonProperty{
//   													"text": jsii.String("text"),
//   													"value": jsii.String("value"),
//   												},
//   											},
//   											"imageUrl": jsii.String("imageUrl"),
//   											"subtitle": jsii.String("subtitle"),
//   										},
//   										"plainTextMessage": &PlainTextMessageProperty{
//   											"value": jsii.String("value"),
//   										},
//   										"ssmlMessage": &SSMLMessageProperty{
//   											"value": jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						"allowInterrupt": jsii.Boolean(false),
//   					},
//
//   					// the properties below are optional
//   					"isActive": jsii.Boolean(false),
//   					"stillWaitingResponse": &StillWaitingResponseSpecificationProperty{
//   						"frequencyInSeconds": jsii.Number(123),
//   						"messageGroupsList": []interface{}{
//   							&MessageGroupProperty{
//   								"message": &MessageProperty{
//   									"customPayload": &CustomPayloadProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"imageResponseCard": &ImageResponseCardProperty{
//   										"title": jsii.String("title"),
//
//   										// the properties below are optional
//   										"buttons": []interface{}{
//   											&ButtonProperty{
//   												"text": jsii.String("text"),
//   												"value": jsii.String("value"),
//   											},
//   										},
//   										"imageUrl": jsii.String("imageUrl"),
//   										"subtitle": jsii.String("subtitle"),
//   									},
//   									"plainTextMessage": &PlainTextMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"ssmlMessage": &SSMLMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   								},
//
//   								// the properties below are optional
//   								"variations": []interface{}{
//   									&MessageProperty{
//   										"customPayload": &CustomPayloadProperty{
//   											"value": jsii.String("value"),
//   										},
//   										"imageResponseCard": &ImageResponseCardProperty{
//   											"title": jsii.String("title"),
//
//   											// the properties below are optional
//   											"buttons": []interface{}{
//   												&ButtonProperty{
//   													"text": jsii.String("text"),
//   													"value": jsii.String("value"),
//   												},
//   											},
//   											"imageUrl": jsii.String("imageUrl"),
//   											"subtitle": jsii.String("subtitle"),
//   										},
//   										"plainTextMessage": &PlainTextMessageProperty{
//   											"value": jsii.String("value"),
//   										},
//   										"ssmlMessage": &SSMLMessageProperty{
//   											"value": jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						"timeoutInSeconds": jsii.Number(123),
//
//   						// the properties below are optional
//   						"allowInterrupt": jsii.Boolean(false),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			"slotTypeId": jsii.String("slotTypeId"),
//   			"slotTypeName": jsii.String("slotTypeName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslotsetting.html
//
type CfnBot_SubSlotSettingProperty struct {
	// The expression text for defining the constituent sub slots in the composite slot using logical AND and OR operators.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslotsetting.html#cfn-lex-bot-subslotsetting-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Specifications for the constituent sub slots of a composite slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslotsetting.html#cfn-lex-bot-subslotsetting-slotspecifications
	//
	SlotSpecifications interface{} `field:"optional" json:"slotSpecifications" yaml:"slotSpecifications"`
}

