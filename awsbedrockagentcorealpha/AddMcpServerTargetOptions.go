package awsbedrockagentcorealpha


// Options for adding an MCP Server target to a gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var credentialProviderConfig ICredentialProviderConfig
//
//   addMcpServerTargetOptions := &AddMcpServerTargetOptions{
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		credentialProviderConfig,
//   	},
//   	Endpoint: jsii.String("endpoint"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	GatewayTargetName: jsii.String("gatewayTargetName"),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type AddMcpServerTargetOptions struct {
	// Credential providers for authentication.
	//
	// MCP servers require explicit authentication configuration.
	// OAuth2 is strongly recommended over NoAuth.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"required" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// The HTTPS endpoint URL of the MCP server.
	//
	// The endpoint must:
	// - Use HTTPS protocol
	// - Be properly URL-encoded
	// - Point to an MCP server that implements tool capabilities.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Optional description for the gateway target.
	// Default: - No description.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target Valid characters are a-z, A-Z, 0-9, _ (underscore) and - (hyphen).
	// Default: - auto generate.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
}

