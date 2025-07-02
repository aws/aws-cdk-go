//go:build no_runtime_type_checking

package awscertificatemanager

// Building without runtime type checking enabled, so all the below just return nil

func validateNewKeyAlgorithmParameters(name *string) error {
	return nil
}

