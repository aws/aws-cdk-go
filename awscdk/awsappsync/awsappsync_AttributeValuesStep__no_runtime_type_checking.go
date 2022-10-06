//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AttributeValuesStep) validateIsParameters(val *string) error {
	return nil
}

func validateNewAttributeValuesStepParameters(attr *string, container *string, assignments *[]Assign) error {
	return nil
}

