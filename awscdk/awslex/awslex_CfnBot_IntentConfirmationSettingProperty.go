package awslex


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
//   		messageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   		promptAttemptsSpecification: map[string]interface{}{
//   			"promptAttemptsSpecificationKey": &PromptAttemptSpecificationProperty{
//   				"allowedInputTypes": &AllowedInputTypesProperty{
//   					"allowAudioInput": jsii.Boolean(false),
//   					"allowDtmfInput": jsii.Boolean(false),
//   				},
//
//   				// the properties below are optional
//   				"allowInterrupt": jsii.Boolean(false),
//   				"audioAndDtmfInputSpecification": &AudioAndDTMFInputSpecificationProperty{
//   					"startTimeoutMs": jsii.Number(123),
//
//   					// the properties below are optional
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
//   				},
//   				"textInputSpecification": &TextInputSpecificationProperty{
//   					"startTimeoutMs": jsii.Number(123),
//   				},
//   			},
//   		},
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

