//go:build no_runtime_type_checking

package customresources

// Building without runtime type checking enabled, so all the below just return nil

func validateNewLoggingParameters(props *LoggingProps) error {
	return nil
}

