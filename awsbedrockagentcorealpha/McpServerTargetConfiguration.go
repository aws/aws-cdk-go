package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration for MCP Server-based targets.
//
// MCP (Model Context Protocol) servers provide tools, data access, and custom
// functions for AI agents. When you configure an MCP server as a gateway target,
// the gateway automatically discovers and indexes available tools through
// synchronization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   mcpServerTargetConfiguration := bedrock_agentcore_alpha.NewMcpServerTargetConfiguration(jsii.String("endpoint"))
//
// Experimental.
type McpServerTargetConfiguration interface {
	McpTargetConfiguration
	// The HTTPS endpoint URL of the MCP server.
	// Experimental.
	Endpoint() *string
	// The target type.
	// Experimental.
	TargetType() McpTargetType
	// Binds this configuration to a construct scope No additional permissions are needed for MCP server targets.
	// Experimental.
	Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig
	// Renders the MCP-specific configuration.
	// Experimental.
	RenderMcpConfiguration() interface{}
}

// The jsii proxy struct for McpServerTargetConfiguration
type jsiiProxy_McpServerTargetConfiguration struct {
	jsiiProxy_McpTargetConfiguration
}

func (j *jsiiProxy_McpServerTargetConfiguration) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_McpServerTargetConfiguration) TargetType() McpTargetType {
	var returns McpTargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}


// Experimental.
func NewMcpServerTargetConfiguration(endpoint *string) McpServerTargetConfiguration {
	_init_.Initialize()

	if err := validateNewMcpServerTargetConfigurationParameters(endpoint); err != nil {
		panic(err)
	}
	j := jsiiProxy_McpServerTargetConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpServerTargetConfiguration",
		[]interface{}{endpoint},
		&j,
	)

	return &j
}

// Experimental.
func NewMcpServerTargetConfiguration_Override(m McpServerTargetConfiguration, endpoint *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpServerTargetConfiguration",
		[]interface{}{endpoint},
		m,
	)
}

// Create an MCP server target configuration.
//
// Returns: A new McpServerTargetConfiguration instance.
// Experimental.
func McpServerTargetConfiguration_Create(endpoint *string) McpServerTargetConfiguration {
	_init_.Initialize()

	if err := validateMcpServerTargetConfiguration_CreateParameters(endpoint); err != nil {
		panic(err)
	}
	var returns McpServerTargetConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpServerTargetConfiguration",
		"create",
		[]interface{}{endpoint},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_McpServerTargetConfiguration) Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig {
	if err := m.validateBindParameters(scope, gateway); err != nil {
		panic(err)
	}
	var returns *TargetConfigurationConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{scope, gateway},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_McpServerTargetConfiguration) RenderMcpConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderMcpConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

