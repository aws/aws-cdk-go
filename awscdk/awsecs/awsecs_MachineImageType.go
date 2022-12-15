package awsecs


// The machine image type.
//
// Example:
//   var cluster cluster
//
//
//   cluster.addCapacity(jsii.String("graviton-cluster"), &addCapacityOptions{
//   	minCapacity: jsii.Number(2),
//   	instanceType: ec2.NewInstanceType(jsii.String("c6g.large")),
//   	machineImageType: ecs.machineImageType_BOTTLEROCKET,
//   })
//
// Experimental.
type MachineImageType string

const (
	// Amazon ECS-optimized Amazon Linux 2 AMI.
	// Experimental.
	MachineImageType_AMAZON_LINUX_2 MachineImageType = "AMAZON_LINUX_2"
	// Bottlerocket AMI.
	// Experimental.
	MachineImageType_BOTTLEROCKET MachineImageType = "BOTTLEROCKET"
)

