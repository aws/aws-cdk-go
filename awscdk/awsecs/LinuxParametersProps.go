package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The properties for defining Linux-specific options that are applied to the container.
//
// Example:
//   var taskDefinition taskDefinition
//
//
//   taskDefinition.AddContainer(jsii.String("container"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryLimitMiB: jsii.Number(1024),
//   	LinuxParameters: ecs.NewLinuxParameters(this, jsii.String("LinuxParameters"), &LinuxParametersProps{
//   		InitProcessEnabled: jsii.Boolean(true),
//   		SharedMemorySize: jsii.Number(1024),
//   		MaxSwap: awscdk.Size_Mebibytes(jsii.Number(5000)),
//   		Swappiness: jsii.Number(90),
//   	}),
//   })
//
type LinuxParametersProps struct {
	// Specifies whether to run an init process inside the container that forwards signals and reaps processes.
	// Default: false.
	//
	InitProcessEnabled *bool `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// The total amount of swap memory a container can use.
	//
	// This parameter
	// will be translated to the --memory-swap option to docker run.
	//
	// This parameter is only supported when you are using the EC2 launch type.
	// Accepted values are positive integers.
	// Default: No swap.
	//
	MaxSwap awscdk.Size `field:"optional" json:"maxSwap" yaml:"maxSwap"`
	// The value for the size of the /dev/shm volume.
	// Default: No shared memory.
	//
	SharedMemorySize *float64 `field:"optional" json:"sharedMemorySize" yaml:"sharedMemorySize"`
	// This allows you to tune a container's memory swappiness behavior.
	//
	// This parameter
	// maps to the --memory-swappiness option to docker run. The swappiness relates
	// to the kernel's tendency to swap memory. A value of 0 will cause swapping to
	// not happen unless absolutely necessary. A value of 100 will cause pages to
	// be swapped very aggressively.
	//
	// This parameter is only supported when you are using the EC2 launch type.
	// Accepted values are whole numbers between 0 and 100. If a value is not
	// specified for maxSwap then this parameter is ignored.
	// Default: 60.
	//
	Swappiness *float64 `field:"optional" json:"swappiness" yaml:"swappiness"`
}

