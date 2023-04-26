// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// A container instance host device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//
//   device := &Device{
//   	HostPath: jsii.String("hostPath"),
//
//   	// the properties below are optional
//   	ContainerPath: jsii.String("containerPath"),
//   	Permissions: []devicePermission{
//   		batch_alpha.*devicePermission_READ,
//   	},
//   }
//
// Experimental.
type Device struct {
	// The path for the device on the host container instance.
	// Experimental.
	HostPath *string `field:"required" json:"hostPath" yaml:"hostPath"`
	// The path inside the container at which to expose the host device.
	// Experimental.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The explicit permissions to provide to the container for the device.
	//
	// By default, the container has permissions for read, write, and mknod for the device.
	// Experimental.
	Permissions *[]DevicePermission `field:"optional" json:"permissions" yaml:"permissions"`
}

