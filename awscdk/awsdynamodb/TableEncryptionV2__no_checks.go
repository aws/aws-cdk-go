//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func validateTableEncryptionV2_CustomerManagedKeyParameters(tableKey awskms.IKey) error {
	return nil
}

