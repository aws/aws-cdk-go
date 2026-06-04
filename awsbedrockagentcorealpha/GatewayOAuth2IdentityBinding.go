package awsbedrockagentcorealpha


// Provider ARN, secret ARN, and OAuth scopes for wiring a Token Vault OAuth2 identity into a gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   gatewayOAuth2IdentityBinding := &GatewayOAuth2IdentityBinding{
//   	ProviderArn: jsii.String("providerArn"),
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   	SecretArn: jsii.String("secretArn"),
//
//   	// the properties below are optional
//   	CustomParameters: map[string]*string{
//   		"customParametersKey": jsii.String("customParameters"),
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type GatewayOAuth2IdentityBinding struct {
	// OAuth2 credential provider ARN.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// OAuth scopes to request when invoking through the gateway.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Secrets Manager secret ARN for OAuth2 client credentials.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// Optional custom parameters for the OAuth flow.
	// Default: - no custom parameters.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CustomParameters *map[string]*string `field:"optional" json:"customParameters" yaml:"customParameters"`
}

