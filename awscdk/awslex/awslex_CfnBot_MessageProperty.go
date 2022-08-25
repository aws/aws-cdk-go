package awslex


// The object that provides message text and it's type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageProperty := &messageProperty{
//   	customPayload: &customPayloadProperty{
//   		value: jsii.String("value"),
//   	},
//   	imageResponseCard: &imageResponseCardProperty{
//   		title: jsii.String("title"),
//
//   		// the properties below are optional
//   		buttons: []interface{}{
//   			&buttonProperty{
//   				text: jsii.String("text"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		imageUrl: jsii.String("imageUrl"),
//   		subtitle: jsii.String("subtitle"),
//   	},
//   	plainTextMessage: &plainTextMessageProperty{
//   		value: jsii.String("value"),
//   	},
//   	ssmlMessage: &sSMLMessageProperty{
//   		value: jsii.String("value"),
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
	// A message in Speech Synthesis Markup Language (SSML) format.
	SsmlMessage interface{} `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

