//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func validateFilterPattern_BooleanValueParameters(jsonField *string, value *bool) error {
	return nil
}

func validateFilterPattern_ExistsParameters(jsonField *string) error {
	return nil
}

func validateFilterPattern_IsNullParameters(jsonField *string) error {
	return nil
}

func validateFilterPattern_LiteralParameters(logPatternString *string) error {
	return nil
}

func validateFilterPattern_NotExistsParameters(jsonField *string) error {
	return nil
}

func validateFilterPattern_NumberValueParameters(jsonField *string, comparison *string, value *float64) error {
	return nil
}

func validateFilterPattern_StringValueParameters(jsonField *string, comparison *string, value *string) error {
	return nil
}

