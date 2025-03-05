//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_CustomerManagedEncryptionConfiguration) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewCustomerManagedEncryptionConfigurationParameters(kmsKey awskms.IKey) error {
	return nil
}

