package awsbedrockagentcorealpha


// Protocol configuration for Agent Runtime.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-protocolconfiguration
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type ProtocolType string

const (
	// Model Context Protocol.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ProtocolType_MCP ProtocolType = "MCP"
	// HTTP protocol.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ProtocolType_HTTP ProtocolType = "HTTP"
	// A2A protocol.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ProtocolType_A2A ProtocolType = "A2A"
	// Agent User Interaction (AGUI) protocol.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ProtocolType_AGUI ProtocolType = "AGUI"
)

