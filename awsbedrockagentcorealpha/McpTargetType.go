package awsbedrockagentcorealpha


// MCP target types.
// Experimental.
type McpTargetType string

const (
	// OpenAPI schema target type.
	// Experimental.
	McpTargetType_OPENAPI_SCHEMA McpTargetType = "OPENAPI_SCHEMA"
	// Smithy model target type.
	// Experimental.
	McpTargetType_SMITHY_MODEL McpTargetType = "SMITHY_MODEL"
	// Lambda function target type.
	// Experimental.
	McpTargetType_LAMBDA McpTargetType = "LAMBDA"
	// MCP server target type.
	// Experimental.
	McpTargetType_MCP_SERVER McpTargetType = "MCP_SERVER"
)

