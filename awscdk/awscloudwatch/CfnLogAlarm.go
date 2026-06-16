package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::CloudWatch::LogAlarm.
//
// A LogAlarm evaluates scheduled query results from CloudWatch Logs and triggers actions when thresholds are breached.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLogAlarm := awscdk.Aws_cloudwatch.NewCfnLogAlarm(this, jsii.String("MyCfnLogAlarm"), &CfnLogAlarmProps{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	QueryResultsToAlarm: jsii.Number(123),
//   	QueryResultsToEvaluate: jsii.Number(123),
//   	ScheduledQueryConfiguration: &ScheduledQueryConfigurationProperty{
//   		AggregationExpression: jsii.String("aggregationExpression"),
//   		LogGroupIdentifiers: []*string{
//   			jsii.String("logGroupIdentifiers"),
//   		},
//   		QueryLanguage: jsii.String("queryLanguage"),
//   		QueryString: jsii.String("queryString"),
//   		ScheduleConfiguration: &ScheduleConfigurationProperty{
//   			ScheduleExpression: jsii.String("scheduleExpression"),
//
//   			// the properties below are optional
//   			EndTimeOffset: jsii.Number(123),
//   			StartTimeOffset: jsii.Number(123),
//   		},
//   		ScheduledQueryRoleArn: jsii.String("scheduledQueryRoleArn"),
//   	},
//   	Threshold: jsii.Number(123),
//
//   	// the properties below are optional
//   	ActionLogLineCount: jsii.Number(123),
//   	ActionLogLineRoleArn: jsii.String("actionLogLineRoleArn"),
//   	ActionsEnabled: jsii.Boolean(false),
//   	AlarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	AlarmDescription: jsii.String("alarmDescription"),
//   	AlarmName: jsii.String("alarmName"),
//   	InsufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
//   	},
//   	OkActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TreatMissingData: jsii.String("treatMissingData"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html
//
type CfnLogAlarm interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawscloudwatch.ILogAlarmRef
	awscdk.ITaggableV2
	// The number of log lines to include in alarm notifications.
	ActionLogLineCount() *float64
	SetActionLogLineCount(val *float64)
	// The ARN of the IAM role that grants CloudWatch permissions to fetch log lines for alarm notifications.
	ActionLogLineRoleArn() *string
	SetActionLogLineRoleArn(val *string)
	// Indicates whether actions should be executed during any changes to the alarm state.
	ActionsEnabled() interface{}
	SetActionsEnabled(val interface{})
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state.
	AlarmActions() *[]*string
	SetAlarmActions(val *[]*string)
	// The description of the log alarm.
	AlarmDescription() *string
	SetAlarmDescription(val *string)
	// The name of the log alarm.
	AlarmName() *string
	SetAlarmName(val *string)
	// The ARN of the log alarm.
	AttrArn() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	CfnResourceType() *string
	// The arithmetic operation to use when comparing the specified threshold and the query results.
	ComparisonOperator() *string
	SetComparisonOperator(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	Env() *interfaces.ResourceEnvironment
	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
	InsufficientDataActions() *[]*string
	SetInsufficientDataActions(val *[]*string)
	// A reference to a LogAlarm resource.
	LogAlarmRef() *interfacesawscloudwatch.LogAlarmReference
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
	// The actions to execute when this alarm transitions to the OK state from any other state.
	OkActions() *[]*string
	SetOkActions(val *[]*string)
	// The number of query results that must be breaching to trigger the alarm.
	QueryResultsToAlarm() *float64
	SetQueryResultsToAlarm(val *float64)
	// The number of query results over which data is compared to the specified threshold.
	QueryResultsToEvaluate() *float64
	SetQueryResultsToEvaluate(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The scheduled query configuration for the log alarm.
	ScheduledQueryConfiguration() interface{}
	SetScheduledQueryConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of key-value pairs to associate with the log alarm.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// The value to compare against the results of the scheduled query evaluation.
	Threshold() *float64
	SetThreshold(val *float64)
	// Sets how this alarm is to handle missing data points.
	TreatMissingData() *string
	SetTreatMissingData(val *string)
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
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	// Sets the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// strength instead of the global default from the consuming stack's context.
	ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength)
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
	CfnPropertyName(cdkPropertyName *string) *string
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnLogAlarm
type jsiiProxy_CfnLogAlarm struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawscloudwatchILogAlarmRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnLogAlarm) ActionLogLineCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actionLogLineCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) ActionLogLineRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionLogLineRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) ActionsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) AlarmActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) AlarmDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) AlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) ComparisonOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) InsufficientDataActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insufficientDataActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) LogAlarmRef() *interfacesawscloudwatch.LogAlarmReference {
	var returns *interfacesawscloudwatch.LogAlarmReference
	_jsii_.Get(
		j,
		"logAlarmRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) OkActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"okActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) QueryResultsToAlarm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryResultsToAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) QueryResultsToEvaluate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryResultsToEvaluate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) ScheduledQueryConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledQueryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) TreatMissingData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatMissingData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarm) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudWatch::LogAlarm`.
func NewCfnLogAlarm(scope constructs.Construct, id *string, props *CfnLogAlarmProps) CfnLogAlarm {
	_init_.Initialize()

	if err := validateNewCfnLogAlarmParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLogAlarm{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.CfnLogAlarm",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudWatch::LogAlarm`.
