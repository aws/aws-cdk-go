package awsbatch


// Options to configure an EcsVolume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsVolumeOptions := &EcsVolumeOptions{
//   	ContainerPath: jsii.String("containerPath"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Readonly: jsii.Boolean(false),
//   }
//
type EcsVolumeOptions struct {
	// the path on the container where this volume is mounted.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// the name of this volume.
	Name *string `field:"required" json:"name" yaml:"name"`
	// if set, the container will have readonly access to the volume.
	// Default: false.
	//
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
}

