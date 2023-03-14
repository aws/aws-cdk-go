//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func validateHttpParameter_FromSecretParameters(value awscdk.SecretValue) error {
	return nil
}

func validateHttpParameter_FromStringParameters(value *string) error {
	return nil
}

