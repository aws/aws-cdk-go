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
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   myLambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_22_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String(`
//   	    exports.handler = async (event) => ({ statusCode: 200 });
//   	  `)),
//   })
//
//   myToolSchema := agentcore.ToolSchema_FromInline([]ToolDefinition{
//   	&ToolDefinition{
//   		Name: jsii.String("my_tool"),
//   		Description: jsii.String("My custom tool"),
//   		InputSchema: &SchemaDefinition{
//   			Type: agentcore.SchemaDefinitionType_OBJECT,
//   			Properties: map[string]interface{}{
//   			},
//   		},
//   	},
//   })
//
//   // Create a custom Lambda configuration
//   customConfig := agentcore.LambdaTargetConfiguration_Create(myLambdaFunction, myToolSchema)
//
//   // Use the GatewayTarget constructor directly
//   target := agentcore.NewGatewayTarget(this, jsii.String("AdvancedTarget"), &GatewayTargetProps{
//   	Gateway: gateway,
//   	GatewayTargetName: jsii.String("advanced-target"),
//   	TargetConfiguration: customConfig,
//   	 // Manually created configuration
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		agentcore.GatewayCredentialProvider_FromIamRole(),
//   	},
//   })
//
// Experimental.
type LambdaTargetConfiguration interface {
	McpTargetConfiguration
	// The Lambda function that implements the MCP server logic.
	// Experimental.
	LambdaFunction() awslambda.IFunction
	// The target type.
	// Experimental.
	TargetType() McpTargetType
	// The tool schema that defines the available tools.
	// Experimental.
	ToolSchema() ToolSchema
	// Binds this configuration to a construct scope Sets up necessary permissions for the gateway to invoke the Lambda function.
	// Experimental.
	Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig
	// Renders the MCP-specific configuration.
	// Experimental.
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


// Experimental.
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

// Experimental.
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
// Experimental.
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

