package interfacesawsguardduty


// A reference to a Detector resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   detectorReference := &DetectorReference{
//   	DetectorId: jsii.String("detectorId"),
//   }
//
type DetectorReference struct {
	// The Id of the Detector resource.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
}

