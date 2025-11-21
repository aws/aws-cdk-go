package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration for OpenAPI-based MCP targets.
//
// This configuration exposes an OpenAPI/REST API as MCP tools,
// allowing the gateway to transform API operations into tool calls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var apiSchema ApiSchema
//
//   openApiTargetConfiguration := bedrock_agentcore_alpha.NewOpenApiTargetConfiguration(apiSchema, jsii.Boolean(false))
//
// Experimental.
type OpenApiTargetConfiguration interface {
	McpTargetConfiguration
	// The OpenAPI schema that defines the API.
	// Experimental.
	ApiSchema() ApiSchema
	// The target type.
	// Experimental.
	TargetType() McpTargetType
	// Binds this configuration to a construct scope Sets up necessary permissions for the gateway to access the API schema.
	// Experimental.
	Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig
	// Renders the MCP-specific configuration.
	// Experimental.
	RenderMcpConfiguration() interface{}
}

// The jsii proxy struct for OpenApiTargetConfiguration
type jsiiProxy_OpenApiTargetConfiguration struct {
	jsiiProxy_McpTargetConfiguration
}

func (j *jsiiProxy_OpenApiTargetConfiguration) ApiSchema() ApiSchema {
	var returns ApiSchema
	_jsii_.Get(
		j,
		"apiSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenApiTargetConfiguration) TargetType() McpTargetType {
	var returns McpTargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}


// Experimental.
func NewOpenApiTargetConfiguration(apiSchema ApiSchema, validateSchema *bool) OpenApiTargetConfiguration {
	_init_.Initialize()

	if err := validateNewOpenApiTargetConfigurationParameters(apiSchema); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenApiTargetConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OpenApiTargetConfiguration",
		[]interface{}{apiSchema, validateSchema},
		&j,
	)

	return &j
}

// Experimental.
func NewOpenApiTargetConfiguration_Override(o OpenApiTargetConfiguration, apiSchema ApiSchema, validateSchema *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OpenApiTargetConfiguration",
		[]interface{}{apiSchema, validateSchema},
		o,
	)
}

// Create an OpenAPI target configuration.
//
// Returns: A new OpenApiTargetConfiguration instance.
// Experimental.
func OpenApiTargetConfiguration_Create(apiSchema ApiSchema, validateSchema *bool) OpenApiTargetConfiguration {
	_init_.Initialize()

	if err := validateOpenApiTargetConfiguration_CreateParameters(apiSchema); err != nil {
		panic(err)
	}
	var returns OpenApiTargetConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.OpenApiTargetConfiguration",
		"create",
		[]interface{}{apiSchema, validateSchema},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenApiTargetConfiguration) Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig {
	if err := o.validateBindParameters(scope, gateway); err != nil {
		panic(err)
	}
	var returns *TargetConfigurationConfig

	_jsii_.Invoke(
		o,
		"bind",
		[]interface{}{scope, gateway},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenApiTargetConfiguration) RenderMcpConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"renderMcpConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

