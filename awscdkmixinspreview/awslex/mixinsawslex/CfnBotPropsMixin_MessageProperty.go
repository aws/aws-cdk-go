package mixinsawslex


// The object that provides message text and its type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   messageProperty := &MessageProperty{
//   	CustomPayload: &CustomPayloadProperty{
//   		Value: jsii.String("value"),
//   	},
//   	ImageResponseCard: &ImageResponseCardProperty{
//   		Buttons: []interface{}{
//   			&ButtonProperty{
//   				Text: jsii.String("text"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ImageUrl: jsii.String("imageUrl"),
//   		Subtitle: jsii.String("subtitle"),
//   		Title: jsii.String("title"),
//   	},
//   	PlainTextMessage: &PlainTextMessageProperty{
//   		Value: jsii.String("value"),
//   	},
//   	SsmlMessage: &SSMLMessageProperty{
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-message.html
//
type CfnBotPropsMixin_MessageProperty struct {
	// A message in a custom format defined by the client application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-message.html#cfn-lex-bot-message-custompayload
	//
	CustomPayload interface{} `field:"optional" json:"customPayload" yaml:"customPayload"`
	// A message that defines a response card that the client application can show to the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-message.html#cfn-lex-bot-message-imageresponsecard
	//
	ImageResponseCard interface{} `field:"optional" json:"imageResponseCard" yaml:"imageResponseCard"`
	// A message in plain text format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-message.html#cfn-lex-bot-message-plaintextmessage
	//
	PlainTextMessage interface{} `field:"optional" json:"plainTextMessage" yaml:"plainTextMessage"`
	// A message in Speech Synthesis Markup Language (SSML).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-message.html#cfn-lex-bot-message-ssmlmessage
	//
	SsmlMessage interface{} `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

