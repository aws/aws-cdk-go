package awsconnect


// Information about the send notification action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sendNotificationActionProperty := &sendNotificationActionProperty{
//   	content: jsii.String("content"),
//   	contentType: jsii.String("contentType"),
//   	deliveryMethod: jsii.String("deliveryMethod"),
//   	recipient: &notificationRecipientTypeProperty{
//   		userArns: []*string{
//   			jsii.String("userArns"),
//   		},
//   		userTags: map[string]*string{
//   			"userTagsKey": jsii.String("userTags"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	subject: jsii.String("subject"),
//   }
//
type CfnRule_SendNotificationActionProperty struct {
	// Notification content.
	//
	// Supports variable injection. For more information, see [JSONPath reference](https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-variable-injection.html) in the *Amazon Connect Administrators Guide* .
	Content *string `field:"required" json:"content" yaml:"content"`
	// Content type format.
	//
	// *Allowed value* : `PLAIN_TEXT`.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Notification delivery method.
	//
	// *Allowed value* : `EMAIL`.
	DeliveryMethod *string `field:"required" json:"deliveryMethod" yaml:"deliveryMethod"`
	// Notification recipient.
	Recipient interface{} `field:"required" json:"recipient" yaml:"recipient"`
	// The subject of the email if the delivery method is `EMAIL` .
	//
	// Supports variable injection. For more information, see [JSONPath reference](https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-variable-injection.html) in the *Amazon Connect Administrators Guide* .
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

