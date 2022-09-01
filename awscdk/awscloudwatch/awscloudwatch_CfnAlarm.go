package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::CloudWatch::Alarm`.
//
// The `AWS::CloudWatch::Alarm` type specifies an alarm and associates it with the specified metric or metric math expression.
//
// When this operation creates an alarm, the alarm state is immediately set to `INSUFFICIENT_DATA` . The alarm is then evaluated and its state is set appropriately. Any actions associated with the new state are then executed.
//
// When you update an existing alarm, its state is left unchanged, but the update completely overwrites the previous configuration of the alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarm := awscdk.Aws_cloudwatch.NewCfnAlarm(this, jsii.String("MyCfnAlarm"), &cfnAlarmProps{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	evaluationPeriods: jsii.Number(123),
//
//   	// the properties below are optional
//   	actionsEnabled: jsii.Boolean(false),
//   	alarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	alarmDescription: jsii.String("alarmDescription"),
//   	alarmName: jsii.String("alarmName"),
//   	datapointsToAlarm: jsii.Number(123),
//   	dimensions: []interface{}{
//   		&dimensionProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	evaluateLowSampleCountPercentile: jsii.String("evaluateLowSampleCountPercentile"),
//   	extendedStatistic: jsii.String("extendedStatistic"),
//   	insufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
//   	},
//   	metricName: jsii.String("metricName"),
//   	metrics: []interface{}{
//   		&metricDataQueryProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			accountId: jsii.String("accountId"),
//   			expression: jsii.String("expression"),
//   			label: jsii.String("label"),
//   			metricStat: &metricStatProperty{
//   				metric: &metricProperty{
//   					dimensions: []interface{}{
//   						&dimensionProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					metricName: jsii.String("metricName"),
//   					namespace: jsii.String("namespace"),
//   				},
//   				period: jsii.Number(123),
//   				stat: jsii.String("stat"),
//
//   				// the properties below are optional
//   				unit: jsii.String("unit"),
//   			},
//   			period: jsii.Number(123),
//   			returnData: jsii.Boolean(false),
//   		},
//   	},
//   	namespace: jsii.String("namespace"),
//   	okActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   	period: jsii.Number(123),
//   	statistic: jsii.String("statistic"),
//   	threshold: jsii.Number(123),
//   	thresholdMetricId: jsii.String("thresholdMetricId"),
//   	treatMissingData: jsii.String("treatMissingData"),
//   	unit: jsii.String("unit"),
//   })
//
type CfnAlarm interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Indicates whether actions should be executed during any changes to the alarm state.
	//
	// The default is TRUE.
	ActionsEnabled() interface{}
	SetActionsEnabled(val interface{})
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state.
	//
	// Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutMetricAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html) in the *Amazon CloudWatch API Reference* .
	AlarmActions() *[]*string
	SetAlarmActions(val *[]*string)
	// The description of the alarm.
	AlarmDescription() *string
	SetAlarmDescription(val *string)
	// The name of the alarm.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the alarm name.
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	AlarmName() *string
	SetAlarmName(val *string)
	// The ARN of the CloudWatch alarm, such as `arn:aws:cloudwatch:us-west-2:123456789012:alarm:myCloudWatchAlarm-CPUAlarm-UXMMZK36R55Z` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The arithmetic operation to use when comparing the specified statistic and threshold.
	//
	// The specified statistic value is used as the first operand.
	//
	// You can specify the following values: `GreaterThanThreshold` , `GreaterThanOrEqualToThreshold` , `LessThanThreshold` , or `LessThanOrEqualToThreshold` .
	ComparisonOperator() *string
	SetComparisonOperator(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The number of datapoints that must be breaching to trigger the alarm.
	//
	// This is used only if you are setting an "M out of N" alarm. In that case, this value is the M, and the value that you set for `EvaluationPeriods` is the N value. For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *Amazon CloudWatch User Guide* .
	//
	// If you omit this parameter, CloudWatch uses the same value here that you set for `EvaluationPeriods` , and the alarm goes to alarm state if that many consecutive periods are breaching.
	DatapointsToAlarm() *float64
	SetDatapointsToAlarm(val *float64)
	// The dimensions for the metric associated with the alarm.
	//
	// For an alarm based on a math expression, you can't specify `Dimensions` . Instead, you use `Metrics` .
	Dimensions() interface{}
	SetDimensions(val interface{})
	// Used only for alarms based on percentiles.
	//
	// If `ignore` , the alarm state does not change during periods with too few data points to be statistically significant. If `evaluate` or this parameter is not used, the alarm is always evaluated and possibly changes state no matter how many data points are available.
	EvaluateLowSampleCountPercentile() *string
	SetEvaluateLowSampleCountPercentile(val *string)
	// The number of periods over which data is compared to the specified threshold.
	//
	// If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an "M out of N" alarm, this value is the N, and `DatapointsToAlarm` is the M.
	//
	// For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *Amazon CloudWatch User Guide* .
	EvaluationPeriods() *float64
	SetEvaluationPeriods(val *float64)
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	//
	// For an alarm based on a metric, you must specify either `Statistic` or `ExtendedStatistic` but not both.
	//
	// For an alarm based on a math expression, you can't specify `ExtendedStatistic` . Instead, you use `Metrics` .
	ExtendedStatistic() *string
	SetExtendedStatistic(val *string)
	// The actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN).
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
	LogicalId() *string
	// The name of the metric associated with the alarm.
	//
	// This is required for an alarm based on a metric. For an alarm based on a math expression, you use `Metrics` instead and you can't specify `MetricName` .
	MetricName() *string
	SetMetricName(val *string)
	// An array that enables you to create an alarm based on the result of a metric math expression.
	//
	// Each item in the array either retrieves a metric or performs a math expression.
	//
	// If you specify the `Metrics` parameter, you cannot specify `MetricName` , `Dimensions` , `Period` , `Namespace` , `Statistic` , `ExtendedStatistic` , or `Unit` .
	Metrics() interface{}
	SetMetrics(val interface{})
	// The namespace of the metric associated with the alarm.
	//
	// This is required for an alarm based on a metric. For an alarm based on a math expression, you can't specify `Namespace` and you use `Metrics` instead.
	//
	// For a list of namespaces for metrics from AWS services, see [AWS Services That Publish CloudWatch Metrics.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)
	Namespace() *string
	SetNamespace(val *string)
	// The tree node.
	Node() constructs.Node
	// The actions to execute when this alarm transitions to the `OK` state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN).
	OkActions() *[]*string
	SetOkActions(val *[]*string)
	// The period, in seconds, over which the statistic is applied.
	//
	// This is required for an alarm based on a metric. Valid values are 10, 30, 60, and any multiple of 60.
	//
	// For an alarm based on a math expression, you can't specify `Period` , and instead you use the `Metrics` parameter.
	//
	// *Minimum:* 10.
	Period() *float64
	SetPeriod(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The statistic for the metric associated with the alarm, other than percentile. For percentile statistics, use `ExtendedStatistic` .
	//
	// For an alarm based on a metric, you must specify either `Statistic` or `ExtendedStatistic` but not both.
	//
	// For an alarm based on a math expression, you can't specify `Statistic` . Instead, you use `Metrics` .
	Statistic() *string
	SetStatistic(val *string)
	// The value to compare with the specified statistic.
	Threshold() *float64
	SetThreshold(val *float64)
	// In an alarm based on an anomaly detection model, this is the ID of the `ANOMALY_DETECTION_BAND` function used as the threshold for the alarm.
	ThresholdMetricId() *string
	SetThresholdMetricId(val *string)
	// Sets how this alarm is to handle missing data points.
	//
	// Valid values are `breaching` , `notBreaching` , `ignore` , and `missing` . For more information, see [Configuring How CloudWatch Alarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the *Amazon CloudWatch User Guide* .
	//
	// If you omit this parameter, the default behavior of `missing` is used.
	TreatMissingData() *string
	SetTreatMissingData(val *string)
	// The unit of the metric associated with the alarm.
	//
	// Specify this only if you are creating an alarm based on a single metric. Do not specify this if you are specifying a `Metrics` array.
	//
	// You can specify the following values: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.
	Unit() *string
	SetUnit(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
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

func (j *jsiiProxy_CfnAlarm) ActionsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) AlarmActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) AlarmDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescription",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnAlarm) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
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

func (j *jsiiProxy_CfnAlarm) Dimensions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) EvaluateLowSampleCountPercentile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluateLowSampleCountPercentile",
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

func (j *jsiiProxy_CfnAlarm) ExtendedStatistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedStatistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) InsufficientDataActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insufficientDataActions",
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

func (j *jsiiProxy_CfnAlarm) Metrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
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

func (j *jsiiProxy_CfnAlarm) OkActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"okActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarm) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
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

func (j *jsiiProxy_CfnAlarm) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
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

func (j *jsiiProxy_CfnAlarm) ThresholdMetricId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdMetricId",
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

func (j *jsiiProxy_CfnAlarm) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
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

func (j *jsiiProxy_CfnAlarm) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudWatch::Alarm`.
func NewCfnAlarm(scope constructs.Construct, id *string, props *CfnAlarmProps) CfnAlarm {
	_init_.Initialize()

	if err := validateNewCfnAlarmParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAlarm{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.CfnAlarm",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudWatch::Alarm`.
func NewCfnAlarm_Override(c CfnAlarm, scope constructs.Construct, id *string, props *CfnAlarmProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.CfnAlarm",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAlarm)SetActionsEnabled(val interface{}) {
	if err := j.validateSetActionsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionsEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetAlarmActions(val *[]*string) {
	_jsii_.Set(
		j,
		"alarmActions",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetAlarmDescription(val *string) {
	_jsii_.Set(
		j,
		"alarmDescription",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetAlarmName(val *string) {
	_jsii_.Set(
		j,
		"alarmName",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetComparisonOperator(val *string) {
	if err := j.validateSetComparisonOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonOperator",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetDatapointsToAlarm(val *float64) {
	_jsii_.Set(
		j,
		"datapointsToAlarm",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetDimensions(val interface{}) {
	if err := j.validateSetDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetEvaluateLowSampleCountPercentile(val *string) {
	_jsii_.Set(
		j,
		"evaluateLowSampleCountPercentile",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetEvaluationPeriods(val *float64) {
	if err := j.validateSetEvaluationPeriodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationPeriods",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetExtendedStatistic(val *string) {
	_jsii_.Set(
		j,
		"extendedStatistic",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetInsufficientDataActions(val *[]*string) {
	_jsii_.Set(
		j,
		"insufficientDataActions",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetMetricName(val *string) {
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetMetrics(val interface{}) {
	if err := j.validateSetMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metrics",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetOkActions(val *[]*string) {
	_jsii_.Set(
		j,
		"okActions",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetPeriod(val *float64) {
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetStatistic(val *string) {
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetThreshold(val *float64) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetThresholdMetricId(val *string) {
	_jsii_.Set(
		j,
		"thresholdMetricId",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetTreatMissingData(val *string) {
	_jsii_.Set(
		j,
		"treatMissingData",
		val,
	)
}

func (j *jsiiProxy_CfnAlarm)SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
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

	if err := validateCfnAlarm_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.CfnAlarm",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAlarm_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnAlarm_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.CfnAlarm",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnAlarm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAlarm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.CfnAlarm",
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
		"aws-cdk-lib.aws_cloudwatch.CfnAlarm",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAlarm) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAlarm) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAlarm) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAlarm) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAlarm) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAlarm) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAlarm) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAlarm) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnAlarm) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnAlarm) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAlarm) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAlarm) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

