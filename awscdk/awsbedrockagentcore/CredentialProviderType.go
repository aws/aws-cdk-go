package awsbedrockagentcore


// Credential provider types supported by gateway target.
type CredentialProviderType string

const (
	// API Key authentication.
	CredentialProviderType_API_KEY CredentialProviderType = "API_KEY"
	// OAuth authentication.
	CredentialProviderType_OAUTH CredentialProviderType = "OAUTH"
	// IAM role authentication.
	CredentialProviderType_GATEWAY_IAM_ROLE CredentialProviderType = "GATEWAY_IAM_ROLE"
)

