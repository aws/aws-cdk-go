package awsbedrockagentcorealpha


// Search types supported by MCP gateway.
// Experimental.
type McpGatewaySearchType string

const (
	// Semantic search type.
	//
	// When semantic search is enabled, your gateway can search the tools via
	// the gateway SDK based off of a natural language phrase.
	// Experimental.
	McpGatewaySearchType_SEMANTIC McpGatewaySearchType = "SEMANTIC"
)

