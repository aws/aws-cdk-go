package interfacesawsguardduty


// A reference to a Filter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterReference := &FilterReference{
//   	DetectorId: jsii.String("detectorId"),
//   	FilterName: jsii.String("filterName"),
//   }
//
type FilterReference struct {
	// The DetectorId of the Filter resource.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The Name of the Filter resource.
	FilterName *string `field:"required" json:"filterName" yaml:"filterName"`
}

