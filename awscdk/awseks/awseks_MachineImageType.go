package awseks


// The machine image type.
//
// Example:
//   var cluster cluster
//
//   cluster.addAutoScalingGroupCapacity(jsii.String("BottlerocketNodes"), &autoScalingGroupCapacityOptions{
//   	instanceType: ec2.NewInstanceType(jsii.String("t3.small")),
//   	minCapacity: jsii.Number(2),
//   	machineImageType: eks.machineImageType_BOTTLEROCKET,
//   })
//
type MachineImageType string

const (
	// Amazon EKS-optimized Linux AMI.
	MachineImageType_AMAZON_LINUX_2 MachineImageType = "AMAZON_LINUX_2"
	// Bottlerocket AMI.
	MachineImageType_BOTTLEROCKET MachineImageType = "BOTTLEROCKET"
)

