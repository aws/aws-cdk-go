package awscognito


// Signing algorithms for SAML requests.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   // specify the metadata as a file content
//   // specify the metadata as a file content
//   cognito.NewUserPoolIdentityProviderSaml(this, jsii.String("userpoolIdpFile"), &UserPoolIdentityProviderSamlProps{
//   	UserPool: userpool,
//   	Metadata: cognito.UserPoolIdentityProviderSamlMetadata_File(jsii.String("my-file-contents")),
//   	// Whether to require encrypted SAML assertions from IdP
//   	EncryptedResponses: jsii.Boolean(true),
//   	// The signing algorithm for the SAML requests
//   	RequestSigningAlgorithm: cognito.SigningAlgorithm_RSA_SHA256,
//   	// Enable IdP initiated SAML auth flow
//   	IdpInitiated: jsii.Boolean(true),
//   })
//
//   // specify the metadata as a URL
//   // specify the metadata as a URL
//   cognito.NewUserPoolIdentityProviderSaml(this, jsii.String("userpoolidpUrl"), &UserPoolIdentityProviderSamlProps{
//   	UserPool: userpool,
//   	Metadata: cognito.UserPoolIdentityProviderSamlMetadata_Url(jsii.String("https://my-metadata-url.com")),
//   })
//
type SigningAlgorithm string

const (
	// RSA with SHA-256.
	SigningAlgorithm_RSA_SHA256 SigningAlgorithm = "RSA_SHA256"
)

