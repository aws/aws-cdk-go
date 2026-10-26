package interfacesawslicensemanagerusersubscriptions


// A reference to a LicenseServerEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseServerEndpointReference := &LicenseServerEndpointReference{
//   	IdentityProviderArn: jsii.String("identityProviderArn"),
//   	LicenseServerEndpointArn: jsii.String("licenseServerEndpointArn"),
//   }
//
type LicenseServerEndpointReference struct {
	// The IdentityProviderArn of the LicenseServerEndpoint resource.
	IdentityProviderArn *string `field:"required" json:"identityProviderArn" yaml:"identityProviderArn"`
	// The LicenseServerEndpointArn of the LicenseServerEndpoint resource.
	LicenseServerEndpointArn *string `field:"required" json:"licenseServerEndpointArn" yaml:"licenseServerEndpointArn"`
}

