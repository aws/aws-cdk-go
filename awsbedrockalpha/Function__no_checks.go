//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateNewFunctionParameters(props *FunctionProps) error {
	return nil
}

