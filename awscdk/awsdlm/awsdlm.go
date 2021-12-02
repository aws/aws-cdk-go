package awsdlm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdlm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::DLM::LifecyclePolicy`.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	PolicyDetails() interface{}
	SetPolicyDetails(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	State() *string
	SetState(val *string)
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnLifecyclePolicy
type jsiiProxy_CfnLifecyclePolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLifecyclePolicy) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) PolicyDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DLM::LifecyclePolicy`.
func NewCfnLifecyclePolicy(scope constructs.Construct, id *string, props *CfnLifecyclePolicyProps) CfnLifecyclePolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnLifecyclePolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DLM::LifecyclePolicy`.
func NewCfnLifecyclePolicy_Override(c CfnLifecyclePolicy, scope constructs.Construct, id *string, props *CfnLifecyclePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLifecyclePolicy) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnLifecyclePolicy) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnLifecyclePolicy) SetPolicyDetails(val interface{}) {
	_jsii_.Set(
		j,
		"policyDetails",
		val,
	)
}

func (j *jsiiProxy_CfnLifecyclePolicy) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLifecyclePolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLifecyclePolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
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
func CfnLifecyclePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLifecyclePolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_dlm.CfnLifecyclePolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLifecyclePolicy) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLifecyclePolicy) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLifecyclePolicy) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLifecyclePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLifecyclePolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLifecyclePolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLifecyclePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnLifecyclePolicy) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnLifecyclePolicy) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnLifecyclePolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLifecyclePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLifecyclePolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnLifecyclePolicy) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnLifecyclePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLifecyclePolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_ActionProperty struct {
	// `CfnLifecyclePolicy.ActionProperty.CrossRegionCopy`.
	CrossRegionCopy interface{} `json:"crossRegionCopy"`
	// `CfnLifecyclePolicy.ActionProperty.Name`.
	Name *string `json:"name"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_CreateRuleProperty struct {
	// `CfnLifecyclePolicy.CreateRuleProperty.CronExpression`.
	CronExpression *string `json:"cronExpression"`
	// `CfnLifecyclePolicy.CreateRuleProperty.Interval`.
	Interval *float64 `json:"interval"`
	// `CfnLifecyclePolicy.CreateRuleProperty.IntervalUnit`.
	IntervalUnit *string `json:"intervalUnit"`
	// `CfnLifecyclePolicy.CreateRuleProperty.Location`.
	Location *string `json:"location"`
	// `CfnLifecyclePolicy.CreateRuleProperty.Times`.
	Times *[]*string `json:"times"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_CrossRegionCopyActionProperty struct {
	// `CfnLifecyclePolicy.CrossRegionCopyActionProperty.EncryptionConfiguration`.
	EncryptionConfiguration interface{} `json:"encryptionConfiguration"`
	// `CfnLifecyclePolicy.CrossRegionCopyActionProperty.RetainRule`.
	RetainRule interface{} `json:"retainRule"`
	// `CfnLifecyclePolicy.CrossRegionCopyActionProperty.Target`.
	Target *string `json:"target"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_CrossRegionCopyDeprecateRuleProperty struct {
	// `CfnLifecyclePolicy.CrossRegionCopyDeprecateRuleProperty.Interval`.
	Interval *float64 `json:"interval"`
	// `CfnLifecyclePolicy.CrossRegionCopyDeprecateRuleProperty.IntervalUnit`.
	IntervalUnit *string `json:"intervalUnit"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_CrossRegionCopyRetainRuleProperty struct {
	// `CfnLifecyclePolicy.CrossRegionCopyRetainRuleProperty.Interval`.
	Interval *float64 `json:"interval"`
	// `CfnLifecyclePolicy.CrossRegionCopyRetainRuleProperty.IntervalUnit`.
	IntervalUnit *string `json:"intervalUnit"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_CrossRegionCopyRuleProperty struct {
	// `CfnLifecyclePolicy.CrossRegionCopyRuleProperty.CmkArn`.
	CmkArn *string `json:"cmkArn"`
	// `CfnLifecyclePolicy.CrossRegionCopyRuleProperty.CopyTags`.
	CopyTags interface{} `json:"copyTags"`
	// `CfnLifecyclePolicy.CrossRegionCopyRuleProperty.DeprecateRule`.
	DeprecateRule interface{} `json:"deprecateRule"`
	// `CfnLifecyclePolicy.CrossRegionCopyRuleProperty.Encrypted`.
	Encrypted interface{} `json:"encrypted"`
	// `CfnLifecyclePolicy.CrossRegionCopyRuleProperty.RetainRule`.
	RetainRule interface{} `json:"retainRule"`
	// `CfnLifecyclePolicy.CrossRegionCopyRuleProperty.Target`.
	Target *string `json:"target"`
	// `CfnLifecyclePolicy.CrossRegionCopyRuleProperty.TargetRegion`.
	TargetRegion *string `json:"targetRegion"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_DeprecateRuleProperty struct {
	// `CfnLifecyclePolicy.DeprecateRuleProperty.Count`.
	Count *float64 `json:"count"`
	// `CfnLifecyclePolicy.DeprecateRuleProperty.Interval`.
	Interval *float64 `json:"interval"`
	// `CfnLifecyclePolicy.DeprecateRuleProperty.IntervalUnit`.
	IntervalUnit *string `json:"intervalUnit"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_EncryptionConfigurationProperty struct {
	// `CfnLifecyclePolicy.EncryptionConfigurationProperty.CmkArn`.
	CmkArn *string `json:"cmkArn"`
	// `CfnLifecyclePolicy.EncryptionConfigurationProperty.Encrypted`.
	Encrypted interface{} `json:"encrypted"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_EventParametersProperty struct {
	// `CfnLifecyclePolicy.EventParametersProperty.DescriptionRegex`.
	DescriptionRegex *string `json:"descriptionRegex"`
	// `CfnLifecyclePolicy.EventParametersProperty.EventType`.
	EventType *string `json:"eventType"`
	// `CfnLifecyclePolicy.EventParametersProperty.SnapshotOwner`.
	SnapshotOwner *[]*string `json:"snapshotOwner"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_EventSourceProperty struct {
	// `CfnLifecyclePolicy.EventSourceProperty.Parameters`.
	Parameters interface{} `json:"parameters"`
	// `CfnLifecyclePolicy.EventSourceProperty.Type`.
	Type *string `json:"type"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_FastRestoreRuleProperty struct {
	// `CfnLifecyclePolicy.FastRestoreRuleProperty.AvailabilityZones`.
	AvailabilityZones *[]*string `json:"availabilityZones"`
	// `CfnLifecyclePolicy.FastRestoreRuleProperty.Count`.
	Count *float64 `json:"count"`
	// `CfnLifecyclePolicy.FastRestoreRuleProperty.Interval`.
	Interval *float64 `json:"interval"`
	// `CfnLifecyclePolicy.FastRestoreRuleProperty.IntervalUnit`.
	IntervalUnit *string `json:"intervalUnit"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_ParametersProperty struct {
	// `CfnLifecyclePolicy.ParametersProperty.ExcludeBootVolume`.
	ExcludeBootVolume interface{} `json:"excludeBootVolume"`
	// `CfnLifecyclePolicy.ParametersProperty.NoReboot`.
	NoReboot interface{} `json:"noReboot"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_PolicyDetailsProperty struct {
	// `CfnLifecyclePolicy.PolicyDetailsProperty.Actions`.
	Actions interface{} `json:"actions"`
	// `CfnLifecyclePolicy.PolicyDetailsProperty.EventSource`.
	EventSource interface{} `json:"eventSource"`
	// `CfnLifecyclePolicy.PolicyDetailsProperty.Parameters`.
	Parameters interface{} `json:"parameters"`
	// `CfnLifecyclePolicy.PolicyDetailsProperty.PolicyType`.
	PolicyType *string `json:"policyType"`
	// `CfnLifecyclePolicy.PolicyDetailsProperty.ResourceLocations`.
	ResourceLocations *[]*string `json:"resourceLocations"`
	// `CfnLifecyclePolicy.PolicyDetailsProperty.ResourceTypes`.
	ResourceTypes *[]*string `json:"resourceTypes"`
	// `CfnLifecyclePolicy.PolicyDetailsProperty.Schedules`.
	Schedules interface{} `json:"schedules"`
	// `CfnLifecyclePolicy.PolicyDetailsProperty.TargetTags`.
	TargetTags interface{} `json:"targetTags"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_RetainRuleProperty struct {
	// `CfnLifecyclePolicy.RetainRuleProperty.Count`.
	Count *float64 `json:"count"`
	// `CfnLifecyclePolicy.RetainRuleProperty.Interval`.
	Interval *float64 `json:"interval"`
	// `CfnLifecyclePolicy.RetainRuleProperty.IntervalUnit`.
	IntervalUnit *string `json:"intervalUnit"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_ScheduleProperty struct {
	// `CfnLifecyclePolicy.ScheduleProperty.CopyTags`.
	CopyTags interface{} `json:"copyTags"`
	// `CfnLifecyclePolicy.ScheduleProperty.CreateRule`.
	CreateRule interface{} `json:"createRule"`
	// `CfnLifecyclePolicy.ScheduleProperty.CrossRegionCopyRules`.
	CrossRegionCopyRules interface{} `json:"crossRegionCopyRules"`
	// `CfnLifecyclePolicy.ScheduleProperty.DeprecateRule`.
	DeprecateRule interface{} `json:"deprecateRule"`
	// `CfnLifecyclePolicy.ScheduleProperty.FastRestoreRule`.
	FastRestoreRule interface{} `json:"fastRestoreRule"`
	// `CfnLifecyclePolicy.ScheduleProperty.Name`.
	Name *string `json:"name"`
	// `CfnLifecyclePolicy.ScheduleProperty.RetainRule`.
	RetainRule interface{} `json:"retainRule"`
	// `CfnLifecyclePolicy.ScheduleProperty.ShareRules`.
	ShareRules interface{} `json:"shareRules"`
	// `CfnLifecyclePolicy.ScheduleProperty.TagsToAdd`.
	TagsToAdd interface{} `json:"tagsToAdd"`
	// `CfnLifecyclePolicy.ScheduleProperty.VariableTags`.
	VariableTags interface{} `json:"variableTags"`
}

// TODO: EXAMPLE
//
type CfnLifecyclePolicy_ShareRuleProperty struct {
	// `CfnLifecyclePolicy.ShareRuleProperty.TargetAccounts`.
	TargetAccounts *[]*string `json:"targetAccounts"`
	// `CfnLifecyclePolicy.ShareRuleProperty.UnshareInterval`.
	UnshareInterval *float64 `json:"unshareInterval"`
	// `CfnLifecyclePolicy.ShareRuleProperty.UnshareIntervalUnit`.
	UnshareIntervalUnit *string `json:"unshareIntervalUnit"`
}

// Properties for defining a `AWS::DLM::LifecyclePolicy`.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicyProps struct {
	// `AWS::DLM::LifecyclePolicy.Description`.
	Description *string `json:"description"`
	// `AWS::DLM::LifecyclePolicy.ExecutionRoleArn`.
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// `AWS::DLM::LifecyclePolicy.PolicyDetails`.
	PolicyDetails interface{} `json:"policyDetails"`
	// `AWS::DLM::LifecyclePolicy.State`.
	State *string `json:"state"`
	// `AWS::DLM::LifecyclePolicy.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

