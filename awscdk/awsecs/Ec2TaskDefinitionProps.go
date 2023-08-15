package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The properties for a task definition run on an EC2 cluster.
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
type Ec2TaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	// Default: - An execution role will be automatically created if you use ECR images in your task definition.
	//
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Default: - Automatically generated name.
	//
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The configuration details for the App Mesh proxy.
	// Default: - No proxy configuration.
	//
	ProxyConfiguration ProxyConfiguration `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Default: - A task role is automatically created for you.
	//
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	// Default: - No volumes are passed to the Docker daemon on a container instance.
	//
	Volumes *[]*Volume `field:"optional" json:"volumes" yaml:"volumes"`
	// The inference accelerators to use for the containers in the task.
	//
	// Not supported in Fargate.
	// Default: - No inference accelerators.
	//
	InferenceAccelerators *[]*InferenceAccelerator `field:"optional" json:"inferenceAccelerators" yaml:"inferenceAccelerators"`
	// The IPC resource namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	// Default: - IpcMode used by the task is not specified.
	//
	IpcMode IpcMode `field:"optional" json:"ipcMode" yaml:"ipcMode"`
	// The Docker networking mode to use for the containers in the task.
	//
	// The valid values are NONE, BRIDGE, AWS_VPC, and HOST.
	// Default: - NetworkMode.BRIDGE for EC2 tasks, AWS_VPC for Fargate tasks.
	//
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The process namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	// Default: - PidMode used by the task is not specified.
	//
	PidMode PidMode `field:"optional" json:"pidMode" yaml:"pidMode"`
	// An array of placement constraint objects to use for the task.
	//
	// You can
	// specify a maximum of 10 constraints per task (this limit includes
	// constraints in the task definition and those specified at run time).
	// Default: - No placement constraints.
	//
	PlacementConstraints *[]PlacementConstraint `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
}

