package awsfsx

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsfsx"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::FSx::FileCache.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFileCache := awscdk.Aws_fsx.NewCfnFileCache(this, jsii.String("MyCfnFileCache"), &CfnFileCacheProps{
//   	FileCacheType: jsii.String("fileCacheType"),
//   	FileCacheTypeVersion: jsii.String("fileCacheTypeVersion"),
//   	StorageCapacity: jsii.Number(123),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	CopyTagsToDataRepositoryAssociations: jsii.Boolean(false),
//   	DataRepositoryAssociations: []interface{}{
//   		&DataRepositoryAssociationProperty{
//   			DataRepositoryPath: jsii.String("dataRepositoryPath"),
//   			FileCachePath: jsii.String("fileCachePath"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LustreConfiguration: &LustreConfigurationProperty{
//   		DeploymentType: jsii.String("deploymentType"),
//   		MetadataConfiguration: &MetadataConfigurationProperty{
//   			StorageCapacity: jsii.Number(123),
//   		},
//   		PerUnitStorageThroughput: jsii.Number(123),
//
//   		// the properties below are optional
//   		LogConfiguration: &LogConfigurationProperty{
//   			Level: jsii.String("level"),
//
//   			// the properties below are optional
//   			Destination: jsii.String("destination"),
//   		},
//   		MountName: jsii.String("mountName"),
//   		WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filecache.html
//
type CfnFileCache interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsfsx.IFileCacheRef
	awscdk.ITaggableV2
	// The time the cache was created.
	AttrCreationTime() *string
	// A list of IDs of data repository associations that are associated with this cache.
	AttrDataRepositoryAssociationIds() *[]*string
	// The DNS name for the cache.
	AttrDnsName() *string
	// The ID of the file cache.
	AttrFileCacheId() *string
	// The lifecycle status of the cache.
	AttrLifecycle() *string
	AttrLustreConfigurationLogConfiguration() awscdk.IResolvable
	AttrLustreConfigurationMountName() *string
	// Network interface IDs for the cache.
	AttrNetworkInterfaceIds() *[]*string
	// The AWS account ID of the cache owner.
	AttrOwnerId() *string
	// The Amazon Resource Name (ARN) of the file cache.
	AttrResourceArn() *string
	// The VPC ID of the cache.
	AttrVpcId() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	CfnResourceType() *string
	// Whether tags should be copied to data repository associations.
	CopyTagsToDataRepositoryAssociations() interface{}
	SetCopyTagsToDataRepositoryAssociations(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Data repository associations for the cache.
	DataRepositoryAssociations() interface{}
	SetDataRepositoryAssociations(val interface{})
	Env() *interfaces.ResourceEnvironment
	// A reference to a FileCache resource.
	FileCacheRef() *interfacesawsfsx.FileCacheReference
	// The type of cache, which must be LUSTRE.
	FileCacheType() *string
	SetFileCacheType(val *string)
	// The Lustre version of the cache, which must be 2.12.
	FileCacheTypeVersion() *string
	SetFileCacheTypeVersion(val *string)
	// The KMS key ID for encrypting data.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	LustreConfiguration() interface{}
	SetLustreConfiguration(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A list of security group IDs for the cache.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The storage capacity of the cache in gibibytes (GiB).
	StorageCapacity() *float64
	SetStorageCapacity(val *float64)
	// A list of subnet IDs for the cache.
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// Tags to associate with the file cache.
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

// The jsii proxy struct for CfnFileCache
type jsiiProxy_CfnFileCache struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsfsxIFileCacheRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnFileCache) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) AttrDataRepositoryAssociationIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDataRepositoryAssociationIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) AttrDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) AttrFileCacheId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFileCacheId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) AttrLifecycle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) AttrLustreConfigurationLogConfiguration() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLustreConfigurationLogConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) AttrLustreConfigurationMountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLustreConfigurationMountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) AttrNetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrNetworkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) AttrOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) AttrResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) AttrVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) CopyTagsToDataRepositoryAssociations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToDataRepositoryAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) DataRepositoryAssociations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataRepositoryAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) FileCacheRef() *interfacesawsfsx.FileCacheReference {
	var returns *interfacesawsfsx.FileCacheReference
	_jsii_.Get(
		j,
		"fileCacheRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) FileCacheType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileCacheType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) FileCacheTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileCacheTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) LustreConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lustreConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) StorageCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFileCache) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::FSx::FileCache`.
func NewCfnFileCache(scope constructs.Construct, id *string, props *CfnFileCacheProps) CfnFileCache {
	_init_.Initialize()

	if err := validateNewCfnFileCacheParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFileCache{}

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.CfnFileCache",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::FSx::FileCache`.
func NewCfnFileCache_Override(c CfnFileCache, scope constructs.Construct, id *string, props *CfnFileCacheProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.CfnFileCache",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFileCache)SetCopyTagsToDataRepositoryAssociations(val interface{}) {
	if err := j.validateSetCopyTagsToDataRepositoryAssociationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToDataRepositoryAssociations",
		val,
	)
}

func (j *jsiiProxy_CfnFileCache)SetDataRepositoryAssociations(val interface{}) {
	if err := j.validateSetDataRepositoryAssociationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRepositoryAssociations",
		val,
	)
}

func (j *jsiiProxy_CfnFileCache)SetFileCacheType(val *string) {
	if err := j.validateSetFileCacheTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileCacheType",
		val,
	)
}

