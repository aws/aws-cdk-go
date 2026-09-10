//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"
)

func validatePolicyPrincipal_EntityParameters(entityType *string, entityId *string) error {
	if entityType == nil {
		return fmt.Errorf("parameter entityType is required, but nil was provided")
	}

	if entityId == nil {
		return fmt.Errorf("parameter entityId is required, but nil was provided")
	}

	return nil
}

func validatePolicyPrincipal_EntityTypeParameters(entityType *string) error {
	if entityType == nil {
		return fmt.Errorf("parameter entityType is required, but nil was provided")
	}

	return nil
}

func validatePolicyPrincipal_InGroupParameters(groupType *string, groupId *string) error {
	if groupType == nil {
		return fmt.Errorf("parameter groupType is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

