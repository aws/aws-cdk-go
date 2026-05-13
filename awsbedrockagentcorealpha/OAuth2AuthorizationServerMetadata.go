package awsbedrockagentcorealpha


// Static OAuth2 authorization server metadata for custom credential providers.
//
// Example:
//   agentcore.OAuth2CredentialProvider_UsingCustom(this, jsii.String("CustomOAuthMeta"), &CustomOAuth2CredentialProviderProps{
//   	ClientId: jsii.String("your-client-id"),
//   	ClientSecret: cdk.SecretValue_UnsafePlainText(jsii.String("your-client-secret")),
//   	AuthorizationServerMetadata: &OAuth2AuthorizationServerMetadata{
//   		Issuer: jsii.String("https://idp.example.com"),
//   		AuthorizationEndpoint: jsii.String("https://idp.example.com/oauth2/authorize"),
//   		TokenEndpoint: jsii.String("https://idp.example.com/oauth2/token"),
//   	},
//   })
//
// See: https://www.rfc-editor.org/rfc/rfc8414
//
// Experimental.
type OAuth2AuthorizationServerMetadata struct {
	// The authorization endpoint URL.
	// Experimental.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The issuer URL for the OAuth2 authorization server.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The token endpoint URL.
	// Experimental.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The supported response types.
	// Default: - not specified.
	//
	// Experimental.
	ResponseTypes *[]*string `field:"optional" json:"responseTypes" yaml:"responseTypes"`
}

