package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for instantiating Gateway Protocols.
type GatewayProtocol interface {
}

// The jsii proxy struct for GatewayProtocol
type jsiiProxy_GatewayProtocol struct {
	_ byte // padding
}

func NewGatewayProtocol_Override(g GatewayProtocol) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayProtocol",
		nil, // no parameters
		g,
	)
}

// Create an MCP protocol configuration.
//
// Returns: IGatewayProtocolConfig configured for MCP.
func GatewayProtocol_Mcp(props *McpConfiguration) IGatewayProtocolConfig {
	_init_.Initialize()

	if err := validateGatewayProtocol_McpParameters(props); err != nil {
		panic(err)
	}
	var returns IGatewayProtocolConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayProtocol",
		"mcp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

