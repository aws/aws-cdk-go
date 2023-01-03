package awsbatch


// Details on a Docker volume mount point that's used in a job's container properties.
//
// This parameter maps to `Volumes` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/api/docker_remote_api_v1.19/#create-a-container) section of the Docker Remote API and the `--volume` option to docker run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mountPointsProperty := &mountPointsProperty{
//   	containerPath: jsii.String("containerPath"),
//   	readOnly: jsii.Boolean(false),
//   	sourceVolume: jsii.String("sourceVolume"),
//   }
//
type CfnJobDefinition_MountPointsProperty struct {
	// The path on the container where the host volume is mounted.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// If this value is `true` , the container has read-only access to the volume.
	//
	// Otherwise, the container can write to the volume. The default value is `false` .
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The name of the volume to mount.
	SourceVolume *string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
}

