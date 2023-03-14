package awselasticache


// A list of `PreferredAvailabilityZones` objects that specifies the configuration of a node group in the resharded cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reshardingConfigurationProperty := &reshardingConfigurationProperty{
//   	nodeGroupId: jsii.String("nodeGroupId"),
//   	preferredAvailabilityZones: []*string{
//   		jsii.String("preferredAvailabilityZones"),
//   	},
//   }
//
type CfnGlobalReplicationGroup_ReshardingConfigurationProperty struct {
	// Either the ElastiCache for Redis supplied 4-digit id or a user supplied id for the node group these configuration values apply to.
	NodeGroupId *string `field:"optional" json:"nodeGroupId" yaml:"nodeGroupId"`
	// A list of preferred availability zones for the nodes in this cluster.
	PreferredAvailabilityZones *[]*string `field:"optional" json:"preferredAvailabilityZones" yaml:"preferredAvailabilityZones"`
}

