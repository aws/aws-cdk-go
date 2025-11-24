package mixinsawschatbot


// > AWS Chatbot is now  . [Learn more](https://docs.aws.amazon.com//chatbot/latest/adminguide/service-rename.html) >  > `Type` attribute values remain unchanged.
//
// Defines when a custom action button should be attached to a notification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customActionAttachmentProperty := &CustomActionAttachmentProperty{
//   	ButtonText: jsii.String("buttonText"),
//   	Criteria: []interface{}{
//   		&CustomActionAttachmentCriteriaProperty{
//   			Operator: jsii.String("operator"),
//   			Value: jsii.String("value"),
//   			VariableName: jsii.String("variableName"),
//   		},
//   	},
//   	NotificationType: jsii.String("notificationType"),
//   	Variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachment.html
//
type CfnCustomActionPropsMixin_CustomActionAttachmentProperty struct {
	// The text of the button that appears on the notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachment.html#cfn-chatbot-customaction-customactionattachment-buttontext
	//
	ButtonText *string `field:"optional" json:"buttonText" yaml:"buttonText"`
	// The criteria for when a button should be shown based on values in the notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachment.html#cfn-chatbot-customaction-customactionattachment-criteria
	//
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// The type of notification that the custom action should be attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachment.html#cfn-chatbot-customaction-customactionattachment-notificationtype
	//
	NotificationType *string `field:"optional" json:"notificationType" yaml:"notificationType"`
	// The variables to extract from the notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachment.html#cfn-chatbot-customaction-customactionattachment-variables
	//
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

