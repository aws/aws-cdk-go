package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props to configure an EcsFargateContainerDefinition.
//
// Example:
//   jobDefn := batch.NewEcsJobDefinition(this, jsii.String("JobDefn"), &EcsJobDefinitionProps{
//   	Container: batch.NewEcsFargateContainerDefinition(this, jsii.String("myFargateContainer"), &EcsFargateContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/amazonlinux/amazonlinux:latest")),
//   		Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   		Cpu: jsii.Number(256),
//   		EphemeralStorageSize: cdk.Size_Gibibytes(jsii.Number(100)),
//   		FargateCpuArchitecture: ecs.CpuArchitecture_ARM64(),
//   		FargateOperatingSystemFamily: ecs.OperatingSystemFamily_LINUX(),
//   	}),
//   })
//
type EcsFargateContainerDefinitionProps struct {
	// The number of vCPUs reserved for the container.
	//
	// Each vCPU is equivalent to 1,024 CPU shares.
	// For containers running on EC2 resources, you must specify at least one vCPU.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
	// The image that this container will run.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// The memory hard limit present to the container.
	//
	// If your container attempts to exceed the memory specified, the container is terminated.
	// You must specify at least 4 MiB of memory for a job.
	Memory awscdk.Size `field:"required" json:"memory" yaml:"memory"`
	// The command that's passed to the container.
	// See: https://docs.docker.com/engine/reference/builder/#cmd
	//
	// Default: - no command.
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The environment variables to pass to a container.
	//
	// Cannot start with `AWS_BATCH`.
	// We don't recommend using plaintext environment variables for sensitive information, such as credential data.
	// Default: - no environment variables.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The role used by Amazon ECS container and AWS Fargate agents to make AWS API calls on your behalf.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/execution-IAM-role.html
	//
	// Default: - a Role will be created.
	//
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The role that the container can assume.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html
	//
	// Default: - no job role.
	//
	JobRole awsiam.IRole `field:"optional" json:"jobRole" yaml:"jobRole"`
	// Linux-specific modifications that are applied to the container, such as details for device mappings.
	// Default: none.
	//
	LinuxParameters LinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The loging configuration for this Job.
	// Default: - the log configuration of the Docker daemon.
	//
	Logging awsecs.LogDriver `field:"optional" json:"logging" yaml:"logging"`
	// Gives the container readonly access to its root filesystem.
	// Default: false.
	//
	ReadonlyRootFilesystem *bool `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// A map from environment variable names to the secrets for the container.
	//
	// Allows your job definitions
	// to reference the secret by the environment variable name defined in this property.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/specifying-sensitive-data.html
	//
	// Default: - no secrets.
	//
	Secrets *map[string]Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// The user name to use inside the container.
	// Default: - no user.
	//
	User *string `field:"optional" json:"user" yaml:"user"`
	// The volumes to mount to this container.
	//
	// Automatically added to the job definition.
	// Default: - no volumes.
	//
	Volumes *[]EcsVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// Indicates whether the job has a public IP address.
	//
	// For a job that's running on Fargate resources in a private subnet to send outbound traffic to the internet
	// (for example, to pull container images), the private subnet requires a NAT gateway be attached to route requests to the internet.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html
	//
	// Default: false.
	//
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The size for ephemeral storage.
	// Default: - 20 GiB.
	//
	EphemeralStorageSize awscdk.Size `field:"optional" json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
	// The vCPU architecture of Fargate Runtime.
	// Default: - X86_64.
	//
	FargateCpuArchitecture awsecs.CpuArchitecture `field:"optional" json:"fargateCpuArchitecture" yaml:"fargateCpuArchitecture"`
	// The operating system for the compute environment.
	// Default: - LINUX.
	//
	FargateOperatingSystemFamily awsecs.OperatingSystemFamily `field:"optional" json:"fargateOperatingSystemFamily" yaml:"fargateOperatingSystemFamily"`
	// Which version of Fargate to use when running this container.
	// Default: LATEST.
	//
	FargatePlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"fargatePlatformVersion" yaml:"fargatePlatformVersion"`
}

