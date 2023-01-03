package awsbatch


// Linux-specific modifications that are applied to the container, such as details for device mappings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linuxParametersProperty := &linuxParametersProperty{
//   	devices: []interface{}{
//   		&deviceProperty{
//   			containerPath: jsii.String("containerPath"),
//   			hostPath: jsii.String("hostPath"),
//   			permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   		},
//   	},
//   	initProcessEnabled: jsii.Boolean(false),
//   	maxSwap: jsii.Number(123),
//   	sharedMemorySize: jsii.Number(123),
//   	swappiness: jsii.Number(123),
//   	tmpfs: []interface{}{
//   		&tmpfsProperty{
//   			containerPath: jsii.String("containerPath"),
//   			size: jsii.Number(123),
//
//   			// the properties below are optional
//   			mountOptions: []*string{
//   				jsii.String("mountOptions"),
//   			},
//   		},
//   	},
//   }
//
type CfnJobDefinition_LinuxParametersProperty struct {
	// Any host devices to expose to the container.
	//
	// This parameter maps to `Devices` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `--device` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources and shouldn't be provided.
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// If true, run an `init` process inside the container that forwards signals and reaps processes.
	//
	// This parameter maps to the `--init` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) . This parameter requires version 1.25 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log into your container instance and run the following command: `sudo docker version | grep "Server API version"`
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// The total amount of swap memory (in MiB) a container can use.
	//
	// This parameter is translated to the `--memory-swap` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) where the value is the sum of the container memory plus the `maxSwap` value. For more information, see [`--memory-swap` details](https://docs.aws.amazon.com/https://docs.docker.com/config/containers/resource_constraints/#--memory-swap-details) in the Docker documentation.
	//
	// If a `maxSwap` value of `0` is specified, the container doesn't use swap. Accepted values are `0` or any positive integer. If the `maxSwap` parameter is omitted, the container doesn't use the swap configuration for the container instance it is running on. A `maxSwap` value must be set for the `swappiness` parameter to be used.
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources and shouldn't be provided.
	MaxSwap *float64 `field:"optional" json:"maxSwap" yaml:"maxSwap"`
	// The value for the size (in MiB) of the `/dev/shm` volume.
	//
	// This parameter maps to the `--shm-size` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources and shouldn't be provided.
	SharedMemorySize *float64 `field:"optional" json:"sharedMemorySize" yaml:"sharedMemorySize"`
	// This allows you to tune a container's memory swappiness behavior.
	//
	// A `swappiness` value of `0` causes swapping not to happen unless absolutely necessary. A `swappiness` value of `100` causes pages to be swapped very aggressively. Accepted values are whole numbers between `0` and `100` . If the `swappiness` parameter isn't specified, a default value of `60` is used. If a value isn't specified for `maxSwap` , then this parameter is ignored. If `maxSwap` is set to 0, the container doesn't use swap. This parameter maps to the `--memory-swappiness` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
	//
	// Consider the following when you use a per-container swap configuration.
	//
	// - Swap space must be enabled and allocated on the container instance for the containers to use.
	//
	// > The Amazon ECS optimized AMIs don't have swap enabled by default. You must enable swap on the instance to use this feature. For more information, see [Instance store swap volumes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-store-swap-volumes.html) in the *Amazon EC2 User Guide for Linux Instances* or [How do I allocate memory to work as swap space in an Amazon EC2 instance by using a swap file?](https://docs.aws.amazon.com/premiumsupport/knowledge-center/ec2-memory-swap-file/)
	// - The swap space parameters are only supported for job definitions using EC2 resources.
	// - If the `maxSwap` and `swappiness` parameters are omitted from a job definition, each container will have a default `swappiness` value of 60, and the total swap usage will be limited to two times the memory reservation of the container.
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources and shouldn't be provided.
	Swappiness *float64 `field:"optional" json:"swappiness" yaml:"swappiness"`
	// The container path, mount options, and size (in MiB) of the tmpfs mount.
	//
	// This parameter maps to the `--tmpfs` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources and shouldn't be provided.
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

