package awsecs


// The details of a tmpfs mount for a container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tmpfs := &tmpfs{
//   	containerPath: jsii.String("containerPath"),
//   	size: jsii.Number(123),
//
//   	// the properties below are optional
//   	mountOptions: []tmpfsMountOption{
//   		awscdk.Aws_ecs.*tmpfsMountOption_DEFAULTS,
//   	},
//   }
//
type Tmpfs struct {
	// The absolute file path where the tmpfs volume is to be mounted.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// The size (in MiB) of the tmpfs volume.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// The list of tmpfs volume mount options.
	//
	// For more information, see
	// [TmpfsMountOptions](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Tmpfs.html).
	MountOptions *[]TmpfsMountOption `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

