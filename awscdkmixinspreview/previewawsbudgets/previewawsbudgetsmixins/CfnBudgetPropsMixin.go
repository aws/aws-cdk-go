package previewawsbudgetsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbudgets/previewawsbudgetsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Budgets::Budget` resource allows customers to take pre-defined actions that will trigger once a budget threshold has been exceeded.
//
// creates, replaces, or deletes budgets for Billing and Cost Management. For more information, see [Managing Your Costs with Budgets](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/budgets-managing-costs.html) in the *Billing and Cost Management User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var costFilters interface{}
//   var expressionProperty_ ExpressionProperty
//   var plannedBudgetLimits interface{}
//
//   cfnBudgetPropsMixin := awscdkmixinspreview.Mixins.NewCfnBudgetPropsMixin(&CfnBudgetMixinProps{
//   	Budget: &BudgetDataProperty{
//   		AutoAdjustData: &AutoAdjustDataProperty{
//   			AutoAdjustType: jsii.String("autoAdjustType"),
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
//   		BudgetType: jsii.String("budgetType"),
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
//   		TimeUnit: jsii.String("timeUnit"),
//   	},
//   	NotificationsWithSubscribers: []interface{}{
//   		&NotificationWithSubscribersProperty{
//   			Notification: &NotificationProperty{
//   				ComparisonOperator: jsii.String("comparisonOperator"),
//   				NotificationType: jsii.String("notificationType"),
//   				Threshold: jsii.Number(123),
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
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html
//
type CfnBudgetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBudgetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBudgetPropsMixin
type jsiiProxy_CfnBudgetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBudgetPropsMixin) Props() *CfnBudgetMixinProps {
	var returns *CfnBudgetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Budgets::Budget`.
func NewCfnBudgetPropsMixin(props *CfnBudgetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBudgetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBudgetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBudgetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Budgets::Budget`.
func NewCfnBudgetPropsMixin_Override(c CfnBudgetPropsMixin, props *CfnBudgetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBudgetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBudgetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBudgetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBudgetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnBudgetPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

