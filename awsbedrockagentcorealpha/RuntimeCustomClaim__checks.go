//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"
)

func validateRuntimeCustomClaim_WithStringArrayValueParameters(name *string, values *[]*string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateRuntimeCustomClaim_WithStringValueParameters(name *string, value *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

