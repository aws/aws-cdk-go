package awselasticache


// Properties for defining a `CfnGlobalReplicationGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGlobalReplicationGroupProps := &CfnGlobalReplicationGroupProps{
//   	Members: []interface{}{
//   		&GlobalReplicationGroupMemberProperty{
//   			ReplicationGroupId: jsii.String("replicationGroupId"),
//   			ReplicationGroupRegion: jsii.String("replicationGroupRegion"),
//   			Role: jsii.String("role"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AutomaticFailoverEnabled: jsii.Boolean(false),
//   	CacheNodeType: jsii.String("cacheNodeType"),
//   	CacheParameterGroupName: jsii.String("cacheParameterGroupName"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	GlobalNodeGroupCount: jsii.Number(123),
//   	GlobalReplicationGroupDescription: jsii.String("globalReplicationGroupDescription"),
//   	GlobalReplicationGroupIdSuffix: jsii.String("globalReplicationGroupIdSuffix"),
//   	RegionalConfigurations: []interface{}{
//   		&RegionalConfigurationProperty{
//   			ReplicationGroupId: jsii.String("replicationGroupId"),
//   			ReplicationGroupRegion: jsii.String("replicationGroupRegion"),
//   			ReshardingConfigurations: []interface{}{
//   				&ReshardingConfigurationProperty{
//   					NodeGroupId: jsii.String("nodeGroupId"),
//   					PreferredAvailabilityZones: []*string{
//   						jsii.String("preferredAvailabilityZones"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnGlobalReplicationGroupProps struct {
	// The replication groups that comprise the Global datastore.
	Members interface{} `field:"required" json:"members" yaml:"members"`
	// Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails.
	//
	// `AutomaticFailoverEnabled` must be enabled for Redis (cluster mode enabled) replication groups.
	AutomaticFailoverEnabled interface{} `field:"optional" json:"automaticFailoverEnabled" yaml:"automaticFailoverEnabled"`
	// The cache node type of the Global datastore.
	CacheNodeType *string `field:"optional" json:"cacheNodeType" yaml:"cacheNodeType"`
	// The name of the cache parameter group to use with the Global datastore.
	//
	// It must be compatible with the major engine version used by the Global datastore.
	CacheParameterGroupName *string `field:"optional" json:"cacheParameterGroupName" yaml:"cacheParameterGroupName"`
	// The Elasticache Redis engine version.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The number of node groups that comprise the Global Datastore.
	GlobalNodeGroupCount *float64 `field:"optional" json:"globalNodeGroupCount" yaml:"globalNodeGroupCount"`
	// The optional description of the Global datastore.
	GlobalReplicationGroupDescription *string `field:"optional" json:"globalReplicationGroupDescription" yaml:"globalReplicationGroupDescription"`
	// The suffix name of a Global Datastore.
	//
	// The suffix guarantees uniqueness of the Global Datastore name across multiple regions.
	GlobalReplicationGroupIdSuffix *string `field:"optional" json:"globalReplicationGroupIdSuffix" yaml:"globalReplicationGroupIdSuffix"`
	// The Regions that comprise the Global Datastore.
	RegionalConfigurations interface{} `field:"optional" json:"regionalConfigurations" yaml:"regionalConfigurations"`
}

