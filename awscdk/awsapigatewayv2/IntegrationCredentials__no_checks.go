//go:build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func validateIntegrationCredentials_FromRoleParameters(role awsiam.IRoleRef) error {
	return nil
}

