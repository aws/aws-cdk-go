package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MCP (Model Context Protocol) configuration implementation.
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
//   		AllowedScopes: []*string{
//   			jsii.String("read"),
//   			jsii.String("write"),
//   		},
//   	}),
//   	KmsKey: encryptionKey,
//   	ExceptionLevel: agentcore.GatewayExceptionLevel_DEBUG,
//   })
//
// Experimental.
type McpProtocolConfiguration interface {
	IGatewayProtocolConfig
	// Instructions for using the MCP gateway.
	// Experimental.
	Instructions() *string
	// The protocol type.
	// Experimental.
	ProtocolType() *string
	// The search type for the MCP gateway.
	// Experimental.
	SearchType() *string
	// The supported MCP protocol versions.
	// Experimental.
	SupportedVersions() *[]MCPProtocolVersion
}

// The jsii proxy struct for McpProtocolConfiguration
type jsiiProxy_McpProtocolConfiguration struct {
	jsiiProxy_IGatewayProtocolConfig
}

func (j *jsiiProxy_McpProtocolConfiguration) Instructions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_McpProtocolConfiguration) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_McpProtocolConfiguration) SearchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_McpProtocolConfiguration) SupportedVersions() *[]MCPProtocolVersion {
	var returns *[]MCPProtocolVersion
	_jsii_.Get(
		j,
		"supportedVersions",
		&returns,
	)
	return returns
}


// Experimental.
func NewMcpProtocolConfiguration(props *McpConfiguration) McpProtocolConfiguration {
	_init_.Initialize()

	if err := validateNewMcpProtocolConfigurationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_McpProtocolConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpProtocolConfiguration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewMcpProtocolConfiguration_Override(m McpProtocolConfiguration, props *McpConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.McpProtocolConfiguration",
		[]interface{}{props},
		m,
	)
}

