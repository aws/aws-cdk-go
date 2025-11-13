package interfacesawsiam


// A reference to a OIDCProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oIDCProviderReference := &OIDCProviderReference{
//   	OidcProviderArn: jsii.String("oidcProviderArn"),
//   }
//
type OIDCProviderReference struct {
	// The Arn of the OIDCProvider resource.
	OidcProviderArn *string `field:"required" json:"oidcProviderArn" yaml:"oidcProviderArn"`
}

