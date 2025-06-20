//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateTaskInput_FromJsonPathAtParameters(path *string) error {
	return nil
}

func validateTaskInput_FromObjectParameters(obj *map[string]interface{}) error {
	return nil
}

func validateTaskInput_FromTextParameters(text *string) error {
	return nil
}

