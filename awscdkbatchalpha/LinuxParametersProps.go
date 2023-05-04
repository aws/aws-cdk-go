// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The properties for defining Linux-specific options that are applied to the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var size size
//
//   linuxParametersProps := &LinuxParametersProps{
//   	InitProcessEnabled: jsii.Boolean(false),
//   	MaxSwap: size,
//   	SharedMemorySize: size,
//   	Swappiness: jsii.Number(123),
//   }
//
// Experimental.
type LinuxParametersProps struct {
	// Specifies whether to run an init process inside the container that forwards signals and reaps processes.
	// Experimental.
	InitProcessEnabled *bool `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// The total amount of swap memory a container can use.
	//
	// This parameter
	// will be translated to the --memory-swap option to docker run.
	//
	// This parameter is only supported when you are using the EC2 launch type.
	// Accepted values are positive integers.
	// Experimental.
	MaxSwap awscdk.Size `field:"optional" json:"maxSwap" yaml:"maxSwap"`
	// The value for the size of the /dev/shm volume.
	// Experimental.
	SharedMemorySize awscdk.Size `field:"optional" json:"sharedMemorySize" yaml:"sharedMemorySize"`
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
	// Experimental.
	Swappiness *float64 `field:"optional" json:"swappiness" yaml:"swappiness"`
}

