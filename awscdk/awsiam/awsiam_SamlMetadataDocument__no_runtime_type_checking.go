//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func validateSamlMetadataDocument_FromFileParameters(path *string) error {
	return nil
}

func validateSamlMetadataDocument_FromXmlParameters(xml *string) error {
	return nil
}

