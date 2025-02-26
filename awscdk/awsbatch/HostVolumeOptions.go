package awsbatch


// Options for configuring an ECS HostVolume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostVolumeOptions := &HostVolumeOptions{
//   	ContainerPath: jsii.String("containerPath"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	HostPath: jsii.String("hostPath"),
//   	Readonly: jsii.Boolean(false),
//   }
//
type HostVolumeOptions struct {
	// the path on the container where this volume is mounted.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// the name of this volume.
	Name *string `field:"required" json:"name" yaml:"name"`
	// if set, the container will have readonly access to the volume.
	// Default: false.
	//
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
	// The path on the host machine this container will have access to.
	// Default: - Docker will choose the host path.
	// The data may not persist after the containers that use it stop running.
	//
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
}

