package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   linuxParametersProperty := &LinuxParametersProperty{
//   	Capabilities: &KernelCapabilitiesProperty{
//   		Add: []*string{
//   			jsii.String("add"),
//   		},
//   		Drop: []*string{
//   			jsii.String("drop"),
//   		},
//   	},
//   	Devices: []interface{}{
//   		&DeviceProperty{
//   			ContainerPath: jsii.String("containerPath"),
//   			HostPath: jsii.String("hostPath"),
//   			Permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   		},
//   	},
//   	InitProcessEnabled: jsii.Boolean(false),
//   	Tmpfs: []interface{}{
//   		&TmpfsProperty{
//   			ContainerPath: jsii.String("containerPath"),
//   			MountOptions: []*string{
//   				jsii.String("mountOptions"),
//   			},
//   			Size: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-linuxparameters.html
//
type CfnDaemonTaskDefinitionPropsMixin_LinuxParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-linuxparameters.html#cfn-ecs-daemontaskdefinition-linuxparameters-capabilities
	//
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-linuxparameters.html#cfn-ecs-daemontaskdefinition-linuxparameters-devices
	//
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-linuxparameters.html#cfn-ecs-daemontaskdefinition-linuxparameters-initprocessenabled
	//
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-linuxparameters.html#cfn-ecs-daemontaskdefinition-linuxparameters-tmpfs
	//
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

