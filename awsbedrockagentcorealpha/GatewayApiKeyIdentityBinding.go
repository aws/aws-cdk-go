package awsbedrockagentcorealpha


// Provider and secret ARNs for wiring a Token Vault API key identity into a gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   gatewayApiKeyIdentityBinding := &GatewayApiKeyIdentityBinding{
//   	ProviderArn: jsii.String("providerArn"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// Experimental.
type GatewayApiKeyIdentityBinding struct {
	// API key credential provider ARN.
	// Experimental.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// Secrets Manager secret ARN for the API key material.
	// Experimental.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

