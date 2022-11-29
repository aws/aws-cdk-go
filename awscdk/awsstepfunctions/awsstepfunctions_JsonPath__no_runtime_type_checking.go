//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateJsonPath_FormatParameters(formatString *string) error {
	return nil
}

func validateJsonPath_IsEncodedJsonPathParameters(value *string) error {
	return nil
}

func validateJsonPath_JsonToStringParameters(value interface{}) error {
	return nil
}

func validateJsonPath_ListAtParameters(path *string) error {
	return nil
}

func validateJsonPath_NumberAtParameters(path *string) error {
	return nil
}

func validateJsonPath_ObjectAtParameters(path *string) error {
	return nil
}

func validateJsonPath_StringAtParameters(path *string) error {
	return nil
}

func validateJsonPath_StringToJsonParameters(jsonString *string) error {
	return nil
}

