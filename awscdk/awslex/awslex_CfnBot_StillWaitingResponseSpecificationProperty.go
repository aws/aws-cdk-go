package awslex


// Defines the messages that Amazon Lex sends to a user to remind them that the bot is waiting for a response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stillWaitingResponseSpecificationProperty := &stillWaitingResponseSpecificationProperty{
//   	frequencyInSeconds: jsii.Number(123),
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
//   	timeoutInSeconds: jsii.Number(123),
//
//   	// the properties below are optional
//   	allowInterrupt: jsii.Boolean(false),
//   }
//
type CfnBot_StillWaitingResponseSpecificationProperty struct {
	// How often a message should be sent to the user.
	//
	// Minimum of 1 second, maximum of 5 minutes.
	FrequencyInSeconds *float64 `field:"required" json:"frequencyInSeconds" yaml:"frequencyInSeconds"`
	// A collection of responses that Amazon Lex can send to the user.
	//
	// Amazon Lex chooses the actual response to send at runtime.
	MessageGroupsList interface{} `field:"required" json:"messageGroupsList" yaml:"messageGroupsList"`
	// If Amazon Lex waits longer than this length of time for a response, it will stop sending messages.
	TimeoutInSeconds *float64 `field:"required" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// Indicates that the user can interrupt the response by speaking while the message is being played.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
}

