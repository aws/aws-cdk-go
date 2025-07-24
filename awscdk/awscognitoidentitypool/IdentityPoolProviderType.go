package awscognitoidentitypool


// Types of Identity Pool Login Providers.
type IdentityPoolProviderType string

const (
	// Facebook provider type.
	IdentityPoolProviderType_FACEBOOK IdentityPoolProviderType = "FACEBOOK"
	// Google provider type.
	IdentityPoolProviderType_GOOGLE IdentityPoolProviderType = "GOOGLE"
	// Amazon provider type.
	IdentityPoolProviderType_AMAZON IdentityPoolProviderType = "AMAZON"
	// Apple provider type.
	IdentityPoolProviderType_APPLE IdentityPoolProviderType = "APPLE"
	// Twitter provider type.
	IdentityPoolProviderType_TWITTER IdentityPoolProviderType = "TWITTER"
	// Open Id provider type.
	IdentityPoolProviderType_OPEN_ID IdentityPoolProviderType = "OPEN_ID"
	// Saml provider type.
	IdentityPoolProviderType_SAML IdentityPoolProviderType = "SAML"
	// User Pool provider type.
	IdentityPoolProviderType_USER_POOL IdentityPoolProviderType = "USER_POOL"
	// Custom provider type.
	IdentityPoolProviderType_CUSTOM IdentityPoolProviderType = "CUSTOM"
)

