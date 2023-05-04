//go:build no_runtime_type_checking

package awsconfig

// Building without runtime type checking enabled, so all the below just return nil

func validateResourceType_OfParameters(type_ *string) error {
	return nil
}

