package awscodestarnotifications


// Information about the SNS topic or AWS Chatbot client associated with a notification target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationRuleTargetConfig := &notificationRuleTargetConfig{
//   	targetAddress: jsii.String("targetAddress"),
//   	targetType: jsii.String("targetType"),
//   }
//
type NotificationRuleTargetConfig struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic or AWS Chatbot client.
	TargetAddress *string `field:"required" json:"targetAddress" yaml:"targetAddress"`
	// The target type.
	//
	// Can be an Amazon SNS topic or AWS Chatbot client.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

