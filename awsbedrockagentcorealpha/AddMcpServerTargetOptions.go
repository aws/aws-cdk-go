package awsbedrockagentcorealpha


// Options for adding an MCP Server target to a gateway.
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
//   // Add an MCP server target directly to the gateway
//   mcpTarget := gateway.AddMcpServerTarget(jsii.String("MyMcpServer"), &AddMcpServerTargetOptions{
//   	GatewayTargetName: jsii.String("my-mcp-server"),
//   	Description: jsii.String("External MCP server integration"),
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
//   // Grant sync permission to a Lambda function that will trigger synchronization
//   syncFunction := lambda.NewFunction(this, jsii.String("SyncFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_12(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String(`
//   	import boto3
//
//   	def handler(event, context):
//   	    client = boto3.client('bedrock-agentcore')
//   	    response = client.synchronize_gateway_targets(
//   	        gatewayIdentifier=event['gatewayId'],
//   	        targetIds=[event['targetId']]
//   	    )
//   	    return response
//   	  `)),
//   })
//
//   mcpTarget.GrantSync(syncFunction)
//
// Experimental.
type AddMcpServerTargetOptions struct {
	// Credential providers for authentication.
	//
	// MCP servers require explicit authentication configuration.
	// OAuth2 is strongly recommended over NoAuth.
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
	// Optional description for the gateway target.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the gateway target Valid characters are a-z, A-Z, 0-9, _ (underscore) and - (hyphen).
	// Default: - auto generate.
	//
	// Experimental.
	GatewayTargetName *string `field:"optional" json:"gatewayTargetName" yaml:"gatewayTargetName"`
}

