//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateRuntimeManagementMode_ManualParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func validateNewRuntimeManagementModeParameters(mode *string) error {
	if mode == nil {
		return fmt.Errorf("parameter mode is required, but nil was provided")
	}

	return nil
}

