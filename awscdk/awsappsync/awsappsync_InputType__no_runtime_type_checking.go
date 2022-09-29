//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InputType) validateAddFieldParameters(options *AddFieldOptions) error {
	return nil
}

func (i *jsiiProxy_InputType) validateAttributeParameters(options *BaseTypeOptions) error {
	return nil
}

func validateNewInputTypeParameters(name *string, props *IntermediateTypeOptions) error {
	return nil
}

