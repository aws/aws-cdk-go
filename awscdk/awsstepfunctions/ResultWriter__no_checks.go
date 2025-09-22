//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateNewResultWriterParameters(props *ResultWriterProps) error {
	return nil
}

