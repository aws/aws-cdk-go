package previewawselasticachemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawselasticache/previewawselasticachemixins/internal"
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
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnReplicationGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnReplicationGroupPropsMixin(&CfnReplicationGroupMixinProps{
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
//   	ReplicationGroupDescription: jsii.String("replicationGroupDescription"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitEncryptionEnabled: jsii.Boolean(false),
//   	TransitEncryptionMode: jsii.String("transitEncryptionMode"),
//   	UserGroupIds: []*string{
//   		jsii.String("userGroupIds"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html
//
type CfnReplicationGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnReplicationGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnReplicationGroupPropsMixin
type jsiiProxy_CfnReplicationGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnReplicationGroupPropsMixin) Props() *CfnReplicationGroupMixinProps {
	var returns *CfnReplicationGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElastiCache::ReplicationGroup`.
func NewCfnReplicationGroupPropsMixin(props *CfnReplicationGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnReplicationGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnReplicationGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReplicationGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnReplicationGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElastiCache::ReplicationGroup`.
func NewCfnReplicationGroupPropsMixin_Override(c CfnReplicationGroupPropsMixin, props *CfnReplicationGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnReplicationGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnReplicationGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReplicationGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnReplicationGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicationGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnReplicationGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReplicationGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

