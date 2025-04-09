package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The details of a tmpfs mount for a container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var size size
//
//   tmpfs := &Tmpfs{
//   	ContainerPath: jsii.String("containerPath"),
//   	Size: size,
//
//   	// the properties below are optional
//   	MountOptions: []tmpfsMountOption{
//   		awscdk.Aws_batch.*tmpfsMountOption_DEFAULTS,
//   	},
//   }
//
type Tmpfs struct {
	// The absolute file path where the tmpfs volume is to be mounted.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// The size (in MiB) of the tmpfs volume.
	Size awscdk.Size `field:"required" json:"size" yaml:"size"`
	// The list of tmpfs volume mount options.
	//
	// For more information, see
	// [TmpfsMountOptions](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Tmpfs.html).
	// Default: none.
	//
	MountOptions *[]TmpfsMountOption `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

