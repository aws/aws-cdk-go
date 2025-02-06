package awsnotifications


// Provides additional information about the current `NotificationHub` status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationHubStatusSummaryProperty := &NotificationHubStatusSummaryProperty{
//   	NotificationHubStatus: jsii.String("notificationHubStatus"),
//   	NotificationHubStatusReason: jsii.String("notificationHubStatusReason"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notifications-notificationhub-notificationhubstatussummary.html
//
type CfnNotificationHub_NotificationHubStatusSummaryProperty struct {
	// Indicates the current status of the `NotificationHub` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notifications-notificationhub-notificationhubstatussummary.html#cfn-notifications-notificationhub-notificationhubstatussummary-notificationhubstatus
	//
	NotificationHubStatus *string `field:"required" json:"notificationHubStatus" yaml:"notificationHubStatus"`
	// An explanation for the current status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notifications-notificationhub-notificationhubstatussummary.html#cfn-notifications-notificationhub-notificationhubstatussummary-notificationhubstatusreason
	//
	NotificationHubStatusReason *string `field:"required" json:"notificationHubStatusReason" yaml:"notificationHubStatusReason"`
}

