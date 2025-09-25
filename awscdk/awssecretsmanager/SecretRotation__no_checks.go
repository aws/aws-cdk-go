//go:build no_runtime_type_checking

package awssecretsmanager

// Building without runtime type checking enabled, so all the below just return nil

func validateSecretRotation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSecretRotationParameters(scope constructs.Construct, id *string, props *SecretRotationProps) error {
	return nil
}

