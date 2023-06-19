package awsecs


// The details of a tmpfs mount for a container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tmpfs := &Tmpfs{
//   	ContainerPath: jsii.String("containerPath"),
//   	Size: jsii.Number(123),
//
//   	// the properties below are optional
//   	MountOptions: []tmpfsMountOption{
//   		awscdk.Aws_ecs.*tmpfsMountOption_DEFAULTS,
//   	},
//   }
//
// Experimental.
type Tmpfs struct {
	// The absolute file path where the tmpfs volume is to be mounted.
	// Experimental.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// The size (in MiB) of the tmpfs volume.
	// Experimental.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// The list of tmpfs volume mount options.
	//
	// For more information, see
	// [TmpfsMountOptions](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Tmpfs.html).
	// Experimental.
	MountOptions *[]TmpfsMountOption `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

