package awsbedrockagentcorealpha


// MCP protocol versions.
//
// The Model Context Protocol uses string-based version identifiers following the format YYYY-MM-DD,
// to indicate the last date backwards incompatible changes were made.
// Versions are available at https://github.com/modelcontextprotocol/modelcontextprotocol/releases
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type MCPProtocolVersion string

const (
	// The latest version of the MCP protocol.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MCPProtocolVersion_MCP_2025_06_18 MCPProtocolVersion = "MCP_2025_06_18"
	// MCP version 2025-03-26.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MCPProtocolVersion_MCP_2025_03_26 MCPProtocolVersion = "MCP_2025_03_26"
)

