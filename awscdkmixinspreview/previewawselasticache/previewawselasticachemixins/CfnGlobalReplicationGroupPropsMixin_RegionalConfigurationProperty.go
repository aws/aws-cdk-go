package previewawselasticachemixins


// A list of the replication groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   regionalConfigurationProperty := &RegionalConfigurationProperty{
//   	ReplicationGroupId: jsii.String("replicationGroupId"),
//   	ReplicationGroupRegion: jsii.String("replicationGroupRegion"),
//   	ReshardingConfigurations: []interface{}{
//   		&ReshardingConfigurationProperty{
//   			NodeGroupId: jsii.String("nodeGroupId"),
//   			PreferredAvailabilityZones: []*string{
//   				jsii.String("preferredAvailabilityZones"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-globalreplicationgroup-regionalconfiguration.html
//
type CfnGlobalReplicationGroupPropsMixin_RegionalConfigurationProperty struct {
	// The name of the secondary cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-globalreplicationgroup-regionalconfiguration.html#cfn-elasticache-globalreplicationgroup-regionalconfiguration-replicationgroupid
	//
	ReplicationGroupId *string `field:"optional" json:"replicationGroupId" yaml:"replicationGroupId"`
	// The Amazon region where the cluster is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-globalreplicationgroup-regionalconfiguration.html#cfn-elasticache-globalreplicationgroup-regionalconfiguration-replicationgroupregion
	//
	ReplicationGroupRegion *string `field:"optional" json:"replicationGroupRegion" yaml:"replicationGroupRegion"`
	// A list of PreferredAvailabilityZones objects that specifies the configuration of a node group in the resharded cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-globalreplicationgroup-regionalconfiguration.html#cfn-elasticache-globalreplicationgroup-regionalconfiguration-reshardingconfigurations
	//
	ReshardingConfigurations interface{} `field:"optional" json:"reshardingConfigurations" yaml:"reshardingConfigurations"`
}

