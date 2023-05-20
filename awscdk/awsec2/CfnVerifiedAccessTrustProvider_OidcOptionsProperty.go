package awsec2


// Describes the options for an OpenID Connect-compatible user-identity trust provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oidcOptionsProperty := &OidcOptionsProperty{
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	Issuer: jsii.String("issuer"),
//   	Scope: jsii.String("scope"),
//   	TokenEndpoint: jsii.String("tokenEndpoint"),
//   	UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   }
//
type CfnVerifiedAccessTrustProvider_OidcOptionsProperty struct {
	// The OIDC authorization endpoint.
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The client identifier.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The OpenID Connect (OIDC) scope specified.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The OIDC token endpoint.
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The OIDC user info endpoint.
	UserInfoEndpoint *string `field:"optional" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
}

