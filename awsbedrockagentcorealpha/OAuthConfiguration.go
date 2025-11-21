package awsbedrockagentcorealpha


// OAuth configuration.
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
type OAuthConfiguration struct {
	// The OAuth credential provider ARN.
	//
	// This is returned when creating the OAuth credential provider via Console or API.
	// Format: arn:aws:bedrock-agentcore:region:account:token-vault/id/oauth2credentialprovider/name
	// Required: Yes.
	// Experimental.
	ProviderArn *string `field:"required" json:"providerArn" yaml:"providerArn"`
	// The OAuth scopes for the credential provider. These scopes define the level of access requested from the OAuth provider.
	//
	// Array Members: Minimum number of 0 items. Maximum number of 100 items.
	// Length Constraints: Minimum length of 1. Maximum length of 64.
	// Required: Yes.
	// Experimental.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The ARN of the Secrets Manager secret containing OAuth credentials (client ID and secret).
	//
	// This is returned when creating the OAuth credential provider via Console or API.
	// Format: arn:aws:secretsmanager:region:account:secret:name
	// Required: Yes.
	// Experimental.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// Custom parameters for the OAuth flow.
	// Default: - No custom parameters.
	//
	// Experimental.
	CustomParameters *map[string]*string `field:"optional" json:"customParameters" yaml:"customParameters"`
}

