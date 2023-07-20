package awselasticache


// A list of `PreferredAvailabilityZones` objects that specifies the configuration of a node group in the resharded cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reshardingConfigurationProperty := &ReshardingConfigurationProperty{
//   	NodeGroupId: jsii.String("nodeGroupId"),
//   	PreferredAvailabilityZones: []*string{
//   		jsii.String("preferredAvailabilityZones"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-globalreplicationgroup-reshardingconfiguration.html
//
type CfnGlobalReplicationGroup_ReshardingConfigurationProperty struct {
	// Either the ElastiCache for Redis supplied 4-digit id or a user supplied id for the node group these configuration values apply to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-globalreplicationgroup-reshardingconfiguration.html#cfn-elasticache-globalreplicationgroup-reshardingconfiguration-nodegroupid
	//
	NodeGroupId *string `field:"optional" json:"nodeGroupId" yaml:"nodeGroupId"`
	// A list of preferred availability zones for the nodes in this cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-globalreplicationgroup-reshardingconfiguration.html#cfn-elasticache-globalreplicationgroup-reshardingconfiguration-preferredavailabilityzones
	//
	PreferredAvailabilityZones *[]*string `field:"optional" json:"preferredAvailabilityZones" yaml:"preferredAvailabilityZones"`
}

