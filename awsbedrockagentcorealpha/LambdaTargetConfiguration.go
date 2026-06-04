package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration for Lambda-based MCP targets.
//
// This configuration wraps a Lambda function as MCP tools,
// allowing the gateway to invoke the function to provide tool capabilities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ Function
//   var toolSchema ToolSchema
//
//   lambdaTargetConfiguration := bedrock_agentcore_alpha.NewLambdaTargetConfiguration(function_, toolSchema)
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type LambdaTargetConfiguration interface {
	McpTargetConfiguration
	// The Lambda function that implements the MCP server logic.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	LambdaFunction() awslambda.IFunction
	// The target type.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	TargetType() McpTargetType
	// The tool schema that defines the available tools.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	ToolSchema() ToolSchema
	// Binds this configuration to a construct scope Sets up necessary permissions for the gateway to invoke the Lambda function.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig
	// Renders the MCP-specific configuration.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	RenderMcpConfiguration() interface{}
}

// The jsii proxy struct for LambdaTargetConfiguration
type jsiiProxy_LambdaTargetConfiguration struct {
	jsiiProxy_McpTargetConfiguration
}

func (j *jsiiProxy_LambdaTargetConfiguration) LambdaFunction() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaTargetConfiguration) TargetType() McpTargetType {
	var returns McpTargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaTargetConfiguration) ToolSchema() ToolSchema {
	var returns ToolSchema
	_jsii_.Get(
		j,
		"toolSchema",
		&returns,
	)
	return returns
}


// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func NewLambdaTargetConfiguration(lambdaFunction awslambda.IFunction, toolSchema ToolSchema) LambdaTargetConfiguration {
	_init_.Initialize()

	if err := validateNewLambdaTargetConfigurationParameters(lambdaFunction, toolSchema); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaTargetConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LambdaTargetConfiguration",
		[]interface{}{lambdaFunction, toolSchema},
		&j,
	)

	return &j
}

// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func NewLambdaTargetConfiguration_Override(l LambdaTargetConfiguration, lambdaFunction awslambda.IFunction, toolSchema ToolSchema) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LambdaTargetConfiguration",
		[]interface{}{lambdaFunction, toolSchema},
		l,
	)
}

// Create a Lambda target configuration.
//
// Returns: A new LambdaTargetConfiguration instance.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func LambdaTargetConfiguration_Create(lambdaFunction awslambda.IFunction, toolSchema ToolSchema) LambdaTargetConfiguration {
	_init_.Initialize()

	if err := validateLambdaTargetConfiguration_CreateParameters(lambdaFunction, toolSchema); err != nil {
		panic(err)
	}
	var returns LambdaTargetConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.LambdaTargetConfiguration",
		"create",
		[]interface{}{lambdaFunction, toolSchema},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaTargetConfiguration) Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig {
	if err := l.validateBindParameters(scope, gateway); err != nil {
		panic(err)
	}
	var returns *TargetConfigurationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, gateway},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaTargetConfiguration) RenderMcpConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderMcpConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

