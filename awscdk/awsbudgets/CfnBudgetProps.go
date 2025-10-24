package awsbudgets


// Properties for defining a `CfnBudget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var costFilters interface{}
//   var expressionProperty_ ExpressionProperty
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
//   		BillingViewArn: jsii.String("billingViewArn"),
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
//   		FilterExpression: &ExpressionProperty{
//   			And: []interface{}{
//   				expressionProperty_,
//   			},
//   			CostCategories: &CostCategoryValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Dimensions: &ExpressionDimensionValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Not: expressionProperty_,
//   			Or: []interface{}{
//   				expressionProperty_,
//   			},
//   			Tags: &TagValuesProperty{
//   				Key: jsii.String("key"),
//   				MatchOptions: []*string{
//   					jsii.String("matchOptions"),
//   				},
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		Metrics: []*string{
//   			jsii.String("metrics"),
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
//   	ResourceTags: []interface{}{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//
//   			// the properties below are optional
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html
//
type CfnBudgetProps struct {
	// The budget object that you want to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-budget
	//
	Budget interface{} `field:"required" json:"budget" yaml:"budget"`
	// A notification that you want to associate with a budget.
	//
	// A budget can have up to five notifications, and each notification can have one SNS subscriber and up to 10 email subscribers. If you include notifications and subscribers in your `CreateBudget` call, AWS creates the notifications and subscribers for you.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-notificationswithsubscribers
	//
	NotificationsWithSubscribers interface{} `field:"optional" json:"notificationsWithSubscribers" yaml:"notificationsWithSubscribers"`
	// An optional list of tags to associate with the specified budget.
	//
	// Each tag consists of a key and a value, and each key must be unique for the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

