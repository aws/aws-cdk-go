//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func validateLoggingFormat_FromJsonParameters(jsonLoggingFormat *map[string]*string) error {
	return nil
}

func validateLoggingFormat_FromTextParameters(text *string) error {
	return nil
}

