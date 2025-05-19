package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// A Batch MachineImage that is compatible with ECS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var machineImage iMachineImage
//
//   ecsMachineImage := &EcsMachineImage{
//   	Image: machineImage,
//   	ImageType: awscdk.Aws_batch.EcsMachineImageType_ECS_AL2,
//   }
//
type EcsMachineImage struct {
	// The machine image to use.
	// Default: - chosen by batch.
	//
	Image awsec2.IMachineImage `field:"optional" json:"image" yaml:"image"`
	// Tells Batch which instance type to launch this image on.
	// Default: - 'ECS_AL2' for non-gpu instances, 'ECS_AL2_NVIDIA' for gpu instances.
	//
	ImageType EcsMachineImageType `field:"optional" json:"imageType" yaml:"imageType"`
}

