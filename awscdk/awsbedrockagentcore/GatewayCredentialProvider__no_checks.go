//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validateGatewayCredentialProvider_FromApiKeyIdentityParameters(provider IApiKeyCredentialProvider, options *FromApiKeyIdentityOptions) error {
	return nil
}

func validateGatewayCredentialProvider_FromApiKeyIdentityArnParameters(props *ApiKeyCredentialProviderOptions) error {
	return nil
}

func validateGatewayCredentialProvider_FromOauthIdentityParameters(provider IOAuth2CredentialProvider, options *FromOauthIdentityOptions) error {
	return nil
}

func validateGatewayCredentialProvider_FromOauthIdentityArnParameters(props *OAuthConfiguration) error {
	return nil
}

