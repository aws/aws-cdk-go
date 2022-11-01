//go:build no_runtime_type_checking

package awssecretsmanager

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretRotationApplication) validateApplicationArnForPartitionParameters(partition *string) error {
	return nil
}

func (s *jsiiProxy_SecretRotationApplication) validateSemanticVersionForPartitionParameters(partition *string) error {
	return nil
}

func validateNewSecretRotationApplicationParameters(applicationId *string, semanticVersion *string, options *SecretRotationApplicationOptions) error {
	return nil
}

