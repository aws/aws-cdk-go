//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateResult_FromArrayParameters(value *[]interface{}) error {
	return nil
}

func validateResult_FromBooleanParameters(value *bool) error {
	return nil
}

func validateResult_FromNumberParameters(value *float64) error {
	return nil
}

func validateResult_FromObjectParameters(value *map[string]interface{}) error {
	return nil
}

func validateResult_FromStringParameters(value *string) error {
	return nil
}

func validateNewResultParameters(value interface{}) error {
	return nil
}

