package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// A Batch MachineImage that is compatible with EKS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var machineImage iMachineImage
//
//   eksMachineImage := &EksMachineImage{
//   	Image: machineImage,
//   	ImageType: batch_alpha.EksMachineImageType_EKS_AL2,
//   }
//
// Experimental.
type EksMachineImage struct {
	// The machine image to use.
	// Experimental.
	Image awsec2.IMachineImage `field:"optional" json:"image" yaml:"image"`
	// Tells Batch which instance type to launch this image on.
	// Experimental.
	ImageType EksMachineImageType `field:"optional" json:"imageType" yaml:"imageType"`
}

