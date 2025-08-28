package awsecs


// The base details of where a volume will be mounted within a container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseMountPoint := &BaseMountPoint{
//   	ContainerPath: jsii.String("containerPath"),
//   	ReadOnly: jsii.Boolean(false),
//   }
//
type BaseMountPoint struct {
	// The path on the container to mount the host volume at.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// Specifies whether to give the container read-only access to the volume.
	//
	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume.
	ReadOnly *bool `field:"required" json:"readOnly" yaml:"readOnly"`
}

