package interfacesawsguardduty


// A reference to a Master resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   masterReference := &MasterReference{
//   	DetectorId: jsii.String("detectorId"),
//   	MasterId: jsii.String("masterId"),
//   }
//
type MasterReference struct {
	// The DetectorId of the Master resource.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The MasterId of the Master resource.
	MasterId *string `field:"required" json:"masterId" yaml:"masterId"`
}

