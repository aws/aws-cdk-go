package awslex


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
//   	messageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   	promptAttemptsSpecification: map[string]interface{}{
//   		"promptAttemptsSpecificationKey": &PromptAttemptSpecificationProperty{
//   			"allowedInputTypes": &AllowedInputTypesProperty{
//   				"allowAudioInput": jsii.Boolean(false),
//   				"allowDtmfInput": jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			"allowInterrupt": jsii.Boolean(false),
//   			"audioAndDtmfInputSpecification": &AudioAndDTMFInputSpecificationProperty{
//   				"startTimeoutMs": jsii.Number(123),
//
//   				// the properties below are optional
//   				"audioSpecification": &AudioSpecificationProperty{
//   					"endTimeoutMs": jsii.Number(123),
//   					"maxLengthMs": jsii.Number(123),
//   				},
//   				"dtmfSpecification": &DTMFSpecificationProperty{
//   					"deletionCharacter": jsii.String("deletionCharacter"),
//   					"endCharacter": jsii.String("endCharacter"),
//   					"endTimeoutMs": jsii.Number(123),
//   					"maxLength": jsii.Number(123),
//   				},
//   			},
//   			"textInputSpecification": &TextInputSpecificationProperty{
//   				"startTimeoutMs": jsii.Number(123),
//   			},
//   		},
//   	},
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
	// `CfnBot.PromptSpecificationProperty.MessageSelectionStrategy`.
	MessageSelectionStrategy *string `field:"optional" json:"messageSelectionStrategy" yaml:"messageSelectionStrategy"`
	// `CfnBot.PromptSpecificationProperty.PromptAttemptsSpecification`.
	PromptAttemptsSpecification interface{} `field:"optional" json:"promptAttemptsSpecification" yaml:"promptAttemptsSpecification"`
}

