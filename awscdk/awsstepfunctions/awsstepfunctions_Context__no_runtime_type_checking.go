//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateContext_NumberAtParameters(path *string) error {
	return nil
}

func validateContext_StringAtParameters(path *string) error {
	return nil
}

