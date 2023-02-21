package awsecs


// The `MountPoint` property specifies details on a volume mount point that is used in a container definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mountPointProperty := &mountPointProperty{
//   	containerPath: jsii.String("containerPath"),
//   	readOnly: jsii.Boolean(false),
//   	sourceVolume: jsii.String("sourceVolume"),
//   }
//
type CfnTaskDefinition_MountPointProperty struct {
	// The path on the container to mount the host volume at.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// If this value is `true` , the container has read-only access to the volume.
	//
	// If this value is `false` , then the container can write to the volume. The default value is `false` .
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The name of the volume to mount.
	//
	// Must be a volume name referenced in the `name` parameter of task definition `volume` .
	SourceVolume *string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
}

