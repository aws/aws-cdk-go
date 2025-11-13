package interfacesawsec2


// A reference to a VerifiedAccessTrustProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   verifiedAccessTrustProviderReference := &VerifiedAccessTrustProviderReference{
//   	VerifiedAccessTrustProviderId: jsii.String("verifiedAccessTrustProviderId"),
//   }
//
type VerifiedAccessTrustProviderReference struct {
	// The VerifiedAccessTrustProviderId of the VerifiedAccessTrustProvider resource.
	VerifiedAccessTrustProviderId *string `field:"required" json:"verifiedAccessTrustProviderId" yaml:"verifiedAccessTrustProviderId"`
}

