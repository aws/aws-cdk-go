package awsbedrockagentcorealpha


// Static OAuth2 authorization server metadata for custom credential providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   oAuth2AuthorizationServerMetadata := &OAuth2AuthorizationServerMetadata{
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	Issuer: jsii.String("issuer"),
//   	TokenEndpoint: jsii.String("tokenEndpoint"),
//
//   	// the properties below are optional
//   	ResponseTypes: []*string{
//   		jsii.String("responseTypes"),
//   	},
//   }
//
// See: https://www.rfc-editor.org/rfc/rfc8414
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type OAuth2AuthorizationServerMetadata struct {
	// The authorization endpoint URL.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The issuer URL for the OAuth2 authorization server.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The token endpoint URL.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The supported response types.
	// Default: - not specified.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ResponseTypes *[]*string `field:"optional" json:"responseTypes" yaml:"responseTypes"`
}

