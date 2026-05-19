package awsbedrockagentcore


// Options for adding an MCP Server target to a gateway.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   oauth := agentcore.OAuth2CredentialProvider_UsingGithub(this, jsii.String("GhOAuth"), &GithubOAuth2CredentialProviderProps{
//   	OAuth2CredentialProviderName: jsii.String("github-oauth"),
//   	ClientId: jsii.String("your-client-id"),
//   	ClientSecret: cdk.SecretValue_UnsafePlainText(jsii.String("your-client-secret")),
//   })
//
//   gateway.AddMcpServerTarget(jsii.String("Mcp"), &AddMcpServerTargetOptions{
//   	GatewayTargetName: jsii.String("mcp-server"),
//   	Description: jsii.String("MCP with GitHub OAuth"),
//   	Endpoint: jsii.String("https://my-mcp-server.example.com"),
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		agentcore.GatewayCredentialProvider_FromOauthIdentity(oauth, &FromOauthIdentityOptions{
//   			Scopes: []*string{
//   				jsii.String("read:user"),
//   			},
//   		}),
//   	},
//   })
//
type AddMcpServerTargetOptions struct {
	// Credential providers for authentication.
	//
	// MCP servers require explicit authentication configuration.
	// OAuth2 is strongly recommended over NoAuth.
	CredentialProviderConfigurations *[]ICredentialProviderConfig `field:"required" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// The HTTPS endpoint URL of the MCP server.
	//
	// The endpoint must:
	// - Use HTTPS protocol
	// - Be properly URL-encoded
	// - Point to an MCP server that implements tool capabilities.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Optional description for the gateway target.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target Valid characters are a-z, A-Z, 0-9, _ (underscore) and - (hyphen).
	// Default: - auto generate.
	//
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
}

