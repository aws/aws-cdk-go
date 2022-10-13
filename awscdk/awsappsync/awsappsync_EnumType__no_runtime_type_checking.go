//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EnumType) validateAddFieldParameters(options *AddFieldOptions) error {
	return nil
}

func (e *jsiiProxy_EnumType) validateAttributeParameters(options *BaseTypeOptions) error {
	return nil
}

func validateNewEnumTypeParameters(name *string, options *EnumTypeOptions) error {
	return nil
}

