package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating different Gateway Credential Providers.
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
type GatewayCredentialProvider interface {
}

// The jsii proxy struct for GatewayCredentialProvider
type jsiiProxy_GatewayCredentialProvider struct {
	_ byte // padding
}

// Experimental.
func NewGatewayCredentialProvider_Override(g GatewayCredentialProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayCredentialProvider",
		nil, // no parameters
		g,
	)
}

// Create an API key credential provider from Identity ARN Use this method when you have the Identity ARN as a string.
//
// Returns: ICredentialProviderConfig configured for API key authentication.
// Experimental.
func GatewayCredentialProvider_FromApiKeyIdentityArn(props *ApiKeyCredentialProviderProps) ICredentialProviderConfig {
	_init_.Initialize()

	if err := validateGatewayCredentialProvider_FromApiKeyIdentityArnParameters(props); err != nil {
		panic(err)
	}
	var returns ICredentialProviderConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayCredentialProvider",
		"fromApiKeyIdentityArn",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an IAM role credential provider.
//
// Returns: IIamRoleCredentialProvider configured for IAM role authentication.
// Experimental.
func GatewayCredentialProvider_FromIamRole() ICredentialProviderConfig {
	_init_.Initialize()

	var returns ICredentialProviderConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayCredentialProvider",
		"fromIamRole",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create an OAuth credential provider from Identity ARN Use this method when you have the Identity ARN as a string.
//
// Returns: ICredentialProviderConfig configured for OAuth authentication.
// Experimental.
func GatewayCredentialProvider_FromOauthIdentityArn(props *OAuthConfiguration) ICredentialProviderConfig {
	_init_.Initialize()

	if err := validateGatewayCredentialProvider_FromOauthIdentityArnParameters(props); err != nil {
		panic(err)
	}
	var returns ICredentialProviderConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayCredentialProvider",
		"fromOauthIdentityArn",
		[]interface{}{props},
		&returns,
	)

	return returns
}

