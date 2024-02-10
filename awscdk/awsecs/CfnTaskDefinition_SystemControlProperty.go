package awsecs


// A list of namespaced kernel parameters to set in the container.
//
// This parameter maps to `Sysctls` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--sysctl` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) . For example, you can configure `net.ipv4.tcp_keepalive_time` setting to maintain longer lived connections.
//
// We don't recommend that you specify network-related `systemControls` parameters for multiple containers in a single task that also uses either the `awsvpc` or `host` network mode. Doing this has the following disadvantages:
//
// - For tasks that use the `awsvpc` network mode including Fargate, if you set `systemControls` for any container, it applies to all containers in the task. If you set different `systemControls` for multiple containers in a single task, the container that's started last determines which `systemControls` take effect.
// - For tasks that use the `host` network mode, the network namespace `systemControls` aren't supported.
//
// If you're setting an IPC resource namespace to use for the containers in the task, the following conditions apply to your system controls. For more information, see [IPC mode](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#task_definition_ipcmode) .
//
// - For tasks that use the `host` IPC mode, IPC namespace `systemControls` aren't supported.
// - For tasks that use the `task` IPC mode, IPC namespace `systemControls` values apply to all containers within a task.
//
// > This parameter is not supported for Windows containers. > This parameter is only supported for tasks that are hosted on AWS Fargate if the tasks are using platform version `1.4.0` or later (Linux). This isn't supported for Windows containers on Fargate.
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

