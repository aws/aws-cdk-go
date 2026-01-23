package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for defining a Gateway.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_22_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String(`
//   	        exports.handler = async (event) => {
//   	            return {
//   	                statusCode: 200,
//   	                body: JSON.stringify({ message: 'Hello from Lambda!' })
//   	            };
//   	        };
//   	    `)),
//   })
//
//   // Create a gateway target with Lambda and tool schema
//   target := agentcore.GatewayTarget_ForLambda(this, jsii.String("MyLambdaTarget"), &GatewayTargetLambdaProps{
//   	GatewayTargetName: jsii.String("my-lambda-target"),
//   	Description: jsii.String("Target for Lambda function integration"),
//   	Gateway: gateway,
//   	LambdaFunction: lambdaFunction,
//   	ToolSchema: agentcore.ToolSchema_FromLocalAsset(path.join(__dirname, jsii.String("schemas"), jsii.String("my-tool-schema.json"))),
//   })
//
// Experimental.
type GatewayProps struct {
	// The authorizer configuration for the gateway.
	// Default: - A default authorizer will be created using Cognito.
	//
	// Experimental.
	AuthorizerConfiguration IGatewayAuthorizerConfig `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
	// Optional description for the gateway Valid characters are a-z, A-Z, 0-9, _ (underscore), - (hyphen) and spaces The description can have up to 200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The verbosity of exception messages Use DEBUG mode to see granular exception messages from a Gateway.
	// Default: - Exception messages are sanitized for presentation to end users.
	//
	// Experimental.
	ExceptionLevel GatewayExceptionLevel `field:"optional" json:"exceptionLevel" yaml:"exceptionLevel"`
	// The name of the gateway Valid characters are a-z, A-Z, 0-9, _ (underscore) and - (hyphen) The name must be unique within your account.
	// Default: - auto generate.
	//
	// Experimental.
	GatewayName *string `field:"optional" json:"gatewayName" yaml:"gatewayName"`
	// Interceptor configurations for the gateway.
	//
	// Interceptors allow you to run custom code during each gateway invocation:
	// - REQUEST interceptors execute before the gateway calls the target
	// - RESPONSE interceptors execute after the target responds
	//
	// A gateway can have at most one REQUEST interceptor and one RESPONSE interceptor.
	// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/gateway-interceptors.html
	//
	// Default: - No interceptors.
	//
	// Experimental.
	InterceptorConfigurations *[]IInterceptor `field:"optional" json:"interceptorConfigurations" yaml:"interceptorConfigurations"`
	// The AWS KMS key used to encrypt data associated with the gateway.
	// Default: - No encryption.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The protocol configuration for the gateway.
	// Default: - A default protocol configuration will be created using MCP with following params
	// supportedVersions: [MCPProtocolVersion.MCP_2025_03_26],
	// searchType: McpGatewaySearchType.SEMANTIC,
	// instructions: "Default gateway to connect to external MCP tools",.
	//
	// Experimental.
	ProtocolConfiguration IGatewayProtocolConfig `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// The IAM role that provides permissions for the gateway to access AWS services.
	// Default: - A new role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Tags for the gateway A list of key:value pairs of tags to apply to this Gateway resource.
	// Default: - No tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

