package awsecs


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-device.html
//
type CfnDaemonTaskDefinition_DeviceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-device.html#cfn-ecs-daemontaskdefinition-device-containerpath
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-device.html#cfn-ecs-daemontaskdefinition-device-hostpath
	//
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-device.html#cfn-ecs-daemontaskdefinition-device-permissions
	//
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

