//go:build no_runtime_type_checking

package awssecretsmanager

// Building without runtime type checking enabled, so all the below just return nil

func validateSecretStringValueBeta1_FromTokenParameters(secretValueFromToken *string) error {
	return nil
}

func validateSecretStringValueBeta1_FromUnsafePlaintextParameters(secretValue *string) error {
	return nil
}

