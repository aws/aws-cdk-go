package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Props to configure an EksContainerDefinition.
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
type EksContainerDefinitionProps struct {
	// The image that this container will run.
	// Experimental.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
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
	// Experimental.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
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
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The hard limit of CPUs to present to this container. Must be an even multiple of 0.25.
	//
	// If your container attempts to exceed this limit, it will be terminated.
	//
	// At least one of `cpuReservation` and `cpuLimit` is required.
	// If both are specified, then `cpuLimit` must be at least as large as `cpuReservation`.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	// Experimental.
	CpuLimit *float64 `field:"optional" json:"cpuLimit" yaml:"cpuLimit"`
	// The soft limit of CPUs to reserve for the container Must be an even multiple of 0.25.
	//
	// The container will given at least this many CPUs, but may consume more.
	//
	// At least one of `cpuReservation` and `cpuLimit` is required.
	// If both are specified, then `cpuLimit` must be at least as large as `cpuReservation`.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	// Experimental.
	CpuReservation *float64 `field:"optional" json:"cpuReservation" yaml:"cpuReservation"`
	// The environment variables to pass to this container.
	//
	// *Note*: Environment variables cannot start with "AWS_BATCH".
	// This naming convention is reserved for variables that AWS Batch sets.
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// The hard limit of GPUs to present to this container.
	//
	// If your container attempts to exceed this limit, it will be terminated.
	//
	// If both `gpuReservation` and `gpuLimit` are specified, then `gpuLimit` must be equal to `gpuReservation`.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	// Experimental.
	GpuLimit *float64 `field:"optional" json:"gpuLimit" yaml:"gpuLimit"`
	// The soft limit of CPUs to reserve for the container Must be an even multiple of 0.25.
	//
	// The container will given at least this many CPUs, but may consume more.
	//
	// If both `gpuReservation` and `gpuLimit` are specified, then `gpuLimit` must be equal to `gpuReservation`.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	// Experimental.
	GpuReservation *float64 `field:"optional" json:"gpuReservation" yaml:"gpuReservation"`
	// The image pull policy for this container.
	// See: https://kubernetes.io/docs/concepts/containers/images/#updating-images
	//
	// Experimental.
	ImagePullPolicy ImagePullPolicy `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
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
	// Experimental.
	MemoryLimit awscdk.Size `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
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
	// Experimental.
	MemoryReservation awscdk.Size `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// The name of this container.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// If specified, gives this container elevated permissions on the host container instance.
	//
	// The level of permissions are similar to the root user permissions.
	//
	// This parameter maps to `privileged` policy in the Privileged pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// See: https://kubernetes.io/docs/concepts/security/pod-security-policy/#volumes-and-file-systems
	//
	// Experimental.
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// If specified, gives this container readonly access to its root file system.
	//
	// This parameter maps to `ReadOnlyRootFilesystem` policy in the Volumes and file systems pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// See: https://kubernetes.io/docs/concepts/security/pod-security-policy/#volumes-and-file-systems
	//
	// Experimental.
	ReadonlyRootFilesystem *bool `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// If specified, the container is run as the specified group ID (`gid`).
	//
	// If this parameter isn't specified, the default is the group that's specified in the image metadata.
	// This parameter maps to `RunAsGroup` and `MustRunAs` policy in the Users and groups pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// See: https://kubernetes.io/docs/concepts/security/pod-security-policy/#users-and-groups
	//
	// Experimental.
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// If specified, the container is run as a user with a `uid` other than 0.
	//
	// Otherwise, no such rule is enforced.
	// This parameter maps to `RunAsUser` and `MustRunAsNonRoot` policy in the Users and groups pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// See: https://kubernetes.io/docs/concepts/security/pod-security-policy/#users-and-groups
	//
	// Experimental.
	RunAsRoot *bool `field:"optional" json:"runAsRoot" yaml:"runAsRoot"`
	// If specified, this container is run as the specified user ID (`uid`).
	//
	// This parameter maps to `RunAsUser` and `MustRunAs` policy in the Users and groups pod security policies in the Kubernetes documentation.
	//
	// *Note*: this is only compatible with Kubernetes < v1.25
	// See: https://kubernetes.io/docs/concepts/security/pod-security-policy/#users-and-groups
	//
	// Experimental.
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	// The Volumes to mount to this container.
	//
	// Automatically added to the Pod.
	// See: https://kubernetes.io/docs/concepts/storage/volumes/
	//
	// Experimental.
	Volumes *[]EksVolume `field:"optional" json:"volumes" yaml:"volumes"`
}

