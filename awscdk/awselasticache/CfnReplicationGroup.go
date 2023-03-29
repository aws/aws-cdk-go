package awselasticache

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::ElastiCache::ReplicationGroup`.
//
// The `AWS::ElastiCache::ReplicationGroup` resource creates an Amazon ElastiCache Redis replication group. A Redis (cluster mode disabled) replication group is a collection of cache clusters, where one of the clusters is a primary read-write cluster and the others are read-only replicas.
//
// A Redis (cluster mode enabled) cluster is comprised of from 1 to 90 shards (API/CLI: node groups). Each shard has a primary node and up to 5 read-only replica nodes. The configuration can range from 90 shards and 0 replicas to 15 shards and 5 replicas, which is the maximum number or replicas allowed.
//
// The node or shard limit can be increased to a maximum of 500 per cluster if the Redis engine version is 5.0.6 or higher. For example, you can choose to configure a 500 node cluster that ranges between 83 shards (one primary and 5 replicas per shard) and 500 shards (single primary and no replicas). Make sure there are enough available IP addresses to accommodate the increase. Common pitfalls include the subnets in the subnet group have too small a CIDR range or the subnets are shared and heavily used by other clusters. For more information, see [Creating a Subnet Group](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.Creating.html) . For versions below 5.0.6, the limit is 250 per cluster.
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
type CfnReplicationGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A flag that enables encryption at rest when set to `true` .
	//
	// You cannot modify the value of `AtRestEncryptionEnabled` after the replication group is created. To enable encryption at rest on a replication group you must set `AtRestEncryptionEnabled` to `true` when you create the replication group.
	//
	// *Required:* Only available when creating a replication group in an Amazon VPC using redis version `3.2.6` or `4.x` onward.
	//
	// Default: `false`.
	AtRestEncryptionEnabled() interface{}
	SetAtRestEncryptionEnabled(val interface{})
	// The DNS hostname of the cache node.
	//
	// > Redis (cluster mode disabled) replication groups don't have this attribute. Therefore, `Fn::GetAtt` returns a value for this attribute only if the replication group is clustered. Otherwise, `Fn::GetAtt` fails. For Redis (cluster mode disabled) replication groups, use the `PrimaryEndpoint` or `ReadEndpoint` attributes.
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
	//
	// `AuthToken` can be specified only on replication groups where `TransitEncryptionEnabled` is `true` . For more information, see [Authenticating Users with the Redis AUTH Command](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/auth.html) .
	//
	// > For HIPAA compliance, you must specify `TransitEncryptionEnabled` as `true` , an `AuthToken` , and a `CacheSubnetGroup` .
	//
	// Password constraints:
	//
	// - Must be only printable ASCII characters.
	// - Must be at least 16 characters and no more than 128 characters in length.
	// - Nonalphanumeric characters are restricted to (!, &, #, $, ^, <, >, -, ).
	//
	// For more information, see [AUTH password](https://docs.aws.amazon.com/http://redis.io/commands/AUTH) at http://redis.io/commands/AUTH.
	//
	// > If ADDING the AuthToken, update requires [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	AuthToken() *string
	SetAuthToken(val *string)
	// Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails.
	//
	// `AutomaticFailoverEnabled` must be enabled for Redis (cluster mode enabled) replication groups.
	//
	// Default: false.
	AutomaticFailoverEnabled() interface{}
	SetAutomaticFailoverEnabled(val interface{})
	// If you are running Redis engine version 6.0 or later, set this parameter to yes if you want to opt-in to the next minor version upgrade campaign. This parameter is disabled for previous versions.
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	// The compute and memory capacity of the nodes in the node group (shard).
	//
	// The following node types are supported by ElastiCache. Generally speaking, the current generation types provide more memory and computational power at lower cost when compared to their equivalent previous generation counterparts.
	//
	// - General purpose:
	//
	// - Current generation:
	//
	// *M6g node types:* `cache.m6g.large` , `cache.m6g.xlarge` , `cache.m6g.2xlarge` , `cache.m6g.4xlarge` , `cache.m6g.12xlarge` , `cache.m6g.24xlarge`
	//
	// *M5 node types:* `cache.m5.large` , `cache.m5.xlarge` , `cache.m5.2xlarge` , `cache.m5.4xlarge` , `cache.m5.12xlarge` , `cache.m5.24xlarge`
	//
	// *M4 node types:* `cache.m4.large` , `cache.m4.xlarge` , `cache.m4.2xlarge` , `cache.m4.4xlarge` , `cache.m4.10xlarge`
	//
	// *T4g node types:* `cache.t4g.micro` , `cache.t4g.small` , `cache.t4g.medium`
	//
	// *T3 node types:* `cache.t3.micro` , `cache.t3.small` , `cache.t3.medium`
	//
	// *T2 node types:* `cache.t2.micro` , `cache.t2.small` , `cache.t2.medium`
	// - Previous generation: (not recommended)
	//
	// *T1 node types:* `cache.t1.micro`
	//
	// *M1 node types:* `cache.m1.small` , `cache.m1.medium` , `cache.m1.large` , `cache.m1.xlarge`
	//
	// *M3 node types:* `cache.m3.medium` , `cache.m3.large` , `cache.m3.xlarge` , `cache.m3.2xlarge`
	// - Compute optimized:
	//
	// - Previous generation: (not recommended)
	//
	// *C1 node types:* `cache.c1.xlarge`
	// - Memory optimized:
	//
	// - Current generation:
	//
	// *R6gd node types:* `cache.r6gd.xlarge` , `cache.r6gd.2xlarge` , `cache.r6gd.4xlarge` , `cache.r6gd.8xlarge` , `cache.r6gd.12xlarge` , `cache.r6gd.16xlarge`
	//
	// > The `r6gd` family is available in the following regions: `us-east-2` , `us-east-1` , `us-west-2` , `us-west-1` , `eu-west-1` , `eu-central-1` , `ap-northeast-1` , `ap-southeast-1` , `ap-southeast-2` .
	//
	// *R6g node types:* `cache.r6g.large` , `cache.r6g.xlarge` , `cache.r6g.2xlarge` , `cache.r6g.4xlarge` , `cache.r6g.12xlarge` , `cache.r6g.24xlarge`
	//
	// *R5 node types:* `cache.r5.large` , `cache.r5.xlarge` , `cache.r5.2xlarge` , `cache.r5.4xlarge` , `cache.r5.12xlarge` , `cache.r5.24xlarge`
	//
	// *R4 node types:* `cache.r4.large` , `cache.r4.xlarge` , `cache.r4.2xlarge` , `cache.r4.4xlarge` , `cache.r4.8xlarge` , `cache.r4.16xlarge`
	// - Previous generation: (not recommended)
	//
	// *M2 node types:* `cache.m2.xlarge` , `cache.m2.2xlarge` , `cache.m2.4xlarge`
	//
	// *R3 node types:* `cache.r3.large` , `cache.r3.xlarge` , `cache.r3.2xlarge` , `cache.r3.4xlarge` , `cache.r3.8xlarge`
	//
	// For region availability, see [Supported Node Types by Amazon Region](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion)
	CacheNodeType() *string
	SetCacheNodeType(val *string)
	// The name of the parameter group to associate with this replication group.
	//
	// If this argument is omitted, the default cache parameter group for the specified engine is used.
	//
	// If you are running Redis version 3.2.4 or later, only one node group (shard), and want to use a default parameter group, we recommend that you specify the parameter group by name.
	//
	// - To create a Redis (cluster mode disabled) replication group, use `CacheParameterGroupName=default.redis3.2` .
	// - To create a Redis (cluster mode enabled) replication group, use `CacheParameterGroupName=default.redis3.2.cluster.on` .
	CacheParameterGroupName() *string
	SetCacheParameterGroupName(val *string)
	// A list of cache security group names to associate with this replication group.
	CacheSecurityGroupNames() *[]*string
	SetCacheSecurityGroupNames(val *[]*string)
	// The name of the cache subnet group to be used for the replication group.
	//
	// > If you're going to launch your cluster in an Amazon VPC, you need to create a subnet group before you start creating a cluster. For more information, see [AWS::ElastiCache::SubnetGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html) .
	CacheSubnetGroupName() *string
	SetCacheSubnetGroupName(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Enables data tiering.
	//
	// Data tiering is only supported for replication groups using the r6gd node type. This parameter must be set to true when using r6gd nodes. For more information, see [Data tiering](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/data-tiering.html) .
	DataTieringEnabled() interface{}
	SetDataTieringEnabled(val interface{})
	// The name of the cache engine to be used for the clusters in this replication group.
	//
	// The value must be set to `Redis` .
	Engine() *string
	SetEngine(val *string)
	// The version number of the cache engine to be used for the clusters in this replication group.
	//
	// To view the supported cache engine versions, use the `DescribeCacheEngineVersions` operation.
	//
	// *Important:* You can upgrade to a newer engine version (see [Selecting a Cache Engine and Version](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement) ) in the *ElastiCache User Guide* , but you cannot downgrade to an earlier engine version. If you want to use an earlier engine version, you must delete the existing cluster or replication group and create it anew with the earlier engine version.
	EngineVersion() *string
	SetEngineVersion(val *string)
	// The name of the Global datastore.
	GlobalReplicationGroupId() *string
	SetGlobalReplicationGroupId(val *string)
	// The network type you choose when creating a replication group, either `ipv4` | `ipv6` .
	//
	// IPv6 is supported for workloads using Redis engine version 6.2 onward or Memcached engine version 1.6.6 on all instances built on the [Nitro system](https://docs.aws.amazon.com/ec2/nitro/) .
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
	//
	// For more information, see [Minimizing Downtime: Multi-AZ](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/AutoFailover.html) .
	MultiAzEnabled() interface{}
	SetMultiAzEnabled(val interface{})
	// Must be either `ipv4` | `ipv6` | `dual_stack` .
	//
	// IPv6 is supported for workloads using Redis engine version 6.2 onward or Memcached engine version 1.6.6 on all instances built on the [Nitro system](https://docs.aws.amazon.com/ec2/nitro/) .
	NetworkType() *string
	SetNetworkType(val *string)
	// The tree node.
	Node() constructs.Node
	// `NodeGroupConfiguration` is a property of the `AWS::ElastiCache::ReplicationGroup` resource that configures an Amazon ElastiCache (ElastiCache) Redis cluster node group.
	//
	// If you set [UseOnlineResharding](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-useonlineresharding) to `true` , you can update `NodeGroupConfiguration` without interruption. When `UseOnlineResharding` is set to `false` , or is not specified, updating `NodeGroupConfiguration` results in [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	NodeGroupConfiguration() interface{}
	SetNodeGroupConfiguration(val interface{})
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
	//
	// > The Amazon SNS topic owner must be the same as the cluster owner.
	NotificationTopicArn() *string
	SetNotificationTopicArn(val *string)
	// The number of clusters this replication group initially has.
	//
	// This parameter is not used if there is more than one node group (shard). You should use `ReplicasPerNodeGroup` instead.
	//
	// If `AutomaticFailoverEnabled` is `true` , the value of this parameter must be at least 2. If `AutomaticFailoverEnabled` is `false` you can omit this parameter (it will default to 1), or you can explicitly set it to a value between 2 and 6.
	//
	// The maximum permitted value for `NumCacheClusters` is 6 (1 primary plus 5 replicas).
	NumCacheClusters() *float64
	SetNumCacheClusters(val *float64)
	// An optional parameter that specifies the number of node groups (shards) for this Redis (cluster mode enabled) replication group.
	//
	// For Redis (cluster mode disabled) either omit this parameter or set it to 1.
	//
	// If you set [UseOnlineResharding](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-useonlineresharding) to `true` , you can update `NumNodeGroups` without interruption. When `UseOnlineResharding` is set to `false` , or is not specified, updating `NumNodeGroups` results in [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	//
	// Default: 1.
	NumNodeGroups() *float64
	SetNumNodeGroups(val *float64)
	// The port number on which each member of the replication group accepts connections.
	Port() *float64
	SetPort(val *float64)
	// A list of EC2 Availability Zones in which the replication group's clusters are created.
	//
	// The order of the Availability Zones in the list is the order in which clusters are allocated. The primary cluster is created in the first AZ in the list.
	//
	// This parameter is not used if there is more than one node group (shard). You should use `NodeGroupConfiguration` instead.
	//
	// > If you are creating your replication group in an Amazon VPC (recommended), you can only locate clusters in Availability Zones associated with the subnets in the selected subnet group.
	// >
	// > The number of Availability Zones listed must equal the value of `NumCacheClusters` .
	//
	// Default: system chosen Availability Zones.
	PreferredCacheClusterAZs() *[]*string
	SetPreferredCacheClusterAZs(val *[]*string)
	// Specifies the weekly time range during which maintenance on the cluster is performed.
	//
	// It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.
	//
	// Valid values for `ddd` are:
	//
	// - `sun`
	// - `mon`
	// - `tue`
	// - `wed`
	// - `thu`
	// - `fri`
	// - `sat`
	//
	// Example: `sun:23:00-mon:01:30`.
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	// The identifier of the cluster that serves as the primary for this replication group.
	//
	// This cluster must already exist and have a status of `available` .
	//
	// This parameter is not required if `NumCacheClusters` , `NumNodeGroups` , or `ReplicasPerNodeGroup` is specified.
	PrimaryClusterId() *string
	SetPrimaryClusterId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// An optional parameter that specifies the number of replica nodes in each node group (shard).
	//
	// Valid values are 0 to 5.
	ReplicasPerNodeGroup() *float64
	SetReplicasPerNodeGroup(val *float64)
	// A user-created description for the replication group.
	ReplicationGroupDescription() *string
	SetReplicationGroupDescription(val *string)
	// The replication group identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - A name must contain from 1 to 40 alphanumeric characters or hyphens.
	// - The first character must be a letter.
	// - A name cannot end with a hyphen or contain two consecutive hyphens.
	ReplicationGroupId() *string
	SetReplicationGroupId(val *string)
	// One or more Amazon VPC security groups associated with this replication group.
	//
	// Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud (Amazon VPC).
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// A list of Amazon Resource Names (ARN) that uniquely identify the Redis RDB snapshot files stored in Amazon S3.
	//
	// The snapshot files are used to populate the new replication group. The Amazon S3 object name in the ARN cannot contain any commas. The new replication group will have the number of node groups (console: shards) specified by the parameter *NumNodeGroups* or the number of node groups configured by *NodeGroupConfiguration* regardless of the number of ARNs specified here.
	//
	// Example of an Amazon S3 ARN: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns() *[]*string
	SetSnapshotArns(val *[]*string)
	// The name of a snapshot from which to restore data into the new replication group.
	//
	// The snapshot status changes to `restoring` while the new replication group is being created.
	SnapshotName() *string
	SetSnapshotName(val *string)
	// The number of days for which ElastiCache retains automatic snapshots before deleting them.
	//
	// For example, if you set `SnapshotRetentionLimit` to 5, a snapshot that was taken today is retained for 5 days before being deleted.
	//
	// Default: 0 (i.e., automatic backups are disabled for this cluster).
	SnapshotRetentionLimit() *float64
	SetSnapshotRetentionLimit(val *float64)
	// The cluster ID that is used as the daily snapshot source for the replication group.
	//
	// This parameter cannot be set for Redis (cluster mode enabled) replication groups.
	SnapshottingClusterId() *string
	SetSnapshottingClusterId(val *string)
	// The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).
	//
	// Example: `05:00-09:00`
	//
	// If you do not specify this parameter, ElastiCache automatically chooses an appropriate time range.
	SnapshotWindow() *string
	SetSnapshotWindow(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of tags to be added to this resource.
	//
	// Tags are comma-separated key,value pairs (e.g. Key= `myKey` , Value= `myKeyValue` . You can include multiple tags as shown following: Key= `myKey` , Value= `myKeyValue` Key= `mySecondKey` , Value= `mySecondKeyValue` . Tags on replication groups will be replicated to all nodes.
	Tags() awscdk.TagManager
	// A flag that enables in-transit encryption when set to `true` .
	//
	// You cannot modify the value of `TransitEncryptionEnabled` after the cluster is created. To enable in-transit encryption on a cluster you must set `TransitEncryptionEnabled` to `true` when you create a cluster.
	//
	// This parameter is valid only if the `Engine` parameter is `redis` , the `EngineVersion` parameter is `3.2.6` or `4.x` onward, and the cluster is being created in an Amazon VPC.
	//
	// If you enable in-transit encryption, you must also specify a value for `CacheSubnetGroup` .
	//
	// *Required:* Only available when creating a replication group in an Amazon VPC using redis version `3.2.6` or `4.x` onward.
	//
	// Default: `false`
	//
	// > For HIPAA compliance, you must specify `TransitEncryptionEnabled` as `true` , an `AuthToken` , and a `CacheSubnetGroup` .
	TransitEncryptionEnabled() interface{}
	SetTransitEncryptionEnabled(val interface{})
	// A setting that allows you to migrate your clients to use in-transit encryption, with no downtime.
	//
	// When setting `TransitEncryptionEnabled` to `true` , you can set your `TransitEncryptionMode` to `preferred` in the same request, to allow both encrypted and unencrypted connections at the same time. Once you migrate all your Redis clients to use encrypted connections you can modify the value to `required` to allow encrypted connections only.
	//
	// Setting `TransitEncryptionMode` to `required` is a two-step process that requires you to first set the `TransitEncryptionMode` to `preferred` , after that you can set `TransitEncryptionMode` to `required` .
	//
	// This process will not trigger the replacement of the replication group.
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
	internal.Type__awscdkIInspectable
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


// Create a new `AWS::ElastiCache::ReplicationGroup`.
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

// Create a new `AWS::ElastiCache::ReplicationGroup`.
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

// Check whether the given construct is a CfnResource.
func CfnReplicationGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationGroup_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticache.CfnReplicationGroup",
		"isCfnResource",
		[]interface{}{construct},
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

