package awselasticache

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The resource representing a serverless cache.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServerlessCache := awscdk.Aws_elasticache.NewCfnServerlessCache(this, jsii.String("MyCfnServerlessCache"), &CfnServerlessCacheProps{
//   	Engine: jsii.String("engine"),
//   	ServerlessCacheName: jsii.String("serverlessCacheName"),
//
//   	// the properties below are optional
//   	CacheUsageLimits: &CacheUsageLimitsProperty{
//   		DataStorage: &DataStorageProperty{
//   			Unit: jsii.String("unit"),
//
//   			// the properties below are optional
//   			Maximum: jsii.Number(123),
//   			Minimum: jsii.Number(123),
//   		},
//   		EcpuPerSecond: &ECPUPerSecondProperty{
//   			Maximum: jsii.Number(123),
//   			Minimum: jsii.Number(123),
//   		},
//   	},
//   	DailySnapshotTime: jsii.String("dailySnapshotTime"),
//   	Description: jsii.String("description"),
//   	Endpoint: &EndpointProperty{
//   		Address: jsii.String("address"),
//   		Port: jsii.String("port"),
//   	},
//   	FinalSnapshotName: jsii.String("finalSnapshotName"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MajorEngineVersion: jsii.String("majorEngineVersion"),
//   	ReaderEndpoint: &EndpointProperty{
//   		Address: jsii.String("address"),
//   		Port: jsii.String("port"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SnapshotArnsToRestore: []*string{
//   		jsii.String("snapshotArnsToRestore"),
//   	},
//   	SnapshotRetentionLimit: jsii.Number(123),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserGroupId: jsii.String("userGroupId"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html
//
type CfnServerlessCache interface {
	awscdk.CfnResource
	IServerlessCacheRef
	awscdk.IInspectable
	awscdk.ITaggableV2
	// The Amazon Resource Name (ARN) of the serverless cache.
	AttrArn() *string
	// When the serverless cache was created.
	AttrCreateTime() *string
	// The DNS hostname of the cache node.
	AttrEndpointAddress() *string
	// The port number that the cache engine is listening on.
	AttrEndpointPort() *string
	// The name and version number of the engine the serverless cache is compatible with.
	AttrFullEngineVersion() *string
	// The DNS hostname of the cache node.
	AttrReaderEndpointAddress() *string
	// The port number that the cache engine is listening on.
	AttrReaderEndpointPort() *string
	// The current status of the serverless cache.
	//
	// The allowed values are CREATING, AVAILABLE, DELETING, CREATE-FAILED and MODIFYING.
	AttrStatus() *string
	// The cache usage limit for the serverless cache.
	CacheUsageLimits() interface{}
	SetCacheUsageLimits(val interface{})
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The daily time that a cache snapshot will be created.
	DailySnapshotTime() *string
	SetDailySnapshotTime(val *string)
	// A description of the serverless cache.
	Description() *string
	SetDescription(val *string)
	// Represents the information required for client programs to connect to a cache node.
	Endpoint() interface{}
	SetEndpoint(val interface{})
	// The engine the serverless cache is compatible with.
	Engine() *string
	SetEngine(val *string)
	// The name of the final snapshot taken of a cache before the cache is deleted.
	FinalSnapshotName() *string
	SetFinalSnapshotName(val *string)
	// The ID of the AWS Key Management Service (KMS) key that is used to encrypt data at rest in the serverless cache.
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
	// The version number of the engine the serverless cache is compatible with.
	MajorEngineVersion() *string
	SetMajorEngineVersion(val *string)
	// The tree node.
	Node() constructs.Node
	// Represents the information required for client programs to connect to a cache node.
	ReaderEndpoint() interface{}
	SetReaderEndpoint(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The IDs of the EC2 security groups associated with the serverless cache.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// The unique identifier of the serverless cache.
	ServerlessCacheName() *string
	SetServerlessCacheName(val *string)
	// A reference to a ServerlessCache resource.
	ServerlessCacheRef() *ServerlessCacheReference
	// The ARN of the snapshot from which to restore data into the new cache.
	SnapshotArnsToRestore() *[]*string
	SetSnapshotArnsToRestore(val *[]*string)
	// The current setting for the number of serverless cache snapshots the system will retain.
	SnapshotRetentionLimit() *float64
	SetSnapshotRetentionLimit(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// If no subnet IDs are given and your VPC is in us-west-1, then ElastiCache will select 2 default subnets across AZs in your VPC.
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// A list of tags to be added to this resource.
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
	// The identifier of the user group associated with the serverless cache.
	UserGroupId() *string
	SetUserGroupId(val *string)
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

// The jsii proxy struct for CfnServerlessCache
type jsiiProxy_CfnServerlessCache struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_IServerlessCacheRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnServerlessCache) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) AttrCreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) AttrEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) AttrEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) AttrFullEngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFullEngineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) AttrReaderEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrReaderEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) AttrReaderEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrReaderEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) CacheUsageLimits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheUsageLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) DailySnapshotTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailySnapshotTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) Endpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) FinalSnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) MajorEngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"majorEngineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) ReaderEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) ServerlessCacheName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessCacheName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) ServerlessCacheRef() *ServerlessCacheReference {
	var returns *ServerlessCacheReference
	_jsii_.Get(
		j,
		"serverlessCacheRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) SnapshotArnsToRestore() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArnsToRestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) SnapshotRetentionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerlessCache) UserGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userGroupId",
		&returns,
	)
	return returns
}


