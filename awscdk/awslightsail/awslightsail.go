package awslightsail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Lightsail::Alarm`.
//
// The `AWS::Lightsail::Alarm` resource specifies an alarm that can be used to monitor a single metric for one of your Lightsail resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarm := awscdk.Aws_lightsail.NewCfnAlarm(this, jsii.String("MyCfnAlarm"), &cfnAlarmProps{
//   	alarmName: jsii.String("alarmName"),
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	evaluationPeriods: jsii.Number(123),
//   	metricName: jsii.String("metricName"),
//   	monitoredResourceName: jsii.String("monitoredResourceName"),
//   	threshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	contactProtocols: []*string{
//   		jsii.String("contactProtocols"),
//   	},
//   	datapointsToAlarm: jsii.Number(123),
//   	notificationEnabled: jsii.Boolean(false),
//   	notificationTriggers: []*string{
//   		jsii.String("notificationTriggers"),
//   	},
//   	treatMissingData: jsii.String("treatMissingData"),
//   })
//
type CfnAlarm interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the alarm.
	AlarmName() *string
	SetAlarmName(val *string)
	// The Amazon Resource Name (ARN) of the alarm.
	AttrAlarmArn() *string
	// The current state of the alarm.
	//
	// An alarm has the following possible states:
	//
	// - `ALARM` - The metric is outside of the defined threshold.
	// - `INSUFFICIENT_DATA` - The alarm has recently started, the metric is not available, or not enough data is available for the metric to determine the alarm state.
	// - `OK` - The metric is within the defined threshold.
	AttrState() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The arithmetic operation to use when comparing the specified statistic and threshold.
	ComparisonOperator() *string
	SetComparisonOperator(val *string)
	// The contact protocols for the alarm, such as `Email` , `SMS` (text messaging), or both.
	//
	// *Allowed Values* : `Email` | `SMS`.
	ContactProtocols() *[]*string
	SetContactProtocols(val *[]*string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The number of data points within the evaluation periods that must be breaching to cause the alarm to go to the `ALARM` state.
	DatapointsToAlarm() *float64
	SetDatapointsToAlarm(val *float64)
	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods() *float64
	SetEvaluationPeriods(val *float64)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the metric associated with the alarm.
	MetricName() *string
	SetMetricName(val *string)
	// The name of the Lightsail resource that the alarm monitors.
	MonitoredResourceName() *string
	SetMonitoredResourceName(val *string)
	// The tree node.
	Node() constructs.Node
	// A Boolean value indicating whether the alarm is enabled.
	NotificationEnabled() interface{}
	SetNotificationEnabled(val interface{})
	// The alarm states that trigger a notification.
	//
	// > To specify the `OK` and `INSUFFICIENT_DATA` values, you must also specify `ContactProtocols` values. Otherwise, the `OK` and `INSUFFICIENT_DATA` values will not take effect and the stack will drift.
	//
	// *Allowed Values* : `OK` | `ALARM` | `INSUFFICIENT_DATA`.
	NotificationTriggers() *[]*string
	SetNotificationTriggers(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The value against which the specified statistic is compared.
	Threshold() *float64
	SetThreshold(val *float64)
	// Specifies how the alarm handles missing data points.
	//
	// An alarm can treat missing data in the following ways:
	//
	// - `breaching` - Assumes the missing data is not within the threshold. Missing data counts towards the number of times that the metric is not within the threshold.
	// - `notBreaching` - Assumes the missing data is within the threshold. Missing data does not count towards the number of times that the metric is not within the threshold.
	// - `ignore` - Ignores the missing data. Maintains the current alarm state.
	// - `missing` - Missing data is treated as missing.
	TreatMissingData() *string
	SetTreatMissingData(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAlarm
type jsiiProxy_CfnAlarm struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAlarm) AlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) AttrAlarmArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAlarmArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) ComparisonOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) ContactProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contactProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) DatapointsToAlarm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) EvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) MonitoredResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoredResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) NotificationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) NotificationTriggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) TreatMissingData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatMissingData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Alarm`.
func NewCfnAlarm(scope constructs.Construct, id *string, props *CfnAlarmProps) CfnAlarm {
	_init_.Initialize()

	j := jsiiProxy_CfnAlarm{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnAlarm",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Alarm`.
func NewCfnAlarm_Override(c CfnAlarm, scope constructs.Construct, id *string, props *CfnAlarmProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnAlarm",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAlarm) SetAlarmName(val *string) {
	_jsii_.Set(
		j,
		"alarmName",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm) SetComparisonOperator(val *string) {
	_jsii_.Set(
		j,
		"comparisonOperator",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm) SetContactProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"contactProtocols",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm) SetDatapointsToAlarm(val *float64) {
	_jsii_.Set(
		j,
		"datapointsToAlarm",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm) SetEvaluationPeriods(val *float64) {
	_jsii_.Set(
		j,
		"evaluationPeriods",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm) SetMetricName(val *string) {
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm) SetMonitoredResourceName(val *string) {
	_jsii_.Set(
		j,
		"monitoredResourceName",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm) SetNotificationEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"notificationEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm) SetNotificationTriggers(val *[]*string) {
	_jsii_.Set(
		j,
		"notificationTriggers",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm) SetThreshold(val *float64) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm) SetTreatMissingData(val *string) {
	_jsii_.Set(
		j,
		"treatMissingData",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAlarm_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnAlarm",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAlarm_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnAlarm",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnAlarm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnAlarm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAlarm_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnAlarm",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAlarm) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAlarm) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAlarm) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAlarm) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAlarm) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAlarm) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAlarm) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAlarm) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarm) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarm) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAlarm) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAlarm) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarm) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlarm) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnAlarm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarmProps := &cfnAlarmProps{
//   	alarmName: jsii.String("alarmName"),
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	evaluationPeriods: jsii.Number(123),
//   	metricName: jsii.String("metricName"),
//   	monitoredResourceName: jsii.String("monitoredResourceName"),
//   	threshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	contactProtocols: []*string{
//   		jsii.String("contactProtocols"),
//   	},
//   	datapointsToAlarm: jsii.Number(123),
//   	notificationEnabled: jsii.Boolean(false),
//   	notificationTriggers: []*string{
//   		jsii.String("notificationTriggers"),
//   	},
//   	treatMissingData: jsii.String("treatMissingData"),
//   }
//
type CfnAlarmProps struct {
	// The name of the alarm.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// The arithmetic operation to use when comparing the specified statistic and threshold.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The name of the metric associated with the alarm.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The name of the Lightsail resource that the alarm monitors.
	MonitoredResourceName *string `field:"required" json:"monitoredResourceName" yaml:"monitoredResourceName"`
	// The value against which the specified statistic is compared.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// The contact protocols for the alarm, such as `Email` , `SMS` (text messaging), or both.
	//
	// *Allowed Values* : `Email` | `SMS`.
	ContactProtocols *[]*string `field:"optional" json:"contactProtocols" yaml:"contactProtocols"`
	// The number of data points within the evaluation periods that must be breaching to cause the alarm to go to the `ALARM` state.
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// A Boolean value indicating whether the alarm is enabled.
	NotificationEnabled interface{} `field:"optional" json:"notificationEnabled" yaml:"notificationEnabled"`
	// The alarm states that trigger a notification.
	//
	// > To specify the `OK` and `INSUFFICIENT_DATA` values, you must also specify `ContactProtocols` values. Otherwise, the `OK` and `INSUFFICIENT_DATA` values will not take effect and the stack will drift.
	//
	// *Allowed Values* : `OK` | `ALARM` | `INSUFFICIENT_DATA`.
	NotificationTriggers *[]*string `field:"optional" json:"notificationTriggers" yaml:"notificationTriggers"`
	// Specifies how the alarm handles missing data points.
	//
	// An alarm can treat missing data in the following ways:
	//
	// - `breaching` - Assumes the missing data is not within the threshold. Missing data counts towards the number of times that the metric is not within the threshold.
	// - `notBreaching` - Assumes the missing data is within the threshold. Missing data does not count towards the number of times that the metric is not within the threshold.
	// - `ignore` - Ignores the missing data. Maintains the current alarm state.
	// - `missing` - Missing data is treated as missing.
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
}

// A CloudFormation `AWS::Lightsail::Bucket`.
//
// The `AWS::Lightsail::Bucket` resource specifies a bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBucket := awscdk.Aws_lightsail.NewCfnBucket(this, jsii.String("MyCfnBucket"), &cfnBucketProps{
//   	bucketName: jsii.String("bucketName"),
//   	bundleId: jsii.String("bundleId"),
//
//   	// the properties below are optional
//   	accessRules: &accessRulesProperty{
//   		allowPublicOverrides: jsii.Boolean(false),
//   		objectAccess: jsii.String("objectAccess"),
//   	},
//   	objectVersioning: jsii.Boolean(false),
//   	readOnlyAccessAccounts: []*string{
//   		jsii.String("readOnlyAccessAccounts"),
//   	},
//   	resourcesReceivingAccess: []*string{
//   		jsii.String("resourcesReceivingAccess"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnBucket interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// An object that describes the access rules for the bucket.
	AccessRules() interface{}
	SetAccessRules(val interface{})
	// A Boolean value indicating whether the bundle that is currently applied to your distribution can be changed to another bundle.
	AttrAbleToUpdateBundle() awscdk.IResolvable
	// The Amazon Resource Name (ARN) of the bucket.
	AttrBucketArn() *string
	// The URL of the bucket.
	AttrUrl() *string
	// The name of the bucket.
	BucketName() *string
	SetBucketName(val *string)
	// The bundle ID for the bucket (for example, `small_1_0` ).
	//
	// A bucket bundle specifies the monthly cost, storage space, and data transfer quota for a bucket.
	BundleId() *string
	SetBundleId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Indicates whether object versioning is enabled for the bucket.
	//
	// The following options can be configured:
	//
	// - `Enabled` - Object versioning is enabled.
	// - `Suspended` - Object versioning was previously enabled but is currently suspended. Existing object versions are retained.
	// - `NeverEnabled` - Object versioning has never been enabled.
	ObjectVersioning() interface{}
	SetObjectVersioning(val interface{})
	// An array of AWS account IDs that have read-only access to the bucket.
	ReadOnlyAccessAccounts() *[]*string
	SetReadOnlyAccessAccounts(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// An array of Lightsail instances that have access to the bucket.
	ResourcesReceivingAccess() *[]*string
	SetResourcesReceivingAccess(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBucket
type jsiiProxy_CfnBucket struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBucket) AccessRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) AttrAbleToUpdateBundle() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrAbleToUpdateBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) AttrBucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrBucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) AttrUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) ObjectVersioning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectVersioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) ReadOnlyAccessAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readOnlyAccessAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) ResourcesReceivingAccess() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesReceivingAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Bucket`.
func NewCfnBucket(scope constructs.Construct, id *string, props *CfnBucketProps) CfnBucket {
	_init_.Initialize()

	j := jsiiProxy_CfnBucket{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnBucket",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Bucket`.
func NewCfnBucket_Override(c CfnBucket, scope constructs.Construct, id *string, props *CfnBucketProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnBucket",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBucket) SetAccessRules(val interface{}) {
	_jsii_.Set(
		j,
		"accessRules",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetBucketName(val *string) {
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetObjectVersioning(val interface{}) {
	_jsii_.Set(
		j,
		"objectVersioning",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetReadOnlyAccessAccounts(val *[]*string) {
	_jsii_.Set(
		j,
		"readOnlyAccessAccounts",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetResourcesReceivingAccess(val *[]*string) {
	_jsii_.Set(
		j,
		"resourcesReceivingAccess",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnBucket_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnBucket",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBucket_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnBucket",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnBucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnBucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBucket_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnBucket",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBucket) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBucket) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBucket) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBucket) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBucket) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBucket) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBucket) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBucket) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucket) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucket) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBucket) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBucket) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucket) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucket) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `AccessRules` is a property of the [AWS::Lightsail::Bucket](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-bucket.html) resource. It describes access rules for a bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessRulesProperty := &accessRulesProperty{
//   	allowPublicOverrides: jsii.Boolean(false),
//   	objectAccess: jsii.String("objectAccess"),
//   }
//
type CfnBucket_AccessRulesProperty struct {
	// A Boolean value indicating whether the access control list (ACL) permissions that are applied to individual objects override the `GetObject` option that is currently specified.
	//
	// When this is true, you can use the [PutObjectAcl](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectAcl.html) Amazon S3 API operation to set individual objects to public (read-only) or private, using either the `public-read` ACL or the `private` ACL.
	AllowPublicOverrides interface{} `field:"optional" json:"allowPublicOverrides" yaml:"allowPublicOverrides"`
	// Specifies the anonymous access to all objects in a bucket.
	//
	// The following options can be specified:
	//
	// - `public` - Sets all objects in the bucket to public (read-only), making them readable by everyone on the internet.
	//
	// If the `GetObject` value is set to `public` , then all objects in the bucket default to public regardless of the `allowPublicOverrides` value.
	// - `private` - Sets all objects in the bucket to private, making them readable only by you and anyone that you grant access to.
	//
	// If the `GetObject` value is set to `private` , and the `allowPublicOverrides` value is set to `true` , then all objects in the bucket default to private unless they are configured with a `public-read` ACL. Individual objects with a `public-read` ACL are readable by everyone on the internet.
	ObjectAccess *string `field:"optional" json:"objectAccess" yaml:"objectAccess"`
}

// Properties for defining a `CfnBucket`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBucketProps := &cfnBucketProps{
//   	bucketName: jsii.String("bucketName"),
//   	bundleId: jsii.String("bundleId"),
//
//   	// the properties below are optional
//   	accessRules: &accessRulesProperty{
//   		allowPublicOverrides: jsii.Boolean(false),
//   		objectAccess: jsii.String("objectAccess"),
//   	},
//   	objectVersioning: jsii.Boolean(false),
//   	readOnlyAccessAccounts: []*string{
//   		jsii.String("readOnlyAccessAccounts"),
//   	},
//   	resourcesReceivingAccess: []*string{
//   		jsii.String("resourcesReceivingAccess"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucketProps struct {
	// The name of the bucket.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The bundle ID for the bucket (for example, `small_1_0` ).
	//
	// A bucket bundle specifies the monthly cost, storage space, and data transfer quota for a bucket.
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// An object that describes the access rules for the bucket.
	AccessRules interface{} `field:"optional" json:"accessRules" yaml:"accessRules"`
	// Indicates whether object versioning is enabled for the bucket.
	//
	// The following options can be configured:
	//
	// - `Enabled` - Object versioning is enabled.
	// - `Suspended` - Object versioning was previously enabled but is currently suspended. Existing object versions are retained.
	// - `NeverEnabled` - Object versioning has never been enabled.
	ObjectVersioning interface{} `field:"optional" json:"objectVersioning" yaml:"objectVersioning"`
	// An array of AWS account IDs that have read-only access to the bucket.
	ReadOnlyAccessAccounts *[]*string `field:"optional" json:"readOnlyAccessAccounts" yaml:"readOnlyAccessAccounts"`
	// An array of Lightsail instances that have access to the bucket.
	ResourcesReceivingAccess *[]*string `field:"optional" json:"resourcesReceivingAccess" yaml:"resourcesReceivingAccess"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Lightsail::Certificate`.
//
// The `AWS::Lightsail::Certificate` resource specifies an SSL/TLS certificate that you can use with a content delivery network (CDN) distribution and a container service.
//
// > For information about certificates that you can use with a load balancer, see [AWS::Lightsail::LoadBalancerTlsCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-loadbalancertlscertificate.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificate := awscdk.Aws_lightsail.NewCfnCertificate(this, jsii.String("MyCfnCertificate"), &cfnCertificateProps{
//   	certificateName: jsii.String("certificateName"),
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	subjectAlternativeNames: []*string{
//   		jsii.String("subjectAlternativeNames"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnCertificate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the certificate.
	AttrCertificateArn() *string
	// The validation status of the certificate.
	AttrStatus() *string
	// The name of the certificate.
	CertificateName() *string
	SetCertificateName(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The domain name of the certificate.
	DomainName() *string
	SetDomainName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of strings that specify the alternate domains (such as `example.org` ) and subdomains (such as `blog.example.com` ) of the certificate.
	SubjectAlternativeNames() *[]*string
	SetSubjectAlternativeNames(val *[]*string)
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCertificate
type jsiiProxy_CfnCertificate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCertificate) AttrCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) SubjectAlternativeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Certificate`.
func NewCfnCertificate(scope constructs.Construct, id *string, props *CfnCertificateProps) CfnCertificate {
	_init_.Initialize()

	j := jsiiProxy_CfnCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Certificate`.
func NewCfnCertificate_Override(c CfnCertificate, scope constructs.Construct, id *string, props *CfnCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnCertificate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCertificate) SetCertificateName(val *string) {
	_jsii_.Set(
		j,
		"certificateName",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnCertificate) SetSubjectAlternativeNames(val *[]*string) {
	_jsii_.Set(
		j,
		"subjectAlternativeNames",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCertificate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnCertificate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCertificate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnCertificate",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCertificate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnCertificate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCertificate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCertificate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCertificate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCertificate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCertificate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCertificate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCertificate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &cfnCertificateProps{
//   	certificateName: jsii.String("certificateName"),
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	subjectAlternativeNames: []*string{
//   		jsii.String("subjectAlternativeNames"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCertificateProps struct {
	// The name of the certificate.
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// The domain name of the certificate.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// An array of strings that specify the alternate domains (such as `example.org` ) and subdomains (such as `blog.example.com` ) of the certificate.
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Lightsail::Container`.
//
// The `AWS::Lightsail::Container` resource specifies a container service.
//
// A Lightsail container service is a compute resource to which you can deploy containers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainer := awscdk.Aws_lightsail.NewCfnContainer(this, jsii.String("MyCfnContainer"), &cfnContainerProps{
//   	power: jsii.String("power"),
//   	scale: jsii.Number(123),
//   	serviceName: jsii.String("serviceName"),
//
//   	// the properties below are optional
//   	containerServiceDeployment: &containerServiceDeploymentProperty{
//   		containers: []interface{}{
//   			&containerProperty{
//   				command: []*string{
//   					jsii.String("command"),
//   				},
//   				containerName: jsii.String("containerName"),
//   				environment: []interface{}{
//   					&environmentVariableProperty{
//   						value: jsii.String("value"),
//   						variable: jsii.String("variable"),
//   					},
//   				},
//   				image: jsii.String("image"),
//   				ports: []interface{}{
//   					&portInfoProperty{
//   						port: jsii.String("port"),
//   						protocol: jsii.String("protocol"),
//   					},
//   				},
//   			},
//   		},
//   		publicEndpoint: &publicEndpointProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			healthCheckConfig: &healthCheckConfigProperty{
//   				healthyThreshold: jsii.Number(123),
//   				intervalSeconds: jsii.Number(123),
//   				path: jsii.String("path"),
//   				successCodes: jsii.String("successCodes"),
//   				timeoutSeconds: jsii.Number(123),
//   				unhealthyThreshold: jsii.Number(123),
//   			},
//   		},
//   	},
//   	isDisabled: jsii.Boolean(false),
//   	publicDomainNames: []interface{}{
//   		&publicDomainNameProperty{
//   			certificateName: jsii.String("certificateName"),
//   			domainNames: []*string{
//   				jsii.String("domainNames"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnContainer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the container.
	AttrContainerArn() *string
	// The publicly accessible URL of the container service.
	//
	// If no public endpoint is specified in the current deployment, this URL returns a 404 response.
	AttrUrl() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// An object that describes the current container deployment of the container service.
	ContainerServiceDeployment() interface{}
	SetContainerServiceDeployment(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A Boolean value indicating whether the container service is disabled.
	IsDisabled() interface{}
	SetIsDisabled(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// The power specification of the container service.
	//
	// The power specifies the amount of RAM, the number of vCPUs, and the base price of the container service.
	Power() *string
	SetPower(val *string)
	// The public domain name of the container service, such as `example.com` and `www.example.com` .
	//
	// You can specify up to four public domain names for a container service. The domain names that you specify are used when you create a deployment with a container that is configured as the public endpoint of your container service.
	//
	// If you don't specify public domain names, then you can use the default domain of the container service.
	//
	// > You must create and validate an SSL/TLS certificate before you can use public domain names with your container service. Use the [AWS::Lightsail::Certificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-certificate.html) resource to create a certificate for the public domain names that you want to use with your container service.
	PublicDomainNames() interface{}
	SetPublicDomainNames(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The scale specification of the container service.
	//
	// The scale specifies the allocated compute nodes of the container service.
	Scale() *float64
	SetScale(val *float64)
	// The name of the container service.
	ServiceName() *string
	SetServiceName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnContainer
type jsiiProxy_CfnContainer struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnContainer) AttrContainerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrContainerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) AttrUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) ContainerServiceDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerServiceDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) IsDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) Power() *string {
	var returns *string
	_jsii_.Get(
		j,
		"power",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) PublicDomainNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicDomainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) Scale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Container`.
func NewCfnContainer(scope constructs.Construct, id *string, props *CfnContainerProps) CfnContainer {
	_init_.Initialize()

	j := jsiiProxy_CfnContainer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnContainer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Container`.
func NewCfnContainer_Override(c CfnContainer, scope constructs.Construct, id *string, props *CfnContainerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnContainer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnContainer) SetContainerServiceDeployment(val interface{}) {
	_jsii_.Set(
		j,
		"containerServiceDeployment",
		val,
	)
}

func (j *jsiiProxy_CfnContainer) SetIsDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"isDisabled",
		val,
	)
}

func (j *jsiiProxy_CfnContainer) SetPower(val *string) {
	_jsii_.Set(
		j,
		"power",
		val,
	)
}

func (j *jsiiProxy_CfnContainer) SetPublicDomainNames(val interface{}) {
	_jsii_.Set(
		j,
		"publicDomainNames",
		val,
	)
}

func (j *jsiiProxy_CfnContainer) SetScale(val *float64) {
	_jsii_.Set(
		j,
		"scale",
		val,
	)
}

func (j *jsiiProxy_CfnContainer) SetServiceName(val *string) {
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnContainer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnContainer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnContainer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnContainer",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnContainer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnContainer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContainer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnContainer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContainer) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnContainer) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnContainer) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnContainer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnContainer) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnContainer) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnContainer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnContainer) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainer) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainer) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnContainer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnContainer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainer) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainer) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `Container` is a property of the [ContainerServiceDeployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-containerservicedeployment.html) property. It describes the settings of a container that will be launched, or that is launched, to an Amazon Lightsail container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerProperty := &containerProperty{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	containerName: jsii.String("containerName"),
//   	environment: []interface{}{
//   		&environmentVariableProperty{
//   			value: jsii.String("value"),
//   			variable: jsii.String("variable"),
//   		},
//   	},
//   	image: jsii.String("image"),
//   	ports: []interface{}{
//   		&portInfoProperty{
//   			port: jsii.String("port"),
//   			protocol: jsii.String("protocol"),
//   		},
//   	},
//   }
//
type CfnContainer_ContainerProperty struct {
	// The launch command for the container.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The name of the container.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The environment variables of the container.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The name of the image used for the container.
	//
	// Container images that are sourced from (registered and stored on) your container service start with a colon ( `:` ). For example, if your container service name is `container-service-1` , the container image label is `mystaticsite` , and you want to use the third version ( `3` ) of the registered container image, then you should specify `:container-service-1.mystaticsite.3` . To use the latest version of a container image, specify `latest` instead of a version number (for example, `:container-service-1.mystaticsite.latest` ). Your container service will automatically use the highest numbered version of the registered container image.
	//
	// Container images that are sourced from a public registry like Docker Hub don’t start with a colon. For example, `nginx:latest` or `nginx` .
	Image *string `field:"optional" json:"image" yaml:"image"`
	// An object that describes the open firewall ports and protocols of the container.
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}

// `ContainerServiceDeployment` is a property of the [AWS::Lightsail::Container](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html) resource. It describes a container deployment configuration of a container service.
//
// A deployment specifies the settings, such as the ports and launch command, of containers that are deployed to your container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerServiceDeploymentProperty := &containerServiceDeploymentProperty{
//   	containers: []interface{}{
//   		&containerProperty{
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			containerName: jsii.String("containerName"),
//   			environment: []interface{}{
//   				&environmentVariableProperty{
//   					value: jsii.String("value"),
//   					variable: jsii.String("variable"),
//   				},
//   			},
//   			image: jsii.String("image"),
//   			ports: []interface{}{
//   				&portInfoProperty{
//   					port: jsii.String("port"),
//   					protocol: jsii.String("protocol"),
//   				},
//   			},
//   		},
//   	},
//   	publicEndpoint: &publicEndpointProperty{
//   		containerName: jsii.String("containerName"),
//   		containerPort: jsii.Number(123),
//   		healthCheckConfig: &healthCheckConfigProperty{
//   			healthyThreshold: jsii.Number(123),
//   			intervalSeconds: jsii.Number(123),
//   			path: jsii.String("path"),
//   			successCodes: jsii.String("successCodes"),
//   			timeoutSeconds: jsii.Number(123),
//   			unhealthyThreshold: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnContainer_ContainerServiceDeploymentProperty struct {
	// An object that describes the configuration for the containers of the deployment.
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// An object that describes the endpoint of the deployment.
	PublicEndpoint interface{} `field:"optional" json:"publicEndpoint" yaml:"publicEndpoint"`
}

// `EnvironmentVariable` is a property of the [Container](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-container.html) property. It describes the environment variables of a container on a container service which are key-value parameters that provide dynamic configuration of the application or script run by the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentVariableProperty := &environmentVariableProperty{
//   	value: jsii.String("value"),
//   	variable: jsii.String("variable"),
//   }
//
type CfnContainer_EnvironmentVariableProperty struct {
	// The environment variable value.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// The environment variable key.
	Variable *string `field:"optional" json:"variable" yaml:"variable"`
}

// `HealthCheckConfig` is a property of the [PublicEndpoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-publicendpoint.html) property. It describes the healthcheck configuration of a container deployment on a container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckConfigProperty := &healthCheckConfigProperty{
//   	healthyThreshold: jsii.Number(123),
//   	intervalSeconds: jsii.Number(123),
//   	path: jsii.String("path"),
//   	successCodes: jsii.String("successCodes"),
//   	timeoutSeconds: jsii.Number(123),
//   	unhealthyThreshold: jsii.Number(123),
//   }
//
type CfnContainer_HealthCheckConfigProperty struct {
	// The number of consecutive health check successes required before moving the container to the `Healthy` state.
	//
	// The default value is `2` .
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The approximate interval, in seconds, between health checks of an individual container.
	//
	// You can specify between `5` and `300` seconds. The default value is `5` .
	IntervalSeconds *float64 `field:"optional" json:"intervalSeconds" yaml:"intervalSeconds"`
	// The path on the container on which to perform the health check.
	//
	// The default value is `/` .
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The HTTP codes to use when checking for a successful response from a container.
	//
	// You can specify values between `200` and `499` . You can specify multiple values (for example, `200,202` ) or a range of values (for example, `200-299` ).
	SuccessCodes *string `field:"optional" json:"successCodes" yaml:"successCodes"`
	// The amount of time, in seconds, during which no response means a failed health check.
	//
	// You can specify between `2` and `60` seconds. The default value is `2` .
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The number of consecutive health check failures required before moving the container to the `Unhealthy` state.
	//
	// The default value is `2` .
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

// `PortInfo` is a property of the [Container](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-container.html) property. It describes the ports to open and the protocols to use for a container on a Amazon Lightsail container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portInfoProperty := &portInfoProperty{
//   	port: jsii.String("port"),
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnContainer_PortInfoProperty struct {
	// The open firewall ports of the container.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol name for the open ports.
	//
	// *Allowed values* : `HTTP` | `HTTPS` | `TCP` | `UDP`.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

// `PublicDomainName` is a property of the [AWS::Lightsail::Container](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-container.html) resource. It describes the public domain names to use with a container service, such as `example.com` and `www.example.com` . It also describes the certificates to use with a container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicDomainNameProperty := &publicDomainNameProperty{
//   	certificateName: jsii.String("certificateName"),
//   	domainNames: []*string{
//   		jsii.String("domainNames"),
//   	},
//   }
//
type CfnContainer_PublicDomainNameProperty struct {
	// The name of the certificate for the public domains.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// The public domain names to use with the container service.
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
}

// `PublicEndpoint` is a property of the [ContainerServiceDeployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-containerservicedeployment.html) property. It describes describes the settings of the public endpoint of a container on a container service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicEndpointProperty := &publicEndpointProperty{
//   	containerName: jsii.String("containerName"),
//   	containerPort: jsii.Number(123),
//   	healthCheckConfig: &healthCheckConfigProperty{
//   		healthyThreshold: jsii.Number(123),
//   		intervalSeconds: jsii.Number(123),
//   		path: jsii.String("path"),
//   		successCodes: jsii.String("successCodes"),
//   		timeoutSeconds: jsii.Number(123),
//   		unhealthyThreshold: jsii.Number(123),
//   	},
//   }
//
type CfnContainer_PublicEndpointProperty struct {
	// The name of the container entry of the deployment that the endpoint configuration applies to.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port of the specified container to which traffic is forwarded to.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// An object that describes the health check configuration of the container.
	HealthCheckConfig interface{} `field:"optional" json:"healthCheckConfig" yaml:"healthCheckConfig"`
}

// Properties for defining a `CfnContainer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerProps := &cfnContainerProps{
//   	power: jsii.String("power"),
//   	scale: jsii.Number(123),
//   	serviceName: jsii.String("serviceName"),
//
//   	// the properties below are optional
//   	containerServiceDeployment: &containerServiceDeploymentProperty{
//   		containers: []interface{}{
//   			&containerProperty{
//   				command: []*string{
//   					jsii.String("command"),
//   				},
//   				containerName: jsii.String("containerName"),
//   				environment: []interface{}{
//   					&environmentVariableProperty{
//   						value: jsii.String("value"),
//   						variable: jsii.String("variable"),
//   					},
//   				},
//   				image: jsii.String("image"),
//   				ports: []interface{}{
//   					&portInfoProperty{
//   						port: jsii.String("port"),
//   						protocol: jsii.String("protocol"),
//   					},
//   				},
//   			},
//   		},
//   		publicEndpoint: &publicEndpointProperty{
//   			containerName: jsii.String("containerName"),
//   			containerPort: jsii.Number(123),
//   			healthCheckConfig: &healthCheckConfigProperty{
//   				healthyThreshold: jsii.Number(123),
//   				intervalSeconds: jsii.Number(123),
//   				path: jsii.String("path"),
//   				successCodes: jsii.String("successCodes"),
//   				timeoutSeconds: jsii.Number(123),
//   				unhealthyThreshold: jsii.Number(123),
//   			},
//   		},
//   	},
//   	isDisabled: jsii.Boolean(false),
//   	publicDomainNames: []interface{}{
//   		&publicDomainNameProperty{
//   			certificateName: jsii.String("certificateName"),
//   			domainNames: []*string{
//   				jsii.String("domainNames"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnContainerProps struct {
	// The power specification of the container service.
	//
	// The power specifies the amount of RAM, the number of vCPUs, and the base price of the container service.
	Power *string `field:"required" json:"power" yaml:"power"`
	// The scale specification of the container service.
	//
	// The scale specifies the allocated compute nodes of the container service.
	Scale *float64 `field:"required" json:"scale" yaml:"scale"`
	// The name of the container service.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// An object that describes the current container deployment of the container service.
	ContainerServiceDeployment interface{} `field:"optional" json:"containerServiceDeployment" yaml:"containerServiceDeployment"`
	// A Boolean value indicating whether the container service is disabled.
	IsDisabled interface{} `field:"optional" json:"isDisabled" yaml:"isDisabled"`
	// The public domain name of the container service, such as `example.com` and `www.example.com` .
	//
	// You can specify up to four public domain names for a container service. The domain names that you specify are used when you create a deployment with a container that is configured as the public endpoint of your container service.
	//
	// If you don't specify public domain names, then you can use the default domain of the container service.
	//
	// > You must create and validate an SSL/TLS certificate before you can use public domain names with your container service. Use the [AWS::Lightsail::Certificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-certificate.html) resource to create a certificate for the public domain names that you want to use with your container service.
	PublicDomainNames interface{} `field:"optional" json:"publicDomainNames" yaml:"publicDomainNames"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Lightsail::Database`.
//
// The `AWS::Lightsail::Database` resource specifies an Amazon Lightsail database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatabase := awscdk.Aws_lightsail.NewCfnDatabase(this, jsii.String("MyCfnDatabase"), &cfnDatabaseProps{
//   	masterDatabaseName: jsii.String("masterDatabaseName"),
//   	masterUsername: jsii.String("masterUsername"),
//   	relationalDatabaseBlueprintId: jsii.String("relationalDatabaseBlueprintId"),
//   	relationalDatabaseBundleId: jsii.String("relationalDatabaseBundleId"),
//   	relationalDatabaseName: jsii.String("relationalDatabaseName"),
//
//   	// the properties below are optional
//   	availabilityZone: jsii.String("availabilityZone"),
//   	backupRetention: jsii.Boolean(false),
//   	caCertificateIdentifier: jsii.String("caCertificateIdentifier"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	relationalDatabaseParameters: []interface{}{
//   		&relationalDatabaseParameterProperty{
//   			allowedValues: jsii.String("allowedValues"),
//   			applyMethod: jsii.String("applyMethod"),
//   			applyType: jsii.String("applyType"),
//   			dataType: jsii.String("dataType"),
//   			description: jsii.String("description"),
//   			isModifiable: jsii.Boolean(false),
//   			parameterName: jsii.String("parameterName"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	rotateMasterUserPassword: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDatabase interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the database (for example, `arn:aws:lightsail:us-east-2:123456789101:RelationalDatabase/244ad76f-8aad-4741-809f-12345EXAMPLE` ).
	AttrDatabaseArn() *string
	// The Availability Zone for the database.
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// A Boolean value indicating whether automated backup retention is enabled for the database.
	BackupRetention() interface{}
	SetBackupRetention(val interface{})
	// The certificate associated with the database.
	CaCertificateIdentifier() *string
	SetCaCertificateIdentifier(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The meaning of this parameter differs according to the database engine you use.
	//
	// *MySQL*
	//
	// The name of the database to create when the Lightsail database resource is created. If this parameter isn't specified, no database is created in the database resource.
	//
	// Constraints:
	//
	// - Must contain 1-64 letters or numbers.
	// - Must begin with a letter. Subsequent characters can be letters, underscores, or numbers (0-9).
	// - Can't be a word reserved by the specified database engine.
	//
	// For more information about reserved words in MySQL, see the Keywords and Reserved Words articles for [MySQL 5.6](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.6/en/keywords.html) , [MySQL 5.7](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.7/en/keywords.html) , and [MySQL 8.0](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/8.0/en/keywords.html) .
	//
	// *PostgreSQL*
	//
	// The name of the database to create when the Lightsail database resource is created. If this parameter isn't specified, a database named `postgres` is created in the database resource.
	//
	// Constraints:
	//
	// - Must contain 1-63 letters or numbers.
	// - Must begin with a letter. Subsequent characters can be letters, underscores, or numbers (0-9).
	// - Can't be a word reserved by the specified database engine.
	//
	// For more information about reserved words in PostgreSQL, see the SQL Key Words articles for [PostgreSQL 9.6](https://docs.aws.amazon.com/https://www.postgresql.org/docs/9.6/sql-keywords-appendix.html) , [PostgreSQL 10](https://docs.aws.amazon.com/https://www.postgresql.org/docs/10/sql-keywords-appendix.html) , [PostgreSQL 11](https://docs.aws.amazon.com/https://www.postgresql.org/docs/11/sql-keywords-appendix.html) , and [PostgreSQL 12](https://docs.aws.amazon.com/https://www.postgresql.org/docs/12/sql-keywords-appendix.html) .
	MasterDatabaseName() *string
	SetMasterDatabaseName(val *string)
	// The name for the primary user.
	//
	// *MySQL*
	//
	// Constraints:
	//
	// - Required for MySQL.
	// - Must be 1-16 letters or numbers. Can contain underscores.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// For more information about reserved words in MySQL 5.6 or 5.7, see the Keywords and Reserved Words articles for [MySQL 5.6](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.6/en/keywords.html) , [MySQL 5.7](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.7/en/keywords.html) , or [MySQL 8.0](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/8.0/en/keywords.html) .
	//
	// *PostgreSQL*
	//
	// Constraints:
	//
	// - Required for PostgreSQL.
	// - Must be 1-63 letters or numbers. Can contain underscores.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// For more information about reserved words in MySQL 5.6 or 5.7, see the Keywords and Reserved Words articles for [PostgreSQL 9.6](https://docs.aws.amazon.com/https://www.postgresql.org/docs/9.6/sql-keywords-appendix.html) , [PostgreSQL 10](https://docs.aws.amazon.com/https://www.postgresql.org/docs/10/sql-keywords-appendix.html) , [PostgreSQL 11](https://docs.aws.amazon.com/https://www.postgresql.org/docs/11/sql-keywords-appendix.html) , and [PostgreSQL 12](https://docs.aws.amazon.com/https://www.postgresql.org/docs/12/sql-keywords-appendix.html) .
	MasterUsername() *string
	SetMasterUsername(val *string)
	// The password for the primary user of the database.
	//
	// The password can include any printable ASCII character except the following: /, ", or @. It cannot contain spaces.
	//
	// > The `MasterUserPassword` and `RotateMasterUserPassword` parameters cannot be used together in the same template.
	//
	// *MySQL*
	//
	// Constraints: Must contain 8-41 characters.
	//
	// *PostgreSQL*
	//
	// Constraints: Must contain 8-128 characters.
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	// The tree node.
	Node() constructs.Node
	// The daily time range during which automated backups are created for the database (for example, `16:00-16:30` ).
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	// The weekly time range during which system maintenance can occur for the database, formatted as follows: `ddd:hh24:mi-ddd:hh24:mi` .
	//
	// For example, `Tue:17:00-Tue:17:30` .
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	// A Boolean value indicating whether the database is accessible to anyone on the internet.
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The blueprint ID for the database (for example, `mysql_8_0` ).
	RelationalDatabaseBlueprintId() *string
	SetRelationalDatabaseBlueprintId(val *string)
	// The bundle ID for the database (for example, `medium_1_0` ).
	RelationalDatabaseBundleId() *string
	SetRelationalDatabaseBundleId(val *string)
	// The name of the instance.
	RelationalDatabaseName() *string
	SetRelationalDatabaseName(val *string)
	// An array of parameters for the database.
	RelationalDatabaseParameters() interface{}
	SetRelationalDatabaseParameters(val interface{})
	// A Boolean value indicating whether to change the primary user password to a new, strong password generated by Lightsail .
	//
	// > The `RotateMasterUserPassword` and `MasterUserPassword` parameters cannot be used together in the same template.
	RotateMasterUserPassword() interface{}
	SetRotateMasterUserPassword(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDatabase
type jsiiProxy_CfnDatabase struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDatabase) AttrDatabaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDatabaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) BackupRetention() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CaCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) MasterDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) RelationalDatabaseBlueprintId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseBlueprintId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) RelationalDatabaseBundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseBundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) RelationalDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) RelationalDatabaseParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relationalDatabaseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) RotateMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotateMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Database`.
func NewCfnDatabase(scope constructs.Construct, id *string, props *CfnDatabaseProps) CfnDatabase {
	_init_.Initialize()

	j := jsiiProxy_CfnDatabase{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Database`.
func NewCfnDatabase_Override(c CfnDatabase, scope constructs.Construct, id *string, props *CfnDatabaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDatabase) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetBackupRetention(val interface{}) {
	_jsii_.Set(
		j,
		"backupRetention",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetCaCertificateIdentifier(val *string) {
	_jsii_.Set(
		j,
		"caCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetMasterDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"masterDatabaseName",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetMasterUsername(val *string) {
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetMasterUserPassword(val *string) {
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetPreferredBackupWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetPubliclyAccessible(val interface{}) {
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetRelationalDatabaseBlueprintId(val *string) {
	_jsii_.Set(
		j,
		"relationalDatabaseBlueprintId",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetRelationalDatabaseBundleId(val *string) {
	_jsii_.Set(
		j,
		"relationalDatabaseBundleId",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetRelationalDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"relationalDatabaseName",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetRelationalDatabaseParameters(val interface{}) {
	_jsii_.Set(
		j,
		"relationalDatabaseParameters",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetRotateMasterUserPassword(val interface{}) {
	_jsii_.Set(
		j,
		"rotateMasterUserPassword",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDatabase_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDatabase_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatabase_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDatabase) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDatabase) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDatabase) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDatabase) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDatabase) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDatabase) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDatabase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDatabase) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatabase) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatabase) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDatabase) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDatabase) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatabase) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatabase) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `RelationalDatabaseParameter` is a property of the [AWS::Lightsail::Database](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html) resource. It describes parameters for the database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relationalDatabaseParameterProperty := &relationalDatabaseParameterProperty{
//   	allowedValues: jsii.String("allowedValues"),
//   	applyMethod: jsii.String("applyMethod"),
//   	applyType: jsii.String("applyType"),
//   	dataType: jsii.String("dataType"),
//   	description: jsii.String("description"),
//   	isModifiable: jsii.Boolean(false),
//   	parameterName: jsii.String("parameterName"),
//   	parameterValue: jsii.String("parameterValue"),
//   }
//
type CfnDatabase_RelationalDatabaseParameterProperty struct {
	// The valid range of values for the parameter.
	AllowedValues *string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// Indicates when parameter updates are applied.
	//
	// Can be `immediate` or `pending-reboot` .
	ApplyMethod *string `field:"optional" json:"applyMethod" yaml:"applyMethod"`
	// Specifies the engine-specific parameter type.
	ApplyType *string `field:"optional" json:"applyType" yaml:"applyType"`
	// The valid data type of the parameter.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// A description of the parameter.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A Boolean value indicating whether the parameter can be modified.
	IsModifiable interface{} `field:"optional" json:"isModifiable" yaml:"isModifiable"`
	// The name of the parameter.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The value for the parameter.
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

// Properties for defining a `CfnDatabase`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatabaseProps := &cfnDatabaseProps{
//   	masterDatabaseName: jsii.String("masterDatabaseName"),
//   	masterUsername: jsii.String("masterUsername"),
//   	relationalDatabaseBlueprintId: jsii.String("relationalDatabaseBlueprintId"),
//   	relationalDatabaseBundleId: jsii.String("relationalDatabaseBundleId"),
//   	relationalDatabaseName: jsii.String("relationalDatabaseName"),
//
//   	// the properties below are optional
//   	availabilityZone: jsii.String("availabilityZone"),
//   	backupRetention: jsii.Boolean(false),
//   	caCertificateIdentifier: jsii.String("caCertificateIdentifier"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	relationalDatabaseParameters: []interface{}{
//   		&relationalDatabaseParameterProperty{
//   			allowedValues: jsii.String("allowedValues"),
//   			applyMethod: jsii.String("applyMethod"),
//   			applyType: jsii.String("applyType"),
//   			dataType: jsii.String("dataType"),
//   			description: jsii.String("description"),
//   			isModifiable: jsii.Boolean(false),
//   			parameterName: jsii.String("parameterName"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	rotateMasterUserPassword: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDatabaseProps struct {
	// The meaning of this parameter differs according to the database engine you use.
	//
	// *MySQL*
	//
	// The name of the database to create when the Lightsail database resource is created. If this parameter isn't specified, no database is created in the database resource.
	//
	// Constraints:
	//
	// - Must contain 1-64 letters or numbers.
	// - Must begin with a letter. Subsequent characters can be letters, underscores, or numbers (0-9).
	// - Can't be a word reserved by the specified database engine.
	//
	// For more information about reserved words in MySQL, see the Keywords and Reserved Words articles for [MySQL 5.6](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.6/en/keywords.html) , [MySQL 5.7](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.7/en/keywords.html) , and [MySQL 8.0](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/8.0/en/keywords.html) .
	//
	// *PostgreSQL*
	//
	// The name of the database to create when the Lightsail database resource is created. If this parameter isn't specified, a database named `postgres` is created in the database resource.
	//
	// Constraints:
	//
	// - Must contain 1-63 letters or numbers.
	// - Must begin with a letter. Subsequent characters can be letters, underscores, or numbers (0-9).
	// - Can't be a word reserved by the specified database engine.
	//
	// For more information about reserved words in PostgreSQL, see the SQL Key Words articles for [PostgreSQL 9.6](https://docs.aws.amazon.com/https://www.postgresql.org/docs/9.6/sql-keywords-appendix.html) , [PostgreSQL 10](https://docs.aws.amazon.com/https://www.postgresql.org/docs/10/sql-keywords-appendix.html) , [PostgreSQL 11](https://docs.aws.amazon.com/https://www.postgresql.org/docs/11/sql-keywords-appendix.html) , and [PostgreSQL 12](https://docs.aws.amazon.com/https://www.postgresql.org/docs/12/sql-keywords-appendix.html) .
	MasterDatabaseName *string `field:"required" json:"masterDatabaseName" yaml:"masterDatabaseName"`
	// The name for the primary user.
	//
	// *MySQL*
	//
	// Constraints:
	//
	// - Required for MySQL.
	// - Must be 1-16 letters or numbers. Can contain underscores.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// For more information about reserved words in MySQL 5.6 or 5.7, see the Keywords and Reserved Words articles for [MySQL 5.6](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.6/en/keywords.html) , [MySQL 5.7](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.7/en/keywords.html) , or [MySQL 8.0](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/8.0/en/keywords.html) .
	//
	// *PostgreSQL*
	//
	// Constraints:
	//
	// - Required for PostgreSQL.
	// - Must be 1-63 letters or numbers. Can contain underscores.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// For more information about reserved words in MySQL 5.6 or 5.7, see the Keywords and Reserved Words articles for [PostgreSQL 9.6](https://docs.aws.amazon.com/https://www.postgresql.org/docs/9.6/sql-keywords-appendix.html) , [PostgreSQL 10](https://docs.aws.amazon.com/https://www.postgresql.org/docs/10/sql-keywords-appendix.html) , [PostgreSQL 11](https://docs.aws.amazon.com/https://www.postgresql.org/docs/11/sql-keywords-appendix.html) , and [PostgreSQL 12](https://docs.aws.amazon.com/https://www.postgresql.org/docs/12/sql-keywords-appendix.html) .
	MasterUsername *string `field:"required" json:"masterUsername" yaml:"masterUsername"`
	// The blueprint ID for the database (for example, `mysql_8_0` ).
	RelationalDatabaseBlueprintId *string `field:"required" json:"relationalDatabaseBlueprintId" yaml:"relationalDatabaseBlueprintId"`
	// The bundle ID for the database (for example, `medium_1_0` ).
	RelationalDatabaseBundleId *string `field:"required" json:"relationalDatabaseBundleId" yaml:"relationalDatabaseBundleId"`
	// The name of the instance.
	RelationalDatabaseName *string `field:"required" json:"relationalDatabaseName" yaml:"relationalDatabaseName"`
	// The Availability Zone for the database.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// A Boolean value indicating whether automated backup retention is enabled for the database.
	BackupRetention interface{} `field:"optional" json:"backupRetention" yaml:"backupRetention"`
	// The certificate associated with the database.
	CaCertificateIdentifier *string `field:"optional" json:"caCertificateIdentifier" yaml:"caCertificateIdentifier"`
	// The password for the primary user of the database.
	//
	// The password can include any printable ASCII character except the following: /, ", or @. It cannot contain spaces.
	//
	// > The `MasterUserPassword` and `RotateMasterUserPassword` parameters cannot be used together in the same template.
	//
	// *MySQL*
	//
	// Constraints: Must contain 8-41 characters.
	//
	// *PostgreSQL*
	//
	// Constraints: Must contain 8-128 characters.
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// The daily time range during which automated backups are created for the database (for example, `16:00-16:30` ).
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur for the database, formatted as follows: `ddd:hh24:mi-ddd:hh24:mi` .
	//
	// For example, `Tue:17:00-Tue:17:30` .
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// A Boolean value indicating whether the database is accessible to anyone on the internet.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// An array of parameters for the database.
	RelationalDatabaseParameters interface{} `field:"optional" json:"relationalDatabaseParameters" yaml:"relationalDatabaseParameters"`
	// A Boolean value indicating whether to change the primary user password to a new, strong password generated by Lightsail .
	//
	// > The `RotateMasterUserPassword` and `MasterUserPassword` parameters cannot be used together in the same template.
	RotateMasterUserPassword interface{} `field:"optional" json:"rotateMasterUserPassword" yaml:"rotateMasterUserPassword"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Lightsail::Disk`.
//
// The `AWS::Lightsail::Disk` resource specifies a disk that can be attached to an Amazon Lightsail instance that is in the same AWS Region and Availability Zone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDisk := awscdk.Aws_lightsail.NewCfnDisk(this, jsii.String("MyCfnDisk"), &cfnDiskProps{
//   	diskName: jsii.String("diskName"),
//   	sizeInGb: jsii.Number(123),
//
//   	// the properties below are optional
//   	addOns: []interface{}{
//   		&addOnProperty{
//   			addOnType: jsii.String("addOnType"),
//
//   			// the properties below are optional
//   			autoSnapshotAddOnRequest: &autoSnapshotAddOnProperty{
//   				snapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   			},
//   			status: jsii.String("status"),
//   		},
//   	},
//   	availabilityZone: jsii.String("availabilityZone"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDisk interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// An array of add-ons for the disk.
	//
	// > If the disk has an add-on enabled when performing a delete disk request, the add-on is automatically disabled before the disk is deleted.
	AddOns() interface{}
	SetAddOns(val interface{})
	// The instance to which the disk is attached.
	AttrAttachedTo() *string
	// (Deprecated) The attachment state of the disk.
	//
	// > In releases prior to November 14, 2017, this parameter returned `attached` for system disks in the API response. It is now deprecated, but still included in the response. Use `isAttached` instead.
	AttrAttachmentState() *string
	// The Amazon Resource Name (ARN) of the disk.
	AttrDiskArn() *string
	// The input/output operations per second (IOPS) of the disk.
	AttrIops() *float64
	// A Boolean value indicating whether the disk is attached to an instance.
	AttrIsAttached() awscdk.IResolvable
	// The path of the disk.
	AttrPath() *string
	// The resource type of the disk (for example, `Disk` ).
	AttrResourceType() *string
	// The state of the disk (for example, `in-use` ).
	AttrState() *string
	// The support code of the disk.
	//
	// Include this code in your email to support when you have questions about a disk or another resource in Lightsail . This code helps our support team to look up your Lightsail information.
	AttrSupportCode() *string
	// The AWS Region and Availability Zone location for the disk (for example, `us-east-1a` ).
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the disk.
	DiskName() *string
	SetDiskName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The size of the disk in GB.
	SizeInGb() *float64
	SetSizeInGb(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDisk
type jsiiProxy_CfnDisk struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDisk) AddOns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrAttachedTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAttachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrAttachmentState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAttachmentState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrDiskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDiskArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrIsAttached() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrSupportCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSupportCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) DiskName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) SizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Disk`.
func NewCfnDisk(scope constructs.Construct, id *string, props *CfnDiskProps) CfnDisk {
	_init_.Initialize()

	j := jsiiProxy_CfnDisk{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Disk`.
func NewCfnDisk_Override(c CfnDisk, scope constructs.Construct, id *string, props *CfnDiskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDisk) SetAddOns(val interface{}) {
	_jsii_.Set(
		j,
		"addOns",
		val,
	)
}

func (j *jsiiProxy_CfnDisk) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnDisk) SetDiskName(val *string) {
	_jsii_.Set(
		j,
		"diskName",
		val,
	)
}

func (j *jsiiProxy_CfnDisk) SetSizeInGb(val *float64) {
	_jsii_.Set(
		j,
		"sizeInGb",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDisk_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDisk_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnDisk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDisk_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDisk) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDisk) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDisk) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDisk) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDisk) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDisk) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDisk) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDisk) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDisk) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDisk) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDisk) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDisk) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDisk) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDisk) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDisk) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `AddOn` is a property of the [AWS::Lightsail::Disk](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disk.html) resource. It describes the add-ons for a disk.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addOnProperty := &addOnProperty{
//   	addOnType: jsii.String("addOnType"),
//
//   	// the properties below are optional
//   	autoSnapshotAddOnRequest: &autoSnapshotAddOnProperty{
//   		snapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   	},
//   	status: jsii.String("status"),
//   }
//
type CfnDisk_AddOnProperty struct {
	// The add-on type (for example, `AutoSnapshot` ).
	//
	// > `AutoSnapshot` is the only add-on that can be enabled for a disk.
	AddOnType *string `field:"required" json:"addOnType" yaml:"addOnType"`
	// The parameters for the automatic snapshot add-on, such as the daily time when an automatic snapshot will be created.
	AutoSnapshotAddOnRequest interface{} `field:"optional" json:"autoSnapshotAddOnRequest" yaml:"autoSnapshotAddOnRequest"`
	// The status of the add-on.
	//
	// Valid Values: `Enabled` | `Disabled`.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

// `AutoSnapshotAddOn` is a property of the [AddOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-addon.html) property. It describes the automatic snapshot add-on for a disk.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoSnapshotAddOnProperty := &autoSnapshotAddOnProperty{
//   	snapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   }
//
type CfnDisk_AutoSnapshotAddOnProperty struct {
	// The daily time when an automatic snapshot will be created.
	//
	// Constraints:
	//
	// - Must be in `HH:00` format, and in an hourly increment.
	// - Specified in Coordinated Universal Time (UTC).
	// - The snapshot will be automatically created between the time specified and up to 45 minutes after.
	SnapshotTimeOfDay *string `field:"optional" json:"snapshotTimeOfDay" yaml:"snapshotTimeOfDay"`
}

// Properties for defining a `CfnDisk`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDiskProps := &cfnDiskProps{
//   	diskName: jsii.String("diskName"),
//   	sizeInGb: jsii.Number(123),
//
//   	// the properties below are optional
//   	addOns: []interface{}{
//   		&addOnProperty{
//   			addOnType: jsii.String("addOnType"),
//
//   			// the properties below are optional
//   			autoSnapshotAddOnRequest: &autoSnapshotAddOnProperty{
//   				snapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   			},
//   			status: jsii.String("status"),
//   		},
//   	},
//   	availabilityZone: jsii.String("availabilityZone"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDiskProps struct {
	// The name of the disk.
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// The size of the disk in GB.
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// An array of add-ons for the disk.
	//
	// > If the disk has an add-on enabled when performing a delete disk request, the add-on is automatically disabled before the disk is deleted.
	AddOns interface{} `field:"optional" json:"addOns" yaml:"addOns"`
	// The AWS Region and Availability Zone location for the disk (for example, `us-east-1a` ).
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Lightsail::Distribution`.
//
// The `AWS::Lightsail::Distribution` resource specifies a content delivery network (CDN) distribution. You can create distributions only in the `us-east-1` AWS Region.
//
// A distribution is a globally distributed network of caching servers that improve the performance of your website or web application hosted on a Lightsail instance, static content hosted on a Lightsail bucket, or through a Lightsail load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDistribution := awscdk.Aws_lightsail.NewCfnDistribution(this, jsii.String("MyCfnDistribution"), &cfnDistributionProps{
//   	bundleId: jsii.String("bundleId"),
//   	defaultCacheBehavior: &cacheBehaviorProperty{
//   		behavior: jsii.String("behavior"),
//   	},
//   	distributionName: jsii.String("distributionName"),
//   	origin: &inputOriginProperty{
//   		name: jsii.String("name"),
//   		protocolPolicy: jsii.String("protocolPolicy"),
//   		regionName: jsii.String("regionName"),
//   	},
//
//   	// the properties below are optional
//   	cacheBehaviors: []interface{}{
//   		&cacheBehaviorPerPathProperty{
//   			behavior: jsii.String("behavior"),
//   			path: jsii.String("path"),
//   		},
//   	},
//   	cacheBehaviorSettings: &cacheSettingsProperty{
//   		allowedHttpMethods: jsii.String("allowedHttpMethods"),
//   		cachedHttpMethods: jsii.String("cachedHttpMethods"),
//   		defaultTtl: jsii.Number(123),
//   		forwardedCookies: &cookieObjectProperty{
//   			cookiesAllowList: []*string{
//   				jsii.String("cookiesAllowList"),
//   			},
//   			option: jsii.String("option"),
//   		},
//   		forwardedHeaders: &headerObjectProperty{
//   			headersAllowList: []*string{
//   				jsii.String("headersAllowList"),
//   			},
//   			option: jsii.String("option"),
//   		},
//   		forwardedQueryStrings: &queryStringObjectProperty{
//   			option: jsii.Boolean(false),
//   			queryStringsAllowList: []*string{
//   				jsii.String("queryStringsAllowList"),
//   			},
//   		},
//   		maximumTtl: jsii.Number(123),
//   		minimumTtl: jsii.Number(123),
//   	},
//   	certificateName: jsii.String("certificateName"),
//   	ipAddressType: jsii.String("ipAddressType"),
//   	isEnabled: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDistribution interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Indicates whether you can update the distribution’s current bundle to another bundle.
	AttrAbleToUpdateBundle() awscdk.IResolvable
	// The Amazon Resource Name (ARN) of the distribution.
	AttrDistributionArn() *string
	// The status of the distribution.
	AttrStatus() *string
	// The ID of the bundle applied to the distribution.
	BundleId() *string
	SetBundleId(val *string)
	// An array of objects that describe the per-path cache behavior of the distribution.
	CacheBehaviors() interface{}
	SetCacheBehaviors(val interface{})
	// An object that describes the cache behavior settings of the distribution.
	CacheBehaviorSettings() interface{}
	SetCacheBehaviorSettings(val interface{})
	// The name of the SSL/TLS certificate attached to the distribution.
	CertificateName() *string
	SetCertificateName(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An object that describes the default cache behavior of the distribution.
	DefaultCacheBehavior() interface{}
	SetDefaultCacheBehavior(val interface{})
	// The name of the distribution.
	DistributionName() *string
	SetDistributionName(val *string)
	// The IP address type of the distribution.
	//
	// The possible values are `ipv4` for IPv4 only, and `dualstack` for IPv4 and IPv6.
	IpAddressType() *string
	SetIpAddressType(val *string)
	// A Boolean value indicating whether the distribution is enabled.
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// An object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer.
	//
	// The distribution pulls, caches, and serves content from the origin.
	Origin() interface{}
	SetOrigin(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDistribution
type jsiiProxy_CfnDistribution struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDistribution) AttrAbleToUpdateBundle() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrAbleToUpdateBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) AttrDistributionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDistributionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) CacheBehaviors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheBehaviors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) CacheBehaviorSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheBehaviorSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) DefaultCacheBehavior() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) DistributionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) Origin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Distribution`.
func NewCfnDistribution(scope constructs.Construct, id *string, props *CfnDistributionProps) CfnDistribution {
	_init_.Initialize()

	j := jsiiProxy_CfnDistribution{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDistribution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Distribution`.
func NewCfnDistribution_Override(c CfnDistribution, scope constructs.Construct, id *string, props *CfnDistributionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDistribution",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDistribution) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_CfnDistribution) SetCacheBehaviors(val interface{}) {
	_jsii_.Set(
		j,
		"cacheBehaviors",
		val,
	)
}

func (j *jsiiProxy_CfnDistribution) SetCacheBehaviorSettings(val interface{}) {
	_jsii_.Set(
		j,
		"cacheBehaviorSettings",
		val,
	)
}

func (j *jsiiProxy_CfnDistribution) SetCertificateName(val *string) {
	_jsii_.Set(
		j,
		"certificateName",
		val,
	)
}

func (j *jsiiProxy_CfnDistribution) SetDefaultCacheBehavior(val interface{}) {
	_jsii_.Set(
		j,
		"defaultCacheBehavior",
		val,
	)
}

func (j *jsiiProxy_CfnDistribution) SetDistributionName(val *string) {
	_jsii_.Set(
		j,
		"distributionName",
		val,
	)
}

func (j *jsiiProxy_CfnDistribution) SetIpAddressType(val *string) {
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_CfnDistribution) SetIsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnDistribution) SetOrigin(val interface{}) {
	_jsii_.Set(
		j,
		"origin",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDistribution_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDistribution",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDistribution_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDistribution",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDistribution_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnDistribution",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDistribution) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDistribution) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDistribution) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDistribution) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDistribution) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDistribution) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDistribution) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDistribution) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistribution) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistribution) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDistribution) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDistribution) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistribution) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistribution) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `CacheBehaviorPerPath` is a property of the [AWS::Lightsail::Distribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html) resource. It describes the per-path cache behavior of an Amazon Lightsail content delivery network (CDN) distribution.
//
// Use a per-path cache behavior to override the default cache behavior of a distribution, or to add an exception to it. For example, if you set the `CacheBehavior` to `cache` , you can use a per-path cache behavior to specify a directory, file, or file type that your distribution will cache. If you don’t want your distribution to cache a specified directory, file, or file type, set the per-path cache behavior to `dont-cache` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cacheBehaviorPerPathProperty := &cacheBehaviorPerPathProperty{
//   	behavior: jsii.String("behavior"),
//   	path: jsii.String("path"),
//   }
//
type CfnDistribution_CacheBehaviorPerPathProperty struct {
	// The cache behavior for the specified path.
	//
	// You can specify one of the following per-path cache behaviors:
	//
	// - *`cache`* - This behavior caches the specified path.
	// - *`dont-cache`* - This behavior doesn't cache the specified path.
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// The path to a directory or file to cache, or not cache.
	//
	// Use an asterisk symbol to specify wildcard directories ( `path/to/assets/*` ), and file types ( `*.html` , `*jpg` , `*js` ). Directories and file paths are case-sensitive.
	//
	// Examples:
	//
	// - Specify the following to cache all files in the document root of an Apache web server running on a instance.
	//
	// `var/www/html/`
	// - Specify the following file to cache only the index page in the document root of an Apache web server.
	//
	// `var/www/html/index.html`
	// - Specify the following to cache only the .html files in the document root of an Apache web server.
	//
	// `var/www/html/*.html`
	// - Specify the following to cache only the .jpg, .png, and .gif files in the images sub-directory of the document root of an Apache web server.
	//
	// `var/www/html/images/*.jpg`
	//
	// `var/www/html/images/*.png`
	//
	// `var/www/html/images/*.gif`
	//
	// Specify the following to cache all files in the images subdirectory of the document root of an Apache web server.
	//
	// `var/www/html/images/`.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

// `CacheBehavior` is a property of the [AWS::Lightsail::Distribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html) resource. It describes the default cache behavior of an Amazon Lightsail content delivery network (CDN) distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cacheBehaviorProperty := &cacheBehaviorProperty{
//   	behavior: jsii.String("behavior"),
//   }
//
type CfnDistribution_CacheBehaviorProperty struct {
	// The cache behavior of the distribution.
	//
	// The following cache behaviors can be specified:
	//
	// - *`cache`* - This option is best for static sites. When specified, your distribution caches and serves your entire website as static content. This behavior is ideal for websites with static content that doesn't change depending on who views it, or for websites that don't use cookies, headers, or query strings to personalize content.
	// - *`dont-cache`* - This option is best for sites that serve a mix of static and dynamic content. When specified, your distribution caches and serves only the content that is specified in the distribution’s `CacheBehaviorPerPath` parameter. This behavior is ideal for websites or web applications that use cookies, headers, and query strings to personalize content for individual users.
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
}

// `CacheSettings` is a property of the [AWS::Lightsail::Distribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html) resource. It describes the cache settings of an Amazon Lightsail content delivery network (CDN) distribution.
//
// These settings apply only to your distribution’s `CacheBehaviors` that have a `Behavior` of `cache` . This includes the `DefaultCacheBehavior` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cacheSettingsProperty := &cacheSettingsProperty{
//   	allowedHttpMethods: jsii.String("allowedHttpMethods"),
//   	cachedHttpMethods: jsii.String("cachedHttpMethods"),
//   	defaultTtl: jsii.Number(123),
//   	forwardedCookies: &cookieObjectProperty{
//   		cookiesAllowList: []*string{
//   			jsii.String("cookiesAllowList"),
//   		},
//   		option: jsii.String("option"),
//   	},
//   	forwardedHeaders: &headerObjectProperty{
//   		headersAllowList: []*string{
//   			jsii.String("headersAllowList"),
//   		},
//   		option: jsii.String("option"),
//   	},
//   	forwardedQueryStrings: &queryStringObjectProperty{
//   		option: jsii.Boolean(false),
//   		queryStringsAllowList: []*string{
//   			jsii.String("queryStringsAllowList"),
//   		},
//   	},
//   	maximumTtl: jsii.Number(123),
//   	minimumTtl: jsii.Number(123),
//   }
//
type CfnDistribution_CacheSettingsProperty struct {
	// The HTTP methods that are processed and forwarded to the distribution's origin.
	//
	// You can specify the following options:
	//
	// - `GET,HEAD` - The distribution forwards the `GET` and `HEAD` methods.
	// - `GET,HEAD,OPTIONS` - The distribution forwards the `GET` , `HEAD` , and `OPTIONS` methods.
	// - `GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE` - The distribution forwards the `GET` , `HEAD` , `OPTIONS` , `PUT` , `PATCH` , `POST` , and `DELETE` methods.
	//
	// If you specify `GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE` , you might need to restrict access to your distribution's origin so users can't perform operations that you don't want them to. For example, you might not want users to have permission to delete objects from your origin.
	AllowedHttpMethods *string `field:"optional" json:"allowedHttpMethods" yaml:"allowedHttpMethods"`
	// The HTTP method responses that are cached by your distribution.
	//
	// You can specify the following options:
	//
	// - `GET,HEAD` - The distribution caches responses to the `GET` and `HEAD` methods.
	// - `GET,HEAD,OPTIONS` - The distribution caches responses to the `GET` , `HEAD` , and `OPTIONS` methods.
	CachedHttpMethods *string `field:"optional" json:"cachedHttpMethods" yaml:"cachedHttpMethods"`
	// The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.
	//
	// > The value specified applies only when the origin does not add HTTP headers such as `Cache-Control max-age` , `Cache-Control s-maxage` , and `Expires` to objects.
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// An object that describes the cookies that are forwarded to the origin.
	//
	// Your content is cached based on the cookies that are forwarded.
	ForwardedCookies interface{} `field:"optional" json:"forwardedCookies" yaml:"forwardedCookies"`
	// An object that describes the headers that are forwarded to the origin.
	//
	// Your content is cached based on the headers that are forwarded.
	ForwardedHeaders interface{} `field:"optional" json:"forwardedHeaders" yaml:"forwardedHeaders"`
	// An object that describes the query strings that are forwarded to the origin.
	//
	// Your content is cached based on the query strings that are forwarded.
	ForwardedQueryStrings interface{} `field:"optional" json:"forwardedQueryStrings" yaml:"forwardedQueryStrings"`
	// The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	//
	// The value specified applies only when the origin adds HTTP headers such as `Cache-Control max-age` , `Cache-Control s-maxage` , and `Expires` to objects.
	MaximumTtl *float64 `field:"optional" json:"maximumTtl" yaml:"maximumTtl"`
	// The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	//
	// A value of `0` must be specified for `minimumTTL` if the distribution is configured to forward all headers to the origin.
	MinimumTtl *float64 `field:"optional" json:"minimumTtl" yaml:"minimumTtl"`
}

// `CookieObject` is a property of the [CacheSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-cachesettings.html) property. It describes whether an Amazon Lightsail content delivery network (CDN) distribution forwards cookies to the origin and, if so, which ones.
//
// For the cookies that you specify, your distribution caches separate versions of the specified content based on the cookie values in viewer requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cookieObjectProperty := &cookieObjectProperty{
//   	cookiesAllowList: []*string{
//   		jsii.String("cookiesAllowList"),
//   	},
//   	option: jsii.String("option"),
//   }
//
type CfnDistribution_CookieObjectProperty struct {
	// The specific cookies to forward to your distribution's origin.
	CookiesAllowList *[]*string `field:"optional" json:"cookiesAllowList" yaml:"cookiesAllowList"`
	// Specifies which cookies to forward to the distribution's origin for a cache behavior.
	//
	// Use one of the following configurations for your distribution:
	//
	// - *`all`* - Forwards all cookies to your origin.
	// - *`none`* - Doesn’t forward cookies to your origin.
	// - *`allow-list`* - Forwards only the cookies that you specify using the `CookiesAllowList` parameter.
	Option *string `field:"optional" json:"option" yaml:"option"`
}

// `HeaderObject` is a property of the [CacheSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-cachesettings.html) property. It describes the request headers used by your distribution, which caches your content based on the request headers.
//
// For the headers that you specify, your distribution caches separate versions of the specified content based on the header values in viewer requests. For example, suppose that viewer requests for logo.jpg contain a custom product header that has a value of either acme or apex. Also, suppose that you configure your distribution to cache your content based on values in the product header. Your distribution forwards the product header to the origin and caches the response from the origin once for each header value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerObjectProperty := &headerObjectProperty{
//   	headersAllowList: []*string{
//   		jsii.String("headersAllowList"),
//   	},
//   	option: jsii.String("option"),
//   }
//
type CfnDistribution_HeaderObjectProperty struct {
	// The specific headers to forward to your distribution's origin.
	HeadersAllowList *[]*string `field:"optional" json:"headersAllowList" yaml:"headersAllowList"`
	// The headers that you want your distribution to forward to your origin.
	//
	// Your distribution caches your content based on these headers.
	//
	// Use one of the following configurations for your distribution:
	//
	// - *`all`* - Forwards all headers to your origin..
	// - *`none`* - Forwards only the default headers.
	// - *`allow-list`* - Forwards only the headers that you specify using the `HeadersAllowList` parameter.
	Option *string `field:"optional" json:"option" yaml:"option"`
}

// `InputOrigin` is a property of the [AWS::Lightsail::Distribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html) resource. It describes the origin resource of an Amazon Lightsail content delivery network (CDN) distribution.
//
// An origin can be a instance, bucket, or load balancer. A distribution pulls content from an origin, caches it, and serves it to viewers through a worldwide network of edge servers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputOriginProperty := &inputOriginProperty{
//   	name: jsii.String("name"),
//   	protocolPolicy: jsii.String("protocolPolicy"),
//   	regionName: jsii.String("regionName"),
//   }
//
type CfnDistribution_InputOriginProperty struct {
	// The name of the origin resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.
	ProtocolPolicy *string `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
	// The AWS Region name of the origin resource.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

// `QueryStringObject` is a property of the [CacheSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-cachesettings.html) property. It describes the query string parameters that an Amazon Lightsail content delivery network (CDN) distribution to bases caching on.
//
// For the query strings that you specify, your distribution caches separate versions of the specified content based on the query string values in viewer requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryStringObjectProperty := &queryStringObjectProperty{
//   	option: jsii.Boolean(false),
//   	queryStringsAllowList: []*string{
//   		jsii.String("queryStringsAllowList"),
//   	},
//   }
//
type CfnDistribution_QueryStringObjectProperty struct {
	// Indicates whether the distribution forwards and caches based on query strings.
	Option interface{} `field:"optional" json:"option" yaml:"option"`
	// The specific query strings that the distribution forwards to the origin.
	//
	// Your distribution caches content based on the specified query strings.
	//
	// If the `option` parameter is true, then your distribution forwards all query strings, regardless of what you specify using the `QueryStringsAllowList` parameter.
	QueryStringsAllowList *[]*string `field:"optional" json:"queryStringsAllowList" yaml:"queryStringsAllowList"`
}

// Properties for defining a `CfnDistribution`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDistributionProps := &cfnDistributionProps{
//   	bundleId: jsii.String("bundleId"),
//   	defaultCacheBehavior: &cacheBehaviorProperty{
//   		behavior: jsii.String("behavior"),
//   	},
//   	distributionName: jsii.String("distributionName"),
//   	origin: &inputOriginProperty{
//   		name: jsii.String("name"),
//   		protocolPolicy: jsii.String("protocolPolicy"),
//   		regionName: jsii.String("regionName"),
//   	},
//
//   	// the properties below are optional
//   	cacheBehaviors: []interface{}{
//   		&cacheBehaviorPerPathProperty{
//   			behavior: jsii.String("behavior"),
//   			path: jsii.String("path"),
//   		},
//   	},
//   	cacheBehaviorSettings: &cacheSettingsProperty{
//   		allowedHttpMethods: jsii.String("allowedHttpMethods"),
//   		cachedHttpMethods: jsii.String("cachedHttpMethods"),
//   		defaultTtl: jsii.Number(123),
//   		forwardedCookies: &cookieObjectProperty{
//   			cookiesAllowList: []*string{
//   				jsii.String("cookiesAllowList"),
//   			},
//   			option: jsii.String("option"),
//   		},
//   		forwardedHeaders: &headerObjectProperty{
//   			headersAllowList: []*string{
//   				jsii.String("headersAllowList"),
//   			},
//   			option: jsii.String("option"),
//   		},
//   		forwardedQueryStrings: &queryStringObjectProperty{
//   			option: jsii.Boolean(false),
//   			queryStringsAllowList: []*string{
//   				jsii.String("queryStringsAllowList"),
//   			},
//   		},
//   		maximumTtl: jsii.Number(123),
//   		minimumTtl: jsii.Number(123),
//   	},
//   	certificateName: jsii.String("certificateName"),
//   	ipAddressType: jsii.String("ipAddressType"),
//   	isEnabled: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDistributionProps struct {
	// The ID of the bundle applied to the distribution.
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// An object that describes the default cache behavior of the distribution.
	DefaultCacheBehavior interface{} `field:"required" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// The name of the distribution.
	DistributionName *string `field:"required" json:"distributionName" yaml:"distributionName"`
	// An object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer.
	//
	// The distribution pulls, caches, and serves content from the origin.
	Origin interface{} `field:"required" json:"origin" yaml:"origin"`
	// An array of objects that describe the per-path cache behavior of the distribution.
	CacheBehaviors interface{} `field:"optional" json:"cacheBehaviors" yaml:"cacheBehaviors"`
	// An object that describes the cache behavior settings of the distribution.
	CacheBehaviorSettings interface{} `field:"optional" json:"cacheBehaviorSettings" yaml:"cacheBehaviorSettings"`
	// The name of the SSL/TLS certificate attached to the distribution.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// The IP address type of the distribution.
	//
	// The possible values are `ipv4` for IPv4 only, and `dualstack` for IPv4 and IPv6.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// A Boolean value indicating whether the distribution is enabled.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Lightsail::Instance`.
//
// The `AWS::Lightsail::Instance` resource specifies an Amazon Lightsail instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstance := awscdk.Aws_lightsail.NewCfnInstance(this, jsii.String("MyCfnInstance"), &cfnInstanceProps{
//   	blueprintId: jsii.String("blueprintId"),
//   	bundleId: jsii.String("bundleId"),
//   	instanceName: jsii.String("instanceName"),
//
//   	// the properties below are optional
//   	addOns: []interface{}{
//   		&addOnProperty{
//   			addOnType: jsii.String("addOnType"),
//
//   			// the properties below are optional
//   			autoSnapshotAddOnRequest: &autoSnapshotAddOnProperty{
//   				snapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   			},
//   			status: jsii.String("status"),
//   		},
//   	},
//   	availabilityZone: jsii.String("availabilityZone"),
//   	hardware: &hardwareProperty{
//   		cpuCount: jsii.Number(123),
//   		disks: []interface{}{
//   			&diskProperty{
//   				diskName: jsii.String("diskName"),
//   				path: jsii.String("path"),
//
//   				// the properties below are optional
//   				attachedTo: jsii.String("attachedTo"),
//   				attachmentState: jsii.String("attachmentState"),
//   				iops: jsii.Number(123),
//   				isSystemDisk: jsii.Boolean(false),
//   				sizeInGb: jsii.String("sizeInGb"),
//   			},
//   		},
//   		ramSizeInGb: jsii.Number(123),
//   	},
//   	keyPairName: jsii.String("keyPairName"),
//   	networking: &networkingProperty{
//   		ports: []interface{}{
//   			&portProperty{
//   				accessDirection: jsii.String("accessDirection"),
//   				accessFrom: jsii.String("accessFrom"),
//   				accessType: jsii.String("accessType"),
//   				cidrListAliases: []*string{
//   					jsii.String("cidrListAliases"),
//   				},
//   				cidrs: []*string{
//   					jsii.String("cidrs"),
//   				},
//   				commonName: jsii.String("commonName"),
//   				fromPort: jsii.Number(123),
//   				ipv6Cidrs: []*string{
//   					jsii.String("ipv6Cidrs"),
//   				},
//   				protocol: jsii.String("protocol"),
//   				toPort: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		monthlyTransfer: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userData: jsii.String("userData"),
//   })
//
type CfnInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// An array of add-ons for the instance.
	//
	// > If the instance has an add-on enabled when performing a delete instance request, the add-on is automatically disabled before the instance is deleted.
	AddOns() interface{}
	SetAddOns(val interface{})
	// The number of vCPUs the instance has.
	AttrHardwareCpuCount() *float64
	// The amount of RAM in GB on the instance (for example, `1.0` ).
	AttrHardwareRamSizeInGb() *float64
	// The Amazon Resource Name (ARN) of the instance (for example, `arn:aws:lightsail:us-east-2:123456789101:Instance/244ad76f-8aad-4741-809f-12345EXAMPLE` ).
	AttrInstanceArn() *string
	// A Boolean value indicating whether the instance has a static IP assigned to it.
	AttrIsStaticIp() awscdk.IResolvable
	// The AWS Region and Availability Zone where the instance is located.
	AttrLocationAvailabilityZone() *string
	// The AWS Region of the instance.
	AttrLocationRegionName() *string
	// The amount of allocated monthly data transfer (in GB) for an instance.
	AttrNetworkingMonthlyTransferGbPerMonthAllocated() *string
	// The private IP address of the instance.
	AttrPrivateIpAddress() *string
	// The public IP address of the instance.
	AttrPublicIpAddress() *string
	// The resource type of the instance (for example, `Instance` ).
	AttrResourceType() *string
	// The name of the SSH key pair used by the instance.
	AttrSshKeyName() *string
	// The status code of the instance.
	AttrStateCode() *float64
	// The state of the instance (for example, `running` or `pending` ).
	AttrStateName() *string
	// The support code of the instance.
	//
	// Include this code in your email to support when you have questions about an instance or another resource in Lightsail . This code helps our support team to look up your Lightsail information.
	AttrSupportCode() *string
	// The user name for connecting to the instance (for example, `ec2-user` ).
	AttrUserName() *string
	// The Availability Zone for the instance.
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// The blueprint ID for the instance (for example, `os_amlinux_2016_03` ).
	BlueprintId() *string
	SetBlueprintId(val *string)
	// The bundle ID for the instance (for example, `micro_1_0` ).
	BundleId() *string
	SetBundleId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The hardware properties for the instance, such as the vCPU count, attached disks, and amount of RAM.
	//
	// > The instance restarts when performing an attach disk or detach disk request. This resets the public IP address of your instance if a static IP isn't attached to it.
	Hardware() interface{}
	SetHardware(val interface{})
	// The name of the instance.
	InstanceName() *string
	SetInstanceName(val *string)
	// The name of the key pair to use for the instance.
	//
	// If no key pair name is specified, the Regional Lightsail default key pair is used.
	KeyPairName() *string
	SetKeyPairName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The public ports and the monthly amount of data transfer allocated for the instance.
	Networking() interface{}
	SetNetworking(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The optional launch script for the instance.
	//
	// Specify a launch script to configure an instance with additional user data. For example, you might want to specify `apt-get -y update` as a launch script.
	//
	// > Depending on the blueprint of your instance, the command to get software on your instance varies. Amazon Linux and CentOS use `yum` , Debian and Ubuntu use `apt-get` , and FreeBSD uses `pkg` .
	UserData() *string
	SetUserData(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnInstance
type jsiiProxy_CfnInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInstance) AddOns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrHardwareCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrHardwareCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrHardwareRamSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrHardwareRamSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrIsStaticIp() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsStaticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrLocationAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrLocationRegionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationRegionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrNetworkingMonthlyTransferGbPerMonthAllocated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNetworkingMonthlyTransferGbPerMonthAllocated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPrivateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrSshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrStateCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrStateCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrStateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrSupportCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSupportCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) BlueprintId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueprintId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Hardware() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardware",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) InstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) KeyPairName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPairName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Networking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Instance`.
func NewCfnInstance(scope constructs.Construct, id *string, props *CfnInstanceProps) CfnInstance {
	_init_.Initialize()

	j := jsiiProxy_CfnInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Instance`.
func NewCfnInstance_Override(c CfnInstance, scope constructs.Construct, id *string, props *CfnInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstance) SetAddOns(val interface{}) {
	_jsii_.Set(
		j,
		"addOns",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetBlueprintId(val *string) {
	_jsii_.Set(
		j,
		"blueprintId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetHardware(val interface{}) {
	_jsii_.Set(
		j,
		"hardware",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetInstanceName(val *string) {
	_jsii_.Set(
		j,
		"instanceName",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetKeyPairName(val *string) {
	_jsii_.Set(
		j,
		"keyPairName",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetNetworking(val interface{}) {
	_jsii_.Set(
		j,
		"networking",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetUserData(val *string) {
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstance) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInstance) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInstance) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInstance) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInstance) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstance) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstance) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstance) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `AddOn` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the add-ons for an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addOnProperty := &addOnProperty{
//   	addOnType: jsii.String("addOnType"),
//
//   	// the properties below are optional
//   	autoSnapshotAddOnRequest: &autoSnapshotAddOnProperty{
//   		snapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   	},
//   	status: jsii.String("status"),
//   }
//
type CfnInstance_AddOnProperty struct {
	// The add-on type (for example, `AutoSnapshot` ).
	//
	// > `AutoSnapshot` is the only add-on that can be enabled for an instance.
	AddOnType *string `field:"required" json:"addOnType" yaml:"addOnType"`
	// The parameters for the automatic snapshot add-on, such as the daily time when an automatic snapshot will be created.
	AutoSnapshotAddOnRequest interface{} `field:"optional" json:"autoSnapshotAddOnRequest" yaml:"autoSnapshotAddOnRequest"`
	// The status of the add-on.
	//
	// Valid Values: `Enabled` | `Disabled`.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

// `AutoSnapshotAddOn` is a property of the [AddOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-addon.html) property. It describes the automatic snapshot add-on for an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoSnapshotAddOnProperty := &autoSnapshotAddOnProperty{
//   	snapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   }
//
type CfnInstance_AutoSnapshotAddOnProperty struct {
	// The daily time when an automatic snapshot will be created.
	//
	// Constraints:
	//
	// - Must be in `HH:00` format, and in an hourly increment.
	// - Specified in Coordinated Universal Time (UTC).
	// - The snapshot will be automatically created between the time specified and up to 45 minutes after.
	SnapshotTimeOfDay *string `field:"optional" json:"snapshotTimeOfDay" yaml:"snapshotTimeOfDay"`
}

// `Disk` is a property of the [Hardware](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-hardware.html) property. It describes a disk attached to an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   diskProperty := &diskProperty{
//   	diskName: jsii.String("diskName"),
//   	path: jsii.String("path"),
//
//   	// the properties below are optional
//   	attachedTo: jsii.String("attachedTo"),
//   	attachmentState: jsii.String("attachmentState"),
//   	iops: jsii.Number(123),
//   	isSystemDisk: jsii.Boolean(false),
//   	sizeInGb: jsii.String("sizeInGb"),
//   }
//
type CfnInstance_DiskProperty struct {
	// The unique name of the disk.
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// The disk path.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The resources to which the disk is attached.
	AttachedTo *string `field:"optional" json:"attachedTo" yaml:"attachedTo"`
	// (Deprecated) The attachment state of the disk.
	//
	// > In releases prior to November 14, 2017, this parameter returned `attached` for system disks in the API response. It is now deprecated, but still included in the response. Use `isAttached` instead.
	AttachmentState *string `field:"optional" json:"attachmentState" yaml:"attachmentState"`
	// The input/output operations per second (IOPS) of the disk.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// A Boolean value indicating whether this disk is a system disk (has an operating system loaded on it).
	IsSystemDisk interface{} `field:"optional" json:"isSystemDisk" yaml:"isSystemDisk"`
	// The size of the disk in GB.
	SizeInGb *string `field:"optional" json:"sizeInGb" yaml:"sizeInGb"`
}

// `Hardware` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the hardware properties for the instance, such as the vCPU count, attached disks, and amount of RAM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hardwareProperty := &hardwareProperty{
//   	cpuCount: jsii.Number(123),
//   	disks: []interface{}{
//   		&diskProperty{
//   			diskName: jsii.String("diskName"),
//   			path: jsii.String("path"),
//
//   			// the properties below are optional
//   			attachedTo: jsii.String("attachedTo"),
//   			attachmentState: jsii.String("attachmentState"),
//   			iops: jsii.Number(123),
//   			isSystemDisk: jsii.Boolean(false),
//   			sizeInGb: jsii.String("sizeInGb"),
//   		},
//   	},
//   	ramSizeInGb: jsii.Number(123),
//   }
//
type CfnInstance_HardwareProperty struct {
	// The number of vCPUs the instance has.
	//
	// > The `CpuCount` property is read-only and should not be specified in a create instance or update instance request.
	CpuCount *float64 `field:"optional" json:"cpuCount" yaml:"cpuCount"`
	// The disks attached to the instance.
	//
	// The instance restarts when performing an attach disk or detach disk request. This resets the public IP address of your instance if a static IP isn't attached to it.
	Disks interface{} `field:"optional" json:"disks" yaml:"disks"`
	// The amount of RAM in GB on the instance (for example, `1.0` ).
	//
	// > The `RamSizeInGb` property is read-only and should not be specified in a create instance or update instance request.
	RamSizeInGb *float64 `field:"optional" json:"ramSizeInGb" yaml:"ramSizeInGb"`
}

// `Location` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the location for an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &locationProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	regionName: jsii.String("regionName"),
//   }
//
type CfnInstance_LocationProperty struct {
	// The Availability Zone for the instance.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The name of the AWS Region for the instance.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

// `MonthlyTransfer` is a property of the [Networking](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-networking.html) property. It describes the amount of allocated monthly data transfer (in GB) for an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monthlyTransferProperty := &monthlyTransferProperty{
//   	gbPerMonthAllocated: jsii.String("gbPerMonthAllocated"),
//   }
//
type CfnInstance_MonthlyTransferProperty struct {
	// The amount of allocated monthly data transfer (in GB) for an instance.
	GbPerMonthAllocated *string `field:"optional" json:"gbPerMonthAllocated" yaml:"gbPerMonthAllocated"`
}

// `Networking` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the public ports and the monthly amount of data transfer allocated for the instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkingProperty := &networkingProperty{
//   	ports: []interface{}{
//   		&portProperty{
//   			accessDirection: jsii.String("accessDirection"),
//   			accessFrom: jsii.String("accessFrom"),
//   			accessType: jsii.String("accessType"),
//   			cidrListAliases: []*string{
//   				jsii.String("cidrListAliases"),
//   			},
//   			cidrs: []*string{
//   				jsii.String("cidrs"),
//   			},
//   			commonName: jsii.String("commonName"),
//   			fromPort: jsii.Number(123),
//   			ipv6Cidrs: []*string{
//   				jsii.String("ipv6Cidrs"),
//   			},
//   			protocol: jsii.String("protocol"),
//   			toPort: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	monthlyTransfer: jsii.Number(123),
//   }
//
type CfnInstance_NetworkingProperty struct {
	// An array of ports to open on the instance.
	Ports interface{} `field:"required" json:"ports" yaml:"ports"`
	// The monthly amount of data transfer, in GB, allocated for the instance.
	MonthlyTransfer *float64 `field:"optional" json:"monthlyTransfer" yaml:"monthlyTransfer"`
}

// `Port` is a property of the [Networking](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-networking.html) property. It describes information about ports for an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portProperty := &portProperty{
//   	accessDirection: jsii.String("accessDirection"),
//   	accessFrom: jsii.String("accessFrom"),
//   	accessType: jsii.String("accessType"),
//   	cidrListAliases: []*string{
//   		jsii.String("cidrListAliases"),
//   	},
//   	cidrs: []*string{
//   		jsii.String("cidrs"),
//   	},
//   	commonName: jsii.String("commonName"),
//   	fromPort: jsii.Number(123),
//   	ipv6Cidrs: []*string{
//   		jsii.String("ipv6Cidrs"),
//   	},
//   	protocol: jsii.String("protocol"),
//   	toPort: jsii.Number(123),
//   }
//
type CfnInstance_PortProperty struct {
	// The access direction ( `inbound` or `outbound` ).
	//
	// > Lightsail currently supports only `inbound` access direction.
	AccessDirection *string `field:"optional" json:"accessDirection" yaml:"accessDirection"`
	// The location from which access is allowed.
	//
	// For example, `Anywhere (0.0.0.0/0)` , or `Custom` if a specific IP address or range of IP addresses is allowed.
	AccessFrom *string `field:"optional" json:"accessFrom" yaml:"accessFrom"`
	// The type of access ( `Public` or `Private` ).
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// An alias that defines access for a preconfigured range of IP addresses.
	//
	// The only alias currently supported is `lightsail-connect` , which allows IP addresses of the browser-based RDP/SSH client in the Lightsail console to connect to your instance.
	CidrListAliases *[]*string `field:"optional" json:"cidrListAliases" yaml:"cidrListAliases"`
	// The IPv4 address, or range of IPv4 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol.
	//
	// > The `ipv6Cidrs` parameter lists the IPv6 addresses that are allowed to connect to an instance.
	//
	// Examples:
	//
	// - To allow the IP address `192.0.2.44` , specify `192.0.2.44` or `192.0.2.44/32` .
	// - To allow the IP addresses `192.0.2.0` to `192.0.2.255` , specify `192.0.2.0/24` .
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
	// The common name of the port information.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// The first port in a range of open ports on an instance.
	//
	// Allowed ports:
	//
	// - TCP and UDP - `0` to `65535`
	// - ICMP - The ICMP type for IPv4 addresses. For example, specify `8` as the `fromPort` (ICMP type), and `-1` as the `toPort` (ICMP code), to enable ICMP Ping.
	// - ICMPv6 - The ICMP type for IPv6 addresses. For example, specify `128` as the `fromPort` (ICMPv6 type), and `0` as `toPort` (ICMPv6 code).
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The IPv6 address, or range of IPv6 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol.
	//
	// Only devices with an IPv6 address can connect to an instance through IPv6; otherwise, IPv4 should be used.
	//
	// > The `cidrs` parameter lists the IPv4 addresses that are allowed to connect to an instance.
	Ipv6Cidrs *[]*string `field:"optional" json:"ipv6Cidrs" yaml:"ipv6Cidrs"`
	// The IP protocol name.
	//
	// The name can be one of the following:
	//
	// - `tcp` - Transmission Control Protocol (TCP) provides reliable, ordered, and error-checked delivery of streamed data between applications running on hosts communicating by an IP network. If you have an application that doesn't require reliable data stream service, use UDP instead.
	// - `all` - All transport layer protocol types.
	// - `udp` - With User Datagram Protocol (UDP), computer applications can send messages (or datagrams) to other hosts on an Internet Protocol (IP) network. Prior communications are not required to set up transmission channels or data paths. Applications that don't require reliable data stream service can use UDP, which provides a connectionless datagram service that emphasizes reduced latency over reliability. If you do require reliable data stream service, use TCP instead.
	// - `icmp` - Internet Control Message Protocol (ICMP) is used to send error messages and operational information indicating success or failure when communicating with an instance. For example, an error is indicated when an instance could not be reached. When you specify `icmp` as the `protocol` , you must specify the ICMP type using the `fromPort` parameter, and ICMP code using the `toPort` parameter.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The last port in a range of open ports on an instance.
	//
	// Allowed ports:
	//
	// - TCP and UDP - `0` to `65535`
	// - ICMP - The ICMP code for IPv4 addresses. For example, specify `8` as the `fromPort` (ICMP type), and `-1` as the `toPort` (ICMP code), to enable ICMP Ping.
	// - ICMPv6 - The ICMP code for IPv6 addresses. For example, specify `128` as the `fromPort` (ICMPv6 type), and `0` as `toPort` (ICMPv6 code).
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

// `State` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the status code and the state (for example, `running` ) of an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stateProperty := &stateProperty{
//   	code: jsii.Number(123),
//   	name: jsii.String("name"),
//   }
//
type CfnInstance_StateProperty struct {
	// The status code of the instance.
	Code *float64 `field:"optional" json:"code" yaml:"code"`
	// The state of the instance (for example, `running` or `pending` ).
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Properties for defining a `CfnInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceProps := &cfnInstanceProps{
//   	blueprintId: jsii.String("blueprintId"),
//   	bundleId: jsii.String("bundleId"),
//   	instanceName: jsii.String("instanceName"),
//
//   	// the properties below are optional
//   	addOns: []interface{}{
//   		&addOnProperty{
//   			addOnType: jsii.String("addOnType"),
//
//   			// the properties below are optional
//   			autoSnapshotAddOnRequest: &autoSnapshotAddOnProperty{
//   				snapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   			},
//   			status: jsii.String("status"),
//   		},
//   	},
//   	availabilityZone: jsii.String("availabilityZone"),
//   	hardware: &hardwareProperty{
//   		cpuCount: jsii.Number(123),
//   		disks: []interface{}{
//   			&diskProperty{
//   				diskName: jsii.String("diskName"),
//   				path: jsii.String("path"),
//
//   				// the properties below are optional
//   				attachedTo: jsii.String("attachedTo"),
//   				attachmentState: jsii.String("attachmentState"),
//   				iops: jsii.Number(123),
//   				isSystemDisk: jsii.Boolean(false),
//   				sizeInGb: jsii.String("sizeInGb"),
//   			},
//   		},
//   		ramSizeInGb: jsii.Number(123),
//   	},
//   	keyPairName: jsii.String("keyPairName"),
//   	networking: &networkingProperty{
//   		ports: []interface{}{
//   			&portProperty{
//   				accessDirection: jsii.String("accessDirection"),
//   				accessFrom: jsii.String("accessFrom"),
//   				accessType: jsii.String("accessType"),
//   				cidrListAliases: []*string{
//   					jsii.String("cidrListAliases"),
//   				},
//   				cidrs: []*string{
//   					jsii.String("cidrs"),
//   				},
//   				commonName: jsii.String("commonName"),
//   				fromPort: jsii.Number(123),
//   				ipv6Cidrs: []*string{
//   					jsii.String("ipv6Cidrs"),
//   				},
//   				protocol: jsii.String("protocol"),
//   				toPort: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		monthlyTransfer: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userData: jsii.String("userData"),
//   }
//
type CfnInstanceProps struct {
	// The blueprint ID for the instance (for example, `os_amlinux_2016_03` ).
	BlueprintId *string `field:"required" json:"blueprintId" yaml:"blueprintId"`
	// The bundle ID for the instance (for example, `micro_1_0` ).
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// The name of the instance.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// An array of add-ons for the instance.
	//
	// > If the instance has an add-on enabled when performing a delete instance request, the add-on is automatically disabled before the instance is deleted.
	AddOns interface{} `field:"optional" json:"addOns" yaml:"addOns"`
	// The Availability Zone for the instance.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The hardware properties for the instance, such as the vCPU count, attached disks, and amount of RAM.
	//
	// > The instance restarts when performing an attach disk or detach disk request. This resets the public IP address of your instance if a static IP isn't attached to it.
	Hardware interface{} `field:"optional" json:"hardware" yaml:"hardware"`
	// The name of the key pair to use for the instance.
	//
	// If no key pair name is specified, the Regional Lightsail default key pair is used.
	KeyPairName *string `field:"optional" json:"keyPairName" yaml:"keyPairName"`
	// The public ports and the monthly amount of data transfer allocated for the instance.
	Networking interface{} `field:"optional" json:"networking" yaml:"networking"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The optional launch script for the instance.
	//
	// Specify a launch script to configure an instance with additional user data. For example, you might want to specify `apt-get -y update` as a launch script.
	//
	// > Depending on the blueprint of your instance, the command to get software on your instance varies. Amazon Linux and CentOS use `yum` , Debian and Ubuntu use `apt-get` , and FreeBSD uses `pkg` .
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

// A CloudFormation `AWS::Lightsail::LoadBalancer`.
//
// The `AWS::Lightsail::LoadBalancer` resource specifies a load balancer that can be used with Lightsail instances.
//
// > You cannot attach a TLS certificate to a load balancer using the `AWS::Lightsail::LoadBalancer` resource type. Instead, use the `AWS::Lightsail::LoadBalancerTlsCertificate` resource type to create a certificate and attach it to a load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoadBalancer := awscdk.Aws_lightsail.NewCfnLoadBalancer(this, jsii.String("MyCfnLoadBalancer"), &cfnLoadBalancerProps{
//   	instancePort: jsii.Number(123),
//   	loadBalancerName: jsii.String("loadBalancerName"),
//
//   	// the properties below are optional
//   	attachedInstances: []*string{
//   		jsii.String("attachedInstances"),
//   	},
//   	healthCheckPath: jsii.String("healthCheckPath"),
//   	ipAddressType: jsii.String("ipAddressType"),
//   	sessionStickinessEnabled: jsii.Boolean(false),
//   	sessionStickinessLbCookieDurationSeconds: jsii.String("sessionStickinessLbCookieDurationSeconds"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tlsPolicyName: jsii.String("tlsPolicyName"),
//   })
//
type CfnLoadBalancer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Lightsail instances to attach to the load balancer.
	AttachedInstances() *[]*string
	SetAttachedInstances(val *[]*string)
	// The Amazon Resource Name (ARN) of the load balancer.
	AttrLoadBalancerArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The path on the attached instance where the health check will be performed.
	//
	// If no path is specified, the load balancer tries to make a request to the default (root) page ( `/index.html` ).
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	// The port that the load balancer uses to direct traffic to your Lightsail instances.
	//
	// For HTTP traffic, specify port `80` . For HTTPS traffic, specify port `443` .
	InstancePort() *float64
	SetInstancePort(val *float64)
	// The IP address type of the load balancer.
	//
	// The possible values are `ipv4` for IPv4 only, and `dualstack` for both IPv4 and IPv6.
	IpAddressType() *string
	SetIpAddressType(val *string)
	// The name of the load balancer.
	LoadBalancerName() *string
	SetLoadBalancerName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A Boolean value indicating whether session stickiness is enabled.
	//
	// Enable session stickiness (also known as *session affinity* ) to bind a user's session to a specific instance. This ensures that all requests from the user during the session are sent to the same instance.
	SessionStickinessEnabled() interface{}
	SetSessionStickinessEnabled(val interface{})
	// The time period, in seconds, after which the load balancer session stickiness cookie should be considered stale.
	//
	// If you do not specify this parameter, the default value is 0, which indicates that the sticky session should last for the duration of the browser session.
	SessionStickinessLbCookieDurationSeconds() *string
	SetSessionStickinessLbCookieDurationSeconds(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags() awscdk.TagManager
	// The name of the TLS security policy for the load balancer.
	TlsPolicyName() *string
	SetTlsPolicyName(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnLoadBalancer
type jsiiProxy_CfnLoadBalancer struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLoadBalancer) AttachedInstances() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attachedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) AttrLoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) InstancePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) SessionStickinessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionStickinessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) SessionStickinessLbCookieDurationSeconds() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionStickinessLbCookieDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) TlsPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::LoadBalancer`.
func NewCfnLoadBalancer(scope constructs.Construct, id *string, props *CfnLoadBalancerProps) CfnLoadBalancer {
	_init_.Initialize()

	j := jsiiProxy_CfnLoadBalancer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::LoadBalancer`.
func NewCfnLoadBalancer_Override(c CfnLoadBalancer, scope constructs.Construct, id *string, props *CfnLoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetAttachedInstances(val *[]*string) {
	_jsii_.Set(
		j,
		"attachedInstances",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetHealthCheckPath(val *string) {
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetInstancePort(val *float64) {
	_jsii_.Set(
		j,
		"instancePort",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetIpAddressType(val *string) {
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetLoadBalancerName(val *string) {
	_jsii_.Set(
		j,
		"loadBalancerName",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetSessionStickinessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"sessionStickinessEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetSessionStickinessLbCookieDurationSeconds(val *string) {
	_jsii_.Set(
		j,
		"sessionStickinessLbCookieDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetTlsPolicyName(val *string) {
	_jsii_.Set(
		j,
		"tlsPolicyName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLoadBalancer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLoadBalancer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancer",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoadBalancer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoadBalancer) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancer) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancer) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancer) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancer) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLoadBalancer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoadBalancerProps := &cfnLoadBalancerProps{
//   	instancePort: jsii.Number(123),
//   	loadBalancerName: jsii.String("loadBalancerName"),
//
//   	// the properties below are optional
//   	attachedInstances: []*string{
//   		jsii.String("attachedInstances"),
//   	},
//   	healthCheckPath: jsii.String("healthCheckPath"),
//   	ipAddressType: jsii.String("ipAddressType"),
//   	sessionStickinessEnabled: jsii.Boolean(false),
//   	sessionStickinessLbCookieDurationSeconds: jsii.String("sessionStickinessLbCookieDurationSeconds"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tlsPolicyName: jsii.String("tlsPolicyName"),
//   }
//
type CfnLoadBalancerProps struct {
	// The port that the load balancer uses to direct traffic to your Lightsail instances.
	//
	// For HTTP traffic, specify port `80` . For HTTPS traffic, specify port `443` .
	InstancePort *float64 `field:"required" json:"instancePort" yaml:"instancePort"`
	// The name of the load balancer.
	LoadBalancerName *string `field:"required" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The Lightsail instances to attach to the load balancer.
	AttachedInstances *[]*string `field:"optional" json:"attachedInstances" yaml:"attachedInstances"`
	// The path on the attached instance where the health check will be performed.
	//
	// If no path is specified, the load balancer tries to make a request to the default (root) page ( `/index.html` ).
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The IP address type of the load balancer.
	//
	// The possible values are `ipv4` for IPv4 only, and `dualstack` for both IPv4 and IPv6.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// A Boolean value indicating whether session stickiness is enabled.
	//
	// Enable session stickiness (also known as *session affinity* ) to bind a user's session to a specific instance. This ensures that all requests from the user during the session are sent to the same instance.
	SessionStickinessEnabled interface{} `field:"optional" json:"sessionStickinessEnabled" yaml:"sessionStickinessEnabled"`
	// The time period, in seconds, after which the load balancer session stickiness cookie should be considered stale.
	//
	// If you do not specify this parameter, the default value is 0, which indicates that the sticky session should last for the duration of the browser session.
	SessionStickinessLbCookieDurationSeconds *string `field:"optional" json:"sessionStickinessLbCookieDurationSeconds" yaml:"sessionStickinessLbCookieDurationSeconds"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the TLS security policy for the load balancer.
	TlsPolicyName *string `field:"optional" json:"tlsPolicyName" yaml:"tlsPolicyName"`
}

// A CloudFormation `AWS::Lightsail::LoadBalancerTlsCertificate`.
//
// The `AWS::Lightsail::LoadBalancerTlsCertificate` resource specifies a TLS certificate that can be used with a Lightsail load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoadBalancerTlsCertificate := awscdk.Aws_lightsail.NewCfnLoadBalancerTlsCertificate(this, jsii.String("MyCfnLoadBalancerTlsCertificate"), &cfnLoadBalancerTlsCertificateProps{
//   	certificateDomainName: jsii.String("certificateDomainName"),
//   	certificateName: jsii.String("certificateName"),
//   	loadBalancerName: jsii.String("loadBalancerName"),
//
//   	// the properties below are optional
//   	certificateAlternativeNames: []*string{
//   		jsii.String("certificateAlternativeNames"),
//   	},
//   	httpsRedirectionEnabled: jsii.Boolean(false),
//   	isAttached: jsii.Boolean(false),
//   })
//
type CfnLoadBalancerTlsCertificate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the SSL/TLS certificate.
	AttrLoadBalancerTlsCertificateArn() *string
	// The validation status of the SSL/TLS certificate.
	//
	// Valid Values: `PENDING_VALIDATION` | `ISSUED` | `INACTIVE` | `EXPIRED` | `VALIDATION_TIMED_OUT` | `REVOKED` | `FAILED` | `UNKNOWN`.
	AttrStatus() *string
	// An array of alternative domain names and subdomain names for your SSL/TLS certificate.
	//
	// In addition to the primary domain name, you can have up to nine alternative domain names. Wildcards (such as `*.example.com` ) are not supported.
	CertificateAlternativeNames() *[]*string
	SetCertificateAlternativeNames(val *[]*string)
	// The domain name for the SSL/TLS certificate.
	//
	// For example, `example.com` or `www.example.com` .
	CertificateDomainName() *string
	SetCertificateDomainName(val *string)
	// The name of the SSL/TLS certificate.
	CertificateName() *string
	SetCertificateName(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A Boolean value indicating whether HTTPS redirection is enabled for the load balancer that the TLS certificate is attached to.
	HttpsRedirectionEnabled() interface{}
	SetHttpsRedirectionEnabled(val interface{})
	// A Boolean value indicating whether the SSL/TLS certificate is attached to a Lightsail load balancer.
	IsAttached() interface{}
	SetIsAttached(val interface{})
	// The name of the load balancer that the SSL/TLS certificate is attached to.
	LoadBalancerName() *string
	SetLoadBalancerName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnLoadBalancerTlsCertificate
type jsiiProxy_CfnLoadBalancerTlsCertificate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) AttrLoadBalancerTlsCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoadBalancerTlsCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) CertificateAlternativeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) CertificateDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) HttpsRedirectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsRedirectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) IsAttached() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::LoadBalancerTlsCertificate`.
func NewCfnLoadBalancerTlsCertificate(scope constructs.Construct, id *string, props *CfnLoadBalancerTlsCertificateProps) CfnLoadBalancerTlsCertificate {
	_init_.Initialize()

	j := jsiiProxy_CfnLoadBalancerTlsCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancerTlsCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::LoadBalancerTlsCertificate`.
func NewCfnLoadBalancerTlsCertificate_Override(c CfnLoadBalancerTlsCertificate, scope constructs.Construct, id *string, props *CfnLoadBalancerTlsCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancerTlsCertificate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) SetCertificateAlternativeNames(val *[]*string) {
	_jsii_.Set(
		j,
		"certificateAlternativeNames",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) SetCertificateDomainName(val *string) {
	_jsii_.Set(
		j,
		"certificateDomainName",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) SetCertificateName(val *string) {
	_jsii_.Set(
		j,
		"certificateName",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) SetHttpsRedirectionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"httpsRedirectionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) SetIsAttached(val interface{}) {
	_jsii_.Set(
		j,
		"isAttached",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancerTlsCertificate) SetLoadBalancerName(val *string) {
	_jsii_.Set(
		j,
		"loadBalancerName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLoadBalancerTlsCertificate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancerTlsCertificate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLoadBalancerTlsCertificate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancerTlsCertificate",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnLoadBalancerTlsCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancerTlsCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoadBalancerTlsCertificate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnLoadBalancerTlsCertificate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancerTlsCertificate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLoadBalancerTlsCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoadBalancerTlsCertificateProps := &cfnLoadBalancerTlsCertificateProps{
//   	certificateDomainName: jsii.String("certificateDomainName"),
//   	certificateName: jsii.String("certificateName"),
//   	loadBalancerName: jsii.String("loadBalancerName"),
//
//   	// the properties below are optional
//   	certificateAlternativeNames: []*string{
//   		jsii.String("certificateAlternativeNames"),
//   	},
//   	httpsRedirectionEnabled: jsii.Boolean(false),
//   	isAttached: jsii.Boolean(false),
//   }
//
type CfnLoadBalancerTlsCertificateProps struct {
	// The domain name for the SSL/TLS certificate.
	//
	// For example, `example.com` or `www.example.com` .
	CertificateDomainName *string `field:"required" json:"certificateDomainName" yaml:"certificateDomainName"`
	// The name of the SSL/TLS certificate.
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// The name of the load balancer that the SSL/TLS certificate is attached to.
	LoadBalancerName *string `field:"required" json:"loadBalancerName" yaml:"loadBalancerName"`
	// An array of alternative domain names and subdomain names for your SSL/TLS certificate.
	//
	// In addition to the primary domain name, you can have up to nine alternative domain names. Wildcards (such as `*.example.com` ) are not supported.
	CertificateAlternativeNames *[]*string `field:"optional" json:"certificateAlternativeNames" yaml:"certificateAlternativeNames"`
	// A Boolean value indicating whether HTTPS redirection is enabled for the load balancer that the TLS certificate is attached to.
	HttpsRedirectionEnabled interface{} `field:"optional" json:"httpsRedirectionEnabled" yaml:"httpsRedirectionEnabled"`
	// A Boolean value indicating whether the SSL/TLS certificate is attached to a Lightsail load balancer.
	IsAttached interface{} `field:"optional" json:"isAttached" yaml:"isAttached"`
}

// A CloudFormation `AWS::Lightsail::StaticIp`.
//
// The `AWS::Lightsail::StaticIp` resource specifies a static IP that can be attached to an Amazon Lightsail instance that is in the same AWS Region and Availability Zone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStaticIp := awscdk.Aws_lightsail.NewCfnStaticIp(this, jsii.String("MyCfnStaticIp"), &cfnStaticIpProps{
//   	staticIpName: jsii.String("staticIpName"),
//
//   	// the properties below are optional
//   	attachedTo: jsii.String("attachedTo"),
//   })
//
type CfnStaticIp interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The instance that the static IP is attached to.
	AttachedTo() *string
	SetAttachedTo(val *string)
	// The IP address of the static IP.
	AttrIpAddress() *string
	// A Boolean value indicating whether the static IP is attached to an instance.
	AttrIsAttached() awscdk.IResolvable
	// The Amazon Resource Name (ARN) of the static IP (for example, `arn:aws:lightsail:us-east-2:123456789101:StaticIp/244ad76f-8aad-4741-809f-12345EXAMPLE` ).
	AttrStaticIpArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The name of the static IP.
	StaticIpName() *string
	SetStaticIpName(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnStaticIp
type jsiiProxy_CfnStaticIp struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStaticIp) AttachedTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) AttrIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) AttrIsAttached() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) AttrStaticIpArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStaticIpArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) StaticIpName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticIpName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::StaticIp`.
func NewCfnStaticIp(scope constructs.Construct, id *string, props *CfnStaticIpProps) CfnStaticIp {
	_init_.Initialize()

	j := jsiiProxy_CfnStaticIp{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::StaticIp`.
func NewCfnStaticIp_Override(c CfnStaticIp, scope constructs.Construct, id *string, props *CfnStaticIpProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStaticIp) SetAttachedTo(val *string) {
	_jsii_.Set(
		j,
		"attachedTo",
		val,
	)
}

func (j *jsiiProxy_CfnStaticIp) SetStaticIpName(val *string) {
	_jsii_.Set(
		j,
		"staticIpName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStaticIp_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStaticIp_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnStaticIp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStaticIp_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStaticIp) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStaticIp) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStaticIp) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStaticIp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStaticIp) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStaticIp) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStaticIp) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStaticIp) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStaticIp) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStaticIp) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStaticIp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStaticIp) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStaticIp) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStaticIp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStaticIp) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnStaticIp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStaticIpProps := &cfnStaticIpProps{
//   	staticIpName: jsii.String("staticIpName"),
//
//   	// the properties below are optional
//   	attachedTo: jsii.String("attachedTo"),
//   }
//
type CfnStaticIpProps struct {
	// The name of the static IP.
	StaticIpName *string `field:"required" json:"staticIpName" yaml:"staticIpName"`
	// The instance that the static IP is attached to.
	AttachedTo *string `field:"optional" json:"attachedTo" yaml:"attachedTo"`
}

