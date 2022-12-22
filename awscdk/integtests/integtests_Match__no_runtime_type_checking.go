//go:build no_runtime_type_checking

package integtests

// Building without runtime type checking enabled, so all the below just return nil

func validateMatch_ArrayWithParameters(pattern *[]interface{}) error {
	return nil
}

func validateMatch_ObjectLikeParameters(pattern *map[string]interface{}) error {
	return nil
}

func validateMatch_StringLikeRegexpParameters(pattern *string) error {
	return nil
}

