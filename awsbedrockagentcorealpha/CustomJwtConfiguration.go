package awsbedrockagentcorealpha


// Custom JWT authorizer configuration.
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
type CustomJwtConfiguration struct {
	// This URL is used to fetch OpenID Connect configuration or authorization server metadata for validating incoming tokens.
	//
	// Pattern: .+/\.well-known/openid-configuration
	// Required: Yes.
	// Experimental.
	DiscoveryUrl *string `field:"required" json:"discoveryUrl" yaml:"discoveryUrl"`
	// Represents individual audience values that are validated in the incoming JWT token validation process.
	// Default: - No audience validation.
	//
	// Experimental.
	AllowedAudience *[]*string `field:"optional" json:"allowedAudience" yaml:"allowedAudience"`
	// Represents individual client IDs that are validated in the incoming JWT token validation process.
	// Default: - No client ID validation.
	//
	// Experimental.
	AllowedClients *[]*string `field:"optional" json:"allowedClients" yaml:"allowedClients"`
}

