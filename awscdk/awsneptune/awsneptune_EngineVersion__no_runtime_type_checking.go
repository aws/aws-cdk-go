//go:build no_runtime_type_checking

package awsneptune

// Building without runtime type checking enabled, so all the below just return nil

func validateNewEngineVersionParameters(version *string) error {
	return nil
}

