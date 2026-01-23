package awsbedrockagentcorealpha


// Properties for creating an MCP Server-based Gateway Target.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   // OAuth2 authentication (recommended)
//   // Note: Create the OAuth provider using AWS console or Identity L2 construct when available
//   oauthProviderArn := "arn:aws:bedrock-agentcore:us-east-1:123456789012:token-vault/abc123/oauth2credentialprovider/my-oauth"
//   oauthSecretArn := "arn:aws:secretsmanager:us-east-1:123456789012:secret:my-oauth-secret-abc123"
//
//   // Create a gateway target with MCP Server
//   mcpTarget := agentcore.GatewayTarget_ForMcpServer(this, jsii.String("MyMcpServer"), &GatewayTargetMcpServerProps{
//   	GatewayTargetName: jsii.String("my-mcp-server"),
//   	Description: jsii.String("External MCP server integration"),
//   	Gateway: gateway,
//   	Endpoint: jsii.String("https://my-mcp-server.example.com"),
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		agentcore.GatewayCredentialProvider_FromOauthIdentityArn(&OAuthConfiguration{
//   			ProviderArn: oauthProviderArn,
//   			SecretArn: oauthSecretArn,
//   			Scopes: []*string{
//   				jsii.String("mcp-runtime-server/invoke"),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type GatewayTargetMcpServerProps struct {
	// Optional description for the gateway target The description can have up to 200 characters.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target The name must be unique within the gateway Pattern: ^([0-9a-zA-Z][-]?){1,100}$.
	// Default: - auto generate.
	//
	// Experimental.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
	// Credential providers for authentication.
	// Experimental.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"required" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// The HTTPS endpoint URL of the MCP server.
	//
	// The endpoint must:
	// - Use HTTPS protocol
	// - Be properly URL-encoded
	// - Point to an MCP server that implements tool capabilities.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The gateway this target belongs to.
	// Experimental.
	Gateway IGateway `field:"required" json:"gateway" yaml:"gateway"`
}

