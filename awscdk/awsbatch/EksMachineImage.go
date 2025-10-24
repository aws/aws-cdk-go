package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// A Batch MachineImage that is compatible with EKS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var machineImage IMachineImage
//
//   eksMachineImage := &EksMachineImage{
//   	Image: machineImage,
//   	ImageType: awscdk.Aws_batch.EksMachineImageType_EKS_AL2,
//   }
//
type EksMachineImage struct {
	// The machine image to use.
	// Default: - chosen by batch.
	//
	Image awsec2.IMachineImage `field:"optional" json:"image" yaml:"image"`
	// Tells Batch which instance type to launch this image on.
	// Default: - 'EKS_AL2' for non-gpu instances, 'EKS_AL2_NVIDIA' for gpu instances.
	//
	ImageType EksMachineImageType `field:"optional" json:"imageType" yaml:"imageType"`
}

