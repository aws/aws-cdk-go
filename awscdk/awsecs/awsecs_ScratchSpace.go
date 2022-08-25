package awsecs


// The temporary disk space mounted to the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scratchSpace := &scratchSpace{
//   	containerPath: jsii.String("containerPath"),
//   	name: jsii.String("name"),
//   	readOnly: jsii.Boolean(false),
//   	sourcePath: jsii.String("sourcePath"),
//   }
//
type ScratchSpace struct {
	// The path on the container to mount the scratch volume at.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// The name of the scratch volume to mount.
	//
	// Must be a volume name referenced in the name parameter of task definition volume.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies whether to give the container read-only access to the scratch volume.
	//
	// If this value is true, the container has read-only access to the scratch volume.
	// If this value is false, then the container can write to the scratch volume.
	ReadOnly *bool `field:"required" json:"readOnly" yaml:"readOnly"`
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
}

