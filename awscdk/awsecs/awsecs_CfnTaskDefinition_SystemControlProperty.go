package awsecs


// A list of namespaced kernel parameters to set in the container.
//
// This parameter maps to `Sysctls` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--sysctl` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
//
// We don't recommend that you specify network-related `systemControls` parameters for multiple containers in a single task. This task also uses either the `awsvpc` or `host` network mode. It does it for the following reasons.
//
// - For tasks that use the `awsvpc` network mode, if you set `systemControls` for any container, it applies to all containers in the task. If you set different `systemControls` for multiple containers in a single task, the container that's started last determines which `systemControls` take effect.
// - For tasks that use the `host` network mode, the `systemControls` parameter applies to the container instance's kernel parameter and that of all containers of any tasks running on that container instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemControlProperty := &systemControlProperty{
//   	namespace: jsii.String("namespace"),
//   	value: jsii.String("value"),
//   }
//
type CfnTaskDefinition_SystemControlProperty struct {
	// The namespaced kernel parameter to set a `value` for.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The value for the namespaced kernel parameter that's specified in `namespace` .
	Value *string `field:"optional" json:"value" yaml:"value"`
}

