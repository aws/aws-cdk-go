package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration for Smithy-based MCP targets.
//
// This configuration exposes a Smithy-modeled API as MCP tools,
// allowing the gateway to transform Smithy operations into tool calls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var apiSchema ApiSchema
//
//   smithyTargetConfiguration := bedrock_agentcore_alpha.NewSmithyTargetConfiguration(apiSchema)
//
// Experimental.
type SmithyTargetConfiguration interface {
	McpTargetConfiguration
	// The Smithy model that defines the API.
	// Experimental.
	SmithyModel() ApiSchema
	// The target type.
	// Experimental.
	TargetType() McpTargetType
	// Binds this configuration to a construct scope Sets up necessary permissions for the gateway to access the Smithy model.
	// Experimental.
	Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig
	// Renders the MCP-specific configuration.
	// Experimental.
	RenderMcpConfiguration() interface{}
}

// The jsii proxy struct for SmithyTargetConfiguration
type jsiiProxy_SmithyTargetConfiguration struct {
	jsiiProxy_McpTargetConfiguration
}

func (j *jsiiProxy_SmithyTargetConfiguration) SmithyModel() ApiSchema {
	var returns ApiSchema
	_jsii_.Get(
		j,
		"smithyModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmithyTargetConfiguration) TargetType() McpTargetType {
	var returns McpTargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}


// Experimental.
func NewSmithyTargetConfiguration(smithyModel ApiSchema) SmithyTargetConfiguration {
	_init_.Initialize()

	if err := validateNewSmithyTargetConfigurationParameters(smithyModel); err != nil {
		panic(err)
	}
	j := jsiiProxy_SmithyTargetConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SmithyTargetConfiguration",
		[]interface{}{smithyModel},
		&j,
	)

	return &j
}

// Experimental.
func NewSmithyTargetConfiguration_Override(s SmithyTargetConfiguration, smithyModel ApiSchema) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SmithyTargetConfiguration",
		[]interface{}{smithyModel},
		s,
	)
}

// Create a Smithy target configuration.
//
// Returns: A new SmithyTargetConfiguration instance.
// Experimental.
func SmithyTargetConfiguration_Create(smithyModel ApiSchema) SmithyTargetConfiguration {
	_init_.Initialize()

	if err := validateSmithyTargetConfiguration_CreateParameters(smithyModel); err != nil {
		panic(err)
	}
	var returns SmithyTargetConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.SmithyTargetConfiguration",
		"create",
		[]interface{}{smithyModel},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmithyTargetConfiguration) Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig {
	if err := s.validateBindParameters(scope, gateway); err != nil {
		panic(err)
	}
	var returns *TargetConfigurationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, gateway},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmithyTargetConfiguration) RenderMcpConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderMcpConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

