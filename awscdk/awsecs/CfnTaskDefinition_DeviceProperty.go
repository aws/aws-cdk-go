package awsecs


// The `Device` property specifies an object representing a container instance host device.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html
//
type CfnTaskDefinition_DeviceProperty struct {
	// The path inside the container at which to expose the host device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html#cfn-ecs-taskdefinition-device-containerpath
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// The path for the device on the host container instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html#cfn-ecs-taskdefinition-device-hostpath
	//
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
	// The explicit permissions to provide to the container for the device.
	//
	// By default, the container has permissions for `read` , `write` , and `mknod` for the device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html#cfn-ecs-taskdefinition-device-permissions
	//
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

