//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func validateComparablePrincipal_DedupeStringForParameters(x IPrincipal) error {
	return nil
}

func validateComparablePrincipal_IsComparablePrincipalParameters(x IPrincipal) error {
	return nil
}

