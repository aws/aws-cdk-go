package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Abstract base class for MCP target configurations Provides common functionality for all MCP target types.
type McpTargetConfiguration interface {
	ITargetConfiguration
	// The target type.
	TargetType() McpTargetType
	// Binds the configuration to a construct scope Sets up permissions and dependencies.
	Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig
	// Renders the MCP-specific configuration.
	RenderMcpConfiguration() interface{}
}

// The jsii proxy struct for McpTargetConfiguration
type jsiiProxy_McpTargetConfiguration struct {
	jsiiProxy_ITargetConfiguration
}

func (j *jsiiProxy_McpTargetConfiguration) TargetType() McpTargetType {
	var returns McpTargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}


func NewMcpTargetConfiguration_Override(m McpTargetConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.McpTargetConfiguration",
		nil, // no parameters
		m,
	)
}

func (m *jsiiProxy_McpTargetConfiguration) Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig {
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

func (m *jsiiProxy_McpTargetConfiguration) RenderMcpConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderMcpConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

