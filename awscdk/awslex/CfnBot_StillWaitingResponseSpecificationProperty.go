package awslex


// Defines the messages that Amazon Lex sends to a user to remind them that the bot is waiting for a response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stillWaitingResponseSpecificationProperty := &StillWaitingResponseSpecificationProperty{
//   	FrequencyInSeconds: jsii.Number(123),
//   	MessageGroupsList: []interface{}{
//   		&MessageGroupProperty{
//   			Message: &MessageProperty{
//   				CustomPayload: &CustomPayloadProperty{
//   					Value: jsii.String("value"),
//   				},
//   				ImageResponseCard: &ImageResponseCardProperty{
//   					Title: jsii.String("title"),
//
//   					// the properties below are optional
//   					Buttons: []interface{}{
//   						&ButtonProperty{
//   							Text: jsii.String("text"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					ImageUrl: jsii.String("imageUrl"),
//   					Subtitle: jsii.String("subtitle"),
//   				},
//   				PlainTextMessage: &PlainTextMessageProperty{
//   					Value: jsii.String("value"),
//   				},
//   				SsmlMessage: &SSMLMessageProperty{
//   					Value: jsii.String("value"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Variations: []interface{}{
//   				&MessageProperty{
//   					CustomPayload: &CustomPayloadProperty{
//   						Value: jsii.String("value"),
//   					},
//   					ImageResponseCard: &ImageResponseCardProperty{
//   						Title: jsii.String("title"),
//
//   						// the properties below are optional
//   						Buttons: []interface{}{
//   							&ButtonProperty{
//   								Text: jsii.String("text"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						ImageUrl: jsii.String("imageUrl"),
//   						Subtitle: jsii.String("subtitle"),
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
//   	TimeoutInSeconds: jsii.Number(123),
//
//   	// the properties below are optional
//   	AllowInterrupt: jsii.Boolean(false),
//   }
//
type CfnBot_StillWaitingResponseSpecificationProperty struct {
	// How often a message should be sent to the user.
	//
	// Minimum of 1 second, maximum of 5 minutes.
	FrequencyInSeconds *float64 `field:"required" json:"frequencyInSeconds" yaml:"frequencyInSeconds"`
	// One or more message groups, each containing one or more messages, that define the prompts that Amazon Lex sends to the user.
	MessageGroupsList interface{} `field:"required" json:"messageGroupsList" yaml:"messageGroupsList"`
	// If Amazon Lex waits longer than this length of time for a response, it will stop sending messages.
	TimeoutInSeconds *float64 `field:"required" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// Indicates that the user can interrupt the response by speaking while the message is being played.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
}

