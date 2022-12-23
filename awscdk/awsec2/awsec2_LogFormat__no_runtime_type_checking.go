//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateLogFormat_CustomParameters(formatString *string) error {
	return nil
}

func validateLogFormat_FieldParameters(field *string) error {
	return nil
}

func validateNewLogFormatParameters(value *string) error {
	return nil
}

