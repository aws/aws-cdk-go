package interfacesawsec2


// A reference to a VerifiedAccessEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   verifiedAccessEndpointReference := &VerifiedAccessEndpointReference{
//   	VerifiedAccessEndpointId: jsii.String("verifiedAccessEndpointId"),
//   }
//
type VerifiedAccessEndpointReference struct {
	// The VerifiedAccessEndpointId of the VerifiedAccessEndpoint resource.
	VerifiedAccessEndpointId *string `field:"required" json:"verifiedAccessEndpointId" yaml:"verifiedAccessEndpointId"`
}

