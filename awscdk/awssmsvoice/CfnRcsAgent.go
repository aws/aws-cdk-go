package awssmsvoice

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssmsvoice/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssmsvoice"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AWS RCS Agent in AWS End User Messaging SMS.
//
// The RCS agent is the top-level resource representing your brand for RCS messaging, and serves as an origination identity for sending RCS messages. A newly created agent is in CREATED status; it advances through TESTING, PARTIAL and ACTIVE only as testing and country launch registrations are submitted and approved by carriers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRcsAgent := awscdk.Aws_smsvoice.NewCfnRcsAgent(this, jsii.String("MyCfnRcsAgent"), &CfnRcsAgentProps{
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	OptOutListName: jsii.String("optOutListName"),
//   	SelfManagedOptOutsEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TwoWayChannelArn: jsii.String("twoWayChannelArn"),
//   	TwoWayChannelRole: jsii.String("twoWayChannelRole"),
//   	TwoWayEnabled: jsii.Boolean(false),
//   	TwoWayMediaS3BucketName: jsii.String("twoWayMediaS3BucketName"),
//   	TwoWayMediaS3KeyPrefix: jsii.String("twoWayMediaS3KeyPrefix"),
//   	TwoWayMediaS3Role: jsii.String("twoWayMediaS3Role"),
//   	TwoWayRcsEventsEnabled: []*string{
//   		jsii.String("twoWayRcsEventsEnabled"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-rcsagent.html
//
type CfnRcsAgent interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawssmsvoice.IRcsAgentRef
	awscdk.ITaggableV2
	// The time when the RCS agent was created.
	AttrCreatedTimestamp() *string
	// The unique identifier of the pool associated with the RCS agent.
	//
	// This is populated only while the agent is associated with a pool, which is done through the pool APIs rather than through this resource.
	AttrPoolId() *string
	// The Amazon Resource Name (ARN) of the RCS agent.
	AttrRcsAgentArn() *string
	// The unique identifier for the RCS agent, in the format rcs- followed by a hexadecimal string.
	AttrRcsAgentId() *string
	// The current status of the RCS agent.
	//
	// CREATED means the agent exists but no registration has been submitted. PENDING means a registration is awaiting processing. TESTING means the testing registration is approved and the agent can message registered test devices. PARTIAL means at least one country launch is active. ACTIVE means all submitted country launches are active. DELETED means the agent has been deleted.
	AttrStatus() *string
	// Information about the testing agent (RCS for Business ID) associated with the RCS agent.
	//
	// A testing agent is created by submitting a testing registration, and allows sending to registered test devices without carrier approval.
	AttrTestingAgent() awscdk.IResolvable
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// When set to true the RCS agent can't be deleted.
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
	Env() *interfaces.ResourceEnvironment
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
	// The name of the opt-out list associated with the RCS agent.
	OptOutListName() *string
	SetOptOutListName(val *string)
	// A reference to a RcsAgent resource.
	RcsAgentRef() *interfacesawssmsvoice.RcsAgentReference
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// When set to true you're responsible for responding to HELP and STOP requests, and for tracking and honoring opt-out requests.
	SelfManagedOptOutsEnabled() interface{}
	SetSelfManagedOptOutsEnabled(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to the RCS agent.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// The Amazon Resource Name (ARN) of the two way channel where inbound messages are delivered.
	TwoWayChannelArn() *string
	SetTwoWayChannelArn(val *string)
	// The Amazon Resource Name (ARN) of an IAM role for the service to assume in order to post inbound messages to the two way channel.
	TwoWayChannelRole() *string
	SetTwoWayChannelRole(val *string)
	// When set to true two-way messaging is enabled for the RCS agent.
	TwoWayEnabled() interface{}
	SetTwoWayEnabled(val interface{})
	// The name of the Amazon S3 bucket where inbound RCS media objects are written.
	TwoWayMediaS3BucketName() *string
	SetTwoWayMediaS3BucketName(val *string)
	// The key prefix used for inbound RCS media objects in the Amazon S3 bucket.
	TwoWayMediaS3KeyPrefix() *string
	SetTwoWayMediaS3KeyPrefix(val *string)
	// The Amazon Resource Name (ARN) of the IAM role used to write inbound RCS media files to the Amazon S3 bucket.
	TwoWayMediaS3Role() *string
	SetTwoWayMediaS3Role(val *string)
	// The list of RCS event types enabled for two-way messaging.
	TwoWayRcsEventsEnabled() *[]*string
	SetTwoWayRcsEventsEnabled(val *[]*string)
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
	// This method has been renamed to `addResourceDependency` to more clearly
	// set it apart from `construct.node.addDependency`. See the documentation
	// of that function for more details.
	// Deprecated: Use `addResourceDependency` instead.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method has been renamed to `addResourceDependency`, which makes it
	// more clear that this method operates at a different level from the
	// construct-level `construct.node.addDependency()` mechanism.
	// Deprecated: Use `addResourceDependency` instead.
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
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method only adds dependencies between L1 resources. If you are
	// looking for a generic construct-to-construct dependency mechanism that works
	// for all constructs including L2s, use `construct.node.addDependency` instead.
	AddResourceDependency(target awscdk.CfnResource, reason *string)
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
	// Retrieves an array of resources and stacks this resource depends on.
	//
	// For resources depended on directly, returns the `CfnResource` object. For
	// dependencies on other stacks, returns the `Stack` object. The order of the
	// array is not guaranteed.
	ObtainDependencies() *[]interface{}
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	// Deprecated: Use `removeResourceDependency` instead.
	RemoveDependency(target awscdk.CfnResource)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveResourceDependency(target awscdk.CfnResource)
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

// The jsii proxy struct for CfnRcsAgent
type jsiiProxy_CfnRcsAgent struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawssmsvoiceIRcsAgentRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnRcsAgent) AttrCreatedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) AttrPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) AttrRcsAgentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRcsAgentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) AttrRcsAgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRcsAgentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) AttrTestingAgent() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrTestingAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) OptOutListName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optOutListName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) RcsAgentRef() *interfacesawssmsvoice.RcsAgentReference {
	var returns *interfacesawssmsvoice.RcsAgentReference
	_jsii_.Get(
		j,
		"rcsAgentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) SelfManagedOptOutsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedOptOutsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) TwoWayChannelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) TwoWayChannelRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) TwoWayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twoWayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) TwoWayMediaS3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayMediaS3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) TwoWayMediaS3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayMediaS3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) TwoWayMediaS3Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayMediaS3Role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) TwoWayRcsEventsEnabled() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"twoWayRcsEventsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRcsAgent) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::SMSVOICE::RcsAgent`.
func NewCfnRcsAgent(scope constructs.Construct, id *string, props *CfnRcsAgentProps) CfnRcsAgent {
	_init_.Initialize()

	if err := validateNewCfnRcsAgentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRcsAgent{}

	_jsii_.Create(
		"aws-cdk-lib.aws_smsvoice.CfnRcsAgent",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SMSVOICE::RcsAgent`.
func NewCfnRcsAgent_Override(c CfnRcsAgent, scope constructs.Construct, id *string, props *CfnRcsAgentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_smsvoice.CfnRcsAgent",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRcsAgent)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnRcsAgent)SetOptOutListName(val *string) {
	_jsii_.Set(
		j,
		"optOutListName",
		val,
	)
}

func (j *jsiiProxy_CfnRcsAgent)SetSelfManagedOptOutsEnabled(val interface{}) {
	if err := j.validateSetSelfManagedOptOutsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfManagedOptOutsEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnRcsAgent)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnRcsAgent)SetTwoWayChannelArn(val *string) {
	_jsii_.Set(
		j,
		"twoWayChannelArn",
		val,
	)
}

