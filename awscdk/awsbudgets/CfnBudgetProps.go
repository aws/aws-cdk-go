package awsbudgets


// Properties for defining a `CfnBudget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var costFilters interface{}
//   var plannedBudgetLimits interface{}
//
//   cfnBudgetProps := &CfnBudgetProps{
//   	Budget: &BudgetDataProperty{
//   		BudgetType: jsii.String("budgetType"),
//   		TimeUnit: jsii.String("timeUnit"),
//
//   		// the properties below are optional
//   		AutoAdjustData: &AutoAdjustDataProperty{
//   			AutoAdjustType: jsii.String("autoAdjustType"),
//
//   			// the properties below are optional
//   			HistoricalOptions: &HistoricalOptionsProperty{
//   				BudgetAdjustmentPeriod: jsii.Number(123),
//   			},
//   		},
//   		BudgetLimit: &SpendProperty{
//   			Amount: jsii.Number(123),
//   			Unit: jsii.String("unit"),
//   		},
//   		BudgetName: jsii.String("budgetName"),
//   		CostFilters: costFilters,
//   		CostTypes: &CostTypesProperty{
//   			IncludeCredit: jsii.Boolean(false),
//   			IncludeDiscount: jsii.Boolean(false),
//   			IncludeOtherSubscription: jsii.Boolean(false),
//   			IncludeRecurring: jsii.Boolean(false),
//   			IncludeRefund: jsii.Boolean(false),
//   			IncludeSubscription: jsii.Boolean(false),
//   			IncludeSupport: jsii.Boolean(false),
//   			IncludeTax: jsii.Boolean(false),
//   			IncludeUpfront: jsii.Boolean(false),
//   			UseAmortized: jsii.Boolean(false),
//   			UseBlended: jsii.Boolean(false),
//   		},
//   		PlannedBudgetLimits: plannedBudgetLimits,
//   		TimePeriod: &TimePeriodProperty{
//   			End: jsii.String("end"),
//   			Start: jsii.String("start"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	NotificationsWithSubscribers: []interface{}{
//   		&NotificationWithSubscribersProperty{
//   			Notification: &NotificationProperty{
//   				ComparisonOperator: jsii.String("comparisonOperator"),
//   				NotificationType: jsii.String("notificationType"),
//   				Threshold: jsii.Number(123),
//
//   				// the properties below are optional
//   				ThresholdType: jsii.String("thresholdType"),
//   			},
//   			Subscribers: []interface{}{
//   				&SubscriberProperty{
//   					Address: jsii.String("address"),
//   					SubscriptionType: jsii.String("subscriptionType"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBudgetProps struct {
	// The budget object that you want to create.
	Budget interface{} `field:"required" json:"budget" yaml:"budget"`
	// A notification that you want to associate with a budget.
	//
	// A budget can have up to five notifications, and each notification can have one SNS subscriber and up to 10 email subscribers. If you include notifications and subscribers in your `CreateBudget` call, AWS creates the notifications and subscribers for you.
	NotificationsWithSubscribers interface{} `field:"optional" json:"notificationsWithSubscribers" yaml:"notificationsWithSubscribers"`
}

