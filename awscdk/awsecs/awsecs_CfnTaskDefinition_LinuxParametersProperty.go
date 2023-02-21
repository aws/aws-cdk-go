package awsecs


// The `LinuxParameters` property specifies Linux-specific options that are applied to the container, such as Linux [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linuxParametersProperty := &linuxParametersProperty{
//   	capabilities: &kernelCapabilitiesProperty{
//   		add: []*string{
//   			jsii.String("add"),
//   		},
//   		drop: []*string{
//   			jsii.String("drop"),
//   		},
//   	},
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
//   			size: jsii.Number(123),
//
//   			// the properties below are optional
//   			containerPath: jsii.String("containerPath"),
//   			mountOptions: []*string{
//   				jsii.String("mountOptions"),
//   			},
//   		},
//   	},
//   }
//
type CfnTaskDefinition_LinuxParametersProperty struct {
	// The Linux capabilities for the container that are added to or dropped from the default configuration provided by Docker.
	//
	// > For tasks that use the Fargate launch type, `capabilities` is supported for all platform versions but the `add` parameter is only supported if using platform version 1.4.0 or later.
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Any host devices to expose to the container.
	//
	// This parameter maps to `Devices` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `--device` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you're using tasks that use the Fargate launch type, the `devices` parameter isn't supported.
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// Run an `init` process inside the container that forwards signals and reaps processes.
	//
	// This parameter maps to the `--init` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) . This parameter requires version 1.25 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// The total amount of swap memory (in MiB) a container can use.
	//
	// This parameter will be translated to the `--memory-swap` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) where the value would be the sum of the container memory plus the `maxSwap` value.
	//
	// If a `maxSwap` value of `0` is specified, the container will not use swap. Accepted values are `0` or any positive integer. If the `maxSwap` parameter is omitted, the container will use the swap configuration for the container instance it is running on. A `maxSwap` value must be set for the `swappiness` parameter to be used.
	//
	// > If you're using tasks that use the Fargate launch type, the `maxSwap` parameter isn't supported.
	MaxSwap *float64 `field:"optional" json:"maxSwap" yaml:"maxSwap"`
	// The value for the size (in MiB) of the `/dev/shm` volume.
	//
	// This parameter maps to the `--shm-size` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you are using tasks that use the Fargate launch type, the `sharedMemorySize` parameter is not supported.
	SharedMemorySize *float64 `field:"optional" json:"sharedMemorySize" yaml:"sharedMemorySize"`
	// This allows you to tune a container's memory swappiness behavior.
	//
	// A `swappiness` value of `0` will cause swapping to not happen unless absolutely necessary. A `swappiness` value of `100` will cause pages to be swapped very aggressively. Accepted values are whole numbers between `0` and `100` . If the `swappiness` parameter is not specified, a default value of `60` is used. If a value is not specified for `maxSwap` then this parameter is ignored. This parameter maps to the `--memory-swappiness` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you're using tasks that use the Fargate launch type, the `swappiness` parameter isn't supported.
	Swappiness *float64 `field:"optional" json:"swappiness" yaml:"swappiness"`
	// The container path, mount options, and size (in MiB) of the tmpfs mount.
	//
	// This parameter maps to the `--tmpfs` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) .
	//
	// > If you're using tasks that use the Fargate launch type, the `tmpfs` parameter isn't supported.
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

