package awsecs


// The `Device` property specifies an object representing a container instance host device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceProperty := &deviceProperty{
//   	containerPath: jsii.String("containerPath"),
//   	hostPath: jsii.String("hostPath"),
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   }
//
type CfnTaskDefinition_DeviceProperty struct {
	// The path inside the container at which to expose the host device.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The path for the device on the host container instance.
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
	// The explicit permissions to provide to the container for the device.
	//
	// By default, the container has permissions for `read` , `write` , and `mknod` for the device.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

