package awsconnect


// Information about the send notification action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sendNotificationActionProperty := &SendNotificationActionProperty{
//   	Content: jsii.String("content"),
//   	ContentType: jsii.String("contentType"),
//   	DeliveryMethod: jsii.String("deliveryMethod"),
//   	Recipient: &NotificationRecipientTypeProperty{
//   		UserArns: []*string{
//   			jsii.String("userArns"),
//   		},
//   		UserTags: map[string]*string{
//   			"userTagsKey": jsii.String("userTags"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Subject: jsii.String("subject"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html
//
type CfnRule_SendNotificationActionProperty struct {
	// Notification content.
	//
	// Supports variable injection. For more information, see [JSONPath reference](https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-variable-injection.html) in the *Amazon Connect Administrators Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html#cfn-connect-rule-sendnotificationaction-content
	//
	Content *string `field:"required" json:"content" yaml:"content"`
	// Content type format.
	//
	// *Allowed value* : `PLAIN_TEXT`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html#cfn-connect-rule-sendnotificationaction-contenttype
	//
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Notification delivery method.
	//
	// *Allowed value* : `EMAIL`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html#cfn-connect-rule-sendnotificationaction-deliverymethod
	//
	DeliveryMethod *string `field:"required" json:"deliveryMethod" yaml:"deliveryMethod"`
	// Notification recipient.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html#cfn-connect-rule-sendnotificationaction-recipient
	//
	Recipient interface{} `field:"required" json:"recipient" yaml:"recipient"`
	// The subject of the email if the delivery method is `EMAIL` .
	//
	// Supports variable injection. For more information, see [JSONPath reference](https://docs.aws.amazon.com/connect/latest/adminguide/contact-lens-variable-injection.html) in the *Amazon Connect Administrators Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-sendnotificationaction.html#cfn-connect-rule-sendnotificationaction-subject
	//
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

