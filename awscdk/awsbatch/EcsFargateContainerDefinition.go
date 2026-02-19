package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A container orchestrated by ECS that uses Fargate resources.
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
type EcsFargateContainerDefinition interface {
	constructs.Construct
	IEcsContainerDefinition
	IEcsFargateContainerDefinition
	// Indicates whether the job has a public IP address.
	//
	// For a job that's running on Fargate resources in a private subnet to send outbound traffic to the internet
	// (for example, to pull container images), the private subnet requires a NAT gateway be attached to route requests to the internet.
	AssignPublicIp() *bool
	// The command that's passed to the container.
	Command() *[]*string
	// The number of vCPUs reserved for the container.
	//
	// Each vCPU is equivalent to 1,024 CPU shares.
	// For containers running on EC2 resources, you must specify at least one vCPU.
	Cpu() *float64
	// Whether to enable ecs exec for this container.
	EnableExecuteCommand() *bool
	// The environment variables to pass to a container.
	//
	// Cannot start with `AWS_BATCH`.
	// We don't recommend using plaintext environment variables for sensitive information, such as credential data.
	Environment() *map[string]*string
	// The size for ephemeral storage.
	EphemeralStorageSize() awscdk.Size
	// The role used by Amazon ECS container and AWS Fargate agents to make AWS API calls on your behalf.
	ExecutionRole() awsiam.IRole
	// The vCPU architecture of Fargate Runtime.
	FargateCpuArchitecture() awsecs.CpuArchitecture
	// The operating system for the compute environment.
	FargateOperatingSystemFamily() awsecs.OperatingSystemFamily
	// Which version of Fargate to use when running this container.
	FargatePlatformVersion() awsecs.FargatePlatformVersion
	// The image that this container will run.
	Image() awsecs.ContainerImage
	// The role that the container can assume.
	JobRole() awsiam.IRole
	// Linux-specific modifications that are applied to the container, such as details for device mappings.
	LinuxParameters() LinuxParameters
	// The configuration of the log driver.
	LogDriverConfig() *awsecs.LogDriverConfig
	// The memory hard limit present to the container.
	//
	// If your container attempts to exceed the memory specified, the container is terminated.
	// You must specify at least 4 MiB of memory for a job.
	Memory() awscdk.Size
	// The tree node.
	Node() constructs.Node
	// Gives the container readonly access to its root filesystem.
	ReadonlyRootFilesystem() *bool
	// A map from environment variable names to the secrets for the container.
	//
	// Allows your job definitions
	// to reference the secret by the environment variable name defined in this property.
	Secrets() *map[string]Secret
	// The user name to use inside the container.
	User() *string
	// The volumes to mount to this container.
	//
	// Automatically added to the job definition.
	Volumes() *[]EcsVolume
	// Add a Volume to this container.
	AddVolume(volume EcsVolume)
	// Returns a string representation of this construct.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for EcsFargateContainerDefinition
type jsiiProxy_EcsFargateContainerDefinition struct {
	internal.Type__constructsConstruct
	jsiiProxy_IEcsContainerDefinition
	jsiiProxy_IEcsFargateContainerDefinition
}

func (j *jsiiProxy_EcsFargateContainerDefinition) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) EnableExecuteCommand() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) EphemeralStorageSize() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"ephemeralStorageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) FargateCpuArchitecture() awsecs.CpuArchitecture {
	var returns awsecs.CpuArchitecture
	_jsii_.Get(
		j,
		"fargateCpuArchitecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) FargateOperatingSystemFamily() awsecs.OperatingSystemFamily {
	var returns awsecs.OperatingSystemFamily
	_jsii_.Get(
		j,
		"fargateOperatingSystemFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) FargatePlatformVersion() awsecs.FargatePlatformVersion {
	var returns awsecs.FargatePlatformVersion
	_jsii_.Get(
		j,
		"fargatePlatformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Image() awsecs.ContainerImage {
	var returns awsecs.ContainerImage
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) JobRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"jobRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) LinuxParameters() LinuxParameters {
	var returns LinuxParameters
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) LogDriverConfig() *awsecs.LogDriverConfig {
	var returns *awsecs.LogDriverConfig
	_jsii_.Get(
		j,
		"logDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Memory() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) ReadonlyRootFilesystem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Secrets() *map[string]Secret {
	var returns *map[string]Secret
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsFargateContainerDefinition) Volumes() *[]EcsVolume {
	var returns *[]EcsVolume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


func NewEcsFargateContainerDefinition(scope constructs.Construct, id *string, props *EcsFargateContainerDefinitionProps) EcsFargateContainerDefinition {
	_init_.Initialize()

	if err := validateNewEcsFargateContainerDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsFargateContainerDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.EcsFargateContainerDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEcsFargateContainerDefinition_Override(e EcsFargateContainerDefinition, scope constructs.Construct, id *string, props *EcsFargateContainerDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.EcsFargateContainerDefinition",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func EcsFargateContainerDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsFargateContainerDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.EcsFargateContainerDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EcsFargateContainerDefinition_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_batch.EcsFargateContainerDefinition",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EcsFargateContainerDefinition) AddVolume(volume EcsVolume) {
	if err := e.validateAddVolumeParameters(volume); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addVolume",
		[]interface{}{volume},
	)
}

func (e *jsiiProxy_EcsFargateContainerDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsFargateContainerDefinition) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

