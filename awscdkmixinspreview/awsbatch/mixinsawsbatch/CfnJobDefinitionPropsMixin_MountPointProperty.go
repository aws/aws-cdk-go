package mixinsawsbatch


// Details for a Docker volume mount point that's used in a job's container properties.
//
// This parameter maps to `Volumes` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.43/#tag/Container/operation/ContainerCreate) section of the *Docker Remote API* and the `--volume` option to docker run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mountPointProperty := &MountPointProperty{
//   	ContainerPath: jsii.String("containerPath"),
//   	ReadOnly: jsii.Boolean(false),
//   	SourceVolume: jsii.String("sourceVolume"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoint.html
//
type CfnJobDefinitionPropsMixin_MountPointProperty struct {
	// The path on the container where the host volume is mounted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoint.html#cfn-batch-jobdefinition-mountpoint-containerpath
	//
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// If this value is `true` , the container has read-only access to the volume.
	//
	// Otherwise, the container can write to the volume. The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoint.html#cfn-batch-jobdefinition-mountpoint-readonly
	//
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The name of the volume to mount.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoint.html#cfn-batch-jobdefinition-mountpoint-sourcevolume
	//
	SourceVolume *string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
}

