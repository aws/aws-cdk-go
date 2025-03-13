//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func validateCredentials_FromGeneratedSecretParameters(username *string, options *CredentialsBaseOptions) error {
	return nil
}

func validateCredentials_FromPasswordParameters(username *string, password awscdk.SecretValue) error {
	return nil
}

func validateCredentials_FromSecretParameters(secret awssecretsmanager.ISecret) error {
	return nil
}

func validateCredentials_FromUsernameParameters(username *string, options *CredentialsFromUsernameOptions) error {
	return nil
}

