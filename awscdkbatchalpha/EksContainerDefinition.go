package awscdkbatchalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A container that can be run with EKS orchestration on EC2 resources.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   jobDefn := batch.NewEksJobDefinition(this, jsii.String("eksf2"), &EksJobDefinitionProps{
//   	Container: batch.NewEksContainerDefinition(this, jsii.String("container"), &EksContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		Volumes: []eksVolume{
//   			batch.*eksVolume_EmptyDir(&EmptyDirVolumeOptions{
//   				Name: jsii.String("myEmptyDirVolume"),
//   				MountPath: jsii.String("/mount/path"),
//   				Medium: batch.EmptyDirMediumType_MEMORY,
//   				Readonly: jsii.Boolean(true),
//   				SizeLimit: cdk.Size_Mebibytes(jsii.Number(2048)),
//   			}),
//   		},
//   	}),
//   })
//
// Experimental.
type EksContainerDefinition interface {
	constructs.Construct
	IEksContainerDefinition
	// An array of arguments to the entrypoint.
	//
	// If this isn't specified, the CMD of the container image is used.
	// This corresponds to the args member in the Entrypoint portion of the Pod in Kubernetes.
	// Environment variable references are expanded using the container's environment.
	// If the referenced environment variable doesn't exist, the reference in the command isn't changed.
	// For example, if the reference is to "$(NAME1)" and the NAME1 environment variable doesn't exist,
	// the command string will remain "$(NAME1)." $$ is replaced with $, and the resulting string isn't expanded.
	// or example, $$(VAR_NAME) is passed as $(VAR_NAME) whether or not the VAR_NAME environment variable exists.
	// Experimental.
	Args() *[]*string
	// The entrypoint for the container.
	//
	// This isn't run within a shell.
	// If this isn't specified, the `ENTRYPOINT` of the container image is used.
	// Environment variable references are expanded using the container's environment.
	// If the referenced environment variable doesn't exist, the reference in the command isn't changed.
	// For example, if the reference is to `"$(NAME1)"` and the `NAME1` environment variable doesn't exist,
	// the command string will remain `"$(NAME1)."` `$$` is replaced with `$` and the resulting string isn't expanded.
	// For example, `$$(VAR_NAME)` will be passed as `$(VAR_NAME)` whether or not the `VAR_NAME` environment variable exists.
	//
	// The entrypoint can't be updated.
	// Experimental.
	Command() *[]*string
	// The hard limit of CPUs to present to this container. Must be an even multiple of 0.25.
	//
	// If your container attempts to exceed this limit, it will be terminated.
	//
	// At least one of `cpuReservation` and `cpuLimit` is required.
	// If both are specified, then `cpuLimit` must be at least as large as `cpuReservation`.
	// Experimental.
	CpuLimit() *float64
	// The soft limit of CPUs to reserve for the container Must be an even multiple of 0.25.
	//
	// The container will given at least this many CPUs, but may consume more.
	//
	// At least one of `cpuReservation` and `cpuLimit` is required.
	// If both are specified, then `cpuLimit` must be at least as large as `cpuReservation`.
	// Experimental.
	CpuReservation() *float64
	// The environment variables to pass to this container.
	//
	// *Note*: Environment variables cannot start with "AWS_BATCH".
	// This naming convention is reserved for variables that AWS Batch sets.
	// Experimental.
	Env() *map[string]*string
	// The hard limit of GPUs to present to this container.
	//
	// If your container attempts to exceed this limit, it will be terminated.
	//
	// If both `gpuReservation` and `gpuLimit` are specified, then `gpuLimit` must be equal to `gpuReservation`.
	// Experimental.
	GpuLimit() *float64
	// The soft limit of CPUs to reserve for the container Must be an even multiple of 0.25.
	//
	// The container will given at least this many CPUs, but may consume more.
	//
	// If both `gpuReservation` and `gpuLimit` are specified, then `gpuLimit` must be equal to `gpuReservation`.
	// Experimental.
	GpuReservation() *float64
	// The image that this container will run.
	// Experimental.
	Image() awsecs.ContainerImage
	// The image pull policy for this container.
	// Experimental.
	ImagePullPolicy() ImagePullPolicy
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, it will be terminated.
	//
	// Must be larger that 4 MiB
	//
	// At least one of `memoryLimit` and `memoryReservation` is required
	//
	// *Note*: To maximize your resource utilization, provide your jobs with as much memory as possible
	// for the specific instance type that you are using.
	// Experimental.
	MemoryLimit() awscdk.Size
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// Your container will be given at least this much memory, but may consume more.
	//
	// Must be larger that 4 MiB
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of `memoryLimit` and `memoryReservation` is required.
	// If both are specified, then `memoryLimit` must be equal to `memoryReservation`
	//
	// *Note*: To maximize your resource utilization, provide your jobs with as much memory as possible
	// for the specific instance type that you are using.
	// Experimental.
	MemoryReservation() awscdk.Size
	// The name of this container.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// If specified, gives this container elevated permissions on the host container instance.
	//
	// The level of permissions are similar to the root user permissions.
	//
	// This parameter maps to `privileged` policy in the Privileged pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// Experimental.
	Privileged() *bool
	// If specified, gives this container readonly access to its root file system.
	//
	// This parameter maps to `ReadOnlyRootFilesystem` policy in the Volumes and file systems pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// Experimental.
	ReadonlyRootFilesystem() *bool
	// If specified, the container is run as the specified group ID (`gid`).
	//
	// If this parameter isn't specified, the default is the group that's specified in the image metadata.
	// This parameter maps to `RunAsGroup` and `MustRunAs` policy in the Users and groups pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// Experimental.
	RunAsGroup() *float64
	// If specified, the container is run as a user with a `uid` other than 0.
	//
	// Otherwise, no such rule is enforced.
	// This parameter maps to `RunAsUser` and `MustRunAsNonRoot` policy in the Users and groups pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// Experimental.
	RunAsRoot() *bool
	// If specified, this container is run as the specified user ID (`uid`).
	//
	// This parameter maps to `RunAsUser` and `MustRunAs` policy in the Users and groups pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// Experimental.
	RunAsUser() *float64
	// The Volumes to mount to this container.
	//
	// Automatically added to the Pod.
	// Experimental.
	Volumes() *[]EksVolume
	// Mount a Volume to this container.
	//
	// Automatically added to the Pod.
	// Experimental.
	AddVolume(volume EksVolume)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EksContainerDefinition
type jsiiProxy_EksContainerDefinition struct {
	internal.Type__constructsConstruct
	jsiiProxy_IEksContainerDefinition
}

func (j *jsiiProxy_EksContainerDefinition) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) CpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) CpuReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) Env() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) GpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) GpuReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpuReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) Image() awsecs.ContainerImage {
	var returns awsecs.ContainerImage
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) ImagePullPolicy() ImagePullPolicy {
	var returns ImagePullPolicy
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) MemoryLimit() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) MemoryReservation() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) Privileged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) ReadonlyRootFilesystem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) RunAsGroup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) RunAsRoot() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"runAsRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) RunAsUser() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksContainerDefinition) Volumes() *[]EksVolume {
	var returns *[]EksVolume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Experimental.
func NewEksContainerDefinition(scope constructs.Construct, id *string, props *EksContainerDefinitionProps) EksContainerDefinition {
	_init_.Initialize()

	if err := validateNewEksContainerDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EksContainerDefinition{}

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.EksContainerDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEksContainerDefinition_Override(e EksContainerDefinition, scope constructs.Construct, id *string, props *EksContainerDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-batch-alpha.EksContainerDefinition",
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
func EksContainerDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksContainerDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-batch-alpha.EksContainerDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksContainerDefinition) AddVolume(volume EksVolume) {
	if err := e.validateAddVolumeParameters(volume); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addVolume",
		[]interface{}{volume},
	)
}

func (e *jsiiProxy_EksContainerDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

