//go:build no_runtime_type_checking

package assertions

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Tags) validateHasValuesParameters(tags interface{}) error {
	return nil
}

func validateTags_FromStackParameters(stack awscdk.Stack) error {
	return nil
}

