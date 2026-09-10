//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"
)

func validatePolicyResource_AnyOfTypeParameters(entityType *string) error {
	if entityType == nil {
		return fmt.Errorf("parameter entityType is required, but nil was provided")
	}

	return nil
}

func validatePolicyResource_InstanceParameters(entityType *string, entityArn *string) error {
	if entityType == nil {
		return fmt.Errorf("parameter entityType is required, but nil was provided")
	}

	if entityArn == nil {
		return fmt.Errorf("parameter entityArn is required, but nil was provided")
	}

	return nil
}

