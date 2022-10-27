//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObjectType) validateAddFieldParameters(options *AddFieldOptions) error {
	return nil
}

func (o *jsiiProxy_ObjectType) validateAttributeParameters(options *BaseTypeOptions) error {
	return nil
}

func (o *jsiiProxy_ObjectType) validateGenerateResolverParameters(api IGraphqlApi, fieldName *string, options *ResolvableFieldOptions) error {
	return nil
}

func validateNewObjectTypeParameters(name *string, props *ObjectTypeOptions) error {
	return nil
}

