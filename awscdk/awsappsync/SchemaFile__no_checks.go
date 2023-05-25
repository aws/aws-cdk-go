//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SchemaFile) validateBindParameters(api IGraphqlApi, _options *SchemaBindOptions) error {
	return nil
}

func validateSchemaFile_FromAssetParameters(filePath *string) error {
	return nil
}

func (j *jsiiProxy_SchemaFile) validateSetDefinitionParameters(val *string) error {
	return nil
}

func validateNewSchemaFileParameters(options *SchemaProps) error {
	return nil
}

