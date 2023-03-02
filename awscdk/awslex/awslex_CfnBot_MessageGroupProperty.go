package awslex


// Provides one or more messages that Amazon Lex should send to the user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageGroupProperty := &messageGroupProperty{
//   	message: &messageProperty{
//   		customPayload: &customPayloadProperty{
//   			value: jsii.String("value"),
//   		},
//   		imageResponseCard: &imageResponseCardProperty{
//   			title: jsii.String("title"),
//
//   			// the properties below are optional
//   			buttons: []interface{}{
//   				&buttonProperty{
//   					text: jsii.String("text"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			imageUrl: jsii.String("imageUrl"),
//   			subtitle: jsii.String("subtitle"),
//   		},
//   		plainTextMessage: &plainTextMessageProperty{
//   			value: jsii.String("value"),
//   		},
//   		ssmlMessage: &sSMLMessageProperty{
//   			value: jsii.String("value"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	variations: []interface{}{
//   		&messageProperty{
//   			customPayload: &customPayloadProperty{
//   				value: jsii.String("value"),
//   			},
//   			imageResponseCard: &imageResponseCardProperty{
//   				title: jsii.String("title"),
//
//   				// the properties below are optional
//   				buttons: []interface{}{
//   					&buttonProperty{
//   						text: jsii.String("text"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				imageUrl: jsii.String("imageUrl"),
//   				subtitle: jsii.String("subtitle"),
//   			},
//   			plainTextMessage: &plainTextMessageProperty{
//   				value: jsii.String("value"),
//   			},
//   			ssmlMessage: &sSMLMessageProperty{
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnBot_MessageGroupProperty struct {
	// The primary message that Amazon Lex should send to the user.
	Message interface{} `field:"required" json:"message" yaml:"message"`
	// Message variations to send to the user.
	//
	// When variations are defined, Amazon Lex chooses the primary message or one of the variations to send to the user.
	Variations interface{} `field:"optional" json:"variations" yaml:"variations"`
}

