package awsappsync


// Configuration for OpenID Connect authorization in AppSync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openIdConnectConfig := &openIdConnectConfig{
//   	oidcProvider: jsii.String("oidcProvider"),
//
//   	// the properties below are optional
//   	clientId: jsii.String("clientId"),
//   	tokenExpiryFromAuth: jsii.Number(123),
//   	tokenExpiryFromIssue: jsii.Number(123),
//   }
//
// Experimental.
type OpenIdConnectConfig struct {
	// The issuer for the OIDC configuration.
	//
	// The issuer returned by discovery must exactly match the value of `iss` in the OIDC token.
	// Experimental.
	OidcProvider *string `field:"required" json:"oidcProvider" yaml:"oidcProvider"`
	// The client identifier of the Relying party at the OpenID identity provider.
	//
	// A regular expression can be specified so AppSync can validate against multiple client identifiers at a time.
	//
	// Example:
	//   -"ABCD|CDEF"
	//
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The number of milliseconds an OIDC token is valid after being authenticated by OIDC provider.
	//
	// `auth_time` claim in OIDC token is required for this validation to work.
	// Experimental.
	TokenExpiryFromAuth *float64 `field:"optional" json:"tokenExpiryFromAuth" yaml:"tokenExpiryFromAuth"`
	// The number of milliseconds an OIDC token is valid after being issued to a user.
	//
	// This validation uses `iat` claim of OIDC token.
	// Experimental.
	TokenExpiryFromIssue *float64 `field:"optional" json:"tokenExpiryFromIssue" yaml:"tokenExpiryFromIssue"`
}