func (j *jsiiProxy_CfnFileCache)SetFileCacheTypeVersion(val *string) {
	if err := j.validateSetFileCacheTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileCacheTypeVersion",
		val,
	)
}

func (j *jsiiProxy_CfnFileCache)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnFileCache)SetLustreConfiguration(val interface{}) {
	if err := j.validateSetLustreConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lustreConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnFileCache)SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnFileCache)SetStorageCapacity(val *float64) {
	if err := j.validateSetStorageCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnFileCache)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnFileCache)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func CfnFileCache_ArnForFileCache(resource interfacesawsfsx.IFileCacheRef) *string {
	_init_.Initialize()

	if err := validateCfnFileCache_ArnForFileCacheParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.CfnFileCache",
		"arnForFileCache",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Creates a new IFileCacheRef from a fileCacheId.
func CfnFileCache_FromFileCacheId(scope constructs.Construct, id *string, fileCacheId *string) interfacesawsfsx.IFileCacheRef {
	_init_.Initialize()

	if err := validateCfnFileCache_FromFileCacheIdParameters(scope, id, fileCacheId); err != nil {
		panic(err)
	}
	var returns interfacesawsfsx.IFileCacheRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.CfnFileCache",
		"fromFileCacheId",
		[]interface{}{scope, id, fileCacheId},
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
func CfnFileCache_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFileCache_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.CfnFileCache",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnFileCache.
func CfnFileCache_IsCfnFileCache(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFileCache_IsCfnFileCacheParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.CfnFileCache",
		"isCfnFileCache",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnFileCache_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFileCache_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.CfnFileCache",
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
func CfnFileCache_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFileCache_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.CfnFileCache",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFileCache_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_fsx.CfnFileCache",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFileCache) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFileCache) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFileCache) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFileCache) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFileCache) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFileCache) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFileCache) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFileCache) AddResourceDependency(target awscdk.CfnResource, reason *string) {
	if err := c.validateAddResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addResourceDependency",
		[]interface{}{target, reason},
	)
}

func (c *jsiiProxy_CfnFileCache) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnFileCache) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFileCache) CfnPropertyName(cdkPropertyName *string) *string {
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

func (c *jsiiProxy_CfnFileCache) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnFileCache) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnFileCache) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFileCache) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileCache) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFileCache) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFileCache) RemoveResourceDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeResourceDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFileCache) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnFileCache) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnFileCache) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileCache) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFileCache) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnFileCache) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

