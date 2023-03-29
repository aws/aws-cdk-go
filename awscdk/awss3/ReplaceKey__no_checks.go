//go:build no_runtime_type_checking

package awss3

// Building without runtime type checking enabled, so all the below just return nil

func validateReplaceKey_PrefixWithParameters(keyReplacement *string) error {
	return nil
}

func validateReplaceKey_WithParameters(keyReplacement *string) error {
	return nil
}

