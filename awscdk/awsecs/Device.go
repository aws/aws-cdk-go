package awsecs


// A container instance host device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   device := &Device{
//   	HostPath: jsii.String("hostPath"),
//
//   	// the properties below are optional
//   	ContainerPath: jsii.String("containerPath"),
//   	Permissions: []devicePermission{
//   		awscdk.Aws_ecs.*devicePermission_READ,
//   	},
//   }
//
type Device struct {
	// The path for the device on the host container instance.
	HostPath *string `field:"required" json:"hostPath" yaml:"hostPath"`
	// The path inside the container at which to expose the host device.
	// Default: Same path as the host.
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The explicit permissions to provide to the container for the device.
	//
	// By default, the container has permissions for read, write, and mknod for the device.
	// Default: Readonly.
	//
	Permissions *[]DevicePermission `field:"optional" json:"permissions" yaml:"permissions"`
}

