package awsbudgets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbudgets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Budgets::Budget`.
//
// The `AWS::Budgets::Budget` resource allows customers to take pre-defined actions that will trigger once a budget threshold has been exceeded. creates, replaces, or deletes budgets for Billing and Cost Management. For more information, see [Managing Your Costs with Budgets](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/budgets-managing-costs.html) in the *AWS Billing and Cost Management User Guide* .
//
// TODO: EXAMPLE
//
type CfnBudget interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Budget() interface{}
	SetBudget(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	NotificationsWithSubscribers() interface{}
	SetNotificationsWithSubscribers(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBudget
type jsiiProxy_CfnBudget struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBudget) Budget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"budget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudget) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudget) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudget) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudget) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudget) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudget) NotificationsWithSubscribers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationsWithSubscribers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudget) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudget) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudget) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Budgets::Budget`.
func NewCfnBudget(scope constructs.Construct, id *string, props *CfnBudgetProps) CfnBudget {
	_init_.Initialize()

	j := jsiiProxy_CfnBudget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_budgets.CfnBudget",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Budgets::Budget`.
func NewCfnBudget_Override(c CfnBudget, scope constructs.Construct, id *string, props *CfnBudgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_budgets.CfnBudget",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBudget) SetBudget(val interface{}) {
	_jsii_.Set(
		j,
		"budget",
		val,
	)
}

func (j *jsiiProxy_CfnBudget) SetNotificationsWithSubscribers(val interface{}) {
	_jsii_.Set(
		j,
		"notificationsWithSubscribers",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnBudget_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_budgets.CfnBudget",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBudget_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_budgets.CfnBudget",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnBudget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_budgets.CfnBudget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBudget_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_budgets.CfnBudget",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnBudget) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnBudget) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnBudget) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnBudget) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnBudget) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnBudget) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnBudget) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnBudget) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnBudget) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnBudget) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnBudget) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBudget) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnBudget) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnBudget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBudget) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Represents the output of the `CreateBudget` operation.
//
// The content consists of the detailed metadata and data file information, and the current status of the `budget` object.
//
// This is the ARN pattern for a budget:
//
// `arn:aws:budgets::AccountId:budget/budgetName`
//
// TODO: EXAMPLE
//
type CfnBudget_BudgetDataProperty struct {
	// Whether this budget tracks costs, usage, RI utilization, RI coverage, Savings Plans utilization, or Savings Plans coverage.
	BudgetType *string `json:"budgetType"`
	// The length of time until a budget resets the actual and forecasted spend.
	//
	// `DAILY` is available only for `RI_UTILIZATION` and `RI_COVERAGE` budgets.
	TimeUnit *string `json:"timeUnit"`
	// The total amount of cost, usage, RI utilization, RI coverage, Savings Plans utilization, or Savings Plans coverage that you want to track with your budget.
	//
	// `BudgetLimit` is required for cost or usage budgets, but optional for RI or Savings Plans utilization or coverage budgets. RI and Savings Plans utilization or coverage budgets default to `100` , which is the only valid value for RI or Savings Plans utilization or coverage budgets. You can't use `BudgetLimit` with `PlannedBudgetLimits` for `CreateBudget` and `UpdateBudget` actions.
	BudgetLimit interface{} `json:"budgetLimit"`
	// The name of a budget.
	//
	// The value must be unique within an account. `BudgetName` can't include `:` and `\` characters. If you don't include value for `BudgetName` in the template, Billing and Cost Management assigns your budget a randomly generated name.
	BudgetName *string `json:"budgetName"`
	// The cost filters, such as `Region` , `Service` , `member account` , `Tag` , or `Cost Category` , that are applied to a budget.
	//
	// AWS Budgets supports the following services as a `Service` filter for RI budgets:
	//
	// - Amazon EC2
	// - Amazon Redshift
	// - Amazon Relational Database Service
	// - Amazon ElastiCache
	// - Amazon OpenSearch Service
	CostFilters interface{} `json:"costFilters"`
	// The types of costs that are included in this `COST` budget.
	//
	// `USAGE` , `RI_UTILIZATION` , `RI_COVERAGE` , `SAVINGS_PLANS_UTILIZATION` , and `SAVINGS_PLANS_COVERAGE` budgets do not have `CostTypes` .
	CostTypes interface{} `json:"costTypes"`
	// A map containing multiple `BudgetLimit` , including current or future limits.
	//
	// `PlannedBudgetLimits` is available for cost or usage budget and supports monthly and quarterly `TimeUnit` .
	//
	// For monthly budgets, provide 12 months of `PlannedBudgetLimits` values. This must start from the current month and include the next 11 months. The `key` is the start of the month, `UTC` in epoch seconds.
	//
	// For quarterly budgets, provide 4 quarters of `PlannedBudgetLimits` value entries in standard calendar quarter increments. This must start from the current quarter and include the next 3 quarters. The `key` is the start of the quarter, `UTC` in epoch seconds.
	//
	// If the planned budget expires before 12 months for monthly or 4 quarters for quarterly, provide the `PlannedBudgetLimits` values only for the remaining periods.
	//
	// If the budget begins at a date in the future, provide `PlannedBudgetLimits` values from the start date of the budget.
	//
	// After all of the `BudgetLimit` values in `PlannedBudgetLimits` are used, the budget continues to use the last limit as the `BudgetLimit` . At that point, the planned budget provides the same experience as a fixed budget.
	//
	// `DescribeBudget` and `DescribeBudgets` response along with `PlannedBudgetLimits` will also contain `BudgetLimit` representing the current month or quarter limit present in `PlannedBudgetLimits` . This only applies to budgets created with `PlannedBudgetLimits` . Budgets created without `PlannedBudgetLimits` will only contain `BudgetLimit` , and no `PlannedBudgetLimits` .
	PlannedBudgetLimits interface{} `json:"plannedBudgetLimits"`
	// The period of time that is covered by a budget.
	//
	// The period has a start date and an end date. The start date must come before the end date. There are no restrictions on the end date.
	//
	// The start date for a budget. If you created your budget and didn't specify a start date, the start date defaults to the start of the chosen time period (MONTHLY, QUARTERLY, or ANNUALLY). For example, if you create your budget on January 24, 2019, choose `MONTHLY` , and don't set a start date, the start date defaults to `01/01/19 00:00 UTC` . The defaults are the same for the AWS Billing and Cost Management console and the API.
	//
	// You can change your start date with the `UpdateBudget` operation.
	//
	// After the end date, AWS deletes the budget and all associated notifications and subscribers.
	TimePeriod interface{} `json:"timePeriod"`
}

// The types of cost that are included in a `COST` budget, such as tax and subscriptions.
//
// `USAGE` , `RI_UTILIZATION` , `RI_COVERAGE` , `SAVINGS_PLANS_UTILIZATION` , and `SAVINGS_PLANS_COVERAGE` budgets do not have `CostTypes` .
//
// TODO: EXAMPLE
//
type CfnBudget_CostTypesProperty struct {
	// Specifies whether a budget includes credits.
	//
	// The default value is `true` .
	IncludeCredit interface{} `json:"includeCredit"`
	// Specifies whether a budget includes discounts.
	//
	// The default value is `true` .
	IncludeDiscount interface{} `json:"includeDiscount"`
	// Specifies whether a budget includes non-RI subscription costs.
	//
	// The default value is `true` .
	IncludeOtherSubscription interface{} `json:"includeOtherSubscription"`
	// Specifies whether a budget includes recurring fees such as monthly RI fees.
	//
	// The default value is `true` .
	IncludeRecurring interface{} `json:"includeRecurring"`
	// Specifies whether a budget includes refunds.
	//
	// The default value is `true` .
	IncludeRefund interface{} `json:"includeRefund"`
	// Specifies whether a budget includes subscriptions.
	//
	// The default value is `true` .
	IncludeSubscription interface{} `json:"includeSubscription"`
	// Specifies whether a budget includes support subscription fees.
	//
	// The default value is `true` .
	IncludeSupport interface{} `json:"includeSupport"`
	// Specifies whether a budget includes taxes.
	//
	// The default value is `true` .
	IncludeTax interface{} `json:"includeTax"`
	// Specifies whether a budget includes upfront RI costs.
	//
	// The default value is `true` .
	IncludeUpfront interface{} `json:"includeUpfront"`
	// Specifies whether a budget uses the amortized rate.
	//
	// The default value is `false` .
	UseAmortized interface{} `json:"useAmortized"`
	// Specifies whether a budget uses a blended rate.
	//
	// The default value is `false` .
	UseBlended interface{} `json:"useBlended"`
}

// A notification that is associated with a budget. A budget can have up to ten notifications.
//
// Each notification must have at least one subscriber. A notification can have one SNS subscriber and up to 10 email subscribers, for a total of 11 subscribers.
//
// For example, if you have a budget for 200 dollars and you want to be notified when you go over 160 dollars, create a notification with the following parameters:
//
// - A notificationType of `ACTUAL`
// - A `thresholdType` of `PERCENTAGE`
// - A `comparisonOperator` of `GREATER_THAN`
// - A notification `threshold` of `80`
//
// TODO: EXAMPLE
//
type CfnBudget_NotificationProperty struct {
	// The comparison that is used for this notification.
	ComparisonOperator *string `json:"comparisonOperator"`
	// Whether the notification is for how much you have spent ( `ACTUAL` ) or for how much you're forecasted to spend ( `FORECASTED` ).
	NotificationType *string `json:"notificationType"`
	// The threshold that is associated with a notification.
	//
	// Thresholds are always a percentage, and many customers find value being alerted between 50% - 200% of the budgeted amount. The maximum limit for your threshold is 1,000,000% above the budgeted amount.
	Threshold *float64 `json:"threshold"`
	// The type of threshold for a notification.
	//
	// For `ABSOLUTE_VALUE` thresholds, AWS notifies you when you go over or are forecasted to go over your total cost threshold. For `PERCENTAGE` thresholds, AWS notifies you when you go over or are forecasted to go over a certain percentage of your forecasted spend. For example, if you have a budget for 200 dollars and you have a `PERCENTAGE` threshold of 80%, AWS notifies you when you go over 160 dollars.
	ThresholdType *string `json:"thresholdType"`
}

// A notification with subscribers.
//
// A notification can have one SNS subscriber and up to 10 email subscribers, for a total of 11 subscribers.
//
// TODO: EXAMPLE
//
type CfnBudget_NotificationWithSubscribersProperty struct {
	// The notification that is associated with a budget.
	Notification interface{} `json:"notification"`
	// A list of subscribers who are subscribed to this notification.
	Subscribers interface{} `json:"subscribers"`
}

// The amount of cost or usage that is measured for a budget.
//
// For example, a `Spend` for `3 GB` of S3 usage would have the following parameters:
//
// - An `Amount` of `3`
// - A `unit` of `GB`
//
// TODO: EXAMPLE
//
type CfnBudget_SpendProperty struct {
	// The cost or usage amount that is associated with a budget forecast, actual spend, or budget threshold.
	Amount *float64 `json:"amount"`
	// The unit of measurement that is used for the budget forecast, actual spend, or budget threshold, such as USD or GB.
	Unit *string `json:"unit"`
}

// The `Subscriber` property type specifies who to notify for a Billing and Cost Management budget notification.
//
// The subscriber consists of a subscription type, and either an Amazon SNS topic or an email address.
//
// For example, an email subscriber would have the following parameters:
//
// - A `subscriptionType` of `EMAIL`
// - An `address` of `example@example.com`
//
// TODO: EXAMPLE
//
type CfnBudget_SubscriberProperty struct {
	// The address that AWS sends budget notifications to, either an SNS topic or an email.
	//
	// When you create a subscriber, the value of `Address` can't contain line breaks.
	Address *string `json:"address"`
	// The type of notification that AWS sends to a subscriber.
	SubscriptionType *string `json:"subscriptionType"`
}

// The period of time that is covered by a budget.
//
// The period has a start date and an end date. The start date must come before the end date. There are no restrictions on the end date.
//
// TODO: EXAMPLE
//
type CfnBudget_TimePeriodProperty struct {
	// The end date for a budget.
	//
	// If you didn't specify an end date, AWS set your end date to `06/15/87 00:00 UTC` . The defaults are the same for the AWS Billing and Cost Management console and the API.
	//
	// After the end date, AWS deletes the budget and all associated notifications and subscribers. You can change your end date with the `UpdateBudget` operation.
	End *string `json:"end"`
	// The start date for a budget.
	//
	// If you created your budget and didn't specify a start date, the start date defaults to the start of the chosen time period (MONTHLY, QUARTERLY, or ANNUALLY). For example, if you create your budget on January 24, 2019, choose `MONTHLY` , and don't set a start date, the start date defaults to `01/01/19 00:00 UTC` . The defaults are the same for the AWS Billing and Cost Management console and the API.
	//
	// You can change your start date with the `UpdateBudget` operation.
	//
	// Valid values depend on the value of `BudgetType` :
	//
	// - If `BudgetType` is `COST` or `USAGE` : Valid values are `MONTHLY` , `QUARTERLY` , and `ANNUALLY` .
	// - If `BudgetType` is `RI_UTILIZATION` or `RI_COVERAGE` : Valid values are `DAILY` , `MONTHLY` , `QUARTERLY` , and `ANNUALLY` .
	Start *string `json:"start"`
}

// Properties for defining a `CfnBudget`.
//
// TODO: EXAMPLE
//
type CfnBudgetProps struct {
	// The budget object that you want to create.
	Budget interface{} `json:"budget"`
	// A notification that you want to associate with a budget.
	//
	// A budget can have up to five notifications, and each notification can have one SNS subscriber and up to 10 email subscribers. If you include notifications and subscribers in your `CreateBudget` call, AWS creates the notifications and subscribers for you.
	NotificationsWithSubscribers interface{} `json:"notificationsWithSubscribers"`
}

// A CloudFormation `AWS::Budgets::BudgetsAction`.
//
// The `AWS::Budgets::BudgetsAction` resource enables you to take predefined actions that are initiated when a budget threshold has been exceeded. For more information, see [Managing Your Costs with Budgets](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/budgets-managing-costs.html) in the *AWS Billing and Cost Management User Guide* .
//
// TODO: EXAMPLE
//
type CfnBudgetsAction interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ActionThreshold() interface{}
	SetActionThreshold(val interface{})
	ActionType() *string
	SetActionType(val *string)
	ApprovalModel() *string
	SetApprovalModel(val *string)
	AttrActionId() *string
	BudgetName() *string
	SetBudgetName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Definition() interface{}
	SetDefinition(val interface{})
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	NotificationType() *string
	SetNotificationType(val *string)
	Ref() *string
	Stack() awscdk.Stack
	Subscribers() interface{}
	SetSubscribers(val interface{})
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBudgetsAction
type jsiiProxy_CfnBudgetsAction struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBudgetsAction) ActionThreshold() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) ActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) ApprovalModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) AttrActionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrActionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) BudgetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) Definition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) NotificationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) Subscribers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscribers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsAction) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Budgets::BudgetsAction`.
func NewCfnBudgetsAction(scope constructs.Construct, id *string, props *CfnBudgetsActionProps) CfnBudgetsAction {
	_init_.Initialize()

	j := jsiiProxy_CfnBudgetsAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_budgets.CfnBudgetsAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Budgets::BudgetsAction`.
