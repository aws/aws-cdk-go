//go:build no_runtime_type_checking

package awscdkcognitoidentitypoolalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateIdentityPoolProviderUrl_CustomParameters(url *string) error {
	return nil
}

func validateIdentityPoolProviderUrl_OpenIdParameters(url *string) error {
	return nil
}

func validateIdentityPoolProviderUrl_SamlParameters(url *string) error {
	return nil
}

func validateIdentityPoolProviderUrl_UserPoolParameters(userPool awscognito.UserPool, userPoolClient awscognito.UserPoolClient) error {
	return nil
}

func validateNewIdentityPoolProviderUrlParameters(type_ IdentityPoolProviderType, value *string) error {
	return nil
}

