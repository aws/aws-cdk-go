package awsbatch


// An object that represents a container instance host device.
//
// > This object isn't applicable to jobs that are running on Fargate resources and shouldn't be provided.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceProperty := &DeviceProperty{
//   	ContainerPath: jsii.String("containerPath"),
//   	HostPath: jsii.String("hostPath"),
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-device.html
//
type CfnJobDefinition_DeviceProperty struct {
	// The path inside the container that's used to expose the host device.
	//
	// By default, the `hostPath` value is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-device.html#cfn-batch-jobdefinition-device-containerpath
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The path for the device on the host container instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-device.html#cfn-batch-jobdefinition-device-hostpath
	//
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
	// The explicit permissions to provide to the container for the device.
	//
	// By default, the container has permissions for `read` , `write` , and `mknod` for the device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-device.html#cfn-batch-jobdefinition-device-permissions
	//
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

