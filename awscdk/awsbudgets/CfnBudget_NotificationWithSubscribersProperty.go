package awsbudgets


// A notification with subscribers.
//
// A notification can have one SNS subscriber and up to 10 email subscribers, for a total of 11 subscribers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationWithSubscribersProperty := &NotificationWithSubscribersProperty{
//   	Notification: &NotificationProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		NotificationType: jsii.String("notificationType"),
//   		Threshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		ThresholdType: jsii.String("thresholdType"),
//   	},
//   	Subscribers: []interface{}{
//   		&SubscriberProperty{
//   			Address: jsii.String("address"),
//   			SubscriptionType: jsii.String("subscriptionType"),
//   		},
//   	},
//   }
//
type CfnBudget_NotificationWithSubscribersProperty struct {
	// The notification that's associated with a budget.
	Notification interface{} `field:"required" json:"notification" yaml:"notification"`
	// A list of subscribers who are subscribed to this notification.
	Subscribers interface{} `field:"required" json:"subscribers" yaml:"subscribers"`
}

