//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateProvideItems_JsonArrayParameters(array *[]interface{}) error {
	return nil
}

func validateProvideItems_JsonataParameters(jsonataExpression *string) error {
	return nil
}

