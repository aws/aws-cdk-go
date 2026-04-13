package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tmpfsProperty := &TmpfsProperty{
//   	Size: jsii.Number(123),
//
//   	// the properties below are optional
//   	ContainerPath: jsii.String("containerPath"),
//   	MountOptions: []*string{
//   		jsii.String("mountOptions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-tmpfs.html
//
type CfnDaemonTaskDefinition_TmpfsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-tmpfs.html#cfn-ecs-daemontaskdefinition-tmpfs-size
	//
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-tmpfs.html#cfn-ecs-daemontaskdefinition-tmpfs-containerpath
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-tmpfs.html#cfn-ecs-daemontaskdefinition-tmpfs-mountoptions
	//
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

