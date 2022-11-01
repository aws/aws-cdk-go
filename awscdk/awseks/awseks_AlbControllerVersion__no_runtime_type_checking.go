//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func validateAlbControllerVersion_OfParameters(version *string) error {
	return nil
}

