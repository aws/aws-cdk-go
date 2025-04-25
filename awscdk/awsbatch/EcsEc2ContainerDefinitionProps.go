package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props to configure an EcsEc2ContainerDefinition.
//
// Example:
//   var vpc iVpc
//
//
//   ecsJob := batch.NewEcsJobDefinition(this, jsii.String("JobDefn"), &EcsJobDefinitionProps{
//   	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("containerDefn"), &EcsEc2ContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/amazonlinux/amazonlinux:latest")),
//   		Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   		Cpu: jsii.Number(256),
//   	}),
//   })
//
//   queue := batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
//   	ComputeEnvironments: []orderedComputeEnvironment{
//   		&orderedComputeEnvironment{
//   			ComputeEnvironment: batch.NewManagedEc2EcsComputeEnvironment(this, jsii.String("managedEc2CE"), &ManagedEc2EcsComputeEnvironmentProps{
//   				Vpc: *Vpc,
//   			}),
//   			Order: jsii.Number(1),
//   		},
//   	},
//   	Priority: jsii.Number(10),
//   })
//
//   user := iam.NewUser(this, jsii.String("MyUser"))
//   ecsJob.GrantSubmitJob(user, queue)
//
type EcsEc2ContainerDefinitionProps struct {
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
	// The number of physical GPUs to reserve for the container.
	//
	// Make sure that the number of GPUs reserved for all containers in a job doesn't exceed
	// the number of available GPUs on the compute resource that the job is launched on.
	// Default: - no gpus.
	//
	Gpu *float64 `field:"optional" json:"gpu" yaml:"gpu"`
	// When this parameter is true, the container is given elevated permissions on the host container instance (similar to the root user).
	// Default: false.
	//
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// Limits to set for the user this docker container will run as.
	// Default: - no ulimits.
	//
	Ulimits *[]*Ulimit `field:"optional" json:"ulimits" yaml:"ulimits"`
}

