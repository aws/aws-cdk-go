//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func validateNewActionParameters(props *ActionProps) error {
	return nil
}

