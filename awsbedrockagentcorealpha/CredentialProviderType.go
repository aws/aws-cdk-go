package awsbedrockagentcorealpha


// Credential provider types supported by gateway target.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type CredentialProviderType string

const (
	// API Key authentication.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderType_API_KEY CredentialProviderType = "API_KEY"
	// OAuth authentication.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderType_OAUTH CredentialProviderType = "OAUTH"
	// IAM role authentication.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderType_GATEWAY_IAM_ROLE CredentialProviderType = "GATEWAY_IAM_ROLE"
)

