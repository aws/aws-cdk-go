package interfacesawsbedrockagentcore


// A reference to a OAuth2CredentialProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuth2CredentialProviderReference := &OAuth2CredentialProviderReference{
//   	CredentialProviderArn: jsii.String("credentialProviderArn"),
//   }
//
type OAuth2CredentialProviderReference struct {
	// The CredentialProviderArn of the OAuth2CredentialProvider resource.
	CredentialProviderArn *string `field:"required" json:"credentialProviderArn" yaml:"credentialProviderArn"`
}

