//go:build no_runtime_type_checking

package awsglue

// Building without runtime type checking enabled, so all the below just return nil

func validateNewInputFormatParameters(className *string) error {
	return nil
}

