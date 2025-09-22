package awsguardduty


// A reference to a TrustedEntitySet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trustedEntitySetReference := &TrustedEntitySetReference{
//   	DetectorId: jsii.String("detectorId"),
//   	TrustedEntitySetId: jsii.String("trustedEntitySetId"),
//   }
//
type TrustedEntitySetReference struct {
	// The DetectorId of the TrustedEntitySet resource.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The Id of the TrustedEntitySet resource.
	TrustedEntitySetId *string `field:"required" json:"trustedEntitySetId" yaml:"trustedEntitySetId"`
}

