//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UnionType) validateAddFieldParameters(options *AddFieldOptions) error {
	return nil
}

func (u *jsiiProxy_UnionType) validateAttributeParameters(options *BaseTypeOptions) error {
	return nil
}

func validateNewUnionTypeParameters(name *string, options *UnionTypeOptions) error {
	return nil
}

