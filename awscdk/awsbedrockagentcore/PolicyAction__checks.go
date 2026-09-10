//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"
)

func validatePolicyAction_AnyOfParameters(actions *[]*string) error {
	if actions == nil {
		return fmt.Errorf("parameter actions is required, but nil was provided")
	}

	return nil
}

func validatePolicyAction_OneParameters(action *string) error {
	if action == nil {
		return fmt.Errorf("parameter action is required, but nil was provided")
	}

	return nil
}

