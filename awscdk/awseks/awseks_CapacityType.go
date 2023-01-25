package awseks


// Capacity type of the managed node group.
//
// Example:
//   var cluster cluster
//
//   cluster.addNodegroupCapacity(jsii.String("extra-ng-spot"), &nodegroupOptions{
//   	instanceTypes: []instanceType{
//   		ec2.NewInstanceType(jsii.String("c5.large")),
//   		ec2.NewInstanceType(jsii.String("c5a.large")),
//   		ec2.NewInstanceType(jsii.String("c5d.large")),
//   	},
//   	minSize: jsii.Number(3),
//   	capacityType: eks.capacityType_SPOT,
//   })
//
// Experimental.
type CapacityType string

const (
	// spot instances.
	// Experimental.
	CapacityType_SPOT CapacityType = "SPOT"
	// on-demand instances.
	// Experimental.
	CapacityType_ON_DEMAND CapacityType = "ON_DEMAND"
)

