package previewawselasticachemixins


// Properties for CfnGlobalReplicationGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGlobalReplicationGroupMixinProps := &CfnGlobalReplicationGroupMixinProps{
//   	AutomaticFailoverEnabled: jsii.Boolean(false),
//   	CacheNodeType: jsii.String("cacheNodeType"),
//   	CacheParameterGroupName: jsii.String("cacheParameterGroupName"),
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	GlobalNodeGroupCount: jsii.Number(123),
//   	GlobalReplicationGroupDescription: jsii.String("globalReplicationGroupDescription"),
//   	GlobalReplicationGroupIdSuffix: jsii.String("globalReplicationGroupIdSuffix"),
//   	Members: []interface{}{
//   		&GlobalReplicationGroupMemberProperty{
//   			ReplicationGroupId: jsii.String("replicationGroupId"),
//   			ReplicationGroupRegion: jsii.String("replicationGroupRegion"),
//   			Role: jsii.String("role"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html
//
type CfnGlobalReplicationGroupMixinProps struct {
	// Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails.
	//
	// `AutomaticFailoverEnabled` must be enabled for Valkey or Redis OSS (cluster mode enabled) replication groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html#cfn-elasticache-globalreplicationgroup-automaticfailoverenabled
	//
	AutomaticFailoverEnabled interface{} `field:"optional" json:"automaticFailoverEnabled" yaml:"automaticFailoverEnabled"`
	// The cache node type of the Global datastore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html#cfn-elasticache-globalreplicationgroup-cachenodetype
	//
	CacheNodeType *string `field:"optional" json:"cacheNodeType" yaml:"cacheNodeType"`
	// The name of the cache parameter group to use with the Global datastore.
	//
	// It must be compatible with the major engine version used by the Global datastore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html#cfn-elasticache-globalreplicationgroup-cacheparametergroupname
	//
	CacheParameterGroupName *string `field:"optional" json:"cacheParameterGroupName" yaml:"cacheParameterGroupName"`
	// The ElastiCache engine.
	//
	// For Valkey or Redis OSS only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html#cfn-elasticache-globalreplicationgroup-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The Elasticache Valkey or Redis OSS engine version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html#cfn-elasticache-globalreplicationgroup-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The number of node groups that comprise the Global Datastore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html#cfn-elasticache-globalreplicationgroup-globalnodegroupcount
	//
	GlobalNodeGroupCount *float64 `field:"optional" json:"globalNodeGroupCount" yaml:"globalNodeGroupCount"`
	// The optional description of the Global datastore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html#cfn-elasticache-globalreplicationgroup-globalreplicationgroupdescription
	//
	GlobalReplicationGroupDescription *string `field:"optional" json:"globalReplicationGroupDescription" yaml:"globalReplicationGroupDescription"`
	// The suffix name of a Global Datastore.
	//
	// The suffix guarantees uniqueness of the Global Datastore name across multiple regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html#cfn-elasticache-globalreplicationgroup-globalreplicationgroupidsuffix
	//
	GlobalReplicationGroupIdSuffix *string `field:"optional" json:"globalReplicationGroupIdSuffix" yaml:"globalReplicationGroupIdSuffix"`
	// The replication groups that comprise the Global datastore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html#cfn-elasticache-globalreplicationgroup-members
	//
	Members interface{} `field:"optional" json:"members" yaml:"members"`
	// The Regions that comprise the Global Datastore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html#cfn-elasticache-globalreplicationgroup-regionalconfigurations
	//
	RegionalConfigurations interface{} `field:"optional" json:"regionalConfigurations" yaml:"regionalConfigurations"`
}

