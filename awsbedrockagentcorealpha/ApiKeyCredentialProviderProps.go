package awsbedrockagentcorealpha


// API key credential provider ARNs for gateway outbound auth (Token Vault identity).
//
// Pass this to {@link GatewayCredentialProvider.fromApiKeyIdentityArn } or to {@link ApiKeyCredentialProviderConfiguration}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var apiKeyCredentialLocation ApiKeyCredentialLocation
//
//   apiKeyCredentialProviderProps := &ApiKeyCredentialProviderProps{
//   	ProviderArn: jsii.String("providerArn"),
//   	SecretArn: jsii.String("secretArn"),
//
//   	// the properties below are optional
//   	CredentialLocation: apiKeyCredentialLocation,
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type ApiKeyCredentialProviderProps struct {
	// The API key credential provider ARN.
	//
	// This is returned when creating the API key credential provider via Console or API.
	// Format: arn:aws:bedrock-agentcore:region:account:token-vault/id/apikeycredentialprovider/name.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// The ARN of the Secrets Manager secret containing the API key.
	//
	// This is returned when creating the API key credential provider via Console or API.
	// Format: arn:aws:secretsmanager:region:account:secret:name.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The location of the API key credential.
	//
	// This field specifies where in the request the API key should be placed.
	// Default: - HEADER.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialLocation ApiKeyCredentialLocation `field:"optional" json:"credentialLocation" yaml:"credentialLocation"`
}