func NewCfnBudgetsAction_Override(c CfnBudgetsAction, scope constructs.Construct, id *string, props *CfnBudgetsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_budgets.CfnBudgetsAction",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBudgetsAction) SetActionThreshold(val interface{}) {
	_jsii_.Set(
		j,
		"actionThreshold",
		val,
	)
}

func (j *jsiiProxy_CfnBudgetsAction) SetActionType(val *string) {
	_jsii_.Set(
		j,
		"actionType",
		val,
	)
}

func (j *jsiiProxy_CfnBudgetsAction) SetApprovalModel(val *string) {
	_jsii_.Set(
		j,
		"approvalModel",
		val,
	)
}

func (j *jsiiProxy_CfnBudgetsAction) SetBudgetName(val *string) {
	_jsii_.Set(
		j,
		"budgetName",
		val,
	)
}

func (j *jsiiProxy_CfnBudgetsAction) SetDefinition(val interface{}) {
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

func (j *jsiiProxy_CfnBudgetsAction) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnBudgetsAction) SetNotificationType(val *string) {
	_jsii_.Set(
		j,
		"notificationType",
		val,
	)
}

func (j *jsiiProxy_CfnBudgetsAction) SetSubscribers(val interface{}) {
	_jsii_.Set(
		j,
		"subscribers",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnBudgetsAction_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_budgets.CfnBudgetsAction",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBudgetsAction_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_budgets.CfnBudgetsAction",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnBudgetsAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_budgets.CfnBudgetsAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBudgetsAction_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_budgets.CfnBudgetsAction",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnBudgetsAction) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnBudgetsAction) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnBudgetsAction) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnBudgetsAction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnBudgetsAction) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnBudgetsAction) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnBudgetsAction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnBudgetsAction) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnBudgetsAction) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnBudgetsAction) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnBudgetsAction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBudgetsAction) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnBudgetsAction) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnBudgetsAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBudgetsAction) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The trigger threshold of the action.
//
// TODO: EXAMPLE
//
type CfnBudgetsAction_ActionThresholdProperty struct {
	// The type of threshold for a notification.
	Type *string `json:"type"`
	// The threshold of a notification.
	Value *float64 `json:"value"`
}

// The definition is where you specify all of the type-specific parameters.
//
// TODO: EXAMPLE
//
type CfnBudgetsAction_DefinitionProperty struct {
	// The AWS Identity and Access Management ( IAM ) action definition details.
	IamActionDefinition interface{} `json:"iamActionDefinition"`
	// The service control policies (SCP) action definition details.
	ScpActionDefinition interface{} `json:"scpActionDefinition"`
	// The Amazon EC2 Systems Manager ( SSM ) action definition details.
	SsmActionDefinition interface{} `json:"ssmActionDefinition"`
}

// The AWS Identity and Access Management ( IAM ) action definition details.
//
// TODO: EXAMPLE
//
type CfnBudgetsAction_IamActionDefinitionProperty struct {
	// The Amazon Resource Name (ARN) of the policy to be attached.
	PolicyArn *string `json:"policyArn"`
	// A list of groups to be attached.
	//
	// There must be at least one group.
	Groups *[]*string `json:"groups"`
	// A list of roles to be attached.
	//
	// There must be at least one role.
	Roles *[]*string `json:"roles"`
	// A list of users to be attached.
	//
	// There must be at least one user.
	Users *[]*string `json:"users"`
}

// The service control policies (SCP) action definition details.
//
// TODO: EXAMPLE
//
type CfnBudgetsAction_ScpActionDefinitionProperty struct {
	// The policy ID attached.
	PolicyId *string `json:"policyId"`
	// A list of target IDs.
	TargetIds *[]*string `json:"targetIds"`
}

// The Amazon EC2 Systems Manager ( SSM ) action definition details.
//
// TODO: EXAMPLE
//
type CfnBudgetsAction_SsmActionDefinitionProperty struct {
	// The EC2 and RDS instance IDs.
	InstanceIds *[]*string `json:"instanceIds"`
	// The Region to run the ( SSM ) document.
	Region *string `json:"region"`
	// The action subType.
	Subtype *string `json:"subtype"`
}

// The subscriber to a budget notification.
//
// The subscriber consists of a subscription type and either an Amazon SNS topic or an email address.
//
// For example, an email subscriber would have the following parameters:
//
// - A `subscriptionType` of `EMAIL`
// - An `address` of `example@example.com`
//
// TODO: EXAMPLE
//
type CfnBudgetsAction_SubscriberProperty struct {
	// The address that AWS sends budget notifications to, either an SNS topic or an email.
	//
	// When you create a subscriber, the value of `Address` can't contain line breaks.
	Address *string `json:"address"`
	// The type of notification that AWS sends to a subscriber.
	Type *string `json:"type"`
}

// Properties for defining a `CfnBudgetsAction`.
//
// TODO: EXAMPLE
//
type CfnBudgetsActionProps struct {
	// The trigger threshold of the action.
	ActionThreshold interface{} `json:"actionThreshold"`
	// The type of action.
	//
	// This defines the type of tasks that can be carried out by this action. This field also determines the format for definition.
	ActionType *string `json:"actionType"`
	// A string that represents the budget name.
	//
	// ":" and "\" characters aren't allowed.
	BudgetName *string `json:"budgetName"`
	// Specifies all of the type-specific parameters.
	Definition interface{} `json:"definition"`
	// The role passed for action execution and reversion.
	//
	// Roles and actions must be in the same account.
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// The type of a notification.
	NotificationType *string `json:"notificationType"`
	// A list of subscribers.
	Subscribers interface{} `json:"subscribers"`
	// This specifies if the action needs manual or automatic approval.
	ApprovalModel *string `json:"approvalModel"`
}