func NewCfnLogAlarm_Override(c CfnLogAlarm, scope constructs.Construct, id *string, props *CfnLogAlarmProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.CfnLogAlarm",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetActionLogLineCount(val *float64) {
	_jsii_.Set(
		j,
		"actionLogLineCount",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetActionLogLineRoleArn(val *string) {
	_jsii_.Set(
		j,
		"actionLogLineRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetActionsEnabled(val interface{}) {
	if err := j.validateSetActionsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionsEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetAlarmActions(val *[]*string) {
	_jsii_.Set(
		j,
		"alarmActions",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetAlarmDescription(val *string) {
	_jsii_.Set(
		j,
		"alarmDescription",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetAlarmName(val *string) {
	_jsii_.Set(
		j,
		"alarmName",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetComparisonOperator(val *string) {
	if err := j.validateSetComparisonOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonOperator",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetInsufficientDataActions(val *[]*string) {
	_jsii_.Set(
		j,
		"insufficientDataActions",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetOkActions(val *[]*string) {
	_jsii_.Set(
		j,
		"okActions",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetQueryResultsToAlarm(val *float64) {
	if err := j.validateSetQueryResultsToAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryResultsToAlarm",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetQueryResultsToEvaluate(val *float64) {
	if err := j.validateSetQueryResultsToEvaluateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryResultsToEvaluate",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetScheduledQueryConfiguration(val interface{}) {
	if err := j.validateSetScheduledQueryConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledQueryConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_CfnLogAlarm)SetTreatMissingData(val *string) {
	_jsii_.Set(
		j,
		"treatMissingData",
		val,
	)
}

func CfnLogAlarm_ArnForLogAlarm(resource interfacesawscloudwatch.ILogAlarmRef) *string {
	_init_.Initialize()

	if err := validateCfnLogAlarm_ArnForLogAlarmParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.CfnLogAlarm",
		"arnForLogAlarm",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLogAlarm_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLogAlarm_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.CfnLogAlarm",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnLogAlarm.
func CfnLogAlarm_IsCfnLogAlarm(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLogAlarm_IsCfnLogAlarmParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.CfnLogAlarm",
		"isCfnLogAlarm",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnLogAlarm_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLogAlarm_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.CfnLogAlarm",
		"isCfnResource",
		[]interface{}{x},
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
func CfnLogAlarm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLogAlarm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.CfnLogAlarm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLogAlarm_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.CfnLogAlarm",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLogAlarm) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLogAlarm) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLogAlarm) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLogAlarm) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLogAlarm) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLogAlarm) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLogAlarm) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLogAlarm) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnLogAlarm) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLogAlarm) CfnPropertyName(cdkPropertyName *string) *string {
	if err := c.validateCfnPropertyNameParameters(cdkPropertyName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"cfnPropertyName",
		[]interface{}{cdkPropertyName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogAlarm) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogAlarm) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnLogAlarm) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLogAlarm) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogAlarm) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogAlarm) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLogAlarm) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLogAlarm) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnLogAlarm) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnLogAlarm) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogAlarm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLogAlarm) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnLogAlarm) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

