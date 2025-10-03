package awselasticache

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ElastiCache::ReplicationGroup` resource creates an Amazon ElastiCache (Valkey or Redis OSS) replication group.
//
// A Valkey or Redis OSS (cluster mode disabled) replication group is a collection of cache clusters, where one of the clusters is a primary read-write cluster and the others are read-only replicas.
//
// A Valkey or Redis OSS (cluster mode enabled) cluster is comprised of from 1 to 90 shards (API/CLI: node groups). Each shard has a primary node and up to 5 read-only replica nodes. The configuration can range from 90 shards and 0 replicas to 15 shards and 5 replicas, which is the maximum number or replicas allowed.
//
// The node or shard limit can be increased to a maximum of 500 per cluster if the engine version is Valkey 7.2 or higher, or Redis OSS 5.0.6 or higher. For example, you can choose to configure a 500 node cluster that ranges between 83 shards (one primary and 5 replicas per shard) and 500 shards (single primary and no replicas). Make sure there are enough available IP addresses to accommodate the increase. Common pitfalls include the subnets in the subnet group have too small a CIDR range or the subnets are shared and heavily used by other clusters. For more information, see [Creating a Subnet Group](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/SubnetGroups.Creating.html) . For versions below 5.0.6, the limit is 250 per cluster.
//
// To request a limit increase, see [Amazon Service Limits](https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html) and choose the limit type *Nodes per cluster per instance type* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationGroup := awscdk.Aws_elasticache.NewCfnReplicationGroup(this, jsii.String("MyCfnReplicationGroup"), &CfnReplicationGroupProps{
//   	ReplicationGroupDescription: jsii.String("replicationGroupDescription"),
//
//   	// the properties below are optional
//   	AtRestEncryptionEnabled: jsii.Boolean(false),
//   	AuthToken: jsii.String("authToken"),
//   	AutomaticFailoverEnabled: jsii.Boolean(false),
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	CacheNodeType: jsii.String("cacheNodeType"),
//   	CacheParameterGroupName: jsii.String("cacheParameterGroupName"),
//   	CacheSecurityGroupNames: []*string{
//   		jsii.String("cacheSecurityGroupNames"),
//   	},
//   	CacheSubnetGroupName: jsii.String("cacheSubnetGroupName"),
//   	ClusterMode: jsii.String("clusterMode"),
//   	DataTieringEnabled: jsii.Boolean(false),
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	GlobalReplicationGroupId: jsii.String("globalReplicationGroupId"),
//   	IpDiscovery: jsii.String("ipDiscovery"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LogDeliveryConfigurations: []interface{}{
//   		&LogDeliveryConfigurationRequestProperty{
//   			DestinationDetails: &DestinationDetailsProperty{
//   				CloudWatchLogsDetails: &CloudWatchLogsDestinationDetailsProperty{
//   					LogGroup: jsii.String("logGroup"),
//   				},
//   				KinesisFirehoseDetails: &KinesisFirehoseDestinationDetailsProperty{
//   					DeliveryStream: jsii.String("deliveryStream"),
//   				},
//   			},
//   			DestinationType: jsii.String("destinationType"),
//   			LogFormat: jsii.String("logFormat"),
//   			LogType: jsii.String("logType"),
//   		},
//   	},
//   	MultiAzEnabled: jsii.Boolean(false),
//   	NetworkType: jsii.String("networkType"),
//   	NodeGroupConfiguration: []interface{}{
//   		&NodeGroupConfigurationProperty{
//   			NodeGroupId: jsii.String("nodeGroupId"),
//   			PrimaryAvailabilityZone: jsii.String("primaryAvailabilityZone"),
//   			ReplicaAvailabilityZones: []*string{
//   				jsii.String("replicaAvailabilityZones"),
//   			},
//   			ReplicaCount: jsii.Number(123),
//   			Slots: jsii.String("slots"),
//   		},
//   	},
//   	NotificationTopicArn: jsii.String("notificationTopicArn"),
//   	NumCacheClusters: jsii.Number(123),
//   	NumNodeGroups: jsii.Number(123),
//   	Port: jsii.Number(123),
//   	PreferredCacheClusterAZs: []*string{
//   		jsii.String("preferredCacheClusterAZs"),
//   	},
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	PrimaryClusterId: jsii.String("primaryClusterId"),
//   	ReplicasPerNodeGroup: jsii.Number(123),
//   	ReplicationGroupId: jsii.String("replicationGroupId"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SnapshotArns: []*string{
//   		jsii.String("snapshotArns"),
//   	},
//   	SnapshotName: jsii.String("snapshotName"),
//   	SnapshotRetentionLimit: jsii.Number(123),
//   	SnapshottingClusterId: jsii.String("snapshottingClusterId"),
//   	SnapshotWindow: jsii.String("snapshotWindow"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitEncryptionEnabled: jsii.Boolean(false),
//   	TransitEncryptionMode: jsii.String("transitEncryptionMode"),
//   	UserGroupIds: []*string{
//   		jsii.String("userGroupIds"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html
//
type CfnReplicationGroup interface {
	awscdk.CfnResource
	IReplicationGroupRef
	awscdk.IInspectable
	awscdk.ITaggable
	// A flag that enables encryption at rest when set to `true` .
	AtRestEncryptionEnabled() interface{}
	SetAtRestEncryptionEnabled(val interface{})
	// The DNS hostname of the cache node.
	//
	// > Valkey or Redis OSS (cluster mode disabled) replication groups don't have this attribute. Therefore, `Fn::GetAtt` returns a value for this attribute only if the replication group is clustered. Otherwise, `Fn::GetAtt` fails. For Valkey or Redis OSS (cluster mode disabled) replication groups, use the `PrimaryEndpoint` or `ReadEndpoint` attributes.
	AttrConfigurationEndPointAddress() *string
	// The port number that the cache engine is listening on.
	AttrConfigurationEndPointPort() *string
	// The DNS address of the primary read-write cache node.
	AttrPrimaryEndPointAddress() *string
	// The number of the port that the primary read-write cache engine is listening on.
	AttrPrimaryEndPointPort() *string
	// A string with a list of endpoints for the primary and read-only replicas.
	//
	// The order of the addresses maps to the order of the ports from the `ReadEndPoint.Ports` attribute.
	AttrReadEndPointAddresses() *string
	// A string with a list of endpoints for the read-only replicas.
	//
	// The order of the addresses maps to the order of the ports from the `ReadEndPoint.Ports` attribute.
	AttrReadEndPointAddressesList() *[]*string
	// A string with a list of ports for the read-only replicas.
	//
	// The order of the ports maps to the order of the addresses from the `ReadEndPoint.Addresses` attribute.
	AttrReadEndPointPorts() *string
	// A string with a list of ports for the read-only replicas.
	//
	// The order of the ports maps to the order of the addresses from the ReadEndPoint.Addresses attribute.
	AttrReadEndPointPortsList() *[]*string
	// The address of the reader endpoint.
	AttrReaderEndPointAddress() *string
	// The port used by the reader endpoint.
	AttrReaderEndPointPort() *string
	// *Reserved parameter.* The password used to access a password protected server.
	AuthToken() *string
	SetAuthToken(val *string)
	// Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails.
	AutomaticFailoverEnabled() interface{}
	SetAutomaticFailoverEnabled(val interface{})
	// If you are running Valkey 7.2 or later, or Redis OSS 6.0 or later, set this parameter to yes if you want to opt-in to the next minor version upgrade campaign. This parameter is disabled for previous versions.
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	// The compute and memory capacity of the nodes in the node group (shard).
	CacheNodeType() *string
	SetCacheNodeType(val *string)
	// The name of the parameter group to associate with this replication group.
	CacheParameterGroupName() *string
	SetCacheParameterGroupName(val *string)
	// A list of cache security group names to associate with this replication group.
	CacheSecurityGroupNames() *[]*string
	SetCacheSecurityGroupNames(val *[]*string)
	// The name of the cache subnet group to be used for the replication group.
	CacheSubnetGroupName() *string
	SetCacheSubnetGroupName(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The mode can be enabled or disabled.
	ClusterMode() *string
	SetClusterMode(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Enables data tiering.
	DataTieringEnabled() interface{}
	SetDataTieringEnabled(val interface{})
	// The name of the cache engine to be used for the clusters in this replication group.
	Engine() *string
	SetEngine(val *string)
	// The version number of the cache engine to be used for the clusters in this replication group.
	EngineVersion() *string
	SetEngineVersion(val *string)
	// The name of the Global datastore.
	GlobalReplicationGroupId() *string
	SetGlobalReplicationGroupId(val *string)
	// The network type you choose when creating a replication group, either `ipv4` | `ipv6` .
	IpDiscovery() *string
	SetIpDiscovery(val *string)
	// The ID of the KMS key used to encrypt the disk on the cluster.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// Specifies the destination, format and type of the logs.
	LogDeliveryConfigurations() interface{}
	SetLogDeliveryConfigurations(val interface{})
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
	// A flag indicating if you have Multi-AZ enabled to enhance fault tolerance.
	MultiAzEnabled() interface{}
	SetMultiAzEnabled(val interface{})
	// Must be either `ipv4` | `ipv6` | `dual_stack` .
	NetworkType() *string
	SetNetworkType(val *string)
	// The tree node.
	Node() constructs.Node
	// `NodeGroupConfiguration` is a property of the `AWS::ElastiCache::ReplicationGroup` resource that configures an Amazon ElastiCache (ElastiCache) Valkey or Redis OSS cluster node group.
	NodeGroupConfiguration() interface{}
	SetNodeGroupConfiguration(val interface{})
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
	NotificationTopicArn() *string
	SetNotificationTopicArn(val *string)
	// The number of clusters this replication group initially has.
	NumCacheClusters() *float64
	SetNumCacheClusters(val *float64)
	// An optional parameter that specifies the number of node groups (shards) for this Valkey or Redis OSS (cluster mode enabled) replication group.
	NumNodeGroups() *float64
	SetNumNodeGroups(val *float64)
	// The port number on which each member of the replication group accepts connections.
	Port() *float64
	SetPort(val *float64)
	// A list of EC2 Availability Zones in which the replication group's clusters are created.
	PreferredCacheClusterAZs() *[]*string
	SetPreferredCacheClusterAZs(val *[]*string)
	// Specifies the weekly time range during which maintenance on the cluster is performed.
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	// The identifier of the cluster that serves as the primary for this replication group.
	PrimaryClusterId() *string
	SetPrimaryClusterId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// An optional parameter that specifies the number of replica nodes in each node group (shard).
	ReplicasPerNodeGroup() *float64
	SetReplicasPerNodeGroup(val *float64)
	// A user-created description for the replication group.
	ReplicationGroupDescription() *string
	SetReplicationGroupDescription(val *string)
	// The replication group identifier.
	//
	// This parameter is stored as a lowercase string.
	ReplicationGroupId() *string
	SetReplicationGroupId(val *string)
	// A reference to a ReplicationGroup resource.
	ReplicationGroupRef() *ReplicationGroupReference
	// One or more Amazon VPC security groups associated with this replication group.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// A list of Amazon Resource Names (ARN) that uniquely identify the Valkey or Redis OSS RDB snapshot files stored in Amazon S3.
	SnapshotArns() *[]*string
	SetSnapshotArns(val *[]*string)
	// The name of a snapshot from which to restore data into the new replication group.
	SnapshotName() *string
	SetSnapshotName(val *string)
	// The number of days for which ElastiCache retains automatic snapshots before deleting them.
	SnapshotRetentionLimit() *float64
	SetSnapshotRetentionLimit(val *float64)
	// The cluster ID that is used as the daily snapshot source for the replication group.
	SnapshottingClusterId() *string
	SetSnapshottingClusterId(val *string)
	// The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).
	SnapshotWindow() *string
	SetSnapshotWindow(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// A list of tags to be added to this resource.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// A flag that enables in-transit encryption when set to `true` .
	TransitEncryptionEnabled() interface{}
	SetTransitEncryptionEnabled(val interface{})
	// A setting that allows you to migrate your clients to use in-transit encryption, with no downtime.
	TransitEncryptionMode() *string
	SetTransitEncryptionMode(val *string)
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
	// The ID of user group to associate with the replication group.
	UserGroupIds() *[]*string
	SetUserGroupIds(val *[]*string)
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

// The jsii proxy struct for CfnReplicationGroup
type jsiiProxy_CfnReplicationGroup struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_IReplicationGroupRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnReplicationGroup) AtRestEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atRestEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AttrConfigurationEndPointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConfigurationEndPointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AttrConfigurationEndPointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConfigurationEndPointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AttrPrimaryEndPointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPrimaryEndPointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AttrPrimaryEndPointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPrimaryEndPointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AttrReadEndPointAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrReadEndPointAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AttrReadEndPointAddressesList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrReadEndPointAddressesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AttrReadEndPointPorts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrReadEndPointPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AttrReadEndPointPortsList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrReadEndPointPortsList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AttrReaderEndPointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrReaderEndPointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AttrReaderEndPointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrReaderEndPointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AuthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AutomaticFailoverEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticFailoverEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) CacheNodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) CacheParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) CacheSecurityGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheSecurityGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) CacheSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) ClusterMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) DataTieringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTieringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) GlobalReplicationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalReplicationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) IpDiscovery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) LogDeliveryConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeliveryConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) MultiAzEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) NodeGroupConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeGroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) NotificationTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) NumCacheClusters() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCacheClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) NumNodeGroups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numNodeGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) PreferredCacheClusterAZs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredCacheClusterAZs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) PrimaryClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) ReplicasPerNodeGroup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasPerNodeGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) ReplicationGroupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) ReplicationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) ReplicationGroupRef() *ReplicationGroupReference {
	var returns *ReplicationGroupReference
	_jsii_.Get(
		j,
		"replicationGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) SnapshotArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) SnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) SnapshotRetentionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) SnapshottingClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshottingClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) SnapshotWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) TransitEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) TransitEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroup) UserGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userGroupIds",
		&returns,
	)
	return returns
}


func NewCfnReplicationGroup(scope constructs.Construct, id *string, props *CfnReplicationGroupProps) CfnReplicationGroup {
	_init_.Initialize()

	if err := validateNewCfnReplicationGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReplicationGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticache.CfnReplicationGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnReplicationGroup_Override(c CfnReplicationGroup, scope constructs.Construct, id *string, props *CfnReplicationGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticache.CfnReplicationGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetAtRestEncryptionEnabled(val interface{}) {
	if err := j.validateSetAtRestEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atRestEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetAuthToken(val *string) {
	_jsii_.Set(
		j,
		"authToken",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetAutomaticFailoverEnabled(val interface{}) {
	if err := j.validateSetAutomaticFailoverEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticFailoverEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetCacheNodeType(val *string) {
	_jsii_.Set(
		j,
		"cacheNodeType",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetCacheParameterGroupName(val *string) {
	_jsii_.Set(
		j,
		"cacheParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetCacheSecurityGroupNames(val *[]*string) {
	_jsii_.Set(
		j,
		"cacheSecurityGroupNames",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetCacheSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"cacheSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetClusterMode(val *string) {
	_jsii_.Set(
		j,
		"clusterMode",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetDataTieringEnabled(val interface{}) {
	if err := j.validateSetDataTieringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTieringEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetEngine(val *string) {
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetGlobalReplicationGroupId(val *string) {
	_jsii_.Set(
		j,
		"globalReplicationGroupId",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetIpDiscovery(val *string) {
	_jsii_.Set(
		j,
		"ipDiscovery",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetLogDeliveryConfigurations(val interface{}) {
	if err := j.validateSetLogDeliveryConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logDeliveryConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetMultiAzEnabled(val interface{}) {
	if err := j.validateSetMultiAzEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAzEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetNetworkType(val *string) {
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetNodeGroupConfiguration(val interface{}) {
	if err := j.validateSetNodeGroupConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroupConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetNotificationTopicArn(val *string) {
	_jsii_.Set(
		j,
		"notificationTopicArn",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetNumCacheClusters(val *float64) {
	_jsii_.Set(
		j,
		"numCacheClusters",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetNumNodeGroups(val *float64) {
	_jsii_.Set(
		j,
		"numNodeGroups",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetPreferredCacheClusterAZs(val *[]*string) {
	_jsii_.Set(
		j,
		"preferredCacheClusterAZs",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetPrimaryClusterId(val *string) {
	_jsii_.Set(
		j,
		"primaryClusterId",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetReplicasPerNodeGroup(val *float64) {
	_jsii_.Set(
		j,
		"replicasPerNodeGroup",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetReplicationGroupDescription(val *string) {
	if err := j.validateSetReplicationGroupDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationGroupDescription",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetReplicationGroupId(val *string) {
	_jsii_.Set(
		j,
		"replicationGroupId",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetSnapshotArns(val *[]*string) {
	_jsii_.Set(
		j,
		"snapshotArns",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetSnapshotName(val *string) {
	_jsii_.Set(
		j,
		"snapshotName",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetSnapshotRetentionLimit(val *float64) {
	_jsii_.Set(
		j,
		"snapshotRetentionLimit",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetSnapshottingClusterId(val *string) {
	_jsii_.Set(
		j,
		"snapshottingClusterId",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetSnapshotWindow(val *string) {
	_jsii_.Set(
		j,
		"snapshotWindow",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetTransitEncryptionEnabled(val interface{}) {
	if err := j.validateSetTransitEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetTransitEncryptionMode(val *string) {
	_jsii_.Set(
		j,
		"transitEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationGroup)SetUserGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"userGroupIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnReplicationGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationGroup_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticache.CfnReplicationGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnReplicationGroup_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationGroup_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticache.CfnReplicationGroup",
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
func CfnReplicationGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticache.CfnReplicationGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicationGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticache.CfnReplicationGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReplicationGroup) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnReplicationGroup) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnReplicationGroup) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationGroup) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationGroup) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnReplicationGroup) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnReplicationGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationGroup) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

