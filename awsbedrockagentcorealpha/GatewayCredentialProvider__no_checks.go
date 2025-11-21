//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateGatewayCredentialProvider_FromApiKeyIdentityArnParameters(props *ApiKeyCredentialProviderProps) error {
	return nil
}

func validateGatewayCredentialProvider_FromOauthIdentityArnParameters(props *OAuthConfiguration) error {
	return nil
}

