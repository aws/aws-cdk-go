package awscodestarnotifications


// Information about the  topics or  clients associated with a notification rule.
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
	// The Amazon Resource Name (ARN) of the  topic or  client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codestarnotifications-notificationrule-target.html#cfn-codestarnotifications-notificationrule-target-targetaddress
	//
	TargetAddress *string `field:"required" json:"targetAddress" yaml:"targetAddress"`
	// The target type. Can be an Amazon Simple Notification Service topic or  client.
	//
	// - Amazon Simple Notification Service topics are specified as `SNS` .
	// - clients are specified as `AWSChatbotSlack` .
	// - clients for Microsoft Teams are specified as `AWSChatbotMicrosoftTeams` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codestarnotifications-notificationrule-target.html#cfn-codestarnotifications-notificationrule-target-targettype
	//
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

