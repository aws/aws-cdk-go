package awsbedrockagentcorealpha


// Custom JWT authorizer configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var gatewayCustomClaim GatewayCustomClaim
//
//   customJwtConfiguration := &CustomJwtConfiguration{
//   	DiscoveryUrl: jsii.String("discoveryUrl"),
//
//   	// the properties below are optional
//   	AllowedAudience: []*string{
//   		jsii.String("allowedAudience"),
//   	},
//   	AllowedClients: []*string{
//   		jsii.String("allowedClients"),
//   	},
//   	AllowedScopes: []*string{
//   		jsii.String("allowedScopes"),
//   	},
//   	CustomClaims: []GatewayCustomClaim{
//   		gatewayCustomClaim,
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type CustomJwtConfiguration struct {
	// This URL is used to fetch OpenID Connect configuration or authorization server metadata for validating incoming tokens.
	//
	// Pattern: .+/\.well-known/openid-configuration
	// Required: Yes.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	DiscoveryUrl *string `field:"required" json:"discoveryUrl" yaml:"discoveryUrl"`
	// Represents individual audience values that are validated in the incoming JWT token validation process.
	// Default: - No audience validation.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// Represents individual client IDs that are validated in the incoming JWT token validation process.
	// Default: - No client ID validation.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
	// Represents individual scopes that are validated in the incoming JWT token validation process.
	// Default: - No scope validation.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// Custom claims for additional JWT token validation.
	//
	// Allows you to validate additional fields in JWT tokens beyond the standard audience, client, and scope validations.
	// Default: - No custom claim validation.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CustomClaims *[]GatewayCustomClaim `field:"optional" json:"customClaims" yaml:"customClaims"`
}

