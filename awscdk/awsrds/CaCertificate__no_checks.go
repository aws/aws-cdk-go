//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func validateCaCertificate_OfParameters(identifier *string) error {
	return nil
}

