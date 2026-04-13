package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mountPointProperty := &MountPointProperty{
//   	ContainerPath: jsii.String("containerPath"),
//   	ReadOnly: jsii.Boolean(false),
//   	SourceVolume: jsii.String("sourceVolume"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-mountpoint.html
//
type CfnDaemonTaskDefinitionPropsMixin_MountPointProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-mountpoint.html#cfn-ecs-daemontaskdefinition-mountpoint-containerpath
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-mountpoint.html#cfn-ecs-daemontaskdefinition-mountpoint-readonly
	//
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-mountpoint.html#cfn-ecs-daemontaskdefinition-mountpoint-sourcevolume
	//
	SourceVolume *string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
}

