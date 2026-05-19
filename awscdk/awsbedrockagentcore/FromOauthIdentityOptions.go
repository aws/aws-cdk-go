package awsbedrockagentcore


// OAuth scopes (and optional custom parameters) when binding an {@link IOAuth2CredentialProvider} to a gateway target.
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
type FromOauthIdentityOptions struct {
	// OAuth scopes the gateway should request for this target.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Additional OAuth parameters for the provider.
	// Default: - none.
	//
	CustomParameters *map[string]*string `field:"optional" json:"customParameters" yaml:"customParameters"`
}

