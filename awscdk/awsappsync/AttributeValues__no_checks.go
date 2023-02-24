//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AttributeValues) validateAttributeParameters(attr *string) error {
	return nil
}

func validateNewAttributeValuesParameters(container *string) error {
	return nil
}

