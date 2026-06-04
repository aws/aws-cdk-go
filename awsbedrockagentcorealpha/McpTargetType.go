package awsbedrockagentcorealpha


// MCP target types.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type McpTargetType string

const (
	// OpenAPI schema target type.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	McpTargetType_OPENAPI_SCHEMA McpTargetType = "OPENAPI_SCHEMA"
	// Smithy model target type.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	McpTargetType_SMITHY_MODEL McpTargetType = "SMITHY_MODEL"
	// Lambda function target type.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	McpTargetType_LAMBDA McpTargetType = "LAMBDA"
	// MCP server target type.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	McpTargetType_MCP_SERVER McpTargetType = "MCP_SERVER"
	// API Gateway target type.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	McpTargetType_API_GATEWAY McpTargetType = "API_GATEWAY"
)

