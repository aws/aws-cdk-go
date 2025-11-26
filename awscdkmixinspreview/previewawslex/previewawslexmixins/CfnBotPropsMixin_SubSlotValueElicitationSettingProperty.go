package previewawslexmixins


// Subslot elicitation settings.
//
// `DefaultValueSpecification` is a list of default values for a constituent sub slot in a composite slot. Default values are used when Amazon Lex hasn't determined a value for a slot. You can specify default values from context variables, session attributes, and defined values. This is similar to `DefaultValueSpecification` for slots.
//
// `PromptSpecification` is the prompt that Amazon Lex uses to elicit the sub slot value from the user. This is similar to `PromptSpecification` for slots.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subSlotValueElicitationSettingProperty := &SubSlotValueElicitationSettingProperty{
//   	DefaultValueSpecification: &SlotDefaultValueSpecificationProperty{
//   		DefaultValueList: []interface{}{
//   			&SlotDefaultValueProperty{
//   				DefaultValue: jsii.String("defaultValue"),
//   			},
//   		},
//   	},
//   	PromptSpecification: &PromptSpecificationProperty{
//   		AllowInterrupt: jsii.Boolean(false),
//   		MaxRetries: jsii.Number(123),
//   		MessageGroupsList: []interface{}{
//   			&MessageGroupProperty{
//   				Message: &MessageProperty{
//   					CustomPayload: &CustomPayloadProperty{
//   						Value: jsii.String("value"),
//   					},
//   					ImageResponseCard: &ImageResponseCardProperty{
//   						Buttons: []interface{}{
//   							&ButtonProperty{
//   								Text: jsii.String("text"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						ImageUrl: jsii.String("imageUrl"),
//   						Subtitle: jsii.String("subtitle"),
//   						Title: jsii.String("title"),
//   					},
//   					PlainTextMessage: &PlainTextMessageProperty{
//   						Value: jsii.String("value"),
//   					},
//   					SsmlMessage: &SSMLMessageProperty{
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Variations: []interface{}{
//   					&MessageProperty{
//   						CustomPayload: &CustomPayloadProperty{
//   							Value: jsii.String("value"),
//   						},
//   						ImageResponseCard: &ImageResponseCardProperty{
//   							Buttons: []interface{}{
//   								&ButtonProperty{
//   									Text: jsii.String("text"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							ImageUrl: jsii.String("imageUrl"),
//   							Subtitle: jsii.String("subtitle"),
//   							Title: jsii.String("title"),
//   						},
//   						PlainTextMessage: &PlainTextMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   						SsmlMessage: &SSMLMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		MessageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   		PromptAttemptsSpecification: map[string]interface{}{
//   			"promptAttemptsSpecificationKey": &PromptAttemptSpecificationProperty{
//   				"allowedInputTypes": &AllowedInputTypesProperty{
//   					"allowAudioInput": jsii.Boolean(false),
//   					"allowDtmfInput": jsii.Boolean(false),
//   				},
//   				"allowInterrupt": jsii.Boolean(false),
//   				"audioAndDtmfInputSpecification": &AudioAndDTMFInputSpecificationProperty{
//   					"audioSpecification": &AudioSpecificationProperty{
//   						"endTimeoutMs": jsii.Number(123),
//   						"maxLengthMs": jsii.Number(123),
//   					},
//   					"dtmfSpecification": &DTMFSpecificationProperty{
//   						"deletionCharacter": jsii.String("deletionCharacter"),
//   						"endCharacter": jsii.String("endCharacter"),
//   						"endTimeoutMs": jsii.Number(123),
//   						"maxLength": jsii.Number(123),
//   					},
//   					"startTimeoutMs": jsii.Number(123),
//   				},
//   				"textInputSpecification": &TextInputSpecificationProperty{
//   					"startTimeoutMs": jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	SampleUtterances: []interface{}{
//   		&SampleUtteranceProperty{
//   			Utterance: jsii.String("utterance"),
//   		},
//   	},
//   	WaitAndContinueSpecification: &WaitAndContinueSpecificationProperty{
//   		ContinueResponse: &ResponseSpecificationProperty{
//   			AllowInterrupt: jsii.Boolean(false),
//   			MessageGroupsList: []interface{}{
//   				&MessageGroupProperty{
//   					Message: &MessageProperty{
//   						CustomPayload: &CustomPayloadProperty{
//   							Value: jsii.String("value"),
//   						},
//   						ImageResponseCard: &ImageResponseCardProperty{
//   							Buttons: []interface{}{
//   								&ButtonProperty{
//   									Text: jsii.String("text"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							ImageUrl: jsii.String("imageUrl"),
//   							Subtitle: jsii.String("subtitle"),
//   							Title: jsii.String("title"),
//   						},
//   						PlainTextMessage: &PlainTextMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   						SsmlMessage: &SSMLMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Variations: []interface{}{
//   						&MessageProperty{
//   							CustomPayload: &CustomPayloadProperty{
//   								Value: jsii.String("value"),
//   							},
//   							ImageResponseCard: &ImageResponseCardProperty{
//   								Buttons: []interface{}{
//   									&ButtonProperty{
//   										Text: jsii.String("text"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								ImageUrl: jsii.String("imageUrl"),
//   								Subtitle: jsii.String("subtitle"),
//   								Title: jsii.String("title"),
//   							},
//   							PlainTextMessage: &PlainTextMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   							SsmlMessage: &SSMLMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		IsActive: jsii.Boolean(false),
//   		StillWaitingResponse: &StillWaitingResponseSpecificationProperty{
//   			AllowInterrupt: jsii.Boolean(false),
//   			FrequencyInSeconds: jsii.Number(123),
//   			MessageGroupsList: []interface{}{
//   				&MessageGroupProperty{
//   					Message: &MessageProperty{
//   						CustomPayload: &CustomPayloadProperty{
//   							Value: jsii.String("value"),
//   						},
//   						ImageResponseCard: &ImageResponseCardProperty{
//   							Buttons: []interface{}{
//   								&ButtonProperty{
//   									Text: jsii.String("text"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							ImageUrl: jsii.String("imageUrl"),
//   							Subtitle: jsii.String("subtitle"),
//   							Title: jsii.String("title"),
//   						},
//   						PlainTextMessage: &PlainTextMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   						SsmlMessage: &SSMLMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Variations: []interface{}{
//   						&MessageProperty{
//   							CustomPayload: &CustomPayloadProperty{
//   								Value: jsii.String("value"),
//   							},
//   							ImageResponseCard: &ImageResponseCardProperty{
//   								Buttons: []interface{}{
//   									&ButtonProperty{
//   										Text: jsii.String("text"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								ImageUrl: jsii.String("imageUrl"),
//   								Subtitle: jsii.String("subtitle"),
//   								Title: jsii.String("title"),
//   							},
//   							PlainTextMessage: &PlainTextMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   							SsmlMessage: &SSMLMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			TimeoutInSeconds: jsii.Number(123),
//   		},
//   		WaitingResponse: &ResponseSpecificationProperty{
//   			AllowInterrupt: jsii.Boolean(false),
//   			MessageGroupsList: []interface{}{
//   				&MessageGroupProperty{
//   					Message: &MessageProperty{
//   						CustomPayload: &CustomPayloadProperty{
//   							Value: jsii.String("value"),
//   						},
//   						ImageResponseCard: &ImageResponseCardProperty{
//   							Buttons: []interface{}{
//   								&ButtonProperty{
//   									Text: jsii.String("text"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							ImageUrl: jsii.String("imageUrl"),
//   							Subtitle: jsii.String("subtitle"),
//   							Title: jsii.String("title"),
//   						},
//   						PlainTextMessage: &PlainTextMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   						SsmlMessage: &SSMLMessageProperty{
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Variations: []interface{}{
//   						&MessageProperty{
//   							CustomPayload: &CustomPayloadProperty{
//   								Value: jsii.String("value"),
//   							},
//   							ImageResponseCard: &ImageResponseCardProperty{
//   								Buttons: []interface{}{
//   									&ButtonProperty{
//   										Text: jsii.String("text"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								ImageUrl: jsii.String("imageUrl"),
//   								Subtitle: jsii.String("subtitle"),
//   								Title: jsii.String("title"),
//   							},
//   							PlainTextMessage: &PlainTextMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   							SsmlMessage: &SSMLMessageProperty{
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslotvalueelicitationsetting.html
//
type CfnBotPropsMixin_SubSlotValueElicitationSettingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslotvalueelicitationsetting.html#cfn-lex-bot-subslotvalueelicitationsetting-defaultvaluespecification
	//
	DefaultValueSpecification interface{} `field:"optional" json:"defaultValueSpecification" yaml:"defaultValueSpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslotvalueelicitationsetting.html#cfn-lex-bot-subslotvalueelicitationsetting-promptspecification
	//
	PromptSpecification interface{} `field:"optional" json:"promptSpecification" yaml:"promptSpecification"`
	// If you know a specific pattern that users might respond to an Amazon Lex request for a sub slot value, you can provide those utterances to improve accuracy.
	//
	// This is optional. In most cases Amazon Lex is capable of understanding user utterances. This is similar to `SampleUtterances` for slots.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslotvalueelicitationsetting.html#cfn-lex-bot-subslotvalueelicitationsetting-sampleutterances
	//
	SampleUtterances interface{} `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-subslotvalueelicitationsetting.html#cfn-lex-bot-subslotvalueelicitationsetting-waitandcontinuespecification
	//
	WaitAndContinueSpecification interface{} `field:"optional" json:"waitAndContinueSpecification" yaml:"waitAndContinueSpecification"`
}

