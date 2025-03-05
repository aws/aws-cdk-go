package awsecs


// The networking mode to use for the containers in the task.
//
// Example:
//   ec2TaskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"), &Ec2TaskDefinitionProps{
//   	NetworkMode: ecs.NetworkMode_BRIDGE,
//   })
//
//   container := ec2TaskDefinition.AddContainer(jsii.String("WebContainer"), &ContainerDefinitionOptions{
//   	// Use an image from DockerHub
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryLimitMiB: jsii.Number(1024),
//   })
//
type NetworkMode string

const (
	// The task's containers do not have external connectivity and port mappings can't be specified in the container definition.
	NetworkMode_NONE NetworkMode = "NONE"
	// The task utilizes Docker's built-in virtual network which runs inside each container instance.
	NetworkMode_BRIDGE NetworkMode = "BRIDGE"
	// The task is allocated an elastic network interface.
	NetworkMode_AWS_VPC NetworkMode = "AWS_VPC"
	// The task bypasses Docker's built-in virtual network and maps container ports directly to the EC2 instance's network interface directly.
	//
	// In this mode, you can't run multiple instantiations of the same task on a
	// single container instance when port mappings are used.
	NetworkMode_HOST NetworkMode = "HOST"
	// The task utilizes NAT network mode required by Windows containers.
	//
	// This is the only supported network mode for Windows containers. For more information, see
	// [Task Definition Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#network_mode).
	NetworkMode_NAT NetworkMode = "NAT"
)

