//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func validateDefaultEncryptedResourceFactories_GetParameters(type_ *string) error {
	return nil
}

func validateDefaultEncryptedResourceFactories_HasParameters(type_ *string) error {
	return nil
}

func validateDefaultEncryptedResourceFactories_SetParameters(type_ *string, factory IEncryptedResourceFactory) error {
	return nil
}

