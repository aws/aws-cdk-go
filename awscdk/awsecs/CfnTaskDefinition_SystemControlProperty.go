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
//   systemControlProperty := &SystemControlProperty{
//   	Namespace: jsii.String("namespace"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-systemcontrol.html
//
type CfnTaskDefinition_SystemControlProperty struct {
	// The namespaced kernel parameter to set a `value` for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-systemcontrol.html#cfn-ecs-taskdefinition-systemcontrol-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The namespaced kernel parameter to set a `value` for.
	//
	// Valid IPC namespace values: `"kernel.msgmax" | "kernel.msgmnb" | "kernel.msgmni" | "kernel.sem" | "kernel.shmall" | "kernel.shmmax" | "kernel.shmmni" | "kernel.shm_rmid_forced"` , and `Sysctls` that start with `"fs.mqueue.*"`
	//
	// Valid network namespace values: `Sysctls` that start with `"net.*"`
	//
	// All of these values are supported by Fargate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-systemcontrol.html#cfn-ecs-taskdefinition-systemcontrol-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

