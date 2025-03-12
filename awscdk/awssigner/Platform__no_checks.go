//go:build no_runtime_type_checking

package awssigner

// Building without runtime type checking enabled, so all the below just return nil

func validatePlatform_OfParameters(platformId *string) error {
	return nil
}

