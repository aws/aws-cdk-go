//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func validateSnapshotCredentials_FromGeneratedPasswordParameters(username *string, options *SnapshotCredentialsFromGeneratedPasswordOptions) error {
	return nil
}

func validateSnapshotCredentials_FromGeneratedSecretParameters(username *string, options *SnapshotCredentialsFromGeneratedPasswordOptions) error {
	return nil
}

func validateSnapshotCredentials_FromPasswordParameters(password awscdk.SecretValue) error {
	return nil
}

func validateSnapshotCredentials_FromSecretParameters(secret awssecretsmanager.ISecret) error {
	return nil
}

