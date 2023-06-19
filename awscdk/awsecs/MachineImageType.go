package awsecs


// The machine image type.
//
// Example:
//   var cluster cluster
//
//
//   cluster.AddCapacity(jsii.String("graviton-cluster"), &AddCapacityOptions{
//   	MinCapacity: jsii.Number(2),
//   	InstanceType: ec2.NewInstanceType(jsii.String("c6g.large")),
//   	MachineImageType: ecs.MachineImageType_BOTTLEROCKET,
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

