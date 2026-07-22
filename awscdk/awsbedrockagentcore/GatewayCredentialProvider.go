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
//   // Outbound auth: ApiKeyCredentialProvider + bindForGatewayApiKeyTarget, or ARNs from console/API
//   apiKeyIdentityArn := "arn:aws:bedrock-agentcore:us-east-1:123456789012:token-vault/abc123/apikeycredentialprovider/my-apikey"
//   apiKeySecretArn := "arn:aws:secretsmanager:us-east-1:123456789012:secret:my-apikey-secret-abc123"
//
//   opneapiSchema := agentcore.ApiSchema_FromLocalAsset(path.join(__dirname, jsii.String("mySchema.yml")))
//   opneapiSchema.Bind(this)
//
//   // Create a gateway target with OpenAPI Schema
//   target := agentcore.GatewayTarget_ForOpenApi(this, jsii.String("MyTarget"), &GatewayTargetOpenApiProps{
//   	GatewayTargetName: jsii.String("my-api-target"),
//   	Description: jsii.String("Target for external API integration"),
//   	Gateway: gateway,
//   	 // Note: you need to pass the gateway reference
//   	ApiSchema: opneapiSchema,
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		agentcore.GatewayCredentialProvider_FromApiKeyIdentityArn(&ApiKeyCredentialProviderOptions{
//   			ProviderArn: apiKeyIdentityArn,
//   			SecretArn: apiKeySecretArn,
//   		}),
//   	},
//   })
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
// The gateway authenticates outbound requests using its own execution role (SigV4).
// Provide `service` and optionally `region` to explicitly choose the SigV4 signing
// service / region instead of relying on the gateway's inference from the target
// endpoint. Useful for cross-region calls and for targets where the service can't be
// inferred from the URL. Explicit `service` / `region` is only supported for MCP Server
// and OpenAPI targets; other target types must use the bare `fromIamRole()`.
func GatewayCredentialProvider_FromIamRole(props *GatewayIamRoleCredentialProviderProps) ICredentialProviderConfig {
	_init_.Initialize()

	if err := validateGatewayCredentialProvider_FromIamRoleParameters(props); err != nil {
		panic(err)
	}
	var returns ICredentialProviderConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.GatewayCredentialProvider",
		"fromIamRole",
		[]interface{}{props},
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

