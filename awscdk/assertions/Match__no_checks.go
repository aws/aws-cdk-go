//go:build no_runtime_type_checking

package assertions

// Building without runtime type checking enabled, so all the below just return nil

func validateMatch_ArrayEqualsParameters(pattern *[]interface{}) error {
	return nil
}

func validateMatch_ArrayWithParameters(pattern *[]interface{}) error {
	return nil
}

func validateMatch_ExactParameters(pattern interface{}) error {
	return nil
}

func validateMatch_NotParameters(pattern interface{}) error {
	return nil
}

func validateMatch_ObjectEqualsParameters(pattern *map[string]interface{}) error {
	return nil
}

func validateMatch_ObjectLikeParameters(pattern *map[string]interface{}) error {
	return nil
}

func validateMatch_SerializedJsonParameters(pattern interface{}) error {
	return nil
}

func validateMatch_StringLikeRegexpParameters(pattern *string) error {
	return nil
}

