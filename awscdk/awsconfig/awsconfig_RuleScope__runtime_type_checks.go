//go:build !no_runtime_type_checking

package awsconfig

import (
	"fmt"
)

func validateRuleScope_FromResourceParameters(resourceType ResourceType) error {
	if resourceType == nil {
		return fmt.Errorf("parameter resourceType is required, but nil was provided")
	}

	return nil
}

func validateRuleScope_FromResourcesParameters(resourceTypes *[]ResourceType) error {
	if resourceTypes == nil {
		return fmt.Errorf("parameter resourceTypes is required, but nil was provided")
	}

	return nil
}

func validateRuleScope_FromTagParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

