package awsconnect


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
	// `CfnRule.SendNotificationActionProperty.Content`.
	Content *string `field:"required" json:"content" yaml:"content"`
	// `CfnRule.SendNotificationActionProperty.ContentType`.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// `CfnRule.SendNotificationActionProperty.DeliveryMethod`.
	DeliveryMethod *string `field:"required" json:"deliveryMethod" yaml:"deliveryMethod"`
	// `CfnRule.SendNotificationActionProperty.Recipient`.
	Recipient interface{} `field:"required" json:"recipient" yaml:"recipient"`
	// `CfnRule.SendNotificationActionProperty.Subject`.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

