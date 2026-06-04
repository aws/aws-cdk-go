package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MCP (Model Context Protocol) configuration implementation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   mcpProtocolConfiguration := bedrock_agentcore_alpha.NewMcpProtocolConfiguration(&McpConfiguration{
//   	Instructions: jsii.String("instructions"),
//   	SearchType: bedrock_agentcore_alpha.McpGatewaySearchType_SEMANTIC,
//   	SupportedVersions: []MCPProtocolVersion{
//   		bedrock_agentcore_alpha.MCPProtocolVersion_MCP_2025_06_18,
//   	},
//   })
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type McpProtocolConfiguration interface {
	IGatewayProtocolConfig
	// Instructions for using the MCP gateway.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Instructions() *string
	// The protocol type.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ProtocolType() *string
	// The search type for the MCP gateway.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	SearchType() *string
	// The supported MCP protocol versions.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	SupportedVersions() *[]MCPProtocolVersion
}

// The jsii proxy struct for McpProtocolConfiguration
type jsiiProxy_McpProtocolConfiguration struct {
	jsiiProxy_IGatewayProtocolConfig
}

func (j *jsiiProxy_McpProtocolConfiguration) Instructions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_McpProtocolConfiguration) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_McpProtocolConfiguration) SearchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_McpProtocolConfiguration) SupportedVersions() *[]MCPProtocolVersion {
	var returns *[]MCPProtocolVersion
	_jsii_.Get(
		j,
		"supportedVersions",
		&returns,
	)
	return returns
}


// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func NewMcpProtocolConfiguration(props *McpConfiguration) McpProtocolConfiguration {
	_init_.Initialize()

	if err := validateNewMcpProtocolConfigurationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_McpProtocolConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpProtocolConfiguration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func NewMcpProtocolConfiguration_Override(m McpProtocolConfiguration, props *McpConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpProtocolConfiguration",
		[]interface{}{props},
		m,
	)
}

