//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func validateValues_FromSearchParameters(expression *string, populateFrom *string) error {
	return nil
}

func validateValues_FromSearchComponentsParameters(components *SearchComponents) error {
	return nil
}

func validateValues_FromValuesParameters(values *[]*VariableValue) error {
	return nil
}

