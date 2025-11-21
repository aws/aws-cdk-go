package awsbedrockagentcorealpha


// Credential provider types supported by gateway target.
// Experimental.
type CredentialProviderType string

const (
	// API Key authentication.
	// Experimental.
	CredentialProviderType_API_KEY CredentialProviderType = "API_KEY"
	// OAuth authentication.
	// Experimental.
	CredentialProviderType_OAUTH CredentialProviderType = "OAUTH"
	// IAM role authentication.
	// Experimental.
	CredentialProviderType_GATEWAY_IAM_ROLE CredentialProviderType = "GATEWAY_IAM_ROLE"
)

