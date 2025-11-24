package mixinsawslex


// Specifies a list of message groups that Amazon Lex sends to a user to elicit a response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   promptSpecificationProperty := &PromptSpecificationProperty{
//   	AllowInterrupt: jsii.Boolean(false),
//   	MaxRetries: jsii.Number(123),
//   	MessageGroupsList: []interface{}{
//   		&MessageGroupProperty{
//   			Message: &MessageProperty{
//   				CustomPayload: &CustomPayloadProperty{
//   					Value: jsii.String("value"),
//   				},
//   				ImageResponseCard: &ImageResponseCardProperty{
//   					Buttons: []interface{}{
//   						&ButtonProperty{
//   							Text: jsii.String("text"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					ImageUrl: jsii.String("imageUrl"),
//   					Subtitle: jsii.String("subtitle"),
//   					Title: jsii.String("title"),
//   				},
//   				PlainTextMessage: &PlainTextMessageProperty{
//   					Value: jsii.String("value"),
//   				},
//   				SsmlMessage: &SSMLMessageProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Variations: []interface{}{
//   				&MessageProperty{
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
//   			},
//   		},
//   	},
//   	MessageSelectionStrategy: jsii.String("messageSelectionStrategy"),
//   	PromptAttemptsSpecification: map[string]interface{}{
//   		"promptAttemptsSpecificationKey": &PromptAttemptSpecificationProperty{
//   			"allowedInputTypes": &AllowedInputTypesProperty{
//   				"allowAudioInput": jsii.Boolean(false),
//   				"allowDtmfInput": jsii.Boolean(false),
//   			},
//   			"allowInterrupt": jsii.Boolean(false),
//   			"audioAndDtmfInputSpecification": &AudioAndDTMFInputSpecificationProperty{
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
//   				"startTimeoutMs": jsii.Number(123),
//   			},
//   			"textInputSpecification": &TextInputSpecificationProperty{
//   				"startTimeoutMs": jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-promptspecification.html
//
type CfnBotPropsMixin_PromptSpecificationProperty struct {
	// Indicates whether the user can interrupt a speech prompt from the bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-promptspecification.html#cfn-lex-bot-promptspecification-allowinterrupt
	//
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// The maximum number of times the bot tries to elicit a response from the user using this prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-promptspecification.html#cfn-lex-bot-promptspecification-maxretries
	//
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// A collection of messages that Amazon Lex can send to the user.
	//
	// Amazon Lex chooses the actual message to send at runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-promptspecification.html#cfn-lex-bot-promptspecification-messagegroupslist
	//
	MessageGroupsList interface{} `field:"optional" json:"messageGroupsList" yaml:"messageGroupsList"`
	// Indicates how a message is selected from a message group among retries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-promptspecification.html#cfn-lex-bot-promptspecification-messageselectionstrategy
	//
	MessageSelectionStrategy *string `field:"optional" json:"messageSelectionStrategy" yaml:"messageSelectionStrategy"`
	// Specifies the advanced settings on each attempt of the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-promptspecification.html#cfn-lex-bot-promptspecification-promptattemptsspecification
	//
	PromptAttemptsSpecification interface{} `field:"optional" json:"promptAttemptsSpecification" yaml:"promptAttemptsSpecification"`
}

