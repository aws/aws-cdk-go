package awsbedrockagentcorealpha


// Protocol configuration for Agent Runtime.
// Experimental.
type ProtocolType string

const (
	// Model Context Protocol.
	// Experimental.
	ProtocolType_MCP ProtocolType = "MCP"
	// HTTP protocol.
	// Experimental.
	ProtocolType_HTTP ProtocolType = "HTTP"
)

