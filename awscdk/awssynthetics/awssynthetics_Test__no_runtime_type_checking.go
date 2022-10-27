//go:build no_runtime_type_checking

package awssynthetics

// Building without runtime type checking enabled, so all the below just return nil

func validateTest_CustomParameters(options *CustomTestOptions) error {
	return nil
}

