package awsbatch


// EKS container properties are used in job definitions for Amazon EKS based job definitions to describe the properties for a container node in the pod that's launched as part of a job.
//
// This can't be specified for Amazon ECS based job definitions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var limits interface{}
//   var requests interface{}
//
//   eksContainerProperty := &EksContainerProperty{
//   	Image: jsii.String("image"),
//
//   	// the properties below are optional
//   	Args: []*string{
//   		jsii.String("args"),
//   	},
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Env: []interface{}{
//   		&EksContainerEnvironmentVariableProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ImagePullPolicy: jsii.String("imagePullPolicy"),
//   	Name: jsii.String("name"),
//   	Resources: &ResourcesProperty{
//   		Limits: limits,
//   		Requests: requests,
//   	},
//   	SecurityContext: &SecurityContextProperty{
//   		Privileged: jsii.Boolean(false),
//   		ReadOnlyRootFilesystem: jsii.Boolean(false),
//   		RunAsGroup: jsii.Number(123),
//   		RunAsNonRoot: jsii.Boolean(false),
//   		RunAsUser: jsii.Number(123),
//   	},
//   	VolumeMounts: []interface{}{
//   		&EksContainerVolumeMountProperty{
//   			MountPath: jsii.String("mountPath"),
//   			Name: jsii.String("name"),
//   			ReadOnly: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnJobDefinition_EksContainerProperty struct {
	// The Docker image used to start the container.
	Image *string `field:"required" json:"image" yaml:"image"`
	// An array of arguments to the entrypoint.
	//
	// If this isn't specified, the `CMD` of the container image is used. This corresponds to the `args` member in the [Entrypoint](https://docs.aws.amazon.com/https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#entrypoint) portion of the [Pod](https://docs.aws.amazon.com/https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/) in Kubernetes. Environment variable references are expanded using the container's environment.
	//
	// If the referenced environment variable doesn't exist, the reference in the command isn't changed. For example, if the reference is to " `$(NAME1)` " and the `NAME1` environment variable doesn't exist, the command string will remain " `$(NAME1)` ." `$$` is replaced with `$` , and the resulting string isn't expanded. For example, `$$(VAR_NAME)` is passed as `$(VAR_NAME)` whether or not the `VAR_NAME` environment variable exists. For more information, see [CMD](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/builder/#cmd) in the *Dockerfile reference* and [Define a command and arguments for a pod](https://docs.aws.amazon.com/https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/) in the *Kubernetes documentation* .
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The entrypoint for the container.
	//
	// This isn't run within a shell. If this isn't specified, the `ENTRYPOINT` of the container image is used. Environment variable references are expanded using the container's environment.
	//
	// If the referenced environment variable doesn't exist, the reference in the command isn't changed. For example, if the reference is to " `$(NAME1)` " and the `NAME1` environment variable doesn't exist, the command string will remain " `$(NAME1)` ." `$$` is replaced with `$` and the resulting string isn't expanded. For example, `$$(VAR_NAME)` will be passed as `$(VAR_NAME)` whether or not the `VAR_NAME` environment variable exists. The entrypoint can't be updated. For more information, see [ENTRYPOINT](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/builder/#entrypoint) in the *Dockerfile reference* and [Define a command and arguments for a container](https://docs.aws.amazon.com/https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/) and [Entrypoint](https://docs.aws.amazon.com/https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#entrypoint) in the *Kubernetes documentation* .
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The environment variables to pass to a container.
	//
	// > Environment variables cannot start with " `AWS_BATCH` ". This naming convention is reserved for variables that AWS Batch sets.
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// The image pull policy for the container.
	//
	// Supported values are `Always` , `IfNotPresent` , and `Never` . This parameter defaults to `IfNotPresent` . However, if the `:latest` tag is specified, it defaults to `Always` . For more information, see [Updating images](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/containers/images/#updating-images) in the *Kubernetes documentation* .
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	// The name of the container.
	//
	// If the name isn't specified, the default name " `Default` " is used. Each container in a pod must have a unique name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type and amount of resources to assign to a container.
	//
	// The supported resources include `memory` , `cpu` , and `nvidia.com/gpu` . For more information, see [Resource management for pods and containers](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) in the *Kubernetes documentation* .
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// The security context for a job.
	//
	// For more information, see [Configure a security context for a pod or container](https://docs.aws.amazon.com/https://kubernetes.io/docs/tasks/configure-pod-container/security-context/) in the *Kubernetes documentation* .
	SecurityContext interface{} `field:"optional" json:"securityContext" yaml:"securityContext"`
	// The volume mounts for the container.
	//
	// AWS Batch supports `emptyDir` , `hostPath` , and `secret` volume types. For more information about volumes and volume mounts in Kubernetes, see [Volumes](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/storage/volumes/) in the *Kubernetes documentation* .
	VolumeMounts interface{} `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

