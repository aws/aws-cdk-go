//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateFilterValue_BooleanParameters(value *bool) error {
	return nil
}

func validateFilterValue_NumberParameters(value *float64) error {
	return nil
}

func validateFilterValue_StringParameters(value *string) error {
	return nil
}

