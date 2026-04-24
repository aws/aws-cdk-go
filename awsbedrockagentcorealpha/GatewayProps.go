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
//   policyEngine := agentcore.NewPolicyEngine(this, jsii.String("MyPolicyEngine"), &PolicyEngineProps{
//   	PolicyEngineName: jsii.String("my_policy_engine"),
//   })
//
//   allowAllPolicy := agentcore.NewPolicy(this, jsii.String("AllowAllPolicy"), &PolicyProps{
//   	PolicyEngine: policyEngine,
//   	PolicyName: jsii.String("allow_all"),
//   	Statement: agentcore.PolicyStatement_Permit().ForAllPrincipals().OnAllActions().OnResource(jsii.String("AgentCore::Gateway"), gateway.GatewayArn),
//   	Description: jsii.String("Allow all actions on specific gateway (development only)"),
//   	ValidationMode: agentcore.PolicyValidationMode_IGNORE_ALL_FINDINGS(),
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
	// The policy engine configuration for this gateway.
	//
	// When provided, the specified policy engine will be associated with this gateway.
	// All agent requests through this gateway will be evaluated against the Cedar policies
	// defined in the policy engine.
	// Default: - No policy engine (requests are not subject to Cedar policy authorization).
	//
	// Experimental.
	PolicyEngineConfiguration *GatewayPolicyEngineConfig `field:"optional" json:"policyEngineConfiguration" yaml:"policyEngineConfiguration"`
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

