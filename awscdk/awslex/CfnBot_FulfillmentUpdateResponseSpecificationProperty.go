package awslex


// Provides settings for a message that is sent periodically to the user while a fulfillment Lambda function is running.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fulfillmentUpdateResponseSpecificationProperty := &FulfillmentUpdateResponseSpecificationProperty{
//   	FrequencyInSeconds: jsii.Number(123),
//   	MessageGroups: []interface{}{
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
//
//   	// the properties below are optional
//   	AllowInterrupt: jsii.Boolean(false),
//   }
//
type CfnBot_FulfillmentUpdateResponseSpecificationProperty struct {
	// The frequency that a message is sent to the user.
	//
	// When the period ends, Amazon Lex chooses a message from the message groups and plays it to the user. If the fulfillment Lambda returns before the first period ends, an update message is not played to the user.
	FrequencyInSeconds *float64 `field:"required" json:"frequencyInSeconds" yaml:"frequencyInSeconds"`
	// 1 - 5 message groups that contain update messages.
	//
	// Amazon Lex chooses one of the messages to play to the user.
	MessageGroups interface{} `field:"required" json:"messageGroups" yaml:"messageGroups"`
	// Determines whether the user can interrupt an update message while it is playing.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
}

