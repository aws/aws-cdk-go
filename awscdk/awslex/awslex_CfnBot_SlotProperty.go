package awslex


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

