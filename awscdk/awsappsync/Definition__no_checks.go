//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validateDefinition_FromFileParameters(filePath *string) error {
	return nil
}

func validateDefinition_FromSchemaParameters(schema ISchema) error {
	return nil
}

func validateDefinition_FromSourceApisParameters(sourceApiOptions *SourceApiOptions) error {
	return nil
}

