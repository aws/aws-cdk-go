//go:build no_runtime_type_checking

package awsecr

// Building without runtime type checking enabled, so all the below just return nil

func validateNewRepositoryEncryptionParameters(value *string) error {
	return nil
}

