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
//   cluster.addCapacity(jsii.String("graviton-cluster"), &addCapacityOptions{
//   	minCapacity: jsii.Number(2),
//   	instanceType: ec2.NewInstanceType(jsii.String("c6g.large")),
//   	machineImage: ecs.ecsOptimizedImage.amazonLinux2(ecs.amiHardwareType_ARM),
//   })
//
type AmiHardwareType string

const (
	// Use the standard Amazon ECS-optimized AMI.
	AmiHardwareType_STANDARD AmiHardwareType = "STANDARD"
	// Use the Amazon ECS GPU-optimized AMI.
	AmiHardwareType_GPU AmiHardwareType = "GPU"
	// Use the Amazon ECS-optimized Amazon Linux 2 (arm64) AMI.
	AmiHardwareType_ARM AmiHardwareType = "ARM"
)

