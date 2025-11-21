package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating Gateway Authorizers.
//
// Example:
//   // Create a KMS key for encryption
//   encryptionKey := kms.NewKey(this, jsii.String("GatewayEncryptionKey"), &KeyProps{
//   	EnableKeyRotation: jsii.Boolean(true),
//   	Description: jsii.String("KMS key for gateway encryption"),
//   })
//
//   // Create gateway with KMS encryption
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-encrypted-gateway"),
//   	Description: jsii.String("Gateway with KMS encryption"),
//   	ProtocolConfiguration: agentcore.NewMcpProtocolConfiguration(&McpConfiguration{
//   		Instructions: jsii.String("Use this gateway to connect to external MCP tools"),
//   		SearchType: agentcore.McpGatewaySearchType_SEMANTIC,
//   		SupportedVersions: []MCPProtocolVersion{
//   			agentcore.MCPProtocolVersion_MCP_2025_03_26,
//   		},
//   	}),
//   	AuthorizerConfiguration: agentcore.GatewayAuthorizer_UsingCustomJwt(&CustomJwtConfiguration{
//   		DiscoveryUrl: jsii.String("https://auth.example.com/.well-known/openid-configuration"),
//   		AllowedAudience: []*string{
//   			jsii.String("my-app"),
//   		},
//   		AllowedClients: []*string{
//   			jsii.String("my-client-id"),
//   		},
//   	}),
//   	KmsKey: encryptionKey,
//   	ExceptionLevel: agentcore.GatewayExceptionLevel_DEBUG,
//   })
//
// Experimental.
type GatewayAuthorizer interface {
}

// The jsii proxy struct for GatewayAuthorizer
type jsiiProxy_GatewayAuthorizer struct {
	_ byte // padding
}

// Experimental.
func NewGatewayAuthorizer_Override(g GatewayAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAuthorizer",
		nil, // no parameters
		g,
	)
}

// AWS IAM authorizer instance.
// Experimental.
func GatewayAuthorizer_UsingAwsIam() IGatewayAuthorizerConfig {
	_init_.Initialize()

	var returns IGatewayAuthorizerConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAuthorizer",
		"usingAwsIam",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create a JWT authorizer from Cognito User Pool.
//
// Returns: CustomJwtAuthorizer configured for Cognito.
// Experimental.
func GatewayAuthorizer_UsingCognito(props *CognitoAuthorizerProps) IGatewayAuthorizerConfig {
	_init_.Initialize()

	if err := validateGatewayAuthorizer_UsingCognitoParameters(props); err != nil {
		panic(err)
	}
	var returns IGatewayAuthorizerConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAuthorizer",
		"usingCognito",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a custom JWT authorizer.
//
// Returns: IGatewayAuthorizerConfig configured for custom JWT.
// Experimental.
func GatewayAuthorizer_UsingCustomJwt(configuration *CustomJwtConfiguration) IGatewayAuthorizerConfig {
	_init_.Initialize()

	if err := validateGatewayAuthorizer_UsingCustomJwtParameters(configuration); err != nil {
		panic(err)
	}
	var returns IGatewayAuthorizerConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAuthorizer",
		"usingCustomJwt",
		[]interface{}{configuration},
		&returns,
	)

	return returns
}

