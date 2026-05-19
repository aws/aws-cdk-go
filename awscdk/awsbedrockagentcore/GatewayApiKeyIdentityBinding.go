package awsbedrockagentcore


// Provider and secret ARNs for wiring a Token Vault API key identity into a gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayApiKeyIdentityBinding := &GatewayApiKeyIdentityBinding{
//   	ProviderArn: jsii.String("providerArn"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
type GatewayApiKeyIdentityBinding struct {
	// API key credential provider ARN.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// Secrets Manager secret ARN for the API key material.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

