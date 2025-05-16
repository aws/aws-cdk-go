//go:build no_runtime_type_checking

package awsbedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateNewFoundationModelIdentifierParameters(modelId *string) error {
	return nil
}

