package awselasticache


// A list of the replication groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   regionalConfigurationProperty := &regionalConfigurationProperty{
//   	replicationGroupId: jsii.String("replicationGroupId"),
//   	replicationGroupRegion: jsii.String("replicationGroupRegion"),
//   	reshardingConfigurations: []interface{}{
//   		&reshardingConfigurationProperty{
//   			nodeGroupId: jsii.String("nodeGroupId"),
//   			preferredAvailabilityZones: []*string{
//   				jsii.String("preferredAvailabilityZones"),
//   			},
//   		},
//   	},
//   }
//
type CfnGlobalReplicationGroup_RegionalConfigurationProperty struct {
	// The name of the secondary cluster.
	ReplicationGroupId *string `field:"optional" json:"replicationGroupId" yaml:"replicationGroupId"`
	// The Amazon region where the cluster is stored.
	ReplicationGroupRegion *string `field:"optional" json:"replicationGroupRegion" yaml:"replicationGroupRegion"`
	// A list of PreferredAvailabilityZones objects that specifies the configuration of a node group in the resharded cluster.
	ReshardingConfigurations interface{} `field:"optional" json:"reshardingConfigurations" yaml:"reshardingConfigurations"`
}

