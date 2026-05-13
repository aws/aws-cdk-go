package awsbedrockagentcorealpha


// Protocol configuration for Agent Runtime.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-bedrockagentcore-runtime.html#cfn-bedrockagentcore-runtime-protocolconfiguration
//
// Experimental.
type ProtocolType string

const (
	// Model Context Protocol.
	// Experimental.
	ProtocolType_MCP ProtocolType = "MCP"
	// HTTP protocol.
	// Experimental.
	ProtocolType_HTTP ProtocolType = "HTTP"
	// A2A protocol.
	// Experimental.
	ProtocolType_A2A ProtocolType = "A2A"
	// Agent User Interaction (AGUI) protocol.
	// Experimental.
	ProtocolType_AGUI ProtocolType = "AGUI"
)

