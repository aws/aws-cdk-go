package awsrolesanywhere


// Customizable notification settings that will be applied to notification events.
//
// IAM Roles Anywhere consumes these settings while notifying across multiple channels - CloudWatch metrics, EventBridge, and Health Dashboard .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationSettingProperty := &NotificationSettingProperty{
//   	Enabled: jsii.Boolean(false),
//   	Event: jsii.String("event"),
//
//   	// the properties below are optional
//   	Channel: jsii.String("channel"),
//   	Threshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-trustanchor-notificationsetting.html
//
type CfnTrustAnchor_NotificationSettingProperty struct {
	// Indicates whether the notification setting is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-trustanchor-notificationsetting.html#cfn-rolesanywhere-trustanchor-notificationsetting-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The event to which this notification setting is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-trustanchor-notificationsetting.html#cfn-rolesanywhere-trustanchor-notificationsetting-event
	//
	Event *string `field:"required" json:"event" yaml:"event"`
	// The specified channel of notification.
	//
	// IAM Roles Anywhere uses CloudWatch metrics, EventBridge, and Health Dashboard to notify for an event.
	//
	// > In the absence of a specific channel, IAM Roles Anywhere applies this setting to 'ALL' channels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-trustanchor-notificationsetting.html#cfn-rolesanywhere-trustanchor-notificationsetting-channel
	//
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// The number of days before a notification event.
	//
	// This value is required for a notification setting that is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rolesanywhere-trustanchor-notificationsetting.html#cfn-rolesanywhere-trustanchor-notificationsetting-threshold
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

