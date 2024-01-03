package awsecs


// Amazon ECS variant.
//
// Example:
//   var cluster cluster
//
//
//   cluster.AddCapacity(jsii.String("bottlerocket-asg"), &AddCapacityOptions{
//   	InstanceType: ec2.NewInstanceType(jsii.String("p3.2xlarge")),
//   	MachineImage: ecs.NewBottleRocketImage(&BottleRocketImageProps{
//   		Variant: ecs.BottlerocketEcsVariant_AWS_ECS_2_NVIDIA,
//   	}),
//   })
//
type BottlerocketEcsVariant string

const (
	// aws-ecs-1 variant.
	BottlerocketEcsVariant_AWS_ECS_1 BottlerocketEcsVariant = "AWS_ECS_1"
	// aws-ecs-1-nvidia variant.
	BottlerocketEcsVariant_AWS_ECS_1_NVIDIA BottlerocketEcsVariant = "AWS_ECS_1_NVIDIA"
	// aws-ecs-2 variant.
	BottlerocketEcsVariant_AWS_ECS_2 BottlerocketEcsVariant = "AWS_ECS_2"
	// aws-ecs-2-nvidia variant.
	BottlerocketEcsVariant_AWS_ECS_2_NVIDIA BottlerocketEcsVariant = "AWS_ECS_2_NVIDIA"
)

