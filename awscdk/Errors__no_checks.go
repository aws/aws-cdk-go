//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateErrors_IsAssertionErrorParameters(x interface{}) error {
	return nil
}

func validateErrors_IsCloudAssemblyErrorParameters(x interface{}) error {
	return nil
}

func validateErrors_IsConstructErrorParameters(x interface{}) error {
	return nil
}

func validateErrors_IsValidationErrorParameters(x interface{}) error {
	return nil
}

