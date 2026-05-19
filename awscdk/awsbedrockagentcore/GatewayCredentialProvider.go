package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating different Gateway Credential Providers.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   // OAuth2 (recommended): use OAuth2CredentialProvider + bindForGatewayOAuthTarget, or ARNs from console/API
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
type GatewayCredentialProvider interface {
}

// The jsii proxy struct for GatewayCredentialProvider
type jsiiProxy_GatewayCredentialProvider struct {
	_ byte // padding
}

func NewGatewayCredentialProvider_Override(g GatewayCredentialProvider) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayCredentialProvider",
		nil, // no parameters
		g,
	)
}

// Create an API key outbound auth configuration from a Token Vault {@link IApiKeyCredentialProvider} construct.
//
// Prefer this over {@link GatewayCredentialProvider.fromApiKeyIdentityArn} when the provider is defined in CDK.
func GatewayCredentialProvider_FromApiKeyIdentity(provider IApiKeyCredentialProvider, options *FromApiKeyIdentityOptions) ICredentialProviderConfig {
	_init_.Initialize()

	if err := validateGatewayCredentialProvider_FromApiKeyIdentityParameters(provider, options); err != nil {
		panic(err)
	}
	var returns ICredentialProviderConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayCredentialProvider",
		"fromApiKeyIdentity",
		[]interface{}{provider, options},
		&returns,
	)

	return returns
}

// Create an API key credential provider from Identity ARN Use this method when you have the Identity ARN as a string.
//
// Returns: ICredentialProviderConfig configured for API key authentication.
func GatewayCredentialProvider_FromApiKeyIdentityArn(props *ApiKeyCredentialProviderOptions) ICredentialProviderConfig {
	_init_.Initialize()

	if err := validateGatewayCredentialProvider_FromApiKeyIdentityArnParameters(props); err != nil {
		panic(err)
	}
	var returns ICredentialProviderConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayCredentialProvider",
		"fromApiKeyIdentityArn",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an IAM role credential provider.
//
// Returns: IIamRoleCredentialProvider configured for IAM role authentication.
func GatewayCredentialProvider_FromIamRole() ICredentialProviderConfig {
	_init_.Initialize()

	var returns ICredentialProviderConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayCredentialProvider",
		"fromIamRole",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create an OAuth outbound auth configuration from a Token Vault {@link IOAuth2CredentialProvider} construct.
//
// Prefer this over {@link GatewayCredentialProvider.fromOauthIdentityArn} when the provider is defined in CDK.
func GatewayCredentialProvider_FromOauthIdentity(provider IOAuth2CredentialProvider, options *FromOauthIdentityOptions) ICredentialProviderConfig {
	_init_.Initialize()

	if err := validateGatewayCredentialProvider_FromOauthIdentityParameters(provider, options); err != nil {
		panic(err)
	}
	var returns ICredentialProviderConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayCredentialProvider",
		"fromOauthIdentity",
		[]interface{}{provider, options},
		&returns,
	)

	return returns
}

// Create an OAuth credential provider from Identity ARN Use this method when you have the Identity ARN as a string.
//
// Returns: ICredentialProviderConfig configured for OAuth authentication.
func GatewayCredentialProvider_FromOauthIdentityArn(props *OAuthConfiguration) ICredentialProviderConfig {
	_init_.Initialize()

	if err := validateGatewayCredentialProvider_FromOauthIdentityArnParameters(props); err != nil {
		panic(err)
	}
	var returns ICredentialProviderConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayCredentialProvider",
		"fromOauthIdentityArn",
		[]interface{}{props},
		&returns,
	)

	return returns
}

