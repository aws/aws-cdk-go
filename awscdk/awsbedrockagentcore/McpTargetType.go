package awsbedrockagentcore


// MCP target types.
type McpTargetType string

const (
	// OpenAPI schema target type.
	McpTargetType_OPENAPI_SCHEMA McpTargetType = "OPENAPI_SCHEMA"
	// Smithy model target type.
	McpTargetType_SMITHY_MODEL McpTargetType = "SMITHY_MODEL"
	// Lambda function target type.
	McpTargetType_LAMBDA McpTargetType = "LAMBDA"
	// MCP server target type.
	McpTargetType_MCP_SERVER McpTargetType = "MCP_SERVER"
	// API Gateway target type.
	McpTargetType_API_GATEWAY McpTargetType = "API_GATEWAY"
)

