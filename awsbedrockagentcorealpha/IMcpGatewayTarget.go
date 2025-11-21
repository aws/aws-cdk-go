package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for MCP gateway targets.
//
// Extends the base gateway target interface with MCP-specific properties.
// MCP targets expose tools using the Model Context Protocol.
// Experimental.
type IMcpGatewayTarget interface {
	IGatewayTarget
	// The type of target (Lambda, OpenAPI, or Smithy).
	// Experimental.
	TargetType() McpTargetType
}

// The jsii proxy for IMcpGatewayTarget
type jsiiProxy_IMcpGatewayTarget struct {
	jsiiProxy_IGatewayTarget
}

func (j *jsiiProxy_IMcpGatewayTarget) TargetType() McpTargetType {
	var returns McpTargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