func (j *jsiiProxy_CfnRcsAgent)SetTwoWayChannelRole(val *string) {
	_jsii_.Set(
		j,
		"twoWayChannelRole",
		val,
	)
}

func (j *jsiiProxy_CfnRcsAgent)SetTwoWayEnabled(val interface{}) {
	if err := j.validateSetTwoWayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnRcsAgent)SetTwoWayMediaS3BucketName(val *string) {
	_jsii_.Set(
		j,
		"twoWayMediaS3BucketName",
		val,
	)
}

func (j *jsiiProxy_CfnRcsAgent)SetTwoWayMediaS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"twoWayMediaS3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnRcsAgent)SetTwoWayMediaS3Role(val *string) {
	_jsii_.Set(
		j,
		"twoWayMediaS3Role",
		val,
	)
}

func (j *jsiiProxy_CfnRcsAgent)SetTwoWayRcsEventsEnabled(val *[]*string) {
	_jsii_.Set(
		j,
		"twoWayRcsEventsEnabled",
		val,
	)
}

func CfnRcsAgent_ArnForRcsAgent(resource interfacesawssmsvoice.IRcsAgentRef) *string {
	_init_.Initialize()

	if err := validateCfnRcsAgent_ArnForRcsAgentParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_smsvoice.CfnRcsAgent",
		"arnForRcsAgent",
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
func CfnRcsAgent_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRcsAgent_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_smsvoice.CfnRcsAgent",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnRcsAgent.
func CfnRcsAgent_IsCfnRcsAgent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRcsAgent_IsCfnRcsAgentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_smsvoice.CfnRcsAgent",
		"isCfnRcsAgent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnRcsAgent_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRcsAgent_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_smsvoice.CfnRcsAgent",
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
func CfnRcsAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRcsAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_smsvoice.CfnRcsAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRcsAgent_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_smsvoice.CfnRcsAgent",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRcsAgent) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRcsAgent) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRcsAgent) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRcsAgent) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRcsAgent) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRcsAgent) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRcsAgent) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRcsAgent) AddResourceDependency(target awscdk.CfnResource, reason *string) {
	if err := c.validateAddResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addResourceDependency",
		[]interface{}{target, reason},
	)
}

func (c *jsiiProxy_CfnRcsAgent) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnRcsAgent) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnRcsAgent) CfnPropertyName(cdkPropertyName *string) *string {
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

func (c *jsiiProxy_CfnRcsAgent) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnRcsAgent) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnRcsAgent) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRcsAgent) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRcsAgent) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRcsAgent) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRcsAgent) RemoveResourceDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeResourceDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRcsAgent) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnRcsAgent) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnRcsAgent) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRcsAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRcsAgent) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnRcsAgent) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

