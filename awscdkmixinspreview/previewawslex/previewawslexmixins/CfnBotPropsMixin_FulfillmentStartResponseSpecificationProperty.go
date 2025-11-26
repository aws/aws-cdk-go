package previewawslexmixins


// Provides settings for a message that is sent to the user when a fulfillment Lambda function starts running.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fulfillmentStartResponseSpecificationProperty := &FulfillmentStartResponseSpecificationProperty{
//   	AllowInterrupt: jsii.Boolean(false),
//   	DelayInSeconds: jsii.Number(123),
//   	MessageGroups: []interface{}{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-fulfillmentstartresponsespecification.html
//
type CfnBotPropsMixin_FulfillmentStartResponseSpecificationProperty struct {
	// Determines whether the user can interrupt the start message while it is playing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-fulfillmentstartresponsespecification.html#cfn-lex-bot-fulfillmentstartresponsespecification-allowinterrupt
	//
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// The delay between when the Lambda fulfillment function starts running and the start message is played.
	//
	// If the Lambda function returns before the delay is over, the start message isn't played.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-fulfillmentstartresponsespecification.html#cfn-lex-bot-fulfillmentstartresponsespecification-delayinseconds
	//
	DelayInSeconds *float64 `field:"optional" json:"delayInSeconds" yaml:"delayInSeconds"`
	// 1 - 5 message groups that contain start messages.
	//
	// Amazon Lex chooses one of the messages to play to the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-fulfillmentstartresponsespecification.html#cfn-lex-bot-fulfillmentstartresponsespecification-messagegroups
	//
	MessageGroups interface{} `field:"optional" json:"messageGroups" yaml:"messageGroups"`
}

