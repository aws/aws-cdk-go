package awsiotevents


// A reference to a DetectorModel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   detectorModelReference := &DetectorModelReference{
//   	DetectorModelName: jsii.String("detectorModelName"),
//   }
//
type DetectorModelReference struct {
	// The DetectorModelName of the DetectorModel resource.
	DetectorModelName *string `field:"required" json:"detectorModelName" yaml:"detectorModelName"`
}

