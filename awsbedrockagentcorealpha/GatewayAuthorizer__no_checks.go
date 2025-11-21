//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateGatewayAuthorizer_UsingCognitoParameters(props *CognitoAuthorizerProps) error {
	return nil
}

func validateGatewayAuthorizer_UsingCustomJwtParameters(configuration *CustomJwtConfiguration) error {
	return nil
}

