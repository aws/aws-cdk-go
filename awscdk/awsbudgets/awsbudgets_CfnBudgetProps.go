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
//   cfnBudgetProps := &cfnBudgetProps{
//   	budget: &budgetDataProperty{
//   		budgetType: jsii.String("budgetType"),
//   		timeUnit: jsii.String("timeUnit"),
//
//   		// the properties below are optional
//   		autoAdjustData: &autoAdjustDataProperty{
//   			autoAdjustType: jsii.String("autoAdjustType"),
//
//   			// the properties below are optional
//   			historicalOptions: &historicalOptionsProperty{
//   				budgetAdjustmentPeriod: jsii.Number(123),
//   			},
//   		},
//   		budgetLimit: &spendProperty{
//   			amount: jsii.Number(123),
//   			unit: jsii.String("unit"),
//   		},
//   		budgetName: jsii.String("budgetName"),
//   		costFilters: costFilters,
//   		costTypes: &costTypesProperty{
//   			includeCredit: jsii.Boolean(false),
//   			includeDiscount: jsii.Boolean(false),
//   			includeOtherSubscription: jsii.Boolean(false),
//   			includeRecurring: jsii.Boolean(false),
//   			includeRefund: jsii.Boolean(false),
//   			includeSubscription: jsii.Boolean(false),
//   			includeSupport: jsii.Boolean(false),
//   			includeTax: jsii.Boolean(false),
//   			includeUpfront: jsii.Boolean(false),
//   			useAmortized: jsii.Boolean(false),
//   			useBlended: jsii.Boolean(false),
//   		},
//   		plannedBudgetLimits: plannedBudgetLimits,
//   		timePeriod: &timePeriodProperty{
//   			end: jsii.String("end"),
//   			start: jsii.String("start"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	notificationsWithSubscribers: []interface{}{
//   		&notificationWithSubscribersProperty{
//   			notification: &notificationProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				notificationType: jsii.String("notificationType"),
//   				threshold: jsii.Number(123),
//
//   				// the properties below are optional
//   				thresholdType: jsii.String("thresholdType"),
//   			},
//   			subscribers: []interface{}{
//   				&subscriberProperty{
//   					address: jsii.String("address"),
//   					subscriptionType: jsii.String("subscriptionType"),
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

