//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func validateResponseType_OfParameters(type_ *string) error {
	return nil
}

