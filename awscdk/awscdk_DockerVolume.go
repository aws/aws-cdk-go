// Version 2 of the AWS Cloud Development Kit library
package awscdk


// A Docker volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerVolume := &dockerVolume{
//   	containerPath: jsii.String("containerPath"),
//   	hostPath: jsii.String("hostPath"),
//
//   	// the properties below are optional
//   	consistency: cdk.dockerVolumeConsistency_CONSISTENT,
//   }
//
type DockerVolume struct {
	// The path where the file or directory is mounted in the container.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// The path to the file or directory on the host machine.
	HostPath *string `field:"required" json:"hostPath" yaml:"hostPath"`
	// Mount consistency.
	//
	// Only applicable for macOS.
	// See: https://docs.docker.com/storage/bind-mounts/#configure-mount-consistency-for-macos
	//
	Consistency DockerVolumeConsistency `field:"optional" json:"consistency" yaml:"consistency"`
}

