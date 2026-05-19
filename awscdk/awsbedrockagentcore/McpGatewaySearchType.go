package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Search types supported by MCP gateway.
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
//   		SearchType: agentcore.McpGatewaySearchType_SEMANTIC(),
//   		SupportedVersions: []MCPProtocolVersion{
//   			agentcore.MCPProtocolVersion_MCP_2025_03_26(),
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
//   	ExceptionLevel: agentcore.GatewayExceptionLevel_DEBUG(),
//   })
//
type McpGatewaySearchType interface {
	// The search type string value.
	Value() *string
	// Returns the string value.
	ToString() *string
}

// The jsii proxy struct for McpGatewaySearchType
type jsiiProxy_McpGatewaySearchType struct {
	_ byte // padding
}

func (j *jsiiProxy_McpGatewaySearchType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom search type not yet defined in this class.
func McpGatewaySearchType_Of(value *string) McpGatewaySearchType {
	_init_.Initialize()

	if err := validateMcpGatewaySearchType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns McpGatewaySearchType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.McpGatewaySearchType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func McpGatewaySearchType_SEMANTIC() McpGatewaySearchType {
	_init_.Initialize()
	var returns McpGatewaySearchType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.McpGatewaySearchType",
		"SEMANTIC",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_McpGatewaySearchType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

