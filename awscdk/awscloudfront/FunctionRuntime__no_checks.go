//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func validateFunctionRuntime_CustomParameters(runtimeString *string) error {
	return nil
}

