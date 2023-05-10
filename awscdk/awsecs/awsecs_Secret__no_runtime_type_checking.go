//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Secret) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateSecret_FromSecretsManagerParameters(secret awssecretsmanager.ISecret) error {
	return nil
}

func validateSecret_FromSecretsManagerVersionParameters(secret awssecretsmanager.ISecret, versionInfo *SecretVersionInfo) error {
	return nil
}

func validateSecret_FromSsmParameterParameters(parameter awsssm.IParameter) error {
	return nil
}

