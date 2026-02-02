package awsbedrockagentcorealpha


// MCP protocol configuration The configuration for the Model Context Protocol (MCP).
//
// This protocol enables communication between Amazon Bedrock Agent and external tools.
//
// Example:
//   // Create a KMS key for encryption
//   encryptionKey := kms.NewKey(this, jsii.String("GatewayEncryptionKey"), &KeyProps{
//   	EnableKeyRotation: jsii.Boolean(true),
//   	Description: jsii.String("KMS key for gateway encryption"),
//   })
//
//   // Create gateway with KMS encryption
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-encrypted-gateway"),
//   	Description: jsii.String("Gateway with KMS encryption"),
//   	ProtocolConfiguration: agentcore.NewMcpProtocolConfiguration(&McpConfiguration{
//   		Instructions: jsii.String("Use this gateway to connect to external MCP tools"),
//   		SearchType: agentcore.McpGatewaySearchType_SEMANTIC,
//   		SupportedVersions: []MCPProtocolVersion{
//   			agentcore.MCPProtocolVersion_MCP_2025_03_26,
//   		},
//   	}),
//   	AuthorizerConfiguration: agentcore.GatewayAuthorizer_UsingCustomJwt(&CustomJwtConfiguration{
//   		DiscoveryUrl: jsii.String("https://auth.example.com/.well-known/openid-configuration"),
//   		AllowedAudience: []*string{
//   			jsii.String("my-app"),
//   		},
//   		AllowedClients: []*string{
//   			jsii.String("my-client-id"),
//   		},
//   		AllowedScopes: []*string{
//   			jsii.String("read"),
//   			jsii.String("write"),
//   		},
//   	}),
//   	KmsKey: encryptionKey,
//   	ExceptionLevel: agentcore.GatewayExceptionLevel_DEBUG,
//   })
//
// Experimental.
type McpConfiguration struct {
	// The instructions for using the Model Context Protocol gateway.
	//
	// These instructions provide guidance on how to interact with the gateway.
	// Default: - No instructions provided.
	//
	// Experimental.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// The search type for the Model Context Protocol gateway.
	//
	// This field specifies how the gateway handles search operations.
	// Default: - No search type specified.
	//
	// Experimental.
	SearchType McpGatewaySearchType `field:"optional" json:"searchType" yaml:"searchType"`
	// The supported versions of the Model Context Protocol.
	//
	// This field specifies which versions of the protocol the gateway can use.
	// Default: - No specific versions specified.
	//
	// Experimental.
	SupportedVersions *[]MCPProtocolVersion `field:"optional" json:"supportedVersions" yaml:"supportedVersions"`
}

