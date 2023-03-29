package awslex


// The object that provides message text and its type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageProperty := &MessageProperty{
//   	CustomPayload: &CustomPayloadProperty{
//   		Value: jsii.String("value"),
//   	},
//   	ImageResponseCard: &ImageResponseCardProperty{
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		Buttons: []interface{}{
//   			&ButtonProperty{
//   				Text: jsii.String("text"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ImageUrl: jsii.String("imageUrl"),
//   		Subtitle: jsii.String("subtitle"),
//   	},
//   	PlainTextMessage: &PlainTextMessageProperty{
//   		Value: jsii.String("value"),
//   	},
//   	SsmlMessage: &SSMLMessageProperty{
//   		Value: jsii.String("value"),
//   	},
//   }
//
type CfnBot_MessageProperty struct {
	// A message in a custom format defined by the client application.
	CustomPayload interface{} `field:"optional" json:"customPayload" yaml:"customPayload"`
	// A message that defines a response card that the client application can show to the user.
	ImageResponseCard interface{} `field:"optional" json:"imageResponseCard" yaml:"imageResponseCard"`
	// A message in plain text format.
	PlainTextMessage interface{} `field:"optional" json:"plainTextMessage" yaml:"plainTextMessage"`
	// A message in Speech Synthesis Markup Language (SSML).
	SsmlMessage interface{} `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

