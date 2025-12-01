package awsmedialive

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::MediaLive::SignalMap Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSignalMap := awscdk.Aws_medialive.NewCfnSignalMap(this, jsii.String("MyCfnSignalMap"), &CfnSignalMapProps{
//   	DiscoveryEntryPointArn: jsii.String("discoveryEntryPointArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CloudWatchAlarmTemplateGroupIdentifiers: []*string{
//   		jsii.String("cloudWatchAlarmTemplateGroupIdentifiers"),
//   	},
//   	Description: jsii.String("description"),
//   	EventBridgeRuleTemplateGroupIdentifiers: []*string{
//   		jsii.String("eventBridgeRuleTemplateGroupIdentifiers"),
//   	},
//   	ForceRediscovery: jsii.Boolean(false),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-signalmap.html
//
type CfnSignalMap interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsmedialive.ISignalMapRef
	awscdk.ITaggableV2
	// A signal map's ARN (Amazon Resource Name).
	AttrArn() *string
	// An alarm template group's id.
	AttrCloudWatchAlarmTemplateGroupIds() *[]*string
	// The date and time of resource creation.
	AttrCreatedAt() *string
	// Error message associated with a failed creation or failed update attempt of a signal map.
	AttrErrorMessage() *string
	// An eventbridge rule template group's id.
	AttrEventBridgeRuleTemplateGroupIds() *[]*string
	// A map representing an incomplete AWS media workflow as a graph.
	AttrFailedMediaResourceMap() awscdk.IResolvable
	// A signal map's id.
	AttrId() *string
	AttrIdentifier() *string
	// The date and time of latest discovery.
	AttrLastDiscoveredAt() *string
	// Represents the latest successful monitor deployment of a signal map.
	AttrLastSuccessfulMonitorDeployment() awscdk.IResolvable
	// A map representing an AWS media workflow as a graph.
	AttrMediaResourceMap() awscdk.IResolvable
	// The date and time of latest resource modification.
	AttrModifiedAt() *string
	// If true, there are pending monitor changes for this signal map that can be deployed.
	AttrMonitorChangesPendingDeployment() awscdk.IResolvable
	// Represents the latest monitor deployment of a signal map.
	AttrMonitorDeployment() awscdk.IResolvable
	// A signal map's current status, which is dependent on its lifecycle actions or associated jobs.
	AttrStatus() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A cloudwatch alarm template group's identifier.
	CloudWatchAlarmTemplateGroupIdentifiers() *[]*string
	SetCloudWatchAlarmTemplateGroupIdentifiers(val *[]*string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A resource's optional description.
	Description() *string
	SetDescription(val *string)
	// A top-level supported Amazon Web Services resource ARN to discover a signal map from.
	DiscoveryEntryPointArn() *string
	SetDiscoveryEntryPointArn(val *string)
	Env() *interfaces.ResourceEnvironment
	// An eventbridge rule template group's identifier.
	EventBridgeRuleTemplateGroupIdentifiers() *[]*string
	SetEventBridgeRuleTemplateGroupIdentifiers(val *[]*string)
	// If true, will force a rediscovery of a signal map if an unchanged discoveryEntryPointArn is provided.
	ForceRediscovery() interface{}
	SetForceRediscovery(val interface{})
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
	// A resource's name.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A reference to a SignalMap resource.
	SignalMapRef() *interfacesawsmedialive.SignalMapReference
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Represents the tags associated with a resource.
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
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
}

// The jsii proxy struct for CfnSignalMap
type jsiiProxy_CfnSignalMap struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsmedialiveISignalMapRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnSignalMap) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrCloudWatchAlarmTemplateGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrCloudWatchAlarmTemplateGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrErrorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrEventBridgeRuleTemplateGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrEventBridgeRuleTemplateGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrFailedMediaResourceMap() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrFailedMediaResourceMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrLastDiscoveredAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastDiscoveredAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrLastSuccessfulMonitorDeployment() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLastSuccessfulMonitorDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrMediaResourceMap() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrMediaResourceMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrMonitorChangesPendingDeployment() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrMonitorChangesPendingDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrMonitorDeployment() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrMonitorDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) CloudWatchAlarmTemplateGroupIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloudWatchAlarmTemplateGroupIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) DiscoveryEntryPointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryEntryPointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) EventBridgeRuleTemplateGroupIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventBridgeRuleTemplateGroupIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) ForceRediscovery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRediscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) SignalMapRef() *interfacesawsmedialive.SignalMapReference {
	var returns *interfacesawsmedialive.SignalMapReference
	_jsii_.Get(
		j,
		"signalMapRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSignalMap) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaLive::SignalMap`.
func NewCfnSignalMap(scope constructs.Construct, id *string, props *CfnSignalMapProps) CfnSignalMap {
	_init_.Initialize()

	if err := validateNewCfnSignalMapParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSignalMap{}

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnSignalMap",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaLive::SignalMap`.
func NewCfnSignalMap_Override(c CfnSignalMap, scope constructs.Construct, id *string, props *CfnSignalMapProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_medialive.CfnSignalMap",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSignalMap)SetCloudWatchAlarmTemplateGroupIdentifiers(val *[]*string) {
	_jsii_.Set(
		j,
		"cloudWatchAlarmTemplateGroupIdentifiers",
		val,
	)
}

func (j *jsiiProxy_CfnSignalMap)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnSignalMap)SetDiscoveryEntryPointArn(val *string) {
	if err := j.validateSetDiscoveryEntryPointArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoveryEntryPointArn",
		val,
	)
}

func (j *jsiiProxy_CfnSignalMap)SetEventBridgeRuleTemplateGroupIdentifiers(val *[]*string) {
	_jsii_.Set(
		j,
		"eventBridgeRuleTemplateGroupIdentifiers",
		val,
	)
}

func (j *jsiiProxy_CfnSignalMap)SetForceRediscovery(val interface{}) {
	if err := j.validateSetForceRediscoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceRediscovery",
		val,
	)
}

func (j *jsiiProxy_CfnSignalMap)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnSignalMap)SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func CfnSignalMap_ArnForSignalMap(resource interfacesawsmedialive.ISignalMapRef) *string {
	_init_.Initialize()

	if err := validateCfnSignalMap_ArnForSignalMapParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnSignalMap",
		"arnForSignalMap",
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
func CfnSignalMap_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSignalMap_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnSignalMap",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnSignalMap_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSignalMap_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnSignalMap",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnSignalMap.
func CfnSignalMap_IsCfnSignalMap(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSignalMap_IsCfnSignalMapParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnSignalMap",
		"isCfnSignalMap",
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
func CfnSignalMap_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSignalMap_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_medialive.CfnSignalMap",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSignalMap_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_medialive.CfnSignalMap",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSignalMap) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSignalMap) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSignalMap) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSignalMap) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSignalMap) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSignalMap) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSignalMap) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSignalMap) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSignalMap) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnSignalMap) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnSignalMap) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSignalMap) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSignalMap) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSignalMap) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSignalMap) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSignalMap) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnSignalMap) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnSignalMap) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSignalMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSignalMap) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

