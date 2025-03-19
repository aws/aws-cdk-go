//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validateValues_AttributeParameters(attr *string) error {
	return nil
}

