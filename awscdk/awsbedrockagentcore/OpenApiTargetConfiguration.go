package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
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
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var apiSchema ApiSchema
//
//   openApiTargetConfiguration := awscdk.Aws_bedrockagentcore.NewOpenApiTargetConfiguration(apiSchema, jsii.Boolean(false))
//
type OpenApiTargetConfiguration interface {
	McpTargetConfiguration
	// The OpenAPI schema that defines the API.
	ApiSchema() ApiSchema
	// The target type.
	TargetType() McpTargetType
	// Binds this configuration to a construct scope Sets up necessary permissions for the gateway to access the API schema.
	Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig
	// Renders the MCP-specific configuration.
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


func NewOpenApiTargetConfiguration(apiSchema ApiSchema, validateSchema *bool) OpenApiTargetConfiguration {
	_init_.Initialize()

	if err := validateNewOpenApiTargetConfigurationParameters(apiSchema); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenApiTargetConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.OpenApiTargetConfiguration",
		[]interface{}{apiSchema, validateSchema},
		&j,
	)

	return &j
}

func NewOpenApiTargetConfiguration_Override(o OpenApiTargetConfiguration, apiSchema ApiSchema, validateSchema *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.OpenApiTargetConfiguration",
		[]interface{}{apiSchema, validateSchema},
		o,
	)
}

// Create an OpenAPI target configuration.
//
// Returns: A new OpenApiTargetConfiguration instance.
func OpenApiTargetConfiguration_Create(apiSchema ApiSchema, validateSchema *bool) OpenApiTargetConfiguration {
	_init_.Initialize()

	if err := validateOpenApiTargetConfiguration_CreateParameters(apiSchema); err != nil {
		panic(err)
	}
	var returns OpenApiTargetConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.OpenApiTargetConfiguration",
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

