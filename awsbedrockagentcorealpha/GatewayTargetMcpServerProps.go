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
// Experimental.
type GatewayTargetMcpServerProps struct {
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Default: - auto generate.
	//
	// Experimental.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// Credential providers for authentication.
	// Experimental.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"required" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// The HTTPS endpoint URL of the MCP server.
	//
	// The endpoint must:
	// - Use HTTPS protocol
	// - Be properly URL-encoded
	// - Point to an MCP server that implements tool capabilities.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The gateway this target belongs to.
	// Experimental.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
}

