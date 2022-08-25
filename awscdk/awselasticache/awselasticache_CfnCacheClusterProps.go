package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCacheCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCacheClusterProps := &cfnCacheClusterProps{
//   	cacheNodeType: jsii.String("cacheNodeType"),
//   	engine: jsii.String("engine"),
//   	numCacheNodes: jsii.Number(123),
//
//   	// the properties below are optional
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	azMode: jsii.String("azMode"),
//   	cacheParameterGroupName: jsii.String("cacheParameterGroupName"),
//   	cacheSecurityGroupNames: []*string{
//   		jsii.String("cacheSecurityGroupNames"),
//   	},
//   	cacheSubnetGroupName: jsii.String("cacheSubnetGroupName"),
//   	clusterName: jsii.String("clusterName"),
//   	engineVersion: jsii.String("engineVersion"),
//   	logDeliveryConfigurations: []interface{}{
//   		&logDeliveryConfigurationRequestProperty{
//   			destinationDetails: &destinationDetailsProperty{
//   				cloudWatchLogsDetails: &cloudWatchLogsDestinationDetailsProperty{
//   					logGroup: jsii.String("logGroup"),
//   				},
//   				kinesisFirehoseDetails: &kinesisFirehoseDestinationDetailsProperty{
//   					deliveryStream: jsii.String("deliveryStream"),
//   				},
//   			},
//   			destinationType: jsii.String("destinationType"),
//   			logFormat: jsii.String("logFormat"),
//   			logType: jsii.String("logType"),
//   		},
//   	},
//   	notificationTopicArn: jsii.String("notificationTopicArn"),
//   	port: jsii.Number(123),
//   	preferredAvailabilityZone: jsii.String("preferredAvailabilityZone"),
//   	preferredAvailabilityZones: []*string{
//   		jsii.String("preferredAvailabilityZones"),
//   	},
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	snapshotArns: []*string{
//   		jsii.String("snapshotArns"),
//   	},
//   	snapshotName: jsii.String("snapshotName"),
//   	snapshotRetentionLimit: jsii.Number(123),
//   	snapshotWindow: jsii.String("snapshotWindow"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
type CfnCacheClusterProps struct {
	// The compute and memory capacity of the nodes in the node group (shard).
	//
	// The following node types are supported by ElastiCache. Generally speaking, the current generation types provide more memory and computational power at lower cost when compared to their equivalent previous generation counterparts. Changing the CacheNodeType of a Memcached instance is currently not supported. If you need to scale using Memcached, we recommend forcing a replacement update by changing the `LogicalResourceId` of the resource.
	//
	// - General purpose:
	//
	// - Current generation:
	//
	// *M6g node types:* `cache.m6g.large` , `cache.m6g.xlarge` , `cache.m6g.2xlarge` , `cache.m6g.4xlarge` , `cache.m6g.8xlarge` , `cache.m6g.12xlarge` , `cache.m6g.16xlarge` , `cache.m6g.24xlarge`
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
	// *R6g node types:* `cache.r6g.large` , `cache.r6g.xlarge` , `cache.r6g.2xlarge` , `cache.r6g.4xlarge` , `cache.r6g.8xlarge` , `cache.r6g.12xlarge` , `cache.r6g.16xlarge` , `cache.r6g.24xlarge`
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
	//
	// *Additional node type info*
	//
	// - All current generation instance types are created in Amazon VPC by default.
	// - Redis append-only files (AOF) are not supported for T1 or T2 instances.
	// - Redis Multi-AZ with automatic failover is not supported on T1 instances.
	// - Redis configuration variables `appendonly` and `appendfsync` are not supported on Redis version 2.8.22 and later.
	CacheNodeType *string `field:"required" json:"cacheNodeType" yaml:"cacheNodeType"`
	// The name of the cache engine to be used for this cluster.
	//
	// Valid values for this parameter are: `memcached` | `redis`.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The number of cache nodes that the cache cluster should have.
	//
	// > However, if the `PreferredAvailabilityZone` and `PreferredAvailabilityZones` properties were not previously specified and you don't specify any new values, an update requires [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	NumCacheNodes *float64 `field:"required" json:"numCacheNodes" yaml:"numCacheNodes"`
	// If you are running Redis engine version 6.0 or later, set this parameter to yes if you want to opt-in to the next minor version upgrade campaign. This parameter is disabled for previous versions.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Specifies whether the nodes in this Memcached cluster are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region.
	//
	// This parameter is only supported for Memcached clusters.
	//
	// If the `AZMode` and `PreferredAvailabilityZones` are not specified, ElastiCache assumes `single-az` mode.
	AzMode *string `field:"optional" json:"azMode" yaml:"azMode"`
	// The name of the parameter group to associate with this cluster.
	//
	// If this argument is omitted, the default parameter group for the specified engine is used. You cannot use any parameter group which has `cluster-enabled='yes'` when creating a cluster.
	CacheParameterGroupName *string `field:"optional" json:"cacheParameterGroupName" yaml:"cacheParameterGroupName"`
	// A list of security group names to associate with this cluster.
	//
	// Use this parameter only when you are creating a cluster outside of an Amazon Virtual Private Cloud (Amazon VPC).
	CacheSecurityGroupNames *[]*string `field:"optional" json:"cacheSecurityGroupNames" yaml:"cacheSecurityGroupNames"`
	// The name of the subnet group to be used for the cluster.
	//
	// Use this parameter only when you are creating a cluster in an Amazon Virtual Private Cloud (Amazon VPC).
	//
	// > If you're going to launch your cluster in an Amazon VPC, you need to create a subnet group before you start creating a cluster. For more information, see [AWS::ElastiCache::SubnetGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html) .
	CacheSubnetGroupName *string `field:"optional" json:"cacheSubnetGroupName" yaml:"cacheSubnetGroupName"`
	// A name for the cache cluster.
	//
	// If you don't specify a name, AWSCloudFormation generates a unique physical ID and uses that ID for the cache cluster. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// The name must contain 1 to 50 alphanumeric characters or hyphens. The name must start with a letter and cannot end with a hyphen or contain two consecutive hyphens.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The version number of the cache engine to be used for this cluster.
	//
	// To view the supported cache engine versions, use the DescribeCacheEngineVersions operation.
	//
	// *Important:* You can upgrade to a newer engine version (see [Selecting a Cache Engine and Version](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement) ), but you cannot downgrade to an earlier engine version. If you want to use an earlier engine version, you must delete the existing cluster or replication group and create it anew with the earlier engine version.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Specifies the destination, format and type of the logs.
	LogDeliveryConfigurations interface{} `field:"optional" json:"logDeliveryConfigurations" yaml:"logDeliveryConfigurations"`
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
	//
	// > The Amazon SNS topic owner must be the same as the cluster owner.
	NotificationTopicArn *string `field:"optional" json:"notificationTopicArn" yaml:"notificationTopicArn"`
	// The port number on which each of the cache nodes accepts connections.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The EC2 Availability Zone in which the cluster is created.
	//
	// All nodes belonging to this cluster are placed in the preferred Availability Zone. If you want to create your nodes across multiple Availability Zones, use `PreferredAvailabilityZones` .
	//
	// Default: System chosen Availability Zone.
	PreferredAvailabilityZone *string `field:"optional" json:"preferredAvailabilityZone" yaml:"preferredAvailabilityZone"`
	// A list of the Availability Zones in which cache nodes are created.
	//
	// The order of the zones in the list is not important.
	//
	// This option is only supported on Memcached.
	//
	// > If you are creating your cluster in an Amazon VPC (recommended) you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group.
	// >
	// > The number of Availability Zones listed must equal the value of `NumCacheNodes` .
	//
	// If you want all the nodes in the same Availability Zone, use `PreferredAvailabilityZone` instead, or repeat the Availability Zone multiple times in the list.
	//
	// Default: System chosen Availability Zones.
	PreferredAvailabilityZones *[]*string `field:"optional" json:"preferredAvailabilityZones" yaml:"preferredAvailabilityZones"`
	// Specifies the weekly time range during which maintenance on the cluster is performed.
	//
	// It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period. Valid values for `ddd` are:
	//
	// Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.
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
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// A single-element string list containing an Amazon Resource Name (ARN) that uniquely identifies a Redis RDB snapshot file stored in Amazon S3.
	//
	// The snapshot file is used to populate the node group (shard). The Amazon S3 object name in the ARN cannot contain any commas.
	//
	// > This parameter is only valid if the `Engine` parameter is `redis` .
	//
	// Example of an Amazon S3 ARN: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns *[]*string `field:"optional" json:"snapshotArns" yaml:"snapshotArns"`
	// The name of a Redis snapshot from which to restore data into the new node group (shard).
	//
	// The snapshot status changes to `restoring` while the new node group (shard) is being created.
	//
	// > This parameter is only valid if the `Engine` parameter is `redis` .
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// The number of days for which ElastiCache retains automatic snapshots before deleting them.
	//
	// For example, if you set `SnapshotRetentionLimit` to 5, a snapshot taken today is retained for 5 days before being deleted.
	//
	// > This parameter is only valid if the `Engine` parameter is `redis` .
	//
	// Default: 0 (i.e., automatic backups are disabled for this cache cluster).
	SnapshotRetentionLimit *float64 `field:"optional" json:"snapshotRetentionLimit" yaml:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).
	//
	// Example: `05:00-09:00`
	//
	// If you do not specify this parameter, ElastiCache automatically chooses an appropriate time range.
	//
	// > This parameter is only valid if the `Engine` parameter is `redis` .
	SnapshotWindow *string `field:"optional" json:"snapshotWindow" yaml:"snapshotWindow"`
	// A list of tags to be added to this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// One or more VPC security groups associated with the cluster.
	//
	// Use this parameter only when you are creating a cluster in an Amazon Virtual Private Cloud (Amazon VPC).
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

