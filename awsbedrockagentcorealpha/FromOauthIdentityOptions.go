package awsbedrockagentcorealpha


// OAuth scopes (and optional custom parameters) when binding an {@link IOAuth2CredentialProvider} to a gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   fromOauthIdentityOptions := &FromOauthIdentityOptions{
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//
//   	// the properties below are optional
//   	CustomParameters: map[string]*string{
//   		"customParametersKey": jsii.String("customParameters"),
//   	},
//   }
//
// Experimental.
type FromOauthIdentityOptions struct {
	// OAuth scopes the gateway should request for this target.
	// Experimental.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Additional OAuth parameters for the provider.
	// Default: - none.
	//
	// Experimental.
	CustomParameters *map[string]*string `field:"optional" json:"customParameters" yaml:"customParameters"`
}

