package awsec2


// The IAM SAML identity provider used for federated authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   federatedAuthenticationRequestProperty := &federatedAuthenticationRequestProperty{
//   	samlProviderArn: jsii.String("samlProviderArn"),
//
//   	// the properties below are optional
//   	selfServiceSamlProviderArn: jsii.String("selfServiceSamlProviderArn"),
//   }
//
type CfnClientVpnEndpoint_FederatedAuthenticationRequestProperty struct {
	// The Amazon Resource Name (ARN) of the IAM SAML identity provider.
	SamlProviderArn *string `field:"required" json:"samlProviderArn" yaml:"samlProviderArn"`
	// The Amazon Resource Name (ARN) of the IAM SAML identity provider for the self-service portal.
	SelfServiceSamlProviderArn *string `field:"optional" json:"selfServiceSamlProviderArn" yaml:"selfServiceSamlProviderArn"`
}