func NewCfnServerlessCache(scope constructs.Construct, id *string, props *CfnServerlessCacheProps) CfnServerlessCache {
	_init_.Initialize()

	if err := validateNewCfnServerlessCacheParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServerlessCache{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticache.CfnServerlessCache",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnServerlessCache_Override(c CfnServerlessCache, scope constructs.Construct, id *string, props *CfnServerlessCacheProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticache.CfnServerlessCache",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetCacheUsageLimits(val interface{}) {
	if err := j.validateSetCacheUsageLimitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheUsageLimits",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetDailySnapshotTime(val *string) {
	_jsii_.Set(
		j,
		"dailySnapshotTime",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetEndpoint(val interface{}) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetFinalSnapshotName(val *string) {
	_jsii_.Set(
		j,
		"finalSnapshotName",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetMajorEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"majorEngineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetReaderEndpoint(val interface{}) {
	if err := j.validateSetReaderEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readerEndpoint",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetServerlessCacheName(val *string) {
	if err := j.validateSetServerlessCacheNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverlessCacheName",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetSnapshotArnsToRestore(val *[]*string) {
	_jsii_.Set(
		j,
		"snapshotArnsToRestore",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetSnapshotRetentionLimit(val *float64) {
	_jsii_.Set(
		j,
		"snapshotRetentionLimit",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnServerlessCache)SetUserGroupId(val *string) {
	_jsii_.Set(
		j,
		"userGroupId",
		val,
	)
}

// Creates a new IServerlessCacheRef from a serverlessCacheName.
func CfnServerlessCache_FromServerlessCacheName(scope constructs.Construct, id *string, serverlessCacheName *string) IServerlessCacheRef {
	_init_.Initialize()

	if err := validateCfnServerlessCache_FromServerlessCacheNameParameters(scope, id, serverlessCacheName); err != nil {
		panic(err)
	}
	var returns IServerlessCacheRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticache.CfnServerlessCache",
		"fromServerlessCacheName",
		[]interface{}{scope, id, serverlessCacheName},
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
func CfnServerlessCache_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServerlessCache_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticache.CfnServerlessCache",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnServerlessCache_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServerlessCache_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticache.CfnServerlessCache",
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
func CfnServerlessCache_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServerlessCache_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticache.CfnServerlessCache",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServerlessCache_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticache.CfnServerlessCache",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServerlessCache) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnServerlessCache) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServerlessCache) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServerlessCache) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnServerlessCache) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnServerlessCache) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnServerlessCache) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnServerlessCache) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnServerlessCache) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnServerlessCache) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnServerlessCache) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnServerlessCache) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServerlessCache) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServerlessCache) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnServerlessCache) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServerlessCache) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnServerlessCache) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnServerlessCache) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServerlessCache) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServerlessCache) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

