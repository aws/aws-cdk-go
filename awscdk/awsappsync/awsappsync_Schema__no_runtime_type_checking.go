//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Schema) validateAddMutationParameters(fieldName *string, field ResolvableField) error {
	return nil
}

func (s *jsiiProxy_Schema) validateAddQueryParameters(fieldName *string, field ResolvableField) error {
	return nil
}

func (s *jsiiProxy_Schema) validateAddSubscriptionParameters(fieldName *string, field Field) error {
	return nil
}

func (s *jsiiProxy_Schema) validateAddToSchemaParameters(addition *string) error {
	return nil
}

func (s *jsiiProxy_Schema) validateAddTypeParameters(type_ IIntermediateType) error {
	return nil
}

func (s *jsiiProxy_Schema) validateBindParameters(api GraphqlApi) error {
	return nil
}

func validateSchema_FromAssetParameters(filePath *string) error {
	return nil
}

func (j *jsiiProxy_Schema) validateSetDefinitionParameters(val *string) error {
	return nil
}

func validateNewSchemaParameters(options *SchemaOptions) error {
	return nil
}

