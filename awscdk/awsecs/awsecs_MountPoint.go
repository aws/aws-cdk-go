package awsecs


// The details of data volume mount points for a container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mountPoint := &mountPoint{
//   	containerPath: jsii.String("containerPath"),
//   	readOnly: jsii.Boolean(false),
//   	sourceVolume: jsii.String("sourceVolume"),
//   }
//
type MountPoint struct {
	// The path on the container to mount the host volume at.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// Specifies whether to give the container read-only access to the volume.
	//
	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume.
	ReadOnly *bool `field:"required" json:"readOnly" yaml:"readOnly"`
	// The name of the volume to mount.
	//
	// Must be a volume name referenced in the name parameter of task definition volume.
	SourceVolume *string `field:"required" json:"sourceVolume" yaml:"sourceVolume"`
}

