package awsiam


// Initialization properties for `OIDCProviderNative`.
//
// Example:
//   nativeProvider := iam.NewOidcProviderNative(this, jsii.String("MyProvider"), &OidcProviderNativeProps{
//   	Url: jsii.String("https://openid/connect"),
//   	ClientIds: []*string{
//   		jsii.String("myclient1"),
//   		jsii.String("myclient2"),
//   	},
//   	Thumbprints: []*string{
//   		jsii.String("aa00aa1122aa00aa1122aa00aa1122aa00aa1122"),
//   	},
//   })
//
type OidcProviderNativeProps struct {
	// The URL of the identity provider.
	//
	// The URL must begin with https:// and
	// should correspond to the iss claim in the provider's OpenID Connect ID
	// tokens. Per the OIDC standard, path components are allowed but query
	// parameters are not. Typically the URL consists of only a hostname, like
	// https://server.example.org or https://example.com.
	//
	// You cannot register the same provider multiple times in a single AWS
	// account. If you try to submit a URL that has already been used for an
	// OpenID Connect provider in the AWS account, you will get an error.
	//
	// Warning: This URL cannot contain any port numbers.
	Url *string `field:"required" json:"url" yaml:"url"`
	// A list of client IDs (also known as audiences).
	//
	// When a mobile or web app
	// registers with an OpenID Connect provider, they establish a value that
	// identifies the application. (This is the value that's sent as the client_id
	// parameter on OAuth requests.)
	//
	// You can register multiple client IDs with the same provider. For example,
	// you might have multiple applications that use the same OIDC provider. You
	// cannot register more than 100 client IDs with a single IAM OIDC provider.
	//
	// Client IDs are up to 255 characters long.
	// Default: - no clients are allowed.
	//
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// The name of the Native OIDC Provider.
	// Default: - A name is automatically generated.
	//
	OidcProviderName *string `field:"optional" json:"oidcProviderName" yaml:"oidcProviderName"`
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificates.
	//
	// Typically this list includes only 1 entry or empty. However, IAM lets
	// you have up to 5 thumbprints for an OIDC provider. This lets you maintain
	// multiple thumbprints if the identity provider is rotating certificates.
	//
	// The server certificate thumbprint is the hex-encoded SHA-1 hash value of
	// the X.509 certificate used by the domain where the OpenID Connect provider
	// makes its keys available. It is always a 40-character string.
	//
	// For example, assume that the OIDC provider is server.example.com and the
	// provider stores its keys at https://keys.server.example.com/openid-connect.
	// In that case, the thumbprint string would be the hex-encoded SHA-1 hash
	// value of the certificate used by https://keys.server.example.com.
	//
	// This property is optional. If it is not included, IAM will retrieve and use
	// the top intermediate certificate authority (CA) thumbprint of the OpenID
	// Connect identity provider server certificate.
	//
	// Obtain the thumbprint of the root certificate authority from the provider's
	// server as described in https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc_verify-thumbprint.html
	// Default: - no thumbprints are allowed. IAM will retrieve and use thumbprint
	// of idenity provider server cerctificate.
	//
	Thumbprints *[]*string `field:"optional" json:"thumbprints" yaml:"thumbprints"`
}

