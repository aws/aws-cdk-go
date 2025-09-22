package awscodestarnotifications


// A reference to a NotificationRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationRuleReference := &NotificationRuleReference{
//   	NotificationRuleArn: jsii.String("notificationRuleArn"),
//   }
//
type NotificationRuleReference struct {
	// The Arn of the NotificationRule resource.
	NotificationRuleArn *string `field:"required" json:"notificationRuleArn" yaml:"notificationRuleArn"`
}

