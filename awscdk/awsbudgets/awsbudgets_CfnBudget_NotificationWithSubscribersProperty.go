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
//   notificationWithSubscribersProperty := &notificationWithSubscribersProperty{
//   	notification: &notificationProperty{
//   		comparisonOperator: jsii.String("comparisonOperator"),
//   		notificationType: jsii.String("notificationType"),
//   		threshold: jsii.Number(123),
//
//   		// the properties below are optional
//   		thresholdType: jsii.String("thresholdType"),
//   	},
//   	subscribers: []interface{}{
//   		&subscriberProperty{
//   			address: jsii.String("address"),
//   			subscriptionType: jsii.String("subscriptionType"),
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

