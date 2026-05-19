package awsbedrockagentcore


// Optional tenant OAuth endpoints for IdPs that use CloudFormation `IncludedOauth2ProviderConfig` with issuer and/or endpoints per the IdP’s outbound documentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   includedOauth2TenantEndpoints := &IncludedOauth2TenantEndpoints{
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	Issuer: jsii.String("issuer"),
//   	TokenEndpoint: jsii.String("tokenEndpoint"),
//   }
//
type IncludedOauth2TenantEndpoints struct {
	// OAuth2 authorization endpoint for your tenant.
	// Default: - not specified; use when your IdP requires an explicit endpoint.
	//
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Token issuer URL for your tenant (often the IdP base or issuer URI).
	// Default: - not specified; use when your IdP requires an explicit issuer.
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// OAuth2 token endpoint for your tenant.
	// Default: - not specified; use when your IdP requires an explicit endpoint.
	//
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
}

