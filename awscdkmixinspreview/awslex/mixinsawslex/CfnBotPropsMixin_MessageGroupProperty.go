package mixinsawslex


// Provides one or more messages that Amazon Lex should send to the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   messageGroupProperty := &MessageGroupProperty{
//   	Message: &MessageProperty{
//   		CustomPayload: &CustomPayloadProperty{
//   			Value: jsii.String("value"),
//   		},
//   		ImageResponseCard: &ImageResponseCardProperty{
//   			Buttons: []interface{}{
//   				&ButtonProperty{
//   					Text: jsii.String("text"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			ImageUrl: jsii.String("imageUrl"),
//   			Subtitle: jsii.String("subtitle"),
//   			Title: jsii.String("title"),
//   		},
//   		PlainTextMessage: &PlainTextMessageProperty{
//   			Value: jsii.String("value"),
//   		},
//   		SsmlMessage: &SSMLMessageProperty{
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Variations: []interface{}{
//   		&MessageProperty{
//   			CustomPayload: &CustomPayloadProperty{
//   				Value: jsii.String("value"),
//   			},
//   			ImageResponseCard: &ImageResponseCardProperty{
//   				Buttons: []interface{}{
//   					&ButtonProperty{
//   						Text: jsii.String("text"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				ImageUrl: jsii.String("imageUrl"),
//   				Subtitle: jsii.String("subtitle"),
//   				Title: jsii.String("title"),
//   			},
//   			PlainTextMessage: &PlainTextMessageProperty{
//   				Value: jsii.String("value"),
//   			},
//   			SsmlMessage: &SSMLMessageProperty{
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-messagegroup.html
//
type CfnBotPropsMixin_MessageGroupProperty struct {
	// The primary message that Amazon Lex should send to the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-messagegroup.html#cfn-lex-bot-messagegroup-message
	//
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// Message variations to send to the user.
	//
	// When variations are defined, Amazon Lex chooses the primary message or one of the variations to send to the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-messagegroup.html#cfn-lex-bot-messagegroup-variations
	//
	Variations interface{} `field:"optional" json:"variations" yaml:"variations"`
}

