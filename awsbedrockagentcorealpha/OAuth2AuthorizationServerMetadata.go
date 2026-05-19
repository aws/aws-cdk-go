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

