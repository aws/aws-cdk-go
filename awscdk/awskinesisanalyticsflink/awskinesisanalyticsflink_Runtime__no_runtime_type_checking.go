//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awskinesisanalyticsflink

// Building without runtime type checking enabled, so all the below just return nil

func validateRuntime_OfParameters(value *string) error {
	return nil
}

