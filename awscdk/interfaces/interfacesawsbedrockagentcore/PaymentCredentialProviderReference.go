package interfacesawsbedrockagentcore


// A reference to a PaymentCredentialProvider resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   paymentCredentialProviderReference := &PaymentCredentialProviderReference{
//   	CredentialProviderArn: jsii.String("credentialProviderArn"),
//   }
//
type PaymentCredentialProviderReference struct {
	// The CredentialProviderArn of the PaymentCredentialProvider resource.
	CredentialProviderArn *string `field:"required" json:"credentialProviderArn" yaml:"credentialProviderArn"`
}

