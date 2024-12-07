package awschatbot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionAttachmentProperty := &CustomActionAttachmentProperty{
//   	ButtonText: jsii.String("buttonText"),
//   	Criteria: []interface{}{
//   		&CustomActionAttachmentCriteriaProperty{
//   			Operator: jsii.String("operator"),
//   			VariableName: jsii.String("variableName"),
//
//   			// the properties below are optional
//   			Value: jsii.String("value"),
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
type CfnCustomAction_CustomActionAttachmentProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachment.html#cfn-chatbot-customaction-customactionattachment-buttontext
	//
	ButtonText *string `field:"optional" json:"buttonText" yaml:"buttonText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachment.html#cfn-chatbot-customaction-customactionattachment-criteria
	//
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachment.html#cfn-chatbot-customaction-customactionattachment-notificationtype
	//
	NotificationType *string `field:"optional" json:"notificationType" yaml:"notificationType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachment.html#cfn-chatbot-customaction-customactionattachment-variables
	//
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

