package awsbedrockagentcorealpha


// Custom JWT authorizer configuration.
//
// Example:
//   // Create a custom execution role
//   executionRole := iam.NewRole(this, jsii.String("GatewayExecutionRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("bedrock-agentcore.amazonaws.com")),
//   	ManagedPolicies: []IManagedPolicy{
//   		iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AmazonBedrockAgentCoreGatewayExecutionRolePolicy")),
//   	},
//   })
//
//   // Create gateway with custom execution role
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   	Description: jsii.String("Gateway with custom execution role"),
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
//   	Role: executionRole,
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
	// Represents individual scopes that are validated in the incoming JWT token validation process.
	// Default: - No scope validation.
	//
	// Experimental.
	AllowedScopes *[]*string `field:"optional" json:"allowedScopes" yaml:"allowedScopes"`
	// Custom claims for additional JWT token validation.
	//
	// Allows you to validate additional fields in JWT tokens beyond the standard audience, client, and scope validations.
	// Default: - No custom claim validation.
	//
	// Experimental.
	CustomClaims *[]GatewayCustomClaim `field:"optional" json:"customClaims" yaml:"customClaims"`
}

