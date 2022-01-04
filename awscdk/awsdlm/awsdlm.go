package awsdlm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsdlm/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::DLM::LifecyclePolicy`.
//
// Specifies a lifecycle policy, which is used to automate operations on Amazon EBS resources.
//
// The properties are required when you add a lifecycle policy and optional when you update a lifecycle policy.
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnLifecyclePolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnLifecyclePolicy(scope awscdk.Construct, id *string, props *CfnLifecyclePolicyProps) CfnLifecyclePolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnLifecyclePolicy{}

	_jsii_.Create(
		"monocdk.aws_dlm.CfnLifecyclePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DLM::LifecyclePolicy`.
func NewCfnLifecyclePolicy_Override(c CfnLifecyclePolicy, scope awscdk.Construct, id *string, props *CfnLifecyclePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_dlm.CfnLifecyclePolicy",
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
// Experimental.
func CfnLifecyclePolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dlm.CfnLifecyclePolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnLifecyclePolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dlm.CfnLifecyclePolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnLifecyclePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dlm.CfnLifecyclePolicy",
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
		"monocdk.aws_dlm.CfnLifecyclePolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnLifecyclePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnLifecyclePolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnLifecyclePolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnLifecyclePolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnLifecyclePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnLifecyclePolicy) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnLifecyclePolicy) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnLifecyclePolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnLifecyclePolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies an action for an event-based policy.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_ActionProperty struct {
	// The rule for copying shared snapshots across Regions.
	CrossRegionCopy interface{} `json:"crossRegionCopy"`
	// A descriptive name for the action.
	Name *string `json:"name"`
}

// Specifies when to create snapshots of EBS volumes.
//
// You must specify either a Cron expression or an interval, interval unit, and start time. You cannot specify both.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_CreateRuleProperty struct {
	// The schedule, as a Cron expression.
	//
	// The schedule interval must be between 1 hour and 1 year. For more information, see [Cron expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions) in the *Amazon CloudWatch User Guide* .
	CronExpression *string `json:"cronExpression"`
	// The interval between snapshots.
	//
	// The supported values are 1, 2, 3, 4, 6, 8, 12, and 24.
	Interval *float64 `json:"interval"`
	// The interval unit.
	IntervalUnit *string `json:"intervalUnit"`
	// Specifies the destination for snapshots created by the policy.
	//
	// To create snapshots in the same Region as the source resource, specify `CLOUD` . To create snapshots on the same Outpost as the source resource, specify `OUTPOST_LOCAL` . If you omit this parameter, `CLOUD` is used by default.
	//
	// If the policy targets resources in an AWS Region , then you must create snapshots in the same Region as the source resource.
	//
	// If the policy targets resources on an Outpost, then you can create snapshots on the same Outpost as the source resource, or in the Region of that Outpost.
	Location *string `json:"location"`
	// The time, in UTC, to start the operation. The supported format is hh:mm.
	//
	// The operation occurs within a one-hour window following the specified time. If you do not specify a time, Amazon DLM selects a time within the next 24 hours.
	Times *[]*string `json:"times"`
}

// Specifies a rule for copying shared snapshots across Regions.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_CrossRegionCopyActionProperty struct {
	// The encryption settings for the copied snapshot.
	EncryptionConfiguration interface{} `json:"encryptionConfiguration"`
	// The target Region.
	Target *string `json:"target"`
	// Specifies the retention rule for cross-Region snapshot copies.
	RetainRule interface{} `json:"retainRule"`
}

// Specifies an AMI deprecation rule for cross-Region AMI copies created by a cross-Region copy rule.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_CrossRegionCopyDeprecateRuleProperty struct {
	// The period after which to deprecate the cross-Region AMI copies.
	//
	// The period must be less than or equal to the cross-Region AMI copy retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.
	Interval *float64 `json:"interval"`
	// The unit of time in which to measure the *Interval* .
	IntervalUnit *string `json:"intervalUnit"`
}

// Specifies the retention rule for cross-Region snapshot copies.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_CrossRegionCopyRetainRuleProperty struct {
	// The amount of time to retain each snapshot.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *float64 `json:"interval"`
	// The unit of time for time-based retention.
	IntervalUnit *string `json:"intervalUnit"`
}

