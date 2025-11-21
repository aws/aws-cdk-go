package awsbedrockagentcorealpha


// MCP protocol versions.
//
// The Model Context Protocol uses string-based version identifiers following the format YYYY-MM-DD,
// to indicate the last date backwards incompatible changes were made.
// Versions are available at https://github.com/modelcontextprotocol/modelcontextprotocol/releases
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
//   	}),
//   	KmsKey: encryptionKey,
//   	ExceptionLevel: agentcore.GatewayExceptionLevel_DEBUG,
//   })
//
// Experimental.
type MCPProtocolVersion string

const (
	// The latest version of the MCP protocol.
	// Experimental.
	MCPProtocolVersion_MCP_2025_06_18 MCPProtocolVersion = "MCP_2025_06_18"
	// MCP version 2025-03-26.
	// Experimental.
	MCPProtocolVersion_MCP_2025_03_26 MCPProtocolVersion = "MCP_2025_03_26"
)

