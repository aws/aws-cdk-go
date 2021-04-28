package awsbudgets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbudgets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Budgets::Budget`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnBudget) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnBudget) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnBudget) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnBudget_BudgetDataProperty struct {
	// `CfnBudget.BudgetDataProperty.BudgetType`.
	BudgetType *string `json:"budgetType"`
	// `CfnBudget.BudgetDataProperty.TimeUnit`.
	TimeUnit *string `json:"timeUnit"`
	// `CfnBudget.BudgetDataProperty.BudgetLimit`.
	BudgetLimit interface{} `json:"budgetLimit"`
	// `CfnBudget.BudgetDataProperty.BudgetName`.
	BudgetName *string `json:"budgetName"`
	// `CfnBudget.BudgetDataProperty.CostFilters`.
	CostFilters interface{} `json:"costFilters"`
	// `CfnBudget.BudgetDataProperty.CostTypes`.
	CostTypes interface{} `json:"costTypes"`
	// `CfnBudget.BudgetDataProperty.PlannedBudgetLimits`.
	PlannedBudgetLimits interface{} `json:"plannedBudgetLimits"`
	// `CfnBudget.BudgetDataProperty.TimePeriod`.
	TimePeriod interface{} `json:"timePeriod"`
}

type CfnBudget_CostTypesProperty struct {
	// `CfnBudget.CostTypesProperty.IncludeCredit`.
	IncludeCredit interface{} `json:"includeCredit"`
	// `CfnBudget.CostTypesProperty.IncludeDiscount`.
	IncludeDiscount interface{} `json:"includeDiscount"`
	// `CfnBudget.CostTypesProperty.IncludeOtherSubscription`.
	IncludeOtherSubscription interface{} `json:"includeOtherSubscription"`
	// `CfnBudget.CostTypesProperty.IncludeRecurring`.
	IncludeRecurring interface{} `json:"includeRecurring"`
	// `CfnBudget.CostTypesProperty.IncludeRefund`.
	IncludeRefund interface{} `json:"includeRefund"`
	// `CfnBudget.CostTypesProperty.IncludeSubscription`.
	IncludeSubscription interface{} `json:"includeSubscription"`
	// `CfnBudget.CostTypesProperty.IncludeSupport`.
	IncludeSupport interface{} `json:"includeSupport"`
	// `CfnBudget.CostTypesProperty.IncludeTax`.
	IncludeTax interface{} `json:"includeTax"`
	// `CfnBudget.CostTypesProperty.IncludeUpfront`.
	IncludeUpfront interface{} `json:"includeUpfront"`
	// `CfnBudget.CostTypesProperty.UseAmortized`.
	UseAmortized interface{} `json:"useAmortized"`
	// `CfnBudget.CostTypesProperty.UseBlended`.
	UseBlended interface{} `json:"useBlended"`
}

type CfnBudget_NotificationProperty struct {
	// `CfnBudget.NotificationProperty.ComparisonOperator`.
	ComparisonOperator *string `json:"comparisonOperator"`
	// `CfnBudget.NotificationProperty.NotificationType`.
	NotificationType *string `json:"notificationType"`
	// `CfnBudget.NotificationProperty.Threshold`.
	Threshold *float64 `json:"threshold"`
	// `CfnBudget.NotificationProperty.ThresholdType`.
	ThresholdType *string `json:"thresholdType"`
}

type CfnBudget_NotificationWithSubscribersProperty struct {
	// `CfnBudget.NotificationWithSubscribersProperty.Notification`.
	Notification interface{} `json:"notification"`
	// `CfnBudget.NotificationWithSubscribersProperty.Subscribers`.
	Subscribers interface{} `json:"subscribers"`
}

type CfnBudget_SpendProperty struct {
	// `CfnBudget.SpendProperty.Amount`.
	Amount *float64 `json:"amount"`
	// `CfnBudget.SpendProperty.Unit`.
	Unit *string `json:"unit"`
}

type CfnBudget_SubscriberProperty struct {
	// `CfnBudget.SubscriberProperty.Address`.
	Address *string `json:"address"`
	// `CfnBudget.SubscriberProperty.SubscriptionType`.
	SubscriptionType *string `json:"subscriptionType"`
}

type CfnBudget_TimePeriodProperty struct {
	// `CfnBudget.TimePeriodProperty.End`.
	End *string `json:"end"`
	// `CfnBudget.TimePeriodProperty.Start`.
	Start *string `json:"start"`
}

// Properties for defining a `AWS::Budgets::Budget`.
type CfnBudgetProps struct {
	// `AWS::Budgets::Budget.Budget`.
	Budget interface{} `json:"budget"`
	// `AWS::Budgets::Budget.NotificationsWithSubscribers`.
	NotificationsWithSubscribers interface{} `json:"notificationsWithSubscribers"`
}

