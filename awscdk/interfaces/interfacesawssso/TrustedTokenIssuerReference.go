package interfacesawssso


// A reference to a TrustedTokenIssuer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trustedTokenIssuerReference := &TrustedTokenIssuerReference{
//   	TrustedTokenIssuerArn: jsii.String("trustedTokenIssuerArn"),
//   }
//
type TrustedTokenIssuerReference struct {
	// The TrustedTokenIssuerArn of the TrustedTokenIssuer resource.
	TrustedTokenIssuerArn *string `field:"required" json:"trustedTokenIssuerArn" yaml:"trustedTokenIssuerArn"`
}

