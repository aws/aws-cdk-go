//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InterfaceType) validateAddFieldParameters(options *AddFieldOptions) error {
	return nil
}

func (i *jsiiProxy_InterfaceType) validateAttributeParameters(options *BaseTypeOptions) error {
	return nil
}

func validateNewInterfaceTypeParameters(name *string, props *IntermediateTypeOptions) error {
	return nil
}

