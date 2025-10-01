package awsguardduty


// A reference to a ThreatEntitySet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   threatEntitySetReference := &ThreatEntitySetReference{
//   	DetectorId: jsii.String("detectorId"),
//   	ThreatEntitySetId: jsii.String("threatEntitySetId"),
//   }
//
type ThreatEntitySetReference struct {
	// The DetectorId of the ThreatEntitySet resource.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The Id of the ThreatEntitySet resource.
	ThreatEntitySetId *string `field:"required" json:"threatEntitySetId" yaml:"threatEntitySetId"`
}

