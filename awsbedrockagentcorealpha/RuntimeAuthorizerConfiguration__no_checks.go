//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateRuntimeAuthorizerConfiguration_UsingCognitoParameters(userPoolId *string, clientId *string) error {
	return nil
}

func validateRuntimeAuthorizerConfiguration_UsingJWTParameters(discoveryUrl *string) error {
	return nil
}

func validateRuntimeAuthorizerConfiguration_UsingOAuthParameters(discoveryUrl *string, clientId *string) error {
	return nil
}

