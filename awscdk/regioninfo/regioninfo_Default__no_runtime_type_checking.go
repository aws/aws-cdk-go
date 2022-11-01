//go:build no_runtime_type_checking

package regioninfo

// Building without runtime type checking enabled, so all the below just return nil

func validateDefault_ServicePrincipalParameters(serviceFqn *string, region *string, urlSuffix *string) error {
	return nil
}

