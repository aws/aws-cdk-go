package awscdkbatchalpha


// Options to configure an EcsVolume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//
//   ecsVolumeOptions := &EcsVolumeOptions{
//   	ContainerPath: jsii.String("containerPath"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Readonly: jsii.Boolean(false),
//   }
//
// Experimental.
type EcsVolumeOptions struct {
	// the path on the container where this volume is mounted.
	// Experimental.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// the name of this volume.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// if set, the container will have readonly access to the volume.
	// Experimental.
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
}

