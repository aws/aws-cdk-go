package awsbedrockagentcore


// Provider ARN, secret ARN, and OAuth scopes for wiring a Token Vault OAuth2 identity into a gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type GatewayOAuth2IdentityBinding struct {
	// OAuth2 credential provider ARN.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// OAuth scopes to request when invoking through the gateway.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Secrets Manager secret ARN for OAuth2 client credentials.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// Optional custom parameters for the OAuth flow.
	// Default: - no custom parameters.
	//
	CustomParameters *map[string]*string `field:"optional" json:"customParameters" yaml:"customParameters"`
}

