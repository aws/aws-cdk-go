//go:build no_runtime_type_checking

package awscdkintegtestsalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateExpectedResult_ArrayWithParameters(expected *[]interface{}) error {
	return nil
}

func validateExpectedResult_ExactParameters(expected interface{}) error {
	return nil
}

func validateExpectedResult_ObjectLikeParameters(expected *map[string]interface{}) error {
	return nil
}

func validateExpectedResult_StringLikeRegexpParameters(expected *string) error {
	return nil
}

func (j *jsiiProxy_ExpectedResult) validateSetResultParameters(val *string) error {
	return nil
}

