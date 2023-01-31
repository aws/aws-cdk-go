package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::CloudWatch::CompositeAlarm`.
//
// The `AWS::CloudWatch::CompositeAlarm` type creates or updates a composite alarm. When you create a composite alarm, you specify a rule expression for the alarm that takes into account the alarm states of other alarms that you have created. The composite alarm goes into ALARM state only if all conditions of the rule are met.
//
// The alarms specified in a composite alarm's rule expression can include metric alarms and other composite alarms.
//
// Using composite alarms can reduce alarm noise. You can create multiple metric alarms, and also create a composite alarm and set up alerts only for the composite alarm. For example, you could create a composite alarm that goes into ALARM state only when more than one of the underlying metric alarms are in ALARM state.
//
// Currently, the only alarm actions that can be taken by composite alarms are notifying SNS topics.
//
// When this operation creates an alarm, the alarm state is immediately set to INSUFFICIENT_DATA. The alarm is then evaluated and its state is set appropriately. Any actions associated with the new state are then executed. For a composite alarm, this initial time after creation is the only time that the alarm can be in INSUFFICIENT_DATA state.
//
// When you update an existing alarm, its state is left unchanged, but the update completely overwrites the previous configuration of the alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCompositeAlarm := awscdk.Aws_cloudwatch.NewCfnCompositeAlarm(this, jsii.String("MyCfnCompositeAlarm"), &cfnCompositeAlarmProps{
//   	alarmRule: jsii.String("alarmRule"),
//
//   	// the properties below are optional
//   	actionsEnabled: jsii.Boolean(false),
//   	actionsSuppressor: jsii.String("actionsSuppressor"),
//   	actionsSuppressorExtensionPeriod: jsii.Number(123),
//   	actionsSuppressorWaitPeriod: jsii.Number(123),
//   	alarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	alarmDescription: jsii.String("alarmDescription"),
//   	alarmName: jsii.String("alarmName"),
//   	insufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
//   	},
//   	okActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   })
//
type CfnCompositeAlarm interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Indicates whether actions should be executed during any changes to the alarm state of the composite alarm.
	//
	// The default is TRUE.
	ActionsEnabled() interface{}
	SetActionsEnabled(val interface{})
	// Actions will be suppressed if the suppressor alarm is in the `ALARM` state.
	//
	// `ActionsSuppressor` can be an AlarmName or an Amazon Resource Name (ARN) from an existing alarm.
	ActionsSuppressor() *string
	SetActionsSuppressor(val *string)
	// The maximum time in seconds that the composite alarm waits after suppressor alarm goes out of the `ALARM` state.
	//
	// After this time, the composite alarm performs its actions.
	//
	// > `ExtensionPeriod` is required only when `ActionsSuppressor` is specified.
	ActionsSuppressorExtensionPeriod() *float64
	SetActionsSuppressorExtensionPeriod(val *float64)
	// The maximum time in seconds that the composite alarm waits for the suppressor alarm to go into the `ALARM` state.
	//
	// After this time, the composite alarm performs its actions.
	//
	// > `WaitPeriod` is required only when `ActionsSuppressor` is specified.
	ActionsSuppressorWaitPeriod() *float64
	SetActionsSuppressorWaitPeriod(val *float64)
	// The actions to execute when this alarm transitions to the ALARM state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the *Amazon CloudWatch API Reference* .
	AlarmActions() *[]*string
	SetAlarmActions(val *[]*string)
	// The description for the composite alarm.
	AlarmDescription() *string
	SetAlarmDescription(val *string)
	// The name for the composite alarm.
	//
	// This name must be unique within your AWS account.
	AlarmName() *string
	SetAlarmName(val *string)
	// An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state.
	//
	// For each alarm that you reference, you designate a function that specifies whether that alarm needs to be in ALARM state, OK state, or INSUFFICIENT_DATA state. You can use operators (AND, OR and NOT) to combine multiple functions in a single expression. You can use parenthesis to logically group the functions in your expression.
	//
	// You can use either alarm names or ARNs to reference the other alarms that are to be evaluated.
	//
	// Functions can include the following:
	//
	// - ALARM("alarm-name or alarm-ARN") is TRUE if the named alarm is in ALARM state.
	// - OK("alarm-name or alarm-ARN") is TRUE if the named alarm is in OK state.
	// - INSUFFICIENT_DATA("alarm-name or alarm-ARN") is TRUE if the named alarm is in INSUFFICIENT_DATA state.
	// - TRUE always evaluates to TRUE.
	// - FALSE always evaluates to FALSE.
	//
	// TRUE and FALSE are useful for testing a complex AlarmRule structure, and for testing your alarm actions.
	//
	// For more information about `AlarmRule` syntax, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the *Amazon CloudWatch API Reference* .
	AlarmRule() *string
	SetAlarmRule(val *string)
	// The ARN of the composite alarm, such as `arn:aws:cloudwatch:us-west-2:123456789012:alarm/CompositeAlarmName` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the *Amazon CloudWatch API Reference* .
	InsufficientDataActions() *[]*string
	SetInsufficientDataActions(val *[]*string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The actions to execute when this alarm transitions to the OK state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the *Amazon CloudWatch API Reference* .
	OkActions() *[]*string
	SetOkActions(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
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
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCompositeAlarm
type jsiiProxy_CfnCompositeAlarm struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCompositeAlarm) ActionsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) ActionsSuppressor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionsSuppressor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) ActionsSuppressorExtensionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actionsSuppressorExtensionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) ActionsSuppressorWaitPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actionsSuppressorWaitPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) AlarmActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) AlarmDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) AlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) AlarmRule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) InsufficientDataActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insufficientDataActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) OkActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"okActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarm) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudWatch::CompositeAlarm`.
func NewCfnCompositeAlarm(scope awscdk.Construct, id *string, props *CfnCompositeAlarmProps) CfnCompositeAlarm {
	_init_.Initialize()

	if err := validateNewCfnCompositeAlarmParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCompositeAlarm{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch.CfnCompositeAlarm",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudWatch::CompositeAlarm`.
func NewCfnCompositeAlarm_Override(c CfnCompositeAlarm, scope awscdk.Construct, id *string, props *CfnCompositeAlarmProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudwatch.CfnCompositeAlarm",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCompositeAlarm)SetActionsEnabled(val interface{}) {
	if err := j.validateSetActionsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionsEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnCompositeAlarm)SetActionsSuppressor(val *string) {
	_jsii_.Set(
		j,
		"actionsSuppressor",
		val,
	)
}

func (j *jsiiProxy_CfnCompositeAlarm)SetActionsSuppressorExtensionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"actionsSuppressorExtensionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCompositeAlarm)SetActionsSuppressorWaitPeriod(val *float64) {
	_jsii_.Set(
		j,
		"actionsSuppressorWaitPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCompositeAlarm)SetAlarmActions(val *[]*string) {
	_jsii_.Set(
		j,
		"alarmActions",
		val,
	)
}

func (j *jsiiProxy_CfnCompositeAlarm)SetAlarmDescription(val *string) {
	_jsii_.Set(
		j,
		"alarmDescription",
		val,
	)
}

func (j *jsiiProxy_CfnCompositeAlarm)SetAlarmName(val *string) {
	_jsii_.Set(
		j,
		"alarmName",
		val,
	)
}

func (j *jsiiProxy_CfnCompositeAlarm)SetAlarmRule(val *string) {
	if err := j.validateSetAlarmRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmRule",
		val,
	)
}

func (j *jsiiProxy_CfnCompositeAlarm)SetInsufficientDataActions(val *[]*string) {
	_jsii_.Set(
		j,
		"insufficientDataActions",
		val,
	)
}

func (j *jsiiProxy_CfnCompositeAlarm)SetOkActions(val *[]*string) {
	_jsii_.Set(
		j,
		"okActions",
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
func CfnCompositeAlarm_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCompositeAlarm_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudwatch.CfnCompositeAlarm",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCompositeAlarm_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnCompositeAlarm_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudwatch.CfnCompositeAlarm",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCompositeAlarm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCompositeAlarm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudwatch.CfnCompositeAlarm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCompositeAlarm_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cloudwatch.CfnCompositeAlarm",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCompositeAlarm) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCompositeAlarm) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCompositeAlarm) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCompositeAlarm) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCompositeAlarm) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCompositeAlarm) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCompositeAlarm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCompositeAlarm) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCompositeAlarm) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

