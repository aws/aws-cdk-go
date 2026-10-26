package awstimestream

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awstimestream/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstimestream"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource type definition for AWS::Timestream::InfluxDBBackup.
//
// An on-demand backup of an Amazon Timestream for InfluxDB DB instance or DB cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInfluxDBBackup := awscdk.Aws_timestream.NewCfnInfluxDBBackup(this, jsii.String("MyCfnInfluxDBBackup"), &CfnInfluxDBBackupProps{
//   	DbResourceId: jsii.String("dbResourceId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	RetentionDays: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbbackup.html
//
type CfnInfluxDBBackup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawstimestream.IInfluxDBBackupRef
	awscdk.ITaggableV2
	// The allocated storage of the resource at the time of backup, in GiB.
	AttrAllocatedStorage() *float64
	// The Amazon Resource Name (ARN) of the backup.
	AttrArn() *string
	// The cluster configuration of the resource at the time of backup.
	AttrClusterConfiguration() awscdk.IResolvable
	// The time when the backup was created.
	AttrCreatedAt() *string
	// The DB instance type of the resource at the time of backup.
	AttrDbInstanceType() *string
	// The identifier of the DB parameter group associated with the backup.
	AttrDbParameterGroupId() *string
	// The storage type of the resource at the time of backup.
	AttrDbStorageType() *string
	// The deployment type of the resource that the backup was created from.
	AttrDeploymentType() *string
	// The engine type of the resource that the backup was created from.
	AttrEngineType() *string
	// The date after which the backup will be automatically deleted, as a calendar date in YYYY-MM-DD format.
	AttrExpiresAfter() *string
	// The failover mode of the resource at the time of backup.
	AttrFailoverMode() *string
	// Service-generated unique identifier of the backup.
	AttrId() *string
	// The ARN of the Secrets Manager secret containing the InfluxDB auth parameters.
	AttrInfluxAuthParametersSecretArn() *string
	// The Amazon Web Services KMS key ARN used for encryption of the resource at the time of backup.
	AttrKmsKeyId() *string
	// The log delivery configuration of the resource at the time of backup.
	AttrLogDeliveryConfiguration() awscdk.IResolvable
	// The maintenance schedule of the resource at the time of backup.
	AttrMaintenanceSchedule() awscdk.IResolvable
	// The network type of the resource at the time of backup.
	AttrNetworkType() *string
	// The port number of the resource at the time of backup.
	AttrPort() *float64
	// Indicates whether the resource was publicly accessible at the time of backup.
	AttrPubliclyAccessible() awscdk.IResolvable
	// The current status of the backup.
	AttrStatus() *string
	// The type of backup.
	//
	// Backups created through this resource type are always ON_DEMAND.
	AttrType() *string
	// The VPC security group IDs associated with the resource at the time of backup.
	AttrVpcSecurityGroupIds() *[]*string
	// The VPC subnet IDs associated with the resource at the time of backup.
	AttrVpcSubnetIds() *[]*string
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
	// The id of the DB instance or DB cluster to back up.
	DbResourceId() *string
	SetDbResourceId(val *string)
	Env() *interfaces.ResourceEnvironment
	// A reference to a InfluxDBBackup resource.
	InfluxDbBackupRef() *interfacesawstimestream.InfluxDBBackupReference
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
	// The name of the backup.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The number of days to retain the backup.
	RetentionDays() *float64
	SetRetentionDays(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of key-value pairs to associate with the backup.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
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

// The jsii proxy struct for CfnInfluxDBBackup
type jsiiProxy_CfnInfluxDBBackup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawstimestreamIInfluxDBBackupRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrAllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrAllocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrClusterConfiguration() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrClusterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrDbInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDbInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrDbParameterGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDbParameterGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrDbStorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDbStorageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrDeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDeploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrEngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEngineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrExpiresAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrExpiresAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrFailoverMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFailoverMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrInfluxAuthParametersSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrInfluxAuthParametersSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrLogDeliveryConfiguration() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLogDeliveryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrMaintenanceSchedule() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrMaintenanceSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrNetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNetworkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrPubliclyAccessible() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrPubliclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrVpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrVpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) AttrVpcSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrVpcSubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) DbResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) InfluxDbBackupRef() *interfacesawstimestream.InfluxDBBackupReference {
	var returns *interfacesawstimestream.InfluxDBBackupReference
	_jsii_.Get(
		j,
		"influxDbBackupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) RetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackup) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Timestream::InfluxDBBackup`.
func NewCfnInfluxDBBackup(scope constructs.Construct, id *string, props *CfnInfluxDBBackupProps) CfnInfluxDBBackup {
	_init_.Initialize()

	if err := validateNewCfnInfluxDBBackupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInfluxDBBackup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_timestream.CfnInfluxDBBackup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Timestream::InfluxDBBackup`.
func NewCfnInfluxDBBackup_Override(c CfnInfluxDBBackup, scope constructs.Construct, id *string, props *CfnInfluxDBBackupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_timestream.CfnInfluxDBBackup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInfluxDBBackup)SetDbResourceId(val *string) {
	if err := j.validateSetDbResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbResourceId",
		val,
	)
}

func (j *jsiiProxy_CfnInfluxDBBackup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnInfluxDBBackup)SetRetentionDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionDays",
		val,
	)
}

func (j *jsiiProxy_CfnInfluxDBBackup)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func CfnInfluxDBBackup_ArnForInfluxDBBackup(resource interfacesawstimestream.IInfluxDBBackupRef) *string {
	_init_.Initialize()

	if err := validateCfnInfluxDBBackup_ArnForInfluxDBBackupParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnInfluxDBBackup",
		"arnForInfluxDBBackup",
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
func CfnInfluxDBBackup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInfluxDBBackup_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnInfluxDBBackup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnInfluxDBBackup.
func CfnInfluxDBBackup_IsCfnInfluxDBBackup(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInfluxDBBackup_IsCfnInfluxDBBackupParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnInfluxDBBackup",
		"isCfnInfluxDBBackup",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnInfluxDBBackup_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInfluxDBBackup_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnInfluxDBBackup",
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
func CfnInfluxDBBackup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInfluxDBBackup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnInfluxDBBackup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInfluxDBBackup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_timestream.CfnInfluxDBBackup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInfluxDBBackup) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) AddResourceDependency(target awscdk.CfnResource, reason *string) {
	if err := c.validateAddResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addResourceDependency",
		[]interface{}{target, reason},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) CfnPropertyName(cdkPropertyName *string) *string {
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

func (c *jsiiProxy_CfnInfluxDBBackup) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnInfluxDBBackup) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnInfluxDBBackup) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInfluxDBBackup) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) RemoveResourceDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeResourceDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnInfluxDBBackup) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInfluxDBBackup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInfluxDBBackup) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

