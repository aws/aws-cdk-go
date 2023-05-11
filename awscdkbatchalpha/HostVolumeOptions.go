// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// Options for configuring an ECS HostVolume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
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
// Experimental.
type HostVolumeOptions struct {
	// the path on the container where this volume is mounted.
	// Experimental.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// the name of this volume.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// if set, the container will have readonly access to the volume.
	// Experimental.
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
	// The path on the host machine this container will have access to.
	// Experimental.
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
}

