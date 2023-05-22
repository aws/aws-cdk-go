package awseks


// The machine image type.
//
// Example:
//   var cluster cluster
//
//   cluster.AddAutoScalingGroupCapacity(jsii.String("BottlerocketNodes"), &AutoScalingGroupCapacityOptions{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t3.small")),
//   	MinCapacity: jsii.Number(2),
//   	MachineImageType: eks.MachineImageType_BOTTLEROCKET,
//   })
//
// Experimental.
type MachineImageType string

const (
	// Amazon EKS-optimized Linux AMI.
	// Experimental.
	MachineImageType_AMAZON_LINUX_2 MachineImageType = "AMAZON_LINUX_2"
	// Bottlerocket AMI.
	// Experimental.
	MachineImageType_BOTTLEROCKET MachineImageType = "BOTTLEROCKET"
)

