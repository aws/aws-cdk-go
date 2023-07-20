package awsecs


// The details for a volume mount point that's used in a container definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mountPointProperty := &MountPointProperty{
//   	ContainerPath: jsii.String("containerPath"),
//   	ReadOnly: jsii.Boolean(false),
//   	SourceVolume: jsii.String("sourceVolume"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-mountpoint.html
//
type CfnTaskDefinition_MountPointProperty struct {
	// The path on the container to mount the host volume at.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-mountpoint.html#cfn-ecs-taskdefinition-mountpoint-containerpath
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// If this value is `true` , the container has read-only access to the volume.
	//
	// If this value is `false` , then the container can write to the volume. The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-mountpoint.html#cfn-ecs-taskdefinition-mountpoint-readonly
	//
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The name of the volume to mount.
	//
	// Must be a volume name referenced in the `name` parameter of task definition `volume` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-mountpoint.html#cfn-ecs-taskdefinition-mountpoint-sourcevolume
	//
	SourceVolume *string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
}

