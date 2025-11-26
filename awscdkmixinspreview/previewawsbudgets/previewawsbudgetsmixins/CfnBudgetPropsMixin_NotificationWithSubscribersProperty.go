package previewawsbudgetsmixins


// A notification with subscribers.
//
// A notification can have one SNS subscriber and up to 10 email subscribers, for a total of 11 subscribers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   notificationWithSubscribersProperty := &NotificationWithSubscribersProperty{
//   	Notification: &NotificationProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		NotificationType: jsii.String("notificationType"),
//   		Threshold: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html
//
type CfnBudgetPropsMixin_NotificationWithSubscribersProperty struct {
	// The notification that's associated with a budget.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html#cfn-budgets-budget-notificationwithsubscribers-notification
	//
	Notification interface{} `field:"optional" json:"notification" yaml:"notification"`
	// A list of subscribers who are subscribed to this notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html#cfn-budgets-budget-notificationwithsubscribers-subscribers
	//
	Subscribers interface{} `field:"optional" json:"subscribers" yaml:"subscribers"`
}

