package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for instantiating Gateway Protocols.
// Experimental.
type GatewayProtocol interface {
}

// The jsii proxy struct for GatewayProtocol
type jsiiProxy_GatewayProtocol struct {
	_ byte // padding
}

// Experimental.
func NewGatewayProtocol_Override(g GatewayProtocol) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayProtocol",
		nil, // no parameters
		g,
	)
}

// Create an MCP protocol configuration.
//
// Returns: IGatewayProtocolConfig configured for MCP.
// Experimental.
func GatewayProtocol_Mcp(props *McpConfiguration) IGatewayProtocolConfig {
	_init_.Initialize()

	if err := validateGatewayProtocol_McpParameters(props); err != nil {
		panic(err)
	}
	var returns IGatewayProtocolConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayProtocol",
		"mcp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