// Specifies a rule for cross-Region snapshot copies.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_CrossRegionCopyRuleProperty struct {
	// To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled, enable encryption using this parameter.
	//
	// Copies of encrypted snapshots are encrypted, even if this parameter is false or if encryption by default is not enabled.
	Encrypted interface{} `json:"encrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption.
	//
	// If this parameter is not specified, the default KMS key for the account is used.
	CmkArn *string `json:"cmkArn"`
	// Indicates whether to copy all user-defined tags from the source snapshot to the cross-Region snapshot copy.
	CopyTags interface{} `json:"copyTags"`
	// The AMI deprecation rule for cross-Region AMI copies created by the rule.
	DeprecateRule interface{} `json:"deprecateRule"`
	// The retention rule that indicates how long snapshot copies are to be retained in the destination Region.
	RetainRule interface{} `json:"retainRule"`
	// The target Region or the Amazon Resource Name (ARN) of the target Outpost for the snapshot copies.
	//
	// Use this parameter instead of *TargetRegion* . Do not specify both.
	Target *string `json:"target"`
	// Avoid using this parameter when creating new policies.
	//
	// Instead, use *Target* to specify a target Region or a target Outpost for snapshot copies.
	//
	// For policies created before the *Target* parameter was introduced, this parameter indicates the target Region for snapshot copies.
	TargetRegion *string `json:"targetRegion"`
}

// Specifies an AMI deprecation rule for a schedule.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_DeprecateRuleProperty struct {
	// If the schedule has a count-based retention rule, this parameter specifies the number of oldest AMIs to deprecate.
	//
	// The count must be less than or equal to the schedule's retention count, and it can't be greater than 1000.
	Count *float64 `json:"count"`
	// If the schedule has an age-based retention rule, this parameter specifies the period after which to deprecate AMIs created by the schedule.
	//
	// The period must be less than or equal to the schedule's retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.
	Interval *float64 `json:"interval"`
	// The unit of time in which to measure the *Interval* .
	IntervalUnit *string `json:"intervalUnit"`
}

// Specifies the encryption settings for shared snapshots that are copied across Regions.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_EncryptionConfigurationProperty struct {
	// To encrypt a copy of an unencrypted snapshot when encryption by default is not enabled, enable encryption using this parameter.
	//
	// Copies of encrypted snapshots are encrypted, even if this parameter is false or when encryption by default is not enabled.
	Encrypted interface{} `json:"encrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption.
	//
	// If this parameter is not specified, the default KMS key for the account is used.
	CmkArn *string `json:"cmkArn"`
}

// Specifies an event that triggers an event-based policy.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_EventParametersProperty struct {
	// The type of event.
	//
	// Currently, only snapshot sharing events are supported.
	EventType *string `json:"eventType"`
	// The IDs of the AWS accounts that can trigger policy by sharing snapshots with your account.
	//
	// The policy only runs if one of the specified AWS accounts shares a snapshot with your account.
	SnapshotOwner *[]*string `json:"snapshotOwner"`
	// The snapshot description that can trigger the policy.
	//
	// The description pattern is specified using a regular expression. The policy runs only if a snapshot with a description that matches the specified pattern is shared with your account.
	//
	// For example, specifying `^.*Created for policy: policy-1234567890abcdef0.*$` configures the policy to run only if snapshots created by policy `policy-1234567890abcdef0` are shared with your account.
	DescriptionRegex *string `json:"descriptionRegex"`
}

// Specifies an event that triggers an event-based policy.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_EventSourceProperty struct {
	// The source of the event.
	//
	// Currently only managed CloudWatch Events rules are supported.
	Type *string `json:"type"`
	// Information about the event.
	Parameters interface{} `json:"parameters"`
}

// Specifies a rule for enabling fast snapshot restore.
//
// You can enable fast snapshot restore based on either a count or a time interval.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_FastRestoreRuleProperty struct {
	// The Availability Zones in which to enable fast snapshot restore.
	AvailabilityZones *[]*string `json:"availabilityZones"`
	// The number of snapshots to be enabled with fast snapshot restore.
	Count *float64 `json:"count"`
	// The amount of time to enable fast snapshot restore.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *float64 `json:"interval"`
	// The unit of time for enabling fast snapshot restore.
	IntervalUnit *string `json:"intervalUnit"`
}

// Specifies optional parameters to add to a policy.
//
// The set of valid parameters depends on the combination of policy type and resource type.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_ParametersProperty struct {
	// [EBS Snapshot Management – Instance policies only] Indicates whether to exclude the root volume from snapshots created using [CreateSnapshots](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSnapshots.html) . The default is false.
	ExcludeBootVolume interface{} `json:"excludeBootVolume"`
	// Applies to AMI lifecycle policies only.
	//
	// Indicates whether targeted instances are rebooted when the lifecycle policy runs. `true` indicates that targeted instances are not rebooted when the policy runs. `false` indicates that target instances are rebooted when the policy runs. The default is `true` (instances are not rebooted).
	NoReboot interface{} `json:"noReboot"`
}

