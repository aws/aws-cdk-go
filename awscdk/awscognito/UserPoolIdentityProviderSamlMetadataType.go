package awscognito


// Metadata types that can be used for a SAML user pool identity provider.
type UserPoolIdentityProviderSamlMetadataType string

const (
	// Metadata provided via a URL.
	UserPoolIdentityProviderSamlMetadataType_URL UserPoolIdentityProviderSamlMetadataType = "URL"
	// Metadata provided via the contents of a file.
	UserPoolIdentityProviderSamlMetadataType_FILE UserPoolIdentityProviderSamlMetadataType = "FILE"
)

