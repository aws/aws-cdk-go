package awsdms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdms"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DMS::ReplicationTask` resource creates an AWS DMS replication task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationTask := awscdk.Aws_dms.NewCfnReplicationTask(this, jsii.String("MyCfnReplicationTask"), &CfnReplicationTaskProps{
//   	MigrationType: jsii.String("migrationType"),
//   	ReplicationInstanceArn: jsii.String("replicationInstanceArn"),
//   	SourceEndpointArn: jsii.String("sourceEndpointArn"),
//   	TableMappings: jsii.String("tableMappings"),
//   	TargetEndpointArn: jsii.String("targetEndpointArn"),
//
//   	// the properties below are optional
//   	CdcStartPosition: jsii.String("cdcStartPosition"),
//   	CdcStartTime: jsii.Number(123),
//   	CdcStopPosition: jsii.String("cdcStopPosition"),
//   	ReplicationTaskIdentifier: jsii.String("replicationTaskIdentifier"),
//   	ReplicationTaskSettings: jsii.String("replicationTaskSettings"),
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskData: jsii.String("taskData"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html
//
type CfnReplicationTask interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsdms.IReplicationTaskRef
	awscdk.ITaggable
	AttrId() *string
	// Indicates when you want a change data capture (CDC) operation to start.
	CdcStartPosition() *string
	SetCdcStartPosition(val *string)
	// Indicates the start time for a change data capture (CDC) operation.
	CdcStartTime() *float64
	SetCdcStartTime(val *float64)
	// Indicates when you want a change data capture (CDC) operation to stop.
	CdcStopPosition() *string
	SetCdcStopPosition(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The migration type.
	MigrationType() *string
	SetMigrationType(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of a replication instance.
	ReplicationInstanceArn() *string
	SetReplicationInstanceArn(val *string)
	// An identifier for the replication task.
	ReplicationTaskIdentifier() *string
	SetReplicationTaskIdentifier(val *string)
	// A reference to a ReplicationTask resource.
	ReplicationTaskRef() *interfacesawsdms.ReplicationTaskReference
	// Overall settings for the task, in JSON format.
	ReplicationTaskSettings() *string
	SetReplicationTaskSettings(val *string)
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	ResourceIdentifier() *string
	SetResourceIdentifier(val *string)
	// An Amazon Resource Name (ARN) that uniquely identifies the source endpoint.
	SourceEndpointArn() *string
	SetSourceEndpointArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The table mappings for the task, in JSON format.
	TableMappings() *string
	SetTableMappings(val *string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// One or more tags to be assigned to the replication task.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// An Amazon Resource Name (ARN) that uniquely identifies the target endpoint.
	TargetEndpointArn() *string
	SetTargetEndpointArn(val *string)
	// Supplemental information that the task requires to migrate the data for certain source and target endpoints.
	TaskData() *string
	SetTaskData(val *string)
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

// The jsii proxy struct for CfnReplicationTask
type jsiiProxy_CfnReplicationTask struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsdmsIReplicationTaskRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnReplicationTask) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CdcStartPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStartPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CdcStartTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CdcStopPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStopPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) MigrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) ReplicationInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) ReplicationTaskIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) ReplicationTaskRef() *interfacesawsdms.ReplicationTaskReference {
	var returns *interfacesawsdms.ReplicationTaskReference
	_jsii_.Get(
		j,
		"replicationTaskRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) ReplicationTaskSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) ResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) SourceEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) TableMappings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) TargetEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) TaskData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationTask) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::DMS::ReplicationTask`.
func NewCfnReplicationTask(scope constructs.Construct, id *string, props *CfnReplicationTaskProps) CfnReplicationTask {
	_init_.Initialize()

	if err := validateNewCfnReplicationTaskParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReplicationTask{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::ReplicationTask`.
func NewCfnReplicationTask_Override(c CfnReplicationTask, scope constructs.Construct, id *string, props *CfnReplicationTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetCdcStartPosition(val *string) {
	_jsii_.Set(
		j,
		"cdcStartPosition",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetCdcStartTime(val *float64) {
	_jsii_.Set(
		j,
		"cdcStartTime",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetCdcStopPosition(val *string) {
	_jsii_.Set(
		j,
		"cdcStopPosition",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetMigrationType(val *string) {
	if err := j.validateSetMigrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrationType",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetReplicationInstanceArn(val *string) {
	if err := j.validateSetReplicationInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationInstanceArn",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetReplicationTaskIdentifier(val *string) {
	_jsii_.Set(
		j,
		"replicationTaskIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetReplicationTaskSettings(val *string) {
	_jsii_.Set(
		j,
		"replicationTaskSettings",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetResourceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"resourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetSourceEndpointArn(val *string) {
	if err := j.validateSetSourceEndpointArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceEndpointArn",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetTableMappings(val *string) {
	if err := j.validateSetTableMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableMappings",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetTargetEndpointArn(val *string) {
	if err := j.validateSetTargetEndpointArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetEndpointArn",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationTask)SetTaskData(val *string) {
	_jsii_.Set(
		j,
		"taskData",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnReplicationTask_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationTask_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnReplicationTask.
func CfnReplicationTask_IsCfnReplicationTask(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationTask_IsCfnReplicationTaskParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
		"isCfnReplicationTask",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnReplicationTask_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationTask_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
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
func CfnReplicationTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationTask_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicationTask_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_dms.CfnReplicationTask",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReplicationTask) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnReplicationTask) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicationTask) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicationTask) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnReplicationTask) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnReplicationTask) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnReplicationTask) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnReplicationTask) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnReplicationTask) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnReplicationTask) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnReplicationTask) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnReplicationTask) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationTask) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationTask) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReplicationTask) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicationTask) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnReplicationTask) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnReplicationTask) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationTask) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

