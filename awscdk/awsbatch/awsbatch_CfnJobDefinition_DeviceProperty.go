package awsbatch


// An object representing a container instance host device.
//
// > This object isn't applicable to jobs that are running on Fargate resources and shouldn't be provided.
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
type CfnJobDefinition_DeviceProperty struct {
	// The path inside the container that's used to expose the host device.
	//
	// By default, the `hostPath` value is used.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The path for the device on the host container instance.
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
	// The explicit permissions to provide to the container for the device.
	//
	// By default, the container has permissions for `read` , `write` , and `mknod` for the device.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

