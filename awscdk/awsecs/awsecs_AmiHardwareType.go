package awsecs


// The ECS-optimized AMI variant to use.
//
// For more information, see
// [Amazon ECS-optimized AMIs](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html).
//
// Example:
//   var cluster cluster
//
//
//   cluster.AddCapacity(jsii.String("graviton-cluster"), &AddCapacityOptions{
//   	MinCapacity: jsii.Number(2),
//   	InstanceType: ec2.NewInstanceType(jsii.String("c6g.large")),
//   	MachineImage: ecs.EcsOptimizedImage_AmazonLinux2(ecs.AmiHardwareType_ARM),
//   })
//
// Experimental.
type AmiHardwareType string

const (
	// Use the standard Amazon ECS-optimized AMI.
	// Experimental.
	AmiHardwareType_STANDARD AmiHardwareType = "STANDARD"
	// Use the Amazon ECS GPU-optimized AMI.
	// Experimental.
	AmiHardwareType_GPU AmiHardwareType = "GPU"
	// Use the Amazon ECS-optimized Amazon Linux 2 (arm64) AMI.
	// Experimental.
	AmiHardwareType_ARM AmiHardwareType = "ARM"
)

