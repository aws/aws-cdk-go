//go:build !no_runtime_type_checking

package awssns

import (
	"fmt"
)

func validatePolicy_FilterParameters(filter SubscriptionFilter) error {
	if filter == nil {
		return fmt.Errorf("parameter filter is required, but nil was provided")
	}

	return nil
}

func validatePolicy_PolicyParameters(policy *map[string]FilterOrPolicy) error {
	if policy == nil {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func validateNewPolicyParameters(policyDoc *map[string]FilterOrPolicy) error {
	if policyDoc == nil {
		return fmt.Errorf("parameter policyDoc is required, but nil was provided")
	}

	return nil
}

