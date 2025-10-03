package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A container that can be run with ECS orchestration.
type IEcsContainerDefinition interface {
	constructs.IConstruct
	// Add a Volume to this container.
	AddVolume(volume EcsVolume)
	// The command that's passed to the container.
	// See: https://docs.docker.com/engine/reference/builder/#cmd
	//
	Command() *[]*string
	// The number of vCPUs reserved for the container.
	//
	// Each vCPU is equivalent to 1,024 CPU shares.
	// For containers running on EC2 resources, you must specify at least one vCPU.
	Cpu() *float64
	// The environment variables to pass to a container.
	//
	// Cannot start with `AWS_BATCH`.
	// We don't recommend using plaintext environment variables for sensitive information, such as credential data.
	// Default: - no environment variables.
	//
	Environment() *map[string]*string
	// The role used by Amazon ECS container and AWS Fargate agents to make AWS API calls on your behalf.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/execution-IAM-role.html
	//
	ExecutionRole() awsiam.IRole
	// The image that this container will run.
	Image() awsecs.ContainerImage
	// The role that the container can assume.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html
	//
	// Default: - no jobRole.
	//
	JobRole() awsiam.IRole
	// Linux-specific modifications that are applied to the container, such as details for device mappings.
	// Default: none.
	//
	LinuxParameters() LinuxParameters
	// The configuration of the log driver.
	LogDriverConfig() *awsecs.LogDriverConfig
	// The memory hard limit present to the container.
	//
	// If your container attempts to exceed the memory specified, the container is terminated.
	// You must specify at least 4 MiB of memory for a job.
	Memory() awscdk.Size
	// Gives the container readonly access to its root filesystem.
	// Default: false.
	//
	ReadonlyRootFilesystem() *bool
	// A map from environment variable names to the secrets for the container.
	//
	// Allows your job definitions
	// to reference the secret by the environment variable name defined in this property.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/specifying-sensitive-data.html
	//
	// Default: - no secrets.
	//
	Secrets() *map[string]Secret
	// The user name to use inside the container.
	// Default: - no user.
	//
	User() *string
	// The volumes to mount to this container.
	//
	// Automatically added to the job definition.
	// Default: - no volumes.
	//
	Volumes() *[]EcsVolume
}

// The jsii proxy for IEcsContainerDefinition
type jsiiProxy_IEcsContainerDefinition struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IEcsContainerDefinition) AddVolume(volume EcsVolume) {
	if err := i.validateAddVolumeParameters(volume); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addVolume",
		[]interface{}{volume},
	)
}

func (j *jsiiProxy_IEcsContainerDefinition) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) Image() awsecs.ContainerImage {
	var returns awsecs.ContainerImage
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) JobRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"jobRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) LinuxParameters() LinuxParameters {
	var returns LinuxParameters
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) LogDriverConfig() *awsecs.LogDriverConfig {
	var returns *awsecs.LogDriverConfig
	_jsii_.Get(
		j,
		"logDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) Memory() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) ReadonlyRootFilesystem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) Secrets() *map[string]Secret {
	var returns *map[string]Secret
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsContainerDefinition) Volumes() *[]EcsVolume {
	var returns *[]EcsVolume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

