package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating different Gateway Credential Providers.
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

// Create an API key outbound auth configuration from a Token Vault {@link IApiKeyCredentialProvider} construct.
//
// Prefer this over {@link GatewayCredentialProvider.fromApiKeyIdentityArn} when the provider is defined in CDK.
// Experimental.
func GatewayCredentialProvider_FromApiKeyIdentity(provider IApiKeyCredentialProvider, options *FromApiKeyIdentityOptions) ICredentialProviderConfig {
	_init_.Initialize()

	if err := validateGatewayCredentialProvider_FromApiKeyIdentityParameters(provider, options); err != nil {
		panic(err)
	}
	var returns ICredentialProviderConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayCredentialProvider",
		"fromApiKeyIdentity",
		[]interface{}{provider, options},
		&returns,
	)

	return returns
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

// Create an OAuth outbound auth configuration from a Token Vault {@link IOAuth2CredentialProvider} construct.
//
// Prefer this over {@link GatewayCredentialProvider.fromOauthIdentityArn} when the provider is defined in CDK.
// Experimental.
func GatewayCredentialProvider_FromOauthIdentity(provider IOAuth2CredentialProvider, options *FromOauthIdentityOptions) ICredentialProviderConfig {
	_init_.Initialize()

	if err := validateGatewayCredentialProvider_FromOauthIdentityParameters(provider, options); err != nil {
		panic(err)
	}
	var returns ICredentialProviderConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayCredentialProvider",
		"fromOauthIdentity",
		[]interface{}{provider, options},
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

