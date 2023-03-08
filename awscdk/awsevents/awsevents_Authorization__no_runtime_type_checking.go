//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func validateAuthorization_ApiKeyParameters(apiKeyName *string, apiKeyValue awscdk.SecretValue) error {
	return nil
}

func validateAuthorization_BasicParameters(username *string, password awscdk.SecretValue) error {
	return nil
}

func validateAuthorization_OauthParameters(props *OAuthAuthorizationProps) error {
	return nil
}

