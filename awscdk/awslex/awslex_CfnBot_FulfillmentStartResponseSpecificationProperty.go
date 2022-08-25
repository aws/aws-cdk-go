package awslex


// Provides settings for a message that is sent to the user when a fulfillment Lambda function starts running.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fulfillmentStartResponseSpecificationProperty := &fulfillmentStartResponseSpecificationProperty{
//   	delayInSeconds: jsii.Number(123),
//   	messageGroups: []interface{}{
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
//   }
//
type CfnBot_FulfillmentStartResponseSpecificationProperty struct {
	// The delay between when the Lambda fulfillment function starts running and the start message is played.
	//
	// If the Lambda function returns before the delay is over, the start message isn't played.
	DelayInSeconds *float64 `field:"required" json:"delayInSeconds" yaml:"delayInSeconds"`
	// One to 5 message groups that contain start messages.
	//
	// Amazon Lex chooses one of the messages to play to the user.
	MessageGroups interface{} `field:"required" json:"messageGroups" yaml:"messageGroups"`
	// Determines whether the user can interrupt the start message while it is playing.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
}

