package awscodestarnotifications


// Standard set of options for `notifyOnXxx` codestar notification handler on construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationRuleOptions := &NotificationRuleOptions{
//   	DetailType: awscdk.Aws_codestarnotifications.DetailType_BASIC,
//   	Enabled: jsii.Boolean(false),
//   	NotificationRuleName: jsii.String("notificationRuleName"),
//   }
//
type NotificationRuleOptions struct {
	// The level of detail to include in the notifications for this resource.
	//
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	// Default: DetailType.FULL
	//
	DetailType DetailType `field:"optional" json:"detailType" yaml:"detailType"`
	// The status of the notification rule.
	//
	// If the enabled is set to DISABLED, notifications aren't sent for the notification rule.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account.
	// Default: - generated from the `id`.
	//
	NotificationRuleName *string `field:"optional" json:"notificationRuleName" yaml:"notificationRuleName"`
}

