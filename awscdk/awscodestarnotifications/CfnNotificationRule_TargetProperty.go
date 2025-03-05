package awscodestarnotifications


// Information about the AWS Chatbot topics or AWS Chatbot clients associated with a notification rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetProperty := &TargetProperty{
//   	TargetAddress: jsii.String("targetAddress"),
//   	TargetType: jsii.String("targetType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codestarnotifications-notificationrule-target.html
//
type CfnNotificationRule_TargetProperty struct {
	// The Amazon Resource Name (ARN) of the AWS Chatbot topic or AWS Chatbot client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codestarnotifications-notificationrule-target.html#cfn-codestarnotifications-notificationrule-target-targetaddress
	//
	TargetAddress *string `field:"required" json:"targetAddress" yaml:"targetAddress"`
	// The target type. Can be an Amazon Simple Notification Service topic or AWS Chatbot client.
	//
	// - Amazon Simple Notification Service topics are specified as `SNS` .
	// - AWS Chatbot clients are specified as `AWSChatbotSlack` .
	// - AWS Chatbot clients for Microsoft Teams are specified as `AWSChatbotMicrosoftTeams` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codestarnotifications-notificationrule-target.html#cfn-codestarnotifications-notificationrule-target-targettype
	//
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

