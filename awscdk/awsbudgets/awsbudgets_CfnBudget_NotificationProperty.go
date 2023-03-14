package awsbudgets


// A notification that's associated with a budget. A budget can have up to ten notifications.
//
// Each notification must have at least one subscriber. A notification can have one SNS subscriber and up to 10 email subscribers, for a total of 11 subscribers.
//
// For example, if you have a budget for 200 dollars and you want to be notified when you go over 160 dollars, create a notification with the following parameters:
//
// - A notificationType of `ACTUAL`
// - A `thresholdType` of `PERCENTAGE`
// - A `comparisonOperator` of `GREATER_THAN`
// - A notification `threshold` of `80`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationProperty := &notificationProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	notificationType: jsii.String("notificationType"),
//   	threshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	thresholdType: jsii.String("thresholdType"),
//   }
//
type CfnBudget_NotificationProperty struct {
	// The comparison that's used for this notification.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Specifies whether the notification is for how much you have spent ( `ACTUAL` ) or for how much that you're forecasted to spend ( `FORECASTED` ).
	NotificationType *string `field:"required" json:"notificationType" yaml:"notificationType"`
	// The threshold that's associated with a notification.
	//
	// Thresholds are always a percentage, and many customers find value being alerted between 50% - 200% of the budgeted amount. The maximum limit for your threshold is 1,000,000% above the budgeted amount.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// The type of threshold for a notification.
	//
	// For `ABSOLUTE_VALUE` thresholds, AWS notifies you when you go over or are forecasted to go over your total cost threshold. For `PERCENTAGE` thresholds, AWS notifies you when you go over or are forecasted to go over a certain percentage of your forecasted spend. For example, if you have a budget for 200 dollars and you have a `PERCENTAGE` threshold of 80%, AWS notifies you when you go over 160 dollars.
	ThresholdType *string `field:"optional" json:"thresholdType" yaml:"thresholdType"`
}

