package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Props to configure an EcsEc2ContainerDefinition.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiNodeJob := batch.NewMultiNodeJobDefinition(this, jsii.String("JobDefinition"), &MultiNodeJobDefinitionProps{
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R4, ec2.InstanceSize_LARGE),
//   	Containers: []multiNodeContainer{
//   		&multiNodeContainer{
//   			Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("mainMPIContainer"), &EcsEc2ContainerDefinitionProps{
//   				Image: ecs.ContainerImage_FromRegistry(jsii.String("yourregsitry.com/yourMPIImage:latest")),
//   				Cpu: jsii.Number(256),
//   				Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   			}),
//   			StartNode: jsii.Number(0),
//   			EndNode: jsii.Number(5),
//   		},
//   	},
//   })
//   // convenience method
//   multiNodeJob.AddContainer(&multiNodeContainer{
//   	StartNode: jsii.Number(6),
//   	EndNode: jsii.Number(10),
//   	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("multiContainer"), &EcsEc2ContainerDefinitionProps{
//   		Image: ecs.ContainerImage_*FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		Cpu: jsii.Number(256),
//   		Memory: cdk.Size_*Mebibytes(jsii.Number(2048)),
//   	}),
//   })
//
// Experimental.
type EcsEc2ContainerDefinitionProps struct {
	// The number of vCPUs reserved for the container.
	//
	// Each vCPU is equivalent to 1,024 CPU shares.
	// For containers running on EC2 resources, you must specify at least one vCPU.
	// Experimental.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
	// The image that this container will run.
	// Experimental.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// The memory hard limit present to the container.
	//
	// If your container attempts to exceed the memory specified, the container is terminated.
	// You must specify at least 4 MiB of memory for a job.
	// Experimental.
	Memory awscdk.Size `field:"required" json:"memory" yaml:"memory"`
	// The command that's passed to the container.
	// See: https://docs.docker.com/engine/reference/builder/#cmd
	//
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The environment variables to pass to a container.
	//
	// Cannot start with `AWS_BATCH`.
	// We don't recommend using plaintext environment variables for sensitive information, such as credential data.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The role used by Amazon ECS container and AWS Fargate agents to make AWS API calls on your behalf.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/execution-IAM-role.html
	//
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The role that the container can assume.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html
	//
	// Experimental.
	JobRole awsiam.IRole `field:"optional" json:"jobRole" yaml:"jobRole"`
	// Linux-specific modifications that are applied to the container, such as details for device mappings.
	// Experimental.
	LinuxParameters LinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The loging configuration for this Job.
	// Experimental.
	Logging awsecs.LogDriver `field:"optional" json:"logging" yaml:"logging"`
	// Gives the container readonly access to its root filesystem.
	// Experimental.
	ReadonlyRootFilesystem *bool `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// The secrets for the container.
	//
	// Can be referenced in your job definition.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/specifying-sensitive-data.html
	//
	// Experimental.
	Secrets *[]awssecretsmanager.ISecret `field:"optional" json:"secrets" yaml:"secrets"`
	// The user name to use inside the container.
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
	// The volumes to mount to this container.
	//
	// Automatically added to the job definition.
	// Experimental.
	Volumes *[]EcsVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// The number of physical GPUs to reserve for the container.
	//
	// Make sure that the number of GPUs reserved for all containers in a job doesn't exceed
	// the number of available GPUs on the compute resource that the job is launched on.
	// Experimental.
	Gpu *float64 `field:"optional" json:"gpu" yaml:"gpu"`
	// When this parameter is true, the container is given elevated permissions on the host container instance (similar to the root user).
	// Experimental.
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// Limits to set for the user this docker container will run as.
	// Experimental.
	Ulimits *[]*Ulimit `field:"optional" json:"ulimits" yaml:"ulimits"`
}

