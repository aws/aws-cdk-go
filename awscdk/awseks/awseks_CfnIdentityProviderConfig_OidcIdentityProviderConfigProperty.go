package awseks


// An object that represents the configuration for an OpenID Connect (OIDC) identity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oidcIdentityProviderConfigProperty := &oidcIdentityProviderConfigProperty{
//   	clientId: jsii.String("clientId"),
//   	issuerUrl: jsii.String("issuerUrl"),
//
//   	// the properties below are optional
//   	groupsClaim: jsii.String("groupsClaim"),
//   	groupsPrefix: jsii.String("groupsPrefix"),
//   	requiredClaims: []interface{}{
//   		&requiredClaimProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	usernameClaim: jsii.String("usernameClaim"),
//   	usernamePrefix: jsii.String("usernamePrefix"),
//   }
//
type CfnIdentityProviderConfig_OidcIdentityProviderConfigProperty struct {
	// This is also known as *audience* .
	//
	// The ID of the client application that makes authentication requests to the OIDC identity provider.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The URL of the OIDC identity provider that allows the API server to discover public signing keys for verifying tokens.
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// The JSON web token (JWT) claim that the provider uses to return your groups.
	GroupsClaim *string `field:"optional" json:"groupsClaim" yaml:"groupsClaim"`
	// The prefix that is prepended to group claims to prevent clashes with existing names (such as `system:` groups).
	//
	// For example, the value `oidc:` creates group names like `oidc:engineering` and `oidc:infra` . The prefix can't contain `system:`
	GroupsPrefix *string `field:"optional" json:"groupsPrefix" yaml:"groupsPrefix"`
	// The key-value pairs that describe required claims in the identity token.
	//
	// If set, each claim is verified to be present in the token with a matching value.
	RequiredClaims interface{} `field:"optional" json:"requiredClaims" yaml:"requiredClaims"`
	// The JSON Web token (JWT) claim that is used as the username.
	UsernameClaim *string `field:"optional" json:"usernameClaim" yaml:"usernameClaim"`
	// The prefix that is prepended to username claims to prevent clashes with existing names.
	//
	// The prefix can't contain `system:`.
	UsernamePrefix *string `field:"optional" json:"usernamePrefix" yaml:"usernamePrefix"`
}

