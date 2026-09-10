//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"
)

func validatePolicyAttribute_ContextParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func validatePolicyAttribute_PrincipalParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

func validatePolicyAttribute_ResourceParameters(attribute *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	return nil
}

