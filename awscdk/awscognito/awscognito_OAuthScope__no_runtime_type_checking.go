//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func validateOAuthScope_CustomParameters(name *string) error {
	return nil
}

func validateOAuthScope_ResourceServerParameters(server IUserPoolResourceServer, scope ResourceServerScope) error {
	return nil
}

