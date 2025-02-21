package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

// A container that can be run with EKS orchestration on EC2 resources.
type IEksContainerDefinition interface {
	constructs.IConstruct
	// Mount a Volume to this container.
	//
	// Automatically added to the Pod.
	AddVolume(volume EksVolume)
	// An array of arguments to the entrypoint.
	//
	// If this isn't specified, the CMD of the container image is used.
	// This corresponds to the args member in the Entrypoint portion of the Pod in Kubernetes.
	// Environment variable references are expanded using the container's environment.
	// If the referenced environment variable doesn't exist, the reference in the command isn't changed.
	// For example, if the reference is to "$(NAME1)" and the NAME1 environment variable doesn't exist,
	// the command string will remain "$(NAME1)." $$ is replaced with $, and the resulting string isn't expanded.
	// or example, $$(VAR_NAME) is passed as $(VAR_NAME) whether or not the VAR_NAME environment variable exists.
	// See: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/
	//
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
	// See: https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#entrypoint
	//
	Command() *[]*string
	// The hard limit of CPUs to present to this container. Must be an even multiple of 0.25.
	//
	// If your container attempts to exceed this limit, it will be terminated.
	//
	// At least one of `cpuReservation` and `cpuLimit` is required.
	// If both are specified, then `cpuLimit` must be at least as large as `cpuReservation`.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	// Default: - No CPU limit.
	//
	CpuLimit() *float64
	// The soft limit of CPUs to reserve for the container Must be an even multiple of 0.25.
	//
	// The container will given at least this many CPUs, but may consume more.
	//
	// At least one of `cpuReservation` and `cpuLimit` is required.
	// If both are specified, then `cpuLimit` must be at least as large as `cpuReservation`.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	// Default: - No CPUs reserved.
	//
	CpuReservation() *float64
	// The environment variables to pass to this container.
	//
	// *Note*: Environment variables cannot start with "AWS_BATCH".
	// This naming convention is reserved for variables that AWS Batch sets.
	Env() *map[string]*string
	// The hard limit of GPUs to present to this container.
	//
	// If your container attempts to exceed this limit, it will be terminated.
	//
	// If both `gpuReservation` and `gpuLimit` are specified, then `gpuLimit` must be equal to `gpuReservation`.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	// Default: - No GPU limit.
	//
	GpuLimit() *float64
	// The soft limit of CPUs to reserve for the container Must be an even multiple of 0.25.
	//
	// The container will given at least this many CPUs, but may consume more.
	//
	// If both `gpuReservation` and `gpuLimit` are specified, then `gpuLimit` must be equal to `gpuReservation`.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	// Default: - No GPUs reserved.
	//
	GpuReservation() *float64
	// The image that this container will run.
	Image() awsecs.ContainerImage
	// The image pull policy for this container.
	// See: https://kubernetes.io/docs/concepts/containers/images/#updating-images
	//
	// Default: - `ALWAYS` if the `:latest` tag is specified, `IF_NOT_PRESENT` otherwise.
	//
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
	// See: https://docs.aws.amazon.com/batch/latest/userguide/memory-management.html
	//
	// Default: - No memory limit.
	//
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
	// See: https://docs.aws.amazon.com/batch/latest/userguide/memory-management.html
	//
	// Default: - No memory reserved.
	//
	MemoryReservation() awscdk.Size
	// The name of this container.
	// Default: : `'Default'`.
	//
	Name() *string
	// If specified, gives this container elevated permissions on the host container instance.
	//
	// The level of permissions are similar to the root user permissions.
	//
	// This parameter maps to `privileged` policy in the Privileged pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// See: https://kubernetes.io/docs/concepts/security/pod-security-policy/#volumes-and-file-systems
	//
	// Default: false.
	//
	Privileged() *bool
	// If specified, gives this container readonly access to its root file system.
	//
	// This parameter maps to `ReadOnlyRootFilesystem` policy in the Volumes and file systems pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// See: https://kubernetes.io/docs/concepts/security/pod-security-policy/#volumes-and-file-systems
	//
	// Default: false.
	//
	ReadonlyRootFilesystem() *bool
	// If specified, the container is run as the specified group ID (`gid`).
	//
	// If this parameter isn't specified, the default is the group that's specified in the image metadata.
	// This parameter maps to `RunAsGroup` and `MustRunAs` policy in the Users and groups pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// See: https://kubernetes.io/docs/concepts/security/pod-security-policy/#users-and-groups
	//
	// Default: none.
	//
	RunAsGroup() *float64
	// If specified, the container is run as a user with a `uid` other than 0.
	//
	// Otherwise, no such rule is enforced.
	// This parameter maps to `RunAsUser` and `MustRunAsNonRoot` policy in the Users and groups pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// See: https://kubernetes.io/docs/concepts/security/pod-security-policy/#users-and-groups
	//
	// Default: - the container is *not* required to run as a non-root user.
	//
	RunAsRoot() *bool
	// If specified, this container is run as the specified user ID (`uid`).
	//
	// This parameter maps to `RunAsUser` and `MustRunAs` policy in the Users and groups pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// See: https://kubernetes.io/docs/concepts/security/pod-security-policy/#users-and-groups
	//
	// Default: - the user that is specified in the image metadata.
	//
	RunAsUser() *float64
	// The Volumes to mount to this container.
	//
	// Automatically added to the Pod.
	// See: https://kubernetes.io/docs/concepts/storage/volumes/
	//
	Volumes() *[]EksVolume
}

// The jsii proxy for IEksContainerDefinition
type jsiiProxy_IEksContainerDefinition struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IEksContainerDefinition) AddVolume(volume EksVolume) {
	if err := i.validateAddVolumeParameters(volume); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addVolume",
		[]interface{}{volume},
	)
}

func (j *jsiiProxy_IEksContainerDefinition) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) CpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) CpuReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) Env() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) GpuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) GpuReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpuReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) Image() awsecs.ContainerImage {
	var returns awsecs.ContainerImage
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) ImagePullPolicy() ImagePullPolicy {
	var returns ImagePullPolicy
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) MemoryLimit() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"memoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) MemoryReservation() awscdk.Size {
	var returns awscdk.Size
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) Privileged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) ReadonlyRootFilesystem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) RunAsGroup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) RunAsRoot() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"runAsRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) RunAsUser() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEksContainerDefinition) Volumes() *[]EksVolume {
	var returns *[]EksVolume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

