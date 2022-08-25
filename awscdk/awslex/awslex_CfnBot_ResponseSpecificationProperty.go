package awslex


// Specifies a list of message groups that Amazon Lex uses to respond to user input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responseSpecificationProperty := &responseSpecificationProperty{
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
//
//   	// the properties below are optional
//   	allowInterrupt: jsii.Boolean(false),
//   }
//
type CfnBot_ResponseSpecificationProperty struct {
	// A collection of responses that Amazon Lex can send to the user.
	//
	// Amazon Lex chooses the actual response to send at runtime.
	MessageGroupsList interface{} `field:"required" json:"messageGroupsList" yaml:"messageGroupsList"`
	// Indicates whether the user can interrupt a speech response from Amazon Lex .
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
}

