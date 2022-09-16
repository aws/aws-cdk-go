package awsecs


// The properties for defining Linux-specific options that are applied to the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linuxParametersProps := &linuxParametersProps{
//   	initProcessEnabled: jsii.Boolean(false),
//   	sharedMemorySize: jsii.Number(123),
//   }
//
// Experimental.
type LinuxParametersProps struct {
	// Specifies whether to run an init process inside the container that forwards signals and reaps processes.
	// Experimental.
	InitProcessEnabled *bool `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// The value for the size (in MiB) of the /dev/shm volume.
	// Experimental.
	SharedMemorySize *float64 `field:"optional" json:"sharedMemorySize" yaml:"sharedMemorySize"`
}

