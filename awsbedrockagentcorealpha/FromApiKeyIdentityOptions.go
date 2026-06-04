package awsbedrockagentcorealpha


// Optional gateway settings when binding an {@link IApiKeyCredentialProvider} to a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var apiKeyCredentialLocation ApiKeyCredentialLocation
//
//   fromApiKeyIdentityOptions := &FromApiKeyIdentityOptions{
//   	CredentialLocation: apiKeyCredentialLocation,
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type FromApiKeyIdentityOptions struct {
	// Where to place the API key on outbound requests.
	// Default: header `Authorization` with `Bearer ` prefix.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialLocation ApiKeyCredentialLocation `field:"optional" json:"credentialLocation" yaml:"credentialLocation"`
}

