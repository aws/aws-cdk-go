package awselasticache


// `NodeGroupConfiguration` is a property of the `AWS::ElastiCache::ReplicationGroup` resource that configures an Amazon ElastiCache (ElastiCache) Redis cluster node group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeGroupConfigurationProperty := &nodeGroupConfigurationProperty{
//   	nodeGroupId: jsii.String("nodeGroupId"),
//   	primaryAvailabilityZone: jsii.String("primaryAvailabilityZone"),
//   	replicaAvailabilityZones: []*string{
//   		jsii.String("replicaAvailabilityZones"),
//   	},
//   	replicaCount: jsii.Number(123),
//   	slots: jsii.String("slots"),
//   }
//
type CfnReplicationGroup_NodeGroupConfigurationProperty struct {
	// Either the ElastiCache for Redis supplied 4-digit id or a user supplied id for the node group these configuration values apply to.
	NodeGroupId *string `field:"optional" json:"nodeGroupId" yaml:"nodeGroupId"`
	// The Availability Zone where the primary node of this node group (shard) is launched.
	PrimaryAvailabilityZone *string `field:"optional" json:"primaryAvailabilityZone" yaml:"primaryAvailabilityZone"`
	// A list of Availability Zones to be used for the read replicas.
	//
	// The number of Availability Zones in this list must match the value of `ReplicaCount` or `ReplicasPerNodeGroup` if not specified.
	ReplicaAvailabilityZones *[]*string `field:"optional" json:"replicaAvailabilityZones" yaml:"replicaAvailabilityZones"`
	// The number of read replica nodes in this node group (shard).
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// A string of comma-separated values where the first set of values are the slot numbers (zero based), and the second set of values are the keyspaces for each slot.
	//
	// The following example specifies three slots (numbered 0, 1, and 2): `0,1,2,0-4999,5000-9999,10000-16,383` .
	//
	// If you don't specify a value, ElastiCache allocates keys equally among each slot.
	//
	// When you use an `UseOnlineResharding` update policy to update the number of node groups without interruption, ElastiCache evenly distributes the keyspaces between the specified number of slots. This cannot be updated later. Therefore, after updating the number of node groups in this way, you should remove the value specified for the `Slots` property of each `NodeGroupConfiguration` from the stack template, as it no longer reflects the actual values in each node group. For more information, see [UseOnlineResharding Policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-useonlineresharding) .
	Slots *string `field:"optional" json:"slots" yaml:"slots"`
}

