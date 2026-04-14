package interfacesawsbedrockagentcore


// A reference to a ApiKeyCredentialProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyCredentialProviderReference := &ApiKeyCredentialProviderReference{
//   	CredentialProviderArn: jsii.String("credentialProviderArn"),
//   }
//
type ApiKeyCredentialProviderReference struct {
	// The CredentialProviderArn of the ApiKeyCredentialProvider resource.
	CredentialProviderArn *string `field:"required" json:"credentialProviderArn" yaml:"credentialProviderArn"`
}

