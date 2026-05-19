package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration for API Gateway-based MCP targets.
//
// This configuration connects your gateway to an Amazon API Gateway REST API stage.
// The gateway translates incoming MCP requests into HTTP requests to your REST API
// and handles response formatting.
//
// Key considerations:
// - API must be in the same account and region as the gateway
// - Only REST APIs are supported (no HTTP or WebSocket APIs)
// - API must use a public endpoint type
// - Methods with both AWS_IAM authorization and API key requirements are not supported
// - Proxy resources (e.g., `/pets/{proxy+}`) are not supported
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var apiGatewayHttpMethod ApiGatewayHttpMethod
//   var restApi RestApi
//
//   apiGatewayTargetConfiguration := awscdk.Aws_bedrockagentcore.NewApiGatewayTargetConfiguration(&ApiGatewayTargetConfigurationProps{
//   	ApiGatewayToolConfiguration: &ApiGatewayToolConfiguration{
//   		ToolFilters: []ApiGatewayToolFilter{
//   			&ApiGatewayToolFilter{
//   				FilterPath: jsii.String("filterPath"),
//   				Methods: []ApiGatewayHttpMethod{
//   					apiGatewayHttpMethod,
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		ToolOverrides: []ApiGatewayToolOverride{
//   			&ApiGatewayToolOverride{
//   				Method: apiGatewayHttpMethod,
//   				Name: jsii.String("name"),
//   				Path: jsii.String("path"),
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   			},
//   		},
//   	},
//   	RestApi: restApi,
//
//   	// the properties below are optional
//   	MetadataConfiguration: &MetadataConfiguration{
//   		AllowedQueryParameters: []*string{
//   			jsii.String("allowedQueryParameters"),
//   		},
//   		AllowedRequestHeaders: []*string{
//   			jsii.String("allowedRequestHeaders"),
//   		},
//   		AllowedResponseHeaders: []*string{
//   			jsii.String("allowedResponseHeaders"),
//   		},
//   	},
//   	Stage: jsii.String("stage"),
//   })
//
type ApiGatewayTargetConfiguration interface {
	McpTargetConfiguration
	// Tool configuration for the API Gateway target.
	ApiGatewayToolConfiguration() *ApiGatewayToolConfiguration
	// Metadata configuration for the API Gateway target.
	MetadataConfiguration() *MetadataConfiguration
	// The ID of the REST API.
	RestApiId() *string
	// The stage name of the REST API.
	Stage() *string
	// The target type.
	TargetType() McpTargetType
	// Binds this configuration to a construct scope Sets up necessary permissions for the gateway to access the API Gateway.
	Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig
	// Renders the MCP-specific configuration.
	RenderMcpConfiguration() interface{}
}

// The jsii proxy struct for ApiGatewayTargetConfiguration
type jsiiProxy_ApiGatewayTargetConfiguration struct {
	jsiiProxy_McpTargetConfiguration
}

func (j *jsiiProxy_ApiGatewayTargetConfiguration) ApiGatewayToolConfiguration() *ApiGatewayToolConfiguration {
	var returns *ApiGatewayToolConfiguration
	_jsii_.Get(
		j,
		"apiGatewayToolConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayTargetConfiguration) MetadataConfiguration() *MetadataConfiguration {
	var returns *MetadataConfiguration
	_jsii_.Get(
		j,
		"metadataConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayTargetConfiguration) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayTargetConfiguration) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayTargetConfiguration) TargetType() McpTargetType {
	var returns McpTargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}


func NewApiGatewayTargetConfiguration(props *ApiGatewayTargetConfigurationProps) ApiGatewayTargetConfiguration {
	_init_.Initialize()

	if err := validateNewApiGatewayTargetConfigurationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayTargetConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayTargetConfiguration",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewApiGatewayTargetConfiguration_Override(a ApiGatewayTargetConfiguration, props *ApiGatewayTargetConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayTargetConfiguration",
		[]interface{}{props},
		a,
	)
}

// Create an API Gateway target configuration.
//
// Returns: A new ApiGatewayTargetConfiguration instance.
func ApiGatewayTargetConfiguration_Create(props *ApiGatewayTargetConfigurationProps) ApiGatewayTargetConfiguration {
	_init_.Initialize()

	if err := validateApiGatewayTargetConfiguration_CreateParameters(props); err != nil {
		panic(err)
	}
	var returns ApiGatewayTargetConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.ApiGatewayTargetConfiguration",
		"create",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayTargetConfiguration) Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig {
	if err := a.validateBindParameters(scope, gateway); err != nil {
		panic(err)
	}
	var returns *TargetConfigurationConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, gateway},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayTargetConfiguration) RenderMcpConfiguration() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderMcpConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

