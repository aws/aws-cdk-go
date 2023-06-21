package awscdkbatchalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A container orchestrated by ECS that uses EC2 resources.
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
type EcsEc2ContainerDefinition interface {
	constructs.Construct
	IEcsContainerDefinition
	IEcsEc2ContainerDefinition
	// The command that's passed to the container.
	// Experimental.
	Command() *[]*string
	// The number of vCPUs reserved for the container.
	//
	// Each vCPU is equivalent to 1,024 CPU shares.
	// For containers running on EC2 resources, you must specify at least one vCPU.
	// Experimental.
	Cpu() *float64
	// The environment variables to pass to a container.
	//
	// Cannot start with `AWS_BATCH`.
	// We don't recommend using plaintext environment variables for sensitive information, such as credential data.
	// Experimental.
	Environment() *map[string]*string
	// The role used by Amazon ECS container and AWS Fargate agents to make AWS API calls on your behalf.
	// Experimental.
	ExecutionRole() awsiam.IRole
	// The number of physical GPUs to reserve for the container.
	//
	// Make sure that the number of GPUs reserved for all containers in a job doesn't exceed
	// the number of available GPUs on the compute resource that the job is launched on.
	// Experimental.
	Gpu() *float64
	// The image that this container will run.
	// Experimental.
	Image() awsecs.ContainerImage
	// The role that the container can assume.
	// Experimental.
	JobRole() awsiam.IRole
	// Linux-specific modifications that are applied to the container, such as details for device mappings.
	// Experimental.
	LinuxParameters() LinuxParameters
	// The configuration of the log driver.
	// Experimental.
	LogDriverConfig() *awsecs.LogDriverConfig
	// The memory hard limit present to the container.
	//
	// If your container attempts to exceed the memory specified, the container is terminated.
	// You must specify at least 4 MiB of memory for a job.
	// Experimental.
	Memory() awscdk.Size
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// When this parameter is true, the container is given elevated permissions on the host container instance (similar to the root user).
	// Experimental.
	Privileged() *bool
	// Gives the container readonly access to its root filesystem.
	// Experimental.
	ReadonlyRootFilesystem() *bool
	// The secrets for the container.
	//
	// Can be referenced in your job definition.
	// Experimental.
	Secrets() *[]awssecretsmanager.ISecret
	// Limits to set for the user this docker container will run as.
	// Experimental.
	Ulimits() *[]*Ulimit
	// The user name to use inside the container.
	// Experimental.
	User() *string
	// The volumes to mount to this container.
	//
	// Automatically added to the job definition.
	// Experimental.
	Volumes() *[]EcsVolume
	// Add a ulimit to this container.
	// Experimental.
	AddUlimit(ulimit *Ulimit)
	// Add a Volume to this container.
	// Experimental.
	AddVolume(volume EcsVolume)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsEc2ContainerDefinition
type jsiiProxy_EcsEc2ContainerDefinition struct {
	internal.Type__constructsConstruct
	jsiiProxy_IEcsContainerDefinition
	jsiiProxy_IEcsEc2ContainerDefinition
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) Gpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) Image() awsecs.ContainerImage {
	var returns awsecs.ContainerImage
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) JobRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"jobRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) LinuxParameters() LinuxParameters {
	var returns LinuxParameters
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) LogDriverConfig() *awsecs.LogDriverConfig {
	var returns *awsecs.LogDriverConfig
	_jsii_.Get(
		j,
		"logDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) Memory() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) Privileged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) ReadonlyRootFilesystem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) Secrets() *[]awssecretsmanager.ISecret {
	var returns *[]awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) Ulimits() *[]*Ulimit {
	var returns *[]*Ulimit
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsEc2ContainerDefinition) Volumes() *[]EcsVolume {
	var returns *[]EcsVolume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Experimental.
func NewEcsEc2ContainerDefinition(scope constructs.Construct, id *string, props *EcsEc2ContainerDefinitionProps) EcsEc2ContainerDefinition {
	_init_.Initialize()

	if err := validateNewEcsEc2ContainerDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsEc2ContainerDefinition{}

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.EcsEc2ContainerDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsEc2ContainerDefinition_Override(e EcsEc2ContainerDefinition, scope constructs.Construct, id *string, props *EcsEc2ContainerDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.EcsEc2ContainerDefinition",
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
// Experimental.
func EcsEc2ContainerDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsEc2ContainerDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.EcsEc2ContainerDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsEc2ContainerDefinition) AddUlimit(ulimit *Ulimit) {
	if err := e.validateAddUlimitParameters(ulimit); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addUlimit",
		[]interface{}{ulimit},
	)
}

func (e *jsiiProxy_EcsEc2ContainerDefinition) AddVolume(volume EcsVolume) {
	if err := e.validateAddVolumeParameters(volume); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addVolume",
		[]interface{}{volume},
	)
}

func (e *jsiiProxy_EcsEc2ContainerDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

