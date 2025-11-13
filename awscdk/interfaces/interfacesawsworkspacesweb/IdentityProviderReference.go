package interfacesawsworkspacesweb


// A reference to a IdentityProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityProviderReference := &IdentityProviderReference{
//   	IdentityProviderArn: jsii.String("identityProviderArn"),
//   }
//
type IdentityProviderReference struct {
	// The IdentityProviderArn of the IdentityProvider resource.
	IdentityProviderArn *string `field:"required" json:"identityProviderArn" yaml:"identityProviderArn"`
}

