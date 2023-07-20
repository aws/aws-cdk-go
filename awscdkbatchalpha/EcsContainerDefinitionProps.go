package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props to configure an EcsContainerDefinition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerImage containerImage
//   var ecsVolume ecsVolume
//   var linuxParameters linuxParameters
//   var logDriver logDriver
//   var role role
//   var secret secret
//   var size size
//
//   ecsContainerDefinitionProps := &EcsContainerDefinitionProps{
//   	Cpu: jsii.Number(123),
//   	Image: containerImage,
//   	Memory: size,
//
//   	// the properties below are optional
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	ExecutionRole: role,
//   	JobRole: role,
//   	LinuxParameters: linuxParameters,
//   	Logging: logDriver,
//   	ReadonlyRootFilesystem: jsii.Boolean(false),
//   	Secrets: map[string]*secret{
//   		"secretsKey": secret,
//   	},
//   	User: jsii.String("user"),
//   	Volumes: []*ecsVolume{
//   		ecsVolume,
//   	},
//   }
//
// Experimental.
type EcsContainerDefinitionProps struct {
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
	// A map from environment variable names to the secrets for the container.
	//
	// Allows your job definitions
	// to reference the secret by the environment variable name defined in this property.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/specifying-sensitive-data.html
	//
	// Experimental.
	Secrets *map[string]Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// The user name to use inside the container.
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
	// The volumes to mount to this container.
	//
	// Automatically added to the job definition.
	// Experimental.
	Volumes *[]EcsVolume `field:"optional" json:"volumes" yaml:"volumes"`
}

