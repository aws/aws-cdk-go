//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validateNewFunctionRuntimeParameters(family FunctionRuntimeFamily, version *string) error {
	return nil
}

