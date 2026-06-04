package awsbedrockagentcorealpha


// Properties for creating an MCP Server-based Gateway Target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var credentialProviderConfig ICredentialProviderConfig
//   var gateway Gateway
//
//   gatewayTargetMcpServerProps := &GatewayTargetMcpServerProps{
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		credentialProviderConfig,
//   	},
//   	Endpoint: jsii.String("endpoint"),
//   	Gateway: gateway,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	GatewayTargetName: jsii.String("gatewayTargetName"),
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type GatewayTargetMcpServerProps struct {
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Default: - auto generate.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// Credential providers for authentication.
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
	// The gateway this target belongs to.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
}

