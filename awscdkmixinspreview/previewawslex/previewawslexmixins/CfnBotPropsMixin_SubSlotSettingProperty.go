package previewawslexmixins


// Specifications for the constituent sub slots and the expression for the composite slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subSlotSettingProperty := &SubSlotSettingProperty{
//   	Expression: jsii.String("expression"),
//   	SlotSpecifications: map[string]interface{}{
//   		"slotSpecificationsKey": &SpecificationsProperty{
//   			"slotTypeId": jsii.String("slotTypeId"),
//   			"valueElicitationSetting": &SubSlotValueElicitationSettingProperty{
//   				"defaultValueSpecification": &SlotDefaultValueSpecificationProperty{
//   					"defaultValueList": []interface{}{
//   						&SlotDefaultValueProperty{
//   							"defaultValue": jsii.String("defaultValue"),
//   						},
//   					},
//   				},
//   				"promptSpecification": &PromptSpecificationProperty{
//   					"allowInterrupt": jsii.Boolean(false),
//   					"maxRetries": jsii.Number(123),
//   					"messageGroupsList": []interface{}{
//   						&MessageGroupProperty{
//   							"message": &MessageProperty{
//   								"customPayload": &CustomPayloadProperty{
//   									"value": jsii.String("value"),
//   								},
//   								"imageResponseCard": &ImageResponseCardProperty{
//   									"buttons": []interface{}{
//   										&ButtonProperty{
//   											"text": jsii.String("text"),
//   											"value": jsii.String("value"),
//   										},
//   									},
//   									"imageUrl": jsii.String("imageUrl"),
//   									"subtitle": jsii.String("subtitle"),
//   									"title": jsii.String("title"),
//   								},
//   								"plainTextMessage": &PlainTextMessageProperty{
//   									"value": jsii.String("value"),
//   								},
//   								"ssmlMessage": &SSMLMessageProperty{
//   									"value": jsii.String("value"),
//   								},
//   							},
//   							"variations": []interface{}{
//   								&MessageProperty{
//   									"customPayload": &CustomPayloadProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"imageResponseCard": &ImageResponseCardProperty{
//   										"buttons": []interface{}{
//   											&ButtonProperty{
//   												"text": jsii.String("text"),
//   												"value": jsii.String("value"),
//   											},
//   										},
//   										"imageUrl": jsii.String("imageUrl"),
//   										"subtitle": jsii.String("subtitle"),
//   										"title": jsii.String("title"),
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
//   					"messageSelectionStrategy": jsii.String("messageSelectionStrategy"),
//   					"promptAttemptsSpecification": map[string]interface{}{
//   						"promptAttemptsSpecificationKey": &PromptAttemptSpecificationProperty{
//   							"allowedInputTypes": &AllowedInputTypesProperty{
//   								"allowAudioInput": jsii.Boolean(false),
//   								"allowDtmfInput": jsii.Boolean(false),
//   							},
//   							"allowInterrupt": jsii.Boolean(false),
//   							"audioAndDtmfInputSpecification": &AudioAndDTMFInputSpecificationProperty{
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
//   								"startTimeoutMs": jsii.Number(123),
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
//   						"allowInterrupt": jsii.Boolean(false),
//   						"messageGroupsList": []interface{}{
//   							&MessageGroupProperty{
//   								"message": &MessageProperty{
//   									"customPayload": &CustomPayloadProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"imageResponseCard": &ImageResponseCardProperty{
//   										"buttons": []interface{}{
//   											&ButtonProperty{
//   												"text": jsii.String("text"),
//   												"value": jsii.String("value"),
//   											},
//   										},
//   										"imageUrl": jsii.String("imageUrl"),
//   										"subtitle": jsii.String("subtitle"),
//   										"title": jsii.String("title"),
//   									},
//   									"plainTextMessage": &PlainTextMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"ssmlMessage": &SSMLMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   								},
//   								"variations": []interface{}{
//   									&MessageProperty{
//   										"customPayload": &CustomPayloadProperty{
//   											"value": jsii.String("value"),
//   										},
//   										"imageResponseCard": &ImageResponseCardProperty{
//   											"buttons": []interface{}{
//   												&ButtonProperty{
//   													"text": jsii.String("text"),
//   													"value": jsii.String("value"),
//   												},
//   											},
//   											"imageUrl": jsii.String("imageUrl"),
//   											"subtitle": jsii.String("subtitle"),
//   											"title": jsii.String("title"),
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
//   					},
//   					"isActive": jsii.Boolean(false),
//   					"stillWaitingResponse": &StillWaitingResponseSpecificationProperty{
//   						"allowInterrupt": jsii.Boolean(false),
//   						"frequencyInSeconds": jsii.Number(123),
//   						"messageGroupsList": []interface{}{
//   							&MessageGroupProperty{
//   								"message": &MessageProperty{
//   									"customPayload": &CustomPayloadProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"imageResponseCard": &ImageResponseCardProperty{
//   										"buttons": []interface{}{
//   											&ButtonProperty{
//   												"text": jsii.String("text"),
//   												"value": jsii.String("value"),
//   											},
//   										},
//   										"imageUrl": jsii.String("imageUrl"),
//   										"subtitle": jsii.String("subtitle"),
//   										"title": jsii.String("title"),
//   									},
//   									"plainTextMessage": &PlainTextMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"ssmlMessage": &SSMLMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   								},
//   								"variations": []interface{}{
//   									&MessageProperty{
//   										"customPayload": &CustomPayloadProperty{
//   											"value": jsii.String("value"),
//   										},
//   										"imageResponseCard": &ImageResponseCardProperty{
//   											"buttons": []interface{}{
//   												&ButtonProperty{
//   													"text": jsii.String("text"),
//   													"value": jsii.String("value"),
//   												},
//   											},
//   											"imageUrl": jsii.String("imageUrl"),
//   											"subtitle": jsii.String("subtitle"),
//   											"title": jsii.String("title"),
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
//   					},
//   					"waitingResponse": &ResponseSpecificationProperty{
//   						"allowInterrupt": jsii.Boolean(false),
//   						"messageGroupsList": []interface{}{
//   							&MessageGroupProperty{
//   								"message": &MessageProperty{
//   									"customPayload": &CustomPayloadProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"imageResponseCard": &ImageResponseCardProperty{
//   										"buttons": []interface{}{
//   											&ButtonProperty{
//   												"text": jsii.String("text"),
//   												"value": jsii.String("value"),
//   											},
//   										},
//   										"imageUrl": jsii.String("imageUrl"),
//   										"subtitle": jsii.String("subtitle"),
//   										"title": jsii.String("title"),
//   									},
//   									"plainTextMessage": &PlainTextMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   									"ssmlMessage": &SSMLMessageProperty{
//   										"value": jsii.String("value"),
//   									},
//   								},
//   								"variations": []interface{}{
//   									&MessageProperty{
//   										"customPayload": &CustomPayloadProperty{
//   											"value": jsii.String("value"),
//   										},
//   										"imageResponseCard": &ImageResponseCardProperty{
//   											"buttons": []interface{}{
//   												&ButtonProperty{
//   													"text": jsii.String("text"),
//   													"value": jsii.String("value"),
//   												},
//   											},
//   											"imageUrl": jsii.String("imageUrl"),
//   											"subtitle": jsii.String("subtitle"),
//   											"title": jsii.String("title"),
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
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslotsetting.html
//
type CfnBotPropsMixin_SubSlotSettingProperty struct {
	// The expression text for defining the constituent sub slots in the composite slot using logical AND and OR operators.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslotsetting.html#cfn-lex-bot-subslotsetting-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Specifications for the constituent sub slots of a composite slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslotsetting.html#cfn-lex-bot-subslotsetting-slotspecifications
	//
	SlotSpecifications interface{} `field:"optional" json:"slotSpecifications" yaml:"slotSpecifications"`
}