// Specifies the configuration of a lifecycle policy.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_PolicyDetailsProperty struct {
	// The actions to be performed when the event-based policy is triggered. You can specify only one action per policy.
	//
	// This parameter is required for event-based policies only. If you are creating a snapshot or AMI policy, omit this parameter.
	Actions interface{} `json:"actions"`
	// The event that triggers the event-based policy.
	//
	// This parameter is required for event-based policies only. If you are creating a snapshot or AMI policy, omit this parameter.
	EventSource interface{} `json:"eventSource"`
	// A set of optional parameters for snapshot and AMI lifecycle policies.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	//
	// > If you are modifying a policy that was created or previously modified using the Amazon Data Lifecycle Manager console, then you must include this parameter and specify either the default values or the new values that you require. You can't omit this parameter or set its values to null.
	Parameters interface{} `json:"parameters"`
	// The valid target resource types and actions a policy can manage.
	//
	// Specify `EBS_SNAPSHOT_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of Amazon EBS snapshots. Specify `IMAGE_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of EBS-backed AMIs. Specify `EVENT_BASED_POLICY` to create an event-based policy that performs specific actions when a defined event occurs in your AWS account .
	//
	// The default is `EBS_SNAPSHOT_MANAGEMENT` .
	PolicyType *string `json:"policyType"`
	// The location of the resources to backup.
	//
	// If the source resources are located in an AWS Region , specify `CLOUD` . If the source resources are located on an Outpost in your account, specify `OUTPOST` .
	//
	// If you specify `OUTPOST` , Amazon Data Lifecycle Manager backs up all resources of the specified type with matching target tags across all of the Outposts in your account.
	ResourceLocations *[]*string `json:"resourceLocations"`
	// The target resource type for snapshot and AMI lifecycle policies.
	//
	// Use `VOLUME` to create snapshots of individual volumes or use `INSTANCE` to create multi-volume snapshots from the volumes for an instance.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	ResourceTypes *[]*string `json:"resourceTypes"`
	// The schedules of policy-defined actions for snapshot and AMI lifecycle policies.
	//
	// A policy can have up to four schedules—one mandatory schedule and up to three optional schedules.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	Schedules interface{} `json:"schedules"`
	// The single tag that identifies targeted resources for this policy.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	TargetTags interface{} `json:"targetTags"`
}

// Specifies the retention rule for a lifecycle policy.
//
// You can retain snapshots based on either a count or a time interval.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_RetainRuleProperty struct {
	// The number of snapshots to retain for each volume, up to a maximum of 1000.
	Count *float64 `json:"count"`
	// The amount of time to retain each snapshot.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *float64 `json:"interval"`
	// The unit of time for time-based retention.
	IntervalUnit *string `json:"intervalUnit"`
}

// Specifies a backup schedule for a snapshot or AMI lifecycle policy.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_ScheduleProperty struct {
	// Copy all user-defined tags on a source volume to snapshots of the volume created by this policy.
	CopyTags interface{} `json:"copyTags"`
	// The creation rule.
	CreateRule interface{} `json:"createRule"`
	// The rule for cross-Region snapshot copies.
	//
	// You can only specify cross-Region copy rules for policies that create snapshots in a Region. If the policy creates snapshots on an Outpost, then you cannot copy the snapshots to a Region or to an Outpost. If the policy creates snapshots in a Region, then snapshots can be copied to up to three Regions or Outposts.
	CrossRegionCopyRules interface{} `json:"crossRegionCopyRules"`
	// The AMI deprecation rule for the schedule.
	DeprecateRule interface{} `json:"deprecateRule"`
	// The rule for enabling fast snapshot restore.
	FastRestoreRule interface{} `json:"fastRestoreRule"`
	// The name of the schedule.
	Name *string `json:"name"`
	// The retention rule.
	RetainRule interface{} `json:"retainRule"`
	// The rule for sharing snapshots with other AWS accounts .
	ShareRules interface{} `json:"shareRules"`
	// The tags to apply to policy-created resources.
	//
	// These user-defined tags are in addition to the AWS -added lifecycle tags.
	TagsToAdd interface{} `json:"tagsToAdd"`
	// A collection of key/value pairs with values determined dynamically when the policy is executed.
	//
	// Keys may be any valid Amazon EC2 tag key. Values must be in one of the two following formats: `$(instance-id)` or `$(timestamp)` . Variable tags are only valid for EBS Snapshot Management – Instance policies.
	VariableTags interface{} `json:"variableTags"`
}

// Specifies a rule for sharing snapshots across AWS accounts .
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicy_ShareRuleProperty struct {
	// The IDs of the AWS accounts with which to share the snapshots.
	TargetAccounts *[]*string `json:"targetAccounts"`
	// The period after which snapshots that are shared with other AWS accounts are automatically unshared.
	UnshareInterval *float64 `json:"unshareInterval"`
	// The unit of time for the automatic unsharing interval.
	UnshareIntervalUnit *string `json:"unshareIntervalUnit"`
}

// Properties for defining a `CfnLifecyclePolicy`.
//
// TODO: EXAMPLE
//
type CfnLifecyclePolicyProps struct {
	// A description of the lifecycle policy.
	//
	// The characters ^[0-9A-Za-z _-]+$ are supported.
	Description *string `json:"description"`
	// The Amazon Resource Name (ARN) of the IAM role used to run the operations specified by the lifecycle policy.
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// The configuration details of the lifecycle policy.
	PolicyDetails interface{} `json:"policyDetails"`
	// The activation state of the lifecycle policy.
	State *string `json:"state"`
	// The tags to apply to the lifecycle policy during creation.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

