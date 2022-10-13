//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateData_IsJsonPathStringParameters(value *string) error {
	return nil
}

func validateData_ListAtParameters(path *string) error {
	return nil
}

func validateData_NumberAtParameters(path *string) error {
	return nil
}

func validateData_StringAtParameters(path *string) error {
	return nil
}

