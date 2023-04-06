//go:build !no_runtime_type_checking

package awssecretsmanager

import (
	"fmt"
)

func validateSecretStringValueBeta1_FromTokenParameters(secretValueFromToken *string) error {
	if secretValueFromToken == nil {
		return fmt.Errorf("parameter secretValueFromToken is required, but nil was provided")
	}

	return nil
}

func validateSecretStringValueBeta1_FromUnsafePlaintextParameters(secretValue *string) error {
	if secretValue == nil {
		return fmt.Errorf("parameter secretValue is required, but nil was provided")
	}

	return nil
}

