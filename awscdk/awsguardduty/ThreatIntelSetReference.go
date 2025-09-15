package awsguardduty


// A reference to a ThreatIntelSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   threatIntelSetReference := &ThreatIntelSetReference{
//   	DetectorId: jsii.String("detectorId"),
//   	ThreatIntelSetId: jsii.String("threatIntelSetId"),
//   }
//
type ThreatIntelSetReference struct {
	// The DetectorId of the ThreatIntelSet resource.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The Id of the ThreatIntelSet resource.
	ThreatIntelSetId *string `field:"required" json:"threatIntelSetId" yaml:"threatIntelSetId"`
}

